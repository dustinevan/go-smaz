package main

import (
	"os"
	"bufio"
	"log"
	"io"
	"strings"
	"fmt"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	codes := make(map[string]int)
	for {
		bytes, err := reader.ReadBytes('\n')
		s := strings.Trim(string(bytes), "\n\"")
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if err == io.EOF {
			PrintResults(codes)
			os.Exit(0)
		}

		for i := 0; i < len(s)-5; i++ {
			this := s[i:(i+5)]
			_, ok := codes[this]
			if ok {
				codes[this]++
				continue
			}
			codes[this] = 1
		}
	}
}
type code struct {
	name        string
	occurrences int
}


func PrintResults(codes map[string]int) {
	results := make([]code, 0)

	for k, v := range codes {
		results = append(results, code{k, v})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].occurrences >= results[j].occurrences
	})
	for _, c := range results {
		fmt.Printf("string: %s, occurrences: %v\n", c.name, c.occurrences)
	}
}