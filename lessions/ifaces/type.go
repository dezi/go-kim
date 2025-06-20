package ifaces

type KimAI interface {
	GetApiKey() (apiKey string)
	GetSystemPrompt() (systemPrompt string)
	EvalDat(userPrompt string) (ok bool, err error)
}
