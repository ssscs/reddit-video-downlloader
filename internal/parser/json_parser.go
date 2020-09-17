// Package json provides functions for reading JSON documents from an URL
package parser

import (
	"net/http"
	"encoding/json"
	"redditdownloader/api/json_schemes/reddit"
	"io/ioutil"
	"log"
	"strings"
)

// ReadUrl reads the body from the provided URL and returns body's elements in bytes
func ReadUrl(url string) (body []byte){
	client := &http.Client{}

	req, errReq := http.NewRequest("GET", url, nil)
	//setting a header to overcome 429 error
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Safari/537.36 OPR/54.0.2952.64")
	//getting the response from the url
	resp, errResp := client.Do(req)
	defer resp.Body.Close()

	if errReq != nil || errResp != nil{
		log.Fatal("Cannot read data from the provided URL ", errReq, errResp)
	}
	//reading bytes from the http response
	body, _ = ioutil.ReadAll(resp.Body)
	return
}

// ParseJson parses JSON documents from provided url and returns an array of map which includes posts details
// this function can parse any post (includes image-.jpg-, video-.mp4-, gif) from 
// any subreddit in the matter of a second.
func ParseJson(url string) reddit.Posts {
	body := ReadUrl(url)
	var j reddit.Json
	var posts reddit.Posts
	//encoding the bytes to create json
	json.Unmarshal(body,  &j)
	// j.Data.Dist is the post count including moderator posts
	for i := 0; i < j.Data.Dist;  i++{
		p := &j.Data.Children[i].PostData
		var post reddit.Post
		// checking if the post belongs to a moderator or a user
		// stickied posts are from the moderators and they explain rules of the subreddit
		// so we dont need to include these posts in our json map
		if p.Is_stickied != true {
			post = reddit.Post {Title: p.Title, ID: p.ID}
			// post is a video
			if p.Is_video == true {
				// video duration must be lower than 60
				if p.Media.RedditVideo.Duration <= 60 {
					// set the url
					post.Url = p.Media.RedditVideo.Url
					if p.Media.RedditVideo.Is_gif != true {
					// video with an audio
						post.Is_video = p.Is_video
						// appending media with the new posts details
						posts = append(posts, post)
					} else {
						// gif doesnt have an audio
						post.Is_gif = p.Media.RedditVideo.Is_gif
						// appending media with the new posts details
						posts = append(posts, post)
					}
				}
			} else {
				if strings.Contains(p.UrlOverridden, ".jpg") {
					post.Is_image = true
					post.UrlOverriddenByDest = p.UrlOverridden
					// appending media with the new posts details
					posts = append(posts, post)
				}
			}	
		}
	}
	return posts
}
func EncodeJson(posts *reddit.Posts) []byte {
    encodedJ, err := json.Marshal(posts)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
		return nil
	}
	return encodedJ
}
func main(){
	url := "https://www.reddit.com/r/Unexpected/.json"
	ParseJson(url)
}
