package channels

import (
	"github.com/dezi/go-kim/utils/log"
	"sync"
	"time"
)

var xChan chan int
var yChan = make(chan int)
var wg sync.WaitGroup

func TestChan() {

	myChan := make(chan int, 10)

	wg.Add(1)
	go consumer(myChan)

	log.Printf("vor reinstopf...")
	myChan <- 5
	log.Printf("nach reinstopf a...")
	myChan <- 55
	log.Printf("nach reinstopf b...")
	myChan <- 56
	myChan <- 57
	myChan <- 58
	log.Printf("nach reinstopf b...")

	close(myChan)

	wg.Wait()
}

func consumer(c chan int) {

	defer wg.Done()

	for {
		time.Sleep(time.Second)
		hubu, ok := <-c
		if !ok {
			log.Printf("channel ist zu, danke...")
			break
		}
		log.Printf("##### hubu=%d", hubu)
	}
}
