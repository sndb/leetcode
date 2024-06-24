int minKBitFlips(int *nums, int numsSize, int k) {
	int c = 0;
	int x = 0;
	for (int i = 0; i < numsSize; i++) {
		int j = i - k;
		if (j >= 0 && nums[j] > 1)
			x ^= 1;
		if (nums[i] == x) {
			if (i + k > numsSize)
				return -1;
			nums[i] = 2;
			x ^= 1;
			c++;
		}
	}
	return c;
}
