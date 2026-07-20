package dnsdumpster

import (
	"testing"

	"github.com/projectdiscovery/subfinder/v2/pkg/subscraping"
)

func TestKeyRequirementIsNoKey(t *testing.T) {
	if got := (&Source{}).KeyRequirement(); got != subscraping.NoKey {
		t.Fatalf("expected NoKey, got %v", got)
	}
}

func TestExtractJWTFromHTMXForm(t *testing.T) {
	html := `<form hx-headers='{&quot;Authorization&quot;:&quot;Bearer token-1&quot;}' class="lookup" data-form-id="mainform"></form>`
	got, err := extractJWT(html)
	if err != nil {
		t.Fatalf("extractJWT returned error: %v", err)
	}
	if got != "Bearer token-1" {
		t.Fatalf("expected JWT header value, got %q", got)
	}
}

func TestExtractJWTRejectsMissingAuthorization(t *testing.T) {
	if _, err := extractJWT(`<form data-form-id="mainform"></form>`); err == nil {
		t.Fatal("expected missing JWT to fail")
	}
}
