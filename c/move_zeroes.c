#include <stdbool.h>

void
moveZeroes(int *nums, int numsSize) {
	int z = 0;
	for (int i = 0; i < numsSize; i++) {
		while (z < numsSize && nums[z] != 0)
			z++;
		if (nums[i] != 0 && i > z) {
			nums[z] = nums[i];
			nums[i] = 0;
		}
	}
}
