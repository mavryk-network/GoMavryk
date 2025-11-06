// Copyright (c) 2020-2024 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package rpc

import "github.com/mavryk-network/gomavryk/mavryk"

// Ensure Endorsement implements the TypedOperation interface.
var _ TypedOperation = (*Endorsement)(nil)

// Endorsement represents an endorsement operation
type Endorsement struct {
	Generic
	Level          int64               `json:"level"`                 // <= v001+
	Endorsement    *InlinedEndorsement `json:"endorsement,omitempty"` // v001+
	Slot           int                 `json:"slot"`                  // v001+
	Round          int                 `json:"round"`                 // v001+
	PayloadHash    mavryk.PayloadHash  `json:"block_payload_hash"`    // v001+
	DalAttestation mavryk.Z            `json:"dal_attestation"`       // v002+
}

func (e Endorsement) GetLevel() int64 {
	if e.Endorsement != nil {
		return e.Endorsement.Operations.Level
	}
	return e.Level
}

// InlinedEndorsement represents and embedded endorsement
type InlinedEndorsement struct {
	Branch     mavryk.BlockHash `json:"branch"`     // the double block
	Operations Endorsement      `json:"operations"` // only level and kind are set
	Signature  mavryk.Signature `json:"signature"`
}
