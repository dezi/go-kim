package maps

import (
	"github.com/dezi/go-kim/utils/log"
	"sync"
	"time"
)

func TestMaps() {

	//myMap := make(map[string]string)

	myMap := map[string]string{
		"a": "abc",
		"b": "bcd",
		"c": "cde",
		"d": "def",
	}

	myMutex := sync.Mutex{}

	myFunc := func() {
		for {
			myMutex.Lock()
			for key, val := range myMap {
				myMap[key+"*"] = val + "*"
				log.Printf("Innnen Key=%s val=%s new=%s", key, val, myMap[key+"*"])
				time.Sleep(time.Millisecond * 100)
			}
			myMutex.Unlock()
		}
	}

	go myFunc()

	for {
		log.Printf("---------------")
		myMutex.Lock()
		for key, val := range myMap {
			log.Printf("Aussen Key=%s val=%s", key, val)
			time.Sleep(time.Millisecond * 100)
		}
		myMutex.Unlock()
	}
}
