package process

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/s3"
)

func ExtractStringFromFile(object *s3.GetObjectOutput) (string, error) {
	decoder := json.NewDecoder(object.Body)

	var transcription Transcription

	err := decoder.Decode(&transcription)
	if err != nil {
		return "", err
	}

	return transcription.Results.Transcripts[0].Transcript, nil

}

type Transcription struct {
	JobName   string `json:"jobName"`
	AccountID string `json:"accountId"`
	Results   struct {
		Transcripts []struct {
			Transcript string `json:"transcript"`
		} `json:"transcripts"`
		Items []struct {
			StartTime    string `json:"start_time,omitempty"`
			EndTime      string `json:"end_time,omitempty"`
			Alternatives []struct {
				Confidence string `json:"confidence"`
				Content    string `json:"content"`
			} `json:"alternatives"`
			Type string `json:"type"`
		} `json:"items"`
	} `json:"results"`
	Status string `json:"status"`
}
