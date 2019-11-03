package Leaderboard

import "sort"

type Leaderboard struct {
	scores map[int]int
	players []int
}

func Constructor() Leaderboard {
	players := make([]int, 10000)
	scores := make(map[int]int)
	return Leaderboard{scores: scores, players: players}
}


func (this *Leaderboard) AddScore(playerId int, score int)  {
	if _, ok := this.scores[playerId]; ok {
		this.scores[playerId] += score
	} else {
		this.scores[playerId] = score
	}
	this.players[playerId] += score
}


func (this *Leaderboard) Top(K int) int {
	temp := make([]int, 10000)
	copy(temp, this.players)
	sort.Ints(temp)
	res := 0
	for i := 0; i < K && i < len(temp); i++ {
		res += temp[len(temp) - 1 - i]
	}
	return res
}


func (this *Leaderboard) Reset(playerId int)  {
	this.players[playerId] = 0
	this.scores[playerId] = 0
}


/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */