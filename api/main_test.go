package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/renatormc/pfila/api/external"
)

func TestMain(t *testing.T) {
	disks, err := external.GetDisks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(disks)
}
