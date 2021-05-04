package main

import (
	"testing"

	"go.viam.com/test"
	"github.com/tonyOreglia/glee/pkg/position"
)

func TestEngine1(t *testing.T) {
	p, err := position.NewPositionFen("r1bqkbnr/pppppppp/2n5/4P3/8/8/PPPP1PPP/RNBQKBNR b KQkq - 2 1")
	test.That(t, err, test.ShouldBeNil)

	_, m := searchForNextMove(p)
	test.That(t, m.String(), test.ShouldNotEqual, "c7c6")
}
