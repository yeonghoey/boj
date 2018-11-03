package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	maxStat = 1000
)

var reader = bufio.NewReader(os.Stdin)

func nextInts() []int {
	s, _ := reader.ReadString('\n')
	fs := strings.Fields(s)
	ns := make([]int, len(fs))
	for i, f := range fs {
		ns[i], _ = strconv.Atoi(f)
	}
	return ns
}

func nextInt() int {
	return nextInts()[0]
}

func nextQuest() Quest {
	ns := nextInts()
	return Quest{Stat{ns[0], ns[1]}, ns[2]}
}

type Quest struct {
	Stat
	PNT int
}

type Stat struct {
	STR int
	INT int
}

func main() {
	N := nextInt()

	quests := make([]Quest, N)
	for i := 0; i < N; i++ {
		quests[i] = nextQuest()
	}

	table := make([][]int, maxStat+1)
	for i := range table {
		table[i] = make([]int, maxStat+1)
		for j := range table[i] {
			table[i][j] = -1
		}
	}

	table[0][1] = 1
	table[1][0] = 1

	best := 0
	queue := []Stat{{1, 1}}
	for len(queue) > 0 {
		stat := queue[0]
		queue = queue[1:]

		STR, INT := stat.STR, stat.INT

		// Already computed
		if table[STR][INT] != -1 {
			continue
		}

		points, nQuests := 2-(STR+INT), 0
		for _, q := range quests {
			if STR >= q.STR || INT >= q.INT {
				points += q.PNT
				nQuests++
			}
		}

		table[STR][INT] = points

		if points > 0 {
			queue = append(queue, Stat{STR + 1, INT})
			queue = append(queue, Stat{STR, INT + 1})
		}

		if nQuests > best {
			best = nQuests
		}

		if best == N {
			break
		}
	}

	fmt.Println(best)
}
