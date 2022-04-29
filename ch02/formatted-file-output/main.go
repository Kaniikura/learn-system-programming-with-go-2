package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "今日の天気は%sで、最高気温は%.1f度、湿度は%dパーセントです。\n", "晴れ", 20.5, 50)
}
