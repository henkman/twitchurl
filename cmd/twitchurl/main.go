package main

import (
	"fmt"
	"os"

	"github.com/henkman/twitchurl"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: twitchurl channel [quality]")
		return
	}
	channel := os.Args[1]
	clientid := os.Getenv("twitchurl")
	streams, err := twitchurl.GetStream(clientid, channel)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(os.Args) == 3 {
		quality := os.Args[2]
		for _, s := range streams {
			if s.Name == quality {
				fmt.Println(s.Url)
				return
			}
		}
		fmt.Println("quality", quality, "not found. listing all")
	}
	for _, s := range streams {
		fmt.Println(s.Name)
	}
}
