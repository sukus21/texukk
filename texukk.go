package texukk

import (
	"fmt"
	"image"
)

type AtlasEntry struct {
	id     uint
	name   string
	source image.Image
	image.Rectangle
}

type atlasLeaf struct {
	image.Rectangle
	entry  *AtlasEntry
	child1 *atlasLeaf
	child2 *atlasLeaf
}

type Atlas struct {
	tree   atlasLeaf
	sprite []AtlasEntry
	names  map[string]*AtlasEntry
	size   image.Point
	ids    uint
}

func NewAtlas(width int, height int) *Atlas {
	atlas := new(Atlas)
	atlas.size.X = width
	atlas.size.Y = height

	atlas.initTree()
	atlas.sprite = make([]AtlasEntry, 0, 128)
	atlas.names = make(map[string]*AtlasEntry)

	return atlas
}

func (a *Atlas) initTree() {
	a.tree = atlasLeaf{
		entry: nil,
		Rectangle: image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: a.size,
		},
	}
}

func (a *Atlas) Add(img image.Image) *AtlasEntry {
	a.sprite = append(a.sprite, AtlasEntry{
		source: img,
		id:     a.ids,
		Rectangle: image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: img.Bounds().Max,
		},
	})
	a.ids++
	return &a.sprite[len(a.sprite)-1]
}

func (a *Atlas) AddNamed(img image.Image, name string) (*AtlasEntry, error) {
	//Check for duplicates
	if _, ok := a.names[name]; ok {
		return nil, fmt.Errorf("texoven: name '%s' already exists in atlas", name)
	}

	a.sprite = append(a.sprite, AtlasEntry{
		source: img,
		id:     a.ids,
		name:   name,
		Rectangle: image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: img.Bounds().Max,
		},
	})
	entry := &a.sprite[len(a.sprite)-1]
	a.names[name] = entry
	a.ids++
	return entry, nil
}

func (a *Atlas) Remove(e *AtlasEntry) bool {
	for i, v := range a.sprite {
		if v.id == e.id {
			a.sprite[i] = a.sprite[len(a.sprite)-1]
			a.sprite = a.sprite[:len(a.sprite)-1]
			return true
		}
	}
	return false
}

func (a *Atlas) RemoveNamed(name string) bool {
	if entry, ok := a.names[name]; ok {
		delete(a.names, name)
		a.Remove(entry)
	}
	return false
}
