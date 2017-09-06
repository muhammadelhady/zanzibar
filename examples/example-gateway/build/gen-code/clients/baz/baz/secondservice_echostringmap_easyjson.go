// Code generated by zanzibar
// @generated
// Checksum : OMKR7rTMjPMqg32DcSfOQw==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package baz

import (
	json "encoding/json"
	fmt "fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	base "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonB0e5590DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStringMap(in *jlexer.Lexer, out *SecondService_EchoStringMap_Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Success = make(map[string]*base.BazResponse)
				} else {
					out.Success = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 *base.BazResponse
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(base.BazResponse)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					(out.Success)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB0e5590EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStringMap(out *jwriter.Writer, in SecondService_EchoStringMap_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"success\":")
	if in.Success == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v2First := true
		for v2Name, v2Value := range in.Success {
			if !v2First {
				out.RawByte(',')
			}
			v2First = false
			out.String(string(v2Name))
			out.RawByte(':')
			if v2Value == nil {
				out.RawString("null")
			} else {
				(*v2Value).MarshalEasyJSON(out)
			}
		}
		out.RawByte('}')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SecondService_EchoStringMap_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB0e5590EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStringMap(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SecondService_EchoStringMap_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB0e5590EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStringMap(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SecondService_EchoStringMap_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB0e5590DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStringMap(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SecondService_EchoStringMap_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB0e5590DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStringMap(l, v)
}
func easyjsonB0e5590DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStringMap1(in *jlexer.Lexer, out *SecondService_EchoStringMap_Args) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var ArgSet bool
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "arg":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Arg = make(map[string]*base.BazResponse)
				} else {
					out.Arg = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 *base.BazResponse
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(base.BazResponse)
						}
						(*v3).UnmarshalEasyJSON(in)
					}
					(out.Arg)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
			ArgSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !ArgSet {
		in.AddError(fmt.Errorf("key 'arg' is required"))
	}
}
func easyjsonB0e5590EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStringMap1(out *jwriter.Writer, in SecondService_EchoStringMap_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"arg\":")
	if in.Arg == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v4First := true
		for v4Name, v4Value := range in.Arg {
			if !v4First {
				out.RawByte(',')
			}
			v4First = false
			out.String(string(v4Name))
			out.RawByte(':')
			if v4Value == nil {
				out.RawString("null")
			} else {
				(*v4Value).MarshalEasyJSON(out)
			}
		}
		out.RawByte('}')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SecondService_EchoStringMap_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB0e5590EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStringMap1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SecondService_EchoStringMap_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB0e5590EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStringMap1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SecondService_EchoStringMap_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB0e5590DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStringMap1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SecondService_EchoStringMap_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB0e5590DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStringMap1(l, v)
}
