package generList

type List[T any] []T

func (l List[T]) DeleteOf(index uint) List[T] {

	newList := append(l[:index], l[index+1:]...)

	length := float64(len(newList))
	capacity := float64(cap(newList))
	if length/capacity <= 0.25 {

		shrinkList := make(List[T], int(length), int(capacity/2))
		for i := range newList {

			shrinkList[i] = newList[i]
		}

		return shrinkList
	}

	return newList
}
