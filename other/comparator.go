package other

type Student struct {
	Id   int
	Age  int
	Name string
}

type IdUpAgeDown []*Student

func (iuad IdUpAgeDown) Less(i, j int) bool {
	if iuad[i].Id != iuad[j].Id {
		return iuad[i].Id < iuad[j].Id
	} else {
		return iuad[j].Age < iuad[i].Age
	}
}

func (iuad IdUpAgeDown) Len() int {
	return len(iuad)
}

func (iuad IdUpAgeDown) Swap(i, j int) {
	iuad[i], iuad[j] = iuad[j], iuad[i]
}
