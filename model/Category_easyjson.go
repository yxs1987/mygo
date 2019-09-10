// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

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

func easyjson21ada79cDecodeMygoModel(in *jlexer.Lexer, out *Category) {
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
		case "category_id":
			out.CategoryId = int64(in.Int64Str())
		case "category_name":
			out.CategoryName = string(in.String())
		case "child_category":
			if in.IsNull() {
				in.Skip()
				out.ChildCategory = nil
			} else {
				in.Delim('[')
				if out.ChildCategory == nil {
					if !in.IsDelim(']') {
						out.ChildCategory = make([]Category, 0, 1)
					} else {
						out.ChildCategory = []Category{}
					}
				} else {
					out.ChildCategory = (out.ChildCategory)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Category
					(v1).UnmarshalEasyJSON(in)
					out.ChildCategory = append(out.ChildCategory, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "image":
			out.Image = string(in.String())
		case "created_at":
			out.CreatedAt = string(in.String())
		case "updated_at":
			out.UpdatedAt = string(in.String())
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
func easyjson21ada79cEncodeMygoModel(out *jwriter.Writer, in Category) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"category_id\":"
		out.RawString(prefix[1:])
		out.Int64Str(int64(in.CategoryId))
	}
	{
		const prefix string = ",\"category_name\":"
		out.RawString(prefix)
		out.String(string(in.CategoryName))
	}
	{
		const prefix string = ",\"child_category\":"
		out.RawString(prefix)
		if in.ChildCategory == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.ChildCategory {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"image\":"
		out.RawString(prefix)
		out.String(string(in.Image))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.String(string(in.CreatedAt))
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.String(string(in.UpdatedAt))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Category) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson21ada79cEncodeMygoModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Category) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson21ada79cEncodeMygoModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Category) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson21ada79cDecodeMygoModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Category) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson21ada79cDecodeMygoModel(l, v)
}