package radicals

import "testing"

var (
	numRadicals    = 252
	numKanji       = 6355
	radkfileParser = RadkfileParser{}
	kradfileParser = KradfileParser{}
)

func TestRadkfileParser(t *testing.T) {
	got, err := ParseRadkfile("radkfile.utf8")
	if err != nil {
		t.Fatalf("ParseRadkfile: %v", err)
	}
	if len(got.Radicals) != numRadicals {
		t.Fatalf("ParseRadkfile length incorrect: got %d, want %d", len(got.Radicals), numRadicals)
	}
	radkfileParser = got
}

func TestKradfileParser(t *testing.T) {
	got, err := ParseKradfile("kradfile.utf8")
	if err != nil {
		t.Fatalf("ParseKradfile: %v", err)
	}
	if len(got.Kanji) != numKanji {
		t.Fatalf("ParseKradfile length incorrect: got %d, want %d", len(got.Kanji), numKanji)
	}
	kradfileParser = got
}
