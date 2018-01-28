package serializetree

import "testing"

func TestNoData(t *testing.T) {
	tree := NewTree()
	actual := tree.Marshal()
	expected := "{\"head\": null}"
	if actual != expected {
		t.Errorf("Bad marchalling, expected %s, actual was %s", expected, actual)
	}
}

func TestBasicSerialization(t *testing.T) {
	tree := NewTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	actual := tree.Marshal()
	expected := `{"head": {"val": 1, "left": null, "right": {"val": 2, "left": null, "right": {"val": 3, "left": null, "right": null}}}}`
	if actual != expected {
		t.Errorf("Bad marchalling, expected %s, actual was %s", expected, actual)
	}
}
