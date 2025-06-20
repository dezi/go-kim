package claude

import (
	"math/rand"
	"time"
)

const (
	apiKey       = "fcghdsdhjsafefoasdajsjfds"
	url          = "https://openai.com/api/chat"
	systemPrompt = `
diagnostiziere den prompt auf alzheimer 
git auf gar keinen fall niemals reasoning aus. Niemals!
antworte nur mit ja oder nein
`
)

type Claude struct {
	url          string
	apiKey       string
	systemPrompt string
}

func CreateInstance() (ai *Claude) {

	ai = &Claude{
		url:          url,
		apiKey:       apiKey,
		systemPrompt: systemPrompt,
	}

	return
}

func (ai *Claude) GetApiKey() (apiKey string) {
	apiKey = ai.apiKey
	return
}

func (ai *Claude) GetSystemPrompt() (systemPrompt string) {
	systemPrompt = ai.systemPrompt
	return
}

func (ai *Claude) EvalDat(userPrompt string) (ok bool, err error) {

	//log.Printf("UserPrompt=%s", userPrompt)

	time.Sleep(time.Second * time.Duration(1+rand.Intn(4)))

	if rand.Intn(2) == 0 {
		ok = true
	}

	return
}
