package bullet_tailor

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// Questionnaire represents a questionnaire, which is a list of questions and answers and an ID to map it to a corresponding resume item.
type Questionnaire struct {
	ID        string     `json:"id" yaml:"id"`
	Questions []Question `json:"questions" yaml:"questions"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (q *Questionnaire) UnmarshalJSON(data []byte) error {
	// Decode into an anonymous struct to avoid infinite recursion into
	// Questionnaire.UnmarshalJSON.
	var aux struct {
		ID        string     `json:"id"`
		Questions []Question `json:"questions"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	q.ID = aux.ID
	q.Questions = aux.Questions

	return q.validate()
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (q *Questionnaire) UnmarshalYAML(value *yaml.Node) error {
	// Decode into an anonymous struct to avoid infinite recursion into
	// Questionnaire.UnmarshalYAML.
	var aux struct {
		ID        string     `yaml:"id"`
		Questions []Question `yaml:"questions"`
	}
	if err := value.Decode(&aux); err != nil {
		return err
	}

	q.ID = aux.ID
	q.Questions = aux.Questions

	return q.validate()
}

// Ensures that the questionnaire is valid.
func (q *Questionnaire) validate() error {
	// Enforce that the questionnaire is not nil
	if q == nil {
		return errors.New("questionnaire is nil")
	}

	// Enforce that the ID is not blank
	q.ID = strings.TrimSpace(q.ID)
	if q.ID == "" {
		return errors.New("id cannot be blank")
	}

	// Enforce that the questions are not empty
	if len(q.Questions) == 0 {
		return errors.New("questions cannot be empty")
	}

	// Re-validate questions in case Questionnaire is constructed programmatically
	for i := range q.Questions {
		if err := q.Questions[i].validate(); err != nil {
			return fmt.Errorf("question[%d]: %w", i, err)
		}
	}

	return nil
}
