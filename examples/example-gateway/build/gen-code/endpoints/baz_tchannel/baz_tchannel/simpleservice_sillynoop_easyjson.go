// Code generated by zanzibar
// @generated
// Checksum : rUNO5U0P0SgjAVWElklGeA==
// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package baz_tchannel

import (
	json "encoding/json"
	fmt "fmt"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannelSimpleServiceSillyNoop(in *jlexer.Lexer, out *SimpleService_SillyNoop_Result) {
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
		case "authErr":
			if in.IsNull() {
				in.Skip()
				out.AuthErr = nil
			} else {
				if out.AuthErr == nil {
					out.AuthErr = new(AuthErr)
				}
				easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannel(in, &*out.AuthErr)
			}
		case "serverErr":
			if in.IsNull() {
				in.Skip()
				out.ServerErr = nil
			} else {
				if out.ServerErr == nil {
					out.ServerErr = new(ServerErr)
				}
				easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannel1(in, &*out.ServerErr)
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
func easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannelSimpleServiceSillyNoop(out *jwriter.Writer, in SimpleService_SillyNoop_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.AuthErr != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"authErr\":")
		if in.AuthErr == nil {
			out.RawString("null")
		} else {
			easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannel(out, *in.AuthErr)
		}
	}
	if in.ServerErr != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"serverErr\":")
		if in.ServerErr == nil {
			out.RawString("null")
		} else {
			easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannel1(out, *in.ServerErr)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_SillyNoop_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannelSimpleServiceSillyNoop(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_SillyNoop_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannelSimpleServiceSillyNoop(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_SillyNoop_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannelSimpleServiceSillyNoop(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_SillyNoop_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannelSimpleServiceSillyNoop(l, v)
}
func easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannel1(in *jlexer.Lexer, out *ServerErr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var MessageSet bool
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
		case "message":
			out.Message = string(in.String())
			MessageSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !MessageSet {
		in.AddError(fmt.Errorf("key 'message' is required"))
	}
}
func easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannel1(out *jwriter.Writer, in ServerErr) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"message\":")
	out.String(string(in.Message))
	out.RawByte('}')
}
func easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannel(in *jlexer.Lexer, out *AuthErr) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	var MessageSet bool
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
		case "message":
			out.Message = string(in.String())
			MessageSet = true
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
	if !MessageSet {
		in.AddError(fmt.Errorf("key 'message' is required"))
	}
}
func easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannel(out *jwriter.Writer, in AuthErr) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"message\":")
	out.String(string(in.Message))
	out.RawByte('}')
}
func easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannelSimpleServiceSillyNoop1(in *jlexer.Lexer, out *SimpleService_SillyNoop_Args) {
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
func easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannelSimpleServiceSillyNoop1(out *jwriter.Writer, in SimpleService_SillyNoop_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SimpleService_SillyNoop_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannelSimpleServiceSillyNoop1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SimpleService_SillyNoop_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson29c9c5f5EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannelSimpleServiceSillyNoop1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SimpleService_SillyNoop_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannelSimpleServiceSillyNoop1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SimpleService_SillyNoop_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson29c9c5f5DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsBazTchannelBazTchannelSimpleServiceSillyNoop1(l, v)
}
