package lighthouse_solver

import (
	"fmt"
	"strings"
	"testing"
)

func TestRotateCard(t *testing.T) {
	card := Card{
		Id:    1,
		Parts: [4]int{SplitRockTop, RoundIslandTop, MarbleheadTop, FortNiagaraTop},
	}

	card.Rotate()

	if part := card.Parts[0]; part != FortNiagaraTop {
		t.Fatalf(`expected %v but got %v`, FortNiagaraTop, part)
	}
	if part := card.Parts[1]; part != SplitRockTop {
		t.Fatalf(`expected %v but got %v`, SplitRockTop, part)
	}
	if part := card.Parts[2]; part != RoundIslandTop {
		t.Fatalf(`expected %v but got %v`, RoundIslandTop, part)
	}
	if part := card.Parts[3]; part != MarbleheadTop {
		t.Fatalf(`expected %v but got %v`, MarbleheadTop, part)
	}
}

func TestMatchCard(t *testing.T) {
	card := &Card{
		Id:    1,
		Parts: [4]int{SplitRockTop, RoundIslandTop, MarbleheadTop, FortNiagaraTop},
	}
	cardToLeft := &Card{
		Id:    2,
		Parts: [4]int{SplitRockTop, SplitRockTop, SplitRockBottom, SplitRockTop},
	}
	cardToTop := &Card{
		Id:    3,
		Parts: [4]int{SplitRockTop, SplitRockTop, SplitRockTop, RoundIslandBottom},
	}

	if !card.MatchesRight(cardToLeft) {
		t.Fatalf(`expected card's left to match cardToLeft's right but it did not`)
	}
	if !card.MatchesBottom(cardToTop) {
		t.Fatalf(`expected card's top to match cardToTop's bottom but it did not`)
	}
	if !card.MatchesRightAndBottom(cardToLeft, cardToTop) {
		t.Fatalf(`expected card to match to cardToLeft and cardToTop but it did not`)
	}
}

func TestHasPart(t *testing.T) {
	card := &Card{
		Id:    1,
		Parts: [4]int{SplitRockTop, RoundIslandBottom, SplitRockTop, SplitRockTop},
	}

	if !card.HasPart(RoundIslandBottom) {
		t.Fatalf(`card reported it does not have RoundIslandBottom, but it does`)
	}
	if card.HasPart(MarbleheadTop) {
		t.Fatalf(`card reported it has MarbleheadTop, but it does not`)
	}
}

func TestPlaceSingleCard(t *testing.T) {
	solver := &Solver{Cards: []*Card{Cards[0]}}
	ok, solution := solver.PlaceCards()
	if !ok {
		t.Fatalf(`solver did not solve the solution`)
	}
	if actual := solution[0][0].Id; actual != 1 {
		t.Fatalf(`expected card 1 in top left but got %v`, actual)
	}
	logSolution(t, solution)
}

func TestPlaceTwoCards(t *testing.T) {
	solver := &Solver{Cards: []*Card{Cards[1], Cards[3]}}
	ok, solution := solver.PlaceCards()
	if !ok {
		t.Fatalf(`solver did not solve the solution`)
	}
	if actual := solution[0][0].Id; actual != 2 {
		t.Fatalf(`expected card 2 in top left but got %v`, actual)
	}
	if actual := solution[0][1].Id; actual != 4 {
		t.Fatalf(`expected card 4 in top middle but got %v`, actual)
	}
	logSolution(t, solution)
}

func TestPlaceTwoCardsThatRequireRotation(t *testing.T) {
	solver := &Solver{Cards: []*Card{Cards[1], Cards[7]}}
	ok, solution := solver.PlaceCards()
	if !ok {
		t.Fatalf(`solver did not solve the solution`)
	}
	if actual := solution[0][0].Id; actual != 2 {
		t.Fatalf(`expected card 2 in top left but got %v`, actual)
	}
	secondCard := solution[0][1]
	if secondCard.Id != 8 {
		t.Fatalf(`expected card 8 in top middle but got %v`, secondCard.Id)
	}
	if secondCard.Parts[0] != SplitRockBottom {
		t.Fatalf(`second card should be rotated to have SplitRockBottom on left, but it is not: %v`, secondCard.Parts)
	}
	logSolution(t, solution)
}

func logSolution(t *testing.T, solution [3][3]*Card) {
	for _, row := range solution {
		builder := strings.Builder{}
		for column, card := range row {
			if column > 0 {
				builder.WriteString(" ")
			}
			if card == nil {
				builder.WriteString("<nil>")
			} else {
				builder.WriteString(fmt.Sprintf(`{%v %v}`, card.Id, card.Parts))
			}
		}
		t.Logf(builder.String())
	}
}
