#include <stdbool.h>

bool
isPowerOfThree(int n) {
	while (n > 0)
		if (n == 1)
			return true;
		else if (n % 3 != 0)
			return false;
		else
			n /= 3;
	return false;
}
