# fdate
ゆるい日付ライブラリ

# example 

```
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
		"2010年1月1日",
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

```


```
$ go run example/main.go
[
    "1972-01-01T00:00:00+09:00"
]
[
    "1972-01-01T00:00:00+09:00"
]
[
    "1980-08-02T00:00:00+09:00",
    "1980-08-24T00:00:00+09:00"
]
[
    "1980-08-02T00:00:00+09:00",
    "1980-08-24T00:00:00+09:00"
]
[
    "2000-01-01T00:00:00+09:00"
]
[
    "2000-02-02T00:00:00+09:00",
    "2000-01-23T00:00:00+09:00",
    "2000-01-03T00:00:00+09:00",
    "2000-12-03T00:00:00+09:00",
    "2000-02-01T00:00:00+09:00",
    "2000-02-12T00:00:00+09:00",
    "2000-02-23T00:00:00+09:00",
    "2000-02-03T00:00:00+09:00",
    "2000-01-02T00:00:00+09:00"
]
[
    "2010-01-01T00:00:00+09:00"
]
```