#include <stdbool.h>
#include <stdlib.h>

int
compar(const void *p1, const void *p2) {
	int *i1 = *(int **)p1;
	int *i2 = *(int **)p2;
	if (i1[1] < i2[1])
		return -1;
	if (i1[1] > i2[1])
		return 1;
	return 0;
}

bool
overlapHelper(int *i, int *j) {
	if (i[0] > j[0] && i[0] < j[1])
		return true;
	if (i[1] > j[0] && i[1] < j[1])
		return true;
	return false;
}

bool
overlap(int *i, int *j) {
	if (i[0] == j[0] && i[1] == j[1])
		return true;
	return overlapHelper(i, j) || overlapHelper(j, i);
}

int
eraseOverlapIntervals(int **intervals, int intervalsSize, int *intervalsColSize) {
	qsort(intervals, intervalsSize, sizeof(int *), compar);

	int  r    = 0;
	int *prev = intervals[0];
	for (int i = 1; i < intervalsSize; i++)
		if (overlap(prev, intervals[i]))
			r++;
		else
			prev = intervals[i];
	return r;
}
