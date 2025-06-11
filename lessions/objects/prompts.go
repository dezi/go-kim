package objects

type EvalPrompt struct {
	claudiApiKey string
	UserInfo     string
	ClaudiRes    string
	GeminiRes    string
	OpenAiRes    string
}

func Eval(userInfo string) (ep *EvalPrompt, err error) {

	ep = &EvalPrompt{
		claudiApiKey: "336427344234",
		UserInfo:     userInfo,
	}

	_ = ep.EvalClaudi()
	_ = ep.EvalGemini()

	return
}

func (ep *EvalPrompt) EvalGemini() (err error) {

	ep.ClaudiRes = "xsyxasdsdads" // todo fetch claudi response.

	return
}

func (ep *EvalPrompt) EvalClaudi() (err error) {

	ep.ClaudiRes = "xsyxasdsdads" // todo fetch claudi response.

	return
}
