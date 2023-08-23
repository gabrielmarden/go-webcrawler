package util

import (
	"fmt"
	"os"
	"time"
)

//Write a slice of data into a file in the current directory. Use the prefix argument to identify your file.
//The sufix is auto generated considering the current time of the function call.
func WriteDataToFile(data []string, prefix string) {
	currentTime := time.Now()
	currentTimeFormat := currentTime.Format("20060102_150405")

	filename := fmt.Sprintf("%s_output_%s.txt", prefix, currentTimeFormat)

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("error on save file: %v", err)
		return
	}
	defer file.Close()

	for _, link := range data {
		_, err := file.WriteString(link + "\n")
		if err != nil {
			fmt.Printf("error on save file: %v", err)
			return
		}
	}

	fmt.Printf("file saved in the current directory. filename: %s", filename)
}
