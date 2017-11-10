package diceware

import (
	"log"
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	for i := 0; i < 10000; i++ {
		if _, err := Generate(16); err != nil {
			t.Fatal(err)
		}
	}
}

func TestRollDie(t *testing.T) {
	t.Parallel()

	for i := 0; i < 10000; i++ {
		r, err := RollDie()
		if err != nil {
			t.Fatal(err)
		}

		if r < 1 || r > 6 {
			t.Fatalf("expected result to be in range (%d)", r)
		}
	}
}

func TestRollWord(t *testing.T) {
	t.Parallel()

	for i := 0; i < 10000; i++ {
		r, err := RollWord(5)
		if err != nil {
			t.Fatal(err)
		}

		if r < 11111 || r > 66666 {
			t.Fatalf("expected result to be in range (%d)", r)
		}
	}
}

func ExampleGenerate() {
	words, err := Generate(6)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", words)
}
