package bullet_tailor

import (
	"strconv"
	"strings"
)

// resultSchema provides the constraint schema for the results from OpenAI's response.
const resultSchema = `{
	"$id": "https://example.com/address.schema.json",
	"$schema": "https://json-schema.org/draft/2020-12/schema",
	"description": "A list of bullet points for a job resume.",
	"type": "array",
	"minItems": {{itemsMin}},
	"maxItems": {{itemsMax}},
	"items": {
		"type": "string",
		"minLength": 1,
		"maxLength": 255
	}
}`

const defaultMin = 5
const defaultMax = 5

// GetResultSchema returns the result schema with the given min and max number of items.
func GetResultSchema(min, max int) string {
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

	// Override templates at runtime without any runtime file access.
	return strings.NewReplacer(
		"{{itemsMin}}", strings.TrimSpace(strconv.Itoa(min)),
		"{{itemsMax}}", strings.TrimSpace(strconv.Itoa(max)),
	).Replace(resultSchema)
}
