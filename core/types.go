// Copyright © 2021 Obol Technologies Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"encoding/hex"
	"fmt"

	"github.com/obolnetwork/charon/app/errors"
)

// DutyType enumerates the different types of duties.
type DutyType int

const (
	DutyUnknown = DutyType(iota)
	DutyAttester
	DutyProposer
	dutySentinal // Must always be last
)

func (d DutyType) Valid() bool {
	return d > DutyUnknown && d < dutySentinal
}

func (d DutyType) String() string {
	return map[DutyType]string{
		DutyUnknown:  "unknown",
		DutyAttester: "attester",
		DutyProposer: "proposer",
	}[d]
}

// AllDutyTypes returns a list of all valid duty types.
func AllDutyTypes() []DutyType {
	var resp []DutyType
	for i := DutyUnknown + 1; i.Valid(); i++ {
		resp = append(resp, i)
	}

	return resp
}

// Duty is the unit of work of the core workflow.
type Duty struct {
	// Slot is the Ethereum consensus layer slot.
	Slot int64
	// Type is the duty type performed in the slot.
	Type DutyType
}

func (d Duty) String() string {
	return fmt.Sprintf("%d/%s", d.Slot, d.Type)
}

const pkLen = 98 // "0x" + hex.Encode([48]byte) = 2+2*48

// NewPubKeyFromBytes returns a new public key from raw bytes.
func NewPubKeyFromBytes(bytes []byte) (PubKey, error) {
	pk := PubKey("0x" + hex.EncodeToString(bytes))
	if len(pk) != pkLen {
		return "", errors.New("invalid public key length")
	}

	return pk, nil
}

// PubKey is the DV root public key, the identifier of a validator in the core workflow.
// It is a hex formatted string, e.g. "0xb82bc680e...".
type PubKey string

// String returns a concise logging friendly version of the public key, e.g. "b82_97f".
func (k PubKey) String() string {
	if len(k) != pkLen {
		return "<invalid public key:" + string(k) + ">"
	}

	return string(k[2:5]) + "_" + string(k[94:97])
}

// Bytes returns the public key as raw bytes.
func (k PubKey) Bytes() ([]byte, error) {
	if len(k) != pkLen {
		return nil, errors.New("invalid public key length")
	}

	b, err := hex.DecodeString(string(k[2:]))
	if err != nil {
		return nil, errors.Wrap(err, "decode public key hex")
	}

	return b, nil
}

// FetchArg contains the arguments required to fetch the duty data,
// it is the result of resolving duties at the start of an epoch.
type FetchArg []byte

// FetchArgSet is a set of fetch args, one per validator.
type FetchArgSet map[PubKey]FetchArg