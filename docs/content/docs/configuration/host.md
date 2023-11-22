---
title: Host
weight: 103
prev: docs/getting-started
next: docs/configuration/application
sidebar:
  open: true
---

{{< callout emoji="ℹ️" type="info" >}}
At the moment you only can create resources via swagger (or cURL, ...). We are working on Frontend and plan to release
it the beginning of 2024.
{{< /callout >}}

## Introduction

In the frontend you should not store API-keys or other credentials. For this you should work with verified hosts. Here
you can use Notify.

## Create

You can add a host to your project. Replace `<project-id>` in the URL with your project id.

POST `v1/settings/projects/<project-id>/hosts`

```json {filename="body"}
{
  "host": "add your host e.g. notify.example.com",
  "stage": "prod"
}
```

## Update

You can't update a host, because of the verification is to complex. You need to delete a host and add a new one.

## Verify

You need to verify a host. For this you need to add your `verify_token` in your
DNS ([configure your own DNS](../../self-hosted/domain#dns)). For this you need go to your domain provider (e.g.
Cloudflare) and add a `TXT`-record with the tokens as the value. The `verify_token` looks like
this: `notify-verification::abc123-de45-fg67-hi89-jklmn01234` after this you can send the verification-request (if it
fails, be aware that a DNS-record can take up to 72 hours until it is active.). Replace `<project-id>` in the URL with
your project id.

For your development/local environment look [here](../..//development/localhost#dns).

PUT `v1/settings/projects/<project-id>/hosts`

```json {filename="body"}
{
  "host": "add your host e.g. notify.example.com",
  "stage": "prod"
}
```

## List

You can list all hosts of your project. Replace `<project-id>` in the URL with your project-id.

GET `v1/settings/projects/<project-id>/hosts`

## Get

You can get a specific host in your project by their ID. Replace `<project-id>` in the URL with your project-id and `<host-id>` with your host-id.

GET `v1/settings/projects/<project-id>/hosts/<host-id>`

## Delete

You can delete a specific host in your project by their ID. Replace `<project-id>` in the URL with your project-id and `<host-id>` with your host-id.

DELETE `v1/settings/projects/<project-id>/hosts/<host-id>`




