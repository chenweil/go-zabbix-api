# Repository Guidelines

## Project Structure & Module Organization
Core library code lives at the repository root as package `zabbix` (for example `base.go`, `host.go`, `item.go`, `trigger.go`).  
Tests are colocated with sources as `*_test.go`, with broader integration coverage in files like `integration_test.go`, `multiversion_test.go`, and `zabbix6_test.go`.  
Executable entrypoints are under `cmd/`:
- `cmd/zabbix-api-server` for HTTP service endpoints
- `cmd/zabbix-cli` for command-line API access

Supporting docs are in `docs/`, and runnable examples are in `examples/`.

## Build, Test, and Development Commands
- `go build ./...` — compile all packages and command binaries.
- `go test ./...` — run the full Go test suite.
- `go test -v -run TestZabbix6` — run Zabbix 6.x focused tests.
- `./run_multiversion_tests.sh` — run curated multi-version/unit/integration checks.
- `go vet ./...` — static checks for suspicious code.
- `gofmt -w .` — format all Go files before commit.

Run local API server:
```bash
cd cmd/zabbix-api-server
go build -o zabbix-api-server .
PORT=8080 ./zabbix-api-server
```

## Coding Style & Naming Conventions
Use standard Go style and keep code `gofmt`-clean (tabs/formatting handled automatically).  
Follow Go naming conventions:
- Exported identifiers: `PascalCase` (`NewAPI`, `HostGroupsGet`)
- Internal helpers: `camelCase`
- Test files: `*_test.go` and descriptive test names (`TestVersion`, `TestHostValidation`)

Prefer extending existing resource patterns (`Host`, `Item`, `Trigger`) rather than introducing new abstractions.

## Testing Guidelines
Integration tests require environment variables:
- `TEST_ZABBIX_URL`
- `TEST_ZABBIX_USER`
- `TEST_ZABBIX_PASSWORD`
- optional `TEST_ZABBIX_VERBOSE=1`

Use a non-production Zabbix instance; tests are designed to be safe but still perform real API operations.

## Commit & Pull Request Guidelines
Recent history uses short prefixed subjects such as `feat: ...`, `merge: ...`, and operational updates.  
Keep commit titles imperative and scoped (example: `fix: handle empty auth token in LoginWithToken`).

For pull requests, include:
- concise problem/solution summary
- linked task or issue ID (if available)
- test evidence (commands + key output)
- notes on env/config impact for `dev/test/prod`

## Security & Configuration Tips
Never commit real credentials or server URLs with embedded secrets.  
Use environment variables for authentication and keep local overrides out of version control.
