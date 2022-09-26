#include <stdlib.h>
#include <string.h>

char *
longestCommonPrefix(char **strs, int strsSize) {
	char *r = calloc(strlen(strs[0]) + 1, sizeof(char));
	for (int i = 0; i < strlen(strs[0]); i++) {
		for (int j = 0; j < strsSize - 1; j++)
			if (strs[j][i] != strs[j + 1][i])
				goto out;
		r[i] = strs[0][i];
	}
out:
	return r;
}
