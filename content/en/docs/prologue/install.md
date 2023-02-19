---
title: "Install"
description: "How to install probe."
draft: false
weight: 110
toc: true
---

Probe support wide variety of OS
* 'linux', 'darwin', 'windows'
* '386', 'amd64', 'arm64'

#### Auto Install

```
curl -L https://raw.githubusercontent.com/abhimanyu003/probe/main/install.sh | bash
```

*This script installs the latest release by default.*

#### Brew

```
brew install abhimanyu003/tap/probe
```

#### Scoop

```
scoop bucket add probe https://github.com/abhimanyu003/scoop-bucket.git
scoop install probe
```

#### Arch Linux

```
yay -S probe-bin
```

#### Snap

> At this point we have to alias while using snap

```
sudo snap install go-probe
sudo snap alias go-probe.probe probe
```

#### Go Install

```
go install github.com/abhimanyu003/probe@latest
```

#### wget

Use wget to download, gzipped pre-compiled binaries.

For instance, `VERSION=v0.0.1` and `BINARY=probe_0.0.1_linux_amd64`

```
wget https://github.com/abhimanyu003/probe/releases/download/${VERSION}/${BINARY}.tar.gz -O - |\
  tar xz && mv probe /usr/bin/probe
```

You can find list of binary and release over [Release Page!](https://github.com/abhimanyu003/probe/releases)

#### Manually

Download the pre-compiled binaries from the [Release!](https://github.com/abhimanyu003/probe/releases) page and copy them
to the desired location.

You can download
* DEB
* RPM
* Pre-compiled binary

Visit [Release!](https://github.com/abhimanyu003/probe/releases)
