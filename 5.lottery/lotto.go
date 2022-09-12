package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(os.Stderr, "Invalid Argumemts") // Stderr: Error 메세지 출력할 때 사용하는 표준입출력 방식
		return
	}

	filename := os.Args[1]
	count, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can not convert count to integer! count: ", count)
	}

	books, err := readBooks(filename)
	if err != nil {
		fmt.Fprint(os.Stderr, "Can not read Book file!", err)
		return
	}

	rand.Seed(time.Now().UnixNano())

	sellected := make([]string, count)
	for i := 0; i < count; i++ {
		n := rand.Intn(len(books))
		sellected[i] = books[n]
		books = append(books[:n], books[n+1:]...)
	}

	for _, book := range sellected {
		fmt.Println(book)
	}
}

func readBooks(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var books []string
	for scanner.Scan() {
		books = append(books, scanner.Text())
	}
	return books, nil
}
