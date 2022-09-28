#include <stdlib.h>

typedef struct Node {
	int          key;
	int          value;
	struct Node *next;
} Node;

Node *
nodeCreate(int key, int value) {
	Node *new  = malloc(sizeof(Node));
	new->key   = key;
	new->value = value;
	new->next  = NULL;
	return new;
}

void
nodeFree(Node *node) {
	if (!node)
		return;
	nodeFree(node->next);
	free(node);
}

typedef struct {
	Node **buckets;
	int    n;
} MyHashMap;

MyHashMap *
myHashMapCreate() {
	int        n       = 512;
	Node     **buckets = calloc(n, sizeof(Node *));
	MyHashMap *map     = malloc(sizeof(MyHashMap));
	map->n             = n;
	map->buckets       = buckets;
	return map;
}

void
myHashMapPut(MyHashMap *obj, int key, int value) {
	int   i = key % obj->n;
	Node *b = obj->buckets[i], *prev = NULL;
	while (b) {
		if (b->key == key) {
			b->value = value;
			return;
		}
		prev = b;
		b    = b->next;
	}
	Node *new = nodeCreate(key, value);
	if (prev)
		prev->next = new;
	else
		obj->buckets[i] = new;
}

int
myHashMapGet(MyHashMap *obj, int key) {
	int   i = key % obj->n;
	Node *b = obj->buckets[i];
	while (b) {
		if (b->key == key)
			return b->value;
		b = b->next;
	}
	return -1;
}

void
myHashMapRemove(MyHashMap *obj, int key) {
	int   i = key % obj->n;
	Node *b = obj->buckets[i], *prev = NULL;
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

void
myHashMapFree(MyHashMap *obj) {
	for (int i = 0; i < obj->n; i++)
		nodeFree(obj->buckets[i]);
	free(obj->buckets);
	free(obj);
}
