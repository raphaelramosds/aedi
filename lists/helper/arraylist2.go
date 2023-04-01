package helper

import "errors"

type ArrayList struct {
	values []int
	size   int
}

// Internal functions

func (al *ArrayList) Init(size int) error {
	if size > 0 {
		al.values = make([]int, size)
		return nil
	} else {
		return errors.New("can't init arraylist with size <= 0")
	}
}

func (al *ArrayList) double() {
	doubledValues := make([]int, len(al.values)*2)
	for i := 0; i < len(al.values); i++ {
		doubledValues[i] = al.values[i]
	}
	al.values = doubledValues
}

// Interface functions

func (al *ArrayList) Add(value int) {
	if al.size == len(al.values) {
		al.double()
	}
	al.values[al.size] = value
	al.size++
}

func (al *ArrayList) AddOnIndex(value int, index int) error {
	if index >= 0 && index <= al.size {
		if al.size == len(al.values) {
			al.double()
		}
		for i := al.size; i > index; i-- {
			al.values[i] = al.values[i-1]
		}
		al.values[index] = value
		al.size++
		return nil
	} else {
		if index < 0 {
			return errors.New("can't add in index < 0")
		} else {
			return errors.New("can't add in index > arraylist.size")
		}
	}
}

func (al *ArrayList) RemoveOnIndex(index int) error {
	if index >= 0 && index < al.size {
		for i := index; i < al.size-1; i++ {
			al.values[i] = al.values[i+1]
		}
		al.size--
		return nil
	} else {
		if index < 0 {
			return errors.New("can't remove from index < 0")
		} else {
			return errors.New("can't remove from index >= arraylist.size")
		}
	}
}

func (al *ArrayList) Get(index int) (int, error) {
	if index >= 0 && index < al.size {
		return al.values[index], nil
	} else {
		if index < 0 {
			return index, errors.New("can't get value from index < 0")
		} else {
			return index, errors.New("can't get value from index >= arraylist.size")
		}
	}
}

func (arraylist *ArrayList) Set(value, index int) error {
	if index >= 0 && index < arraylist.size {
		arraylist.values[index] = value
		return nil
	} else {
		if index < 0 {
			return errors.New("can't set value on arraylist on index < 0")
		} else {
			return errors.New("can't set value on arraylist on index >= arraylist.size")
		}
	}
}

func (arraylist *ArrayList) Size() int {
	return arraylist.size
}
