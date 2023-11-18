---
title: Localhost
weight: 402
prev: docs/development
next: docs/self-hosted
sidebar:
  open: true
---

## DNS

If you want to release a localhost in Notify during development, you must use the `local` stage and start
with `localhost:`. It doesn't matter which port you need.

For example, the host to be registered may look like this:

```yaml
{
  "host": "localhost:8084", // Note that `localhost` is required with a colon
  "stage": "local"
}
```