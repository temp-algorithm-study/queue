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

	// 1. 명령어 수 입력받기
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	queue := []int{}

	for i := 0; i < n; i++ {
		scanner.Scan()
		cmd := strings.Fields(scanner.Text())

		switch cmd[0] {
		case "push":
			x, err := strconv.Atoi(cmd[1])
			if err != nil {
				panic(err)
			}
			queue = append(queue, x)
		case "pop":
			if len(queue) == 0 {
				fmt.Fprintln(writer, -1)
				continue
			}

			front := queue[0]
			queue = queue[1:]
			fmt.Fprintln(writer, front)

		case "size":
			fmt.Fprintln(writer, len(queue))
		case "empty":
			if len(queue) == 0 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		case "front":
			if len(queue) == 0 {
				fmt.Fprintln(writer, -1)
				continue
			}
			front := queue[0]
			fmt.Fprintln(writer, front)

		case "back":
			if len(queue) == 0 {
				fmt.Fprintln(writer, -1)
				continue
			}
			back := queue[len(queue)-1]
			fmt.Fprintln(writer, back)

		}
	}
}
