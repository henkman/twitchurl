package twitchurl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

type Stream struct {
	Name string
	Url  string
}

var (
	reExtractM3U8 = regexp.MustCompile(
		`(?s)#EXT-X-MEDIA:TYPE=.*?NAME="([^"]+)".*?(http://.*?)(?:\n|$)`)
)

func GetStream(clientid, channel string) ([]Stream, error) {
	cli := http.Client{
		Timeout: 10 * time.Second,
	}
	var token struct {
		Token            string `json:"token"`
		Sig              string `json:"sig"`
		MobileRestricted bool   `json:"mobile_restricted"`
	}
	{
		u := fmt.Sprintf("http://api.twitch.tv/api/channels/%s/access_token",
			channel)
		req, err := http.NewRequest("GET", u, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Client-Id", clientid)
		res, err := cli.Do(req)
		if err != nil {
			return nil, err
		}
		err = json.NewDecoder(res.Body).Decode(&token)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
	}
	u := fmt.Sprintf("http://usher.twitch.tv/api/channel/hls/%s.m3u8",
		channel)
	vals := url.Values{
		"player":           {"twitchweb"},
		"token":            {token.Token},
		"sig":              {token.Sig},
		"allow_audio_only": {"true"},
		"allow_source":     {"true"},
		"type":             {"any"},
		"p":                {fmt.Sprint(rand.Int31n(10000000))},
	}
	res, err := cli.Get(u + "?" + vals.Encode())
	if err != nil {
		return nil, err
	}
	m3u8, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, nil
	}
	m := reExtractM3U8.FindAllStringSubmatch(string(m3u8), -1)
	if m == nil {
		return nil, errors.New("m3u8 is invalid")
	}
	ss := make([]Stream, 0, len(m))
	for _, t := range m {
		ss = append(ss, Stream{Name: t[1], Url: t[2]})
	}
	return ss, nil
}
