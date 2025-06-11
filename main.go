package main

import (
	"github.com/dezi/go-kim/lessions/types"
	"github.com/dezi/go-kim/utils/log"
)

func main() {

	log.Printf("Hello world started...")
	defer log.Printf("Hello world finished.")

	types.Test()
}
