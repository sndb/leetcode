package main

func mergeAlternately(word1 string, word2 string) string {
	i := 0
	s := []byte{}
	for i < len(word1) && i < len(word2) {
		s = append(s, word1[i], word2[i])
		i++
	}
	s = append(s, word1[i:]...)
	s = append(s, word2[i:]...)
	return string(s)
}
