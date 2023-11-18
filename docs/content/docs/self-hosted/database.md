---
title: Database
weight: 309
prev: docs/configuration/domain
next: docs/configuration/file
sidebar:
  open: true
---

Notify uses a MongoDB as their database. For this you need to configure your MongoDB credentials as your environment.
You can add a .env-file or you also can add the environment variables directly into the environment.

```env  {filename=".env"}
MONGO_HOST=localhost
MONGO_PORT=27017
MONGO_DATABASE_NAME=notify
MONGO_USERNAME=admin
MONGO_PASSWORD=admin!password
```
