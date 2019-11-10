package Twitter

import "sort"

type Twitter struct {
	follows map[int][]int
	times map[int][]int
	tweets []int
}


/** Initialize your data structure here. */
func Constructor() Twitter {
	times, follows := make(map[int][]int), make(map[int][]int)
	tweets := make([]int, 0)
	return Twitter{follows: follows, times: times, tweets: tweets}
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	this.tweets = append(this.tweets, tweetId)
	this.times[userId] = append(this.times[userId], len(this.tweets) - 1)
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	res, times := make([]int, 0), make([]int, 0)
	for _, v := range this.times[userId] {
		times = append(times, v)
	}
	for _, followId := range this.follows[userId] {
		for _, v := range this.times[followId] {
			times = append(times, v)
		}
	}
	sort.Ints(times)
	count := 0
	for i := len(times) - 1; i >= 0; i-- {
		res = append(res, this.tweets[times[i]])
		count++
		if count == 10 {
			break
		}
	}
	return res
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	if followerId != followeeId && !inArray(followeeId, this.follows[followerId]) {
		this.follows[followerId] = append(this.follows[followerId], followeeId)
	}
}

func inArray(needle int, hayStack []int) bool {
	for _, v := range hayStack {
		if v == needle {
			return true
		}
	}
	return false
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	index := -1
	for i, v := range this.follows[followerId] {
		if v == followeeId {
			index = i
			break
		}
	}
	if index >= 0 {
		this.follows[followerId] = append(this.follows[followerId][:index], this.follows[followerId][index + 1:]...)
	}
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */