package media_utils

import ( 
	"os"
	"os/exec"
	"log"
)

func ScaleVideo(filepath string) {
	cmd := exec.Command("ffmpeg", "-i", "readyToScale.mp4", "-vf", 
		"split[original][copy];[copy]scale=ih*16/9:-1,crop=h=iw*9/16,gblur=sigma=20[blurred];[blurred][original]overlay=(main_w-overlay_w)/2:(main_h-overlay_h)/2",
		"readyToWM.mp4")
	cmd.Dir = filepath
	/*out, err := cmd.CombinedOutput()
	log.Printf(string(out))*/
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
}

func WatermarkVideo(filepath string) {
	cmd := exec.Command("ffmpeg", "-i", "readyToWM.mp4", "-i", "/home/eniax/go/src/redditdownloader/assets/logo.png", "-filter_complex",
		"overlay=main_w-overlay_w-5:main_h-overlay_h-5", "-codec:a", "copy", 
		"output.mp4")
	cmd.Dir = filepath
/*	out, err := cmd.CombinedOutput()
	log.Printf(string(out))*/
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
	deleteOld(filepath)
}

func deleteOld(filepath string) {
	os.Chdir(filepath)
	os.Remove("readyToScale.mp4")
	os.Remove("readyToWM.mp4")
}
/*func main() {
	filepath := "/home/eniax/Downloads/"
	ScaleVideo(filepath)
	WatermarkVideo(filepath)
}*/
