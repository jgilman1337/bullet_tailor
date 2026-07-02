package bullet_tailor

import (
	"encoding/json"
	"errors"
	"strings"

	"gopkg.in/yaml.v3"
)

// Represents a single question-answer pair.
type Questionnaire struct {
	Question string `json:"question" yaml:"question"`
	Answer   string `json:"answer" yaml:"answer"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (q *Questionnaire) UnmarshalJSON(data []byte) error {
	// Decode into an anonymous struct to avoid infinite recursion into
	// Questionnaire.UnmarshalJSON.
	var aux struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Map the decoded data to the questionnaire
	q.Question = aux.Question
	q.Answer = aux.Answer

	return q.validate()
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (q *Questionnaire) UnmarshalYAML(value *yaml.Node) error {
	// Decode into an anonymous struct to avoid infinite recursion into
	// Questionnaire.UnmarshalYAML.
	var aux struct {
		Question string `yaml:"question"`
		Answer   string `yaml:"answer"`
	}
	if err := value.Decode(&aux); err != nil {
		return err
	}

	// Map the decoded data to the questionnaire
	q.Question = aux.Question
	q.Answer = aux.Answer

	return q.validate()
}

// Ensures that the questionnaire is valid.
func (q *Questionnaire) validate() error {
	// Enforce that the questionnaire is not nil
	if q == nil {
		return errors.New("questionnaire is nil")
	}

	q.Question = strings.TrimSpace(q.Question)
	q.Answer = strings.TrimSpace(q.Answer)

	// Enforce that each question has a corresponding (non-blank) answer
	if q.Question == "" {
		return errors.New("question cannot be blank")
	}
	if q.Answer == "" {
		return errors.New("answer cannot be blank")
	}

	return nil
}
