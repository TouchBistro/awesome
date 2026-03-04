# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**awesome** (AWS Config & Clients Provider for Go) is a lightweight wrapper on AWS SDK for Go v2 that eliminates boilerplate for AWS credential configuration and client initialization. It supports 386 auto-generated service client wrappers.

## Commands

```bash
make setup      # Install dev dependencies (golangci-lint)
make build      # Build all packages: go build ./...
make lint       # Run golangci-lint
make test       # Run tests with HTML coverage report (coverage/coverage.html)
make test-ci    # Run tests with stdout coverage output
make clean      # Remove build artifacts and generated code
```

Run a single test:
```bash
go test ./providers/... -run TestFunctionName -v
```

Regenerate client wrappers (not automatic):
```bash
cd codegen && go generate codegen.go
```

## Architecture

### Core Abstractions

**Provider** (`providers/`) — encapsulates `aws.Config` behind a named key. Four implementations:
- `DefaultCredsProvider` — AWS SDK default credentials chain (env → shared config → ECS/EC2 role)
- `EnvironmentCredsProvider` — static credentials from env vars (customizable var names)
- `SharedConfigCredsProvider` — reads `~/.aws/config` and `~/.aws/credentials`
- `AssumeRoleCredsProvider` — wraps another provider and assumes an IAM role

**Provider Registry** (`providers/map.go`) — thread-safe (sync.Mutex) named registry. Key functions: `Get()`, `MustGet()`, `Default()`, `MustDefault()`, `Clone()`, `MustClone()`.

**Client Wrappers** (`clients/_<service>/client.go`) — auto-generated per-AWS-service packages. Each provides:
- `Client(provider, optFns)` — creates/returns cached singleton (sync.Map per provider key)
- `Must(provider, optFns)` — panics on error
- `Delete(provider)` / `Refresh(provider, optFns)` — cache management

**Auto-initialization** (`aws-ccp.go`) — blank-importing the root package runs an `init()` that auto-initializes the `default` provider from command-line flags (via `init/cmd/`).

### Key Design Decisions

- **Package naming**: Client packages use underscore prefix (`_ec2`, `_s3`) to avoid conflicts with AWS SDK package names
- **Binary size**: Each client package imports only its specific AWS service (full SDK import would balloon from ~8MB to ~68MB)
- **Auto-generated code**: All files under `clients/` are generated — do not edit manually
- **Functional options**: All provider constructors use `CredsProviderOptionsFunc` pattern (see `providers/provider_creds_options.go`)

### Configuration Precedence (EnvironmentCredsProvider region)
1. `WithRegionFrom()` env var
2. `AWS_REGION`
3. `WithDefaultRegion()` / `WithRegion()` option
4. `DefaultAWSRegion` constant (`us-east-1`)

### Typical Usage Pattern

```go
// Blank import for auto-initialization of 'default' provider
import _ "github.com/TouchBistro/awesome"
import "github.com/TouchBistro/awesome/clients/_ec2"

// Use the default provider (initialized from CLI flags)
client, err := _ec2.Client(providers.DefaultProviderKey)

// Or create a named provider explicitly
provider, err := providers.NewAssumeRoleCredsProvider(ctx, "prod",
    providers.WithRoleArn("arn:aws:iam::123:role/MyRole"),
    providers.WithBaseProvider(baseProvider),
)
```

## Code Generation

The codegen tool (`codegen/codegen.go`) uses the GitHub API to fetch AWS SDK service package names and generates `client.go` files. Run it from the `codegen/` directory. The generated files follow a strict template — see `codegen/README.md` for details on include/exclude lists and the generation process.
