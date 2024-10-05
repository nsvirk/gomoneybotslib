# Library for Moneybots Apps

## Description

This library provides a set of packages to facilitate the development of bots for the Moneybots platform. It includes functionality for connecting to the Moneybots API, logging events, and maintaining bot state.

It has three packages:

- `connect`: for connecting to the Moneybots API
- `logger`: for logging events
- `state`: for maintaining bot state

## Install

```sh
go get github.com/nsvirk/gomoneybotslib
```

## Package Imports

```go
mbconnect "github.com/nsvirk/gomoneybotslib/pkg/connect"
mblogger "github.com/nsvirk/gomoneybotslib/pkg/logger"
mbstate "github.com/nsvirk/gomoneybotslib/pkg/state"
```

## Examples

```sh
go run examples/connect/main.go
go run examples/logger/main.go
go run examples/state/main.go
```
