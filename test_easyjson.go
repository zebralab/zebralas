// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package main

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

func easyjson785d9294DecodeGithubComStorage(in *jlexer.Lexer, out *TestSeries) {
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
		case "Labels":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Labels = make(map[string]string)
				} else {
					out.Labels = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 string
					v1 = string(in.String())
					(out.Labels)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "Rs":
			if in.IsNull() {
				in.Skip()
				out.Rs = nil
			} else {
				in.Delim('[')
				if out.Rs == nil {
					if !in.IsDelim(']') {
						out.Rs = make([]uint64, 0, 8)
					} else {
						out.Rs = []uint64{}
					}
				} else {
					out.Rs = (out.Rs)[:0]
				}
				for !in.IsDelim(']') {
					var v2 uint64
					v2 = uint64(in.Uint64())
					out.Rs = append(out.Rs, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "TVs":
			if in.IsNull() {
				in.Skip()
				out.TVs = nil
			} else {
				in.Delim('[')
				if out.TVs == nil {
					if !in.IsDelim(']') {
						out.TVs = make([]TimeValue, 0, 4)
					} else {
						out.TVs = []TimeValue{}
					}
				} else {
					out.TVs = (out.TVs)[:0]
				}
				for !in.IsDelim(']') {
					var v3 TimeValue
					easyjson785d9294DecodeGithubComStorage1(in, &v3)
					out.TVs = append(out.TVs, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Ref":
			out.Ref = uint64(in.Uint64())
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
func easyjson785d9294EncodeGithubComStorage(out *jwriter.Writer, in TestSeries) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Labels\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Labels == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v4First := true
			for v4Name, v4Value := range in.Labels {
				if v4First {
					v4First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v4Name))
				out.RawByte(':')
				out.String(string(v4Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"Rs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Rs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Rs {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Uint64(uint64(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"TVs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.TVs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v7, v8 := range in.TVs {
				if v7 > 0 {
					out.RawByte(',')
				}
				easyjson785d9294EncodeGithubComStorage1(out, v8)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Ref\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Ref))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TestSeries) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson785d9294EncodeGithubComStorage(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TestSeries) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson785d9294EncodeGithubComStorage(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TestSeries) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson785d9294DecodeGithubComStorage(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TestSeries) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson785d9294DecodeGithubComStorage(l, v)
}
func easyjson785d9294DecodeGithubComStorage1(in *jlexer.Lexer, out *TimeValue) {
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
		case "Time":
			out.Time = int64(in.Int64())
		case "Value":
			out.Value = float64(in.Float64())
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
func easyjson785d9294EncodeGithubComStorage1(out *jwriter.Writer, in TimeValue) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Time))
	}
	{
		const prefix string = ",\"Value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float64(float64(in.Value))
	}
	out.RawByte('}')
}
func easyjson785d9294DecodeGithubComStorage2(in *jlexer.Lexer, out *TestDataSet) {
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
		case "TSes":
			if in.IsNull() {
				in.Skip()
				out.TSes = nil
			} else {
				in.Delim('[')
				if out.TSes == nil {
					if !in.IsDelim(']') {
						out.TSes = make([]*TestSeries, 0, 8)
					} else {
						out.TSes = []*TestSeries{}
					}
				} else {
					out.TSes = (out.TSes)[:0]
				}
				for !in.IsDelim(']') {
					var v9 *TestSeries
					if in.IsNull() {
						in.Skip()
						v9 = nil
					} else {
						if v9 == nil {
							v9 = new(TestSeries)
						}
						(*v9).UnmarshalEasyJSON(in)
					}
					out.TSes = append(out.TSes, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Count":
			out.Count = int(in.Int())
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
func easyjson785d9294EncodeGithubComStorage2(out *jwriter.Writer, in TestDataSet) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"TSes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.TSes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v10, v11 := range in.TSes {
				if v10 > 0 {
					out.RawByte(',')
				}
				if v11 == nil {
					out.RawString("null")
				} else {
					(*v11).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Count\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Count))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TestDataSet) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson785d9294EncodeGithubComStorage2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TestDataSet) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson785d9294EncodeGithubComStorage2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TestDataSet) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson785d9294DecodeGithubComStorage2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TestDataSet) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson785d9294DecodeGithubComStorage2(l, v)
}
