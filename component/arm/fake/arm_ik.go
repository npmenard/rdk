package fake

import (
	"context"

	// for arm model.
	_ "embed"

	"github.com/edaniels/golog"
	"github.com/pkg/errors"

	"go.viam.com/rdk/component/arm"
	"go.viam.com/rdk/config"
	"go.viam.com/rdk/motionplan"
	commonpb "go.viam.com/rdk/proto/api/common/v1"
	pb "go.viam.com/rdk/proto/api/component/v1"
	"go.viam.com/rdk/referenceframe"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/robot"
)

//go:embed arm_model.json
var armikModelJSON []byte

func init() {
	registry.RegisterComponent(arm.Subtype, "fake_ik", registry.Component{
		Constructor: func(ctx context.Context, r robot.Robot, config config.Component, logger golog.Logger) (interface{}, error) {
			if config.Attributes.Bool("fail_new", false) {
				return nil, errors.New("whoops")
			}
			return NewArmIK(ctx, config, logger)
		},
	})
}

// fakeModel returns the kinematics model.
func fakeModel() (referenceframe.Model, error) {
	return referenceframe.ParseJSON(armikModelJSON, "")
}

// NewArmIK returns a new fake arm.
func NewArmIK(ctx context.Context, cfg config.Component, logger golog.Logger) (arm.Arm, error) {
	name := cfg.Name
	model, err := fakeModel()
	if err != nil {
		return nil, err
	}
	mp, err := motionplan.NewCBiRRTMotionPlanner(model, 4, logger)
	if err != nil {
		return nil, err
	}

	return &ArmIK{
		Name:     name,
		position: &commonpb.Pose{},
		joints:   &pb.ArmJointPositions{Degrees: []float64{0, 0, 0, 0, 0, 0}},
		mp:       mp,
		model:    model,
	}, nil
}

// ArmIK is a fake arm that can simply read and set properties.
type ArmIK struct {
	Name       string
	position   *commonpb.Pose
	joints     *pb.ArmJointPositions
	mp         motionplan.MotionPlanner
	CloseCount int
	model      referenceframe.Model
}

// ModelFrame returns the dynamic frame of the model.
func (a *ArmIK) ModelFrame() referenceframe.Model {
	return a.model
}

// CurrentPosition returns the set position.
func (a *ArmIK) CurrentPosition(ctx context.Context) (*commonpb.Pose, error) {
	joints, err := a.CurrentJointPositions(ctx)
	if err != nil {
		return nil, err
	}
	return motionplan.ComputePosition(a.mp.Frame(), joints)
}

// MoveToPosition sets the position.
func (a *ArmIK) MoveToPosition(ctx context.Context, pos *commonpb.Pose) error {
	joints, err := a.CurrentJointPositions(ctx)
	if err != nil {
		return err
	}
	solution, err := a.mp.Plan(ctx, pos, referenceframe.JointPosToInputs(joints), nil)
	if err != nil {
		return err
	}
	return arm.GoToWaypoints(ctx, a, solution)
}

// MoveToJointPositions sets the joints.
func (a *ArmIK) MoveToJointPositions(ctx context.Context, joints *pb.ArmJointPositions) error {
	a.joints = joints
	return nil
}

// CurrentJointPositions returns the set joints.
func (a *ArmIK) CurrentJointPositions(ctx context.Context) (*pb.ArmJointPositions, error) {
	return a.joints, nil
}

// JointMoveDelta returns an error.
func (a *ArmIK) JointMoveDelta(ctx context.Context, joint int, amountDegs float64) error {
	return errors.New("arm JointMoveDelta does nothing")
}

// CurrentInputs TODO.
func (a *ArmIK) CurrentInputs(ctx context.Context) ([]referenceframe.Input, error) {
	res, err := a.CurrentJointPositions(ctx)
	if err != nil {
		return nil, err
	}
	return referenceframe.JointPosToInputs(res), nil
}

// GoToInputs TODO.
func (a *ArmIK) GoToInputs(ctx context.Context, goal []referenceframe.Input) error {
	return a.MoveToJointPositions(ctx, referenceframe.InputsToJointPos(goal))
}

// Close does nothing.
func (a *ArmIK) Close() {
	a.CloseCount++
}