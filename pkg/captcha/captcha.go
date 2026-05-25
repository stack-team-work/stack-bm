package captcha

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type item struct {
	answer string
	expire time.Time
}

var store sync.Map

func init() {
	go func() {
		for {
			time.Sleep(30 * time.Second)
			store.Range(func(key, value interface{}) bool {
				if time.Now().After(value.(item).expire) {
					store.Delete(key)
				}
				return true
			})
		}
	}()
}

func Generate() (id string, question string) {
	a := rand.Intn(20) + 1
	b := rand.Intn(20) + 1
	id = fmt.Sprintf("%d", time.Now().UnixNano())
	question = fmt.Sprintf("%d + %d = ?", a, b)
	store.Store(id, item{answer: fmt.Sprintf("%d", a+b), expire: time.Now().Add(5 * time.Minute)})
	return
}

func Verify(id, answer string) bool {
	v, ok := store.Load(id)
	if !ok {
		return false
	}
	it := v.(item)
	if time.Now().After(it.expire) {
		store.Delete(id)
		return false
	}
	if it.answer == answer {
		store.Delete(id)
		return true
	}
	return false
}
