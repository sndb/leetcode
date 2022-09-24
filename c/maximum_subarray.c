int
maxSubArray(int *nums, int numsSize) {
	int best = nums[0], curr = nums[0];
	for (int i = 1; i < numsSize; i++) {
		curr = nums[i] < nums[i] + curr ? nums[i] + curr : nums[i];
		best = best < curr ? curr : best;
	}
	return best;
}
