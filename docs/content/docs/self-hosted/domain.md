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
      subject: true
      message: true
```

## DNS

Per default Notify uses the DNS server from Google (`8.8.8.8:53`).

## Activity

You can also activate whether the subject or the message should be logged in the activities.