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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	input := strings.Fields(scanner.Text())
	n, err := strconv.Atoi(input[0])
	if err != nil {
		panic(err)
	}

	k, err := strconv.Atoi(input[1])
	if err != nil {
		panic(err)
	}

	queue := []int{}
	answer := []int{}
	for i := 1; i <= n; i++ {
		queue = append(queue, i)
	}

	for x := 0; x < n; x++ {
		for y := 0; y < k; y++ {
			value := queue[0]
			queue = queue[1:]

			/// 마지막이라면
			if y == k-1 {
				answer = append(answer, value)
				break
			}

			queue = append(queue, value)
		}
	}

	fmt.Fprint(writer, "<")
	for i, v := range answer {
		if i > 0 {
			fmt.Fprint(writer, ", ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprint(writer, ">")

}
