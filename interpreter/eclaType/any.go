package eclaType

import (
	"errors"
	"github.com/Eclalang/Ecla/parser"
)

type Any struct {
	Name  string
	Value Type
}

func (a *Any) String() string {
	return string(a.Value.GetString())
}

func (a *Any) GetString() String {
	return a.Value.GetString()
}

func (a *Any) GetValue() any {
	return a.Value.GetValue()
}

func (a *Any) SetValue(value any) error {
	return a.Value.SetValue(value)
}

func (a *Any) GetType() string {
	return a.Value.GetType()
}

func (a *Any) GetIndex(i Type) (*Type, error) {
	return a.Value.GetIndex(i)
}

// SetAny sets the value of the variable
func (a *Any) SetAny(value Type) error {
	switch value.(type) {
	case *Var:
		a.Value = value.(*Any).Value
		return nil
	}
	typ2 := a.Value.GetType()
	if value.IsNull() {
		a.Value = NewNullType(typ2)
		return nil
	}
	a.Value = value
	return nil
}

// Add adds two Type objects
func (a *Any) Add(other Type) (Type, error) {
	return a.Value.Add(other)
}

// Sub subtracts two Type objects
func (a *Any) Sub(other Type) (Type, error) {
	return a.Value.Sub(other)
}

// Mul multiplies two Type objects
func (a *Any) Mul(other Type) (Type, error) {
	return a.Value.Mul(other)
}

// Div divides two Type objects
func (a *Any) Div(other Type) (Type, error) {
	return a.Value.Div(other)
}

// Mod modulos two Type objects
func (a *Any) Mod(other Type) (Type, error) {
	return a.Value.Mod(other)
}

// DivEc divides two Type objects
func (a *Any) DivEc(other Type) (Type, error) {
	return a.Value.DivEc(other)
}

// Eq returns true if the two Type objects are equal
func (a *Any) Eq(other Type) (Type, error) {
	return a.Value.Eq(other)
}

// NotEq returns true if the two Type objects are not equal
func (a *Any) NotEq(other Type) (Type, error) {
	return a.Value.NotEq(other)
}

// Gt returns true if the first Type object is greater than the second
func (a *Any) Gt(other Type) (Type, error) {
	return a.Value.Gt(other)
}

// GtEq returns true if the first Type object is greater than or equal to the second
func (a *Any) GtEq(other Type) (Type, error) {
	return a.Value.GtEq(other)
}

// Lw returns true if the first Type object is lower than the second
func (a *Any) Lw(other Type) (Type, error) {
	return a.Value.Lw(other)
}

// LwEq returns true if the first Type object is lower than or equal to the second
func (a *Any) LwEq(other Type) (Type, error) {
	return a.Value.LwEq(other)
}

// And returns true if the two Type objects are true
func (a *Any) And(other Type) (Type, error) {
	return a.Value.And(other)
}

// Or returns true if either Type objects is true
func (a *Any) Or(other Type) (Type, error) {
	return a.Value.Or(other)
}

// Not returns the opposite of the Type object
func (a *Any) Not() (Type, error) {
	return a.Value.Not()
}

func (a *Any) Decrement() {
	var err error
	a.Value, err = a.Value.Sub(NewInt("1"))
	if err != nil {
		panic(err)
	}
}

func (a *Any) Increment() {
	var err error
	a.Value, err = a.Value.Add(NewInt("1"))
	if err != nil {
		panic(err)
	}
}

func (a *Any) Append(other Type) (Type, error) {
	return a.Value.Append(other)
}

func (a *Any) IsNull() bool {
	return a.Value.IsNull()
}

func (a *Any) IsFunction() bool {
	return a.Value.GetType() == "function"
}

func (a *Any) GetFunction() *Function {
	switch a.Value.(type) {
	case *Function:
		return a.Value.(*Function)
	}
	return nil
}

// NewAny creates a new variable
func NewAny(name string, Type string, value Type) (*Any, error) {
	if Type == parser.String {
		return &Any{
			Name:  name,
			Value: value.GetString(),
		}, nil
	}
	if Type == parser.Float && value.GetType() == parser.Int {
		return &Any{
			Name:  name,
			Value: NewFloat(value.String()),
		}, nil

	}

	if Type == "" {
		Type = value.GetType()
	} else if Type != value.GetType() && !value.IsNull() {
		return nil, errors.New("cannot create variable of type " + Type + " with value of type " + value.GetType())
	}
	if value.IsNull() {
		value = NewNullType(Type)
	}

	switch value.(type) {
	case *Var:
		value = value.(*Var).Value
	}
	return &Any{Name: name, Value: value}, nil
}

func NewAnyEmpty(name string, Type string) (*Any, error) {
	return &Any{Name: name, Value: NewNullType(Type)}, nil
}
