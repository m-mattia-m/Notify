---
title: Mailgun
weight: 107
prev: docs/getting-started
next: docs/configuration/application
sidebar:
  open: true
---

## How to connect Mailgun

{{% steps %}}

### Step 1

[Log in](https://login.mailgun.com/login/) to your Mailgun account. If you don't have one
yet, [register](https://signup.mailgun.com/new/signup).

### Step 2

If you don't have your domain
connected, [add](https://help.mailgun.com/hc/en-us/articles/203637190-How-Do-I-Add-or-Delete-a-Domain-) one.

### Step 3

Create an [API-Key](https://app.mailgun.com/settings/api_security) for Notify.

### Step 4

Create a [Notify project](../../project), if you don't have one yet.

### Step 5

Add your Mailgun API-Key to your project. At the moment, you need to configure it via Swagger (or cURL, ...). (We are
working on a frontend.). Replace `<your-project-id>` in the url with your notify-project-id

POST `/v1/settings/projects/<your-project-id>/integrations/mailgun`

```json
{
  "api_key": "Your generated Mailgun API-key",
  "domain": "add your connected Mailgun domain",
  "sender_email": "set the email from which the mail should be sent",
  "sender_name": "set the name, who is the sender of the email"
}
```

{{% /steps %}}