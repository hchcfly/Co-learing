package models

import (
	"fmt"
	"strings"
)


func HandleTagsListData(tags []string) map[string]int {
	var tagsMap = make(map[string]int)
	for _, tag := range tags {
		fmt.Println("--------->",tag)
		tagList := strings.Split(tag,"&")
		fmt.Println("--------->",tagList)
		for _, value := range tagList {
			tagsMap[value]++
		}
	}
	return tagsMap
}