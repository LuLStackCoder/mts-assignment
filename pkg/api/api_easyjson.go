// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package api

import (
	json "encoding/json"
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

func easyjsonC1cedd36DecodeGithubComLuLStackCoderMtsAssignmentPkgApi(in *jlexer.Lexer, out *URLData) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "url":
			out.URL = string(in.String())
		case "body":
			out.Body = string(in.String())
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
func easyjsonC1cedd36EncodeGithubComLuLStackCoderMtsAssignmentPkgApi(out *jwriter.Writer, in URLData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix[1:])
		out.String(string(in.URL))
	}
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix)
		out.String(string(in.Body))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v URLData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComLuLStackCoderMtsAssignmentPkgApi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v URLData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComLuLStackCoderMtsAssignmentPkgApi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *URLData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComLuLStackCoderMtsAssignmentPkgApi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *URLData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComLuLStackCoderMtsAssignmentPkgApi(l, v)
}
func easyjsonC1cedd36DecodeGithubComLuLStackCoderMtsAssignmentPkgApi1(in *jlexer.Lexer, out *HandleUrlsResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]URLData, 0, 2)
					} else {
						out.Data = []URLData{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v1 URLData
					(v1).UnmarshalEasyJSON(in)
					out.Data = append(out.Data, v1)
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
func easyjsonC1cedd36EncodeGithubComLuLStackCoderMtsAssignmentPkgApi1(out *jwriter.Writer, in HandleUrlsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix[1:])
		if in.Data == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Data {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HandleUrlsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComLuLStackCoderMtsAssignmentPkgApi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HandleUrlsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComLuLStackCoderMtsAssignmentPkgApi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HandleUrlsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComLuLStackCoderMtsAssignmentPkgApi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HandleUrlsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComLuLStackCoderMtsAssignmentPkgApi1(l, v)
}
func easyjsonC1cedd36DecodeGithubComLuLStackCoderMtsAssignmentPkgApi2(in *jlexer.Lexer, out *HandleUrlsRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(HandleUrlsRequest, 0, 4)
			} else {
				*out = HandleUrlsRequest{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 string
			v4 = string(in.String())
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC1cedd36EncodeGithubComLuLStackCoderMtsAssignmentPkgApi2(out *jwriter.Writer, in HandleUrlsRequest) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			out.String(string(v6))
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v HandleUrlsRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComLuLStackCoderMtsAssignmentPkgApi2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HandleUrlsRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComLuLStackCoderMtsAssignmentPkgApi2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HandleUrlsRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComLuLStackCoderMtsAssignmentPkgApi2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HandleUrlsRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComLuLStackCoderMtsAssignmentPkgApi2(l, v)
}
