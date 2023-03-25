package service

import (
	"errors"
	interfaces "smhome/app/interface"
)

// Factory method Design Pattern
// param typ string.

func NewEntityContext(typ string) (interfaces.IEntity, error) {
	switch typ {
	case "user":
		return newUser(), nil
	case "action":
		return newAction(), nil
	case "sensors":
		return newSensors(), nil
	}
	return nil, errors.New("data type not in models")
}
