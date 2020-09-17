package media_utils

import (
	"os/exec"
	"log"
)

func WatermarkImage(filepath, url string) error{
	cmd := exec.Command("ffmpeg", "-i", url, "-i", "/home/eniax/go/src/redditdownloader/assets/logo.png", "-filter_complex",
		"overlay=main_w-overlay_w-5:main_h-overlay_h-5", 
		"output.jpg")
	cmd.Dir = filepath
/*	out, err := cmd.CombinedOutput()
	log.Printf(string(out))*/
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	return err
}	
/*func main() {
	filepath := "/home/eniax/Downloads/"
	WatermarkImage(filepath)
}*/