package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	duplicateVowel bool = true
	removeVowel    bool = false
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	var vI int
	for s.Scan() {
		word := []byte(s.Text())
		if randBool() {
			vI = -1
			for i, char := range word {
				switch char {
				case 'a', 'e', 'i', 'u', 'o', 'A', 'E', 'I', 'O', 'U':
					if randBool() {
						vI = i
					}
				}
			}
		}

		if vI >= 0 {
			switch randBool() {
			case duplicateVowel:
				word = append(word[:vI+1], word[vI:]...)
			case removeVowel:
				word = append(word[:vI], word[vI+1:]...)
			}
		}
		fmt.Println(string(word))
	}


}
