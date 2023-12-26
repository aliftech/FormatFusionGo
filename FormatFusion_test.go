package lib

import (
    "testing"
)

func TestYAMLToJSON(t *testing.T) {
    yamlData := []byte(`
name: John Doe
age: 30
hobbies:
  - coding
  - reading
`)

    expectedJSON := `{"age":30,"hobbies":["coding","reading"],"name":"John Doe"}`

    jsonData, err := yamlToJson(yamlData)
    if err != nil {
        t.Errorf("Error converting YAML to JSON: %v", err)
    } else {
        if string(jsonData) != expectedJSON {
            t.Errorf("Incorrect JSON output:\nExpected:\n%s\nGot:\n%s", expectedJSON, string(jsonData))
        }
    }
}

// TestJSONToYAML tests the JSON to YAML conversion
func TestJSONToYAML(t *testing.T) {
    jsonData := []byte(`{"name": "Jane Doe", "age": 25, "city": "New York"}`)

    expectedYAML := `age: 25
city: New York
name: Jane Doe
`

    yamlData, err := jsonToYAML(jsonData)
    if err != nil {
        t.Errorf("Error converting JSON to YAML: %v", err)
    } else {
        if string(yamlData) != expectedYAML {
            t.Errorf("Incorrect YAML output:\nExpected:\n%s\nGot:\n%s", expectedYAML, string(yamlData))
        }
    }
}