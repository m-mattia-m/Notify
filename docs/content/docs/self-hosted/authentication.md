---
title: Authentication
weight: 306
prev: docs/configuration/logging
next: docs/configuration/frontend
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

Notify is tested with [Zitadel](https://github.com/zitadel/zitadel) (open source) but have adhered to the OIDC
standards. You can try to connect your own OIDC provider.

| Provider                                         | Fully tested | not tested |
|--------------------------------------------------|--------------|------------|
| [Zitadel](https://github.com/zitadel/zitadel)    | ✅            | ❌          |
| [Keycloak](https://github.com/keycloak/keycloak) | ❌            | ✅          |
| [Hydra](https://github.com/ory/hydra)            | ❌            | ✅          |
| ...                                              | ❌            | ✅          |


