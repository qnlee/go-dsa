package transition_words

/*
Given startWord, endWord: is it possible to transition startWord to endWord one letter at a time?
Each transition can only be 1 letter swap, and each transition word must be a valid word.

eg) startWord = "cat", endWord = "dog"
	cat -> hat -> hot -> hog -> dog
	cat -> cot -> cog -> dog
	cat -> pat -> pot -> dot -> dog
*/

type partialDict struct {
	wordSet map[string]bool
}

func makePartialDict(wordList []string) *partialDict {
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}
	return &partialDict{wordSet: wordSet}
}

func (d *partialDict) IsValidWord(word string) bool {
	return d.wordSet[word]
}

func transitionWords(startWord string, endWord string, wordList []string) int {
	if startWord == endWord {
		return 0
	}

	dict := makePartialDict(wordList)
	if !dict.IsValidWord(startWord) || !dict.IsValidWord(endWord) {
		return -1
	}

	beginSet := map[string]bool{startWord: true}
	endSet := map[string]bool{endWord: true}
	visited := make(map[string]bool)
	level := 1
	for len(beginSet) > 0 && len(endSet) > 0 {
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet // swap to always extend the smaller set
		}

		nextLevelSet := make(map[string]bool)

		for word := range beginSet {
			for j := 0; j < len(word); j++ {
				for c := 'a'; c <= 'z'; c++ {
					newWord := word[:j] + string(c) + word[j+1:]
					if endSet[newWord] {
						return level + 1 // reached the endWord, return the level
					}
					if dict.IsValidWord(newWord) && !visited[newWord] {
						nextLevelSet[newWord] = true // enqueue the valid transition word
						visited[newWord] = true      // mark the word as visited
					}
				}
			}
		}
		beginSet = nextLevelSet
		level++
	}

	return -1 // unable to reach the endWord
}
