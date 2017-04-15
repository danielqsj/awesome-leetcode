package reverseWordsIII

import (
	"testing"
)

func Test_reverseWordsIII(t *testing.T) {
	exceptResult := "s'teL ekat edoCteeL tsetnoc"
	result := reverseWordsIII("Let's take LeetCode contest")
	if result != exceptResult {
		t.Errorf("error, except result: %v, actual result: %v", exceptResult, result)
	}
}
