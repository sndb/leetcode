package main

func rotateTheBox(box [][]byte) [][]byte {
	rows := len(box)
	cols := len(box[0])
	result := make([][]byte, cols)
	for c := 0; c < cols; c++ {
		result[c] = make([]byte, rows)
	}
	for r := 0; r < rows; r++ {
		rr := rows - r - 1
		cc := cols - 1
		for c := cols - 1; c >= 0; c-- {
			result[c][rr] = '.'
			switch box[r][c] {
			case '*':
				result[c][rr] = '*'
				cc = c - 1
			case '#':
				result[cc][rr] = '#'
				cc--
			}
		}
	}
	return result
}
