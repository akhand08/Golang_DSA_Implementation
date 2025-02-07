package singlylinkedlist

import "testing"

func TestAppend(t *testing.T) {

	list := New()
	list.Append(10)
	list.Append(20)
	list.Append(30)

	if list.Size() != 3 {
		t.Errorf("Size should be 3, got %d", list.Size())
	}

	if list.Head.Data != 10 {

		t.Errorf("Value of first Element should be 10, got %d", list.Head.Data)
	}

}
