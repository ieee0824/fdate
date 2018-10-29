package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ieee0824/fdate"
)

func main() {
	strs := []string{
		"201811282930",
	}

	for _, v := range strs {
		d, err := fdate.PickPossibleDate(v)
		if err != nil {
			log.Println(err)
		}

		bin, _ := json.MarshalIndent(d, "", "    ")

		fmt.Print(v, " = ")
		fmt.Println(string(bin), "\n")
	}
}
