package ifaces

import (
	"github.com/dezi/go-kim/lessions/ifaces/claude"
	"github.com/dezi/go-kim/lessions/ifaces/openai"
	"github.com/dezi/go-kim/utils/log"
)

func TestIFace() {

	openAiInstance := openai.CreateInstance()
	claudeInstance := claude.CreateInstance()

	userPrompt := "Ich nix wiss"

	var ok bool
	var err error

	ok, err = openAiInstance.EvalDat(userPrompt)
	if err != nil {
		log.Cerror(err)
		return
	}

	log.Printf("OpenAI Userprompt=%s ok=%v", userPrompt, ok)

	ok, err = claudeInstance.EvalDat(userPrompt)
	if err != nil {
		log.Cerror(err)
		return
	}

	log.Printf("Claude Userprompt=%s ok=%v", userPrompt, ok)
}
