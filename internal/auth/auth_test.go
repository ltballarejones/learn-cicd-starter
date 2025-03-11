package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKeyGood(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey <KEY>")
	key, err := GetAPIKey(h)
	want := "<KEY>"
	if !reflect.DeepEqual(want, key) {
		t.Fatalf("expected: %v, got: %v", want, key)
	}
	if !reflect.DeepEqual(nil, err) {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}
}
