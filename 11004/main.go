package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	nextInt := func() int {
		_ = scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		return n
	}

	N := nextInt()
	K := nextInt() - 1

	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}

	a := A[:N]
	sort.Ints(a)
	fmt.Println(a[K])
}
