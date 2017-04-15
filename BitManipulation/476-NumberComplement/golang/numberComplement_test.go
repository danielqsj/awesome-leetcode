package numberComplement

import (
	"testing"
)

func Test_numberComplement(t *testing.T) {
	exceptResult := 1
	result := numberComplement(2)
	if result != exceptResult {
		t.Errorf("error, except result: %v, actual result: %v", exceptResult, result)
	}
}
