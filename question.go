package bullet_tailor

import (
	"encoding/json"
	"errors"
	"strings"

	"gopkg.in/yaml.v3"
)

// Question represents a single question-answer pair in a questionnaire.
type Question struct {
	Question string `json:"question" yaml:"question"`
	Answer   string `json:"answer" yaml:"answer"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (q *Question) UnmarshalJSON(data []byte) error {
	// Decode into an anonymous struct to avoid infinite recursion into
	// Question.UnmarshalJSON.
	var aux struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Map the decoded data to the question
	q.Question = aux.Question
	q.Answer = aux.Answer

	return q.validate()
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (q *Question) UnmarshalYAML(value *yaml.Node) error {
	// Decode into an anonymous struct to avoid infinite recursion into
	// Question.UnmarshalYAML.
	var aux struct {
		Question string `yaml:"question"`
		Answer   string `yaml:"answer"`
	}
	if err := value.Decode(&aux); err != nil {
		return err
	}

	// Map the decoded data to the question
	q.Question = aux.Question
	q.Answer = aux.Answer

	return q.validate()
}

// Ensures that the question is valid.
func (q *Question) validate() error {
	// Enforce that the question is not nil
	if q == nil {
		return errors.New("question is nil")
	}

	// Trim the question and answer
	q.Question = strings.TrimSpace(q.Question)
	q.Answer = strings.TrimSpace(q.Answer)

	// Enforce that the question has a corresponding (non-blank) answer
	if q.Question == "" {
		return errors.New("question cannot be blank")
	}
	if q.Answer == "" {
		return errors.New("answer cannot be blank")
	}

	return nil
}
