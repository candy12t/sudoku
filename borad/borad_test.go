package borad

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	tests := []struct {
		name  string
		borad Borad
		want  bool
	}{
		{
			name: "verified zero",
			borad: Borad{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: true,
		},
		{
			name: "verified",
			borad: Borad{
				{6, 1, 2, 3, 7, 8, 4, 5, 9},
				{9, 4, 5, 6, 1, 2, 7, 8, 3},
				{3, 7, 8, 9, 4, 5, 6, 1, 2},
				{7, 8, 6, 1, 2, 9, 3, 4, 5},
				{1, 2, 3, 4, 5, 6, 9, 7, 8},
				{5, 9, 4, 7, 8, 3, 1, 2, 6},
				{4, 5, 9, 8, 6, 1, 2, 3, 7},
				{8, 3, 7, 2, 9, 4, 5, 6, 1},
				{2, 6, 1, 5, 3, 7, 8, 9, 4},
			},
			want: true,
		},
		{
			name: "not verified row",
			borad: Borad{
				{1, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: false,
		},
		{
			name: "not verified column",
			borad: Borad{
				{1, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: false,
		},
		{
			name: "not verified box",
			borad: Borad{
				{1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.borad.Verify()
			if got != tt.want {
				t.Errorf("got is %v, wannt is %v\n", got, tt.want)
			}
		})
	}
}

func TestBackTrack(t *testing.T) {
	tests := []struct {
		borad Borad
		want  bool
	}{
		{
			borad: Borad{
				{0, 1, 2, 3, 0, 0, 0, 0, 0},
				{0, 4, 5, 6, 0, 0, 0, 8, 0},
				{0, 7, 8, 9, 0, 0, 0, 0, 0},
				{7, 0, 0, 0, 2, 0, 0, 0, 5},
				{0, 0, 0, 4, 5, 6, 0, 0, 0},
				{5, 0, 0, 0, 8, 0, 0, 0, 6},
				{0, 0, 0, 0, 0, 1, 2, 3, 0},
				{0, 3, 0, 0, 0, 4, 5, 6, 0},
				{0, 0, 0, 0, 0, 7, 8, 9, 0},
			},
			want: true,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("pattern #%d", i), func(t *testing.T) {
			got := tt.borad.Backtrack()
			if got != tt.want {
				t.Errorf("got is %v, wannt is %v\n", got, tt.want)
			}
		})
	}
}

func TestDuplicate(t *testing.T) {
	tests := []struct {
		name string
		cnt  [10]int
		want bool
	}{
		{
			name: "not duplicated",
			cnt:  [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: false,
		},
		{
			name: "duplicated",
			cnt:  [10]int{0, 2, 0, 0, 0, 0, 0, 0, 0, 0},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := duplicate(tt.cnt)
			if got != tt.want {
				t.Errorf("got is %v, wannt is %v\n", got, tt.want)
			}
		})
	}
}
