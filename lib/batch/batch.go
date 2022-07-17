package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {

	var wg sync.WaitGroup
	var m sync.Mutex
	sem := make(chan struct{}, pool)
	var i int64 = 0

	for i = 0; i < n; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(id int64) {
			defer wg.Done()
			u := getOne(id)
			<-sem
			m.Lock()
			res = append(res, u)
			m.Unlock()
		}(i)
	}

	wg.Wait()
	return res
}
