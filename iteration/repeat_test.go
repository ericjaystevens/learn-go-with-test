package iteration

import(
	"testing"
)

func TestRepeat(t *testing.T){
	got := Repeat("sh",3)
	want := "shshsh"

	if got != want {
		t.Errorf("got: %q, want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++ {
		Repeat("sh",3)	
	}
}