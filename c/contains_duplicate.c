#include <stdbool.h>
#include <stdlib.h>

int
compar(const void *p1, const void *p2) {
	int a = *(int *)(p1), b = *(int *)(p2);
	if (a < b)
		return -1;
	if (a > b)
		return 1;
	return 0;
}

bool
containsDuplicate(int *nums, int numsSize) {
	qsort(nums, numsSize, sizeof(int), compar);
	for (int i = 0; i < numsSize - 1; i++)
		if (nums[i] == nums[i + 1])
			return true;
	return false;
}
