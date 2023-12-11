package main

import (
	"fmt"
	"sync"
	"time"
)

func (q *TimedQueue) getQueueStatus() string {
	q.mu.Lock()
	defer q.mu.Unlock()

	result := ""
	for _, item := range q.items {
		// result += fmt.Sprint(item.SetTime) + " " + fmt.Sprint(int(time.Since(item.sline).Seconds())) + " || "

		result += fmt.Sprintf("mySET: %d myLifeTime: %d || ", item.SetTime, int(time.Since(item.sline).Seconds()))
	}

	if len(q.items) == 0 {
		result = "Empty || "
	}

	return result + fmt.Sprint(10-int(time.Since(q.stime).Seconds()))
}

type TimedQueue struct {
	mu       sync.Mutex
	items    []Item
	timerSet *time.Timer
	stime    time.Time
	SetTime  int
}

type Item struct {
	Value   string
	SetTime int
	sline   time.Time
}

func NewTimedQueue() *TimedQueue {
	return &TimedQueue{
		items: make([]Item, 0),
	}
}

func (q *TimedQueue) enqueue(item string) {
	q.mu.Lock()
	defer q.mu.Unlock()

	// 만약 큐가 empty라면 setTime를 10초로 설정하고, 아니라면 이전 아이템의 setTime을 가져온다.

	if len(q.items) == 0 {
		q.items = append(q.items, Item{Value: item, SetTime: 0, sline: time.Now()})
		q.stime = time.Now()

		q.SetTime = 10
		q.timerSet = time.AfterFunc(10*time.Second, func() { q.dequeue() })
		fmt.Println("timer set")
	} else {

		// 현재 타이머의 남은 시간
		elapsed := q.SetTime - int(time.Since(q.stime).Seconds())

		fmt.Println("elapsed: ", elapsed)

		var pst int

		// 맨 앞에 하나를 제외한 (현재 카운트 중인) 나머지 아이템들의 setTime을 더한다.
		for _, item := range q.items[1:] {
			pst += item.SetTime
		}

		preSetTime := elapsed + pst

		mySet := 10 - preSetTime

		q.items = append(q.items, Item{Value: item, SetTime: mySet, sline: time.Now()})

	}

}

func (q *TimedQueue) dequeue() {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.items) > 0 {

		itme := q.items[0]
		q.items = q.items[1:]

		fmt.Print("Dequeue: ", int(time.Since(itme.sline).Seconds()), " ", itme.Value, " ", itme.SetTime, " || ")

		q.timerSet.Stop()

		fmt.Println("timer stopped")

		if len(q.items) > 0 {
			q.stime = time.Now()
			q.SetTime = q.items[0].SetTime
			q.timerSet = time.AfterFunc(time.Duration(q.items[0].SetTime)*time.Second, func() { q.dequeue() })
			fmt.Println("timer set")
		}

	}
}
