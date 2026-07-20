# Fork maintenance

## Ownership and scope

Subfinder supplies the engine, CLI, packages, and upstream update stream. BBOT informed selected public-source access behavior; DNSDumpster is the explicit code-level adaptation. Subfinder-plus owns only the maintained delta, tests, release tags, and fork documentation.

This scope preserves compatibility with the upstream architecture while documenting the small, independently maintained source-behavior delta.

## Migration provenance

- Former patch source: `dabsterz/subfinder-plus`
- Imported behavior commit: `8ac9353a78565b4407f8dde949c503c5cb90781e`
- Imported documentation commit: `9afcef859e723f85bbbb2e1f3038f860c369ac44`
- Initial upstream base: `d0ea1029cf87ff965804fc399d12e4502c7436b2`.

The initial upstream base was taken from `upstream/dev`.

## Maintained behavior

- CertSpotter and URLScan can operate anonymously and use configured keys only for authenticated or rate-limited access.
- DNSDumpster uses its public short-lived web flow without accepting a provider API key.

## Release policy

Subfinder-plus uses independent `subfinder-plus-vX.Y.Z` releases. Each release records its upstream base commit and publishes:

- `subfinder-plus-windows-amd64.exe`
- `subfinder-plus-linux-amd64`

## Upstream synchronization

1. Fetch `upstream`.
2. Create `upstream-sync/YYYY-MM-DD` from the team default branch.
3. Merge the upstream default branch without rewriting published team history.
4. Resolve conflicts while retaining only patches that remain necessary.
5. Run all Go tests, vet, build, and source-specific regression tests.
6. Merge through a reviewed pull request.
7. Publish a new independent release only when RADAR should adopt the update.
