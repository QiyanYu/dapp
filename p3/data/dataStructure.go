package data

import (
	"fmt"
	"sync"
)

type DataQueue struct {
	items           map[string]string
	insertedRecords map[string]string
	mux             sync.Mutex
}

func NewDataQueue() DataQueue {
	q := DataQueue{
		items:           make(map[string]string),
		insertedRecords: make(map[string]string),
	}
	return q
}

func (queue *DataQueue) Add(key string, value string) {
	queue.mux.Lock()
	fmt.Println("1")
	fmt.Println("11", queue.insertedRecords[key])
	if queue.insertedRecords[key] == "" || queue.insertedRecords[key] != value {
		fmt.Println("2")
		queue.items[key] = value
		queue.insertedRecords[key] = value
	}
	queue.mux.Unlock()
}

func (queue *DataQueue) Get() (bool, map[string]string) {
	queue.mux.Lock()
	defer queue.mux.Unlock()
	if len(queue.items) == 0 {
		return false, nil
	}
	returnMap := queue.items
	queue.items = map[string]string{}
	return true, returnMap
}

func (queue *DataQueue) Show() {
	for k, v := range queue.items {
		fmt.Println("key: ", k, "value: ", v)
	}
}

func TestDataQueue() {
	data := NewDataQueue()
	data.Add("key", "value")
	data.Add("key2", "value2")
	is, result := data.Get()
	fmt.Println("is ", is, "result ", result)
	fmt.Println("now queue")
	for key, value := range data.items {
		fmt.Printf("Key: %s\tValue: %v\n", key, value)
	}
}
