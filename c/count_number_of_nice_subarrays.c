int numberOfSubarrays(int *nums, int numsSize, int k) {
	int start = 0;
	int tail = 0;
	int sum = 0;
	for (int i = 0; i < numsSize; i++) {
		if (nums[i] & 1) {
			tail = 0;
			k--;
		}
		while (k == 0) {
			if (nums[start] & 1)
				k++;
			tail++;
			start++;
		}
		sum += tail;
	}
	return sum;
}
