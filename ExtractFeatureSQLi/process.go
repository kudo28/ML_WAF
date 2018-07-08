package main

import (
	"os"
	"fmt"
	"io"
	"log"
	"io/ioutil"
	"regexp"
	"strings"
)

func caseInsenstiveContains(a, b string) bool {
	return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}

var getURL = regexp.MustCompile("GET (http://.*) HTTP")

func getLength(s string) int {
	return len(s)
}
func getTableHandling(raw string) string {
	s := ""

	bol := caseInsenstiveContains(raw, "SELECT")
	if bol == true {
		s = "1"
	} else {
		s = "0"
	}
	bol = caseInsenstiveContains(raw, "INSERT")
	if bol == true {
		s = fmt.Sprintf("%s,1", s)
	} else {
		s = fmt.Sprintf("%s,0", s)
	}
	bol = caseInsenstiveContains(raw, "UPDATE")
	if bol == true {
		s = fmt.Sprintf("%s,1", s)
	} else {
		s = fmt.Sprintf("%s,0", s)
	}
	bol = caseInsenstiveContains(raw, "DELETE")
	if bol == true {
		s = fmt.Sprintf("%s,1", s)
	} else {
		s = fmt.Sprintf("%s,0", s)
	}
	bol = caseInsenstiveContains(raw, "CREATE")
	if bol == true {
		s = fmt.Sprintf("%s,1", s)
	} else {
		s = fmt.Sprintf("%s,0", s)
	}
	bol = caseInsenstiveContains(raw, "DROP")
	if bol == true {
		s = fmt.Sprintf("%s,1", s)
	} else {
		s = fmt.Sprintf("%s,0", s)
	}
	bol = caseInsenstiveContains(raw, "AFTER")
	if bol == true {
		s = fmt.Sprintf("%s,1", s)
	} else {
		s = fmt.Sprintf("%s,0", s)
	}
	bol = caseInsenstiveContains(raw, "RENAME")
	if bol == true {
		s = fmt.Sprintf("%s,1", s)
	} else {
		s = fmt.Sprintf("%s,0", s)
	}
	bol = caseInsenstiveContains(raw, "WHERE")
	if bol == true {
		s = fmt.Sprintf("%s,1", s)
	} else {
		s = fmt.Sprintf("%s,0", s)
	}
	return s
}
func NewLogger(logger *log.Logger, logFile string) (*os.File, error) {
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		if _, err := os.Create(logFile); err != nil {
			return nil, fmt.Errorf("error creating log file: %s", err)
		}
	}

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening log file: %s", err)
	}

	logWriter := io.MultiWriter(file)

	logger.SetOutput(logWriter)
	//logger.SetFlags(log.Ldate , log.Ltime)
	return file, nil
}

func processNormal() {
	dat, err := ioutil.ReadFile("normalTrafficTraining.txt")
	if err != nil {
		log.Fatal("cannot read file")
	}
	list := make([]string, 0)
	raw := string(dat)
	result := getURL.FindAllStringSubmatch(raw, -1)
	for _, match := range result {
		list = append(list, match[1])
		raw := match[1]
		logger.Printf("%s,%d,%d,%s", raw, 0, getLength(raw), getTableHandling(raw))
	}
}

func processMalicious() {
	dat, err := ioutil.ReadFile("anomalousTrafficTest.txt")
	if err != nil {
		log.Fatal("cannot read file")
	}
	list := make([]string, 0)
	raw := string(dat)
	result := getURL.FindAllStringSubmatch(raw, -1)
	for _, match := range result {
		list = append(list, match[1])
		raw := match[1]
		logger.Printf("%s,%d,%d,%s", raw, 1, getLength(raw), getTableHandling(raw))
	}
}
