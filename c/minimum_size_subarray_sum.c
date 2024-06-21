int minSubArrayLen(int target, int *nums, int numsSize) {
	int result = numsSize + 1;
	for (int i = 0, j = 0, sum = 0; j < numsSize; j++) {
		sum += nums[j];
		while (sum >= target) {
			int d = j - i + 1;
			if (d < result) result = d;
			sum -= nums[i++];
		}
	}
	return result > numsSize ? 0 : result;
}
