package objects

import "github.com/dezi/go-kim/utils/log"

type MyNixDrin struct {
}

type MySuperClass struct {
	privateVar int
	PublicVar  int
	lastAddRes int
}

func (nix *MyNixDrin) Add() (addRes int) {
	return
}

func (msc *MySuperClass) Add() (addRes int) {
	addRes = msc.privateVar + msc.PublicVar
	msc.lastAddRes = addRes
	return
}

func (msc *MySuperClass) Sub() (addRes int) {
	addRes = msc.privateVar - msc.PublicVar
	return
}

func Test6() {

	msc := &MySuperClass{
		privateVar: 42,
		PublicVar:  0,
		lastAddRes: -1,
	}

	msc2 := MySuperClass{
		privateVar: 42,
		PublicVar:  0,
		lastAddRes: -1,
	}

	res := msc.Add()
	log.Printf("##### res=%d lastAddRes=%d", res, msc.lastAddRes)

	msc.PublicVar = 22

	res = msc.Add()
	log.Printf("##### res=%d lastAddRes=%d", res, msc.lastAddRes)

	res = msc2.Add()

}
