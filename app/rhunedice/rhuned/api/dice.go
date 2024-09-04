package api

import (
	"fmt"

	"github.com/jrecuero/thengine/pkg/tools"
)

type IDice interface {
	GetFaces() []IFace
	GetName() string
	Roll() IFace
	SetFaces([]IFace)
	SetName(string)
	String() string
}

type Dice struct {
	faces []IFace
	name  string
}

//func NewRandom(max int) int {
//    maxBig := big.NewInt(int64(max))
//    randomNumber, err := rand.Int(rand.Reader, maxBig)
//    if err != nil {
//        panic(err)
//    }
//    return int(randomNumber.Int64())
//}

func NewDice(name string, faces []IFace) *Dice {
	return &Dice{
		faces: faces,
		name:  name,
	}
}

func (d *Dice) GetFaces() []IFace {
	return d.faces
}

func (d *Dice) GetName() string {
	return d.name
}

func (d *Dice) Roll() IFace {
	nbrFaces := len(d.faces)
	index := tools.RandomRing.Intn(nbrFaces)
	//index := NewRandom(nbrFaces)
	return d.faces[index]
}

func (d *Dice) SetFaces(faces []IFace) {
	d.faces = faces
}

func (d *Dice) SetName(name string) {
	d.name = name
}

func (d *Dice) String() string {
	return fmt.Sprintf("%s", d.name)
}

var _ IDice = (*Dice)(nil)
