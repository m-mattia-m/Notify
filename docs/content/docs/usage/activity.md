---
title: Activity
weight: 203
prev: docs
next: docs/configuration
sidebar:
  open: true
---

## Introduction

For every request Notify will generate an activity/audit. You can't create or update manually an activity, you can only
list and get the activities.

## Logged data

All activities will be logged in the mongo database.

```json lines {filename="activity-log.json"}
{
  "project_id": "abc123",
  "state": "success",
  "source_type": "slack",
  "target": "ABC123",
  "subject": "Contact form",
  "message": "Hi, I need some technical help.",
  "note": "",
  "updated_at": {
    "$date": "2023-11-18T20:58:16.653Z"
  },
  "created_at": {
    "$date": "2023-11-18T20:58:16.653Z"
  },
  "deleted_at": null
}
```

```json lines {filename="activity-log.json"}
{
  "project_id": "abc123",
  "state": "failed",
  "source_type": "slack",
  "target": "ABC123",
  "subject": "Contact form",
  "message": "Hi, I need some technical help.",
  "note": "NotificationError 'here is a detailed error about the problem'",
  "updated_at": {
    "$date": "2023-11-18T20:56:25.939Z"
  },
  "created_at": {
    "$date": "2023-11-18T20:56:25.939Z"
  },
  "deleted_at": null
}
```

## Hide message and subject

If you are use the self-hosted variant, you can set in the configuration, that the subject and the massage should not be
logged in the activity. [Read more](../../self-hosted/domain#activity)
