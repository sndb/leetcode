#include <stdlib.h>

struct ListNode {
	int              val;
	struct ListNode *next;
};

struct ListNode *
prependLastNode(struct ListNode *head) {
	struct ListNode *prev = NULL, *curr = head;
	while (curr->next) {
		prev = curr;
		curr = curr->next;
	}
	prev->next = NULL;
	curr->next = head;
	return curr;
}

struct ListNode *
reverseList(struct ListNode *head) {
	if (!head || !head->next)
		return head;
	head       = prependLastNode(head);
	head->next = reverseList(head->next);
	return head;
}

struct ListNode *
reverseListFast(struct ListNode *head) {
	struct ListNode *prev = NULL, *curr = head, *next = NULL;
	while (curr) {
		next       = curr->next;
		curr->next = prev;
		prev       = curr;
		curr       = next;
	}
	return prev;
}
