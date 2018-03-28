package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	counts := make(map[string]int)
	fmt.Println(os.Args)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	}else {
		for _, args := range files{
			f, err := os.Open(args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
 		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	fmt.Println(counts)
}


func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}