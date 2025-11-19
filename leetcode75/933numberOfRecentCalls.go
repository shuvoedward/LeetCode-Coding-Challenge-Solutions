package leetcode75

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: []int{},
	}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.queue = append(rc.queue, t)

	for rc.queue[0] < t-3000 {
		rc.queue = rc.queue[1:]
	}

	return len(rc.queue)
}
