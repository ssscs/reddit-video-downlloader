package info_utils

import (
	"os"
	"fmt"
)
func WriteTag(filepath, tag string) {
	os.Chdir(filepath)
	file, _ := os.OpenFile("info.txt", os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprintln(file, fmt.Sprintf("today's tag - %q", tags))
	file.Close()
}
func WriteToFile(filepath, id, title string) {
	os.Chdir(filepath)
	file, _ := os.OpenFile("info.txt", os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprintln(file, fmt.Sprintf("%q - %q", id, title))
	file.Close()
}

