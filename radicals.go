package radicals

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type RadkfileParser struct {
	Radicals map[string][]string
}

func ParseRadkfile(filename string) (r RadkfileParser, err error) {
	radkfile, err := os.Open(filename)
	if err != nil {
		return
	}
	defer radkfile.Close()
	scanner := bufio.NewScanner(radkfile)
	r.Radicals = map[string][]string{}
	cur := ""
	for scanner.Scan() {
		t := scanner.Text()
		switch t[0] {
		case '$':
			s := strings.Split(t, " ")
			cur = s[1] + "_" + s[2]
		default:
			s := strings.Split(scanner.Text(), "")
			r.Radicals[cur] = append(r.Radicals[cur], s...)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
	return
}
