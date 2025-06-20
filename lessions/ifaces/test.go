package ifaces

import (
	"github.com/dezi/go-kim/lessions/ifaces/claude"
	"github.com/dezi/go-kim/lessions/ifaces/openai"
	"github.com/dezi/go-kim/utils/log"
	"sync"
)

func TestIFace() {

	openAiInstance := openai.CreateInstance()
	claudeInstance := claude.CreateInstance()

	wg := sync.WaitGroup{}

	userPrompt := "Ich nix wiss"

	wg.Add(1) // important: before go function!!!
	go func(userPrompt string) {
		defer wg.Done()

		ok, err := openAiInstance.EvalDat(userPrompt)
		if err != nil {
			log.Cerror(err)
			return
		}

		log.Printf("OpenAI Userprompt=%s ok=%v", userPrompt, ok)

	}(userPrompt)

	wg.Add(1) // important: before go function!!!
	go func(userPrompt string) {
		defer wg.Done()

		ok, err := claudeInstance.EvalDat(userPrompt)
		if err != nil {
			log.Cerror(err)
			return
		}

		log.Printf("Claude Userprompt=%s ok=%v", userPrompt, ok)

	}(userPrompt)

	wg.Wait()
}
