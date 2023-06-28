package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func RecordLog(errSend error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}

	fileTime := time.Now().Format("02-01-06-15-04")
	fileName := fmt.Sprintf("%s.txt", fileTime)
	adjacentFolder := filepath.Join(dir, "../src/", "logs")

	if _, err := os.Stat(adjacentFolder); os.IsNotExist(err) {
		err := os.Mkdir(adjacentFolder, 0755)
		if err != nil {
			log.Println(err)
			return
		}
	}

	filePath := filepath.Join(adjacentFolder, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Error on create: ", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	errorBytes := []byte(errSend.Error())
	_, err = writer.WriteString(string(errorBytes))
	if err != nil {
		log.Println("Error on write: ", err)
		return
	}

	err = writer.Flush()
	if err != nil {
		log.Println("Error on flush: ", err)
		return
	}
}