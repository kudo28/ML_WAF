package main

import (
	"io/ioutil"
	"github.com/Sirupsen/logrus"
	"regexp"
	"fmt"
)

var getURL = regexp.MustCompile("GET (http://.*) HTTP")

func getLength(s string) int {
	return len(s)
}
func getToken(s string) []int {
	return nil
}
func main() {
	dat, err := ioutil.ReadFile("normalTrafficTraining.txt")
	if err != nil {
		logrus.Fatal("cannot read file")
	}
	list := make([]string, 0)
	raw := string(dat)
	result := getURL.FindAllStringSubmatch(raw, -1)
	//fmt.Println(result[0])
	for _, match := range result {
		list = append(list, match[1])
		fmt.Println(match[1])
	}
	logrus.Info(len(list))
	//for _, url := range list {
	//	fmt.Println(getLength(url))
	//}
}
