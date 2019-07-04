package AvlTree

import (
	"ericivan/Queue"
	"ericivan/stack"
	"fmt"
	"testing"
)

func TestAvlPut(t *testing.T) {

	AvlMapData := AvlMap{
		LinkList: stack.StackItem{},
		Size:     0,
	}
	for i := 1; i < 10; i++ {
		key := i
		value := string(i)
		AvlMapData.Put(key, value)
	}

	AvlMapData.LevelOrder()

}

func TestQueue(t *testing.T) {
	q := Queue.ItemQueue{}
	q.New()

	AvlMapData := AvlMap{
		LinkList: stack.StackItem{},
		Size:     0,
		Root: &AvlEntry{
			Key:   4,
			Value: "0",
		},
	}

	for i := 1; i < 10; i++ {
		key := i
		value := string(i)
		AvlMapData.Put(key, value)
	}

	q.Enqueue(AvlMapData.Root)

	//pAddr := q.Dequeue()

	stack := AvlMapData.LinkList

	for j := 1; j < 10; j++ {
		data := stack.Pop()

		if data != nil {
			fmt.Println(data.(*AvlEntry))
		}
	}

}

func TestPointer(t *testing.T) {

	entry := &AvlEntry{
		Key:   1,
		Value: "1",
	}

	fmt.Println(*entry)
}

func TestRotate(t *testing.T) {
	AvlMapData := AvlMap{
		LinkList: stack.StackItem{},
		Size:     0,
	}

	AvlMapData.Put(1, "1")

	fmt.Println(AvlMapData.Root)

	p := AvlMapData.Root

	AvlMapData.Put(2, "2")

	p = p.Right
	fmt.Println(p)

	node3 := AvlMapData.Put(3, "3")

	p = p.Right
	fmt.Println(p)

	//LeftRotate(AvlMapData.Root, &AvlMapData)

	AvlMapData.Put(4, "4")

	p = p.Right
	fmt.Println(p)

	AvlMapData.Put(5, "5")

	p = p.Right
	fmt.Println(p)

	LeftRotate(node3, &AvlMapData)
}

func TestPut(t *testing.T) {
	AvlMapData := AvlMap{
		LinkList: stack.StackItem{},
		Size:     0,
	}

	AvlMapData.Put(1, "1")
	AvlMapData.Put(2, "2")
	AvlMapData.Put(3, "3")

	fmt.Println(AvlMapData)

}
