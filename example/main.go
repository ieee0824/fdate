package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ieee0824/fdate"
)

func main() {
	strs := []string{
		"197211",
		"19720101",
		"19800824",
		"1980824",
		"200011",
		"200021232",
	}

	for _, v := range strs {
		d, err := fdate.PickPossibleDate(v)
		if err != nil {
			log.Println(err)
		}

		bin, _ := json.MarshalIndent(d, "", "    ")

		fmt.Println(string(bin))
	}
}