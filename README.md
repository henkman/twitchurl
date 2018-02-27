# twitchurlextract
small package to get stream urls and util to print to console or pipe into video player

## Usage
first you need to get a twitch client id by creating a twitch application and set it as environment variable "twitchurl"

```
$ go get github.com/henkman/twitchurl/cmd/twitchurl
$ twitchurl forsen 720p | vlc --no-video-title-show -
```
