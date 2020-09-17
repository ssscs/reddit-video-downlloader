package media_utils

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"log"
)

// Downloader provides functions that will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadGif(filepath string, url string) {
	// Get the data
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	// Create the file
	os.Chdir(filepath)
	defer os.Chdir("$GOPATH")
	out, _ := os.Create("readyToScale.mp4")
	defer out.Close()

	// Write the body to file
	io.Copy(out, resp.Body)
}

func DownloadImage(filepath string, url string) {
	WatermarkImage(filepath, url)
}

func DownloadVideo(filepath string, url string) {
	cmd := exec.Command("ffmpeg", "-i", url, "-i", strings.Split(url, "DASH_")[0]+"DASH_audio.mp4", "-c:v", "copy", "-c:a", "aac", "readyToScale.mp4")
	cmd.Dir = filepath
	/*out, err := cmd.CombinedOutput()
	log.Printf(string(out))*/
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
}

/*func main() {
	url := "https://v.redd.it/oh6q9uq5xzm51/DASH_720.mp4?source=fallback"
	filepath := "/home/eniax/Downloads"
	usr, _ := user.Current()
	log.Println(usr.HomeDir)	
	DownloadGif(filepath, url)
}*/
