# CODEBUDDY.md This file provides guidance to CodeBuddy when working with code in this repository.

## Common Commands

### Build
```bash
go build .
```
Compiles the package to verify syntax and type correctness. This is a library package, so it produces no binary output.

### Run Tests
```bash
# Run all tests (requires Zabbix server)
export TEST_ZABBIX_URL=http://localhost:8080/zabbix/api_jsonrpc.php
export TEST_ZABBIX_USER=Admin
export TEST_ZABBIX_PASSWORD=zabbix
export TEST_ZABBIX_VERBOSE=1
go test -v

# Run unit tests only (no Zabbix server required)
go test -v -run TestVersionManager
go test -v -run TestFeatureDetection
go test -v -run TestHeaderConversion
go test -v -run TestItemValidation
go test -v -run TestHostValidation
go test -v -run TestBrowserItemValidation
go test -v -run TestAdapterInterface
go test -v -run TestFeatureConstants

# Run a specific test
go test -v -run TestHosts

# Run with coverage
go test -v -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

### Code Quality
```bash
# Format code
gofmt -w .

# Vet code for issues
go vet ./...
```

### Multi-version Test Suite
```bash
./run_multiversion_tests.sh
```
Runs comprehensive unit tests, integration tests, benchmarks, and code checks for Zabbix 6.0/7.0 multi-version support.

## Architecture Overview

This is a Go client library for the Zabbix monitoring API. It provides bindings to interoperate between Go programs and Zabbix servers via JSON-RPC API calls.

### Core Architecture Components

**API Client (`base.go`)**
The `API` struct is the central client that manages authentication, HTTP communication, and version detection. It maintains an auth token after login and provides methods like `CallWithError()` and `CallWithErrorParse()` for making JSON-RPC requests. The client handles request/response serialization, TLS configuration, and optional compression for Zabbix 6.0+.

**Version Management System**
The library implements a sophisticated multi-version support system through `VersionManager` (in `base.go`). Upon login, it automatically detects the Zabbix server version and initializes appropriate adapters. Feature constants like `FeatureHistoryPush`, `FeatureMFA`, and `FeatureBrowserItem` define version-specific capabilities. Methods like `IsZabbix7()`, `IsZabbix6()`, and `SupportsFeature()` allow runtime version checking.

**Adapter Pattern for Multi-version Support**
To handle API differences between Zabbix 6.0 and 7.0, the library uses adapter interfaces:
- `ItemAdapter` with implementations `Zabbix6ItemAdapter` and `Zabbix7ItemAdapter` (in `item.go`)
- `HostAdapter` with implementations `Zabbix6HostAdapter` and `Zabbix7HostAdapter` (in `host.go`)

These adapters automatically convert data formats between versions. For example, HTTP headers changed from a map format in 6.0 (`HeadersV6`) to an array of structs in 7.0 (`HeadersV7`). The adapters handle bidirectional conversion transparently.

### Data Model Structure

**Resource Types**
Each Zabbix API resource has a corresponding Go file with struct definitions and CRUD operations:
- `host.go`: Host struct with status, availability, inventory, proxy configuration
- `item.go`: Item struct with types (ZabbixAgent, HTTPAgent, Browser, etc.), value types, preprocessing
- `trigger.go`: Trigger struct with expressions, priorities, dependencies
- `template.go`: Template struct for host templates
- `host_group.go`: HostGroup for organizing hosts
- `proxy.go`, `proxy_group.go`: Proxy and ProxyGroup for distributed monitoring
- `user.go`, `mfa.go`: User management and multi-factor authentication
- `lld.go`: Low-level discovery rules
- `mediatype.go`, `alert.go`: Alerting and notification media

**Type Constants**
Each resource file defines typed constants for enumerated values. For example, `item.go` defines `ItemType` constants like `ZabbixAgent`, `HTTPAgent`, `Browser` (Zabbix 7.0+), and `ValueType` constants like `NumericFloat`, `Text`, `Log`. These provide type safety and IDE autocomplete.

### Key Implementation Patterns

**CRUD Wrapper Methods**
Each resource follows a consistent naming pattern:
- `ResourcesGet(params Params)` - Get resources with filter parameters
- `ResourceGetByID(id string)` - Get single resource by ID
- `ResourcesCreate(resources)` - Create new resources
- `ResourcesUpdate(resources)` - Update existing resources
- `ResourcesDelete(resources)` - Delete resources
- `ResourcesDeleteByIds(ids []string)` - Delete by IDs

The `Params` type is a `map[string]interface{}` used for API request parameters. Wrapper methods set default values like `output: "extend"` if not provided.

**Error Handling**
The library defines custom error types: `Error` for Zabbix API errors (with Code, Message, Data), `ExpectedOneResult` when queries return unexpected counts, and `ExpectedMore` for batch operations. All API methods return `(result, error)` tuples.

**JSON-RPC Communication**
The `request` and `Response`/`RawResponse` structs handle JSON-RPC 2.0 protocol. `CallWithError()` makes requests and returns raw JSON responses. `CallWithErrorParse()` unmarshals results into provided structs. Request IDs are atomically incremented for concurrency safety.

**Version-specific Field Handling**
Structs include version-prefixed fields for format differences:
```go
type Item struct {
    HeadersV6 HttpHeaders   `json:"headers_v6,omitempty"`  // Zabbix 6.0 map format
    HeadersV7 []HeaderField `json:"headers_v7,omitempty"`  // Zabbix 7.0 array format
}
```

Adapters automatically convert between formats based on the detected server version.

### Test Structure

**Test Files**
- `base_test.go`: Test harness with `getAPI()` helper for integration tests
- `*_test.go`: Resource-specific tests (host_test.go, item_test.go, etc.)
- `multiversion_test.go`: Unit tests for version management, adapters, and format conversion
- `zabbix6_test.go`: Zabbix 6.0 specific feature tests

**Integration Test Setup**
Integration tests require environment variables to connect to a real Zabbix server. Tests are skipped if `TEST_ZABBIX_URL` is not set. The test harness in `base_test.go` creates a shared API client that's reused across tests.

### Important Notes for Development

- **No external dependencies**: The library only uses Go standard library plus `github.com/AlekSi/reflector` for reflection utilities.
- **Go 1.12+**: Minimum Go version as specified in go.mod.
- **API Compatibility**: Tested on Zabbix 3.2, 3.4, 4.0, 4.2, 4.4, 5.0, 6.0, and 7.0.
- **Compression Support**: Enable with `Config.EnableCompression` for Zabbix 6.0+.
- **Zabbix 7.0 Features**: MFA, Proxy Groups, Browser Items, and History Push API are only available when connected to Zabbix 7.0+.
- **Header Format Conversion**: Use `ConvertHeadersToV7()` and `ConvertHeadersToV6()` for manual format conversion when needed.
