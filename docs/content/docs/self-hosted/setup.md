---
title: Setup
weight: 302
prev: docs/self-hosted/
next: docs/self-hosted/application
sidebar:
  open: true
---

## Docker-compose

You need to add the environment and the configuration. [Read more](/docs/self-hosted/file)

```yaml  {filename="docker-compose.yaml"}
version: '3.8'

services:
  mongo:
    image: docker.io/mongo:latest
    environment:
      - MONGO_INITDB_ROOT_USERNAME=${MONGO_USERNAME}
      - MONGO_INITDB_ROOT_PASSWORD=${MONGO_PASSWORD}
      - MONGO_DATABASE_NAME=${MONGO_DATABASE_NAME}
    volumes:
      # database storage
      - mongo-data:/data
    ports:
      - "27017:27017"
    hostname: mongo
    restart: "no"

  notify:
    image: ghcr.io/m-mattia-m/notify-backend:latest
    environment:
      - MONGO_HOST=${MONGO_HOST}
      - MONGO_PORT=${MONGO_PORT}
      - MONGO_DATABASE_NAME=${MONGO_DATABASE_NAME}
      - MONGO_USERNAME=${MONGO_USERNAME}
      - MONGO_PASSWORD=${MONGO_PASSWORD}
    ports:
      - "8080:8080"
    hostname: notify
    restart: "no"

volumes:
  mongo-data: {}

```

## K8s-Manifest

{{< callout type="info" icon="ℹ️" >}}
Coming soon
{{< /callout >}}
