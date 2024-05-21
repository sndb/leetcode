int maxArea(int *height, int heightSize) {
	int i = 0, j = heightSize - 1, r = 0;
	while (i < j) {
		int a;
		if (height[i] > height[j]) {
			a = (j - i) * height[j];
			j--;
		} else {
			a = (j - i) * height[i];
			i++;
		}
		if (a > r)
			r = a;
	}
	return r;
}
