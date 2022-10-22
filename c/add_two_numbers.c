#include <stdlib.h>

struct ListNode {
	int val;
	struct ListNode *next;
};

struct ListNode *
addTwoNumbers(struct ListNode *l1, struct ListNode *l2)
{
	struct ListNode *dummy = malloc(sizeof(*dummy));
	struct ListNode *prev = dummy;
	int carry = 0;
	while (l1 || l2 || carry) {
		struct ListNode *n = malloc(sizeof(*n));
		int l1v = l1 ? l1->val : 0;
		int l2v = l2 ? l2->val : 0;
		int x = l1v + l2v + carry;
		n->val = x % 10;
		carry = x / 10;
		if (l1)
			l1 = l1->next;
		if (l2)
			l2 = l2->next;
		prev->next = n;
		prev = n;
	}
	prev->next = NULL;
	struct ListNode *r = dummy->next;
	free(dummy);
	return r;
}
