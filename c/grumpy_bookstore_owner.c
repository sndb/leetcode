int maxSatisfied(int *customers, int customersSize, int *grumpy, int grumpySize, int minutes) {
	int curr = 0;
	for (int i = 0; i < minutes; i++)
		curr += customers[i];
	for (int i = minutes; i < customersSize; i++)
		curr += customers[i] * (1 - grumpy[i]);

	int max = curr;
	for (int i = 0, j = minutes; j < customersSize; i++, j++) {
		curr -= customers[i] * grumpy[i];
		curr += customers[j] * grumpy[j];
		if (curr > max) max = curr;
	}
	return max;
}
