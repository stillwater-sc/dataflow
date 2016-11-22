/**
 * File		:	$File: //depot/stillwater/dataflow/unit/istore.go $
 *
 * Authors	:	E. Theodore L. Omtzigt
 * Date		:	21 April 2016
 *
 * Source Control Information:
 * Version	:	$Revision: #1 $
 * Latest	:	$Date: 2016/04/21 $
 * Location	:	$Id: //depot/stillwater/dataflow/unit/istore.go#1 $
 *
 * Organization:
 *		Stillwater Supercomputing, Inc.
 *		P.O Box 720
 *		South Freeport, ME 04078-0720
 *
 * Copyright (c) 2006-2016 E. Theodore L. Omtzigt.  All rights reserved.
 *
 *
 *********************************************************************************
 *
 * The IStore is the Content Addressable Memory (CAM) of a processing element.
 * It matches data tokens to instruction tokens, and manages the life cycle of
 * the instruction token.
 *
 * The IStore data token to instruction token life cycle is:
 *      Creation: When a data token is the first right hand side to arrive, the iStore
 *                creates and allocates one or more instruction tokens.
 *      Matching: When subsequent data tokens arrive, the iStore matches them
 *                to their respective recurrence equation instruction token,
 *                and writes the data in the instruction token operand slot
 *		  assigned to the recurrence variable.
 *      Firing:   When all operands in the instruction token are valid,
 * 		  the iStore evicts the instruction token and injects the token
 *		  in the instruction queue.
 *
 */
package unit

import (
	"github.com/stillwater-sc/dataflow/types"
	"fmt"
	"github.com/golang/glog"
)

type IStore struct {
	iMap 	map[types.Tag]*types.Instruction
}

/////////////////////////////////////////////////////////////////
// Selectors

func (i *IStore) PendingInstructions() int { return len(i.iMap) }

/////////////////////////////////////////////////////////////////
// Modifiers

// don't use a max value of the IStore so that we can simulate
// arbitrary large programs.
func (i *IStore) Init() error {
	i.iMap = make(map[types.Tag]*types.Instruction)
}

/*
 * The iStore is the global program graph storage
 */
func (i *IStore) String() string {
	var content string
	for tag, pInstr := range i.iMap {
		content = content + fmt.Sprintf("[%08X] %p %v\n", tag, pInstr, *pInstr)
	}
}

/*
 * A MatchAdapter to isolate the type conversions needed from Packet to individual information
 * components to address the inputs to the IStore CAM functionality.
 *
 * RETURN:
 * a pointer to an instruction
 * a true/false ready status
 * an errorNr in case something went wrong
 */
func (i *IStore) MatchAdapter(packet *types.Packet) (*types.Instruction, bool, error) {
	var tag types.Tag = packet.Tag
	var operand types.Operand = packet.Payload
	// we need to get the slot assignment from the ProgramStore
	var slot types.RhsDescriptor = packet.Slot
	return i.Match(tag, operand, slot)
}

/*
 * Match() is the basic operator on the iStore CAM functionality
 * The client presents a disassembled data token through its Signature, its Operand, and operand slot
 * and the CAM will look up if there is a pending instruction token.
 * - If there is not, a new instruction is created and stored in the CAM.
 * - If there is, the operand is written to the selected operand slot in the existing instruction
 * After the Match, the instruction is checked to see if it has received all operands, and if yes,
 * it is deallocated from the CAM, and returned to the client to be forwarded to the execute phase.
 *
 * The CAM is also one possible source of errors. This is because the CAM must generate instructions.
 * This happens when the first operand arrives for a Recurrence Equation at a particular computational
 * index point. This instruction emission is the result of a lookup of the ReId of the operand
 * against the programs in the ProgramStore. If there is an inconsistency that causes this look-up
 * to fail, we need to make a decision how to handle this, and subsequent operands and other
 * recurrence equation processing.
 * Possible actions:
 * 1- HALT any further processing
 *	Halting would block all the recurrence equation processing that passes through this PE.
 *	This is a good solution for ICE debugging as all the state at the time of the error
 *	could be available for examination.
 * 2- If an INVALID_REID is found, simply ignore that specific instruction and continue
 *	This will cause all streams to continue, and the specific stream that is inconsistent
 *	with the ProgramStore information, to continuously throw this error.
 */
func (i *IStore) Match(tag types.Tag, operand types.Operand, slot types.RhsDescriptor) (*types.Instruction, bool, error) {
	// to turn on this logging for debugging programStore and token errors
	// add "-vmodule istore.go=10" to the cmd line
	glog.V(10).Infoln("Match input", "signature", tag, "operand", operand.String(), "slot", slot)
	var err error
	// set the entry in the iStore
	pInstr, ok := i.iMap[tag]
	if ok {
		pInstr.SetOperand(slot.Slot, operand)
		glog.V(10).Infoln("Updated existing instruction", "instruction", *pInstr)
	} else {
		// create a new instruction and store a pointer in the CAM
		// we'll receive an error when the ReId doesn't point to a valid program.
		pInstr, err = (*i.iGen).Instruction(tag, slot)
		if err == nil {
			pInstr.SetOperand(slot.Slot, operand)
			//		fmt.Printf("Created new instruction: %v at (%T = %p)\n", *pInstr,pInstr,pInstr)
			glog.V(10).Infoln("Created new instruction", *pInstr)
			i.iMap[tag] = pInstr
		} else {
			glog.Errorf("Instruction lookup failure: %v", err.Error())
		}
	}
	if pInstr.Ready() {
		delete (i.iMap, tag)
		glog.V(10).Infoln("Firing instruction", *pInstr)
		return pInstr, true, err
	}
	return nil, false, err
}
