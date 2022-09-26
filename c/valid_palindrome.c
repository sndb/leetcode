#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>
#include <string.h>

bool
isPalindrome(char *s) {
	int len = strlen(s);
	int l = 0, r = len - 1;
	while (l < len && r >= 0) {
		char cl = tolower(s[l]);
		char cr = tolower(s[r]);
		if (!isalnum(cl)) {
			l++;
			continue;
		}
		if (!isalnum(cr)) {
			r--;
			continue;
		}
		if (cl != cr)
			return false;
		l++;
		r--;
	}
	return true;
}

int
main(void) {
	printf("%d\n", isPalindrome("A man, a plan, a canal: Panama"));
}
