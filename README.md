# InitVariables [![Build Action](https://github.com/my20002000/InitVariables/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/my20002000/InitVariables/actions/workflows/go.yml)

## Setup

``` bash
go install github.com/akavel/rsrc@latest
go install mvdan.cc/garble@latest
go install github.com/burrowers/garble@latest
```

## Build

``` bash
rsrc -manifest uac.manifest -ico icon.ico -o uac.syso
go build -ldflags="-s -w " -trimpath 
garble -tiny build
```

``` bash
go build -ldflags "-s -w -H windowsgui -X main.var1=a -X main.var2=b"
CC=musl-gcc go build -tags musl -ldflags="-extldflags --static"
```

## upx

```
upx --best
upx --ultra-brute 
```

