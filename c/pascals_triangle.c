#include <stdlib.h>

int **
generate(int numRows, int *returnSize, int **returnColumnSizes) {
	int **triangle     = malloc(numRows * sizeof(int *));
	*returnColumnSizes = malloc(numRows * sizeof(int));
	*returnSize        = numRows;
	for (int i = 0; i < numRows; i++) {
		int  size               = i + 1;
		int *row                = malloc(size * sizeof(int));
		triangle[i]             = row;
		(*returnColumnSizes)[i] = size;
		for (int j = 0; j < size; j++)
			if (j == 0 || j == size - 1)
				row[j] = 1;
			else
				row[j] = triangle[i - 1][j - 1] + triangle[i - 1][j];
	}
	return triangle;
}
