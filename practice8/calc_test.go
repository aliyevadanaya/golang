package practice8

import "testing"

//func TestSum(t *testing.T) {
//	got := Sum(2, 0)
//	want := 8
//	if got != want {
//		t.Errorf("Got %d, want %d", got, want)
//	}
//}

func TestSumTableDriven(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"both positive", 2, 4, 6},
		{"positive + zero", 1, 0, 1},
		{"both negative", -4, -3, -7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Sum(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	got, err := Divide(12, 3)
	want := 4
	if err != nil {
		t.Errorf("Unexpexted error: %v", err)
	}

	if got != want {
		t.Errorf("Divide(12, 3) = %d, want %d", got, want)
	}

	_, err = Divide(2, 0)
	if err == nil {
		t.Error("Expected error, but nil???")
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"both positive", 5, 3, 2},
		{"positive and zero", 5, 0, 5},
		{"negative and positive", -2, 3, -5},
		{"both negative", -5, -3, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
