package iteration

import (
	"fmt"
	"strings"
	"testing"
)

const repeatCount = 50

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", repeatCount)
	expected := strings.Repeat("a", repeatCount)

	if repeated != expected {
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", repeatCount)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
