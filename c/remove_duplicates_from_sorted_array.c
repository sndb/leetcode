int
removeDuplicates(int *nums, int numsSize) {
	int j = 0;
	for (int i = 0; i < numsSize; i++) {
		if (nums[i] > nums[j]) {
			j++;
			nums[j] = nums[i];
		}
	}
	return j + 1;
}
