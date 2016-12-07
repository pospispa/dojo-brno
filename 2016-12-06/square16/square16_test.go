package square16

import "testing"

func TestSquare16(t *testing.T) {
	tests := []struct {
		in   Solution
		want bool
	}{
		{
			in:   Solution{},
			want: false,
		},
		{
			in: Solution{
				[]Point{
					{0, 0}, {0, 4},
					{1, 4}, {1, 1},
					{2, 1}, {2, 2},
					{3, 2}, {3, 0},
				},
				[]Point{
					{1, 1}, {1, 4},
					{4, 4}, {4, 0},
					{3, 0}, {3, 2},
					{2, 2}, {2, 1},
				},
			},
			want: true,
		},
		{
			in: Solution{
				[]Point{
					{0, 0}, {0, 4},
					{2, 4}, {2, 0},
				},
				[]Point{
					{2, 0}, {2, 4},
					{4, 4}, {4, 0},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		got := IsSolution(tt.in)
		if got != tt.want {
			t.Errorf("IsSolution(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestPerimeter(t *testing.T) {
	p := []Point{
		{0, 0}, {0, 4},
		{1, 4}, {1, 1},
		{2, 1}, {2, 2},
		{3, 2}, {3, 0},
	}
	got := perimeter(p)
	want := 16
	if got != want {
		t.Errorf("perimeter(%v) = %v, want %v", p, got, want)
	}
}

func TestGenerate(t *testing.T) {
	allCombinations := powInt(2, 16)
	if got := len(Generate()); got != allCombinations {
		t.Errorf("Generate() generated (%v) wanted (%v)", got, allCombinations)
	}
}

func TestTwoObjectsOnly(t *testing.T) {
	tests := []struct {
		s    Square4x4
		want bool
	}{
		{
			s: Square4x4{
				{White, White, White, White},
				{White, White, White, White},
				{White, White, White, White},
				{White, White, White, White},
			},
			want: false,
		},
		{
			s: Square4x4{
				{Black, White, White, White},
				{White, Black, White, White},
				{White, White, White, White},
				{White, White, White, White},
			},
			want: false,
		},
		{
			s: Square4x4{
				{Black, Black, White, White},
				{Black, Black, White, White},
				{White, White, White, White},
				{White, White, White, White},
			},
			want: true,
		},
		{
			s: Square4x4{
				{Black, Black, White, White},
				{Black, White, White, White},
				{Black, Black, Black, White},
				{White, White, White, White},
			},
			want: true,
		},
		{
			s: Square4x4{
				{Black, Black, White, White},
				{Black, White, White, White},
				{Black, Black, Black, Black},
				{White, White, White, White},
			},
			want: false,
		},
		{
			s: Square4x4{
				{Black, Black, White, White},
				{Black, White, White, White},
				{Black, Black, Black, Black},
				{Black, Black, White, White},
			},
			want: false,
		},
		{
			s: Square4x4{
				{Black, White, White, White},
				{Black, White, White, White},
				{Black, White, Black, White},
				{Black, Black, Black, White},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		if got := TwoObjectsOnly(tt.s); got != tt.want {
			t.Errorf("TwoObjectsOnly(%v) got (%v) want (%v)", tt.s, got, tt.want)
		}
	}
}

func TestGetAllSolutions(t *testing.T) {
	GetAllSolutions()
}
