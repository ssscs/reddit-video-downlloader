package wrapper

import (
	"redditdownloader/internal/parser"
	"redditdownloader/internal/media_utils"
	"redditdownloader/internal/info_utils"
	"redditdownloader/internal/url_provider"
	"redditdownloader/api/json_schemes/reddit"
	"redditdownloader/tools/timeT"
	"os/user"
	"os"
	"fmt"
	"time"
)

func Hack(shortlink string) {
	// every url has a shortlink that assigned to its absolute url
	// absolute url gathered from url provider
	url := url_provider.GetUrl(shortlink)
	
	// list of posts
	parsestart := time.Now()
	posts := parser.ParseJson(url)
	timeT.TimeTrack(parsestart, "parse")

    user, _:= user.Current()
	filepath := createFolders(user.HomeDir + "/RedditDownloader/", shortlink)
	

	info_utils.WriteTag(filepath, info_utils.GetTag())

	fmt.Println("[info] download started..")
	downloadstart := time.Now()
	for _, post := range posts {
		download(filepath, post)
		info_utils.WriteToFile(filepath, post.ID, post.Title)
	}
	timeT.TimeTrack(downloadstart, "download")
}

func createFolders(filepath, shortlink string) string {
	newfilepath := info_utils.CreateBaseFolder(filepath, shortlink)
	info_utils.CreateFile(newfilepath)
	return newfilepath
}

func download(filepath string, post reddit.Post) {
	os.Chdir(filepath)
	os.MkdirAll("media/"+ post.ID, 0755)

	filepath += "/media/"+ post.ID 
	if post.Is_video == true {
		media_utils.DownloadVideo(filepath, post.Url)
		media_utils.ScaleVideo(filepath)
		media_utils.WatermarkVideo(filepath)
	} else if post.Is_gif == true {
		media_utils.DownloadGif(filepath, post.Url)
		media_utils.ScaleVideo(filepath)
		media_utils.WatermarkVideo(filepath)
	} else if post.Is_image == true {
		media_utils.DownloadImage(filepath, post.UrlOverriddenByDest)
	}
}