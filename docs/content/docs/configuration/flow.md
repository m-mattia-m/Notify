---
title: Flow
weight: 104
prev: docs/getting-started
next: docs/configuration/application
sidebar:
  open: true
---

{{< callout emoji="ℹ️" type="info" >}}
At the moment you only can create resources via swagger (or cURL, ...). We are working on Frontend and plan to release
it the beginning of 2024.
{{< /callout >}}

## Introduction

A flow defines where a message is sent to. You can add multiple flows for a project. You also can create a HTML-template
for Mailgun which are sent or only a simple plain text template. A standard recipient can also be defined which can also
be sent with permission in the notification request.

### Template types

You can configure this template types.

- `text/plain`
- `text/html`

Note that you also can add the type `text/html` to slack, but it is only rendered in the mail providers e.g. Mailgun.

### Message template

You can only replace `{{subject}}` and `{{message}}` in the template. Don't forget it, because we don't check if this in
the template. If you don't want a customized message, you only need to past: `{{messge}}` in the template attribute.

**text/plain**

```text
Hi Notify team\n\nYou have received a message about '{{subject}}' from your contact form: \n\n{{message}
```

**text/html**

```html
<h3>Hi Notify team</h3><p>You have received a message about <b>'{{subject}}'</b> from your contact form:</p><p>
    {{message}</p>
```

## Create

You can create a flow in your project. Replace `<project-id>` in the URL with your project id.

POST `/v1/settings/projects/<project-id>/flows`

```json lines {filename="body"}
{
  // defines if a flow is action and can be triggered
  "active": true,
  // here you can place your
  "message_template": "Hi,\n\nYou have received a message about '{{subject}}' from your contact form: \n\n{{message}",
  // defines the type of the message: text/html, text/plain
  "message_template_type": "text/plain",
  // you need to set a name for your flow. e.g. team-slack-notification
  "name": "Contact form notification",
  // defines if the message can override your default target
  "override_target": false,
  // defines where you want to be notified. e.g. slack, mailgun
  "source_type": "mailgun",
  // here you can past your default receiver email for mailgun or your Slack channel id 
  // (add the integration to your channel before connection Notify with it)
  "target": "you@example.com"
}
```

## Update

You can update a existing flow in your project. Replace `<project-id>` in the URL with your project id and `flow-id`
with your flow-id.

PUT `/v1/settings/projects/<project-id>/flows/<flow-id>`

```json lines {filename="body"}
{
  // defines if a flow is action and can be triggered
  "active": true,
  // here you can place your
  "message_template": "Hi,\n\nYou have received a message about '{{subject}}' from your contact form: \n\n{{message}",
  // defines the type of the message: text/html, text/plain
  "message_template_type": "text/plain",
  // you need to set a name for your flow. e.g. team-slack-notification
  "name": "Contact form notification",
  // defines if the message can override your default target
  "override_target": false,
  // defines where you want to be notified. e.g. slack, mailgun
  "source_type": "mailgun",
  // here you can past your default receiver email for mailgun or your Slack channel id 
  // (add the integration to your channel before connection Notify with it)
  "target": "you@example.com"
}
```

## List

You can list all your flows in a project. Replace `<project-id>` in the URL with your project id.

GET `/v1/settings/projects/<project-id>/flows`

## Get

You can get a specific flow in a project by their ID. Replace `<project-id>` in the URL with your project id
and `flow-id`
with your flow-id.

GET `/v1/settings/projects/<project-id>/flows/<flow-id>`

## Delete

You can delete a specific flow in a project by their ID. Replace `<project-id>` in the URL with your project id
and `flow-id`
with your flow-id.

DELETE `/v1/settings/projects/<project-id>/flows/<flow-id>`
