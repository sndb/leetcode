struct ListNode {
	int              val;
	struct ListNode *next;
};

struct ListNode *
removeNthFromEnd(struct ListNode *head, int n) {
	head = &(struct ListNode){0, head};

	struct ListNode *prev = head, *last = head;
	for (int i = 0; i < n; i++)
		last = last->next;
	while (last->next) {
		prev = prev->next;
		last = last->next;
	}
	prev->next = prev->next->next;
	return head->next;
}
