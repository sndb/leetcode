#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

bool
isAnagram(char *s, char *t) {
	size_t sl    = strlen(s);
	size_t tl    = strlen(t);
	int    a[26] = {0};
	int    b[26] = {0};
	if (sl != tl)
		return false;
	for (int i = 0; i < sl; i++) {
		a[s[i] - 'a']++;
		b[t[i] - 'a']++;
	}
	for (int i = 0; i < 26; i++)
		if (a[i] != b[i])
			return false;
	return true;
}
