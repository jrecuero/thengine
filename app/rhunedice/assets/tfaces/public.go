package tfaces

import (
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/faces"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/rhunes"
	API "github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/constants"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

var (
	AsciiFaceSize = API.NewSize(5, 5)
)

var (
	AsciiAttack  = "  |  \n  |  \n  |  \n -x- \n  |  \n"
	AsciiDefense = "x   x\n x x \n  x  \n x x \nx   x\n"
	AsciiSkill   = "k  kk\nk kk \nkkk  \nk kk \nk  kk\n"
	AsciiStep    = "sssss\ns    \nsssss\n    s\nsssss\n"
	AsciiHealth  = "h   h\nh   h\nhhhhh\nh   h\nh   h\n"
	AsciiNil     = "0000h\n0   0\n0   0\n0   0\n00000\n"
)

var (
	AsciiCanvasAttack  = engine.NewCanvasFromString(AsciiAttack, &constants.WhiteOverBlack)
	AsciiCanvasDefense = engine.NewCanvasFromString(AsciiDefense, &constants.WhiteOverBlack)
	AsciiCanvasSkill   = engine.NewCanvasFromString(AsciiSkill, &constants.WhiteOverBlack)
	AsciiCanvasStep    = engine.NewCanvasFromString(AsciiStep, &constants.WhiteOverBlack)
	AsciiCanvasHealth  = engine.NewCanvasFromString(AsciiHealth, &constants.WhiteOverBlack)
	AsciiCanvasNil     = engine.NewCanvasFromString(AsciiNil, &constants.WhiteOverBlack)

	AsciiCanvasAllFaces = []*engine.Canvas{
		AsciiCanvasAttack,
		AsciiCanvasDefense,
		AsciiCanvasSkill,
		AsciiCanvasStep,
		AsciiCanvasHealth,
		AsciiCanvasNil,
	}
)

type RhuneFrame struct {
	*widgets.Frame
	rhune *api.Rhune
}

func NewRhuneFrame(rhune *api.Rhune, canvas *engine.Canvas, maxTicks int) *RhuneFrame {
	return &RhuneFrame{
		Frame: widgets.NewFrameWithCanvas(canvas, maxTicks),
		rhune: rhune,
	}
}

func (r *RhuneFrame) GetRhune() *api.Rhune {
	return r.rhune
}

const (
	RhuneFrameMaxTicks int = 5
)

var (
	AttackRhuneFrame  *RhuneFrame = NewRhuneFrame(rhunes.AttackRhune, AsciiCanvasAttack, RhuneFrameMaxTicks)
	DefenseRhuneFrame *RhuneFrame = NewRhuneFrame(rhunes.DefenseRhune, AsciiCanvasDefense, RhuneFrameMaxTicks)
	SkillRhuneFrame   *RhuneFrame = NewRhuneFrame(rhunes.SkillRhune, AsciiCanvasSkill, RhuneFrameMaxTicks)
	StepRhuneFrame    *RhuneFrame = NewRhuneFrame(rhunes.StepRhune, AsciiCanvasStep, RhuneFrameMaxTicks)
	HealthRhuneFrame  *RhuneFrame = NewRhuneFrame(rhunes.HealthRhune, AsciiCanvasHealth, RhuneFrameMaxTicks)
	NilRhuneFrame     *RhuneFrame = NewRhuneFrame(rhunes.NilRhune, AsciiCanvasNil, RhuneFrameMaxTicks)

	BaseRhuneFrames []widgets.IFrame = []widgets.IFrame{
		AttackRhuneFrame,
		DefenseRhuneFrame,
		SkillRhuneFrame,
		StepRhuneFrame,
		HealthRhuneFrame,
		NilRhuneFrame,
	}
)

func NewBaseRhuneFames() []widgets.IFrame {
	return BaseRhuneFrames
}

func NewAsciiFramesFromAllFaces(ticks int) []widgets.IFrame {
	var frames []widgets.IFrame = make([]widgets.IFrame, len(AsciiCanvasAllFaces))
	for i, face := range AsciiCanvasAllFaces {
		frames[i] = widgets.NewFrameWithCanvas(face, ticks)
	}
	return frames
}

// GetFaceFromTFace function returns a string with the name of the rhune being
// represented in the given canvas/face.
func GetFaceFromTFace(face *engine.Canvas) string {
	switch face {
	case AsciiCanvasAttack:
		return api.AttackName
	case AsciiCanvasDefense:
		return api.DefenseName
	case AsciiCanvasSkill:
		return api.SkillName
	case AsciiCanvasStep:
		return api.StepName
	case AsciiCanvasHealth:
		return api.HealthName
	case AsciiCanvasNil:
		return api.NilName
	default:
		return api.NilName
	}
}

type Face struct {
	face   api.IFace
	canvas *engine.Canvas
}

func NewFace(face api.IFace, canvas *engine.Canvas) *Face {
	return &Face{
		face:   face,
		canvas: canvas,
	}
}

var (
	AttackFace  = NewFace(faces.AttackFace, AsciiCanvasAttack)
	DefenseFace = NewFace(faces.DefenseFace, AsciiCanvasDefense)
	SkillFace   = NewFace(faces.SkillFace, AsciiCanvasSkill)
	StepFace    = NewFace(faces.StepFace, AsciiCanvasStep)
	HealthFace  = NewFace(faces.HealthFace, AsciiCanvasHealth)
	NilFace     = NewFace(faces.NilFace, AsciiCanvasNil)

	BaseFaces = []*Face{
		AttackFace,
		DefenseFace,
		SkillFace,
		StepFace,
		HealthFace,
		NilFace,
	}
)
