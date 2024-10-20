package main

func parseBoolExpr(expr string) bool {
	if expr == "t" {
		return true
	}
	if expr == "f" {
		return false
	}

	inner := expr[2 : len(expr)-1]
	if expr[0] == '!' {
		return !parseBoolExpr(inner)
	}

	ret := expr[0] == '&'
	for i, j, p := 0, 0, 0; j <= len(inner); j++ {
		if j == len(inner) || p == 0 && inner[j] == ',' {
			if expr[0] == '|' {
				ret = ret || parseBoolExpr(inner[i:j])
			} else {
				ret = ret && parseBoolExpr(inner[i:j])
			}
			i = j + 1
		} else if inner[j] == '(' {
			p++
		} else if inner[j] == ')' {
			p--
		}
	}
	return ret
}
