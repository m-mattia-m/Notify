<h1 align="center">Notify</h1>

<p align="center">
    <a href="https://github.com/m-mattia-m/notify/blob/main/LICENSE" style="text-decoration: inherit;">
        <img alt="Licence" src="https://img.shields.io/github/license/m-mattia-m/notify"/>
    </a>
    <a href="https://github.com/m-mattia-m/notify/actions" style="text-decoration: inherit;">
        <img alt="GitHub Workflow Status (with event)"  src="https://img.shields.io/github/actions/workflow/status/m-mattia-m/notify/docker-build-publish.yaml">
    </a>
    <a href="https://github.com/m-mattia-m/notify/releases" style="text-decoration: inherit;">
        <img alt="Release" src="https://badgen.net/github/release/m-mattia-m/notify/stable" />
    </a>
    <a href="https://goreportcard.com/report/github.com/m-mattia-m/notify" style="text-decoration: inherit;">
        <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/m-mattia-m/notify" />
    </a>
    <a href="https://github.com/m-mattia-m/notify/graphs/contributors" style="text-decoration: inherit;">
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
</p>

# What is Notify?

Notify is a simple secure message gateway witch allows you to send messages from your frontend. Connect Slack or Mailgun
to send Messages. Notify is a simple, lightweight, and secure way to send messages from your frontend.

[See the full documentation here.](https://m-mattia-m.github.io/Notify/)

# Features

- 🔐 **Secure messaging:** You don't need your own gateway for secure messaging from the frontend.
- 💬 **Slack:** You can add your own Slack integration.
- ✉️ **Mailgun:** You can add your own Mailgun integration.
- 🖼️ **Template:** You can add your HTML or TEXT template and replace the content dynamic
- 💨 **Flows:** You can add multiple flows for every project with different providers.
- 📝 **Activity:** We log all notifications for you, so that you have all activities logged.

# Getting started

## Cloud

> ℹ️ At the moment you only can choose the self-hosted option. In the future we will offer you a fully managed cloud
> solution. We plan with the fully managed solution in at the beginning of 2024.

## Self-hosted

You can host your own Notify if you need this.

### Requirements

- 🖥️ You need a server where you can host Notify.
- 🗄️ You need a mongoDB where you can store the data.

# Setup

Notify is Open Source! So you can host Notify for yourself.

## Docker-compose

You need to add the environment and the configuration. You should have this folder structure:

- config
  - config.yaml
- .env
- docker-compose.yaml

> If you use the docker-compose on a server, you need to add a reverse proxy for a secure communication.

```yaml {filename="docker-compose.yaml"}
version: '3.8'

services:
  mongo:
    image: docker.io/mongo:7.0.3
    hostname: mongo
    environment:
      - MONGO_INITDB_ROOT_USERNAME=${MONGO_USERNAME}
      - MONGO_INITDB_ROOT_PASSWORD=${MONGO_PASSWORD}
      - MONGO_DATABASE_NAME=${MONGO_DATABASE_NAME}
    ports:
      - "${MONGO_PORT}:27017"
    volumes:
      - mongo-data:/data
    restart: "no"
    networks:
      - notify

  notify:
    image: ghcr.io/m-mattia-m/notify:v1.0.1
    hostname: notify
    environment:
      - MONGO_HOST=mongo # use here the docker service name (if you don't change anything here it's 'mongo')
      - MONGO_PORT=${MONGO_PORT}
      - MONGO_DATABASE_NAME=${MONGO_DATABASE_NAME}
      - MONGO_USERNAME=${MONGO_USERNAME}
      - MONGO_PASSWORD=${MONGO_PASSWORD}
      - SENTRY_LOGGING_DNS=${SENTRY_LOGGING_DNS}
    ports:
      - "8080:8080"
    volumes:
      - ./config/config.yaml:/app/config.yaml
    restart: "no"
    depends_on:
      - mongo
    networks:
      - notify

volumes:
  mongo-data: { }

networks:
  notify:
    driver: bridge
```

## Kubernetes-Manifest

Here you will find all K8s manifests from the namespace to the ingress. Create a DNS A-record with your hostname
e.g. `api.notify.example.com` and add the IP-address from your load-balancer as the target.

```yaml {filename="k8s-manifest-namespace.yaml"}
apiVersion: v1
kind: Namespace
metadata:
  name: notify
```

```yaml {filename="k8s-manifest-config-map.yaml"}
apiVersion: v1
kind: ConfigMap
metadata:
  name: notify-configuration
  namespace: notify
data:
  config.yaml: |-
    app:
      name: notify
      env: PROD

    server:
      scheme: http
      domain: api.notify.example.com
      port: 8080
      version: v1

    logging:
      enable:
        console: true
        sentry: true

    authentication:
      oidc:
        issuer: https://your-instance.zitadel.cloud
        clientId: 12345@notify

    frontend:
      url: https://notify.example.com

    domain:
      dns:
        verifyDns: 8.8.8.8:53 # this is optional -> if not set then the google standard is used ("8.8.8.8:53")
      activity:
        enable:
          subject: true
          message: true
```

```yaml {filename="k8s-manifest-secret.yaml"}
apiVersion: v1
kind: Secret
metadata:
  name: notify-secrets
  namespace: notify
data: # all values must be base64 encoded
  MONGO_HOST: YXNkZi10ZXN0LXNlY3JldA== # base54 encoded value
  MONGO_PORT: YXNkZi10ZXN0LXNlY3JldA== # base54 encoded value
  MONGO_DATABASE_NAME: YXNkZi10ZXN0LXNlY3JldA== # base54 encoded value
  MONGO_USERNAME: YXNkZi10ZXN0LXNlY3JldA== # base54 encoded value
  MONGO_PASSWORD: YXNkZi10ZXN0LXNlY3JldA== # base54 encoded value
  SENTRY_LOGGING_DNS: YXNkZi10ZXN0LXNlY3JldA== # base54 encoded value
```

```yaml {filename="k8s-manifest-deployment.yaml"}
# Here the application itself is created
apiVersion: apps/v1
kind: Deployment
metadata:
  name: notify
  namespace: notify
spec:
  replicas: 1 # set the number of pods you want
  selector:
    matchLabels:
      app: notify
  template:
    metadata:
      labels:
        app: notify
    spec:
      containers:
        - name: notify
          image: ghcr.io/m-mattia-m/notify:v1.0.1
          env:
            - name: MONGO_HOST
              valueFrom:
                secretKeyRef:
                  name: notify-secrets
                  key: MONGO_HOST
            - name: MONGO_PORT
              valueFrom:
                 secretKeyRef:
                    name: notify-secrets
                    key: MONGO_PORT
            - name: MONGO_DATABASE_NAME
              valueFrom:
                secretKeyRef:
                  name: notify-secrets
                  key: MONGO_DATABASE_NAME
            - name: MONGO_USERNAME
              valueFrom:
                secretKeyRef:
                  name: notify-secrets
                  key: MONGO_USERNAME
            - name: MONGO_PASSWORD # is not required
              valueFrom:
                secretKeyRef:
                  name: notify-secrets
                  key: MONGO_PASSWORD
            - name: SENTRY_LOGGING_DNS
              valueFrom:
                secretKeyRef:
                  name: notify-secrets
                  key: SENTRY_LOGGING_DNS
          volumeMounts:
            - name: config-volume
              mountPath: ./app/config.yaml
              subPath: config.yaml
      volumes:
        - name: config-volume
          configMap:
            name: notify-configuration
            items:
              - key: config.yaml
                path: config.yaml
```

```yaml {filename="k8s-manifest-service.yaml"}
# here is the service-creation to expose the app in the namespace
apiVersion: v1
kind: Service
metadata:
  name: notify
  namespace: notify
spec:
  selector:
    app.kubernetes.io/name: notify
    app: notify
  ports:
    - name: http
      port: 80 # external port
      targetPort: 8080 # internal port (check if in the config.yaml server.port is the same)
```

```yaml {filename="k8s-manifest-ingress.yaml"}
# here is the route-creation to expose the app-service to the internet

apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: notify
  namespace: notify
  annotations:
    cert-manager.io/issuer: letsencrypt-nginx
spec:
  rules:
    - host: api.notify.example.com # set here your domain
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: notify
                port:
                  number: 80
  ingressClassName: nginx
```

# Configuration

Go on the domain where you have hosted your Notify. Then you will be redirected to the swagger of Notify. Here you can
create your resources. All this requests need a valid JWT token from a valid user from
your [OIDC-Provider](https://m-mattia-m.github.io/notify/docs/configuration/providers) (in
cloud: [Zitadel](https://github.com/zitadel/zitadel/)).

> At the moment you only can create resources via swagger (or cURL, ...). We are working on Frontend and plan to release
> it the beginning of 2024.

## Project

### Create

A user can create unlimited projects with unique names (only for user itself unique).

POST `/v1/settings/projects`

```json {filename="body"}
{
  "name": "your-project-name"
}
```

### Update

Replace `<project-id>` in the URL with your project id. If you do not know this ID, you can list all your project
with [list-projects](/docs/configuration/project#list).

PUT `/v1/settings/projects/<project-id>`

```json {filename="body"}
{
  "name": "your-updated-project-name"
}
```

### List

List all your projects.

GET `/v1/settings/projects`

### Get

Get a specific project by their ID. Replace `<project-id>` in the URL with your project id. If you do not know this ID,
you can list all your project with [list-projects](/docs/configuration/project#list).

GET `/v1/settings/projects/<project-id>`

### Delete

Delete a specific project by their ID. Replace `<project-id>` in the URL with your project id. If you do not know this
ID, you can list all your project with [list-projects](/docs/configuration/project#list).

DELETE `/v1/settings/projects`

## Flow

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

### Create

You can create a flow in your project. Replace `<project-id>` in the URL with your project id.

POST `/v1/settings/projects/<project-id>/flows`

```yaml
{
  # defines if a flow is action and can be triggered
  "active": true,
  # here you can place your
  "message_template": "Hi,\n\nYou have received a message about '{{subject}}' from your contact form: \n\n{{message}",
  # defines the type of the message: text/html, text/plain
  "message_template_type": "text/plain",
  # you need to set a name for your flow. e.g. team-slack-notification
  "name": "Contact form notification",
  # defines if the message can override your default target
  "override_target": false,
  # defines where you want to be notified. e.g. slack, mailgun
  "source_type": "mailgun",
  # here you can past your default receiver email for mailgun or your Slack channel id 
  # (add the integration to your channel before connection Notify with it)
  "target": "you@example.com"
}
```

### Update

You can update a existing flow in your project. Replace `<project-id>` in the URL with your project id and `flow-id`
with your flow-id.

PUT `/v1/settings/projects/<project-id>/flows/<flow-id>`

```yaml
{
  # defines if a flow is action and can be triggered
  "active": true,
  # here you can place your
  "message_template": "Hi,\n\nYou have received a message about '{{subject}}' from your contact form: \n\n{{message}",
  # defines the type of the message: text/html, text/plain
  "message_template_type": "text/plain",
  # you need to set a name for your flow. e.g. team-slack-notification
  "name": "Contact form notification",
  # defines if the message can override your default target
  "override_target": false,
  # defines where you want to be notified. e.g. slack, mailgun
  "source_type": "mailgun",
  # here you can past your default receiver email for mailgun or your Slack channel id 
  # (add the integration to your channel before connection Notify with it)
  "target": "you@example.com"
}
```

### List

You can list all your flows in a project. Replace `<project-id>` in the URL with your project id.

GET `/v1/settings/projects/<project-id>/flows`

### Get

You can get a specific flow in a project by their ID. Replace `<project-id>` in the URL with your project id
and `flow-id`
with your flow-id.

GET `/v1/settings/projects/<project-id>/flows/<flow-id>`

### Delete

You can delete a specific flow in a project by their ID. Replace `<project-id>` in the URL with your project id
and `flow-id`
with your flow-id.

DELETE `/v1/settings/projects/<project-id>/flows/<flow-id>`

## Host

In the frontend you should not store API-keys or other credentials. For this you should work with verified hosts. Here
you can use Notify.

### Create

You can add a host to your project. Replace `<project-id>` in the URL with your project id.

POST `v1/settings/projects/<project-id>/hosts`

```json {filename="body"}
{
  "host": "add your host e.g. notify.example.com",
  "stage": "prod"
}
```

### Update

You can't update a host, because of the verification is to complex. You need to delete a host and add a new one.

### Verify

You need to verify a host. For this you need to add your `verify_token` in your
DNS ([configure your own DNS](/docs/self-hosted/domain#dns)). For this you need go to your domain provider (e.g.
Cloudflare) and add a `TXT`-record with the tokens as the value. The `verify_token` looks like
this: `notify-verification::abc123-de45-fg67-hi89-jklmn01234` after this you can send the verification-request (if it
fails, be aware that a DNS-record can take up to 72 hours until it is active.). Replace `<project-id>` in the URL with
your project id.

For your development/local environment look [here](/docs/development/localhost#dns).

PUT `v1/settings/projects/<project-id>/hosts`

```json {filename="body"}
{
  "host": "add your host e.g. notify.example.com",
  "stage": "prod"
}
```

### List

You can list all hosts of your project. Replace `<project-id>` in the URL with your project-id.

GET `v1/settings/projects/<project-id>/hosts`

### Get

You can get a specific host in your project by their ID. Replace `<project-id>` in the URL with your project-id
and `<host-id>` with your host-id.

GET `v1/settings/projects/<project-id>/hosts/<host-id>`

### Delete

You can delete a specific host in your project by their ID. Replace `<project-id>` in the URL with your project-id
and `<host-id>` with your host-id.

DELETE `v1/settings/projects/<project-id>/hosts/<host-id>`

## Providers

| Implemented | Planned    |
|-------------|------------|
| Slack       | Teams      |
| Mailgun     | Discord    |
|             | SMTP       |
|             | (WhatsApp) |
|             | (Twitter)  |
|             | (...)      |

### Slack

1. [Log in](https://slack.com/signin) to your Slack account on your workspace. If you don't have one yet, register on
   your
   workspace `<your-workspace>.slack.com`.
2. Go to your [Slack integrations](https://api.slack.com/apps).
3. Create a Slack integration. After the creation go to ***Features > OAuth & Permissions*** and then
   ***Scopes > Bot Token Scopes***. Here you need to add this permission: `chat:write` to send
   messages. You can also add `channels:join`, so that your integration can join public channels independently (
   optional).
   Go in this page to ***OAuth Tokens for Your Workspace > Bot User OAuth Token*** and copy this token. This token
   should
   have this format: `xoxb-123456abcdef...` You can also add in ***Settings > Basic Information > Display Information***
   a
   bot image, description, background color and the shown name (not functional, only for design).
4. Add your Slack API-Key to your project. At the moment, you need to configure it via Swagger (or cURL, ...). (We are
   working on a frontend.) Replace `<your-project-id>` in the url with your notify-project-id

   POST `/v1/settings/projects/<your-project-id>/integrations/slack`

   ```json
   {
      "bot_user_o_auth_token": "Your generated Mailgun API-key e.g. xoxb-123456abcdef..."
   }
   ```

### Mailgun

1. [Log in](https://login.mailgun.com/login/) to your Mailgun account. If you don't have one
   yet, [register](https://signup.mailgun.com/new/signup).
2. If you don't have your domain
   connected, [add](https://help.mailgun.com/hc/en-us/articles/203637190-How-Do-I-Add-or-Delete-a-Domain-) one.
3. Create an [API-Key](https://app.mailgun.com/settings/api_security) for Notify.
4. Create a [Notify project](/docs/configuration/project), if you don't have one yet.
5. Add your Mailgun API-Key to your project. At the moment, you need to configure it via Swagger (or cURL, ...). (We are
   working on a frontend.). Replace `<your-project-id>` in the url with your notify-project-id
   \
   \
   POST `/v1/settings/projects/<your-project-id>/integrations/mailgun`

   ```json
   {
      "api_key": "Your generated Mailgun API-key",
      "domain": "add your connected Mailgun domain",
      "sender_email": "set the email from which the mail should be sent",
      "sender_name": "set the name, who is the sender of the email"
   }
   ```

# Usage

## Activity (Audit)

For every request Notify will generate an activity. You can't create or update manually an activity, you can only list
and get the activities.

### Logged data

All activities will be logged in the mongo database.

```json
{
  "project_id": "abc123",
  "state": "success",
  "source_type": "slack",
  "target": "ABC123",
  "subject": "Contact form",
  "message": "Hi, I need some technical help.",
  "note": "",
  "updated_at": {
    "$date": "2023-11-18T20:58:16.653Z"
  },
  "created_at": {
    "$date": "2023-11-18T20:58:16.653Z"
  },
  "deleted_at": null
}
```

```json
{
  "project_id": "abc123",
  "state": "failed",
  "source_type": "slack",
  "target": "ABC123",
  "subject": "Contact form",
  "message": "Hi, I need some technical help.",
  "note": "NotificationError 'here is a detailed error about the problem'",
  "updated_at": {
    "$date": "2023-11-18T20:56:25.939Z"
  },
  "created_at": {
    "$date": "2023-11-18T20:56:25.939Z"
  },
  "deleted_at": null
}
```

### Hide message and subject

If you are use the self-hosted variant, you can set in the configuration, that the subject and the massage should not be
logged in the activity. [Read more](/docs/self-hosted/domain#activity)

## Notification

You can create projects with verified hosts and flows. If you have correctly configure this, you can send notification
requests to Notify.

### Request

You can only send notifications from a verified host but you don't need a JWT token. The request looks like this:

POST `/v1/notifications`

```yaml
{
  "project_id": "abc123",
  "subject": "Contact form",
  "message": "Hi, I need some technical help.",
  # you can also remove this attribute
  "target": ""
}
```

### Override target

You can also override your target in the request when your flow this allowed. If you want, you can set multiple targets
but note, that for example mailgun send one mail per recipient. When you want to send multiple targets, you need to
separate them with a Semicolon (;). You need also set for which provider a target is (also when only one or the same
provider are configured). Possible overrides looks like this:

```json
{
  "target": "mailgun:john.dow@example.com;mailgun:support@company.com;slack:A1B2C3;slack:D4E5F6"
}
```

# Self hosted

## Config file

Create a *config.yaml* file in the project `root`, in `./configs/` or in `./config/`.

All possible configurations are listed here:

```yaml
app:
  name: notify # required
  env: DEV # required

server:
  scheme: http # required
  domain: localhost # required
  port: 8080 # required
  version: v1 # required: in the most cases always 'v1'

database:
   mongo:
      authMechanism: SCRAM-SHA-256 # optional
      srv: true # optional
      tls: true # is only required when your DB use TLS.

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

Create a *.env* file in the project `root`.

All possible configurations are listed here:

```env
MONGO_HOST=localhost # required
MONGO_PORT=27017 # required
MONGO_DATABASE_NAME=notify # required
MONGO_USERNAME=admin # required
MONGO_PASSWORD=admin!password # required

# only required when logging.enable.sentry in the config-file is true.
SENTRY_LOGGING_DNS=https://1245@asdf.ingest.sentry.io/67890
```

# Development

If you are developing, you find all the configurations and peculiarities for the development here.

## DNS

If you want to release a localhost in Notify during development, you must use the `local` stage and start
with `localhost:`. It doesn't matter which port you need.

For example, the host to be registered may look like this:

```yaml
{
  "host": "localhost:8084", # Note that `localhost` is required with a colon
  "stage": "local"
}
```

# Contribute

You are welcome to contribute to Notify. Just do what you think is good and important. Go
to [Github](https://github.com/m-mattia-m/Notify) for this.

## Ideas

- add [messaging providers](/docs/configuration/providers) (like Discord, Microsoft Teams, SMTP, Twitter, Instagram,
  WhatsApp, ...)
- test more [OIDC providers](/docs/self-hosted/authentication) and document if it works.
- add [logging providers](/)

## Development

### Swagger

To generate the swagger-docs, use the output-directory `./swagger-docs`, as the documentation is stored in the `./docs`
folder.

- `swag init --output "./swagger-docs" --parseInternal --parseDependency true`

