package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ieee0824/fdate"
)

func main() {
	log.SetFlags(log.Llongfile)
	d, err := fdate.PickPossibleDate("1991/12/24")
	if err != nil {
		log.Println(err)
	}

	bin, _ := json.MarshalIndent(d, "", "    ")

	fmt.Println(string(bin))
}
