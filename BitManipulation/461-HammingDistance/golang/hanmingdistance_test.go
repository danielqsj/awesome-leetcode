package hanmingdistance

import (
	"testing"
)

func Test_hammingDistance(t *testing.T) {
	exceptResult := 2
	result := hammingDistance(1, 4)
	if result != exceptResult {
		t.Errorf("error, except result: %v, actual result: %v", exceptResult, result)
	}
}
