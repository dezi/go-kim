package gosubs

import (
	"github.com/dezi/go-kim/utils/log"
	"math/rand"
	"sync"
)

var waitGroup sync.WaitGroup
var myMap = make(map[int]bool)
var myMapLock sync.Mutex

func Test10() {

	for inx := 0; inx < 10; inx++ {
		waitGroup.Add(1)
		go DoRandom(inx, 500+rand.Intn(500))
	}

	waitGroup.Wait()

	log.Printf("all gosubs done...")
}

func DoRandom(slot, max int) {

	log.Printf("Slot %d started.", slot)
	defer log.Printf("Slot %d done.", slot)

	defer waitGroup.Done()

	for ; max >= 0; max-- {
		myVal := rand.Intn(100)
		myMapLock.Lock()
		myMap[myVal] = true
		myMapLock.Unlock()
		if myVal%100 == 55 {
			log.Printf("Slot=%d Random=%d", slot, myVal)
		}
	}
}
