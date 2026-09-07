# modctl

[![CI](https://github.com/modelpack/modctl/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/modelpack/modctl/actions/workflows/ci.yml)
[![GoDoc](https://godoc.org/github.com/modelpack/modctl?status.svg)](https://godoc.org/github.com/modelpack/modctl)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/modelpack/modctl)

> **This is a ZonHoren-maintained fork of [modelpack/modctl](https://github.com/modelpack/modctl).**
> This `main` branch is an untouched, Dependabot-tracked mirror of upstream
> and does **not** carry ZonHoren's patch. The actual fix — upstream's
> `Pull`, `Fetch`, and their Dragonfly-accelerated counterparts
> (`pullByDragonfly`/`fetchByDragonfly`) parsed a reference's digest but only
> ever passed `ref.Tag()` to `FetchReference()`, so a pure digest reference
> (e.g. `registry.example.com/repo@sha256:...`, no tag) resolved to an empty
> tag string and failed every time with `"invalid reference"` — `Remove`
> (`pkg/backend/rm.go`) already handled this correctly — lives on the
> [`zonhoren-patches`](https://github.com/zonhoren/modctl/tree/zonhoren-patches)
> branch (commits `f77e496`, `7a55d32`; tags `v0.1.2-alpha.0-zonhoren.1`,
> `v0.1.2-alpha.0-zonhoren.2`). Consume that branch or those tags, not
> `main`. This fork exists so that
> [`zonhoren/model-csi-driver`](https://github.com/zonhoren/model-csi-driver)
> can pin it via a `go.mod` `replace` directive to mount OCI model artifacts
> referenced purely by digest.

Modctl is a user-friendly CLI tool for managing OCI model artifacts, which are bundled based on [Model Spec](https://github.com/modelpack/model-spec).
It offers commands such as `build`, `pull`, `push`, and more, making it easy for users to convert their AI models into OCI artifacts.

## Documentation

You can find the full documentation on the [getting started](./docs/getting-started.md).

## Copyright

Copyright © contributors to ModelPack, established as ModelPack a Series of LF Projects, LLC.

## LICENSE

Apache 2.0 License. Please see [LICENSE](LICENSE) for more information.
