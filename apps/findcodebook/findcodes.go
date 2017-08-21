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

var splitchar = map[rune]bool{'/':true, '&':true, '%':true, '.':true, '=':true }

func main() {
	reader := bufio.NewReader(os.Stdin)
	codes := make([]*code, 0)
	for {
		bytes, err := reader.ReadBytes('\n')
		url := strings.Trim(string(bytes), "\n\"")
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if err == io.EOF {
			PrintResults(codes)
			os.Exit(0)
		}
		strs := strings.FieldsFunc(url,
			func(c rune) bool {
				_, ok := splitchar[c]
				return ok
			})
		for _, s := range strs {
			if len(s) > 40 {
				continue
			}
			exists := false
			for _, c := range codes {
				if c.name == s {
					exists = true
					c.occurrences++
				}
			}
			if !exists {
				codes = append(codes, &code{s, 1})
			}
		}
	}
}
type code struct {
	name        string
	occurrences int
}

func (c *code) score() int {
	return (len(c.name)-1)*c.occurrences
}

func PrintResults(codes []*code) {
	sort.Slice(codes, func(i, j int) bool {
		return codes[i].score() >= codes[j].score()
	})
	for _, c := range codes {
		fmt.Printf("string: %s, occurrences: %v, score: %v\n", c.name, c.occurrences, c.score())
	}
}