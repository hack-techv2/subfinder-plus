// Package dnsdumpster logic
package dnsdumpster

import (
	"context"
	"encoding/json"
	"html"
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/projectdiscovery/subfinder/v2/pkg/subscraping"
)

// DNSDumpster's site uses HTMX with a per-session JWT on its main lookup form.
var (
	formTagRegex   = regexp.MustCompile(`(?is)<form\b[^>]*>`)
	mainFormRegex  = regexp.MustCompile(`(?i)\bdata-form-id\s*=\s*["']mainform["']`)
	hxHeadersRegex = regexp.MustCompile(`(?i)\bhx-headers\s*=\s*'([^']*)'`)
)

func extractJWT(page string) (string, error) {
	for _, formTag := range formTagRegex.FindAllString(page, -1) {
		if !mainFormRegex.MatchString(formTag) {
			continue
		}
		match := hxHeadersRegex.FindStringSubmatch(formTag)
		if len(match) < 2 {
			continue
		}
		var parsed map[string]string
		if err := json.Unmarshal([]byte(html.UnescapeString(match[1])), &parsed); err != nil {
			return "", err
		}
		jwt := parsed["Authorization"]
		if jwt == "" {
			return "", errJWTNotFound
		}
		return jwt, nil
	}
	return "", errJWTNotFound
}

// Source is the passive scraping agent
type Source struct {
	timeTaken time.Duration
	errors    int
	results   int
	requests  int
}

// Run function returns all subdomains found with the service.
//
// DNSDumpster does not expose a free public API. The website itself uses
// HTMX with a short-lived JWT token injected into the homepage. We scrape
// the homepage to obtain that JWT, then call the same internal HTMX
// endpoint the site itself uses. This mirrors BBOT's approach.
func (s *Source) Run(ctx context.Context, domain string, session *subscraping.Session) <-chan subscraping.Result {
	results := make(chan subscraping.Result)
	s.errors = 0
	s.results = 0
	s.requests = 0

	go func() {
		defer func(startTime time.Time) {
			s.timeTaken = time.Since(startTime)
			close(results)
		}(time.Now())

		jwt, err := s.fetchJWT(ctx, session)
		if err != nil {
			results <- subscraping.Result{Source: s.Name(), Type: subscraping.Error, Error: err}
			s.errors++
			return
		}

		body, err := s.fetchSubdomains(ctx, session, domain, jwt)
		if err != nil {
			results <- subscraping.Result{Source: s.Name(), Type: subscraping.Error, Error: err}
			s.errors++
			return
		}

		seen := make(map[string]struct{})
		for _, sub := range session.Extractor.Extract(body) {
			if sub == "" {
				continue
			}
			if _, ok := seen[sub]; ok {
				continue
			}
			seen[sub] = struct{}{}
			select {
			case <-ctx.Done():
				return
			case results <- subscraping.Result{Source: s.Name(), Type: subscraping.Subdomain, Value: sub}:
				s.results++
			}
		}
	}()

	return results
}

// fetchJWT scrapes the dnsdumpster homepage and extracts the JWT token
// from the embedded HTMX form. Returns the bare token string suitable for
// use as an Authorization header value.
func (s *Source) fetchJWT(ctx context.Context, session *subscraping.Session) (string, error) {
	s.requests++
	resp, err := session.SimpleGet(ctx, "https://dnsdumpster.com/")
	if err != nil {
		session.DiscardHTTPResponse(resp)
		return "", err
	}
	htmlBytes, err := io.ReadAll(resp.Body)
	session.DiscardHTTPResponse(resp)
	if err != nil {
		return "", err
	}

	return extractJWT(string(htmlBytes))
}

// fetchSubdomains POSTs to the htmld endpoint with the scraped JWT and
// returns the raw HTML body for the caller to extract subdomains from.
func (s *Source) fetchSubdomains(ctx context.Context, session *subscraping.Session, domain, jwt string) (string, error) {
	headers := map[string]string{
		"Authorization":  jwt,
		"Content-Type":   "application/x-www-form-urlencoded",
		"Origin":         "https://dnsdumpster.com",
		"Referer":        "https://dnsdumpster.com/",
		"HX-Request":     "true",
		"HX-Target":      "results",
		"HX-Current-URL": "https://dnsdumpster.com/",
	}
	body := strings.NewReader("target=" + strings.ToLower(domain))

	s.requests++
	resp, err := session.Post(ctx, "https://api.dnsdumpster.com/htmld/", "", headers, body)
	if err != nil {
		session.DiscardHTTPResponse(resp)
		return "", err
	}
	respBytes, err := io.ReadAll(resp.Body)
	session.DiscardHTTPResponse(resp)
	if err != nil {
		return "", err
	}
	return string(respBytes), nil
}

// errJWTNotFound is returned when the homepage scrape did not yield a token.
var errJWTNotFound = &dnsdumpsterError{msg: "could not extract JWT from dnsdumpster homepage"}

type dnsdumpsterError struct{ msg string }

func (e *dnsdumpsterError) Error() string { return e.msg }

// Name returns the name of the source
func (s *Source) Name() string {
	return "dnsdumpster"
}

func (s *Source) IsDefault() bool {
	return true
}

func (s *Source) HasRecursiveSupport() bool {
	return false
}

func (s *Source) KeyRequirement() subscraping.KeyRequirement {
	return subscraping.NoKey
}

func (s *Source) NeedsKey() bool {
	return s.KeyRequirement() == subscraping.RequiredKey
}

func (s *Source) AddApiKeys(_ []string) {
	// no key needed - JWT is scraped from the homepage
}

func (s *Source) Statistics() subscraping.Statistics {
	return subscraping.Statistics{
		Errors:    s.errors,
		Results:   s.results,
		Requests:  s.requests,
		TimeTaken: s.timeTaken,
	}
}
