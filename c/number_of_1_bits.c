#include <stdint.h>

int
hammingWeight(uint32_t n) {
	int i = 0;
	while (n > 0) {
		i += 1 & n;
		n >>= 1;
	}
	return i;
}
