#include <stdio.h>
#include <stdlib.h>

void
rotateBuffer(int **matrix, int n) {
	int **buf = malloc(n * sizeof(int *));
	for (int i = 0; i < n; i++)
		buf[i] = malloc(n * sizeof(int));
	for (int i = 0; i < n; i++)
		for (int j = 0; j < n; j++)
			buf[i][j] = matrix[n - 1 - j][i];
	for (int i = 0; i < n; i++)
		for (int j = 0; j < n; j++)
			matrix[i][j] = buf[i][j];
}

void
rotateShift(int **matrix, int n) {
	/* if x = matrix[i][j] in [-1000, 1000], increment x by 1000 for x to be in [0, 2000] and
	 * shift left by 11 (ceil(log2(x)), (1 << 11 = 2048) > 2000) to store the updated value */
	const int shift = 11, range = 1000;
	for (int i = 0; i < n; i++)
		for (int j = 0; j < n; j++)
			matrix[i][j] = (matrix[i][j] + range) << shift;
	for (int i = 0; i < n; i++)
		for (int j = 0; j < n; j++)
			matrix[i][j] += matrix[n - 1 - j][i] >> shift;
	for (int i = 0; i < n; i++)
		for (int j = 0; j < n; j++)
			matrix[i][j] = (matrix[i][j] & ((1 << shift) - 1)) - range;
}

void
printMatrix(int **matrix, int n) {
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < n; j++)
			printf("%5d ", matrix[i][j]);
		printf("\n");
	}
}

int **
newMatrix(int n) {
	int **matrix = malloc(n * sizeof(int *));
	for (int i = 0; i < n; i++) {
		matrix[i] = malloc(n * sizeof(int));
		for (int j = 0; j < n; j++)
			matrix[i][j] = i * n + j + 1;
	}
	return matrix;
}

int
main(void) {
	void (*rotate)(int **, int) = rotateShift;

	const int n = 3, range = 1000;

	int **matrix = newMatrix(n);
	matrix[0][0] = -range;
	matrix[0][1] = range;

	printMatrix(matrix, n);
	printf("\n");

	rotate(matrix, n);
	printMatrix(matrix, n);
}
