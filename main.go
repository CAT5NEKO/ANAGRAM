package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func generateAnagrams(input string) []string {
	anagramsMap := make(map[string]bool)
	runes := []rune(input)

	var generate func(int)
	generate = func(n int) {
		if n == len(runes)-1 {
			anagram := string(runes)
			if anagram != input {
				anagramsMap[anagram] = true
			}
			return
		}

		for i := n; i < len(runes); i++ {
			runes[n], runes[i] = runes[i], runes[n]
			generate(n + 1)
			runes[n], runes[i] = runes[i], runes[n]
		}
	}

	generate(0)

	anagrams := make([]string, 0, len(anagramsMap))
	for anagram := range anagramsMap {
		anagrams = append(anagrams, anagram)
	}

	sort.Strings(anagrams)

	return anagrams
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("並べ替えたい文字列入れてね: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("何か読み取れへんのやけど😡:", err)
		return
	}

	input = strings.TrimSpace(input)

	anagrams := generateAnagrams(input)

	file, err := os.Create("anagrams.md")
	if err != nil {
		fmt.Println("ファイル作れないんやけど😡:", err)
		return
	}
	defer file.Close()

	if _, err := io.WriteString(file, fmt.Sprintf("#アナグラムを生成  %s\n\n", input)); err != nil {
		fmt.Println("ファイル書き込み失敗しました😡:", err)
		return
	}

	for _, anagram := range anagrams {
		if _, err := io.WriteString(file, fmt.Sprintf("- %s\n", anagram)); err != nil {
			fmt.Println("ファイル書き込み失敗しました😡:", err)
			return
		}
	}

	fmt.Println("アナグラム生成しました。 anagrams.md に残してるよ。")
}
