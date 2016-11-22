/**
 * File		:	$File: //depot/stillwater/dataflow/types/tag.go $
 *
 * Authors	:	E. Theodore L. Omtzigt
 * Date		:	21 April 2016
 *
 * Source Control Information:
 * Version	:	$Revision: #1 $
 * Latest	:	$Date: 2016/04/21 $
 * Location	:	$Id: //depot/stillwater/dataflow/types/tag.go#1 $
 *
 * Organization:
 *		Stillwater Supercomputing, Inc.
 *		P.O Box 720
 *		South Freeport, ME 04078-0720
 *
 * Copyright (c) 2006-2016 E. Theodore L. Omtzigt.  All rights reserved.
 *
 */
package types

import "fmt"

type Tag uint64

func (t Tag) String() string {
	return fmt.Sprintf("0X%08X", t)
}
