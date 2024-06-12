inline void swap(int *a, int i, int j) {
	int tmp = a[i];
	a[i] = a[j];
	a[j] = tmp;
}

void sortColors(int *nums, int numsSize) {
	int i = 0, j = 0, k = numsSize - 1;
	while (j <= k)
		switch (nums[j]) {
		case 0:
			swap(nums, i++, j++);
			break;
		case 1:
			j++;
			break;
		case 2:
			swap(nums, j, k--);
			break;
		}
}
