package tree

import (
	"errors"
	"sort"
)

//Record is a struct containing int fields ID and Parent
type Record struct {
	ID, Parent int
}

//Node is a struct containing int field ID and []*Node field Children.
type Node struct {
	ID       int
	Children []*Node
}

//Build a Node struct tree from a set of Record structs
func Build(recs []Record) (*Node, error) {

	sort.Slice(recs, func(i, j int) bool {
		return recs[i].ID < recs[j].ID
	})

	parents := make(map[int]*Node)
	var root *Node
	for x, rec := range recs {
		if x > 0 && recs[x-1].ID == recs[x].ID {
			return nil, errors.New("duplicated record")
		}
		if recs[x].ID > 0 && recs[x].ID == recs[x].Parent {
			return nil, errors.New("self parent")
		}
		if x > 0 && recs[x].ID-recs[x-1].ID > 1 {
			return nil, errors.New("non-continuous")
		}

		//create root node
		if rec.ID == 0 {
			if rec.Parent != 0 {
				return nil, errors.New("root node has parent")
			}
			root = &Node{ID: 0}
			parents[0] = root
			continue
		}
		p := parents[rec.Parent]
		if p == nil {
			return nil, errors.New("parent not found")
		}
		c := &Node{ID: rec.ID}
		p.Children = append(p.Children, c)
		parents[c.ID] = c
	}
	return root, nil
}
