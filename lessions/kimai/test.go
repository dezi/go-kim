package kimai

import (
	"encoding/json"
	"github.com/dezi/go-kim/lessions/ifaces/claude"
	"github.com/dezi/go-kim/lessions/ifaces/openai"
	"github.com/dezi/go-kim/utils/log"
	"github.com/yunabe/easycsv"
	"os"
	"sync"
)

var wg sync.WaitGroup
var userChannel = make(chan *UserRecord, 100)
var results = make(map[int]*ResultRecord)
var resultsLock sync.Mutex

func DoKimAI() {

	wg.Add(1)
	go CsvReader()

	for slot := 0; slot < 100; slot++ {
		wg.Add(1)
		go Evaluator(slot)
	}

	wg.Wait()

	log.Printf("Auswertung fettich.")

	jsonBytes, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Cerror(err)
		return
	}

	err = os.WriteFile("./X_test_results.json", jsonBytes, 0640)
	if err != nil {
		log.Cerror(err)
		return
	}
}

func Evaluator(slot int) {

	defer wg.Done()

	openAiInstance := openai.CreateInstance()
	claudeInstance := claude.CreateInstance()

	for {

		var ok bool
		var err error
		var current *UserRecord

		current, ok = <-userChannel

		if !ok {
			break
		}

		log.Printf("Process slot=%d id=%d", slot, current.Id)

		openAIOk, err := openAiInstance.EvalDat(current.Prompt)
		if err != nil {
			log.Cerror(err)
			return
		}

		log.Printf("OpenAi slot=%d id=%d ok=%v Userprompt=%s", slot, current.Id, openAIOk, current.Prompt[:20])

		claudeOk, err := claudeInstance.EvalDat(current.Prompt)
		if err != nil {
			log.Cerror(err)
			return
		}

		log.Printf("Claude slot=%d id=%d ok=%v Userprompt=%s", slot, current.Id, openAIOk, current.Prompt[:20])

		result := ResultRecord{
			Id:     current.Id,
			OpenAi: openAIOk,
			Claude: claudeOk,
		}

		resultsLock.Lock()
		results[current.Id] = &result
		resultsLock.Unlock()
	}
}

func CsvReader() {

	defer close(userChannel)
	defer wg.Done()

	csvReader := easycsv.NewReaderFile("./X_test_vignette_prompt.csv")

	for {

		//
		// Create a NEW record for each line from CSV!!!!
		//

		var myEntry UserRecord
		if !csvReader.Read(&myEntry) {
			break
		}

		userChannel <- &myEntry

		log.Printf("Id=%d", myEntry.Id)
	}

	if err := csvReader.Done(); err != nil {
		log.Cerror(err)
		return
	}
}
