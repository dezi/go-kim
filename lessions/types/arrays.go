package types

import (
	"errors"
	"github.com/dezi/go-kim/utils/log"
)

var (
	intArray  []int
	int4Array [4]int

	// java: int[4] int4Array
)

func Test3() {

	int4Array[0] = 12
	int4Array[1] = 13
	int4Array[2] = 14
	int4Array[3] = 15

	log.Printf("int4Array=%v", int4Array)

	//index := 5 * 3
	//int4Array[index] = 16

	intArray = make([]int, 5)

	if intArray == nil {
		err := errors.New("intArray is nix drin")
		log.Cerror(err)
		return
	}

	for inx := 0; inx < len(intArray); inx++ {
		intArray[inx] = inx
	}

	for inx, oldVal := range intArray {
		intArray[inx] = inx + oldVal
	}

	for inx := range intArray {
		intArray[inx] = inx
	}

	intArray = append(intArray, int4Array[:]...)

	log.Printf("intArray=%v", intArray)

	myArray := make([]int, 10)
	mySlice := myArray[3:5]

	mySlice[0] = 3
	mySlice[1] = 4

	log.Printf("myArray=%v", myArray)
	log.Printf("mySlice=%v", mySlice)

	myString := "                                   +++++"

	for {
		if len(myString) == 0 || myString[0] != ' ' {
			break
		}

		myString = myString[1:]
		myString = myString + "*"
	}

	log.Printf("myString=%s", myString)
}
