package lighthouse_solver

import (
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
