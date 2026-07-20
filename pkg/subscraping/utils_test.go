package subscraping

import "testing"

func TestPickRandomOptionalReturnsZeroValueWithoutKeys(t *testing.T) {
	if got := PickRandomOptional([]string{}, "optional-source"); got != "" {
		t.Fatalf("expected empty key, got %q", got)
	}
}

func TestPickRandomOptionalReturnsConfiguredKey(t *testing.T) {
	if got := PickRandomOptional([]string{"key-1"}, "optional-source"); got != "key-1" {
		t.Fatalf("expected configured key, got %q", got)
	}
}
