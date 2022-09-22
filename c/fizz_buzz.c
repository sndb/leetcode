#include <stdio.h>
#include <stdlib.h>

char **
fizzBuzz(int n, int *returnSize) {
	char **a = malloc(sizeof(char *) * n);
	for (int i = 0, x = 1; i < n; i++, x++) {
		if (x % 15 == 0)
			a[i] = "FizzBuzz";
		else if (x % 3 == 0)
			a[i] = "Fizz";
		else if (x % 5 == 0)
			a[i] = "Buzz";
		else {
			a[i] = malloc(sizeof(char) * 6);
			sprintf(a[i], "%d", x);
		}
	}
	*returnSize = n;
	return a;
}
