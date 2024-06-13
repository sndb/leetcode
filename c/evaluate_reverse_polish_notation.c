int evalRPN(char **tokens, int tokensSize) {
	int stack[10000];
	int k = 0;
	for (int i = 0; i < tokensSize; i++) {
		char *tok = tokens[i];
		switch (tok[0]) {
		case '*':
			stack[k - 2] *= stack[k - 1];
			k--;
			break;
		case '/':
			stack[k - 2] /= stack[k - 1];
			k--;
			break;
		case '+':
			stack[k - 2] += stack[k - 1];
			k--;
			break;
		case '-':
			if (tok[1] == '\0') {
				stack[k - 2] -= stack[k - 1];
				k--;
				break;
			}
		default:
			stack[k] = atoi(tok);
			k++;
		}
	}
	return stack[0];
}
