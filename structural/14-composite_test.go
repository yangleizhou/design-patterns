package structural

import "testing"

func TestComposite(t *testing.T) {
	root := NewComponent(CompositeNode, "root")
	p1 := NewComponent(CompositeNode, "p1")
	p2 := NewComponent(CompositeNode, "p2")

	c1 := NewComponent(LeafNode, "c1")
	c2 := NewComponent(LeafNode, "c2")

	root.AddChild(p1)
	root.AddChild(p2)
	root.AddChild(c2)
	p1.AddChild(p2)
	p1.AddChild(c2)
	p2.AddChild(c1)
}

func TestCompositeRing(t *testing.T) {
	root := NewComponent(CompositeNode, "root")
	p1 := NewComponent(CompositeNode, "p1")

	root.AddChild(p1)
	p1.AddChild(root)
}
