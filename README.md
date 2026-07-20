# Subfinder-plus

Subfinder's passive-enumeration engine with selected free-access source behavior adapted from or inspired by BBOT.

[![Go CI](https://github.com/hack-techv2/subfinder-plus/actions/workflows/ci.yml/badge.svg)](https://github.com/hack-techv2/subfinder-plus/actions/workflows/ci.yml)
[![Latest Subfinder-plus release](https://img.shields.io/github/v/release/hack-techv2/subfinder-plus?label=release)](https://github.com/hack-techv2/subfinder-plus/releases/latest)

Subfinder-plus is a HackTech-maintained fork of ProjectDiscovery Subfinder. It preserves Subfinder's CLI and Go architecture while carrying a narrow set of source changes intended to improve useful enumeration without paid service subscriptions. DNSDumpster mirrors BBOT's public web-flow approach; CertSpotter and URLScan use their anonymous endpoints and accept optional keys for higher service limits.

This is not a full merger with BBOT and is not affiliated with or endorsed by ProjectDiscovery or the BBOT project.

> **Update notice:** Update only from this fork's [Subfinder-plus releases](https://github.com/hack-techv2/subfinder-plus/releases). Do not use the inherited `-up` flag: it targets ProjectDiscovery releases, and inherited `-version` output does not uniquely identify a Subfinder-plus release. Use `-duc` to suppress the inherited update check.

## Scope and authorization

Use Subfinder-plus only against domains and systems you own or are explicitly authorized to assess. Passive data sources remain subject to their own terms, quotas, and availability.

## Prerequisites

- Go 1.24.1 for source builds.
- Docker, if building the container image.

## Install

### Download a release

Download `subfinder-plus-windows-amd64.exe` for Windows or `subfinder-plus-linux-amd64` for Linux from the [latest Subfinder-plus release](https://github.com/hack-techv2/subfinder-plus/releases/latest), together with `checksums.txt`. Verify the selected binary against `checksums.txt`, then rename it to match the documented command:

```sh
# Linux
chmod +x subfinder-plus-linux-amd64
mv subfinder-plus-linux-amd64 subfinder-plus
```

```powershell
# Windows PowerShell
Rename-Item subfinder-plus-windows-amd64.exe subfinder-plus.exe
```

Place the renamed file on your `PATH`, then run `subfinder-plus -h`.

### Build from source

```sh
git clone https://github.com/hack-techv2/subfinder-plus.git
cd subfinder-plus
go build -o subfinder-plus ./cmd/subfinder
```

### Build with Docker

```sh
docker build -t subfinder-plus .
docker run --rm subfinder-plus -d example.com -silent
```

## Usage

```sh
subfinder-plus -d example.com -silent
```

Use `subfinder-plus -h` as the authoritative reference for flags available in the release you run. Common options include `-d` for a domain, `-dL` for a domain list, `-o` for output, `-silent` for only discovered names, `-s` to choose sources, and `-es` to exclude sources.

To inspect the current source names before selecting them, run:

```sh
subfinder-plus -ls
```

## Free-access source behavior

The fork's selected free-access behavior covers DNSDumpster, CertSpotter, and URLScan. It does not make every inherited source key-free or remove provider rate limits. See the [free-access source coverage reference](docs/free-source-coverage.md) for the supported behavior and its limits.

## Configuration

Subfinder-plus follows Subfinder's existing configuration layout. By default, it uses `config.yaml` and `provider-config.yaml` in the Subfinder application-config directory; override them with `-config`, `-pc`, `SUBFINDER_CONFIG`, or `SUBFINDER_PROVIDER_CONFIG`.

Optional provider credentials belong in the provider configuration. Supplying them can raise limits where a provider offers that capability; it does not change the authorization requirements for using the service.

## Go library

This fork preserves the existing ProjectDiscovery Subfinder Go module path and package layout for compatibility. Build applications against the module requirements in `go.mod`, and review [examples/main.go](examples/main.go) for the current library usage pattern.

## Documentation

- [Free-access source coverage](docs/free-source-coverage.md)
- [Fork maintenance](docs/fork-maintenance.md)
- [Security policy](SECURITY.md)

## Contributing

Contributions should preserve Subfinder's CLI compatibility and keep the maintained fork delta narrow. See [CONTRIBUTING.md](CONTRIBUTING.md) for setup, validation, and pull-request guidance.

## Security

Report suspected vulnerabilities according to [SECURITY.md](SECURITY.md). Do not disclose credentials, target data, or exploit details in public issues.

## Attribution

Subfinder-plus builds on [ProjectDiscovery Subfinder](https://github.com/projectdiscovery/subfinder). Selected public-source behavior was adapted from or inspired by [BBOT](https://github.com/blacklanternsecurity/bbot) as described in the [fork maintenance reference](docs/fork-maintenance.md). See [THANKS.md](THANKS.md) for additional attribution details.

## License

Subfinder-plus is distributed under the repository's [MIT License](LICENSE.md). Review [DISCLAIMER.md](DISCLAIMER.md) before using third-party sources.
