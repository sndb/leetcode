int minOperations(int *nums, int numsSize, int k) {
	int c = 0;
	for (int i = 0; i < numsSize; i++)
		if (nums[i] < k)
			c++;
	return c;
}
