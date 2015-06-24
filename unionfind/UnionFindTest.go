package main

import (
	"algorithms/unionfind/uf"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	path := "input.txt"
	file, err := os.Open(path)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var uf uf.UF
	if scanner.Scan() {
		uf.N, err = strconv.ParseInt(scanner.Text(), 0, 64)
		check(err)
	}

	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), " ")
		p, err := strconv.ParseInt(pairs[0], 0, 64)
		check(err)

		q, err := strconv.ParseInt(pairs[1], 0, 64)
		check(err)

		if !uf.Connected(p, q) {
			uf.Union(p, q)
			fmt.Println(p, q)
		}
	}
}
