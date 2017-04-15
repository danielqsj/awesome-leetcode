package keyboardRow

import (
"testing"
)

func Test_keyboardRow(t *testing.T) {
	exceptResult := []string{"Alaska", "Dad"}
	result := keyboardRow([]string{"Hello", "Alaska", "Dad", "Peace"})

	if len(result) != len(exceptResult) {
		t.Errorf("error, except result: %v, actual result: %v", exceptResult, result)
		return
	}

	for i := range result {
		if result[i] != exceptResult[i] {
			t.Errorf("error, except result: %v, actual result: %v", exceptResult, result)
			return
		}
	}
}
