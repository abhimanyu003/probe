---
title: "Writing Tests"
description: "This article will introduce how to Write Tests."
weight: 1
---

## Overview

Probe makes it easy to write end-to-end API tests using YAML and JQ.

This section illustrates how to write a simple test suite with Probe, and what are the conventions you should use.

The setup is very simple looks like this:

```yaml
name: name of test
stages: # a test can multiple stages
  - name: name of stage
    request:
      # request builder here
    assert:
      status: # assert response status code ( eg: 200, 404, 500 )
      headers:
        # things you want to assert in header
      body:
        # things you want to assert
```

Probe use `jq` as JSON query processor

`jq` is like sed for JSON data - you can use it to slice and filter and map and transform structured data with the same ease that sed, awk, grep and friends let you play with text.

👉 More details on jq: [https://stedolan.github.io/jq/](https://stedolan.github.io/jq/)

You can use all the functions and power of jq 🙂

### Quick Example

* Create a file `main.yaml` ( filename can be anything )

```yaml
name: Writing Test
stages:
  - name: first stage
    request:
      url: "https://httpbin.org/get"
      method: GET
    assert:
      status: 200
```


* Running test


```
$ probe run
```


## Stages

A test is required to have a `stage`, which allow your to send HTTP Request and Assert Body.
Test can have multiple `stages`, stages run in sequential order.

```yaml
name: name of test
stages: # a test can multiple stages
  - name: first stage
    request:
      # request for first stage here
    assert:
      # things you want to assert

  - name: second stage
    request:
      # request for second stage here
    assert:
      # things you want to assert
```

### Example: Single Stage Test

* Create a file with name `main.yaml` ( filename can be anything )


```yaml
name: Writing Test
stages:
  - name: first stage
    request:
      url: "https://httpbin.org/get"
      method: GET
    assert:
      status: 200
```

### Example: Multiple Stage Test

* Create a file with name `main.yaml` ( filename can be anything )

```yaml
name: Writing Test
stages:
  - name: first stage
    request:
      url: "https://httpbin.org/get"
      method: GET
    assert:
      status: 200

  - name: second stage
    request:
      url: "https://httpbin.org/post"
      method: POST
    assert:
      status: 200
```
* Running test

```
$ probe run
```

