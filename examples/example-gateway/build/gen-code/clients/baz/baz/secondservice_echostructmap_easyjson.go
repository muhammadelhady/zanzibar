// Code generated by zanzibar
// @generated
// Checksum : lpOnr1Pz+9QRqV0i7Oo22A==
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

func easyjson899419cDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStructMap(in *jlexer.Lexer, out *SecondService_EchoStructMap_Result) {
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
				out.Success = nil
			} else {
				in.Delim('[')
				if out.Success == nil {
					if !in.IsDelim(']') {
						out.Success = make([]struct {
							Key   *base.BazResponse
							Value string
						}, 0, 2)
					} else {
						out.Success = []struct {
							Key   *base.BazResponse
							Value string
						}{}
					}
				} else {
					out.Success = (out.Success)[:0]
				}
				for !in.IsDelim(']') {
					var v1 struct {
						Key   *base.BazResponse
						Value string
					}
					easyjson899419cDecode(in, &v1)
					out.Success = append(out.Success, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson899419cEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStructMap(out *jwriter.Writer, in SecondService_EchoStructMap_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"success\":")
	if in.Success == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in.Success {
			if v2 > 0 {
				out.RawByte(',')
			}
			easyjson899419cEncode(out, v3)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SecondService_EchoStructMap_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson899419cEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStructMap(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SecondService_EchoStructMap_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson899419cEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStructMap(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SecondService_EchoStructMap_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson899419cDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStructMap(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SecondService_EchoStructMap_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson899419cDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStructMap(l, v)
}
func easyjson899419cDecode(in *jlexer.Lexer, out *struct {
	Key   *base.BazResponse
	Value string
}) {
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
		case "Key":
			if in.IsNull() {
				in.Skip()
				out.Key = nil
			} else {
				if out.Key == nil {
					out.Key = new(base.BazResponse)
				}
				(*out.Key).UnmarshalEasyJSON(in)
			}
		case "Value":
			out.Value = string(in.String())
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
func easyjson899419cEncode(out *jwriter.Writer, in struct {
	Key   *base.BazResponse
	Value string
}) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Key\":")
	if in.Key == nil {
		out.RawString("null")
	} else {
		(*in.Key).MarshalEasyJSON(out)
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Value\":")
	out.String(string(in.Value))
	out.RawByte('}')
}
func easyjson899419cDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStructMap1(in *jlexer.Lexer, out *SecondService_EchoStructMap_Args) {
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
				out.Arg = nil
			} else {
				in.Delim('[')
				if out.Arg == nil {
					if !in.IsDelim(']') {
						out.Arg = make([]struct {
							Key   *base.BazResponse
							Value string
						}, 0, 2)
					} else {
						out.Arg = []struct {
							Key   *base.BazResponse
							Value string
						}{}
					}
				} else {
					out.Arg = (out.Arg)[:0]
				}
				for !in.IsDelim(']') {
					var v4 struct {
						Key   *base.BazResponse
						Value string
					}
					easyjson899419cDecode(in, &v4)
					out.Arg = append(out.Arg, v4)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson899419cEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStructMap1(out *jwriter.Writer, in SecondService_EchoStructMap_Args) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"arg\":")
	if in.Arg == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in.Arg {
			if v5 > 0 {
				out.RawByte(',')
			}
			easyjson899419cEncode(out, v6)
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SecondService_EchoStructMap_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson899419cEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStructMap1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SecondService_EchoStructMap_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson899419cEncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStructMap1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SecondService_EchoStructMap_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson899419cDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStructMap1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SecondService_EchoStructMap_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson899419cDecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSecondServiceEchoStructMap1(l, v)
}
