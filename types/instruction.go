/**
 * File		:	$File: //depot/stillwater/dataflow/types/instruction.go $
 *
 * Authors	:	E. Theodore L. Omtzigt
 * Date		:	21 April 2016
 *
 * Source Control Information:
 * Version	:	$Revision: #1 $
 * Latest	:	$Date: 2016/04/21 $
 * Location	:	$Id: //depot/stillwater/dataflow/types/instruction.go#1 $
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
)

// MIPS ISA
const (
	OPCODE_NOOP 		uint8 = 0		// zero value type

	// opcode ranges to determine, if we have 1, 2, or 3-operand instructions
	// one operand opcodes
	OPCODE_MOV     		uint8 = 1
	OPCODE_CMOV    		uint8 = 2
	OPCODE_RECIP   		uint8 = 6
	OPCODE_HASH   		uint8 = 7
	OPCODE_SQRT    		uint8 = 12
	OPCODE_TAN    		uint8 = 13
	OPCODE_COS     		uint8 = 14
	OPCODE_SIN    		uint8 = 15
	OPCODE_ONE_OPERAND	uint8= 16

	// two operand opcodes
	OPCODE_IADD2 		uint8 = 16
	OPCODE_ISUB2 		uint8 = 17
	OPCODE_IMUL2 		uint8 = 18
	OPCODE_IDIV2 		uint8 = 19
	OPCODE_FADD2 		uint8 = 20
	OPCODE_FSUB2 		uint8 = 21
	OPCODE_FMUL2 		uint8 = 22
	OPCODE_FDIV2 		uint8 = 23
	OPCODE_CADD2 		uint8 = 24
	OPCODE_CSUB2 		uint8 = 25
	OPCODE_CMUL2 		uint8 = 26
	OPCODE_CDIV2 		uint8 = 27
	OPCODE_ICMP 		uint8 = 28
	OPCODE_FCMP 		uint8 = 29
	OPCODE_CCMP 		uint8 = 30
	OPCODE_LERP		uint8 = 31
	OPCODE_TWO_OPERAND	uint8 = 32

	// three operand opcodes
	OPCODE_IADD3		uint8 = 32
	OPCODE_ISUB3		uint8 = 33
	OPCODE_IMAC3		uint8 = 34
	OPCODE_IDIV3		uint8 = 35
	OPCODE_FADD3		uint8 = 36
	OPCODE_FSUB3		uint8 = 37
	OPCODE_FMAC3		uint8 = 38
	OPCODE_FDIV3		uint8 = 39
	OPCODE_CADD3		uint8 = 40
	OPCODE_CSUB3		uint8 = 41
	OPCODE_CMAC3		uint8 = 42
	OPCODE_CDIV3		uint8 = 43
	OPCODE_THREE_OPERAND	uint8 = 64

	OPCODE_INVALID 		uint8 = 63
)

var (
	OpcodeName OpcodeNames = OpcodeNames{
		"NOOP",					// 0000 0000    0
		"MOV",					// 0000 0001    1
		"CMOV",					// 0000 0010    2
		"",					// 0000 0011    3
		"",					// 0000 0100    4
		"",					// 0000 0101    5
		"RECIP",				// 0000 0110    6
		"HASH",					// 0000 0111    7
		"",					// 0000 1000    8
		"",					// 0000 1001    9
		"",					// 0000 1010    10
		"",					// 0000 1011    11
		"SQRT",					// 0000 1100    12
		"TAN",					// 0000 1101    13
		"COS",					// 0000 1110    14
		"SIN",					// 0000 1111    15

		"IADD2",				// 0001 0000    16
		"ISUB2",				// 0001 0001    17
		"IMUL2",				// 0001 0010    18
		"IDIV2",				// 0001 0011    19
		"FADD2",				// 0001 0100    20
		"FSUB2",				// 0001 0101    21
		"FMUL2",				// 0001 0110    22
		"FDIV2",				// 0001 0111    23
		"CADD2",				// 0001 1000    24
		"CSUB2",				// 0001 1001    25
		"CMUL2",				// 0001 1010    26
		"CDIV2",				// 0001 1011    27
		"ICMP",					// 0001 1100    28
		"FCMP",					// 0001 1101    29
		"CCMP",					// 0001 1110    30
		"LERP",					// 0001 1111    31

		"IADD3",				// 0010 0000    32
		"ISUB3",				// 0010 0001    33
		"IMAC3",				// 0010 0010    34
		"IDIV3",				// 0010 0011    35
		"FADD3",				// 0010 0100    36
		"FSUB3",				// 0010 0101    37
		"FMAC3",				// 0010 0110    38
		"FDIV3",				// 0010 0111    39
		"CADD3",				// 0010 1000    40
		"CSUB3",				// 0010 1001    41
		"CMAC3",				// 0010 1010    42
		"CDIV3",				// 0010 1011    43
		"",					// 0010 1100    44
		"",					// 0010 1101    45
		"",					// 0010 1110    46
		"",					// 0010 1111    47
		"",					// 0011 0000    48
		"",					// 0011 0001    49
		"",					// 0011 0010    50
		"",					// 0011 0011    51
		"",					// 0011 0100    52
		"",					// 0011 0101    53
		"",					// 0011 0110    54
		"",					// 0011 0111    55
		"",					// 0011 1000    56
		"",					// 0011 1001    57
		"",					// 0011 1010    58
		"",					// 0011 1011    59
		"",					// 0011 1100    60
		"",					// 0011 1101    61
		"",					// 0011 1110    62
		"INVALID",				// 0011 1111    63
	}
)

/*
 * Static array of string encodings of the ISA of the simulated DFM
 * Mostly used for human-readable debugging and tracing of execution.
 */
type OpcodeNames [64]string

/*
 * The encoded function for the instruction in the ISA of the Execute Unit
 */
type Opcode  	 uint8

/*
 * An instruction is the data structure that is assembled in the iStore.
 * The iStore is an associative memory keyed on the Tag, which is a
 * unique, 64bit number.
 *
 * When an instruction has received all its operands it is evicted from the iStore
 * and forwarded to the execution units. The result will yield a new data token
 * with a new tag, which is computed along side the result. We therefore
 * piggyback the Tag on the Instruction as well so we have a nice compact
 * abstraction for the information that flows from the iStore through the router,
 * to the execute units, and result tokens flow back to the iStore via the return router.
 */
type Instruction struct {
	Tag	// contains the ReId and IndexVector of the computational event
	opcode 		Opcode
	operands	[3]Operand	// 3 operand ISA
	opValid	 	[3]bool
	slot		uint8
}

/////////////////////////////////////////////////////////////////
// Selectors

func (i *Instruction) GetTag() Tag { return i.Tag }
func (i *Instruction) Slot() uint8 { return i.slot }
func (i *Instruction) Opcode() Opcode { return i.opcode }
func (i *Instruction) Operand(slot int) Operand { return i.operands[slot] }
func (i *Instruction) OperandValid(slot int) bool { return i.opValid[slot] }

func (i Instruction) String() string {
	return fmt.Sprintf("[%s] [%s %s %s] %v @ [%v %v]", OpcodeName[i.opcode], i.operands[0].String(), i.operands[1].String(), i.operands[2].String(), i.opValid, i.Tag, i.slot)
}

// precondition for this logic is that when an instruction is created that all
// operand slots that are NOT being used, the corresponding operand valid bit is set to true
func (i *Instruction) Ready() bool { return i.opValid[0] && i.opValid[1] && i.opValid[2]}

/////////////////////////////////////////////////////////////////
// Modifiers

func (i *Instruction) Tag(tag Tag) { i.Tag = tag }
func (i *Instruction) SetOpcode(opcode Opcode) { i.opcode = opcode }
func (i *Instruction) SetOperand(slot uint8, operand Operand)  {
	i.operands[slot] = operand;
	i.opValid[slot] = true
}
func (i *Instruction) SetOperandValid(slot int)  { i.opValid[slot] = true }
func (i *Instruction) ResetOperandValid(slot int)  { i.opValid[slot] = false }
func (i *Instruction) SetValid(valid [3]bool) { i.opValid = valid }
