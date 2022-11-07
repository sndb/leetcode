int searchInsert(int *nums, int numsSize, int target) {
	int min = 0;
	int max = numsSize - 1;
	while (min <= max) {
		int mid = (min + max) / 2;
		if (nums[mid] > target)
			max = mid - 1;
		else if (nums[mid] < target)
			min = mid + 1;
		else
			return mid;
	}
	return min;
}
