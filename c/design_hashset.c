#include <stdbool.h>
#include <stdlib.h>

typedef struct Node {
	int          key;
	struct Node *next;
} Node;

typedef struct {
	Node **buckets;
	int    n;
} MyHashSet;

Node *
nodeCreate(int key) {
	Node *new = malloc(sizeof(Node));
	new->key  = key;
	new->next = NULL;
	return new;
}

void
nodeFree(Node *node) {
	if (!node)
		return;
	nodeFree(node->next);
	free(node);
}

MyHashSet *
myHashSetCreate() {
	int        n       = 256;
	Node     **buckets = calloc(n, sizeof(Node *));
	MyHashSet *set     = malloc(sizeof(MyHashSet));
	set->n             = n;
	set->buckets       = buckets;
	return set;
}

void
myHashSetAdd(MyHashSet *obj, int key) {
	int   i    = key % obj->n;
	Node *b    = obj->buckets[i];
	Node *prev = NULL;
	while (b) {
		if (b->key == key)
			return;
		prev = b;
		b    = b->next;
	}
	Node *new = nodeCreate(key);
	if (prev)
		prev->next = new;
	else
		obj->buckets[i] = new;
}

void
myHashSetRemove(MyHashSet *obj, int key) {
	int   i    = key % obj->n;
	Node *b    = obj->buckets[i];
	Node *prev = NULL;
	while (b) {
		if (b->key == key) {
			if (prev)
				prev->next = b->next;
			else
				obj->buckets[i] = b->next;
			b->next = NULL;
			nodeFree(b);
			return;
		}
		prev = b;
		b    = b->next;
	}
}

bool
myHashSetContains(MyHashSet *obj, int key) {
	int   i = key % obj->n;
	Node *b = obj->buckets[i];
	while (b) {
		if (b->key == key)
			return true;
		b = b->next;
	}
	return false;
}

void
myHashSetFree(MyHashSet *obj) {
	for (int i = 0; i < obj->n; i++)
		nodeFree(obj->buckets[i]);
	free(obj->buckets);
	free(obj);
}
