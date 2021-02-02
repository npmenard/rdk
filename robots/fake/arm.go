package fake

import (
	"fmt"

	"github.com/viamrobotics/robotcore/arm"
)

type Arm struct {
	position arm.Position
}

func (a *Arm) Close() {
}

func (a *Arm) CurrentPosition() (arm.Position, error) {
	return a.position, nil
}

func (a *Arm) MoveToPosition(c arm.Position) error {
	a.position = c
	return nil
}

func (a *Arm) MoveToJointPositions(joints []float64) error {
	return fmt.Errorf("arm MoveToJointPositions doesn't work")
}

func (a *Arm) CurrentJointPositions() ([]float64, error) {
	return nil, fmt.Errorf("arm CurrentJointPositions doesn't work")
}

func (a *Arm) JointMoveDelta(joint int, amount float64) error {
	return fmt.Errorf("arm JointMoveDelta does nothing")
}
