package test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/jgilman1337/bullet_tailor"
	"gopkg.in/yaml.v3"
)

func TestQuestionnaireUnmarshalJSON_Valid(t *testing.T) {
	input := []byte(`{
		"id": " job-1 ",
		"questions": [
			{"question":" q1 ", "answer":" a1 "}
		]
	}`)

	var q bullet_tailor.Questionnaire
	if err := json.Unmarshal(input, &q); err != nil {
		t.Fatalf("expected valid unmarshal, got error: %v", err)
	}

	if q.ID != "job-1" {
		t.Fatalf("expected trimmed id %q, got %q", "job-1", q.ID)
	}
	if len(q.Questions) != 1 {
		t.Fatalf("expected 1 question, got %d", len(q.Questions))
	}
	if q.Questions[0].Question != "q1" || q.Questions[0].Answer != "a1" {
		t.Fatalf("expected trimmed question/answer, got %#v", q.Questions[0])
	}
}

func TestQuestionnaireUnmarshalJSON_InvalidID(t *testing.T) {
	input := []byte(`{
		"id": "   ",
		"questions": [{"question":"q1","answer":"a1"}]
	}`)

	var q bullet_tailor.Questionnaire
	err := json.Unmarshal(input, &q)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "id cannot be blank") {
		t.Fatalf("expected id validation error, got: %v", err)
	}
}

func TestQuestionnaireUnmarshalYAML_Valid(t *testing.T) {
	input := `
id: job-2
questions:
  - question: " q2 "
    answer: " a2 "
`

	var q bullet_tailor.Questionnaire
	if err := yaml.Unmarshal([]byte(input), &q); err != nil {
		t.Fatalf("expected valid unmarshal, got error: %v", err)
	}

	if q.ID != "job-2" {
		t.Fatalf("expected id %q, got %q", "job-2", q.ID)
	}
	if len(q.Questions) != 1 {
		t.Fatalf("expected 1 question, got %d", len(q.Questions))
	}
	if q.Questions[0].Question != "q2" || q.Questions[0].Answer != "a2" {
		t.Fatalf("expected trimmed question/answer, got %#v", q.Questions[0])
	}
}

func TestQuestionnaireUnmarshalYAML_InvalidQuestion(t *testing.T) {
	input := `
id: job-3
questions:
  - question: " "
    answer: "a3"
`

	var q bullet_tailor.Questionnaire
	err := yaml.Unmarshal([]byte(input), &q)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	// For YAML, the error can come directly from Question.UnmarshalYAML.
	if !strings.Contains(err.Error(), "question cannot be blank") {
		t.Fatalf("expected question validation error, got: %v", err)
	}
}
