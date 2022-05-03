// Code generated by protoc-gen-go-field-setters. DO NOT EDIT.
// versions:
// - protoc-gen-go-field-setters v0.0.0-dev
// - protoc             v3.17.3
// source: scalars.proto

package test

import (
	fmt "fmt"
	fieldsetterplugin "github.com/TheThingsIndustries/protoc-gen-go-field-setters/fieldsetterplugin"
)

// SetFields sets the given fields from src into x.
func (x *MessageWithScalars) SetFields(src *MessageWithScalars, paths ...string) error {
	switch {
	case x == nil && src == nil:
		return nil
	case x == nil:
		return fmt.Errorf("can not set fields into nil MessageWithScalars")
	case src == nil:
		src = &MessageWithScalars{}
	}
	fset := make(fieldsetterplugin.FieldSet)
	for _, field := range fieldsetterplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldsetterplugin.FieldErrorf(x, field, "unknown field")
		case "double_value", "doubleValue":
			x.DoubleValue = src.DoubleValue
		case "double_values", "doubleValues":
			x.DoubleValues = src.DoubleValues
		case "float_value", "floatValue":
			x.FloatValue = src.FloatValue
		case "float_values", "floatValues":
			x.FloatValues = src.FloatValues
		case "int32_value", "int32Value":
			x.Int32Value = src.Int32Value
		case "int32_values", "int32Values":
			x.Int32Values = src.Int32Values
		case "int64_value", "int64Value":
			x.Int64Value = src.Int64Value
		case "int64_values", "int64Values":
			x.Int64Values = src.Int64Values
		case "uint32_value", "uint32Value":
			x.Uint32Value = src.Uint32Value
		case "uint32_values", "uint32Values":
			x.Uint32Values = src.Uint32Values
		case "uint64_value", "uint64Value":
			x.Uint64Value = src.Uint64Value
		case "uint64_values", "uint64Values":
			x.Uint64Values = src.Uint64Values
		case "sint32_value", "sint32Value":
			x.Sint32Value = src.Sint32Value
		case "sint32_values", "sint32Values":
			x.Sint32Values = src.Sint32Values
		case "sint64_value", "sint64Value":
			x.Sint64Value = src.Sint64Value
		case "sint64_values", "sint64Values":
			x.Sint64Values = src.Sint64Values
		case "fixed32_value", "fixed32Value":
			x.Fixed32Value = src.Fixed32Value
		case "fixed32_values", "fixed32Values":
			x.Fixed32Values = src.Fixed32Values
		case "fixed64_value", "fixed64Value":
			x.Fixed64Value = src.Fixed64Value
		case "fixed64_values", "fixed64Values":
			x.Fixed64Values = src.Fixed64Values
		case "sfixed32_value", "sfixed32Value":
			x.Sfixed32Value = src.Sfixed32Value
		case "sfixed32_values", "sfixed32Values":
			x.Sfixed32Values = src.Sfixed32Values
		case "sfixed64_value", "sfixed64Value":
			x.Sfixed64Value = src.Sfixed64Value
		case "sfixed64_values", "sfixed64Values":
			x.Sfixed64Values = src.Sfixed64Values
		case "bool_value", "boolValue":
			x.BoolValue = src.BoolValue
		case "bool_values", "boolValues":
			x.BoolValues = src.BoolValues
		case "string_value", "stringValue":
			x.StringValue = src.StringValue
		case "string_values", "stringValues":
			x.StringValues = src.StringValues
		case "bytes_value", "bytesValue":
			x.BytesValue = src.BytesValue
		case "bytes_values", "bytesValues":
			x.BytesValues = src.BytesValues
		}
		fset.Add(field)
	}
	return nil
}

// SetFields sets the given fields from src into x.
func (x *MessageWithOneofScalars) SetFields(src *MessageWithOneofScalars, paths ...string) error {
	switch {
	case x == nil && src == nil:
		return nil
	case x == nil:
		return fmt.Errorf("can not set fields into nil MessageWithOneofScalars")
	case src == nil:
		src = &MessageWithOneofScalars{}
	}
	fset := make(fieldsetterplugin.FieldSet)
	for _, field := range fieldsetterplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldsetterplugin.FieldErrorf(x, field, "unknown field")
		case "double_value", "doubleValue":
			ov, ok := src.Value.(*MessageWithOneofScalars_DoubleValue)
			if !ok {
				return fieldsetterplugin.FieldErrorf(x, field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "float_value", "floatValue":
			ov, ok := src.Value.(*MessageWithOneofScalars_FloatValue)
			if !ok {
				return fieldsetterplugin.FieldErrorf(x, field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "int32_value", "int32Value":
			ov, ok := src.Value.(*MessageWithOneofScalars_Int32Value)
			if !ok {
				return fieldsetterplugin.FieldErrorf(x, field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "int64_value", "int64Value":
			ov, ok := src.Value.(*MessageWithOneofScalars_Int64Value)
			if !ok {
				return fieldsetterplugin.FieldErrorf(x, field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "uint32_value", "uint32Value":
			ov, ok := src.Value.(*MessageWithOneofScalars_Uint32Value)
			if !ok {
				return fieldsetterplugin.FieldErrorf(x, field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "uint64_value", "uint64Value":
			ov, ok := src.Value.(*MessageWithOneofScalars_Uint64Value)
			if !ok {
				return fieldsetterplugin.FieldErrorf(x, field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "sint32_value", "sint32Value":
			ov, ok := src.Value.(*MessageWithOneofScalars_Sint32Value)
			if !ok {
				return fieldsetterplugin.FieldErrorf(x, field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "sint64_value", "sint64Value":
			ov, ok := src.Value.(*MessageWithOneofScalars_Sint64Value)
			if !ok {
				return fieldsetterplugin.FieldErrorf(x, field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "fixed32_value", "fixed32Value":
			ov, ok := src.Value.(*MessageWithOneofScalars_Fixed32Value)
			if !ok {
				return fieldsetterplugin.FieldErrorf(x, field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "fixed64_value", "fixed64Value":
			ov, ok := src.Value.(*MessageWithOneofScalars_Fixed64Value)
			if !ok {
				return fieldsetterplugin.FieldErrorf(x, field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "sfixed32_value", "sfixed32Value":
			ov, ok := src.Value.(*MessageWithOneofScalars_Sfixed32Value)
			if !ok {
				return fieldsetterplugin.FieldErrorf(x, field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "sfixed64_value", "sfixed64Value":
			ov, ok := src.Value.(*MessageWithOneofScalars_Sfixed64Value)
			if !ok {
				return fieldsetterplugin.FieldErrorf(x, field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "bool_value", "boolValue":
			ov, ok := src.Value.(*MessageWithOneofScalars_BoolValue)
			if !ok {
				return fieldsetterplugin.FieldErrorf(x, field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "string_value", "stringValue":
			ov, ok := src.Value.(*MessageWithOneofScalars_StringValue)
			if !ok {
				return fieldsetterplugin.FieldErrorf(x, field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "bytes_value", "bytesValue":
			ov, ok := src.Value.(*MessageWithOneofScalars_BytesValue)
			if !ok {
				return fieldsetterplugin.FieldErrorf(x, field, "invalid Value of type %T in source struct", src.Value)
			}
			x.Value = ov
		case "value":
			x.Value = src.Value
		}
		fset.Add(field)
	}
	return nil
}

// SetFields sets the given fields from src into x.
func (x *MessageWithScalarMaps) SetFields(src *MessageWithScalarMaps, paths ...string) error {
	switch {
	case x == nil && src == nil:
		return nil
	case x == nil:
		return fmt.Errorf("can not set fields into nil MessageWithScalarMaps")
	case src == nil:
		src = &MessageWithScalarMaps{}
	}
	fset := make(fieldsetterplugin.FieldSet)
	for _, field := range fieldsetterplugin.TopLevelPaths(paths) {
		if fset.Contains(field) {
			continue
		}
		switch field {
		default:
			return fieldsetterplugin.FieldErrorf(x, field, "unknown field")
		case "string_double_map", "stringDoubleMap":
			x.StringDoubleMap = src.StringDoubleMap
		case "string_float_map", "stringFloatMap":
			x.StringFloatMap = src.StringFloatMap
		case "string_int32_map", "stringInt32Map":
			x.StringInt32Map = src.StringInt32Map
		case "int32_string_map", "int32StringMap":
			x.Int32StringMap = src.Int32StringMap
		case "string_int64_map", "stringInt64Map":
			x.StringInt64Map = src.StringInt64Map
		case "int64_string_map", "int64StringMap":
			x.Int64StringMap = src.Int64StringMap
		case "string_uint32_map", "stringUint32Map":
			x.StringUint32Map = src.StringUint32Map
		case "uint32_string_map", "uint32StringMap":
			x.Uint32StringMap = src.Uint32StringMap
		case "string_uint64_map", "stringUint64Map":
			x.StringUint64Map = src.StringUint64Map
		case "uint64_string_map", "uint64StringMap":
			x.Uint64StringMap = src.Uint64StringMap
		case "string_sint32_map", "stringSint32Map":
			x.StringSint32Map = src.StringSint32Map
		case "sint32_string_map", "sint32StringMap":
			x.Sint32StringMap = src.Sint32StringMap
		case "string_sint64_map", "stringSint64Map":
			x.StringSint64Map = src.StringSint64Map
		case "sint64_string_map", "sint64StringMap":
			x.Sint64StringMap = src.Sint64StringMap
		case "string_fixed32_map", "stringFixed32Map":
			x.StringFixed32Map = src.StringFixed32Map
		case "fixed32_string_map", "fixed32StringMap":
			x.Fixed32StringMap = src.Fixed32StringMap
		case "string_fixed64_map", "stringFixed64Map":
			x.StringFixed64Map = src.StringFixed64Map
		case "fixed64_string_map", "fixed64StringMap":
			x.Fixed64StringMap = src.Fixed64StringMap
		case "string_sfixed32_map", "stringSfixed32Map":
			x.StringSfixed32Map = src.StringSfixed32Map
		case "sfixed32_string_map", "sfixed32StringMap":
			x.Sfixed32StringMap = src.Sfixed32StringMap
		case "string_sfixed64_map", "stringSfixed64Map":
			x.StringSfixed64Map = src.StringSfixed64Map
		case "sfixed64_string_map", "sfixed64StringMap":
			x.Sfixed64StringMap = src.Sfixed64StringMap
		case "string_bool_map", "stringBoolMap":
			x.StringBoolMap = src.StringBoolMap
		case "bool_string_map", "boolStringMap":
			x.BoolStringMap = src.BoolStringMap
		case "string_string_map", "stringStringMap":
			x.StringStringMap = src.StringStringMap
		case "string_bytes_map", "stringBytesMap":
			x.StringBytesMap = src.StringBytesMap
		}
		fset.Add(field)
	}
	return nil
}
