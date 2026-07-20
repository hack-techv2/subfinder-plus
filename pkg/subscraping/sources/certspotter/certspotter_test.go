package certspotter

import (
	"testing"

	"github.com/projectdiscovery/subfinder/v2/pkg/subscraping"
)

func TestKeyRequirementIsOptional(t *testing.T) {
	if got := (&Source{}).KeyRequirement(); got != subscraping.OptionalKey {
		t.Fatalf("expected OptionalKey, got %v", got)
	}
}

func TestAPIHeadersOmitAuthorizationWithoutKey(t *testing.T) {
	if got := apiHeaders(""); len(got) != 0 {
		t.Fatalf("expected no headers, got %#v", got)
	}
}

func TestAPIHeadersUseBearerTokenWhenConfigured(t *testing.T) {
	if got := apiHeaders("key-1")["Authorization"]; got != "Bearer key-1" {
		t.Fatalf("expected Bearer header, got %q", got)
	}
}
