package tailor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jgilman1337/skill_tailor/pkg/common"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

// InferArgs are the arguments for the infer function.
type InferArgs struct {
	JobListing    string         `json:"job_listing" yaml:"job_listing"`
	MinBullets    int            `json:"min_bullets" yaml:"min_bullets"`
	MaxBullets    int            `json:"max_bullets" yaml:"max_bullets"`
	Questionnaire *Questionnaire `json:"questionnaire" yaml:"questionnaire"`
	SystemPrompt  string         `json:"system_prompt" yaml:"system_prompt"`
	Timeout       time.Duration  `json:"timeout" yaml:"timeout"`
}

// DefaultInferArgs returns the default set of arguments for the infer function.
func DefaultInferArgs() *InferArgs {
	return &InferArgs{
		MinBullets:   5,
		MaxBullets:   5,
		SystemPrompt: DefaultPrompt,
		Timeout:      30 * time.Second,
	}
}

// NewInferClient creates a new client for the GPT API.
func NewInferClient(auth *common.GPTAuth) *openai.Client {
	client := openai.NewClient(
		option.WithAPIKey(auth.APIKey),
		option.WithBaseURL(auth.Endpoint),
	)
	return &client
}

// Infer is the main function for inferring the template from the questionnaire.
func Infer(client *openai.Client, params *common.GPTParams, args *InferArgs) (*BulletedList, error) {
	// Ensure the client is not nil
	if client == nil {
		return nil, errors.New("client is nil")
	}

	// Ensure the params are not nil
	if params == nil {
		params = common.DefaultGPTParams()
	}

	// Ensure the args are not nil
	if args == nil {
		args = DefaultInferArgs()
	}

	// Ensure the required fields are not nil
	if args.JobListing == "" {
		return nil, errors.New("job listing is required")
	}
	if args.Questionnaire == nil {
		return nil, errors.New("questionnaire is required")
	}

	// Create a context with the timeout
	ctx, cancel := context.WithTimeout(context.Background(), args.Timeout)
	defer cancel()

	// Serialize the questionnaire to JSON
	questionnaireJSON, err := json.Marshal(args.Questionnaire)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize questionnaire: %w", err)
	}

	// Create the questionnaire and job listing messages
	questionnaireMessage := fmt.Sprintf(
		"User's Filled Questionnaire (JSON):\n%s",
		strings.TrimSpace(string(questionnaireJSON)),
	)
	jobListingMessage := fmt.Sprintf(
		"Job Listing:\n%s",
		strings.TrimSpace(args.JobListing),
	)

	// Add schema info
	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:        SchemaName,
		Description: openai.String(SchemaDescription),
		Schema:      GetBulletedListSchema(args.MinBullets, args.MaxBullets),
		Strict:      openai.Bool(true),
	}

	// Construct the response query
	resp, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model:       params.Model,
		MaxTokens:   openai.Int(int64(params.MaxTokens)),
		Temperature: openai.Float(params.Temperature),
		TopP:        openai.Float(params.TopP),
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(args.SystemPrompt),
			openai.UserMessage(questionnaireMessage),
			openai.UserMessage(jobListingMessage),
		},
		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{
				JSONSchema: schemaParam,
			},
		},
	})

	// Check for errors, including if the API took too long to produce a response
	if err != nil {
		if ctx.Err() != nil {
			return nil, fmt.Errorf("time limit exceeded: %w", ctx.Err())
		}
		return nil, fmt.Errorf("chat completion error: %w", err)
	}

	// Get the response and clean it
	rawContent := strings.TrimSpace(resp.Choices[0].Message.Content)
	cleanedContent := clean(rawContent)

	// Deserialize a job description from the output
	var bulletedList BulletedList
	err = json.Unmarshal([]byte(cleanedContent), &bulletedList)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize bulleted list: %w", err)
	}

	return &bulletedList, nil
}

// Some OpenAI-compatible backends may still wrap structured output in
// ```json ... ``` fences even when schema enforcement is requested.
// This function removes the fences and any whitespace around the content.
func clean(text string) string {
	content := text
	if strings.HasPrefix(content, "```") {
		content = strings.TrimPrefix(content, "```json")
		content = strings.TrimPrefix(content, "```")
		content = strings.TrimSuffix(content, "```")
		content = strings.TrimSpace(content)
	}

	return content
}
