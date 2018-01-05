package tdlib

import (
	"encoding/json"
	"testing"
)

func TestClientGetTextEntities(t *testing.T) {
	client := NewClient()

	bytes, _ := json.Marshal(map[string]interface{}{
		"@type": "getTextEntities",
		"text":  "@telegram /test_command https://telegram.org telegram.me",
		"@extra": map[string]interface{}{
			"5": 7.0,
		},
	})

	resp := client.Execute(bytes)

	var item map[string]interface{}

	err := json.Unmarshal([]byte(resp), &item)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
