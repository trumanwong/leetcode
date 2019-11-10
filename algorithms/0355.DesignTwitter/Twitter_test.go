package main

import (
	"leetcode/algorithms/0355.DesignTwitter/Twitter"
	"testing"
)

func TestTwitter(t *testing.T)  {
	tests := []struct{
		operates []string
		values [][]int
		output [][]interface{}
	}{
		{[]string{"Twitter","postTweet","getNewsFeed","follow","postTweet","getNewsFeed","unfollow","getNewsFeed"},
		[][]int{{},{1,5},{1},{1,2},{2,6},{1},{1,2},{1}},
		[][]interface{}{nil,nil,{5},nil,nil,{6, 5},nil,{5}}},
	}

	for _, test := range tests {
		twitter := Twitter.Twitter{}
		for i, operate := range test.operates {
			if operate == "Twitter" {
				twitter = Twitter.Constructor()
			} else if operate == "postTweet" {
				twitter.PostTweet(test.values[i][0], test.values[i][1])
			} else if operate == "getNewsFeed" {
				ret := twitter.GetNewsFeed(test.values[i][0])
				for j, v := range test.output[i] {
					if v.(int) != ret[j] {
						t.Errorf("Got %v for GetNewsFeed(%d); expected %v", ret, test.values[i][0], test.output[i])
						break
					}
				}
			} else if operate == "follow" {
				twitter.Follow(test.values[i][0], test.values[i][1])
			} else if operate == "unfollow" {
				twitter.Unfollow(test.values[i][0], test.values[i][1])
			}
		}
	}
}