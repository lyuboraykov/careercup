package reversewords

import "testing"

func TestReverseWords(t *testing.T) {
	expected := "I ma a nam"
	got := reverseWordsInSentence("I am a man")
	if expected != got {
		t.Error("Wrong words reverse", expected, got)
	}
}
