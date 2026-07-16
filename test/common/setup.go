package common_test

import (
	"testing"

	"github.com/caarlos0/env/v11"
	"github.com/jgilman1337/skill_tailor/pkg/common"
)

// InitGPTConfig initializes the GPT auth + params configuration structs
// from OPENAI_* environment variables for unit tests.
//
// If a variable is missing, it falls back to skill_tailor defaults.
func InitGPTConfig(t *testing.T) (*common.GPTAuth, *common.GPTParams) {
	t.Helper()

	// Define the default auth and param structs
	auth := common.DefaultGPTAuth()
	params := common.DefaultGPTParams()

	// Parse overrides from environment using the existing `env:"..."` struct tags
	// Defaults are preserved because we pre-initialize the structs first
	if err := env.ParseWithOptions(auth, env.Options{Prefix: "OPENAI_"}); err != nil {
		t.Fatalf("failed to parse GPTAuth from env: %v", err)
	}
	if err := env.ParseWithOptions(params, env.Options{Prefix: "OPENAI_"}); err != nil {
		t.Fatalf("failed to parse GPTParams from env: %v", err)
	}

	return auth, params
}
