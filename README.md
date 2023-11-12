# Notify

With Notify, you can send messages securely from your frontend. You don't need your own backend to access Mailgun or
Slack, for example. Notify is a simple, lightweight, and secure way to send messages from your frontend.

## Installation

@ **TODO**

## Config

@ **TODO**

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

## Development

If you want to release a localhost in Notify during development, you must use the `local` stage and start
with `localhost:`. It doesn't matter which port you need.

For example, the host to be registered may look like this:

```yaml
{
  "host": "localhost:8084", // Note that `localhost` is required with a colon
  "stage": "local"
}
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
    dsn: https://42c1e2abd797f14755e647e2b3dca963@o4505611535384576.ingest.sentry.io/4506197653323776 # required
    level: debug # required

authentication:
  oidc:
    issuer: https://zitadel.upcraft.li # required
    clientId: 223437944512405507@notion-database-form # required
  zitadel:
    api: zitadel.upcraft.li:443 # required
    projectId: 198654217882173444 # required
    projectName: Notify # required
    organizationId: 198651022292287492 # required
    organizationName: Notify # required

frontend:
  url: http://localhost:3000 # not currently required, but will be added in the future

domain:
  dns:
    verifyDns: 8.8.8.8:53 # this is optional -> if not set then the google standard is used ("8.8.8.8:53")
```

## Authentication

Currently only [Zitadel (OpenSource)](https://github.com/zitadel/zitadel) is directly supported. However, as it is an OIDC provider, you are welcome to try it with your OIDC provider.





