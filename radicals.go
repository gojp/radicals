package radicals

import (
	"bufio"
	"os"
	"strings"
)

type RadkfileParser struct {
	Radicals map[string][]string
}

type KradfileParser struct {
	Kanji map[string][]string
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
		case '#':
			continue
		case '$':
			s := strings.Split(t, " ")
			cur = s[1] + "_" + s[2]
		default:
			s := strings.Split(t, "")
			r.Radicals[cur] = append(r.Radicals[cur], s...)
		}
	}
	if err = scanner.Err(); err != nil {
		return
	}
	return
}

func ParseKradfile(filename string) (k KradfileParser, err error) {
	kradfile, err := os.Open(filename)
	if err != nil {
		return
	}
	defer kradfile.Close()
	scanner := bufio.NewScanner(kradfile)
	k.Kanji = map[string][]string{}
	for scanner.Scan() {
		t := scanner.Text()
		switch t[0] {
		case '#':
			continue
		default:
			s := strings.Split(t, " : ")
			k.Kanji[s[0]] = s[1:]
		}
	}
	if err = scanner.Err(); err != nil {
		return
	}
	return
}
