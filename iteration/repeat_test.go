package iteration

import "testing"

import "fmt"

func TestRepeat(t *testing.T) {
	got := Repeat("z", 7)
	want := "zzzzzzz"
	
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("z", 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat("z", 11)
	fmt.Println(repeated)
	// Output: zzzzzzzzzzz
}