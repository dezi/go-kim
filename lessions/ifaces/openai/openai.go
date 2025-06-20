package openai

import (
	"math/rand"
	"time"
)

const (
	apiKey       = "fcghdsdhjsafefoasdajsjfds"
	url          = "https://openai.com/api/chat"
	systemPrompt = `
diagnostiziere den prompt auf alzheimer 
git kein reasoning aus
antworte nur mit ja oder nein
`
)

type OpenAi struct {
	url          string
	apiKey       string
	systemPrompt string
}

func CreateInstance() (ai *OpenAi) {

	ai = &OpenAi{
		url:          url,
		apiKey:       apiKey,
		systemPrompt: systemPrompt,
	}

	return
}

func (ai *OpenAi) GetApiKey() (apiKey string) {
	apiKey = ai.apiKey
	return
}

func (ai *OpenAi) GetSystemPrompt() (systemPrompt string) {
	systemPrompt = ai.systemPrompt
	return
}

func (ai *OpenAi) EvalDat(userPrompt string) (ok bool, err error) {

	//log.Printf("UserPrompt=%s", userPrompt)

	time.Sleep(time.Second * time.Duration(1+rand.Intn(4)))

	if rand.Intn(2) == 0 {
		ok = true
	}

	return
}
