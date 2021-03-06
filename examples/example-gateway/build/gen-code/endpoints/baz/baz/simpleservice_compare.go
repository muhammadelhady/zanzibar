// Code generated by thriftrw v1.6.0. DO NOT EDIT.
// @generated

package baz

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type SimpleService_Compare_Args struct {
	Arg1 *BazRequest `json:"arg1,required"`
	Arg2 *BazRequest `json:"arg2,required"`
}

func (v *SimpleService_Compare_Args) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Arg1 == nil {
		return w, errors.New("field Arg1 of SimpleService_Compare_Args is required")
	}
	w, err = v.Arg1.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.Arg2 == nil {
		return w, errors.New("field Arg2 of SimpleService_Compare_Args is required")
	}
	w, err = v.Arg2.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SimpleService_Compare_Args) FromWire(w wire.Value) error {
	var err error
	arg1IsSet := false
	arg2IsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Arg1, err = _BazRequest_Read(field.Value)
				if err != nil {
					return err
				}
				arg1IsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.Arg2, err = _BazRequest_Read(field.Value)
				if err != nil {
					return err
				}
				arg2IsSet = true
			}
		}
	}
	if !arg1IsSet {
		return errors.New("field Arg1 of SimpleService_Compare_Args is required")
	}
	if !arg2IsSet {
		return errors.New("field Arg2 of SimpleService_Compare_Args is required")
	}
	return nil
}

func (v *SimpleService_Compare_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Arg1: %v", v.Arg1)
	i++
	fields[i] = fmt.Sprintf("Arg2: %v", v.Arg2)
	i++
	return fmt.Sprintf("SimpleService_Compare_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *SimpleService_Compare_Args) Equals(rhs *SimpleService_Compare_Args) bool {
	if !v.Arg1.Equals(rhs.Arg1) {
		return false
	}
	if !v.Arg2.Equals(rhs.Arg2) {
		return false
	}
	return true
}

func (v *SimpleService_Compare_Args) MethodName() string {
	return "compare"
}

func (v *SimpleService_Compare_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var SimpleService_Compare_Helper = struct {
	Args           func(arg1 *BazRequest, arg2 *BazRequest) *SimpleService_Compare_Args
	IsException    func(error) bool
	WrapResponse   func(*BazResponse, error) (*SimpleService_Compare_Result, error)
	UnwrapResponse func(*SimpleService_Compare_Result) (*BazResponse, error)
}{}

func init() {
	SimpleService_Compare_Helper.Args = func(arg1 *BazRequest, arg2 *BazRequest) *SimpleService_Compare_Args {
		return &SimpleService_Compare_Args{Arg1: arg1, Arg2: arg2}
	}
	SimpleService_Compare_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *AuthErr:
			return true
		case *OtherAuthErr:
			return true
		default:
			return false
		}
	}
	SimpleService_Compare_Helper.WrapResponse = func(success *BazResponse, err error) (*SimpleService_Compare_Result, error) {
		if err == nil {
			return &SimpleService_Compare_Result{Success: success}, nil
		}
		switch e := err.(type) {
		case *AuthErr:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for SimpleService_Compare_Result.AuthErr")
			}
			return &SimpleService_Compare_Result{AuthErr: e}, nil
		case *OtherAuthErr:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for SimpleService_Compare_Result.OtherAuthErr")
			}
			return &SimpleService_Compare_Result{OtherAuthErr: e}, nil
		}
		return nil, err
	}
	SimpleService_Compare_Helper.UnwrapResponse = func(result *SimpleService_Compare_Result) (success *BazResponse, err error) {
		if result.AuthErr != nil {
			err = result.AuthErr
			return
		}
		if result.OtherAuthErr != nil {
			err = result.OtherAuthErr
			return
		}
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type SimpleService_Compare_Result struct {
	Success      *BazResponse  `json:"success,omitempty"`
	AuthErr      *AuthErr      `json:"authErr,omitempty"`
	OtherAuthErr *OtherAuthErr `json:"otherAuthErr,omitempty"`
}

func (v *SimpleService_Compare_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.AuthErr != nil {
		w, err = v.AuthErr.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.OtherAuthErr != nil {
		w, err = v.OtherAuthErr.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("SimpleService_Compare_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _BazResponse_Read(w wire.Value) (*BazResponse, error) {
	var v BazResponse
	err := v.FromWire(w)
	return &v, err
}

func _OtherAuthErr_Read(w wire.Value) (*OtherAuthErr, error) {
	var v OtherAuthErr
	err := v.FromWire(w)
	return &v, err
}

func (v *SimpleService_Compare_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _BazResponse_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.AuthErr, err = _AuthErr_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.OtherAuthErr, err = _OtherAuthErr_Read(field.Value)
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
	if v.AuthErr != nil {
		count++
	}
	if v.OtherAuthErr != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("SimpleService_Compare_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *SimpleService_Compare_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [3]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.AuthErr != nil {
		fields[i] = fmt.Sprintf("AuthErr: %v", v.AuthErr)
		i++
	}
	if v.OtherAuthErr != nil {
		fields[i] = fmt.Sprintf("OtherAuthErr: %v", v.OtherAuthErr)
		i++
	}
	return fmt.Sprintf("SimpleService_Compare_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *SimpleService_Compare_Result) Equals(rhs *SimpleService_Compare_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	if !((v.AuthErr == nil && rhs.AuthErr == nil) || (v.AuthErr != nil && rhs.AuthErr != nil && v.AuthErr.Equals(rhs.AuthErr))) {
		return false
	}
	if !((v.OtherAuthErr == nil && rhs.OtherAuthErr == nil) || (v.OtherAuthErr != nil && rhs.OtherAuthErr != nil && v.OtherAuthErr.Equals(rhs.OtherAuthErr))) {
		return false
	}
	return true
}

func (v *SimpleService_Compare_Result) MethodName() string {
	return "compare"
}

func (v *SimpleService_Compare_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
