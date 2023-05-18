package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/renatormc/pfila/api/processes/ftkimager"
)

func TestMain(t *testing.T) {
	disks, err := ftkimager.GetDisks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(disks)
}
