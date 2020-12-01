package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main()  {
	var movies = []Movie{
		{Title:"Casablanca", Year:1942, Color:false, Actors:[]string{"Hjadfajd", "Adjflajdl"}},
		{Title:"Dasablanca", Year:1943, Color:true, Actors:[]string{"Hjadfajd"}},
		{Title:"Easablanca", Year:1944, Color:false, Actors:[]string{"Ejadfajd", "Fdjflajdl"}},
	}
	fmt.Printf("%#v\n", movies)
	data, err := json.Marshal(movies) // 字节 slice,
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	data2, err := json.MarshalIndent(movies, "", "	") // 输出整齐格式化的结果, 第二个定义每行输出的前缀字符串, 第三个是定义缩进的字符串.
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data2)

	var titles []struct{Title string}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool  `json:"color,omitempty"` // omitempty 如果为空忽略
	Actors []string
}

