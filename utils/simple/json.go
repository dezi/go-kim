package simple

import (
	"encoding/json"
	"os"
	"strings"
	"tch-go-coaching-dezi/utils/log"
)

func UnmarshalJsonFile(filePath string, data interface{}) (err error) {

	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonBytes, data)
	if err != nil {
		log.Cerror(err)
		return
	}

	return
}

func MarshalJsonCleanFile(data interface{}, file string) (err error) {

	bytes, err := MarshalJsonClean(data)
	if err != nil {
		return
	}

	err = os.WriteFile(file, bytes, 0644)
	return
}

func MarshalJsonClean(data interface{}) (clean []byte, err error) {

	clean, err = json.MarshalIndent(data, "", "  ")
	if err != nil {
		return
	}

	//
	// Get rid of ampersand UTF-code.
	//

	clean = []byte(MarshalDefuck(string(clean)))
	return
}

func MarshalDefuck(fucked string) (unFucked string) {
	unFucked = strings.Replace(fucked, "\\u0026", "&", -1)
	unFucked = strings.Replace(unFucked, "\\u003c", "<", -1)
	unFucked = strings.Replace(unFucked, "\\u003e", ">", -1)
	return
}

func ReMarshal(generic interface{}, specific interface{}) (err error) {

	jsonBytes, err := json.Marshal(generic)
	if err != nil {
		log.Cerror(err)
		return
	}

	err = json.Unmarshal(jsonBytes, specific)
	if err != nil {
		log.Cerror(err)
		return
	}

	return
}
