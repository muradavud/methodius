// create unit test for extract function

package process

import (
	"encoding/json"
	"os"
	"testing"
)

func TestExtractStringFromFile(t *testing.T) {
	file, err := os.Open("tmp/test")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var transcription Transcription

	err = decoder.Decode(&transcription)
	if err != nil {
		t.Error(err)
	}

	if transcription.Results.Transcripts[0].Transcript !=
		"Hey, man, what's up? How are you doing? Glad to see you, man." {
		t.Error("Transcription did not match expected value")
	}
}
