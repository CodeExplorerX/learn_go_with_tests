package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat(5, "a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat(5, "a")
	}
}

func ExampleRepeat() {
	repeated := Repeat(5, "a")
	fmt.Println(repeated)
	// Output: aaaaa
}
