#include <stdbool.h>
#include <stdlib.h>

int *
plusOne(int *digits, int digitsSize, int *returnSize) {
	int *r = malloc((1 + digitsSize) * sizeof(int));
	for (int i = 0; i < digitsSize; i++)
		r[1 + i] = digits[i];
	r[0] = 0;
	r[digitsSize] += 1;
	for (int i = digitsSize; i >= 1; i--) {
		if (r[i] > 9) {
			r[i - 1] += 1;
			r[i] = 0;
		}
	}
	if (r[0] == 0) {
		*returnSize = digitsSize;
		return r + 1;
	}
	*returnSize = digitsSize + 1;
	return r;
}
