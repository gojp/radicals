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
	if len(got) != numKanji {
		t.Fatalf("ParseKradfile length incorrect: got %d, want %d", len(got), numKanji)
	}
}

var radicalToKanjiTests = []struct {
	radical     string
	strokeCount int
	kanji       []string
}{
	{"入", 2, strings.Split("久込入兩兪叺圦懣杁柩滿疚瞞窩糴裲蹣輛陝魎鳰", "")},
	{"マ", 2, strings.Split("桶擬疑凝柔序痛通樋矛勇湧涌予預踊豫舒俑墅慂懋抒揉矜礙糅蛹蹂踴鞣", "")},
	{"門", 8, strings.Split("闇閏閲開閣澗簡間閑関閤潤閃闘閥聞閉問悶門欄蘭們墹嫺嫻憫捫擱椚櫚瀾燗爛癇繝藺襴躙躪閂閇閊閔閖閘閙閠閨閧閭閼閻閹閾闊濶闃闍闌闕闔闖關闡闥闢", "")},
	{"龠", 17, strings.Split("籥鑰龠", "")},
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
		if got := r[tt.radical].StrokeCount; got != tt.strokeCount {
			t.Errorf("TestRadicalsToKanji(%s): got %d, want %d", tt.radical, got, tt.strokeCount)
		}
	}
}

var kanjiToRadicalTests = []struct {
	kanji    string
	radicals []string
}{
	{"亜", strings.Split("｜ 一 口", " ")},
	{"機", strings.Split("ノ 木 丶 幺 戈", " ")},
	{"熙", strings.Split("杰 已 匚 口", " ")},
}

func TestKanjiToRadicals(t *testing.T) {
	r, err := ParseKradfile("kradfile.utf8")
	if err != nil {
		t.Fatalf("ParseKradfile: %v", err)
	}
	for _, tt := range kanjiToRadicalTests {
		if got := r[tt.kanji].Radicals; !sliceEqual(got, tt.radicals) {
			t.Errorf("TestKanjiToRadicals(%s): got %v, want %v", tt.kanji, got, tt.radicals)
		}
	}
}
