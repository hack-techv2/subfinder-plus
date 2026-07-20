package passive

import "testing"

func TestLiveSourceTestsRequireExplicitOptIn(t *testing.T) {
	t.Setenv("SUBFINDER_LIVE_TESTS", "")
	if liveSourceTestsEnabled() {
		t.Fatal("live source tests must be disabled by default")
	}

	t.Setenv("SUBFINDER_LIVE_TESTS", "1")
	if !liveSourceTestsEnabled() {
		t.Fatal("live source tests must run after explicit opt-in")
	}
}
