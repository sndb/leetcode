#include <stdbool.h>

bool isBadVersion(int version);

int
firstBadVersion(int n) {
	int min = 1;
	while (min < n) {
		int mid = n / 2 + min / 2;
		if (isBadVersion(mid))
			n = mid;
		else
			min = mid + 1;
	}
	return min;
}
