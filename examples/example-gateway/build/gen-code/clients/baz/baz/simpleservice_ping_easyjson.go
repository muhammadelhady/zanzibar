// Code generated by zanzibar
// @generated
// Checksum : vuwjJMjGibuq+uQeqHAT1Q==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package baz

import (
	json "encoding/json"
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

func easyjson730cf796DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServicePing(in *jlexer.Lexer, out *SimpleService_Ping_Result) {
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
				if out.Success == nil {
					out.Success = new(base.BazResponse)
				}
				(*out.Success).UnmarshalEasyJSON(in)
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
func easyjson730cf796EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServicePing(out *jwriter.Writer, in SimpleService_Ping_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"success\":")
		if in.Success == nil {
			out.RawString("null")
		} else {
			(*in.Success).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_Ping_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson730cf796EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServicePing(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_Ping_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson730cf796EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServicePing(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_Ping_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson730cf796DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServicePing(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_Ping_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson730cf796DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServicePing(l, v)
}
func easyjson730cf796DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServicePing1(in *jlexer.Lexer, out *SimpleService_Ping_Args) {
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
func easyjson730cf796EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServicePing1(out *jwriter.Writer, in SimpleService_Ping_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_Ping_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson730cf796EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServicePing1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_Ping_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson730cf796EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServicePing1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_Ping_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson730cf796DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServicePing1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_Ping_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson730cf796DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeClientsBazBazSimpleServicePing1(l, v)
}
