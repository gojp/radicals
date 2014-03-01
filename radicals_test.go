package radicals

import (
	"strings"
	"testing"
)

var (
	numRadicals = 252
	numKanji    = 6355
)

func TestRadkfileParser(t *testing.T) {
	got, err := ParseRadkfile("radkfile.utf8")
	if err != nil {
		t.Fatalf("ParseRadkfile: %v", err)
	}
	if len(got) != numRadicals {
		t.Fatalf("ParseRadkfile length incorrect: got %d, want %d", len(got), numRadicals)
	}
}

func TestKradfileParser(t *testing.T) {
	got, err := ParseKradfile("kradfile.utf8")
	if err != nil {
		t.Fatalf("ParseKradfile: %v", err)
	}
	if len(got.Kanji) != numKanji {
		t.Fatalf("ParseKradfile length incorrect: got %d, want %d", len(got.Kanji), numKanji)
	}
}

var radicalToKanjiTests = []struct {
	radical string
	strokes int
	kanji   []string
}{
	{"入", 2, strings.Split("久込入兩兪叺圦懣杁柩滿疚瞞窩糴裲蹣輛陝魎鳰", "")},
	{"マ", 2, strings.Split("桶擬疑凝柔序痛通樋矛勇湧涌予預踊豫舒俑墅慂懋抒揉矜礙糅蛹蹂踴鞣", "")},
}

func sliceEqual(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, x := range s1 {
		if s2[i] != x {
			return false
		}
	}
	return true
}

func TestRadicalsToKanji(t *testing.T) {
	r, err := ParseRadkfile("radkfile.utf8")
	if err != nil {
		t.Fatalf("ParseRadkfile: %v", err)
	}
	for _, tt := range radicalToKanjiTests {
		if got := r[tt.radical].Kanji; !sliceEqual(got, tt.kanji) {
			t.Errorf("TestRadicalsToKanji(%s): got %v, want %v", tt.radical, got, tt.kanji)
		}
	}
}
