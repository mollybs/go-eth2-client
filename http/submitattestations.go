// Copyright © 2020 - 2024 Attestant Limited.
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

package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	"github.com/attestantio/go-eth2-client/spec/phase0"
)

// SubmitAttestations submits attestations.
func (s *Service) SubmitAttestations(ctx context.Context, attestations []*phase0.Attestation) error {
	if err := s.assertIsSynced(ctx); err != nil {
		return err
	}

	specJSON, err := json.Marshal(attestations)
	if err != nil {
		return errors.Join(errors.New("failed to marshal JSON"), err)
	}

	_, err = s.post(ctx, "/eth/v1/beacon/pool/attestations", bytes.NewBuffer(specJSON))
	if err != nil {
		return errors.Join(errors.New("failed to submit beacon attestations"), err)
	}

	return nil
}
