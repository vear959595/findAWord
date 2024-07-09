package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	word := make(map[string]int)
	var findMe string
	fmt.Fscan(os.Stdin, &findMe)
	scanner := bufio.NewReader(os.Stdin)
	s, _ := scanner.ReadString('\n')
	fmt.Printf("%v", s)
	for _, v := range findMe {
		//if _, ok := word[string(v)]; ok {
		word[string(v)] = word[string(v)] + 1
		//}
	}
	count := make(map[string]int)
	for _, v := range s {
		if _, ok := word[string(v)]; ok {
			count[string(v)]++
		}
	}
	oldValue := 0
	newValue := 0

	if len(count) != len(word) {
		oldValue = 0
	} else {
		for kc, valc := range count {
			for kw, valw := range word {
				if kc == kw {
					if oldValue == 0 {
						oldValue = valc / valw
					}

					newValue = valc / valw

					if oldValue > newValue {
						oldValue = newValue
					}
				}
			}
		}

	}
	fmt.Printf("Слово %v встречается %v раз", findMe, oldValue)
}
