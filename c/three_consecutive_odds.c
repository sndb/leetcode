#include <stdbool.h>

bool threeConsecutiveOdds(int *arr, int arrSize) {
	int odds = 0;
	for (int i = 0; i < arrSize; i++) {
		if (arr[i] % 2 == 0)
			odds = 0;
		else if (++odds == 3)
			return true;
	}
	return false;
}
