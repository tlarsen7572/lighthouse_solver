package lighthouse_solver

const SplitRockTop = 1
const SplitRockBottom = -1
const RoundIslandTop = 10
const RoundIslandBottom = -10
const MarbleheadTop = 100
const MarbleheadBottom = -100
const FortNiagaraTop = 1000
const FortNiagaraBottom = -1000

const left = 0
const top = 1
const right = 2
const bottom = 3

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

func (c *Card) MatchesRight(other *Card) bool {
	return c.Parts[left]+other.Parts[right] == 0
}

func (c *Card) MatchesBottom(other *Card) bool {
	return c.Parts[top]+other.Parts[bottom] == 0
}

func (c *Card) MatchesRightAndBottom(toLeft *Card, toTop *Card) bool {
	return c.Parts[top]+c.Parts[left]+toLeft.Parts[right]+toTop.Parts[bottom] == 0
}

func (c *Card) HasPart(part int) bool {
	for _, value := range c.Parts {
		if value == part {
			return true
		}
	}
	return false
}

type Solver struct {
	Cards []*Card
}

func (s *Solver) PlaceCards() (bool, [3][3]*Card) {
	solution := [3][3]*Card{}
	for index, card := range s.Cards {
		if index == 0 {
			solution[0][index] = card
			continue
		}
		priorCard := solution[0][index-1]
		if card.MatchesRight(priorCard) {
			solution[0][index] = card
			continue
		}
		if s.tryRotateAndMatch(card, priorCard) {
			solution[0][index] = card
			continue
		}
		return false, solution
	}
	return true, solution
}

func (s *Solver) tryRotateAndMatch(card *Card, priorCard *Card) bool {
	rotations := 0
	for rotations < 4 {
		card.Rotate()
		if card.MatchesRight(priorCard) {
			return true
		}
		rotations++
	}
	return false
}

var Cards = []*Card{
	{Id: 1, Parts: [4]int{RoundIslandTop, FortNiagaraBottom, SplitRockBottom, MarbleheadBottom}},
	{Id: 2, Parts: [4]int{FortNiagaraBottom, MarbleheadTop, SplitRockTop, RoundIslandTop}},
	{Id: 3, Parts: [4]int{MarbleheadTop, RoundIslandBottom, FortNiagaraTop, SplitRockBottom}},
	{Id: 4, Parts: [4]int{SplitRockBottom, MarbleheadTop, RoundIslandBottom, FortNiagaraTop}},
	{Id: 5, Parts: [4]int{RoundIslandTop, SplitRockTop, MarbleheadBottom, SplitRockTop}},
	{Id: 6, Parts: [4]int{MarbleheadTop, RoundIslandTop, MarbleheadBottom, FortNiagaraTop}},
	{Id: 7, Parts: [4]int{MarbleheadTop, SplitRockBottom, RoundIslandBottom, FortNiagaraTop}},
	{Id: 8, Parts: [4]int{RoundIslandTop, SplitRockBottom, FortNiagaraTop, MarbleheadTop}},
	{Id: 9, Parts: [4]int{MarbleheadBottom, RoundIslandTop, FortNiagaraTop, SplitRockTop}},
}
