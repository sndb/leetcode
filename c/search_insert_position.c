int searchInsert(int *nums, int numsSize, int target) {
	int min = 0;
	int max = numsSize - 1;
	while (min <= max) {
		int i = (min + max) / 2;
		if (nums[i] == target)
			return i;
		if (nums[i] > target)
			max = i - 1;
		else
			min = i + 1;
	}
	return min;
}
