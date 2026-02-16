// Copyright (c) 2020-2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package mavryk

import (
	"fmt"
)

type OpStatus byte

const (
	OpStatusInvalid OpStatus = iota // 0
	OpStatusApplied                 // 1 (success)
	OpStatusFailed
	OpStatusSkipped
	OpStatusBacktracked
)

func (t OpStatus) IsValid() bool {
	return t != OpStatusInvalid
}

func (t OpStatus) IsSuccess() bool {
	return t == OpStatusApplied
}

func (t *OpStatus) UnmarshalText(data []byte) error {
	v := ParseOpStatus(string(data))
	if !v.IsValid() {
		return fmt.Errorf("mavryk: invalid operation status '%s'", string(data))
	}
	*t = v
	return nil
}

func (t OpStatus) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func ParseOpStatus(s string) OpStatus {
	switch s {
	case "applied":
		return OpStatusApplied
	case "failed":
		return OpStatusFailed
	case "skipped":
		return OpStatusSkipped
	case "backtracked":
		return OpStatusBacktracked
	default:
		return OpStatusInvalid
	}
}

func (t OpStatus) String() string {
	switch t {
	case OpStatusApplied:
		return "applied"
	case OpStatusFailed:
		return "failed"
	case OpStatusSkipped:
		return "skipped"
	case OpStatusBacktracked:
		return "backtracked"
	default:
		return ""
	}
}

type OpType byte

// enums are allocated in chronological order
const (
	OpTypeInvalid                         OpType = iota
	OpTypeActivateAccount                        // 1
	OpTypeDoubleBakingEvidence                   // 2
	OpTypeDoubleEndorsementEvidence              // 3
	OpTypeSeedNonceRevelation                    // 4
	OpTypeTransaction                            // 5
	OpTypeOrigination                            // 6
	OpTypeDelegation                             // 7
	OpTypeReveal                                 // 8
	OpTypeEndorsement                            // 9
	OpTypeProposals                              // 10
	OpTypeBallot                                 // 11
	OpTypeFailingNoop                            // 12 v001
	OpTypeEndorsementWithSlot                    // 13 v001
	OpTypeRegisterConstant                       // 14 v001
	OpTypePreendorsement                         // 15 v001
	OpTypeDoublePreendorsementEvidence           // 16 v001
	OpTypeSetDepositsLimit                       // 17 v001 DEPRECATED in v002 (AI)
	OpTypeTransferTicket                         // 26 v001
	OpTypeVdfRevelation                          // 27 v001
	OpTypeIncreasePaidStorage                    // 28 v001
	OpTypeEvent                                  // 29 v001 (only in internal_operation_results)
	OpTypeDrainDelegate                          // 30 v001
	OpTypeUpdateConsensusKey                     // 31 v001
	OpTypeSmartRollupOriginate                   // 32 v001
	OpTypeSmartRollupAddMessages                 // 33 v001
	OpTypeSmartRollupCement                      // 34 v001
	OpTypeSmartRollupPublish                     // 35 v001
	OpTypeSmartRollupRefute                      // 36 v001
	OpTypeSmartRollupTimeout                     // 37 v001
	OpTypeSmartRollupExecuteOutboxMessage        // 38 v001
	OpTypeSmartRollupRecoverBond                 // 39 v001
	OpTypeDalPublishCommitment                   // 40 v002
	OpTypeAttestation                            // 41 v002 ??
	OpTypePreattestation                         // 42 v002
	OpTypeDoublePreattestationEvidence           // 43 v002
	OpTypeDoubleAttestationEvidence              // 44 v002
	OpTypeAttestationWithDal                     // 45 v002 ??
)

var (
	opTypeStrings = map[OpType]string{
		OpTypeInvalid:                         "",
		OpTypeActivateAccount:                 "activate_account",
		OpTypeDoubleBakingEvidence:            "double_baking_evidence",
		OpTypeDoubleEndorsementEvidence:       "double_endorsement_evidence",
		OpTypeSeedNonceRevelation:             "seed_nonce_revelation",
		OpTypeTransaction:                     "transaction",
		OpTypeOrigination:                     "origination",
		OpTypeDelegation:                      "delegation",
		OpTypeReveal:                          "reveal",
		OpTypeEndorsement:                     "endorsement",
		OpTypeProposals:                       "proposals",
		OpTypeBallot:                          "ballot",
		OpTypeFailingNoop:                     "failing_noop",
		OpTypeEndorsementWithSlot:             "endorsement_with_slot",
		OpTypeRegisterConstant:                "register_global_constant",
		OpTypePreendorsement:                  "preendorsement",
		OpTypeDoublePreendorsementEvidence:    "double_preendorsement_evidence",
		OpTypeSetDepositsLimit:                "set_deposits_limit",
		OpTypeTransferTicket:                  "transfer_ticket",
		OpTypeVdfRevelation:                   "vdf_revelation",
		OpTypeIncreasePaidStorage:             "increase_paid_storage",
		OpTypeEvent:                           "event",
		OpTypeDrainDelegate:                   "drain_delegate",
		OpTypeUpdateConsensusKey:              "update_consensus_key",
		OpTypeSmartRollupOriginate:            "smart_rollup_originate",
		OpTypeSmartRollupAddMessages:          "smart_rollup_add_messages",
		OpTypeSmartRollupCement:               "smart_rollup_cement",
		OpTypeSmartRollupPublish:              "smart_rollup_publish",
		OpTypeSmartRollupRefute:               "smart_rollup_refute",
		OpTypeSmartRollupTimeout:              "smart_rollup_timeout",
		OpTypeSmartRollupExecuteOutboxMessage: "smart_rollup_execute_outbox_message",
		OpTypeSmartRollupRecoverBond:          "smart_rollup_recover_bond",
		OpTypeDalPublishCommitment:            "dal_publish_commitment",
		OpTypeAttestation:                     "attestation",
		OpTypeAttestationWithDal:              "attestation_with_dal",
		OpTypePreattestation:                  "preattestation",
		OpTypeDoublePreattestationEvidence:    "double_preattestation_evidence",
		OpTypeDoubleAttestationEvidence:       "double_attestation_evidence",
	}
	opTypeReverseStrings = make(map[string]OpType)
)

func init() {
	for n, v := range opTypeStrings {
		opTypeReverseStrings[v] = n
	}
}

func (t OpType) IsValid() bool {
	return t != OpTypeInvalid
}

func (t *OpType) UnmarshalText(data []byte) error {
	v := ParseOpType(string(data))
	if !v.IsValid() {
		return fmt.Errorf("mavryk: invalid operation type '%s'", string(data))
	}
	*t = v
	return nil
}

func (t OpType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func ParseOpType(s string) OpType {
	t, ok := opTypeReverseStrings[s]
	if !ok {
		t = OpTypeInvalid
	}
	return t
}

func (t OpType) String() string {
	return opTypeStrings[t]
}

var (
	// Atlas v001
	opTagV2 = map[OpType]byte{
		OpTypeSeedNonceRevelation:             1,
		OpTypeDoubleEndorsementEvidence:       2,
		OpTypeDoubleBakingEvidence:            3,
		OpTypeActivateAccount:                 4,
		OpTypeProposals:                       5,
		OpTypeBallot:                          6,
		OpTypeReveal:                          107, // v001
		OpTypeTransaction:                     108, // v001
		OpTypeOrigination:                     109, // v001
		OpTypeDelegation:                      110, // v001
		OpTypeFailingNoop:                     17,  // v001
		OpTypeRegisterConstant:                111, // v001
		OpTypePreendorsement:                  20,  // v001
		OpTypeEndorsement:                     21,  // v001
		OpTypeAttestationWithDal:              23,  // v002
		OpTypeDoublePreendorsementEvidence:    7,   // v001
		OpTypeSetDepositsLimit:                112, // v001
		OpTypeTransferTicket:                  158, // v001
		OpTypeVdfRevelation:                   8,   // v001
		OpTypeIncreasePaidStorage:             113, // v001
		OpTypeDrainDelegate:                   9,   // v001
		OpTypeUpdateConsensusKey:              114, // v001
		OpTypeSmartRollupOriginate:            200, // v001
		OpTypeSmartRollupAddMessages:          201, // v001
		OpTypeSmartRollupCement:               202, // v001
		OpTypeSmartRollupPublish:              203, // v001
		OpTypeSmartRollupRefute:               204, // v001
		OpTypeSmartRollupTimeout:              205, // v001
		OpTypeSmartRollupExecuteOutboxMessage: 206, // v001
		OpTypeSmartRollupRecoverBond:          207, // v001
	}
	// Boreas v002 and up
	opTagV3 = map[OpType]byte{
		OpTypeSeedNonceRevelation:             1,
		OpTypeDoubleAttestationEvidence:       2, // v002
		OpTypeDoubleBakingEvidence:            3,
		OpTypeActivateAccount:                 4,
		OpTypeProposals:                       5,
		OpTypeBallot:                          6,
		OpTypeReveal:                          107, // v001
		OpTypeTransaction:                     108, // v001
		OpTypeOrigination:                     109, // v001
		OpTypeDelegation:                      110, // v001
		OpTypeFailingNoop:                     17,  // v001
		OpTypeRegisterConstant:                111, // v001
		OpTypePreattestation:                  20,  // v002
		OpTypeAttestation:                     21,  // v002
		OpTypeAttestationWithDal:              23,  // v002
		OpTypeDoublePreattestationEvidence:    7,   // v002
		OpTypeSetDepositsLimit:                112, // v001
		OpTypeTransferTicket:                  158, // v001
		OpTypeVdfRevelation:                   8,   // v001
		OpTypeIncreasePaidStorage:             113, // v001
		OpTypeDrainDelegate:                   9,   // v001
		OpTypeUpdateConsensusKey:              114, // v001
		OpTypeSmartRollupOriginate:            200, // v001
		OpTypeSmartRollupAddMessages:          201, // v001
		OpTypeSmartRollupCement:               202, // v001
		OpTypeSmartRollupPublish:              203, // v001
		OpTypeSmartRollupRefute:               204, // v001
		OpTypeSmartRollupTimeout:              205, // v001
		OpTypeSmartRollupExecuteOutboxMessage: 206, // v001
		OpTypeSmartRollupRecoverBond:          207, // v001
		OpTypeDalPublishCommitment:            230, // v001 FIXME: is this correct?
	}
)

func (t OpType) TagVersion(ver int) byte {
	var (
		tag byte
		ok  bool
	)
	switch ver {
	case 2:
		tag, ok = opTagV2[t]
	case 3:
		tag, ok = opTagV3[t]
	default:
		tag, ok = opTagV2[t]
	}
	if !ok {
		return 255
	}
	return tag
}


func (t OpType) Tag() byte {
	tag, ok := opTagV2[t]
	if !ok {
		tag = 255
	}
	return tag
}

var (
	// Atlas v001 and up
	opMinSizeV2 = map[byte]int{
		1:   37,                       // OpTypeSeedNonceRevelation
		2:   9 + 2*(32+43+64),         // OpTypeDoubleEndorsementEvidence
		3:   9 + 2*237,                // OpTypeDoubleBakingEvidence (w/o seed_nonce_hash, min fitness size)
		4:   41,                       // OpTypeActivateAccount
		5:   30,                       // OpTypeProposals
		6:   59,                       // OpTypeBallot
		107: 26 + 32,                  // OpTypeReveal // v001 (assuming shortest pk)
		108: 50,                       // OpTypeTransaction // v001
		109: 28,                       // OpTypeOrigination // v001
		110: 27,                       // OpTypeDelegation // v001
		17:  5,                        // OpTypeFailingNoop  // v001
		111: 30,                       // OpTypeRegisterConstant // v001
		7:   9 + 2*(32+43+64),         // OpTypeDoublePreendorsementEvidence // v001
		20:  43,                       // OpTypePreendorsement // v001
		21:  43,                       // OpTypeEndorsement // v001
		112: 27,                       // OpTypeSetDepositsLimit // v001
		8:   201,                      // OpTypeVdfRevelation // v001
		113: 27 + 22,                  // OpTypeIncreasePaidStorage // v001
		9:   1 + 3*21,                 // OpTypeDrainDelegate // v001
		114: 26 + 32,                  // OpTypeUpdateConsensusKey // v001
		158: 26 + 8 + 22 + 1 + 22 + 4, // OpTypeTransferTicket // v001
		200: 26 + 13,                  // OpTypeSmartRollupOriginate // v001
		201: 26 + 4,                   // OpTypeSmartRollupAddMessages // v001
		202: 26 + 52,                  // OpTypeSmartRollupCement // v001
		203: 26 + 96,                  // OpTypeSmartRollupPublish // v001
		204: 26 + 41,                  // OpTypeSmartRollupRefute // v001
		205: 26 + 62,                  // OpTypeSmartRollupTimeout // v001
		206: 26 + 56,                  // OpTypeSmartRollupExecuteOutboxMessage // v001
		207: 26 + 41,                  // OpTypeSmartRollupRecoverBond // v001

		// FIXME:
		230: 26 + 101, // OpTypeDalPublishCommitment // v002
		23:  43 + 1,   // OpTypeAttestationWithDal // v002 (assuming smallest slot number)
	}
	// Boreas v002 and up
	opMinSizeV3 = map[byte]int{
		1:   37,                       // OpTypeSeedNonceRevelation
		2:   9 + 2*(32+43+64),         // OpTypeDoubleAttestationEvidence
		3:   9 + 2*237,                // OpTypeDoubleBakingEvidence (w/o seed_nonce_hash, min fitness size)
		4:   41,                       // OpTypeActivateAccount
		5:   30,                       // OpTypeProposals
		6:   59,                       // OpTypeBallot
		107: 26 + 32,                  // OpTypeReveal // v001 (assuming shortest pk)
		108: 50,                       // OpTypeTransaction // v001
		109: 28,                       // OpTypeOrigination // v001
		110: 27,                       // OpTypeDelegation // v001
		17:  5,                        // OpTypeFailingNoop  // v001
		111: 30,                       // OpTypeRegisterConstant // v001
		7:   9 + 2*(32+43+64),         // OpTypeDoublePreattestationEvidence // v002
		20:  43,                       // OpTypePreattestation // v002
		21:  43,                       // OpTypeAttestation // v002
		23:  43 + 1,                   // OpTypeAttestationWithDal // v002 (assuming smallest slot number)
		112: 27,                       // OpTypeSetDepositsLimit // v001
		8:   201,                      // OpTypeVdfRevelation // v001
		113: 27 + 22,                  // OpTypeIncreasePaidStorage // v001
		9:   1 + 3*21,                 // OpTypeDrainDelegate // v001
		114: 26 + 32,                  // OpTypeUpdateConsensusKey // v001
		158: 26 + 8 + 22 + 1 + 22 + 4, // OpTypeTransferTicket // v001
		200: 26 + 13,                  // OpTypeSmartRollupOriginate // v001
		201: 26 + 4,                   // OpTypeSmartRollupAddMessages // v001
		202: 26 + 52,                  // OpTypeSmartRollupCement // v001
		203: 26 + 96,                  // OpTypeSmartRollupPublish // v001
		204: 26 + 41,                  // OpTypeSmartRollupRefute // v001
		205: 26 + 62,                  // OpTypeSmartRollupTimeout // v001
		206: 26 + 56,                  // OpTypeSmartRollupExecuteOutboxMessage // v001
		207: 26 + 41,                  // OpTypeSmartRollupRecoverBond // v001
		230: 26 + 101,                 // OpTypeDalPublishCommitment // v002
	}
)

func (t OpType) MinSizeVersion(ver int) int {
	switch ver {
	case 2:
		return opMinSizeV2[t.TagVersion(ver)]
	default:
		return opMinSizeV3[t.TagVersion(ver)]
	}
}

func (t OpType) MinSize() int {
	return opMinSizeV3[t.Tag()]
}

func (t OpType) ListId() int {
	switch t {
	case OpTypeEndorsement, OpTypeEndorsementWithSlot, OpTypePreendorsement,
		OpTypeAttestation, OpTypePreattestation, OpTypeAttestationWithDal:
		return 0
	case OpTypeProposals, OpTypeBallot:
		return 1
	case OpTypeActivateAccount,
		OpTypeDoubleBakingEvidence,
		OpTypeDoubleEndorsementEvidence,
		OpTypeSeedNonceRevelation,
		OpTypeDoublePreendorsementEvidence,
		OpTypeVdfRevelation,
		OpTypeDrainDelegate,
		OpTypeDoubleAttestationEvidence,
		OpTypeDoublePreattestationEvidence:
		return 2
	case OpTypeTransaction, // generic user operations
		OpTypeOrigination,
		OpTypeDelegation,
		OpTypeReveal,
		OpTypeRegisterConstant,
		OpTypeSetDepositsLimit,
		OpTypeTransferTicket,
		OpTypeUpdateConsensusKey,
		OpTypeSmartRollupOriginate,
		OpTypeSmartRollupAddMessages,
		OpTypeSmartRollupCement,
		OpTypeSmartRollupPublish,
		OpTypeSmartRollupRefute,
		OpTypeSmartRollupTimeout,
		OpTypeSmartRollupExecuteOutboxMessage,
		OpTypeSmartRollupRecoverBond,
		OpTypeDalPublishCommitment:
		return 3
	default:
		return -1 // invalid
	}
}

func ParseOpTag(t byte) OpType {
	switch t {
	case 0:
		return OpTypeAttestation
	case 1:
		return OpTypeSeedNonceRevelation
	case 2:
		return OpTypeDoubleAttestationEvidence
	case 3:
		return OpTypeDoubleBakingEvidence
	case 4:
		return OpTypeActivateAccount
	case 5:
		return OpTypeProposals
	case 6:
		return OpTypeBallot
	case 7:
		return OpTypeDoublePreattestationEvidence
	case 8:
		return OpTypeVdfRevelation
	case 9:
		return OpTypeDrainDelegate
	case 10:
		return OpTypeEndorsementWithSlot
	case 17:
		return OpTypeFailingNoop
	case 20:
		return OpTypePreattestation
	case 21:
		return OpTypeEndorsement
	case 23:
		return OpTypeAttestationWithDal
	case 107:
		return OpTypeReveal
	case 108:
		return OpTypeTransaction
	case 109:
		return OpTypeOrigination
	case 110:
		return OpTypeDelegation
	case 111:
		return OpTypeRegisterConstant
	case 112:
		return OpTypeSetDepositsLimit
	case 113:
		return OpTypeIncreasePaidStorage
	case 114:
		return OpTypeUpdateConsensusKey
	case 158:
		return OpTypeTransferTicket
	case 200:
		return OpTypeSmartRollupOriginate
	case 201:
		return OpTypeSmartRollupAddMessages
	case 202:
		return OpTypeSmartRollupCement
	case 203:
		return OpTypeSmartRollupPublish
	case 204:
		return OpTypeSmartRollupRefute
	case 205:
		return OpTypeSmartRollupTimeout
	case 206:
		return OpTypeSmartRollupExecuteOutboxMessage
	case 207:
		return OpTypeSmartRollupRecoverBond
	case 230:
		return OpTypeDalPublishCommitment
	default:
		return OpTypeInvalid
	}
}

func ParseOpTagVersion(t byte, ver int) OpType {
	tags := opTagV2
	switch ver {
	case 2:
		tags = opTagV2
	case 3:
		tags = opTagV3
	}
	for typ, tag := range tags {
		if tag == t {
			return typ
		}
	}
	return OpTypeInvalid
}
