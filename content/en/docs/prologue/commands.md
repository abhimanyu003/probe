---
title: "Commands"
description: "How to use probe commands."
draft: false
weight: 990
toc: true
---

## Run

```
Usage:
  probe run [flags]

Flags:
      --disableLogs     Disable logs file entries.
      --failfast        Do not start new tests after the first test failure.
  -h, --help            help for run
      --parallel uint   Maximum number of tests to run simultaneously
  -v, --test.v          Get verbose output of the tests.
```

**Run all tests in current folder and sub-directory**

```
probe run
```

**Run specific test file**


```
probe run /path/to/test.yaml
```

**Disable Logs**

```
probe run --disableLogs
```

**Verbose output**

```
probe run -v
```

**Failfast**

```
probe run --failfast
```

**Limit parallel run**

```
probe run  --parallel 2
```

## Help


```
probe help
```


## Version

```
probe version
```
