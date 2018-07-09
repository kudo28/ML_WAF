package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"io/ioutil"
	"os"
	"io"
)

var Tokens = []string{"SELECT", "INSERT", "UPDATE", "DELETE", "CREATE", "DROP", "ALTER", "RENAME", "WHERE", "FROM", "UNION", "NOT", "AND", "OR", "XOR","EXEC",
	"!", "&&", "||", "--", "#", "<", ">", "<=>", ">=", "<=", "==", "=", "!=", "<<", ">>", "<>", "%", "*", "?", "|", "&", "-", "+"}

func caseInsenstiveContains(a, b string) bool {
	return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}

var getURL = regexp.MustCompile("GET (http://.*) HTTP")

func getLength(s string) int {
	return len(s)
}

func containAndConcat(raw, word, s string) string {
	bol := caseInsenstiveContains(raw, word)
	if bol == true {
		return fmt.Sprintf("%s,1", s)
	}
	return fmt.Sprintf("%s,0", s)

}

func getTokenEasy(raw string) string {

	s := ""

	bol := caseInsenstiveContains(raw, Tokens[0])
	if bol == true {
		s = "1"
	} else {
		s = "0"
	}

	for i, reserveword := range Tokens {
		if i == 0 {
			continue
		}
		s = containAndConcat(raw, reserveword, s)
	}
	s = getCommentTokens(raw, s)
	return s
}

func getCommentTokens(raw, s string) string {
	var re = regexp.MustCompile(`(?m)/\*.*\*/`)
	check := re.MatchString(raw)
	if check == true {
		return fmt.Sprintf("%s,1", s)
	}
	return fmt.Sprintf("%s,0", s)
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

func processNormal(filename string) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("normal cannot read file")
	}
	list := make([]string, 0)
	raw := string(dat)
	result := getURL.FindAllStringSubmatch(raw, -1)
	for _, match := range result {
		list = append(list, match[1])
		raw := match[1]
		logger.Printf("%s,%d,%d,%s", raw, 0, getLength(raw), getTokenEasy(raw))
	}
	fmt.Println(len(Tokens))
}

func processMalicious(filename string) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("malicious cannot read file")
	}
	list := make([]string, 0)
	raw := string(dat)
	result := getURL.FindAllStringSubmatch(raw, -1)
	for _, match := range result {
		list = append(list, match[1])
		raw := match[1]
		logger.Printf("%s,%d,%d,%s", raw, 1, getLength(raw), getTokenEasy(raw))
	}
}
