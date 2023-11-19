---
title: Setup
weight: 302
prev: docs/self-hosted/
next: docs/self-hosted/application
sidebar:
  open: true
---

## Docker-compose

You need to add the environment and the configuration. [Read more](/docs/self-hosted/file). You should have this folder
structure:

{{< filetree/container >}}
{{< filetree/folder name="config" >}}
{{< filetree/file name="config.yaml" >}}
{{< /filetree/folder >}}
{{< filetree/file name=".env" >}}
{{< filetree/file name="docker-compose.yaml" >}}
{{< /filetree/container >}}

{{< callout type="warning" >}}
If you use the docker-compose on a server, you need to add a reverse proxy for a secure communication.
{{< /callout >}}

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

## K8s-Manifest

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
    - host: notify.formtion.app # set here your domain
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