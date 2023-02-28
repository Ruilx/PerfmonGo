package main

import (
	"PerfmonGo/base/config"
	"fmt"
)

func main() {
	a, err := config.New("agent.json")
	if err != nil {
		fmt.Println(err)
	}
	a.FindKey("a", "b", "c")
}
