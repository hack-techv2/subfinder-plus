# Contributing to Subfinder-plus

Thank you for helping maintain Subfinder-plus. Keep changes focused, preserve the existing command-line interface, and make the fork delta easy to review.

## Before opening a change

- Work from the `dev` branch and open pull requests against `dev`.
- Check whether the change affects the maintained fork delta or should be proposed upstream to ProjectDiscovery Subfinder.
- Do not include credentials, API keys, target data, or other secrets in commits, fixtures, test output, or logs.

## Development setup

Install Go 1.24.1, clone the repository, and work from the `dev` contribution target:

```powershell
git clone https://github.com/hack-techv2/subfinder-plus.git
Set-Location subfinder-plus
git switch dev
```

Use `go run ./cmd/subfinder -h` to confirm that the development checkout builds and exposes the expected command-line interface.

## Validate a change

Run the following commands from the repository root before opening a pull request:

```powershell
go test ./...
go vet ./...
go build ./...
pwsh -NoProfile -File scripts/test-workflow-contract.ps1
```

## Source behavior changes

Add or update tests for every source behavior change. Document provider-facing assumptions, rate-limit behavior, and any anonymous-versus-keyed behavior that affects the result. Keep fixtures and logs free of credentials, API keys, and target data.

Explain in the pull request whether the source change belongs upstream in ProjectDiscovery Subfinder or is intentionally maintained as part of this fork's narrow delta.

## Pull requests

Open focused pull requests against `dev`. Include a clear description of the behavior change, the validation you ran, and any compatibility impact. Keep unrelated formatting, generated files, and dependency changes out of the pull request unless they are required for the change.

## Upstream synchronization

When syncing with ProjectDiscovery Subfinder, preserve the fork's documented behavior and resolve conflicts deliberately. Record why retained changes belong in Subfinder-plus rather than upstream, and avoid expanding the fork delta without a concrete maintenance reason.
