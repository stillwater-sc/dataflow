/**
 * File		:	$File: //depot/stillwater/dataflow/types/operand.go $
 *
 * Authors	:	E. Theodore L. Omtzigt
 * Date		:	4 May 2016
 *
 * Source Control Information:
 * Version	:	$Revision: #1 $
 * Latest	:	$Date: 2016/05/04 $
 * Location	:	$Id: //depot/stillwater/dataflow/types/operand.go#1 $
 *
 * Organization:
 *		Stillwater Supercomputing, Inc.
 *		P.O Box 720
 *		South Freeport, ME 04078-0720
 *
 * Copyright (c) 2006-2016 E. Theodore L. Omtzigt.  All rights reserved.
 *
 * Licence      : Stillwater license as defined in this directory
 */
package types

import (
	"fmt"
	"encoding/json"
)

/*
 * OperandType  encoding identifier for the data format of the operand
 */
type OperandType uint8

const (
	// Data types
	FORMAT_NAN		OperandType = 0x00 // zero value type

	// explicit data types
	FORMAT_UINT8		= 0x01
	FORMAT_UINT16		= 0x02
	FORMAT_UINT32		= 0x03
	FORMAT_UINT64		= 0x04
	FORMAT_INT8		= 0x05
	FORMAT_INT16		= 0x06
	FORMAT_INT32		= 0x07
	FORMAT_INT64		= 0x08
	FORMAT_FLOAT32 		= 0x09
	FORMAT_FLOAT64 		= 0x0A
	FORMAT_UNUM			= 0x10 // TODO: non-Golang type

	FORMAT_RAW8 		= 0x20
	FORMAT_RAW16 		= 0x21

	FORMAT_PROGRAM		= 0x30
)

var OPERANDTYPE_LOOKUP_S2O = map[string]OperandType{
	"NAN":		FORMAT_NAN,
	"UINT8":	FORMAT_UINT8,
	"UINT16":	FORMAT_UINT16,
	"UINT32":	FORMAT_UINT32,
	"UINT64":	FORMAT_UINT64,
	"INT8":		FORMAT_INT8,
	"INT16":	FORMAT_INT16,
	"INT32":	FORMAT_INT32,
	"INT64":	FORMAT_INT64,
	"FLOAT32":	FORMAT_FLOAT32,
	"FLOAT64":	FORMAT_FLOAT64,
	"UNUM":		FORMAT_UNUM,
	"RAW8":		FORMAT_RAW8,
	"RAW16":	FORMAT_RAW16,
	"PROGRAM":	FORMAT_PROGRAM,
}

var OPERANDTYPE_LOOKUP_O2S = map[OperandType]string{
	FORMAT_NAN:		"NAN",
	FORMAT_UINT8:		"UINT8",
	FORMAT_UINT16:		"UINT16",
	FORMAT_UINT32:		"UINT32",
	FORMAT_UINT64:		"UINT64",
	FORMAT_INT8:		"INT8",
	FORMAT_INT16:		"INT16",
	FORMAT_INT32:		"INT32",
	FORMAT_INT64:		"INT64",
	FORMAT_FLOAT32:		"FLOAT32",
	FORMAT_FLOAT64:		"FLOAT64",
	FORMAT_UNUM:		"UNUM",
	FORMAT_RAW8:		"RAW8",
	FORMAT_RAW16:		"RAW16",
	FORMAT_PROGRAM:		"PROGRAM",
}

func (t OperandType) SizeInBytes() (size int) {
	size = 0
	switch t {
	case FORMAT_NAN:
		size = 0
	case FORMAT_UINT8, FORMAT_INT8:
		size = 1
	case FORMAT_UINT16, FORMAT_INT16:
		size = 2
	case FORMAT_UINT32, FORMAT_INT32, FORMAT_FLOAT32:
		size = 4
	case FORMAT_UINT64, FORMAT_INT64, FORMAT_FLOAT64, FORMAT_RAW8:
		size = 8
	case FORMAT_RAW16:
		size = 16
	case FORMAT_UNUM:
		size = 16		// TODO
	}
	return size
}

/*
 * Tagged Union to capture the ability of a hardware register
 * to contain many different types
 */
type Operand struct {
	t 	OperandType		// type identifier
	u8	uint8
	u16 	uint16
	u32 	uint32
	u64 	uint64
	i8 	int8
	i16 	int16
	i32 	int32
	i64 	int64
	f32 	float32
	f64	float64
	raw8 	[8]byte
	raw16   [16]byte
}

/////////////////////////////////////////////////////////////////
// Selectors

func (o Operand) String() (s string) {
	switch o.t {
	case FORMAT_NAN:
		s = fmt.Sprintf("nan")
	case FORMAT_UINT8:
		s = fmt.Sprintf("uint8(%d)", o.u8)
	case FORMAT_UINT16:
		s = fmt.Sprintf("uint16(%d)", o.u16)
	case FORMAT_UINT32:
		s = fmt.Sprintf("uint32(%d)", o.u32)
	case FORMAT_UINT64:
		s = fmt.Sprintf("uint64(%d)", o.u64)
	case FORMAT_INT8:
		s = fmt.Sprintf("int8(%d)", o.i8)
	case FORMAT_INT16:
		s = fmt.Sprintf("int16(%d)", o.i16)
	case FORMAT_INT32:
		s = fmt.Sprintf("int32(%d)", o.i32)
	case FORMAT_INT64:
		s = fmt.Sprintf("int64(%d)", o.i64)
	case FORMAT_FLOAT32:
		s = fmt.Sprintf("fp32(%f)", o.f32)
	case FORMAT_FLOAT64:
		s = fmt.Sprintf("fp64(%g)", o.f64)
	case FORMAT_UNUM:
		s = fmt.Sprintf("unum(%v)", o.raw8)
	case FORMAT_RAW8:
		s = fmt.Sprintf("raw8(%v)", o.raw8)
	case FORMAT_RAW16:
		s = fmt.Sprintf("raw16(%v)", o.raw16)
	case FORMAT_PROGRAM:
		s = fmt.Sprintf("raw8(%v)", o.raw8)
	}
	return
}
func (o Operand) DataType() OperandType { return o.t }
func (o Operand) Uint8() uint8 { return o.u8 }
func (o Operand) Uint16() uint16 { return o.u16 }
func (o Operand) Uint32() uint32{ return o.u32 }
func (o Operand) Uint64() uint64 { return o.u64 }
func (o Operand) Int8() int8 { return o.i8 }
func (o Operand) Int16() int16 { return o.i16 }
func (o Operand) Int32() int32 { return o.i32 }
func (o Operand) Int64() int64 { return o.i64 }
func (o Operand) F32() float32 { return o.f32 }
func (o Operand) F64() float64 { return o.f64 }
func (o Operand) Raw8() [8]byte { return o.raw8 }
func (o Operand) Raw16() [16]byte { return o.raw16 }

/////////////////////////////////////////////////////////////////
// Modifiers

func (o *Operand) SetNan() { o.t = FORMAT_NAN }
func (o *Operand) SetUint8(u8 uint8) { o.t = FORMAT_UINT8; o.u8 = u8 }
func (o *Operand) SetUint16(u16 uint16) { o.t = FORMAT_UINT16; o.u16 = u16 }
func (o *Operand) SetUint32(u32 uint32) { o.t = FORMAT_UINT32; o.u32 = u32 }
func (o *Operand) SetUint64(u64 uint64) { o.t = FORMAT_UINT64; o.u64 = u64 }
func (o *Operand) SetInt8(i8 int8) { o.t = FORMAT_INT8; o.i8 = i8 }
func (o *Operand) SetInt16(i16 int16) { o.t = FORMAT_INT16; o.i16 = i16 }
func (o *Operand) SetInt32(i32 int32) { o.t = FORMAT_INT32; o.i32 = i32 }
func (o *Operand) SetInt64(i64 int64) { o.t = FORMAT_INT64; o.i64 = i64 }
func (o *Operand) SetF32(f32 float32) { o.t = FORMAT_FLOAT32; o.f32 = f32 }
func (o *Operand) SetF64(f64 float64) { o.t = FORMAT_FLOAT64; o.f64 = f64 }
func (o *Operand) SetRaw8(raw8 [8]byte) { o.t = FORMAT_RAW8; o.raw8 = raw8 }
func (o *Operand) SetRaw16(raw16 [16]byte) { o.t = FORMAT_RAW16; o.raw16 = raw16 }


/////////////////////////////////////////////////////////////////
// JSON Marshal/Unmarshal helpers

/*
 *
 * Key observation: all values in JS are floats and ints, so use that baseline with
 * human readable data, packet, and recurrence types.
 */
func (o Operand) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct{
		T 	OperandType	`json:"t"`
		U8	uint8	`json:"u8"`
		U16 	uint16	`json:"u16"`
		U32 	uint32	`json:"u32"`
		U64 	uint64	`json:"u64"`
		I8 	int8	`json:"i8"`
		I16 	int16	`json:"i16"`
		I32 	int32	`json:"i32"`
		I64 	int64	`json:"i64"`
		F32 	float32 `json:"f32"`
		F64	float64 `json:"f64"`
		Raw8 	[8]byte    `json:"raw8"`
		Raw16   [16]byte   `json:"raw16"`
	}{
		T: 	o.t,
		U8:	o.u8,
		U16:	o.u16,
		U32:	o.u32,
		U64:	o.u64,
		I8:	o.i8,
		I16:	o.i16,
		I32:	o.i32,
		I64:	o.i64,
		F32:	o.f32,
		F64:	o.f64,
		Raw8:	o.raw8,
		Raw16:	o.raw16,
	})
}
