# headlesslab/.github

Organisation-wide defaults for [headlesslab](https://github.com/headlesslab) repositories.

[![Smoke](https://github.com/headlesslab/.github/actions/workflows/smoke.yml/badge.svg)](https://github.com/headlesslab/.github/actions/workflows/smoke.yml)

| Path | What it is |
| --- | --- |
| [`SECURITY.md`](SECURITY.md) | The security policy GitHub shows on every headlesslab repository that has no `SECURITY.md` of its own. |
| [`.github/workflows/go.yml`](.github/workflows/go.yml) | The reusable CI workflow for a Go module. The Satellite modules of [wand](https://github.com/headlesslab/wand) call it. |
| [`.golangci.yml`](.golangci.yml) | The default golangci-lint configuration the workflow applies: upstream go-rod's, migrated to the v2 schema. |
| [`smoke/`](smoke) | A tiny Go module that runs through the workflow on every push and pull request, so a broken `go.yml` is caught here first. |

## Reusable Go workflow

`go.yml` is a `workflow_call` workflow. On the caller's module it runs:

| Job | Runner | Go | What |
| --- | --- | --- | --- |
| `test` | `ubuntu-latest`; with `cross-platform: true` also `windows-latest` and `macos-latest` | `floor` and `stable` on `ubuntu-latest`, `stable` on the other runners | `go build ./...`, `go vet ./...`, `go test -race -count=1 -covermode=atomic -coverprofile=coverage.out ./...`. The `ubuntu-latest` / `stable` cell also runs the coverage ratchet: total statement coverage below `coverage-threshold` fails the job. |
| `lint` | `ubuntu-latest` | `stable` | golangci-lint at a version pinned in `go.yml`, with the module's own `.golangci.yml` when it has one and the default configuration in this repository otherwise. |
| `govulncheck` | `ubuntu-latest` | `stable` | `govulncheck ./...` with `golang.org/x/vuln` at a version pinned in `go.yml`. Stable only: Go 1.21's standard library would report its own unpatched vulnerabilities forever. |

`floor` is the module's Go floor: the version its `go.mod` declares, at its latest patch release (`go 1.21` runs on the latest Go 1.21.x). `stable` is the current Go release. Every job runs with `GOTOOLCHAIN=local`, so a dependency that raises its Go floor above the module's fails the run instead of downloading a toolchain. The workflow declares `permissions: contents: read`, checks out without persisted credentials, and pins every action to a full-length commit SHA with the version in a trailing comment.

### Inputs

| Input | Type | Default | Meaning |
| --- | --- | --- | --- |
| `coverage-threshold` | number | required | Minimum total statement coverage in percent, measured on the `ubuntu-latest` / `stable` job. Set it to the module's current coverage (for a snapshot-imported Satellite, the level it was imported at) and only ever raise it. |
| `cross-platform` | boolean | `false` | Also run the `test` job on `windows-latest` and `macos-latest`, on Go `stable`. |
| `working-directory` | string | `.` | Directory of the Go module, relative to the repository root. |
| `godebug` | string | empty | `GODEBUG` for the test binaries on every `test` cell, e.g. `tracebackancestors=1000` for a suite built on [leakcheck](https://github.com/headlesslab/leakcheck). A called workflow does not inherit the caller's `env`, so the variable travels as an input. |

### Calling it

Pin the workflow to a full commit SHA of this repository's `main` and leave the branch name in the comment. Dependabot's `github-actions` updates move a SHA pinned to a branch commit to the head of that branch, so the pin stays current without a floating reference.

```yaml
name: CI

on:
  push:
    branches: [main]
  pull_request:

permissions:
  contents: read

concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true

jobs:
  go:
    uses: headlesslab/.github/.github/workflows/go.yml@<full commit sha> # main
    with:
      coverage-threshold: 90
      # cross-platform: true   # only where OS behaviour matters (file locks, renames)
      # godebug: tracebackancestors=1000   # only for a suite built on leakcheck
```

The resulting check names are `go / test (ubuntu-latest, floor)`, `go / test (ubuntu-latest, stable)`, `go / lint` and `go / govulncheck`, plus `go / test (windows-latest, stable)` and `go / test (macos-latest, stable)` with `cross-platform: true`. Those are the Gates a Satellite's `main` ruleset requires; the same settings script passes them as `-check` flags.

### golangci-lint configuration

The default is [`.golangci.yml`](.golangci.yml): upstream go-rod's configuration run through `golangci-lint migrate`, with `default: all` and upstream's disabled list, so linters added to golangci-lint since are on. A module that needs different rules commits its own `.golangci.yml` (or `.yaml`, `.toml`, `.json`) in its module directory and the workflow uses that instead. The golangci-lint version is pinned in `go.yml`; bumping it is a change to this repository and shows up in the smoke run first.

## Dependabot template

Each Satellite carries this `.github/dependabot.yml`: weekly version updates for the SHA-pinned actions and the reusable-workflow pin, and for Go modules security updates only.

```yaml
version: 2
updates:
  # Keep every SHA-pinned action and the reusable-workflow pin moving.
  - package-ecosystem: github-actions
    directory: /
    schedule:
      interval: weekly

  # Go modules: security updates only. A limit of 0 disables version-update
  # pull requests; security updates are not subject to the limit.
  - package-ecosystem: gomod
    directory: /
    schedule:
      interval: weekly
    open-pull-requests-limit: 0
```

Dependabot security updates themselves are a repository setting, switched on by wand's repository settings script ([`internal/tools/repo-settings`](https://github.com/headlesslab/wand/tree/main/internal/tools/repo-settings), run as documented in wand's [maintainer notes](https://github.com/headlesslab/wand/blob/main/docs/maintainer-notes.md)) rather than by this file. Version bumps of Go dependencies arrive as hand pull requests.

This repository's own [`.github/dependabot.yml`](.github/dependabot.yml) is the same file with the `gomod` directory set to `/smoke`.
