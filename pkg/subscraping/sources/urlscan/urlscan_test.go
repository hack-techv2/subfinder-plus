package urlscan

import (
	"testing"

	"github.com/projectdiscovery/subfinder/v2/pkg/subscraping"
)

func TestKeyRequirementIsOptional(t *testing.T) {
	if got := (&Source{}).KeyRequirement(); got != subscraping.OptionalKey {
		t.Fatalf("expected OptionalKey, got %v", got)
	}
}

func TestAPIHeadersOmitAPIKeyWithoutKey(t *testing.T) {
	if got := apiHeaders(""); len(got) != 0 {
		t.Fatalf("expected no headers, got %#v", got)
	}
}

func TestAPIHeadersUseConfiguredKey(t *testing.T) {
	if got := apiHeaders("key-1")["api-key"]; got != "key-1" {
		t.Fatalf("expected configured api-key header, got %q", got)
	}
}
