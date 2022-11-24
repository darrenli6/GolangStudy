package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"xoinye", "darren"}}

	// 编码过程
	jsonstr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("jsonstr = %s \n", jsonstr)

	// 解码过程
	//jsonstr = {"title":"喜剧之王","year":2000,"rmb":10,"actors":["xoinye","darren"]}
	my_movie := Movie{}

	err = json.Unmarshal(jsonstr, &my_movie)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Printf("%v\n", my_movie)

}
