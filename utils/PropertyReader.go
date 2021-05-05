package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var propertyMap = make(map[string]string)

func ReadProperties() {
	properties, err := os.Open("E:\\bots\\iroh\\resources\\irohbot.properties")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(properties)

	for scanner.Scan() {
		line := scanner.Text()

		//ignore comments
		if line[0] != '#' {
			keyValuePair := strings.Split(line, "=")

			propertyMap[keyValuePair[0]] = keyValuePair[1]
		}
	}
}

func GetProperty(name string) string {
	return propertyMap[name]
}
