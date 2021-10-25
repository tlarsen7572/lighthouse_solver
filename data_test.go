package lighthouse_solver_test

import (
	"lighthouse_solver"
	"testing"
)

func TestRotateCard(t *testing.T) {
	card := lighthouse_solver.Card{
		Id:    1,
		Parts: [4]int{lighthouse_solver.SplitRockTop, lighthouse_solver.RoundIslandTop, lighthouse_solver.MarbleheadTop, lighthouse_solver.FortNiagaraTop},
	}

	card.Rotate()

	if part := card.Parts[0]; part != lighthouse_solver.FortNiagaraTop {
		t.Fatalf(`expected %v but got %v`, lighthouse_solver.FortNiagaraTop, part)
	}
	if part := card.Parts[1]; part != lighthouse_solver.SplitRockTop {
		t.Fatalf(`expected %v but got %v`, lighthouse_solver.SplitRockTop, part)
	}
	if part := card.Parts[2]; part != lighthouse_solver.RoundIslandTop {
		t.Fatalf(`expected %v but got %v`, lighthouse_solver.RoundIslandTop, part)
	}
	if part := card.Parts[3]; part != lighthouse_solver.MarbleheadTop {
		t.Fatalf(`expected %v but got %v`, lighthouse_solver.MarbleheadTop, part)
	}
}
