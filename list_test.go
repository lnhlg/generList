package generList

import "testing"

func TestList_DeleteOf(t *testing.T) {

	list := make(List[int], 5, 10)
	for i := 0; i < 5; i++ {
		list[i] = i
	}

	list = list.DeleteOf(2)
	list = list.DeleteOf(2)
	list = list.DeleteOf(2)

	t.Log(len(list))
	t.Log(cap(list))
}
