int
maxProfit(int *prices, int pricesSize) {
	int min = prices[0], r = 0;
	for (int i = 1; i < pricesSize; i++) {
		int c = prices[i] - min;
		if (c > r)
			r = c;
		if (prices[i] < min)
			min = prices[i];
	}
	return r;
}
