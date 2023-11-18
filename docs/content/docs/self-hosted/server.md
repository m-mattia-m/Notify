---
title: Server
weight: 304
prev: docs/configuration/application
next: docs/configuration/logging
sidebar:
  open: true
---

```yaml {filename="./configs/config.yaml"}
server:
  scheme: http # (http or https) -> since you are submitting credentials it is strongly recommended to use https
  domain: localhost # host domain
  port: 8080 # set the port you want
  version: v1 # version of the API (as long as you don't change anything in the code, it's 'v1' here)
```