# Repository Documentation Cleanup Design

## Objective

Present Subfinder-plus as a professional HackTech-maintained fork whose purpose, free-source scope, upstream relationship, installation path, and maintenance policy are immediately clear.

The cleanup must preserve Subfinder's Go package layout and keep future upstream synchronization practical.

## Project Positioning

Subfinder-plus uses ProjectDiscovery Subfinder as its engine, command-line interface, and passive-enumeration architecture. It adds selected free-access source behavior adapted from or inspired by BBOT.

Documentation must not describe the repository as a wholesale codebase merger or imply endorsement by or affiliation with ProjectDiscovery or the BBOT project. It must distinguish these source behaviors:

- DNSDumpster follows the public HTMX and short-lived JWT workflow used by the website, mirroring BBOT's approach. It requires no user-supplied API key.
- CertSpotter supports anonymous access and accepts an optional API key for higher service limits.
- URLScan supports anonymous access and accepts an optional API key for higher service limits.
- Other sources inherited from Subfinder retain their own credential and service requirements.

The phrase "free-access scope" means that this fork carries only selected behavior available without a paid service subscription. It does not promise that every inherited source works without credentials, has unlimited capacity, or will remain publicly available.

## Documentation Architecture

### Root README

Rewrite `README.md` as the entry point for users. It will contain:

1. Subfinder-plus identity and a one-sentence description.
2. A concise explanation of the Subfinder and BBOT relationship.
3. Meaningful fork-owned CI and release badges only.
4. Authorization and third-party-service notice.
5. Prerequisites and copy-paste installation options.
6. A minimal working command before advanced usage.
7. A summary of free-access source behavior.
8. Configuration paths and API-key behavior.
9. Links to focused reference and maintenance documents.
10. Contribution, security, attribution, disclaimer, and license links.

The README will not embed the full CLI help output because that content drifts as flags change. It will direct readers to `subfinder -h` for authoritative flag reference.

### Focused Documents

Create or normalize these documents:

- `docs/free-source-coverage.md`: factual reference for the fork's selected free-access source behavior, limitations, optional credentials, and inherited-source boundary.
- `docs/fork-maintenance.md`: provenance, versioning, release contents, and upstream synchronization procedure currently stored in `HACKTECH_CHANGES.md`.
- `CONTRIBUTING.md`: concise development setup, validation commands, branch target, documentation expectations, and upstream-sync considerations.
- `SECURITY.md`: private vulnerability-reporting guidance if a confirmed private channel exists; otherwise it will direct reporters to a clearly identified repository-owner contact without inventing an address.

Retain `DISCLAIMER.md`, `LICENSE.md`, and `THANKS.md` at the root because release automation packages them and users expect legal and attribution files there.

## Repository Layout

Keep the existing Go-standard source layout:

```text
cmd/                 CLI entry point
examples/            Go library example
pkg/                 reusable packages and source implementations
scripts/             repository validation scripts
static/              README images inherited from upstream
docs/                fork-specific reference and maintenance documentation
.github/              CI, releases, and contribution templates
```

Do not rename `cmd`, `pkg`, `examples`, or `static`. The small cosmetic benefit would not justify extra divergence from upstream.

Remove `HACKTECH_CHANGES.md` after migrating all accurate content and updating every reference. Consolidate the two overlapping bug-report issue templates into one fork-owned template.

## GitHub-Facing Metadata

- Replace upstream badges with CI and release links for `hack-techv2/subfinder-plus`.
- Update issue and pull-request templates to name Subfinder-plus and link to the fork.
- Preserve the `dev` branch as the contribution target.
- Remove upstream Discord and discussion links where they incorrectly imply that ProjectDiscovery supports this fork.
- Keep direct links to upstream documentation only when explicitly labeled as upstream Subfinder documentation.

## Verification

The implementation is complete only after these checks pass:

1. Search for stale fork-facing ProjectDiscovery issue, release, discussion, and badge links.
2. Verify every relative Markdown link and local image target.
3. Confirm release automation reads the migrated maintenance document.
4. Run `go test ./...`.
5. Run `go vet ./...`.
6. Run `go build ./...`.
7. Run `scripts/test-workflow-contract.ps1`.
8. Review `git diff` for accidental code or generated-artifact changes.

Live passive-source tests remain opt-in because they depend on third-party network services and are not required for a documentation cleanup.

## Non-Goals

- Changing enumeration behavior or adding sources.
- Claiming feature parity with BBOT.
- Removing inherited Subfinder sources that require credentials.
- Renaming the Go module, CLI package, or internal imports.
- Publishing a release or changing the release version.
- Broad source-code or test reorganization.
