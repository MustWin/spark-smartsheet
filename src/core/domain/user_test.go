package domain

import "testing"

func TestTokenPrefix(t *testing.T) {
	tok, err := generateRandomString("foo@bar.baz", 32)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Token: %q", tok)

	prefix, err := extractPrefix(tok)
	if err != nil {
		t.Fatalf("got error (%q): %v", prefix, err)
	}

	if prefix != "foo@bar.baz" {
		t.Fatalf("expected foo@bar.baz, got %s", prefix)
	}
}
