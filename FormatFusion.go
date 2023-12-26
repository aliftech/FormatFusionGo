/**
 * Author: Wahyu Krisna Aji
 * File: FormatFusion.go
 */

package lib

import (
    "encoding/json"
    "fmt"

    "github.com/ghodss/yaml" // Import a reliable YAML library
)

func yamlToJson(yamlData []byte) ([]byte, error) {
    var jsonData interface{}

    // Unmarshal YAML into a generic interface{} to handle any data structure
    if err := yaml.Unmarshal(yamlData, &jsonData); err != nil {
        return nil, fmt.Errorf("error unmarshaling YAML: %w", err)
    }

    // Marshal the parsed JSON into bytes
    jsonBytes, err := json.Marshal(jsonData)
    if err != nil {
        return nil, fmt.Errorf("error marshaling JSON: %w", err)
    }

    return jsonBytes, nil
}

// Function to convert JSON data into YAML
func jsonToYAML(jsonData []byte) ([]byte, error) {
    var yamlData interface{}

    // Unmarshal JSON into a generic interface{}
    if err := json.Unmarshal(jsonData, &yamlData); err != nil {
        return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
    }

    // Marshal the parsed YAML into bytes
    yamlBytes, err := yaml.Marshal(yamlData)
    if err != nil {
        return nil, fmt.Errorf("error marshaling YAML: %w", err)
    }

    return yamlBytes, nil
}