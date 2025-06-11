package gosubs

import (
	"github.com/dezi/go-kim/utils/log"
	"math/rand"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func Test10() {

	for inx := 0; inx < 10; inx++ {
		waitGroup.Add(1)
		go DoRandom(inx, 5+rand.Intn(5))
	}

	waitGroup.Wait()
}

func DoRandom(slot, max int) {

	log.Printf("Slot %d started.", slot)
	defer log.Printf("Slot %d done.", slot)

	defer waitGroup.Done()

	for ; max >= 0; max-- {
		log.Printf("Slot=%d Random=%d", slot, rand.Intn(100))
		time.Sleep(time.Second)
	}
}
