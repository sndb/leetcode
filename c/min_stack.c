#include <stdlib.h>

typedef struct Node {
	int          val;
	struct Node *next;
} Node;

typedef struct {
	Node *top;
	Node *min;
} MinStack;

MinStack *
minStackCreate() {
	MinStack *obj = malloc(sizeof(MinStack));
	obj->top      = NULL;
	obj->min      = NULL;
	return obj;
}

void
minStackPush(MinStack *obj, int val) {
	Node *new = malloc(sizeof(Node));
	new->val  = val;
	new->next = obj->top;
	obj->top  = new;

	Node *min = malloc(sizeof(Node));
	min->val  = obj->min && obj->min->val < val ? obj->min->val : val;
	min->next = obj->min;
	obj->min  = min;
}

void
minStackPop(MinStack *obj) {
	Node *head = obj->top;
	obj->top   = head->next;
	free(head);

	Node *min = obj->min;
	obj->min  = min->next;
	free(min);
}

int
minStackTop(MinStack *obj) {
	return obj->top->val;
}

int
minStackGetMin(MinStack *obj) {
	return obj->min->val;
}

void
minStackFree(MinStack *obj) {
	Node *curr;
	while (obj->top) {
		curr     = obj->top;
		obj->top = obj->top->next;
		free(curr);
	}
	while (obj->min) {
		curr     = obj->min;
		obj->min = obj->min->next;
		free(curr);
	}
	free(obj);
}
