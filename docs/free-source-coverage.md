# Free-access source coverage

## Scope and authorization

This reference describes the maintained free-access behavior in Subfinder-plus, not permission to query a provider or a target. Use the tool only for authorized domains and comply with each provider's applicable terms.

## Reference

| Source | Works without a user key | Optional credential | Fork behavior |
|---|---:|---:|---|
| DNSDumpster | Yes | No | Fetches the site's short-lived JWT and uses its public HTMX request flow, mirroring BBOT's approach. |
| CertSpotter | Yes | Yes | Queries the public issuance endpoint anonymously; a bearer token can raise service limits. |
| URLScan | Yes | Yes | Queries the public search endpoint anonymously; an API key can raise service limits. |

Third-party endpoints, quotas, terms, and availability can change without notice. `-all` also enables inherited sources that may require credentials. Consult `subfinder -ls` and your provider configuration for the current runtime source selection.
