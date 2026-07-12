package test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jgilman1337/bullet_tailor"
)

func TestInferIntegration(t *testing.T) {
	// Read the questionnaire and unmarshal it
	formBytes, err := os.ReadFile("data/questionnaires/swe_job1.json")
	if err != nil {
		t.Fatalf("failed to read questionnaire: %v", err)
	}
	var questionnaire bullet_tailor.Questionnaire
	if err := json.Unmarshal(formBytes, &questionnaire); err != nil {
		t.Fatalf("failed to unmarshal questionnaire: %v", err)
	}

	// Read the job listing and unmarshal it
	jobBytes, err := os.ReadFile("data/job_listings/job1.txt")
	if err != nil {
		t.Fatalf("failed to read job listing: %v", err)
	}

	// Create a GPT config and inference client
	auth, params := InitGPTConfig(t)
	client := bullet_tailor.NewInferClient(auth)

	// Create the inference arguments
	args := &bullet_tailor.InferArgs{
		JobListing:    strings.TrimSpace(string(jobBytes)),
		MinBullets:    5,
		MaxBullets:    5,
		Questionnaire: &questionnaire,
		SystemPrompt:  bullet_tailor.DefaultPrompt,
		Timeout:       60 * time.Second,
	}

	// Infer the bulleted list and ensure the number of bullets is within the expected range
	bulletedList, err := bullet_tailor.Infer(client, params, args)
	if err != nil {
		t.Fatalf("Infer failed: %v", err)
	}
	if got, wantMin, wantMax := len(*bulletedList), args.MinBullets, args.MaxBullets; got < wantMin || got > wantMax {
		t.Fatalf("unexpected bullet count: got %d, expected between %d and %d", got, wantMin, wantMax)
	}

	// Verify the bulleted list is not empty
	for i, b := range *bulletedList {
		if strings.TrimSpace(b) == "" {
			t.Fatalf("bullet[%d] is blank", i)
		}
	}

	// Print the bulleted list to the console
	for i, b := range *bulletedList {
		fmt.Printf("Bullet %d: %s\n", i+1, b)
	}
}
