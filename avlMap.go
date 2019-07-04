package AvlTree

import (
	"ericivan/Queue"
	"ericivan/stack"
	"fmt"
	"math"
)

type AvlMap struct {
	LinkList stack.StackItem
	Root     *AvlEntry
	Size     int
}

func compare(a int, b int) int {
	if a == b {
		return 0
	}
	if a-b > 0 {
		return 1
	} else {
		return -1
	}
}

func (m *AvlMap) Put(key int, value string) *AvlEntry {

	var retNode *AvlEntry

	if m.Root == nil {
		m.Root = NewAvlEntry(key, value)
		retNode = m.Root
		m.LinkList.Push(m.Root)
		m.Size++
	} else {
		p := m.Root

		for p != nil {
			res := compare(key, p.Key)

			if res == 0 {
				p.Value = value

				retNode = p
				break
			} else if res < 0 {
				if p.Left == nil {
					p.Left = NewAvlEntry(key, value)

					p.Left.Parent = p

					m.LinkList.Push(p.Left)

					m.Size++

					retNode = p.Left
					break
				} else {
					p = p.Left
				}
			} else {

				if p.Right == nil {

					p.Right = NewAvlEntry(key, value)

					p.Right.Parent = p

					m.LinkList.Push(p.Right)

					m.Size++

					retNode = p.Right

					break
				} else {
					p = p.Right
				}
			}
		}
	}

	m.fixAfterIntersion(key)

	return retNode
}

func (m *AvlMap) Delete(p *AvlEntry, k int) *AvlEntry {
	if p == nil {
		return nil
	} else {
		compaireResult := compare(k, p.Key)

		if compaireResult == 0 {

			if p.Left == nil && p.Right == nil {
				p = nil
			} else if p.Left != nil && p.Right == nil {
				p = p.Left
			} else if p.Left == nil && p.Right != nil {
				p = p.Right
			} else {
				if m.Size == 1 {
					rightMin := getFirstEntry(p.Right)

					p.Key = rightMin.Key

					p.Value = rightMin.Value

					newRight := m.Delete(p.Right, p.Key)

					p.Right = newRight
				} else {
					leftMax := getLastEntry(p.Left)

					p.Key = leftMax.Key
					p.Value = leftMax.Value

					newLeft := m.Delete(p.Left, p.Key)

					p.Left = newLeft
				}
			}
		} else if compaireResult < 0 {
			newLeft := m.Delete(p.Left, k)
			p.Left = newLeft
		} else {
			newRight := m.Delete(p.Right, k)
			p.Right = newRight
		}

		return p
	}
}

func getFirstEntry(p *AvlEntry) *AvlEntry {
	if p == nil {
		return nil
	}

	for p.Left != nil {
		p = p.Left
	}

	return p
}

func getLastEntry(p *AvlEntry) *AvlEntry {

	if p == nil {
		return nil
	}

	for p.Right != nil {
		p = p.Right
	}

	return p
}

func (root *AvlMap) LevelOrder() {
	queue := Queue.ItemQueue{}
	queue.Enqueue(*root.Root)

	preCount := 1
	pCount := 0

	for !queue.IsEmpty() {
		preCount--

		p := queue.Dequeue()

		avl := p.(AvlEntry)

		fmt.Println(avl)

		if avl.Left != nil {
			queue.Enqueue(*avl.Left)
			pCount++
		}

		if avl.Right != nil {
			queue.Enqueue(*avl.Right)
			pCount++
		}

		if preCount == 0 {

			preCount = pCount
			pCount = 0
			fmt.Println()
		}
	}
}

func (m *AvlMap) fixAfterIntersion(key int) {
	p := m.Root

	itemStack := m.LinkList

	for !itemStack.IsEmpty() {
		p := itemStack.Pop()

		node := p.(*AvlEntry)

		newHeight := math.Max(getHeight(node.Left), getHeight(node.Right)) + 1

		if node.Height > 1 && int(newHeight) == node.Height {
			itemStack.Clear()
			return
		}

		node.Height = int(newHeight)

		d := getHeight(node.Left) - getHeight(node.Right)

		if math.Abs(d) <= 1 {
			continue
		} else {

			if d == 2 {
				if key < node.Left.Key {
					node = RightRotate(node, m)
				} else {
					node = LeftRotate(node.Left, m)
					node = RightRotate(node, m)
				}
			} else if d == -2 {
				if key > node.Right.Key {
					node = LeftRotate(node, m)
				} else {
					node = RightRotate(node.Right, m)
					node = LeftRotate(node, m)
				}
			}

			if !itemStack.IsEmpty() {
				peekPtr := itemStack.Peek()

				peekData := peekPtr.(*AvlEntry)

				if compare(key, peekData.Key) < 0 {
					peekData.Left = node
				} else {
					peekData.Right = node
				}
			}

		}
	}

	m.Root = p
}

func getHeight(entry *AvlEntry) float64 {

	if entry == nil {
		return 0
	} else {
		return float64(entry.Height)
	}
}

func LeftRotate(entry *AvlEntry, avl *AvlMap) *AvlEntry {
	subRight := entry.Right
	subSLeft := subRight.Left
	parent := entry.Parent

	entry.Parent = subRight
	subRight.Left = entry

	subRight.Parent = parent
	parent.Right = subRight

	entry.Right = subSLeft
	if subSLeft != nil {
		subSLeft.Parent = entry
	}

	if entry == avl.Root {
		avl.Root = subRight
	}

	return subRight
}

func RightRotate(entry *AvlEntry, avl *AvlMap) *AvlEntry {

	subLeft := entry.Left
	subSRight := subLeft.Left
	parent := entry.Parent

	entry.Parent = subLeft
	subLeft.Right = entry

	subLeft.Parent = parent
	parent.Left = subLeft

	entry.Left = subSRight

	if subSRight != nil {
		subSRight.Parent = entry
	}

	if avl.Root == entry {
		avl.Root = subLeft
	}

	return subLeft
}

func Parent(entry *AvlEntry) *AvlEntry {
	if entry.Parent != nil {
		return entry.Parent
	}
	return nil
}

func Left(entry *AvlEntry) *AvlEntry {
	if entry.Left != nil {
		return entry.Left
	}
	return nil
}

func Right(entry *AvlEntry) *AvlEntry {

	if entry.Right != nil {
		return entry.Right
	}

	return nil
}
