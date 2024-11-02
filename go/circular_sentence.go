package main

func isCircularSentence(sentence string) bool {
	for i := 1; i < len(sentence)-1; i++ {
		if sentence[i] == ' ' && sentence[i-1] != sentence[i+1] {
			return false
		}
	}
	return sentence[0] == sentence[len(sentence)-1]
}
