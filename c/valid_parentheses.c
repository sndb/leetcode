#include <stdbool.h>
#include <string.h>

bool
isValid(char *s) {
	const char open[]  = "({[";
	const char close[] = ")}]";
	char       stack[10000];
	int        head = 0;
	char      *c;
	for (int i = 0; i < strlen(s); i++) {
		if ((c = strchr(open, s[i]))) {
			stack[++head] = close[c - open];
			continue;
		}
		if ((c = strchr(close, s[i])) && stack[head--] != *c)
			return false;
	}
	return head == 0;
}
