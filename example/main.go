package main

import (
	"fmt"
	"log"

	"github.com/ieee0824/fdate"
)

func main() {
	t, err := fdate.PickPossibleDate("2018ヰ10月2211")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(t)
}
