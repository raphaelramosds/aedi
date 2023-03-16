package helper

type ArrayList struct {
	values []int
	tam    int
}

func (al *ArrayList) Init() {
	al.values = make([]int, 10)
}

func (al *ArrayList) double() {
	doubledValues := make([]int, len(al.values)*2)
	for i := 0; i < len(al.values); i++ {
		doubledValues[i] = al.values[i]
	}
	al.values = doubledValues
}

func (al *ArrayList) Add(value int) {
	if al.tam == len(al.values) {
		al.double()
	}
	al.values[al.tam] = value
	al.tam++
}

func (al *ArrayList) AddOnIndex(value int, index int) {
	if al.tam == len(al.values) {
		al.double()
	}
	for i := al.tam; i > index; i-- {
		al.values[i] = al.values[i-1]
	}
	al.values[index] = value
	al.tam++
}

func (al *ArrayList) Remove() {
	if al.tam > 0 {
		al.tam--
	} //TODO: add error handling
}

func (al *ArrayList) RemoveOnIndex(index int) {
	if index >= 0 && index < al.tam {
		for i := index; i < al.tam; i++ {
			al.values[i] = al.values[i+1]
		}
		al.tam--
	}
}

func (al *ArrayList) Get(index int) int {
	if index >= 0 && index < al.tam {
		return al.values[index]
	}
	return -1 //TODO: add error handling
}

func (al *ArrayList) Set(value, index int) {
	if index >= 0 && index < al.tam {
		al.values[index] = value
	}
	//else //TODO: add error handling
}

func (al *ArrayList) Size() int {
	return al.tam
}
