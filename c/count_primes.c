#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

int
countPrimes(int n) {
	if (n < 3)
		return 0;

	bool *primes = calloc(sizeof(bool), n);
	int   r      = 0;
	for (int i = 2; i < n; i++) {
		if (!primes[i]) {
			for (int p = i, j = 1; p * j < n; j++)
				primes[p * j] = true;
			r++;
		}
	}
	free(primes);
	return r;
}

int
main(void) {
	printf("%d\n", countPrimes(10));
	return 0;
}
