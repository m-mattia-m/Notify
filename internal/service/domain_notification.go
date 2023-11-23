package service

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"notify/internal/model"
	"strings"
	"time"
)

const (
	MESSAGE_REPLACE_KEY = "{{message}}"
	SUBJECT_REPLACE_KEY = "{{subject}}"
)

func (c *Client) SendNotification(host string, notification model.Notification) (*model.SuccessMessage, error) {
	err := validateNotificationRequest(notification)
	if err != nil {
		return nil, err
	}

	hosts, err := c.db.ListHosts(model.Host{ProjectId: notification.ProjectId})
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("failed to verify host")
	}

	if index := ifHostInHosts(host, hosts); hosts == nil || index == -1 || !hosts[index].Verified {
		log.Error(err)
		return nil, fmt.Errorf("failed to verify host")
	}

	flows, err := c.db.ListFlows(model.Flow{ProjectId: notification.ProjectId})
	if err != nil {
		log.Error(err)
		return nil, fmt.Errorf("failed to list flows")
	}

	if flows == nil || len(flows) == 0 {
		return nil, fmt.Errorf("failed to list flows")
	}

	for _, flow := range flows {
		if !flow.Active {
			continue
		}

		var notificationError error
		var sourceTypeErr error
		switch flow.SourceType {
		case "slack":
			credential, err := c.db.GetSlackRevealedCredential(model.SlackCredentials{ProjectId: notification.ProjectId})
			if err != nil {
				notificationError = err
				log.Error(notificationError)
				break
			}

			err = c.sendSlackNotification(*flow, notification, *credential)
			if err != nil {
				notificationError = err
				log.Error(notificationError)
			}
			break
		case "mailgun":
			credential, err := c.db.GetMailgunRevealedCredential(model.MailgunCredentials{ProjectId: notification.ProjectId})
			if err != nil {
				notificationError = err
				log.Error(notificationError)
				break
			}

			err = c.sendMailgunNotification(*flow, notification, *credential)
			if err != nil {
				notificationError = err
				log.Error(notificationError)
				break
			}
			break
		default:
			sourceTypeErr = fmt.Errorf("invalid source-type")
		}

		err = c.logNotificationActivity(model.Activity{
			ProjectId:  notification.ProjectId,
			State:      validateActivityState(notificationError, sourceTypeErr),
			SourceType: flow.SourceType,
			Target:     validateTarget(*flow, notification),
			Subject:    validateActivitySubject(notification.Subject),
			Message:    validateActivityMessage(notification.Message),
			Note:       validateActivityNote(notificationError, sourceTypeErr),
		})
		if err != nil {
			log.Panic(err)
		}

		if err != nil && (notificationError != nil || sourceTypeErr != nil) {
			sourceTypeErrSuffix := ""
			if sourceTypeErr != nil {
				sourceTypeErrSuffix = fmt.Sprintf(": %s", sourceTypeErr.Error())
			}

			return nil, fmt.Errorf("failed to send message%s", sourceTypeErrSuffix)
		}

	}

	return &model.SuccessMessage{
		Message: "successfully sent",
	}, nil

}

func (c *Client) sendSlackNotification(flow model.Flow, notification model.Notification, credential model.SlackCredentials) error {
	target := validateTarget(flow, notification)
	message := strings.ReplaceAll(flow.MessageTemplate, MESSAGE_REPLACE_KEY, notification.Message)
	message = strings.ReplaceAll(message, SUBJECT_REPLACE_KEY, notification.Subject)
	preparedTarget := getSlackTarget(target, flow.Target)

	var errors string
	for _, receiver := range strings.Split(preparedTarget, ";") {
		err := c.slack.SendMessage(receiver, message, credential.BotUserOAuthToken)
		if err != nil {
			errors = fmt.Sprintf("%s;;%s", errors, err)
		}
	}

	if errors != "" {
		return fmt.Errorf(errors)
	}
	return nil
}

func (c *Client) sendMailgunNotification(flow model.Flow, notification model.Notification, credential model.MailgunCredentials) error {
	target := validateTarget(flow, notification)
	message := strings.ReplaceAll(flow.MessageTemplate, MESSAGE_REPLACE_KEY, notification.Message)
	message = strings.ReplaceAll(message, SUBJECT_REPLACE_KEY, notification.Subject)
	preparedTarget := getMailgunTarget(target, flow.Target)

	mg := mailgun.NewMailgun(credential.Domain, credential.ApiKey)

	switch credential.ApiBase {
	case "eu":
		mg.SetAPIBase(mailgun.APIBaseEU)
		break
	case "us":
		mg.SetAPIBase(mailgun.APIBaseUS)
		break
	}

	var errors string
	for _, receiver := range strings.Split(preparedTarget, ";") {
		mailgunMessage := mg.NewMessage(
			fmt.Sprintf("%s <%s>", credential.SenderName, credential.SenderEmail),
			notification.Subject,
			message,
			receiver,
		)
		if flow.MessageTemplateType == "text/html" {
			mailgunMessage = mg.NewMessage(
				fmt.Sprintf("%s <%s>", credential.SenderName, credential.SenderEmail),
				notification.Subject,
				"",
				receiver,
			)
			mailgunMessage.SetHtml(message)
		}
		mailgunMessage.SetReplyTo(credential.ReplyToEmail)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		_, _, err := mg.Send(ctx, mailgunMessage)
		if err != nil {
			errors = fmt.Sprintf("%s;;%s", errors, err)
		}
	}

	if errors != "" {
		return fmt.Errorf(errors)
	}
	return nil
}

func (c *Client) logNotificationActivity(activity model.Activity) error {
	err := c.db.CreateActivity(activity)
	return err
}

func ifHostInHosts(host string, hosts []*model.Host) int {
	for i, currentHost := range hosts {
		if currentHost.Host == host {
			return i
		}
	}
	return -1
}

func validateActivityState(notificationError, sourceTypeErr error) string {
	if notificationError == nil && sourceTypeErr == nil {
		return "success"
	}
	return "failed"
}

func validateTarget(flow model.Flow, notification model.Notification) string {
	if flow.OverrideTarget {
		if notification.Target != "" {
			return notification.Target
		}
	}
	return fmt.Sprintf("%s:%s", flow.SourceType, flow.Target)
}

func validateActivityNote(notificationError, sourceTypeErr error) string {
	var note string
	if notificationError != nil {
		note = fmt.Sprintf("%s \n NotificationError %s", note, notificationError)
	}
	if sourceTypeErr != nil {
		note = fmt.Sprintf("%s \n SourceTypeError %s", note, sourceTypeErr)
	}
	return note
}

func validateActivitySubject(subject string) string {
	if viper.GetBool("domain.activity.enable.subject") {
		return subject
	}
	return "config: subject in activity disabled"
}

func validateActivityMessage(subject string) string {
	if viper.GetBool("domain.activity.enable.message") {
		return subject
	}
	return "config: message in activity disabled"
}

func getSlackTarget(target, fallback string) string {
	cleanedTarget := strings.ReplaceAll(target, " ", "")
	targetParts := strings.Split(cleanedTarget, ";")

	var receiver string
	for _, targetPart := range targetParts {
		if strings.HasPrefix(targetPart, "slack:") {
			receiver = fmt.Sprintf("%s%s;", receiver, strings.Replace(targetPart, "slack:", "", 1))
		}
	}

	if strings.HasSuffix(receiver, ";") {
		receiver = receiver[:len(receiver)-1]
	}

	if receiver == "" {
		receiver = fallback
	}
	return receiver
}

func getMailgunTarget(target, fallback string) string {
	cleanedTarget := strings.ReplaceAll(target, " ", "")
	targetParts := strings.Split(cleanedTarget, ";")

	var receiver string
	for _, targetPart := range targetParts {
		if strings.HasPrefix(targetPart, "mailgun:") {
			receiver = fmt.Sprintf("%s%s;", receiver, strings.Replace(targetPart, "mailgun:", "", 1))
		}
	}

	if strings.HasSuffix(receiver, ";") {
		receiver = receiver[:len(receiver)-1]
	}

	if receiver == "" {
		receiver = fallback
	}
	return receiver
}

func validateNotificationRequest(notification model.Notification) error {
	if notification.ProjectId == "" {
		return fmt.Errorf("projectId is a required attribute")
	}
	if notification.Subject == "" {
		return fmt.Errorf("subject is a required attribute")
	}
	if notification.Message == "" {
		return fmt.Errorf("message is a required attribute")
	}
	return nil
}
