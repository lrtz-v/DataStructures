package ladderLength

func ladderLength(beginWord string, endWord string, wordList []string) int {

	wordMap := make(map[string]int, len(wordList))
	for i, v := range wordList {
		wordMap[v] = i
	}

	if _, ok := wordMap[endWord]; !ok {
		return 0
	}

	
	return -1
}
