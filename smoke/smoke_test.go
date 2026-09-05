package smoke

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "named", in: "wand", want: "hello, wand"},
		{name: "empty falls back to the world", in: "", want: "hello, world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.in); got != tt.want {
				t.Errorf("Greet(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
