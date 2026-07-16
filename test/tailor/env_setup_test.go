package tailor_test

import (
	"testing"

	"github.com/jgilman1337/skill_tailor/pkg/common"
	common_test "github.com/jgilman1337/skill_tailor/test/common"
)

func TestInitGPTConfig_DefaultsWhenEnvUnset(t *testing.T) {
	// Ensure env doesn't accidentally affect this test.
	t.Setenv("OPENAI_ENDPOINT", "")
	t.Setenv("OPENAI_API_KEY", "")
	t.Setenv("OPENAI_MODEL", "")
	t.Setenv("OPENAI_TEMPERATURE", "")
	t.Setenv("OPENAI_MAX_TOKENS", "")
	t.Setenv("OPENAI_TOP_P", "")

	// Get the auth and params from the environment
	auth, params := common_test.InitGPTConfig(t)

	// Verify the auth and params are set correctly
	if auth.Endpoint != common.DefaultGPTAuth().Endpoint {
		t.Fatalf("unexpected auth.Endpoint: got %q want %q", auth.Endpoint, common.DefaultGPTAuth().Endpoint)
	}
	if auth.APIKey != "" {
		t.Fatalf("unexpected auth.APIKey: got %q want empty", auth.APIKey)
	}

	if params.Model != "" {
		t.Fatalf("unexpected params.Model: got %q want empty", params.Model)
	}
	if params.MaxTokens != common.DefaultGPTParams().MaxTokens {
		t.Fatalf("unexpected params.MaxTokens: got %d want %d", params.MaxTokens, common.DefaultGPTParams().MaxTokens)
	}
	if params.Temperature != common.DefaultGPTParams().Temperature {
		t.Fatalf("unexpected params.Temperature: got %v want %v", params.Temperature, common.DefaultGPTParams().Temperature)
	}
	if params.TopP != common.DefaultGPTParams().TopP {
		t.Fatalf("unexpected params.TopP: got %v want %v", params.TopP, common.DefaultGPTParams().TopP)
	}
}

func TestInitGPTConfig_OverridesFromEnv(t *testing.T) {
	// Ensure env doesn't accidentally affect this test.
	t.Setenv("OPENAI_ENDPOINT", "http://example.local/v1")
	t.Setenv("OPENAI_API_KEY", "test-api-key")
	t.Setenv("OPENAI_MODEL", "test-model")
	t.Setenv("OPENAI_TEMPERATURE", "0.9")
	t.Setenv("OPENAI_MAX_TOKENS", "123")
	t.Setenv("OPENAI_TOP_P", "0.2")

	// Get the auth and params from the environment
	auth, params := common_test.InitGPTConfig(t)

	// Verify the auth and params are set correctly
	if auth.Endpoint != "http://example.local/v1" {
		t.Fatalf("unexpected auth.Endpoint: got %q want %q", auth.Endpoint, "http://example.local/v1")
	}
	if auth.APIKey != "test-api-key" {
		t.Fatalf("unexpected auth.APIKey: got %q want %q", auth.APIKey, "test-api-key")
	}

	if params.Model != "test-model" {
		t.Fatalf("unexpected params.Model: got %q want %q", params.Model, "test-model")
	}
	if params.MaxTokens != 123 {
		t.Fatalf("unexpected params.MaxTokens: got %d want %d", params.MaxTokens, 123)
	}
	if params.Temperature != 0.9 {
		t.Fatalf("unexpected params.Temperature: got %v want %v", params.Temperature, 0.9)
	}
	if params.TopP != 0.2 {
		t.Fatalf("unexpected params.TopP: got %v want %v", params.TopP, 0.2)
	}
}
