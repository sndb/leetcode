#include <limits.h>

int
reverse(int x) {
	int r = 0;
	while (x != 0) {
		int v = x % 10;
		if (r < 0 && r < INT_MIN / 10)
			return 0;
		if (r > 0 && r > INT_MAX / 10)
			return 0;
		r *= 10;
		if (r < 0 && r < INT_MIN - v)
			return 0;
		if (r > 0 && r > INT_MAX - v)
			return 0;
		r += v;
		x /= 10;
	}
	return r;
}
