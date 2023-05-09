package main

import (
	"fmt"
)

func levenshteinDistance(s, t string) int {
	m := len(s)
	n := len(t)

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			cost := 1
			if s[i-1] == t[j-1] {
				cost = 0
			}
			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+cost)
		}
	}

	return dp[m][n]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}

func findClosestMatch(commandOutput string, correctResults []string) string {
	closestMatch := ""
	minDistance := -1

	for _, result := range correctResults {
		distance := levenshteinDistance(commandOutput, result)
		if minDistance == -1 || distance < minDistance {
			minDistance = distance
			closestMatch = result
		}
	}

	return closestMatch
}

func main() {
	commandOutput := ""
	for commandOutput != "exit"{
	correctResults := []string{"cd", "mkdir", "ls", "pwd"}

	fmt.Scanf("%s", &commandOutput)

	closestMatch := findClosestMatch(commandOutput, correctResults)
	if(commandOutput == closestMatch){
		fmt.Println("正しいコマンドです")
	}else{
        fmt.Println("そのコマンドは存在しません")
	    fmt.Println("入力値", commandOutput)
	    fmt.Println("もしかして…？", closestMatch)
	}}
	fmt.Println("システムを終了します")
    
}
