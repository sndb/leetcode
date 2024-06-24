int minOperations(int *nums, int numsSize) {
	int c = 0;
	for (int i = 0; i < numsSize; i++) {
		if (nums[i] == 0) {
			if (i + 3 > numsSize)
				return -1;
			for (int j = i; j < i + 3; j++)
				nums[j] ^= 1;
			c++;
		}
	}
	return c;
}
