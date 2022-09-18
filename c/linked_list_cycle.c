#include <stdbool.h>
#include <stdlib.h>

struct ListNode {
	int              val;
	struct ListNode *next;
};

bool
hasCycle(struct ListNode *head) {
	struct ListNode *next = NULL, *seen = NULL;

	while (head) {
		for (struct ListNode *n = seen; n; n = n->next)
			if (head == n)
				return true;

		next       = head->next;
		head->next = seen;
		seen       = head;
		head       = next;
	}
	return false;
}
