package goLeetCode

import "fmt"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	res, wordMap := make([][]string, 0), make(map[string]bool)
	for _, w := range wordList {
		wordMap[w] = true
	}
	if !wordMap[endWord] {
		return res
	}
	queue := make([][]string, 0)
	queue = append(queue, []string{beginWord})
	queueLen := 1
	levelMap := make(map[string]bool)
	for len(queue) > 0 {
		fmt.Println(queue)
		path := queue[0]
		queue = queue[1:]
		lastWord := path[len(path)-1]
		for i := 0; i < len(lastWord); i++ {
			for c := 'a'; c <= 'z'; c++ {
				nextWord := lastWord[:i] + string(c) + lastWord[i+1:]
				if nextWord == endWord {
					path = append(path, endWord)
					res = append(res, path)
					continue
				}
				if wordMap[nextWord] {
					levelMap[nextWord] = true
					newPath := make([]string, len(path))
					copy(newPath, path)
					newPath = append(newPath, nextWord)
					queue = append(queue, newPath)
				}
			}
		}
		queueLen--
		if queueLen == 0 {
			if len(res) > 0 {
				return res
			}
			for k := range levelMap {
				delete(wordMap, k)
			}
			levelMap = make(map[string]bool)
			queueLen = len(queue)
		}
	}
	return res
}
