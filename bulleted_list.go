package skill_tailor

const (
	SchemaID          = "https://github.com/jgilman1337/skill_tailor/result_schema.go"
	SchemaName        = "bulleted_list"
	SchemaDescription = "A list of bullet points for a job resume."

	defaultMin = 5
	defaultMax = 5
)

// BulletedList is a list of bullet points for a job resume.
type BulletedList []string

// GetBulletedListSchema returns the bullet list schema with the given min and max number of items.
func GetBulletedListSchema(min, max int) map[string]any {
	// Use default values if not provided or max is less than min
	if min < 1 {
		min = defaultMin
	}
	if max < 1 {
		max = defaultMax
	}
	if max < min {
		max = min
	}

	// OpenAI's JSON schema response_format requires `schema` to be a JSON object, so encode as a map[string]any
	return map[string]any{
		"$id":         SchemaID,
		"$schema":     "https://json-schema.org/draft/2020-12/schema",
		"title":       SchemaName,
		"description": SchemaDescription,
		"type":        "array",
		"minItems":    min,
		"maxItems":    max,
		"items": map[string]any{
			"type":      "string",
			"minLength": 1,
		},
	}
}
