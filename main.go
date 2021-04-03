package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/maszuari/http-client-ex-v2/model"
	"github.com/maszuari/http-client-ex-v2/util"
)

func main() {
	fmt.Println("Start...")
	start := time.Now()
	url := "https://test.spaceflightnewsapi.net/api/v2/articles?_limit=10"
	out, err := util.FetchData(url)
	if err != nil {
		log.Fatal(err)
	}
	arr, err := model.UnmarshalJSON(out)
	if err != nil {
		log.Fatal(err)
	}

	sli := make([]string, 5)
	for _, obj := range arr {
		sli = append(sli, obj.Title)
		//fmt.Println(obj.Title)
	}
	sort.Strings(sli)
	for _, elm := range sli {
		fmt.Println(elm)
	}
	end := time.Now()
	val := end.Sub(start)
	fmt.Println("End ", val)
}
