# Repository Documentation Cleanup Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make Subfinder-plus a clean, professional fork whose documentation accurately explains its Subfinder foundation, selected BBOT-inspired free-access behavior, installation, maintenance, and contribution model.

**Architecture:** Keep the Go-standard source tree unchanged and organize fork-owned documentation under `docs/`. Use the root README as a concise landing page, focused reference files for details, fork-owned GitHub templates for contributor interactions, and repository scripts for repeatable validation.

**Tech Stack:** Markdown, Go 1.24.1, GitHub Actions YAML, PowerShell 7-compatible validation scripts

## Global Constraints

- Preserve Subfinder's `cmd/`, `pkg/`, `examples/`, and `static/` layout.
- Describe the project as Subfinder with selected free-access behavior adapted from or inspired by BBOT, not as a wholesale codebase merger.
- State that only DNSDumpster explicitly mirrors BBOT's public HTMX and short-lived JWT approach.
- State that CertSpotter and URLScan work anonymously with optional credentials for higher service limits.
- Do not claim that every inherited Subfinder source is keyless, unlimited, or permanently free.
- Preserve ProjectDiscovery and BBOT attribution without implying endorsement or affiliation.
- Keep `DISCLAIMER.md`, `LICENSE.md`, and `THANKS.md` at the repository root.
- Do not change enumeration behavior, Go module paths, CLI package names, source implementations, or release versions.
- Keep live passive-source tests opt-in.

---

### Task 1: Fork-Owned User Documentation

**Files:**
- Modify: `README.md`
- Create: `docs/free-source-coverage.md`
- Create: `docs/fork-maintenance.md`
- Delete: `HACKTECH_CHANGES.md`
- Modify: `.github/workflows/subfinder-plus-release.yml`
- Modify: `scripts/test-workflow-contract.ps1`

**Interfaces:**
- Consumes: Existing CLI flags from `pkg/runner/options.go`, module requirements from `go.mod`, provenance from `HACKTECH_CHANGES.md`, and release filenames from `.github/workflows/subfinder-plus-release.yml`.
- Produces: Stable relative links `docs/free-source-coverage.md` and `docs/fork-maintenance.md`; the release workflow and contract test consume the exact maintenance path `docs/fork-maintenance.md`.

- [ ] **Step 1: Establish a failing documentation contract**

Extend `scripts/test-workflow-contract.ps1` so it reads `docs/fork-maintenance.md`, rejects `HACKTECH_CHANGES.md` references, and checks that `README.md` contains all of these exact identifiers:

```powershell
$readme = Get-Content -Raw (Join-Path $repoRoot 'README.md')
$maintenance = Get-Content -Raw (Join-Path $repoRoot 'docs/fork-maintenance.md')

foreach ($required in @(
    'Subfinder-plus',
    'ProjectDiscovery Subfinder',
    'BBOT',
    'docs/free-source-coverage.md',
    'docs/fork-maintenance.md'
)) {
    if ($readme -notmatch [regex]::Escape($required)) {
        throw "README.md is missing required documentation marker: $required"
    }
}

if ($workflow -match 'HACKTECH_CHANGES\.md') {
    throw 'Release workflow still references HACKTECH_CHANGES.md'
}
if ($workflow -notmatch 'docs/fork-maintenance\.md') {
    throw 'Release workflow does not reference docs/fork-maintenance.md'
}
```

- [ ] **Step 2: Run the contract and verify it fails**

Run:

```powershell
pwsh -NoProfile -File scripts/test-workflow-contract.ps1
```

Expected: non-zero exit because `docs/fork-maintenance.md` does not exist and the README lacks the new fork-owned markers.

- [ ] **Step 3: Rewrite the README landing page**

Replace the current upstream-oriented README with a 50–200 line document using this exact section order:

```markdown
# Subfinder-plus

Subfinder's passive-enumeration engine with selected free-access source behavior adapted from or inspired by BBOT.

[Go CI badge] [latest Subfinder-plus release badge]

Subfinder-plus is a HackTech-maintained fork of ProjectDiscovery Subfinder. It preserves Subfinder's CLI and Go architecture while carrying a narrow set of source changes intended to improve useful enumeration without paid service subscriptions. DNSDumpster mirrors BBOT's public web-flow approach; CertSpotter and URLScan use their anonymous endpoints and accept optional keys for higher service limits.

This is not a full merger with BBOT and is not affiliated with or endorsed by ProjectDiscovery or the BBOT project.

## Scope and authorization
## Prerequisites
## Install
### Download a release
### Build from source
### Build with Docker
## Usage
## Free-access source behavior
## Configuration
## Go library
## Documentation
## Contributing
## Security
## Attribution
## License
```

Use fork-owned badge targets:

```text
https://github.com/hack-techv2/subfinder-plus/actions/workflows/ci.yml
https://github.com/hack-techv2/subfinder-plus/releases/latest
```

Use `go 1.24.1`, `git clone https://github.com/hack-techv2/subfinder-plus.git`, `go build -o subfinder-plus ./cmd/subfinder`, and `docker build -t subfinder-plus .` as the source installation examples. Show `subfinder-plus -d example.com -silent` as the first usage example and identify `subfinder-plus -h` as the authoritative flag reference.

- [ ] **Step 4: Create the free-source reference**

Create `docs/free-source-coverage.md` as a Diátaxis reference page with a scope-and-authorization note, then this table:

```markdown
| Source | Works without a user key | Optional credential | Fork behavior |
|---|---:|---:|---|
| DNSDumpster | Yes | No | Fetches the site's short-lived JWT and uses its public HTMX request flow, mirroring BBOT's approach. |
| CertSpotter | Yes | Yes | Queries the public issuance endpoint anonymously; a bearer token can raise service limits. |
| URLScan | Yes | Yes | Queries the public search endpoint anonymously; an API key can raise service limits. |
```

Explain that third-party endpoints, quotas, terms, and availability can change; `-all` also enables inherited sources that may require credentials; and users should consult `subfinder -ls` plus their provider configuration for current runtime selection.

- [ ] **Step 5: Migrate the maintenance reference**

Move every accurate provenance hash, release filename, version-policy statement, and upstream-sync step from `HACKTECH_CHANGES.md` into `docs/fork-maintenance.md`. Add an explanation section covering:

```text
Subfinder supplies the engine, CLI, packages, and upstream update stream.
BBOT informed selected public-source access behavior; DNSDumpster is the explicit code-level adaptation.
Subfinder-plus owns only the maintained delta, tests, release tags, and fork documentation.
```

Delete `HACKTECH_CHANGES.md` only after confirming the new file contains all three provenance commits.

- [ ] **Step 6: Update the release workflow and contract**

Change the upstream-base extraction command to:

```bash
upstream_commit="$(sed -n 's/.*Initial upstream base: `\([0-9a-f]\{40\}\)`.*/\1/p' docs/fork-maintenance.md)"
```

Keep the existing release artifacts, tag pattern, and release version unchanged.

- [ ] **Step 7: Run focused verification**

Run:

```powershell
pwsh -NoProfile -File scripts/test-workflow-contract.ps1
rg -n 'HACKTECH_CHANGES|projectdiscovery/subfinder/(issues|releases|discussions)' README.md docs .github scripts
git diff --check
```

Expected: contract exits `0`; `rg` returns no stale matches; `git diff --check` exits `0`.

- [ ] **Step 8: Commit the user documentation**

```powershell
git add README.md docs/free-source-coverage.md docs/fork-maintenance.md HACKTECH_CHANGES.md .github/workflows/subfinder-plus-release.yml scripts/test-workflow-contract.ps1
git commit -m "docs: establish Subfinder-plus project identity"
```

### Task 2: Contribution and Security Documentation

**Files:**
- Create: `CONTRIBUTING.md`
- Create: `SECURITY.md`

**Interfaces:**
- Consumes: Validation commands in `Makefile`, `scripts/test-workflow-contract.ps1`, and the `dev` contribution target.
- Produces: Stable root links consumed by `README.md` and GitHub contributors.

- [ ] **Step 1: Create contribution guidance**

Create `CONTRIBUTING.md` with these sections and exact commands:

```markdown
# Contributing to Subfinder-plus
## Before opening a change
## Development setup
## Validate a change
## Source behavior changes
## Pull requests
## Upstream synchronization
```

The validation block must contain:

```powershell
go test ./...
go vet ./...
go build ./...
pwsh -NoProfile -File scripts/test-workflow-contract.ps1
```

Require pull requests against `dev`, tests for source behavior changes, no secrets in fixtures or logs, and an explanation of whether a source change belongs upstream or in the fork delta.

- [ ] **Step 2: Create security guidance**

Create `SECURITY.md` with supported-version and reporting sections. Direct sensitive reports to GitHub's private advisory form:

```text
https://github.com/hack-techv2/subfinder-plus/security/advisories/new
```

Tell users not to disclose credentials, target data, or exploit details in public issues. Separate vulnerabilities in Subfinder-plus from outages or behavior changes in third-party passive sources.

- [ ] **Step 3: Verify root-document links**

Run:

```powershell
$required = @('CONTRIBUTING.md', 'SECURITY.md', 'DISCLAIMER.md', 'LICENSE.md', 'THANKS.md')
$required | ForEach-Object { if (-not (Test-Path -LiteralPath $_)) { throw "Missing $_" } }
rg -n 'CONTRIBUTING\.md|SECURITY\.md|DISCLAIMER\.md|LICENSE\.md|THANKS\.md' README.md
git diff --check
```

Expected: all files exist, README links to each document, and whitespace validation passes.

- [ ] **Step 4: Commit contribution documentation**

```powershell
git add CONTRIBUTING.md SECURITY.md README.md
git commit -m "docs: add contribution and security guidance"
```

### Task 3: Fork-Owned GitHub Templates

**Files:**
- Modify: `.github/PULL_REQUEST_TEMPLATE.md`
- Modify: `.github/ISSUE_TEMPLATE/bug_report.md`
- Modify: `.github/ISSUE_TEMPLATE/feature_request.md`
- Modify: `.github/ISSUE_TEMPLATE/config.yml`
- Delete: `.github/ISSUE_TEMPLATE/issue-report.md`

**Interfaces:**
- Consumes: The `dev` branch policy and validation commands documented in `CONTRIBUTING.md`.
- Produces: One bug template, one feature template, fork-owned support links, and a PR checklist consistent with repository policy.

- [ ] **Step 1: Add a failing stale-link check**

Run:

```powershell
rg -n 'github\.com/projectdiscovery/subfinder/(tree/dev|issues|releases|discussions)|discord\.gg/projectdiscovery' .github
```

Expected: matches in the current templates.

- [ ] **Step 2: Rewrite the pull-request template**

Keep the headings `Proposed changes`, `Proof`, and `Checklist`. Require:

```markdown
- [ ] The pull request targets the `dev` branch.
- [ ] `go test ./...`, `go vet ./...`, and `go build ./...` pass.
- [ ] Source behavior changes include focused tests.
- [ ] User-visible changes include documentation.
- [ ] No credentials, target data, generated binaries, or build artifacts are committed.
```

- [ ] **Step 3: Consolidate issue templates**

Keep `bug_report.md` and delete `issue-report.md`. The remaining bug template must request the Subfinder-plus version, OS, exact command with secrets redacted, current behavior, expected behavior, reproduction steps, and safe diagnostic output.

Update `feature_request.md` to ask for the problem, proposed behavior, free-access relevance, alternatives, and whether upstream Subfinder already supports it.

Update `config.yml` to use fork-owned URLs:

```yaml
blank_issues_enabled: false
contact_links:
  - name: Subfinder-plus documentation
    url: https://github.com/hack-techv2/subfinder-plus#readme
    about: Review installation, usage, configuration, and fork scope before opening an issue.
  - name: Security report
    url: https://github.com/hack-techv2/subfinder-plus/security/advisories/new
    about: Report a vulnerability privately. Do not publish sensitive details in an issue.
```

- [ ] **Step 4: Verify templates**

Run:

```powershell
if (Test-Path '.github/ISSUE_TEMPLATE/issue-report.md') { throw 'Duplicate issue template remains' }
rg -n 'github\.com/projectdiscovery/subfinder/(tree/dev|issues|releases|discussions)|discord\.gg/projectdiscovery' .github
git diff --check
```

Expected: duplicate template is absent; `rg` returns no matches; whitespace validation passes.

- [ ] **Step 5: Commit GitHub metadata**

```powershell
git add .github/PULL_REQUEST_TEMPLATE.md .github/ISSUE_TEMPLATE
git commit -m "docs: align GitHub templates with the fork"
```

### Task 4: Repository-Wide Verification

**Files:**
- Verify only: all changed files

**Interfaces:**
- Consumes: Deliverables from Tasks 1–3.
- Produces: Evidence that documentation, automation, Go code, and repository boundaries remain sound.

- [ ] **Step 1: Validate Markdown relative links and images**

Run a PowerShell link scan over tracked Markdown files. For every non-HTTP target extracted from Markdown links and HTML `src` attributes, resolve it relative to the containing file and fail if the path does not exist. Ignore heading fragments after `#`.

Expected: zero missing local targets.

- [ ] **Step 2: Run repository contract tests**

```powershell
pwsh -NoProfile -File scripts/test-workflow-contract.ps1
```

Expected: exit code `0` and a success message from the contract script.

- [ ] **Step 3: Run Go verification**

```powershell
go test ./...
go vet ./...
go build ./...
```

Expected: all commands exit `0`. Live third-party source tests remain skipped unless explicitly enabled by their existing opt-in environment variable.

- [ ] **Step 4: Audit the final diff**

```powershell
git diff --check HEAD~3..HEAD
git status --short
git diff --stat HEAD~3..HEAD
git diff --name-status HEAD~3..HEAD
```

Expected: no whitespace errors; only documentation, GitHub metadata, and documentation-contract files changed; no binaries or generated artifacts appear.

- [ ] **Step 5: Record final evidence**

Capture the commit hashes for the documentation, contributor guidance, and GitHub template commits, plus the exact exit status of every verification command. Do not publish a release or push the branch as part of this task.
