---
title: Notification
weight: 202
prev: docs
next: docs/configuration
sidebar:
  open: true
---

## Introduction

You can create projects with verified hosts and flows. If you have correctly configure this, you can send notification
requests to Notify.

## Request

You can only send notifications from a verified host but you don't need a JWT token. The request looks like this:

POST `/v1/notifications`

```json lines {filename="body"}
{
  "project_id": "abc123",
  "subject": "Contact form",
  "message": "Hi, I need some technical help.",
  // you can also remove this attribute
  "target": ""
}
```

## Override target

You can also override your target in the request when your flow this allowed. If you want, you can set multiple targets
but note, that for example mailgun send one mail per recipient. When you want to send multiple targets, you need to
separate them with a Semicolon (;). You need also set for which provider a target is (also when only one or the same
provider are configured). Possible overrides looks like this:

```json lines {filename="target"}
{
  "target": "mailgun:john.dow@example.com;mailgun:support@company.com;slack:A1B2C3;slack:D4E5F6"
}
```