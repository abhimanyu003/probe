---
title: "Logs"
description: "This article will introduce how to Logs."
weight: 6
---

## Overview

Probe captures and logs all the request.
You can find logs for your test under `logs` folder

```text
logs/
  2023-02-13T19:21:39+05:30/
    test-1.log
    test-2.log
  2023-02-13T19:22:13+05:30
    test-1.log
    test-2.log
```

### Example

```text
:authority: httpbin.org
:method: POST
:path: /post
:scheme: https
user-agent: probe ( https://github.com/abhimanyu003/probe )
content-type: application/x-www-form-urlencoded
content-length: 18
accept-encoding: gzip

username=abhimanyu

:status: 200
date: Mon, 13 Feb 2023 16:30:47 GMT
content-type: application/json
content-length: 482
server: gunicorn/19.9.0
access-control-allow-origin: *
access-control-allow-credentials: true

{
  "args": {},
  "data": "",
  "files": {},
  "form": {
    "username": "abhimanyu"
  },
  "headers": {
    "Accept-Encoding": "gzip",
    "Content-Length": "18",
    "Content-Type": "application/x-www-form-urlencoded",
    "Host": "httpbin.org",
    "User-Agent": "probe ( https://github.com/abhimanyu003/probe )",
    "X-Amzn-Trace-Id": "Root=1-63ea65b7-1203aa90794001080b5fdee2"
  },
  "json": null,
  "origin": "103.59.75.66",
  "url": "https://httpbin.org/post"
}
```

## Disable Logs

To disable logs you can use CLI flag `--disableLogs`

```
probe run --disableLogs
```
