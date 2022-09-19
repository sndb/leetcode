#include <stdbool.h>
#include <stdlib.h>

struct ListNode {
	int              val;
	struct ListNode *next;
};

struct ListNode *
reverse(struct ListNode *head) {
	struct ListNode *curr = head, *prev = NULL, *next = NULL;
	while (curr) {
		next       = curr->next;
		curr->next = prev;
		prev       = curr;
		curr       = next;
	}
	return prev;
}

bool
equal(struct ListNode *l, struct ListNode *m) {
	while (l && m) {
		if (l->val != m->val)
			return false;
		l = l->next;
		m = m->next;
	}
	return true;
}

bool
isPalindrome(struct ListNode *head) {
	struct ListNode *slow = head, *fast = head;
	while (fast && fast->next) {
		slow = slow->next;
		fast = fast->next->next;
	}
	return equal(head, reverse(slow));
}
