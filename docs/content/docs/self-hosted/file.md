---
title: Config file
weight: 310
prev: docs/configuration/domain
next: docs/development
sidebar:
  open: true
---

## Config file

Create a config.yaml file in the project `root`, in `./configs/` or in `./config/`.

All possible configurations are listed here:

```yaml  {filename="./configs/config.yaml or ./config.yaml"}
app:
  name: notify # required
  env: DEV # required

server:
  scheme: http # required
  domain: localhost # required
  port: 8080 # required
  version: v1 # required: in the most cases always 'v1'

logging:
  enable:
    console: true # required
    sentry: true # required

authentication:
  oidc:
    issuer: https://your-instance.zitadel.cloud # required
    clientId: 12345@notify # required

frontend:
  url: http://localhost:3000 # required

domain:
  dns:
    verifyDns: 8.8.8.8:53 # this is optional -> if not set then the google standard is used ("8.8.8.8:53")
  activity:
    enable:
      subject: true # required
      message: true # required
```

## Environment

Create a .env file in the project `root`.

All possible configurations are listed here:

- `MONGO_TLS_ACTIVE` is only required when your DB use TLS.
- `MONGO_PORT` is optional, if your MongoDb host don't need a port, you can remove this attribute.

```env {filename=".env"}
MONGO_HOST=localhost # required
MONGO_PORT=27017 # optional
MONGO_DATABASE_NAME=notify # required
MONGO_USERNAME=admin # required
MONGO_PASSWORD=admin!password # required
MONGO_TLS_ACTIVE=true # optional

# only required when logging.enable.sentry in the config-file is true.
SENTRY_LOGGING_DNS=https://1245@asdf.ingest.sentry.io/67890
```