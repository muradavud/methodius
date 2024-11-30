package bot

import (
	_ "embed"
	"encoding/json"
)

//go:embed resources/keyvalue.json
var keyValue []byte

func parseKeyValueFile() (map[string]string, error) {
	var result map[string]string
	err := json.Unmarshal(keyValue, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
