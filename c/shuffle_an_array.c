#include <stdlib.h>
#include <string.h>

typedef struct {
	int *base;
	int *nums;
	int  size;
} Solution;

Solution *
solutionCreate(int *nums, int numsSize) {
	Solution *obj = malloc(sizeof(Solution));
	obj->size     = numsSize;
	obj->base     = malloc(sizeof(int) * numsSize);
	obj->nums     = malloc(sizeof(int) * numsSize);
	memcpy(obj->base, nums, sizeof(int) * numsSize);
	memcpy(obj->nums, nums, sizeof(int) * numsSize);
	return obj;
}

int *
solutionReset(Solution *obj, int *retSize) {
	memcpy(obj->nums, obj->base, sizeof(int) * obj->size);
	*retSize = obj->size;
	return obj->nums;
}

int *
solutionShuffle(Solution *obj, int *retSize) {
	for (int i = 0; i < obj->size; i++) {
		int r        = rand() % obj->size;
		int t        = obj->nums[i];
		obj->nums[i] = obj->nums[r];
		obj->nums[r] = t;
	}
	*retSize = obj->size;
	return obj->nums;
}

void
solutionFree(Solution *obj) {
	free(obj->nums);
	free(obj->base);
	free(obj);
}
