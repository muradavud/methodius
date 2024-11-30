package bot

import (
	"bytes"
	_ "embed"

	tele "gopkg.in/telebot.v3"
)

//go:embed resources/loading.gif
var loading []byte

func SendLoading(c tele.Context) (*tele.Message, error) {
	animation := &tele.Video{
		File: tele.FromReader(bytes.NewReader(loading)),
	}

	msg, err := B.Send(c.Recipient(), animation)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
