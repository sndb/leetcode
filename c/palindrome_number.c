#include <stdbool.h>

bool
isPalindrome(int x) {
	int  k = x;
	long n = 0;
	while (k > 0) {
		n = n * 10 + k % 10;
		k /= 10;
	}
	return n == x;
}
