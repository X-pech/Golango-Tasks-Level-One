package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Пара ключ+значение
*/
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

type ConcurrentMap[K comparable, V any] interface {
	Add(K, V)
	Get(K) V
}

/*
Енам для типа запроса
*/
type QueryType int

const (
	GET QueryType = 0
	ADD QueryType = 1
)

/*
Сам запрос - включает в себя типа, канал для ответов и встроенную пару.
Можно сделать через интерфейс, но тогда нужно выносить поля в геттеры-сеттеры
Что для такой задачки мне лень делать
*/
type Query[K comparable, V any] struct {
	queryType       QueryType
	responseChannel chan V
	Pair[K, V]
}

/*
Сама мапа. Включает в себя мапу и канал для запросов.
*/
type QueryMap[K comparable, V any] struct {
	data         map[K]V
	queryChannel chan Query[K, V]
}

/*
При инициализации не забывает запустить конкурентно listener
*/
func NewQueryMap[K comparable, V any](channelSize int) QueryMap[K, V] {
	m := new(QueryMap[K, V])
	m.data = make(map[K]V)
	m.queryChannel = make(chan Query[K, V], channelSize)
	go m.listener()
	return *m
}

/*
Обработчик запросов: ему прилетает запрос, он смотрит на его тип
И соответствующим образом обрабатывает
*/
func (cm *QueryMap[K, V]) listener() {
	for {
		query := <-cm.queryChannel
		switch query.queryType {
		case ADD:
			cm.data[query.Key] = query.Value
		case GET:
			query.responseChannel <- cm.data[query.Key]
		}
	}
}

/*
Add() кидает запрос на добавление в очередь
*/
func (cm *QueryMap[K, V]) Add(key K, value V) {
	cm.queryChannel <- Query[K, V]{
		Pair: Pair[K, V]{
			Key:   key,
			Value: value,
		},
		queryType: ADD,
	}
}

/*
Get кидает запрос на получение в очередь
Что важно - создаётся отдельный канал для получение только нашего ответа
Это необходимо для того чтобы наш ответ не перепутался со всеми остальными
Как отличными по значениию так и по хронологии
*/
func (cm *QueryMap[K, V]) Get(key K) V {
	responseChannel := make(chan V)
	cm.queryChannel <- Query[K, V]{
		Pair: Pair[K, V]{
			Key: key,
		},
		responseChannel: responseChannel,
		queryType:       GET,
	}
	v := <-responseChannel
	close(responseChannel)
	return v
}

/*
Решил переписать во время устных вопросов.
Эта реализация не требует каких-то огромных каналов
А пользуется простым советским RWMutex
Который как раз нужен нам здесь - чтобы конкуретно читать
Но не писать
+ Он гарантирует упорядоченность операций то есть мы
Не прочитаем данные которые *ещё не записались*
*/
type RWMMap[K comparable, V any] struct {
	data map[K]V
	rwm  *sync.RWMutex
}

func NewRWMMap[K comparable, V any]() RWMMap[K, V] {
	w := new(RWMMap[K, V])
	w.data = make(map[K]V)
	w.rwm = new(sync.RWMutex)
	return *w
}

/*Тут мы блочим мапу на запись поэтому Lock*/
func (rw *RWMMap[K, V]) Add(key K, value V) {
	rw.rwm.Lock()
	defer rw.rwm.Unlock()
	rw.data[key] = value
}

/*А тут только на чтение поэтому RLock который позволяет вызывать другие
RLock но блокировать Lock то есть читать можно писать в это время нельзя*/
func (rw *RWMMap[K, V]) Get(key K) V {
	rw.rwm.RLock()
	defer rw.rwm.RUnlock()
	return rw.data[key]
}

/*
тест на основе интерфейса
*/
func testcase(cm ConcurrentMap[int, int]) {
	cm.Add(5, 10)
	rand.Seed(time.Now().Unix())

	N := 10
	M := 10
	c := make(chan Pair[int, int], N*M)
	for i := 0; i < N; i++ {
		go func() {
			for j := 0; j < M; j++ {
				key := rand.Int()
				value := rand.Int()
				cm.Add(key, value)
				checkValue := cm.Get(key)
				if checkValue != value {
					fmt.Printf("Oh no! %d is not equals %d\n", value, checkValue)
				}
				c <- Pair[int, int]{Key: key, Value: value}
			}
		}()
	}

	ncmap := make(map[int]int)
	for i := 0; i < N*M; i++ {
		p := <-c
		ncmap[p.Key] = p.Value
	}

	time.Sleep(3 * time.Second)

	for i := range ncmap {
		res := cm.Get(i)
		if ncmap[i] != res {
			fmt.Printf("ncmap[%d] == %d != %d == cm[%d]\n", i, ncmap[i], res, i)
		}
	}

	fmt.Println(cm.Get(5))
	fmt.Println(cm.Get(10))
}

func main() {
	qm := NewQueryMap[int, int](5)
	testcase(&qm)

	rm := NewRWMMap[int, int]()
	testcase(&rm)
}
