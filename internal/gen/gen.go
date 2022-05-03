// Copyright © 2022 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package gen

import (
	"fmt"
	"strings"

	"github.com/TheThingsIndustries/protoc-gen-go-field-setters/annotations"
	"github.com/TheThingsIndustries/protoc-gen-go-field-setters/internal/gogoproto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Version is the version of the generator.
var Version = "0.0.0-dev"

// Params are the parameters for the generator.
var Params struct {
	Lang string
}

const (
	fmtPackage = protogen.GoImportPath("fmt")
	pp         = protogen.GoImportPath("github.com/TheThingsIndustries/protoc-gen-go-field-setters/fieldsetterplugin")
)

type generator struct {
	gen  *protogen.Plugin
	file *protogen.File
	*protogen.GeneratedFile
}

// GenerateFile generates a file with FieldSetters.
func GenerateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	g := &generator{
		gen:  gen,
		file: file,
	}

	// If the file doesn't have FieldSetters, we can skip it.
	if !g.fileHasFieldSetter() {
		return nil
	}

	// Generate a new file that ends with `_field_setters.pb.go`.
	filename := file.GeneratedFilenamePrefix + "_field_setters.pb.go"
	g.GeneratedFile = gen.NewGeneratedFile(filename, file.GoImportPath)

	// The standard header for generated files.
	g.P("// Code generated by protoc-gen-go-field-setters. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc-gen-go-field-setters v", Version)
	g.P("// - protoc             ", g.protocVersion())
	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}

	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	g.generateFileContent()

	return g.GeneratedFile
}

func containsMessage(s []*protogen.Message, search *protogen.Message) bool {
	for _, el := range s {
		if el == search {
			return true
		}
	}
	return false
}

func appendSubMessages(dst []*protogen.Message, msgs ...*protogen.Message) []*protogen.Message {
	for _, msg := range msgs {
		if !containsMessage(dst, msg) {
			dst = append(dst, msg)
		}
		for _, subMsg := range msg.Messages {
			if !containsMessage(dst, subMsg) {
				dst = appendSubMessages(dst, subMsg)
			}
		}
	}
	return dst
}

func (g *generator) fileHasFieldSetter() bool {
	for _, message := range appendSubMessages(nil, g.file.Messages...) {
		if g.messageHasFieldSetter(message) {
			return true
		}
	}
	return false
}

func (*generator) messageHasFieldSetter(message *protogen.Message) bool {
	// No code is generated for map entries, so we also don't need to generate field_setters.
	if message.Desc.IsMapEntry() {
		return false
	}

	var generateFieldSetter bool

	// If the file has the (thethings.fieldsetters.file) option, and field_setter_all is set, we start with that.
	fileOpts := message.Desc.ParentFile().Options().(*descriptorpb.FileOptions)
	if proto.HasExtension(fileOpts, annotations.E_File) {
		if fileExt, ok := proto.GetExtension(fileOpts, annotations.E_File).(*annotations.FileOptions); ok {
			if fileExt.FieldSetterAll != nil {
				generateFieldSetter = *fileExt.FieldSetterAll
			}
		}
	}

	// If the message has the (thethings.fieldsetters.message) option, the field_setter field can still override to true or false if explicitly set.
	messageOpts := message.Desc.Options().(*descriptorpb.MessageOptions)
	if proto.HasExtension(messageOpts, annotations.E_Message) {
		if messageExt, ok := proto.GetExtension(messageOpts, annotations.E_Message).(*annotations.MessageOptions); ok {
			if messageExt.FieldSetter != nil {
				generateFieldSetter = *messageExt.FieldSetter
			}
		}
	}

	return generateFieldSetter
}

func (g *generator) protocVersion() string {
	v := g.gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}

func (g *generator) generateFileContent() {
	for _, message := range g.file.Messages {
		g.genMessage(message)
	}
}

func fieldGoName(field *protogen.Field) interface{} {
	var fieldGoName interface{} = field.GoName
	if Params.Lang == "gogo" {
		// If we use gogo, the GoName of a field can be overridden with the (gogoproto.customname) option.
		if proto.HasExtension(field.Desc.Options(), gogoproto.E_Customname) {
			fieldGoName = proto.GetExtension(field.Desc.Options(), gogoproto.E_Customname).(string)
		}
		// Fields with the (gogoproto.embed) option should use the name of the embedded message.
		if proto.HasExtension(field.Desc.Options(), gogoproto.E_Embed) {
			if proto.GetExtension(field.Desc.Options(), gogoproto.E_Embed).(bool) {
				fieldGoName = field.Message.GoIdent
			}
		}
	}
	return fieldGoName
}

func fieldIsNullable(field *protogen.Field) bool {
	// Typically, only message fields are nullable (use pointers).
	nullable := field.Desc.Kind() == protoreflect.MessageKind

	// If we use gogo, the nullability of fields (messages, but also bytes fields with custom types) is controlled with the (gogoproto.nullable) option.
	if Params.Lang == "gogo" {
		if proto.HasExtension(field.Desc.Options(), gogoproto.E_Customtype) {
			nullable = true
		}
		if proto.HasExtension(field.Desc.Options(), gogoproto.E_Nullable) {
			nullable = proto.GetExtension(field.Desc.Options(), gogoproto.E_Nullable).(bool)
		}
	}

	return nullable
}

func filterFields(fields []*protogen.Field, selector func(f *protogen.Field) bool) []*protogen.Field {
	out := make([]*protogen.Field, 0, len(fields))
	for _, path := range fields {
		if selector(path) {
			out = append(out, path)
		}
	}
	return out
}

func fieldIsMessage(f *protogen.Field) bool {
	if f.Message == nil || f.Desc.IsList() || f.Desc.IsMap() {
		return false
	}
	if strings.HasPrefix(string(f.Message.Desc.FullName()), "google.protobuf.") {
		return false
	}
	return true
}

func (g *generator) genMessage(message *protogen.Message) { //nolint:gocyclo
	if !g.messageHasFieldSetter(message) {
		return
	}

	// Generate setters for all sub-messages defined in the message.
	for _, message := range message.Messages {
		g.genMessage(message)
	}

	g.P("// SetFields sets the given fields from src into x.")
	g.P("func (x *", message.GoIdent, ") SetFields(src *", message.GoIdent, ", paths ...string) error {")

	// First we deal with nil.
	g.P("switch {")
	g.P("case x == nil && src == nil:")
	// If both source and destination are nil, there's nothing to do.
	g.P("return nil")
	g.P("case x == nil:")
	// If the destination is nil, that's an implementation bug, so perhaps this should even panic.
	g.P("return ", fmtPackage.Ident("Errorf"), `("can not set fields into nil `, message.GoIdent, `")`)
	g.P("case src == nil:")
	// If the source is nil, we just treat it as a "zero" message.
	g.P("src = &", message.GoIdent, "{}")
	g.P("}") // end switch

	// This is the set of fields we'll use to make sure that we don't try setting the same field twice.
	g.P("fset := make(", pp.Ident("FieldSet"), ")")

	// We first look at exact fields (without sub-fields).
	// That way, if paths contains both the field, and sub-fields, we can skip the logic for sub-paths.
	g.P("for _, field := range ", pp.Ident("TopLevelPaths"), "(paths) {")

	// Skip duplicate fields.
	g.P("if fset.Contains(field) {")
	g.P("continue")
	g.P("}")

	g.P("switch field {")

	// Error if we don't know the field.
	g.P("default:")
	g.P("return ", pp.Ident("FieldErrorf"), `(x, field, "unknown field")`)

	for _, field := range message.Fields {
		fieldGoName := fieldGoName(field)

		// The "case" supports both the proto name and the JSON name, if those are different.
		if string(field.Desc.Name()) != field.Desc.JSONName() {
			g.P(`case "`, field.Desc.Name(), `", "`, field.Desc.JSONName(), `":`)
		} else {
			g.P(`case "`, field.Desc.Name(), `":`)
		}

		if field.Oneof != nil {
			// If this field is in a oneof, assert that the source has the oneof set to the expected type, and replace the oneof.
			// Note that, in order to unset the oneof, we use the name of the oneof directly.
			g.P("ov, ok := src.", field.Oneof.GoName, ".(*", field.GoIdent.GoName, ")")
			g.P("if !ok {")
			g.P("return ", pp.Ident("FieldErrorf"), `(x, field, "invalid `, field.Oneof.GoName, ` of type %T in source struct", src.`, field.Oneof.GoName, ")")
			g.P("}")
			g.P("x.", field.Oneof.GoName, " = ov")
		} else {
			// Just set the field in the destination to the same field in the source.
			g.P("x.", fieldGoName, " = src.", fieldGoName)
		}
	}

	for _, oneof := range message.Oneofs {
		// For each oneof, we can set the oneof directly. This is also used to set a oneof to nil.
		g.P(`case "`, oneof.Desc.Name(), `":`) // NOTE: oneofs don't have a JSONName.
		g.P("x.", oneof.GoName, " = src.", oneof.GoName)
	}

	g.P("}") // end switch field
	g.P("fset.Add(field)")
	g.P("}") // end for _, field := range ...

	// After we've completed handling the "exact match" paths, we look at everything that can have nested fields.
	fieldsWithSubFields := filterFields(message.Fields, fieldIsMessage)
	if len(fieldsWithSubFields) > 0 {
		g.P("for _, field := range ", pp.Ident("SubPaths"), "(paths) {")

		g.P("topLevelField := ", pp.Ident("TopLevelField"), "(field)")

		// Skip duplicate fields.
		g.P("if fset.Contains(topLevelField) {")
		g.P("continue")
		g.P("}")

		g.P("switch topLevelField {")

		// Error if we don't know the field.
		g.P("default:")
		g.P("return ", pp.Ident("FieldErrorf"), `(x, field, "unknown field")`)

		for _, field := range fieldsWithSubFields {
			var (
				fieldGoName = fieldGoName(field)
				nullable    = fieldIsNullable(field)
			)

			// The "case" supports both the proto name and the JSON name, if those are different.
			if string(field.Desc.Name()) != field.Desc.JSONName() {
				g.P(`case "`, field.Desc.Name(), `", "`, field.Desc.JSONName(), `":`)
			} else {
				g.P(`case "`, field.Desc.Name(), `":`)
			}

			if field.Oneof != nil {
				if nullable {
					// If the field is nullable and the destination is nil, but the source isn't,
					// we initialize new oneof wrapper containing a new sub-message.
					g.P("if x.", field.Oneof.GoName, " == nil && src.", field.Oneof.GoName, " != nil {")
					g.P("ov := &", field.GoIdent.GoName, "{", fieldGoName, ": &", field.Message.GoIdent, "{}}")
					g.P("x.", field.Oneof.GoName, " = ov")
					g.P("}")
				}
				// We assert that the oneof wrapper in the destination has the expected type.
				g.P("xOV, ok := x.", field.Oneof.GoName, ".(*", field.GoIdent.GoName, ")")
				g.P("if !ok {")
				g.P("return ", pp.Ident("FieldErrorf"), `(x, field, "invalid `, field.Oneof.GoName, ` of type %T in destination struct", x.`, field.Oneof.GoName, ")")
				g.P("}")
				// We assert that the oneof wrapper in the source has the expected type.
				g.P("srcOV, ok := src.", field.Oneof.GoName, ".(*", field.GoIdent.GoName, ")")
				g.P("if !ok {")
				g.P("return ", pp.Ident("FieldErrorf"), `(x, field, "invalid `, field.Oneof.GoName, ` of type %T in source struct", src.`, field.Oneof.GoName, ")")
				g.P("}")
				// We set the sub-fields in the sub-message.
				g.P("if err := xOV.", fieldGoName, ".SetFields(srcOV.", fieldGoName, ", ", pp.Ident("SubPathsOf"), "(paths, topLevelField)...); err != nil {")
				g.P("return ", pp.Ident("WrapFieldError"), "(x, field, err)")
				g.P("}")
			} else if nullable {
				// If the field is nullable and the destination is nil, but the source isn't,
				// we initialize a new sub-message.
				g.P("if x.", fieldGoName, " == nil && src.", fieldGoName, " != nil {")
				g.P("var v ", field.Message.GoIdent)
				g.P("x.", fieldGoName, " = &v")
				g.P("}")
				// We set the sub-fields in the sub-message.
				g.P("if err := x.", fieldGoName, ".SetFields(src.", fieldGoName, ", ", pp.Ident("SubPathsOf"), "(paths, topLevelField)...); err != nil {")
				g.P("return ", pp.Ident("WrapFieldError"), "(x, field, err)")
				g.P("}")
			} else {
				// We set the sub-fields in the sub-message.
				g.P("if err := x.", fieldGoName, ".SetFields(&src.", fieldGoName, ", ", pp.Ident("SubPathsOf"), "(paths, topLevelField)...); err != nil {")
				g.P("return ", pp.Ident("WrapFieldError"), "(x, field, err)")
				g.P("}")
			}
		}

		for _, oneof := range message.Oneofs {
			// Even though the name of the oneof should not be part of a field mask sub-path, we allow it anyway.
			g.P(`case "`, oneof.Desc.Name(), `":`) // NOTE: oneofs don't have a JSONName.
			g.P("if err := x.SetFields(src, ", pp.Ident("SubPathsOf"), `(paths, topLevelField)...); err != nil {`)
			g.P("return err")
			g.P("}")
		}

		g.P("}") // end switch topLevelField
		g.P("fset.Add(topLevelField)")
		g.P("}") // end for _, field := range ...
	}

	g.P("return nil")
	g.P("}") // end func SetFields
}
