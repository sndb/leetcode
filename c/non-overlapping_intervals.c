#include <stdbool.h>
#include <stdlib.h>

int
compar(const void *p1, const void *p2) {
	return (*(int **)p1)[1] - (*(int **)p2)[1];
}

bool
overlap(int *i, int *j) {
	if (i[0] > j[0] && i[0] < j[1]
	    || i[1] > j[0] && i[1] < j[1]
	    || j[0] > i[0] && j[0] < i[1]
	    || j[1] > i[0] && j[1] < i[1]
	    || i[0] == j[0] && i[1] == j[1])
		return true;
	return false;
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
