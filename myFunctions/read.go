package myFunctions

import (
	"os"
	"log"
)

func Read(fileName string) (string, error) {
	
	//Open File.
    file, err := os.Open(fileName)
	if err != nil {
		log.Println("error opening file :", fileName)
		return "", err
	}

	defer file.Close()
	
	//Get file info.
    fileInfo, err := file.Stat()
    if err != nil {
        log.Println("Error getting file stats:", err)
		return "", err
    }

    //Get file size.
    fileSize := fileInfo.Size()
    data := make([]byte, fileSize)

	//Reading data.
    _, err = file.Read(data)
    if err != nil {
		log.Println("Error reading the file:", err)
		return "", err
    }
	return string(data), nil
}
