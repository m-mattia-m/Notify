---
title: Logging
weight: 305
prev: docs/configuration/server
next: docs/configuration/authentication
sidebar:
  open: true
---

```yaml {filename="./configs/config.yaml"}
logging:
  enable:
    console: true # defines if the errors should be logged in the console.
    sentry: false # defines if the errors should be logged to sentry.
```

## Sentry

For the logging in Notify we use the [Sentry](https://sentry.io/)-SDK. If you don't have Sentry, then you can enable
the console logging or you can implement your provider and [contribute](/contribute) to the project. This shouldn't be
so elaborate, because [logrus](https://github.com/sirupsen/logrus/) allows to set a hook.

If you use Sentry, you need to add the DSN in the environment. For this you can add a .env-file or directory in the
environment variables.

```env {filename=".env"}
SENTRY_LOGGING_DNS=https://1245@asdf.ingest.sentry.io/67890
```