package transition_words_test

import (
	"go-dsa/reimplementations/transition_words"
	"testing"
)

func TestTransitionWords(t *testing.T) {
	startWord := "cat"
	endWord := "dog"
	wordList := []string{"cat", "hat", "hot", "hog", "dog", "cot", "cog", "pat", "pot", "dot"}

	result := transition_words.TransitionWords(startWord, endWord, wordList)
	expected := 4

	if result != expected {
		t.Errorf("Expected ladder length: %d, but got: %d", expected, result)
	}
}

func TestTransitionWords_NoSolution(t *testing.T) {
	startWord := "cat"
	endWord := "dog"
	wordList := []string{"cat", "hat", "hot", "hog", "pot", "dot"}

	result := transition_words.TransitionWords(startWord, endWord, wordList)
	expected := -1

	if result != expected {
		t.Errorf("Expected ladder length: %d, but got: %d", expected, result)
	}
}

func TestTransitionWords_StartAndEndSame(t *testing.T) {
	startWord := "cat"
	endWord := "cat"
	wordList := []string{"cat", "hat", "hot", "hog", "dog", "cot", "cog", "pat", "pot", "dot"}

	result := transition_words.TransitionWords(startWord, endWord, wordList)
	expected := 0

	if result != expected {
		t.Errorf("Expected ladder length: %d, but got: %d", expected, result)
	}
}

func TestTransitionWords_EndWordNotInList(t *testing.T) {
	startWord := "cat"
	endWord := "dog"
	wordList := []string{"cat", "hat", "hot", "hog", "pot", "dot"}

	result := transition_words.TransitionWords(startWord, endWord, wordList)
	expected := -1

	if result != expected {
		t.Errorf("Expected ladder length: %d, but got: %d", expected, result)
	}
}

func TestTransitionWords_EmptyWordList(t *testing.T) {
	startWord := "cat"
	endWord := "dog"
	var wordList []string

	result := transition_words.TransitionWords(startWord, endWord, wordList)
	expected := -1

	if result != expected {
		t.Errorf("Expected ladder length: %d, but got: %d", expected, result)
	}
}

func TestTransitionWords_NoTransitionNeeded(t *testing.T) {
	startWord := "cat"
	endWord := "cat"
	wordList := []string{"cat", "hat", "hot", "hog", "dog"}

	result := transition_words.TransitionWords(startWord, endWord, wordList)
	expected := 0

	if result != expected {
		t.Errorf("Expected ladder length: %d, but got: %d", expected, result)
	}
}

func TestTransitionWords_SingleTransition(t *testing.T) {
	startWord := "cat"
	endWord := "dog"
	wordList := []string{"cat", "cot", "cog", "dog"}

	result := transition_words.TransitionWords(startWord, endWord, wordList)
	expected := 4

	if result != expected {
		t.Errorf("Expected ladder length: %d, but got: %d", expected, result)
	}
}

func TestTransitionWords_LongerTransitions(t *testing.T) {
	startWord := "hit"
	endWord := "cog"
	wordList := []string{"hit", "hot", "dot", "dog", "lot", "log", "cog"}

	result := transition_words.TransitionWords(startWord, endWord, wordList)
	expected := 5

	if result != expected {
		t.Errorf("Expected ladder length: %d, but got: %d", expected, result)
	}
}
