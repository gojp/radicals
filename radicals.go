package radicals

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Radical struct {
	StrokeCount int
	Kanji       []string
}

type RadkfileParser map[string]Radical

type Kanji struct {
	Radicals []string
}

type KradfileParser map[string]Kanji

func ParseRadkfile(filename string) (RadkfileParser, error) {
	r := RadkfileParser{}
	radkfile, err := os.Open(filename)
	if err != nil {
		return r, err
	}
	defer radkfile.Close()
	scanner := bufio.NewScanner(radkfile)
	var rad Radical
	var cur = ""
	var strokes = 0
	for scanner.Scan() {
		t := scanner.Text()
		switch t[0] {
		case '#':
			continue
		case '$':
			s := strings.Split(t, " ")
			cur = s[1]
			strokes, err = strconv.Atoi(s[2])
			if err != nil {
				return r, err
			}
			rad.StrokeCount = strokes
			r[cur] = rad
		default:
			s := strings.Split(t, "")
			rad := r[cur]
			k := r[cur].Kanji
			k = append(r[cur].Kanji, s...)
			rad.Kanji = k
			r[cur] = rad
		}
	}
	if err = scanner.Err(); err != nil {
		return r, err
	}
	return r, err
}

func ParseKradfile(filename string) (KradfileParser, error) {
	k := KradfileParser{}
	kradfile, err := os.Open(filename)
	if err != nil {
		return k, err
	}
	defer kradfile.Close()
	scanner := bufio.NewScanner(kradfile)
	for scanner.Scan() {
		t := scanner.Text()
		switch t[0] {
		case '#':
			continue
		default:
			s := strings.Split(t, " : ")
			kanji := Kanji{}
			kanji.Radicals = s[1:]
			k[s[0]] = kanji
		}
	}
	if err = scanner.Err(); err != nil {
		return k, err
	}
	return k, err
}
