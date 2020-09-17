package reddit
// this file contains the outline of the json file
// which will be parsed from an url

type Json struct {
	Data		Data		`json:"data"`
}

type Data struct {
	Dist		int		`json:"dist"`
	Children	[]Children	`json:"children"`
}

type Children struct {
	PostData	PostData	`json:"data"`
}

type PostData struct {
	Title		string		`json:"title"`
	UrlOverridden	string		`json:"url_overridden_by_dest"`
	ID		string		`json:"id"`
	Is_video	bool		`json:"is_video"`
	Is_stickied	bool		`json:"stickied"`
	Media		Media		`json:"media"`
}

type Media struct {
	RedditVideo	RedditVideo	`json:"reddit_video"`
}

type RedditVideo struct {
	Url		string		`json:"fallback_url"`
	Duration	int		`json:"duration"`
	Is_gif		bool		`json:"is_gif"`
}

type Posts []Post

type Post struct {
	Title			string
	UrlOverriddenByDest	string
	Url			string
	ID			string
	Is_video		bool
	Is_gif			bool
	Is_image		bool
}

