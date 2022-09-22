#include <string.h>

int
romanValue(char c) {
	switch (c) {
	case 'I':
		return 1;
	case 'V':
		return 5;
	case 'X':
		return 10;
	case 'L':
		return 50;
	case 'C':
		return 100;
	case 'D':
		return 500;
	case 'M':
		return 1000;
	}
	return 0;
}

int
romanToInt(char *s) {
	int l = strlen(s), x = 0;
	for (int i = 0; i < l; i++) {
		int c = romanValue(s[i]), n = romanValue(s[i + 1]);
		x += n > c ? -c : c;
	}
	return x;
}
