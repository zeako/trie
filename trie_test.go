package trie

import (
	"testing"
)

func TestNew(t *testing.T) {
	trie := New()
	if trie == nil {
		t.Error("trie is nil")
	}
}
func TestTryWrite(t *testing.T) {
	trie := New()
	k, v := "key", int64(1)

	err := trie.TryWrite(k, v)
	if err != nil {
		t.Fatalf("writing failed with error: %s", err)
	}
}

func TestTryRead(t *testing.T) {
	trie := New()
	key, want := "key", int64(1)

	trie.TryWrite(key, want)
	got, err := trie.TryRead(key)
	if err != nil {
		t.Fatalf("reading failed with error: %s", err)
	}
	if want != got {
		t.Fatalf("want: %d, got: %d", want, got)
	}
}