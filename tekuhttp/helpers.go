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

package tekuhttp

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

// CurrentEpoch is a helper that calculates the current epoch.
func (s *Service) CurrentEpoch(ctx context.Context) (uint64, error) {
	genesisTime, err := s.GenesisTime(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to obtain genesis time for current epoch")
	}
	slotDuration, err := s.SlotDuration(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to obtain slot duration for current epoch")
	}
	slotsPerEpoch, err := s.SlotsPerEpoch(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to obtain slots per epoch for current epoch")
	}
	var currentEpoch uint64
	if genesisTime.After(time.Now()) {
		currentEpoch = 0
	} else {
		currentEpoch = uint64(time.Since(genesisTime).Seconds()) / (uint64(slotDuration.Seconds()) * slotsPerEpoch)
	}

	return currentEpoch, nil
}

// CurrentSlot is a helper that calculates the current slot.
func (s *Service) CurrentSlot(ctx context.Context) (uint64, error) {
	genesisTime, err := s.GenesisTime(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to obtain genesis time for current slot")
	}
	slotDuration, err := s.SlotDuration(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to obtain slot duration for current slot")
	}
	var currentSlot uint64
	if genesisTime.After(time.Now()) {
		currentSlot = 0
	} else {
		currentSlot = uint64(time.Since(genesisTime).Seconds()) / uint64(slotDuration.Seconds())
	}

	return currentSlot, nil
}
