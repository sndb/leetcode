#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int toNumber(char c) {
	if (c == '.')
		return -1;
	return c - '0';
}

bool validRow(char **board, int r) {
	bool seen[10] = {0};
	for (int c = 0; c < 9; c++) {
		int n = toNumber(board[r][c]);
		if (n == -1)
			continue;
		if (seen[n])
			return false;
		seen[n] = true;
	}
	return true;
}

bool validColumn(char **board, int c) {
	bool seen[10] = {0};
	for (int r = 0; r < 9; r++) {
		int n = toNumber(board[r][c]);
		if (n == -1)
			continue;
		if (seen[n])
			return false;
		seen[n] = true;
	}
	return true;
}

bool validSquare(char **board, int r, int c) {
	bool seen[10] = {0};
	int or = r * 3, oc = c * 3;
	for (int r = 0; r < 3; r++) {
		for (int c = 0; c < 3; c++) {
			int n = toNumber(board[r + or][c + oc]);
			if (n == -1)
				continue;
			if (seen[n])
				return false;
			seen[n] = true;
		}
	}
	return true;
}

bool isValidSudoku(char **board, int boardSize, int *boardColSize) {
	for (int r = 0; r < 9; r++)
		if (!validRow(board, r))
			return false;
	for (int c = 0; c < 9; c++)
		if (!validColumn(board, c))
			return false;
	for (int i = 0; i < 9; i++)
		if (!validSquare(board, i / 3, i % 3))
			return false;
	return true;
}

int main(void) {
	char *board[9] = {
		"53..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	};
	printf("%d\n", isValidSudoku(board, 0, NULL));
	return 0;
}
