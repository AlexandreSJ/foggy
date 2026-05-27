package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	arr := []string{}

	for amount, err := strconv.Atoi(strings.TrimSpace(scanner.Text())); amount > 0; amount-- {
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		scanner.Scan()

		line := strings.TrimSpace(scanner.Text())
		splittedLine := strings.Split(line, " ")
		if len(splittedLine) != 4 {
			arr = append(arr, "NO")
			continue
		}
		for i := 1; i < len(splittedLine); i++ { // the index got to be 1 to compare with the previous element
			current, err := strconv.Atoi(splittedLine[i])    // otherwise we would get an index out of bounds error
			previous, err := strconv.Atoi(splittedLine[i-1]) // and the len(splittedLine) must be 4 on the previous if
			if err != nil || current <= 0 {
				arr = append(arr, "NO")
				break
			}
			if current != previous {
				arr = append(arr, "NO")
				break
			}
			if current == previous && i == len(splittedLine)-1 {
				arr = append(arr, "YES")
				break
			}

		}
	}

	for i := range arr {
		fmt.Println(arr[i])
	}

}
