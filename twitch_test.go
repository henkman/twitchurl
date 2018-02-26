package twitchurl

import (
	"fmt"
	"os"
	"testing"
)

func TestTwitch(t *testing.T) {
	clientid := os.Getenv("twitchurl")
	streams, err := GetStream(clientid, "forsen")
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	for _, s := range streams {
		fmt.Println(s)
	}
}
