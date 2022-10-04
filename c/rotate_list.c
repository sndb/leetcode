#include <stdlib.h>

struct ListNode {
	int              val;
	struct ListNode *next;
};

struct ListNode *
rotateRight(struct ListNode *head, int k) {
	if (!head)
		return NULL;
	int length = 0;
	for (struct ListNode *c = head; c; c = c->next)
		length++;
	k %= length;
	if (k == 0)
		return head;

	struct ListNode *a = head, *b = head;
	for (int i = 0; i < k; i++)
		b = b->next;
	while (b->next) {
		a = a->next;
		b = b->next;
	}

	struct ListNode *r = a->next;
	a->next            = NULL;
	b->next            = head;
	return r;
}
