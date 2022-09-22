int
hammingDistance(int x, int y) {
	int n = x ^ y;
	int i = 0;
	while (n > 0) {
		i += n & 1;
		n >>= 1;
	}
	return i;
}
