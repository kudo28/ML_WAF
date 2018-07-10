package main

import (
	"fmt"
	"git.cyradar.com/utilities/data/timer"
	"time"
	"log"
)

var logger = new(log.Logger)

func main() {
	defer timer.TimeTrack(time.Now(), fmt.Sprintf("Extract features"))
	apiFile, e := NewLogger(logger, "dataset.csv")
	if e != nil {
		return
	}
	defer apiFile.Close()
	//for _, url := range list {
	//	fmt.Println(getLength(url))
	//}

	header := "URL,label,length"
	for _, token := range Tokens {
		header = fmt.Sprintf("%s,%s", header, token)
	}
	header = fmt.Sprintf("%s,/**/", header)
	logger.Printf("%s\n", header)
	processNormal("normalTrafficTraining.txt")
	processMalicious("raw_payload_train.txt")
	//processMalicious("anomalousTrafficTest.txt")
}
