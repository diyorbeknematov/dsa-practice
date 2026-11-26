type User struct {
    id        int
    following map[int]bool
}

type Tweet struct {
    id   int
    time int
}

type HeapItem struct {
    userId int
    tweet  Tweet
    index  int
}

type Twitter struct {
    users    map[int]*User
    tweets   map[int][]Tweet
    time     int
}


func Constructor() Twitter {
    return Twitter{
		users: make(map[int]*User),
		tweets: make(map[int][]Tweet),
	}
}

type PriorityQueue []*HeapItem

func (pq PriorityQueue) Len() int {
    return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].tweet.time > pq[j].tweet.time
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
    item := x.(*HeapItem)
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[:n-1]
    return item
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
	this.getOrCreateUser(userId)
    tweet := Tweet{
		id: tweetId,
		time: this.time,
	}

	this.tweets[userId] = append(this.tweets[userId], tweet)
	this.time ++

	if len(this.tweets[userId]) > 10 {
		this.tweets[userId] = this.tweets[userId][1:]
	}
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    pq := &PriorityQueue{}
	heap.Init(pq)
	user := this.users[userId]

	if user == nil {
		return []int{}
	}

	userIDs := []int{userId}
	for followeeID := range user.following {
		userIDs = append(userIDs, followeeID)
	}

	tweetIDs := make([]int, 0, 10)

	for _, id := range userIDs {
		if len(this.tweets[id]) > 0 {
			item := HeapItem{
				userId: id,
				tweet: this.tweets[id][len(this.tweets[id])-1],
				index: len(this.tweets[id])-1,
			}
			heap.Push(pq, &item)
		}
	}

	for pq.Len() > 0 && len(tweetIDs) < 10 {
		item := heap.Pop(pq).(*HeapItem)
		tweetIDs = append(tweetIDs, item.tweet.id)

		prevIndex := item.index - 1
		if prevIndex >= 0 {
			prevTweet := this.tweets[item.userId][prevIndex]

			item := HeapItem{
				userId: item.userId,
				tweet: prevTweet,
				index: prevIndex,
			}
			heap.Push(pq, &item)
		}
	}

	return tweetIDs
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
    follower := this.getOrCreateUser(followerId)
	this.getOrCreateUser(followeeId)

	follower.following[followeeId] = true
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    follower, ok := this.users[followerId]
    if !ok {
        return
    }

    if _, ok := this.users[followeeId]; !ok {
        return
    }

    delete(follower.following, followeeId)
}

func (this *Twitter) getOrCreateUser(userId int) *User {
	if _, ok := this.users[userId]; !ok {
		user := &User{
			id: userId,
			following: make(map[int]bool),
		}
		this.users[userId] = user
		return user
	}

	return this.users[userId]
}

