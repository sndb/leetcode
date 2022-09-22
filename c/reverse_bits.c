#include <stdint.h>
#include <stdio.h>

uint32_t
reverseBits(uint32_t n) {
	for (uint8_t i = 0; i < (32 / 2); i++) {
		uint32_t ir = (1L << i);
		uint32_t il = (1L << (31 - i));
		uint32_t r  = n & ir;
		uint32_t l  = n & il;
		if ((!l && r) || (l && !r))
			n ^= ir | il;
	}
	return n;
}
