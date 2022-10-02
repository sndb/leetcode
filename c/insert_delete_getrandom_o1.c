#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define INIT_VEC_CAP 8
#define N_BUCKETS    256

int
hash(int val) {
	return abs(val);
}

typedef struct {
	int *data;
	int  len;
	int  cap;
} Vector;

Vector *
vectorCreate() {
	int     len  = 0;
	int     cap  = INIT_VEC_CAP;
	int    *data = malloc(sizeof(int[cap]));
	Vector *v    = malloc(sizeof(Vector));
	v->len       = len;
	v->cap       = cap;
	v->data      = data;
	return v;
}

void
vectorExpand(Vector *v) {
	v->cap *= 2;
	v->data = realloc(v->data, sizeof(int[v->cap]));
}

int
vectorFind(Vector *v, int val) {
	for (int i = 0; i < v->len; i++)
		if (v->data[i] == val)
			return i;
	return -1;
}

int
vectorFindIndex(Vector *idxs, Vector *vals, int val) {
	for (int i = 0; i < idxs->len; i++)
		if (vals->data[idxs->data[i]] == val)
			return i;
	return -1;
}

int
vectorInsert(Vector *v, int val) {
	if (v->len == v->cap)
		vectorExpand(v);
	int idx = v->len;
	v->len++;
	v->data[idx] = val;
	return idx;
}

void
vectorRemoveIndex(Vector *v, int idx) {
	v->data[idx] = v->data[v->len - 1];
	v->len--;
}

bool
vectorRemove(Vector *v, int val) {
	if (v->len == 0)
		return false;
	int idx = vectorFind(v, val);
	if (idx == -1)
		return false;
	vectorRemoveIndex(v, idx);
	return true;
}

void
vectorFree(Vector *v) {
	free(v->data);
	free(v);
}

typedef struct {
	Vector **buckets;
	Vector  *nums;
	int      n;
} RandomizedSet;

RandomizedSet *
randomizedSetCreate() {
	int      n       = N_BUCKETS;
	Vector **buckets = malloc(sizeof(Vector *[n]));
	for (int i = 0; i < n; i++)
		buckets[i] = vectorCreate();
	RandomizedSet *s = malloc(sizeof(RandomizedSet));
	s->n             = n;
	s->buckets       = buckets;
	s->nums          = vectorCreate();
	return s;
}

bool
randomizedSetInsert(RandomizedSet *obj, int val) {
	unsigned int h = hash(val) % N_BUCKETS;
	if (vectorFindIndex(obj->buckets[h], obj->nums, val) != -1)
		return false;
	int idx = vectorInsert(obj->nums, val);
	vectorInsert(obj->buckets[h], idx);
	return true;
}

bool
randomizedSetRemove(RandomizedSet *obj, int val) {
	unsigned int h = hash(val) % N_BUCKETS;

	int iidx = vectorFindIndex(obj->buckets[h], obj->nums, val);
	if (iidx == -1)
		return false;
	int nidx = obj->buckets[h]->data[iidx];
	vectorRemoveIndex(obj->buckets[h], iidx);

	int lidx = obj->nums->len - 1;
	int last = obj->nums->data[lidx];
	vectorRemoveIndex(obj->nums, nidx);

	if (lidx != nidx) {
		int lhsh = hash(last) % N_BUCKETS;
		vectorRemove(obj->buckets[lhsh], lidx);
		vectorInsert(obj->buckets[lhsh], nidx);
	}
	return true;
}

int
randomizedSetGetRandom(RandomizedSet *obj) {
	return obj->nums->data[rand() % obj->nums->len];
}

void
randomizedSetFree(RandomizedSet *obj) {
	for (int i = 0; i < obj->n; i++)
		vectorFree(obj->buckets[i]);
	vectorFree(obj->nums);
	free(obj->buckets);
	free(obj);
}

int
main(void) {
	float start = (float)clock() / CLOCKS_PER_SEC;

	RandomizedSet *s = randomizedSetCreate();
	for (int i = 0; i < 100000; i++)
		randomizedSetInsert(s, i);
	for (int i = 0; i < 50000; i++)
		randomizedSetRemove(s, i);
	for (int i = 0; i < 5; i++)
		printf("%d\n", randomizedSetGetRandom(s));
	randomizedSetFree(s);

	float end = (float)clock() / CLOCKS_PER_SEC;
	printf("%f\n", end - start);
	return 0;
}
