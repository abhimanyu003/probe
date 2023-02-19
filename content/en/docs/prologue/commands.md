---
title: "Commands"
description: "How to use probe commands."
draft: false
weight: 990
toc: true
---

## Run

```sh
Usage:
  probe run [flags]

Flags:
      --disableLogs     Disable logs file entries.
      --failfast        Do not start new tests after the first test failure.
  -h, --help            help for run
      --parallel uint   Maximum number of tests to run simultaneously
  -v, --test.v          Get verbose output of the tests.
```

### Run all tests in current folder and sub-directory

```sh
probe run
```

### Run specific test file

```sh
probe run /path/to/test.yaml
```

### Disable Logs

```sh
probe run --disableLogs
```

### Verbose output

```sh
probe run -v
```

### Failfast

```sh
probe run --failfast
```

### Limit parallel run

```sh
probe run  --parallel 2
```

### Help

```sh
probe help
```

### Version

```sh
probe version
```
