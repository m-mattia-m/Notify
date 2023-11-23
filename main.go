package main

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	sentrylogrus "github.com/getsentry/sentry-go/logrus"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"notify/internal/api"
	"notify/internal/service"
	"os"
	"time"
)

// @title Notify-API
// @version 1.0
// @description This is the API for Notify. With Notify you can securely send messages from the frontend to your chosen provider to send messages.
// @termsOfService  https://makenotify.io/terms-of-use/

// @contact.name API Support
// @contact.email develop@makenotify.io

// @schemes https
// @host      api.makenotify.io
// @BasePath  /v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @externalDocs.description  Docs
// @externalDocs.url          https://docs.makenotify.io

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Info(".env-file was not found. If you set the env-vars not in a .env-file, you don't have to do anything.")
	}

	err = initConfig()
	if err != nil {
		log.Fatalln("Error reading config file, %s", err)
	}

	checkIfRequiredConfigurationAttributesSet()

	err = initLogger()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	svc := service.NewClient()

	router := api.Router(svc)
	err = router.Run(fmt.Sprintf(":%s", viper.GetString("server.port")))
	if err != nil {
		log.Fatalf("Error starting api: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func initLogger() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(log.StandardLogger().Out)
	if viper.GetBool("logging.enable.debug") {
		log.SetLevel(log.DebugLevel)
	}

	if !viper.GetBool("logging.enable.sentry") {
		return nil
	}

	sentryLevels := []log.Level{log.ErrorLevel, log.FatalLevel, log.PanicLevel, log.InfoLevel}

	sentryLoggingDns, found := os.LookupEnv("SENTRY_LOGGING_DNS")
	if !found && viper.GetBool("logging.enable.sentry") {
		log.Fatal("Error during initialization the logging-client: env 'SENTRY_LOGGING_DNS' not found")
	}

	sentryHook, err := sentrylogrus.New(sentryLevels, sentry.ClientOptions{
		Dsn: sentryLoggingDns,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			event.Environment = viper.GetString("app.env")

			if hint.Context != nil {
				if req, ok := hint.Context.Value(sentry.RequestContextKey).(*http.Request); ok {
					fmt.Println(req)
				}
			}
			if viper.GetBool("logging.enable.console") {
				fmt.Println(
					struct {
						timestamp string
						eventId   string
						message   string
					}{
						timestamp: event.Timestamp.String(),
						eventId:   fmt.Sprintf("%v", event.EventID),
						message:   event.Message,
					})
			}
			return event
		},
		Debug:            viper.GetString("app.env") == "PROD",
		AttachStacktrace: true,
	})
	if err != nil {
		return fmt.Errorf("failed to create sentry-logrus-hook: %v", err)
	}
	defer sentryHook.Flush(5 * time.Second)
	log.AddHook(sentryHook)

	log.RegisterExitHandler(func() { sentryHook.Flush(5 * time.Second) })
	return nil
}

func checkIfRequiredConfigurationAttributesSet() {
	if viper.GetString("app.name") == "" {
		log.Fatal("failed to get required config attribute: 'app.name'")
	}
	if viper.GetString("app.env") == "" {
		log.Fatal("failed to get required config attribute: 'app.env'")
	}
	if viper.GetString("server.scheme") == "" {
		log.Fatal("failed to get required config attribute: 'server.scheme'")
	}
	if viper.GetString("server.domain") == "" {
		log.Fatal("failed to get required config attribute: 'server.domain'")
	}
	if viper.GetString("server.port") == "" {
		log.Fatal("failed to get required config attribute: 'server.port'")
	}
	if viper.GetString("server.version") == "" {
		log.Fatal("failed to get required config attribute: 'server.version'")
	}
	//if viper.GetString("logging.enable.console") == "" {
	//	log.Fatal("failed to get required config attribute: 'logging.enable.console'")
	//}
	//if viper.GetString("logging.enable.sentry") == "" {
	//	log.Fatal("failed to get required config attribute: 'logging.enable.sentry'")
	//}
	if viper.GetString("authentication.oidc.issuer") == "" {
		log.Fatal("failed to get required config attribute: 'authentication.oidc.issuer'")
	}
	if viper.GetString("authentication.oidc.clientId") == "" {
		log.Fatal("failed to get required config attribute: 'authentication.oidc.clientId'")
	}
	if viper.GetString("frontend.url") == "" {
		log.Fatal("failed to get required config attribute: 'frontend.url'")
	}
	//if viper.GetString("domain.activity.enable.subject") == "" {
	//	log.Fatal("failed to get required config attribute: 'domain.activity.enable.subject'")
	//}
	//if viper.GetString("domain.activity.enable.message") == "" {
	//	log.Fatal("failed to get required config attribute: 'domain.activity.enable.message'")
	//}
	//if viper.GetString("domain.swagger.port") == "" {
	//	log.Fatal("failed to get required config attribute: 'failed to get required config attribute'")
	//}
}
