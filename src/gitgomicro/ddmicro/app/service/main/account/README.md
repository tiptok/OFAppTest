# Account Service

This is the Account service

Generated with

```
micro new account
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.dmicro.srv.account
- Type: srv
- Alias: account

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./account-srv
```

Build a docker image
```
make docker
```

## Start

1.Start micro api

```
micro api --handler=http
```

2.Start go run ./account/main.go

3.Start go run ./push/main.go


