package url_provider

var urls = map[string]string{
	"uvid": "https://www.reddit.com/r/Unexpected/.json",
	"wpdi": "https://www.reddit.com/r/WatchPeopleDieInside/.json",
	"pcut": "https://www.reddit.com/r/perfectlycutscreams/.json",
	"nfkl": "https://www.reddit.com/r/nextfuckinglevel/.json",
	"funy": "https://www.reddit.com/r/funny/.json",
	"iasf": "https://www.reddit.com/r/interestingasfuck/.json",
}

func GetUrl(shortlink string) string {
	return urls[shortlink]
}