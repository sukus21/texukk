package texukk

import (
	"fmt"
	"image"
	"image/draw"
)

func (a *Atlas) Render() (image.Image, map[string]image.Rectangle, error) {
	a.initTree()
	namedOutput := make(map[string]image.Rectangle)

	//Write everything to a priority queue
	pq := new(pQueue[*AtlasEntry])
	for i := range a.sprite {
		v := &a.sprite[i]
		pq.Add(v.Dx()*v.Dx()+v.Dy()*v.Dy(), v)
	}

	//Place it on the thing
	dest := image.NewRGBA(image.Rect(0, 0, a.size.X, a.size.Y))
	for _, v := range pq.Sort() {
		if err := a.tree.place(v); err != nil {
			return nil, nil, err
		}
		draw.Draw(
			dest,
			v.Rectangle,
			v.source,
			image.Point{},
			draw.Over,
		)
		if v.name != "" {
			namedOutput[v.name] = v.Rectangle
		}
	}

	return dest, namedOutput, nil
}

func (node *atlasLeaf) place(e *AtlasEntry) error {
	if node == nil {
		return fmt.Errorf("node is empty")
	}

	//Already occupied, fit into children
	if node.entry != nil {

		//Try to insert into first child
		err := node.child1.place(e)
		if err == nil {
			return nil
		}

		//Insert into second child
		return node.child2.place(e)
	}

	//Am I too small for this?
	if node.Dx() < e.Dx() || node.Dy() < e.Dy() {
		return fmt.Errorf("node size <%d,%d> cannot fit element size <%d,%d>", node.Dx(), node.Dy(), e.Dx(), e.Dy())
	}

	//Split into children if size isn't an exact match
	if node.Dx() != e.Dx() {
		node.child1 = new(atlasLeaf)
		node.child1.Min.X = node.Min.X + e.Dx()
		node.child1.Min.Y = node.Min.Y
		node.child1.Max.X = node.Max.X
		node.child1.Max.Y = node.Min.Y + e.Dy()
	}
	if node.Dy() != e.Dy() {
		node.child2 = new(atlasLeaf)
		node.child2.Min.Y = node.Min.Y + e.Dy()
		node.child2.Min.X = node.Min.X
		node.child2.Max = node.Max
	}

	//Insert entry into node
	node.entry = e
	node.Max.X = node.Min.X + e.Dx()
	node.Max.Y = node.Min.Y + e.Dy()
	e.Rectangle = node.Rectangle
	return nil
}
