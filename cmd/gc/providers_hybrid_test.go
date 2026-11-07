package main

import (
	"strings"
	"testing"

	"github.com/gastownhall/gascity/internal/config"
)

func TestNewHybridProviderAcceptsSSHRemoteProvider(t *testing.T) {
	sc := config.SessionConfig{
		RemoteProvider: "ssh:example.invalid",
		RemoteMatch:    "remote-",
	}

	sp, err := newHybridProvider(sc, "test-city", t.TempDir())
	if err != nil {
		t.Fatalf("newHybridProvider: %v", err)
	}
	if sp == nil {
		t.Fatal("newHybridProvider returned a nil provider")
	}
}

func TestHybridRemoteProviderNameDefaultsToK8s(t *testing.T) {
	if got := hybridRemoteProviderName(config.SessionConfig{}); got != "k8s" {
		t.Fatalf("hybridRemoteProviderName(empty) = %q, want k8s", got)
	}
}

func TestNewHybridProviderRejectsNestedHybrid(t *testing.T) {
	_, err := newHybridProvider(config.SessionConfig{RemoteProvider: "hybrid"}, "test-city", t.TempDir())
	if err == nil {
		t.Fatal("newHybridProvider(remote_provider=hybrid) succeeded, want unsupported provider error")
	}
	if !strings.Contains(err.Error(), "unsupported provider") {
		t.Fatalf("newHybridProvider(remote_provider=hybrid) error = %q, want unsupported provider", err)
	}
}
