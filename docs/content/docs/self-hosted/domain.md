---
title: Domain
weight: 308
prev: docs/configuration/frontend
next: docs/configuration/database
sidebar:
  open: true
---

```yaml {filename="./configs/config.yaml"}
domain:
  dns:
    verifyDns: 8.8.8.8:53 # this is optional -> if not set then the google standard is used ("8.8.8.8:53")
  activity:
    enable:
      subject: true # this is optional
      message: true # this is optional
  swagger:
    port: false # this is optional
```

## DNS

Per default Notify uses the DNS server from Google (`8.8.8.8:53`). But if you use the self hosted possibility, you can
set your DNS in the config.

```yaml {filename="./configs/config.yaml"}
domain:
  dns:
    verifyDns: 8.8.8.8:53 # this is optional -> if not set then the google standard is used ("8.8.8.8:53")
```

## Activity

You can also activate whether the subject or the message should be logged in the activities.

```yaml {filename="./configs/config.yaml"}
domain:
  activity:
    enable:
      subject: true # optional -> default: false
      message: true # optional -> default: false
```

## Swagger

It could be, that you don't need the port in the swagger-configuration. For this you can disable it in the config. Note,
the port is needed to start the API.

```yaml {filename="./configs/config.yaml"}
domain:
  swagger:
    port: false # this is optional
```

