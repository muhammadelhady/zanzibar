// Code generated by thriftrw v1.6.0. DO NOT EDIT.
// @generated

package baz

import (
	"errors"
	"fmt"
	"github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type SecondService_EchoStructSet_Args struct {
	Arg []*base.BazResponse `json:"arg,required"`
}

type _Set_BazResponse_ValueList []*base.BazResponse

func (v _Set_BazResponse_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		if x == nil {
			return fmt.Errorf("invalid set item: value is nil")
		}
		w, err := x.ToWire()
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _Set_BazResponse_ValueList) Size() int {
	return len(v)
}

func (_Set_BazResponse_ValueList) ValueType() wire.Type {
	return wire.TStruct
}

func (_Set_BazResponse_ValueList) Close() {
}

func (v *SecondService_EchoStructSet_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Arg == nil {
		return w, errors.New("field Arg of SecondService_EchoStructSet_Args is required")
	}
	w, err = wire.NewValueSet(_Set_BazResponse_ValueList(v.Arg)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Set_BazResponse_Read(s wire.ValueList) ([]*base.BazResponse, error) {
	if s.ValueType() != wire.TStruct {
		return nil, nil
	}
	o := make([]*base.BazResponse, 0, s.Size())
	err := s.ForEach(func(x wire.Value) error {
		i, err := _BazResponse_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	s.Close()
	return o, err
}

func (v *SecondService_EchoStructSet_Args) FromWire(w wire.Value) error {
	var err error
	argIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TSet {
				v.Arg, err = _Set_BazResponse_Read(field.Value.GetSet())
				if err != nil {
					return err
				}
				argIsSet = true
			}
		}
	}
	if !argIsSet {
		return errors.New("field Arg of SecondService_EchoStructSet_Args is required")
	}
	return nil
}

func (v *SecondService_EchoStructSet_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Arg: %v", v.Arg)
	i++
	return fmt.Sprintf("SecondService_EchoStructSet_Args{%v}", strings.Join(fields[:i], ", "))
}

func _Set_BazResponse_Equals(lhs, rhs []*base.BazResponse) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for _, x := range lhs {
		ok := false
		for _, y := range rhs {
			if x.Equals(y) {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}
	return true
}

func (v *SecondService_EchoStructSet_Args) Equals(rhs *SecondService_EchoStructSet_Args) bool {
	if !_Set_BazResponse_Equals(v.Arg, rhs.Arg) {
		return false
	}
	return true
}

func (v *SecondService_EchoStructSet_Args) MethodName() string {
	return "echoStructSet"
}

func (v *SecondService_EchoStructSet_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var SecondService_EchoStructSet_Helper = struct {
	Args           func(arg []*base.BazResponse) *SecondService_EchoStructSet_Args
	IsException    func(error) bool
	WrapResponse   func([]*base.BazResponse, error) (*SecondService_EchoStructSet_Result, error)
	UnwrapResponse func(*SecondService_EchoStructSet_Result) ([]*base.BazResponse, error)
}{}

func init() {
	SecondService_EchoStructSet_Helper.Args = func(arg []*base.BazResponse) *SecondService_EchoStructSet_Args {
		return &SecondService_EchoStructSet_Args{Arg: arg}
	}
	SecondService_EchoStructSet_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	SecondService_EchoStructSet_Helper.WrapResponse = func(success []*base.BazResponse, err error) (*SecondService_EchoStructSet_Result, error) {
		if err == nil {
			return &SecondService_EchoStructSet_Result{Success: success}, nil
		}
		return nil, err
	}
	SecondService_EchoStructSet_Helper.UnwrapResponse = func(result *SecondService_EchoStructSet_Result) (success []*base.BazResponse, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type SecondService_EchoStructSet_Result struct {
	Success []*base.BazResponse `json:"success"`
}

func (v *SecondService_EchoStructSet_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueSet(_Set_BazResponse_ValueList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("SecondService_EchoStructSet_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SecondService_EchoStructSet_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TSet {
				v.Success, err = _Set_BazResponse_Read(field.Value.GetSet())
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("SecondService_EchoStructSet_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *SecondService_EchoStructSet_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("SecondService_EchoStructSet_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *SecondService_EchoStructSet_Result) Equals(rhs *SecondService_EchoStructSet_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && _Set_BazResponse_Equals(v.Success, rhs.Success))) {
		return false
	}
	return true
}

func (v *SecondService_EchoStructSet_Result) MethodName() string {
	return "echoStructSet"
}

func (v *SecondService_EchoStructSet_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
