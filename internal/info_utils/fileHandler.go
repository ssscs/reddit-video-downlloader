package info_utils

import (
	"os"
	"time"
)

func CreateFile(filepath string) {
	_, err := os.Stat(filepath+"/info.txt")

    if err != nil {
        if os.IsNotExist(err) {
        	os.Chdir(filepath)
            os.Create("info.txt")
        }
    }
}


func CreateBaseFolder(filepath , shortlink string) string{
	year := time.Now().Format("2006")
	day := time.Now().Format("01.02")
	filepath += year + "/" + day + "/" + shortlink
	os.MkdirAll(filepath, 0755)
	return filepath
}