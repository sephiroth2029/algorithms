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
	var ufAlg uf.UF
	if scanner.Scan() {
		var n int
		n, err = strconv.Atoi(scanner.Text())
		check(err)

		ufAlg = *uf.NewUF(n)
	}

	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), " ")
		p, err := strconv.Atoi(pairs[0])
		check(err)

		q, err := strconv.Atoi(pairs[1])
		check(err)

		if !ufAlg.Connected(p, q) {
			ufAlg.Union(p, q)
			fmt.Println(p, q)
		}
	}
}
