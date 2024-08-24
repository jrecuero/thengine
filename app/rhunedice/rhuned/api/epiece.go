package api

import "fmt"

type IEquipmentPiece interface {
	GetBucketName() string
	GetDescription() string
	GetName() string
	GetValue() int
	String() string
}

type EquipmentPiece struct {
	bucketName  string
	description string
	name        string
	value       int
}

func NewEquipmentPiece(name string, description string,
	bucketName string, value int) *EquipmentPiece {
	return &EquipmentPiece{
		bucketName:  bucketName,
		description: description,
		name:        name,
		value:       value,
	}
}

func (e *EquipmentPiece) GetBucketName() string {
	return e.bucketName
}

func (e *EquipmentPiece) GetDescription() string {
	return e.description
}

func (e *EquipmentPiece) GetName() string {
	return e.name
}

func (e *EquipmentPiece) GetValue() int {
	return e.value
}

func (e *EquipmentPiece) String() string {
	return fmt.Sprintf("%s %s %d", e.name, e.bucketName, e.value)
}

var _ IEquipmentPiece = (*EquipmentPiece)(nil)
