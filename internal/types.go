package internal

type DaySolution struct {
	Year           int
	Day            int
	SolutionResult any
	RunAll         RunAllFunc
	RunPartOne     PartOneFunc
	RunPartTwo     PartTwoFunc
}

type RunAllFunc func(year int, day int, partOneInput []string, partTwoInput []string) (partOneResult any, partTwoResult any)
type PartOneFunc func(input []string, year int, day int) any
type PartTwoFunc func(input []string, year int, day int) any

func NewDaySolution(year int, day int, allFunc RunAllFunc, oneFunc PartOneFunc, twoFunc PartTwoFunc) *DaySolution {
	return &DaySolution{
		Year:       year,
		Day:        day,
		RunAll:     allFunc,
		RunPartOne: oneFunc,
		RunPartTwo: twoFunc,
	}
}
