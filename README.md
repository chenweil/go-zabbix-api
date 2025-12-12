# Go zabbix api

Note, this is not tested and is adjusted for use of tpretz/terraform-provider-zabbix

[![GoDoc](https://godoc.org/github.com/tpretz/go-zabbix-api?status.svg)](https://godoc.org/github.com/tpretz/go-zabbix-api) [![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE) [![Build Status](https://travis-ci.org/tpretz/go-zabbix-api.svg?branch=master)](https://travis-ci.org/tpretz/go-zabbix-api)

This Go package provides access to Zabbix API.

Tested on Zabbix 3.2, 3.4, 4.0, 4.2, 4.4, 5.0 and **6.0** with full compatibility support.

This package aims to support multiple zabbix resources from its API like trigger, application, host group, host, item, template..

### Zabbix 6.0 Support

This version includes comprehensive support for Zabbix 6.0 API features:

- **UUID Support**: All major API objects now support UUID fields
- **Enhanced Authentication**: Token-based authentication with `LoginWithToken()` and `CheckAuthentication()`
- **HTTP Agent Improvements**: Optional `interfaceid` parameter for HTTP Agent items
- **Extended Item Types**: New calculated item value types (Text, Log, Character)
- **User Permission Updates**: Updated default fields for `user.get`, `mediatype.get`, and `alert.get`
- **HTTP Method Extensions**: Support for PATCH, HEAD, OPTIONS, TRACE, CONNECT methods
- **Compression Support**: Built-in support for gzip and deflate compression
- **Extended Field Lengths**: URL fields now support up to 2048 characters

## Install

Install it: `go get github.com/tpretz/go-zabbix-api`

## Getting started

### Basic Usage

```go
package main

import (
	"fmt"

	"github.com/tpretz/go-zabbix-api"
)

func main() {
	user := "MyZabbixUsername"
	pass := "MyZabbixPassword"
	api := zabbix.NewAPI("http://localhost/api_jsonrpc.php")
	api.Login(user, pass)

	res, err := api.Version()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to zabbix api v%s\n", res)
}
```

### Zabbix 6.0 Enhanced Features

#### Compression Support

```go
// Enable compression for better performance
config := zabbix.Config{
    Url:               "https://zabbix.example.com/api_jsonrpc.php",
    EnableCompression: true,
    AcceptedEncodings: []string{"gzip", "deflate", "identity"},
}

api := zabbix.NewAPI(config)
```

#### Enhanced Authentication

```go
// Login with token support (Zabbix 6.0)
auth, err := api.LoginWithToken("username", "password", "optional-token")

// Check authentication validity
valid, err := api.CheckAuthentication("token")
```

#### HTTP Agent with New Methods

```go
// Create HTTP Agent item with new Zabbix 6.0 methods
item := zabbix.Item{
    Type:         zabbix.HTTPAgent,
    HostID:       "host123",
    Key:          "http.api.test",
    Name:         "API Test",
    Url:          "https://api.example.com/endpoint",
    RequestMethod: zabbix.HTTPMethodPATCH, // New in Zabbix 6.0
    // InterfaceID is optional for HTTP Agent in Zabbix 6.0
}
```

#### UUID Support

```go
// Objects now include UUID fields
host := zabbix.Host{
    HostID: "12345",
    UUID:   "12345678-1234-1234-1234-123456789abc", // Zabbix 6.0
    Host:   "example.com",
}
```

## Migration Guide

### Upgrading from Zabbix 5.x to 6.0

The library maintains backward compatibility with Zabbix 5.x and earlier versions. However, when upgrading to Zabbix 6.0, consider these changes:

1. **UUID Fields**: New UUID fields are added to all major objects but are optional
2. **HTTP Agent Items**: `interfaceid` is now optional for HTTP Agent type items
3. **Authentication**: New token-based authentication methods are available
4. **Compression**: Enable compression for better performance in Zabbix 6.0 environments
5. **HTTP Methods**: Additional HTTP methods (PATCH, HEAD, OPTIONS, TRACE, CONNECT) are now supported

### Example Migration

```go
// Old way (still works)
api := zabbix.NewAPI("http://zabbix/api_jsonrpc.php")

// New way with Zabbix 6.0 enhancements
config := zabbix.Config{
    Url:               "http://zabbix/api_jsonrpc.php",
    EnableCompression: true,  // Enable for Zabbix 6.0
    TlsNoVerify:       false, // Keep existing TLS settings
}
api := zabbix.NewAPI(config)
```

## Tests

### Considerations

You should run tests before using this package.
Zabbix API doesn't match documentation in few details, which are changing in patch releases. 

Tests are not expected to be destructive, but you are advised to run them against not-production instance or at least make a backup.
For a safer and more accurate testing we advice to run tests with following minimum versions which implements strict validation of valuemap for `get` method:

- 4.0.13rc1 [6ead4fd7865](https://git.zabbix.com/projects/ZBX/repos/zabbix/commits/6ead4fd7865f24ba1246832caa867d33ee9773ba)
- 4.2.7rc1 [a1d257bf6a3](https://git.zabbix.com/projects/ZBX/repos/zabbix/commits/a1d257bf6a3972e24a0044aa019d120eaf7a211a)
- 4.4.0alpha3 [db94d75b4bf](https://git.zabbix.com/projects/ZBX/repos/zabbix/commits/db94d75b4bf5bfc72df3e01cd5fd4a57bc3784e3)
- **6.0.x** (Latest stable version recommended for full feature support)

For more information, please see issues [ZBX-3783](https://support.zabbix.com/browse/ZBX-3783) and [ZBX-3685](https://support.zabbix.com/browse/ZBX-3685)

### Zabbix 6.0 Testing

For testing Zabbix 6.0 specific features, use the dedicated test suite:

```bash
# Run Zabbix 6.0 specific tests
go test -v -run TestZabbix6

# Run performance benchmarks
go test -v -run TestZabbix6PerformanceBenchmarks

# Run compression tests
go test -v -run TestZabbix6Compression
```

### Run tests

```bash
export TEST_ZABBIX_URL=http://localhost:8080/zabbix/api_jsonrpc.php
export TEST_ZABBIX_USER=Admin
export TEST_ZABBIX_PASSWORD=zabbix
export TEST_ZABBIX_VERBOSE=1
go test -v
```

`TEST_ZABBIX_URL` may contain HTTP basic auth username and password: `http://username:password@host/api_jsonrpc.php`. Also, in some setups URL should be like `http://host/zabbix/api_jsonrpc.php`.

## References

Documentation is available on [godoc.org](https://godoc.org/github.com/tpretz/go-zabbix-api).
Also, Rafael Fernandes dos Santos wrote a [great article](http://www.sourcecode.net.br/2014/02/zabbix-api-with-golang.html) about using and extending this package.

License: Simplified BSD License (see [LICENSE](LICENSE)).
