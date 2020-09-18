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

package lighthousehttp

import (
	"context"
	"encoding/json"
	"fmt"

	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

// AttestationData obtains attestation data for a slot.
func (s *Service) AttestationData(ctx context.Context, slot uint64, committeeIndex uint64) (*spec.AttestationData, error) {
	respBodyReader, err := s.get(ctx, fmt.Sprintf("/validator/attestation?slot=%d&committee_index=%d", slot, committeeIndex))
	if err != nil {
		return nil, errors.Wrap(err, "failed to request attestation")
	}

	specReader, err := lhToSpec(ctx, respBodyReader)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert lighthouse response to spec response")
	}

	var attestation *spec.Attestation
	if err := json.NewDecoder(specReader).Decode(&attestation); err != nil {
		return nil, errors.Wrap(err, "failed to parse attestation")
	}

	return attestation.Data, nil
}
