package marshal

import (
	"encoding/json"
	"github.com/dezi/go-kim/utils/log"
)

type MyMarshal struct {
	unExported string
	Geheim     string   `json:"-"`
	UserId     int64    `json:"user_id"`
	UserPrompt string   `json:"user_prompt"`
	NixDrin    string   `json:"user_nixdrin,omitempty"`
	Diagnosis  []string `json:"diag"`
}

func (ms *MyMarshal) Export2Json() (jsonStr string, err error) {

	jsonBytes, err := json.MarshalIndent(ms, "", "  ")
	if err != nil {
		log.Cerror(err)
		return
	}

	jsonStr = string(jsonBytes)
	return
}

func Test9() {

	ms := MyMarshal{
		unExported: "secretApiKey",
		Geheim:     "geheim",
		UserId:     4711,
		UserPrompt: "Ich werde blöd....",
		Diagnosis: []string{
			"Paranoide Schizophrenie",
			"Alzheimer",
			"ADHS",
		},
	}

	jStr, err := ms.Export2Json()
	if err != nil {
		log.Cerror(err)
		return
	}

	log.Printf("Marshall\n%s", jStr)

	var msNew MyMarshal

	err = msNew.ImportFromJson(jStr)
	if err != nil {
		log.Cerror(err)
		return
	}

	log.Printf("UnMarshall msNew=%+v", msNew)

	err = msNew.ImportFromJson(jStr)
	if err != nil {
		log.Cerror(err)
		return
	}

}

func (ms *MyMarshal) ImportFromJson(jsonStr string) (err error) {
	err = json.Unmarshal([]byte(jsonStr), ms)
	return
}
