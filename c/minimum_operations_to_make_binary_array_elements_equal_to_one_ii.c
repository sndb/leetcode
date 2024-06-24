int minOperations(int *nums, int numsSize) {
	int c = 0;
	int x = 0;
	for (int i = 0; i < numsSize; i++) {
		if (nums[i] == x) {
			x ^= 1;
			c++;
		}
	}
	return c;
}
