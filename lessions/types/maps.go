package types

import "github.com/dezi/go-kim/utils/log"

var (
	myStringMap  map[string]string
	myString2Map = map[string]map[string]string{
		"a": {
			"ssss": "4",
		},
	}

	myString3Map = make(map[string]map[string]string)

	myString7Map = map[string]string{
		"a": "1",
	}
)

func Test5() {

	myString2Map = make(map[string]map[string]string)

	myString2Map["2"] = make(map[string]string)
	myString2Map["2"]["3"] = "xxxx"

	sMap := myString2Map["2"]
	log.Printf("sMap=%v", sMap)

	sVal, ok := myString2Map["2"]["3"]
	log.Printf("sVal=%s ok=%v", sVal, ok)

	for key1, val := range myString2Map {
		for key2, str := range val {
			log.Printf("key1=%s key2=%s str=%s", key1, key2, str)
		}
	}
}

func Test4() {

	myStringMap = make(map[string]string)

	myStringMap["a"] = "1"
	myStringMap["b"] = "2"
	myStringMap["c"] = "3"
	myStringMap["d"] = "4"

	delete(myStringMap, "c")
	delete(myStringMap, "x")

	log.Printf("---------------")
	for key, val := range myStringMap {
		log.Printf("Key=%s val=%s", key, val)
	}

	dVal, dOk := myStringMap["d"]
	xVal, xOk := myStringMap["x"]

	if xVal, xOk := myStringMap["x"]; xOk {
		log.Printf("x is present = %s", xVal)
	}

	log.Printf("dVal=%s dOk=%v", dVal, dOk)
	log.Printf("xVal=%s xOk=%v", xVal, xOk)
}
