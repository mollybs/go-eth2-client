// Copyright © 2020 Attestant Limited.
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

package v1

import (
	"fmt"
	"strings"

	spec "github.com/attestantio/go-eth2-client/spec/phase0"
)

// ValidatorState defines the state of the validator.
type ValidatorState int

const (
	// ValidatorStateUnknown means no information can be found about the validator.
	ValidatorStateUnknown ValidatorState = iota
	// ValidatorStatePendingInitialized means the validator is not yet in the queue to be activated.
	ValidatorStatePendingInitialized
	// ValidatorStatePendingQueued means the validator is in the queue to be activated.
	ValidatorStatePendingQueued
	// ValidatorStateActiveOngoing means the validator is active.
	ValidatorStateActiveOngoing
	// ValidatorStateActiveExiting means the validator is active but exiting.
	ValidatorStateActiveExiting
	// ValidatorStateActiveSlashed means the validator is active but exiting due to being slashed.
	ValidatorStateActiveSlashed
	// ValidatorStateExitedUnslashed means the validator has exited without being slashed.
	ValidatorStateExitedUnslashed
	// ValidatorStateExitedSlashed means the validator has exited due to being slashed.
	ValidatorStateExitedSlashed
	// ValidatorStateWithdrawalPossible means it is possible to withdraw funds from the validator.
	ValidatorStateWithdrawalPossible
	// ValidatorStateWithdrawalPossible means funds have been withdrawn from the validator.
	ValidatorStateWithdrawalDone
)

var validatorStateStrings = [...]string{
	"Unknown",
	"Pending_initialized",
	"Pending_queued",
	"Active_ongoing",
	"Active_exiting",
	"Active_slashed",
	"Exited_unslashed",
	"Exited_slashed",
	"Withdrawal_possible",
	"Withdrawal_done",
}

// MarshalJSON implements json.Marshaler.
func (v *ValidatorState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", validatorStateStrings[*v])), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (v *ValidatorState) UnmarshalJSON(input []byte) error {
	var err error
	switch strings.ToLower(string(input)) {
	case `"unknown"`:
		*v = ValidatorStateUnknown
	case `"pending_initialized"`:
		*v = ValidatorStatePendingInitialized
	case `"pending_queued"`:
		*v = ValidatorStatePendingQueued
	case `"active_ongoing"`:
		*v = ValidatorStateActiveOngoing
	case `"active_exiting"`:
		*v = ValidatorStateActiveExiting
	case `"active_slashed"`:
		*v = ValidatorStateActiveSlashed
	case `"exited_unslashed"`:
		*v = ValidatorStateExitedUnslashed
	case `"exited_slashed"`:
		*v = ValidatorStateExitedSlashed
	case `"withdrawal_possible"`:
		*v = ValidatorStateWithdrawalPossible
	case `"withdrawal_done"`:
		*v = ValidatorStateWithdrawalDone
	default:
		*v = ValidatorStateUnknown
		err = fmt.Errorf("unrecognised validator state %s", string(input))
	}
	return err
}

func (v ValidatorState) String() string {
	return validatorStateStrings[v]
}

// IsAttesting returns true if the validator should be attesting.
func (v ValidatorState) IsAttesting() bool {
	return v == ValidatorStateActiveOngoing || v == ValidatorStateActiveExiting
}

// ValidatorToState is a helper that calculates the validator state given a validator struct.
func ValidatorToState(validator *spec.Validator, currentEpoch uint64, farFutureEpoch uint64) ValidatorState {
	if validator == nil {
		return ValidatorStateUnknown
	}

	if validator.ActivationEpoch > currentEpoch {
		// Pending.
		if validator.ActivationEligibilityEpoch == farFutureEpoch {
			return ValidatorStatePendingInitialized
		}
		return ValidatorStatePendingQueued
	}

	if validator.ActivationEpoch <= currentEpoch && currentEpoch < validator.ExitEpoch {
		// Active.
		if validator.ExitEpoch == farFutureEpoch {
			return ValidatorStateActiveOngoing
		}
		if validator.Slashed {
			return ValidatorStateActiveSlashed
		}
		return ValidatorStateActiveExiting
	}

	if validator.ExitEpoch <= currentEpoch && currentEpoch < validator.WithdrawableEpoch {
		// Exited.
		if validator.Slashed {
			return ValidatorStateExitedSlashed
		}
		return ValidatorStateExitedUnslashed
	}

	// Withdrawable.  No balance available so state possible.
	return ValidatorStateWithdrawalPossible
}
