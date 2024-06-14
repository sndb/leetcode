char *evaluate(char *t, int *result) {
	int sum = 0;
	int sign = 1;
	for (; *t; t++) {
		if (*t == ' ') {
			continue;
		} else if (*t == '+') {
			sign = 1;
		} else if (*t == '-') {
			sign = -1;
		} else if (*t == '(') {
			int n;
			t = evaluate(t + 1, &n);
			sum += n * sign;
		} else if (*t == ')') {
			break;
		} else {
			int n = 0;
			while (*t && *t >= '0' && *t <= '9') {
				n *= 10;
				n += *t - '0';
				t++;
			}
			t--;
			sum += n * sign;
		}
	}
	*result = sum;
	return t;
}

int calculate(char *s) {
	int result;
	evaluate(s, &result);
	return result;
}
