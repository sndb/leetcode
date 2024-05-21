#include <stdlib.h>

struct ListNode {
	int val;
	struct ListNode *next;
};

typedef struct {
	struct ListNode *list;
} Solution;

Solution *solutionCreate(struct ListNode *head) {
	Solution *s = malloc(sizeof(*s));
	s->list = head;
	return s;
}

int solutionGetRandom(Solution *s) {
	int choice;
	double seen = 0;
	for (struct ListNode *node = s->list; node; node = node->next) {
		seen++;
		double n = 1L << 31;
		double r = random() / n;
		if (r < 1 / seen)
			choice = node->val;
	}
	return choice;
}

void solutionFree(Solution *s) {
	free(s);
}
