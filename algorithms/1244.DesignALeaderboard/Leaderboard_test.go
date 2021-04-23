package main

import (
	"fmt"
	"leetcode/algorithms/1244.DesignALeaderboard/Leaderboard"
	"testing"
)

func TestLeaderboard(t *testing.T) {
	tests := []struct {
		opts  []string
		input [][]int
	}{
		{[]string{"Leaderboard", "addScore", "addScore", "addScore", "addScore", "addScore", "top",
			"reset", "reset", "addScore", "top"},
			[][]int{{}, {1, 73}, {2, 56}, {3, 39}, {4, 51}, {5, 4}, {1}, {1}, {2}, {2, 51}, {3}},
		},
	}

	for _, test := range tests {
		leaderboard := Leaderboard.Leaderboard{}
		for i := 0; i < len(test.opts); i++ {
			if test.opts[i] == "Leaderboard" {
				leaderboard = Leaderboard.Constructor()
			} else if test.opts[i] == "addScore" {
				leaderboard.AddScore(test.input[i][0], test.input[i][1])
			} else if test.opts[i] == "top" {
				fmt.Println("top: ", leaderboard.Top(test.input[i][0]))
			} else if test.opts[i] == "reset" {
				leaderboard.Reset(test.input[i][0])
			}
		}
	}
}
