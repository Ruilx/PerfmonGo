package main

import (
	"PerfmonGo/core/agentConfig"
	"fmt"
)

func main() {
	a, err := agentConfig.New("agent.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a.GetAgentName())
	fmt.Println(a.GetSubmitConfig())
}
