---
title: Slack
weight: 106
prev: docs/getting-started
next: docs/configuration/application
sidebar:
  open: true
---

## How to connect Slack

{{% steps %}}

### Step 1

[Log in](https://slack.com/signin) to your Slack account on your workspace. If you don't have one yet, register on your
workspace `<your-workspace>.slack.com`.

### Step 2

Go to your [Slack integrations](https://api.slack.com/apps).

### Step 3

Create a Slack integration. After the creation go to ***Features > OAuth & Permissions*** and then
***Scopes > Bot Token Scopes***. Here you need to add this permission: `chat:write` to send
messages. You can also add `channels:join`, so that your integration can join public channels independently (optional).
Go in this page to ***OAuth Tokens for Your Workspace > Bot User OAuth Token*** click "install to Workspace" and copy
this token. This token should have this format: `xoxb-123456abcdef...` You can also add in
***Settings > Basic Information > Display Information*** a bot image, description, background color and the shown name (
not functional, only for design).

### Step 4

Add your Slack API-Key to your project. At the moment, you need to configure it via Swagger (or cURL, ...). (We are
working on a frontend.) Replace `<your-project-id>` in the url with your notify-project-id

POST `/v1/settings/projects/<your-project-id>/integrations/slack`

```json
{
  "bot_user_o_auth_token": "Your generated Mailgun API-key e.g. xoxb-123456abcdef..."
}
```

{{% /steps %}}