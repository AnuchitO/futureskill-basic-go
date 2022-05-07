package calc

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 1)
	want := 2
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestAddwithSubTestStyle(t *testing.T) {
	t.Run("1+1=2", func(t *testing.T) {
		got := Add(1, 1)
		want := 2
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestAddWithTableStyle(t *testing.T) {
	var tests = []struct {
		x    int
		y    int
		want int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
	}
	for _, tt := range tests {
		got := Add(tt.x, tt.y)

		if got != tt.want {
			t.Errorf("Add(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
		}
	}
}
