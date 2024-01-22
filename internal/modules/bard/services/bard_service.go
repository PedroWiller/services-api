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
