package borad

import (
	"bytes"
	"strconv"
)

type Borad [9][9]int

func (b Borad) String() string {
	var buf bytes.Buffer
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			buf.WriteString("+---+---+---+\n")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				buf.WriteString("|")
			}
			buf.WriteString(strconv.Itoa(b[i][j]))
		}
		buf.WriteString("|\n")
	}
	buf.WriteString("+---+---+---+\n")
	return buf.String()
}

func (b *Borad) Verify() bool {
	if !b.verifyRow() {
		return false
	}

	if !b.verifyColumn() {
		return false
	}

	if !b.verifyBox() {
		return false
	}
	return true
}

func (b *Borad) Backtrack() bool {
	if b.solved() {
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				for c := 1; c <= 9; c++ {
					b[i][j] = c
					if b.Verify() {
						if b.Backtrack() {
							return true
						}
					}
					b[i][j] = 0
				}
				return false
			}
		}
	}
	return false
}

func (b *Borad) solved() bool {
	if !b.Verify() {
		return false
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func (b *Borad) verifyRow() bool {
	for i := 0; i < 9; i++ {
		cnt := [10]int{0}
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				continue
			}
			cnt[b[i][j]]++
		}
		if duplicate(cnt) {
			return false
		}
	}
	return true
}

func (b *Borad) verifyColumn() bool {
	for i := 0; i < 9; i++ {
		cnt := [10]int{0}
		for j := 0; j < 9; j++ {
			if b[j][i] == 0 {
				continue
			}
			cnt[b[j][i]]++
		}
		if duplicate(cnt) {
			return false
		}
	}
	return true
}

func (b *Borad) verifyBox() bool {
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			cnt := [10]int{0}
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if b[i+k][j+l] == 0 {
						continue
					}
					cnt[b[i+k][j+l]]++
				}
			}
			if duplicate(cnt) {
				return false
			}
		}
	}
	return true
}

func duplicate(cnt [10]int) bool {
	for _, v := range cnt {
		if v >= 2 {
			return true
		}
	}
	return false
}
