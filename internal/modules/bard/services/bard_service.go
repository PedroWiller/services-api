package bardService

import (
	"net/http"
	"net/url"

	gobard "github.com/aquasecurity/gobard"

	"face-id-detection/config/env"
)

type Chatbot struct {
	Headers        http.Header
	ReqID          int
	AT             string
	ConversationID string
	ResponseID     string
	ChoiceID       string
	Proxy          *url.URL
}

var BardSession *gobard.Bard

func Init() {
	PSID := "ewh3T44S2fzbGKg0cwf_5a_6Um2f0CGZ8c7kOyU5q7xoRg03viWkPi1KUgRrLyGifAJw-w."
	PSIDTS := "sidts-CjEBPVxjSorTZRDOwIAxbPCYshorilKHpXXQ5edldKLVwvHWELr5OKILZMoGnOe_CGMoEAA"

	BardSession = gobard.New(env.PSID, env.PSIDTS)
}

func SendAsk(ask string) (string, error) {
	err := BardSession.Ask(ask)
	if err != nil {
		return "", err
	}

	var response string
	for i := 0; i < BardSession.GetNumOfAnswers(); i++ {
		response = BardSession.GetAnswer()
	}

	return response, nil
}
