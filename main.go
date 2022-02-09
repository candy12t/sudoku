package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/candy12t/sudoku/borad"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Error: require sudoku file path")
		os.Exit(1)
	}

	filePath := os.Args[1]
	data, err := readFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	b, err := initBorad(data)
	if err != nil {
		log.Fatal(err)
	}

	b.Backtrack()
	fmt.Print(b)
}

func initBorad(line string) (borad.Borad, error) {
	var b borad.Borad
	br := bytes.NewBufferString(line)
	s := bufio.NewScanner(br)
	s.Split(bufio.ScanRunes)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !s.Scan() {
				break
			}
			token := s.Text()
			if token == "." {
				b[i][j] = 0
				continue
			}
			n, err := strconv.Atoi(token)
			if err != nil {
				return borad.Borad{}, err
			}
			b[i][j] = n
		}
	}
	return b, nil
}

func readFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	line := strings.Replace(string(data), "\n", "", 9)
	return line, nil
}
