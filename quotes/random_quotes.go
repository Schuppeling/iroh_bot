package quotes

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

var quoteMap = make(map[int]string)

func init() {
	readQuotes()
}

func readQuotes() {
	quotes, err := os.Open("E:\\bots\\iroh_bot\\resources\\quotes.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(quotes)
	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()

		quoteMap[lineNum] = line
		lineNum++
	}
}

func GetRandomQuote() string {
	randSource := rand.NewSource(time.Now().UnixNano())
	randomSeed := rand.New(randSource)

	return quoteMap[randomSeed.Intn(50)]
}
