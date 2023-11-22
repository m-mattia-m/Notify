---
title: Project
weight: 102
prev: docs/getting-started
next: docs/configuration/application
sidebar:
  open: true
---

{{< callout emoji="ℹ️" type="info" >}}
At the moment you only can create resources via swagger (or cURL, ...). We are working on Frontend and plan to release
it the beginning of 2024.
{{< /callout >}}

## Create

A user can create unlimited projects with unique names (only for user itself unique).

POST `/v1/settings/projects`

```json {filename="body"}
{
  "name": "your-project-name"
}
```

## Update

Replace `<project-id>` in the URL with your project id. If you do not know this ID, you can list all your project
with [list-projects](../../configuration/project#list).

PUT `/v1/settings/projects/<project-id>`

```json {filename="body"}
{
  "name": "your-updated-project-name"
}
```

## List

List all your projects.

GET `/v1/settings/projects`

## Get

Get a specific project by their ID. Replace `<project-id>` in the URL with your project id. If you do not know this ID,
you can list all your project with [list-projects](../../configuration/project#list).

GET `/v1/settings/projects/<project-id>`

## Delete

Delete a specific project by their ID. Replace `<project-id>` in the URL with your project id. If you do not know this
ID, you can list all your project with [list-projects](../../configuration/project#list).

DELETE `/v1/settings/projects`



