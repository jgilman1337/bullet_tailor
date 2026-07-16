package common

// GPTAuth is the authentication information for the GPT API.
type GPTAuth struct {
	Endpoint string `env:"ENDPOINT" json:"endpoint" yaml:"endpoint"`
	APIKey   string `env:"API_KEY" json:"api_key" yaml:"api_key"`
}

// DefaultGPTAuth returns the default authentication information for the GPT API.
func DefaultGPTAuth() *GPTAuth {
	return &GPTAuth{
		Endpoint: "https://api.openai.com/v1",
	}
}

// GPTParams are the parameters for the GPT API.
type GPTParams struct {
	Model       string  `env:"MODEL" json:"model" yaml:"model"`
	MaxTokens   int     `env:"MAX_TOKENS" json:"max_tokens" yaml:"max_tokens"`
	Temperature float64 `env:"TEMPERATURE" json:"temperature" yaml:"temperature"`
	TopP        float64 `env:"TOP_P" json:"top_p" yaml:"top_p"`
}

// DefaultGPTParams returns the default parameters for the GPT API.
func DefaultGPTParams() *GPTParams {
	return &GPTParams{
		MaxTokens:   2048,
		Temperature: 0.5,
		TopP:        1.0,
	}
}
