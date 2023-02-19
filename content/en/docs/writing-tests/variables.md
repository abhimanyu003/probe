---
title: "Variables"
description: "This article will introduce how to use variables."
weight: 3
---

## Overview

It's possible to have variables `${ VAR }` in your test, sometime you want to export few property from header or body to next stage.

Types of Variables

* Env  Variables. ( System Variable ) `${ env:VAR }`
* Test Variables ( Common for all stages ).
* Run Time Variables / Stage Variables ( Exported on stage run time ).

Exported variable can be used in stage body or header

```yaml
${ env:VAR } // retrieves the value of the os/system variable.
${ VAR } // run time exported variable
```

## Env Variables

It's possible to load OS/system level variable in test. This is also helpful
and allow you to set/load many variable even before test starts

```yaml
${env:USER}
```

Here `${env:USER}` will load the value of `USER` present at OS level.

### Example: Env variables

```yaml
assert:
  body:
    - select: .form.username
      want: ${env:USER}
```

## Test Variables

It's possible to export test level variables that will common for all stages.
These are define at the top of all the stages

```yaml
name: set variables
variables:
  key: value
  secondKey: secondValue
```

### Example: Test Variable

```yaml
name: setting test variables
variables:
  name: abhimanyu
stages:
  - name: send request
    request:
      url: "https://httpbin.org/post"
      method: POST
      formData:
        username: ${name} # we are using exported name variable here
```

## Run Time Variables

It's possible to export any header or body value, that can be used in next stage.

```yaml
export: # export key
  header:
    - select: token
      as: authToken
  body:
    - select: .form.username
      as: username
```

You can also export variable right from assert as well

```yaml
assert:
  body:
    - select: .form.username
      want: "abhimanyu"
      exportAs: "username" # Exporting right from assert
```

### Example: Run Time Variables

```yaml
name: exporting runtime variables
stages:
  - name: send request
    request:
      url: "https://httpbin.org/post"
      method: POST
      formData:
        username: abhimanyu
    assert:
      status: 200
    export: # here we are exporting variables from body
      body:
        - select: .form.username
          as: "username"
```
