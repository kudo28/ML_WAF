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
	apiFile, e := NewLogger(logger, "normal.csv")
	if e != nil {
		return
	}
	defer apiFile.Close()
	//for _, url := range list {
	//	fmt.Println(getLength(url))
	//}
	processNormal()
	processMalicious()
}
