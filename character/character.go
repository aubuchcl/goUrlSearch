package character

import "sort"

//Character struct is a type to use when evaluating char count
type Character struct {
	Char  string
	Count int
}

//NewCharacter init a new character type
func NewCharacter() Character {
	newChar := Character{
		Char:  "",
		Count: 0,
	}
	return newChar
}

//CharSort use to sort stripped and itemized character slices
func CharSort(slc []Character) []Character {
	//redo this function with regex
	sort.SliceStable(slc, func(i, j int) bool {
		return slc[i].Count > slc[j].Count
	})
	//fmt.Println("By Char:", slc)
	return slc
}
