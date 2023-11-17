<h1 align="center">Notify</h1>

<p align="center">
    <a href="https://github.com/m-mattia-m/notify/blob/main/LICENSE">
        <img alt="Licence" src="https://img.shields.io/github/license/m-mattia-m/notify"/>
    </a>
    <a href="https://github.com/m-mattia-m/notify/actions">
        <img alt="GitHub Workflow Status (with event)"  src="https://img.shields.io/github/actions/workflow/status/m-mattia-m/notify/deploy.yaml">
    </a>
    <a href="https://github.com/m-mattia-m/notify/releases">
        <img alt="Release" src="https://badgen.net/github/release/m-mattia-m/notify/stable" />
    </a>
    <a href="https://goreportcard.com/report/github.com/m-mattia-m/notify">
        <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/m-mattia-m/notify" />
    </a>
    <a href="https://github.com/m-mattia-m/notify/graphs/contributors">
        <img alt="GitHub contributors" src="https://img.shields.io/github/contributors/m-mattia-m/notify">
    </a>
</p>

<p align="center">
    <a href="https://github.com/m-mattia-m/notify/issues">
        <img alt="issues" src="https://img.shields.io/github/issues/m-mattia-m/notify?label=issues"/>
    </a>
    <a href="https://github.com/m-mattia-m/notify/issues">
        <img alt="open issues" src="https://img.shields.io/github/issues-raw/m-mattia-m/notify?label=open%20issues"/>
    </a>
    <a href="https://github.com/m-mattia-m/notify/issues">
        <img alt="Licence" src="https://img.shields.io/github/issues-closed/m-mattia-m/notify?label=closed%20issues"/>
    </a>
</p>

With Notify, you can send messages securely from your frontend. You don't need your own backend to access Mailgun or
Slack, for example. Notify is a simple, lightweight, and secure way to send messages from your frontend.

## Getting started

### Installation

@ **TODO**

#### Code

@ **TODO**

#### Image

@ **TODO**

## Usage

### Flow 

#### Fields
- `message_template_type` -> `text/plain`, `text/html`
- `source_type` -> `slack`, `mailgun`

#### Message-Template

You can create a message-template for the notification. For example

`Title: {{subject}}\nHere is your message: {{message}}\Best regards\nNotify`

### Notification request

You can send a notification in this format:

```json5
{
  project_id: "12345", // not secret
  subject: "Important",
  message: "This is the important message.",
  target: "mailgun:john.doe@example.com;mailgun:john@doe.com;slack:A1B2C3D4", // optional, only required when no target specified in the flow | define your notification provider
}
```

### Credentials

#### Mailgun

- `api_base`: `eu`, `us`

## Config

### DNS

If you use Notify in the self-hosted version, you can specify your own DNS server and verify the hosts. If you do not
specify anything, a standard DNS server from Google will be used (8.8.8.8:53). To specify your DNS server, configure the
following in the config.yaml file:
****

```yaml
domain:
  dns:
    verifyDns: 8.8.8.8:53 # this is optional -> if not set then the google standard is used ("8.8.8.8:53")
```

### Config file

Create a config.yaml file in the `root` project or in `./configs/`.

All possible configurations are listed here:

```yaml
app:
  name: notify # required
  env: DEV # required

server:
  scheme: http # required
  domain: localhost # required
  port: 8084 # required
  version: v1 # required
  # URL: $(SCHEME)://$(DOMAIN):$(PORT) -> is mapped together during runtime # is not required, is for information only

db:
  mongo:
    host: localhost # required
    port: 27018 # required
    name: notify # required
    user: user # required
    password: your!user!password # required

logging:
  enable:
    console: true # required
    sentry: false # required
  sentry:
    dsn: https://1234@asdf.ingest.sentry.io/4321 # required
    level: debug # required

authentication:
  oidc:
    issuer: https://instance.zitadel.cloud # required
    clientId: 1234567890@project-name # required

frontend:
  url: http://localhost:3000 # not currently required, but will be added in the future

domain:
  dns:
    verifyDns: 8.8.8.8:53 # this is optional -> if not set then the google standard is used ("8.8.8.8:53")
```

## Authentication

Currently only [Zitadel (OpenSource)](https://github.com/zitadel/zitadel) is directly supported. However, as it is an
OIDC provider, you are welcome to try it with your OIDC provider.

## Development

### Localhost

If you want to release a localhost in Notify during development, you must use the `local` stage and start
with `localhost:`. It doesn't matter which port you need.

For example, the host to be registered may look like this:

```yaml
{
  "host": "localhost:8084", // Note that `localhost` is required with a colon
  "stage": "local"
}
```

## Contribution

### Swagger

To generate the swagger-docs, use the output-directory `./swagger-docs`, as the documentation is stored in the `./docs` folder.

- `swag init --output "./swagger-docs" --parseInternal --parseDependency true`






