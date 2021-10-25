package lighthouse_solver

const SplitRockTop = 1
const SplitRockBottom = -1
const RoundIslandTop = 10
const RoundIslandBottom = -10
const MarbleheadTop = 100
const MarbleheadBottom = -100
const FortNiagaraTop = 1000
const FortNiagaraBottom = -1000

type Card struct {
	Id    int
	Parts [4]int
}

func (c *Card) Rotate() {
	first := c.Parts[3]
	c.Parts[3] = c.Parts[2]
	c.Parts[2] = c.Parts[1]
	c.Parts[1] = c.Parts[0]
	c.Parts[0] = first
}

var cards = []*Card{
	{Id: 6, Parts: [4]int{MarbleheadTop, RoundIslandTop, MarbleheadBottom, FortNiagaraTop}},
	{Id: 7, Parts: [4]int{MarbleheadTop, SplitRockBottom, RoundIslandBottom, FortNiagaraTop}},
	{Id: 5, Parts: [4]int{RoundIslandTop, SplitRockTop, MarbleheadBottom, SplitRockTop}},
	{Id: 8, Parts: [4]int{RoundIslandTop, SplitRockBottom, FortNiagaraTop, MarbleheadTop}},
	{Id: 4, Parts: [4]int{SplitRockBottom, MarbleheadTop, RoundIslandBottom, FortNiagaraTop}},
	{Id: 3, Parts: [4]int{MarbleheadTop, RoundIslandBottom, FortNiagaraTop, SplitRockBottom}},
	{Id: 2, Parts: [4]int{FortNiagaraBottom, MarbleheadTop, SplitRockTop, RoundIslandTop}},
	{Id: 1, Parts: [4]int{RoundIslandTop, FortNiagaraBottom, SplitRockBottom, MarbleheadBottom}},
	{Id: 9, Parts: [4]int{MarbleheadBottom, RoundIslandTop, FortNiagaraTop, SplitRockTop}},
}
