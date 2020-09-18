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

package prysmgrpc

import (
	"context"
	"strconv"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
)

// TargetAggregatorsPerCommittee provides the target number of aggregators for each attestation committee.
func (s *Service) TargetAggregatorsPerCommittee(ctx context.Context) (uint64, error) {
	if s.targetAggregatorsPerCommittee == nil {
		client := ethpb.NewBeaconChainClient(s.conn)
		config, err := client.GetBeaconConfig(ctx, &types.Empty{})
		if err != nil {
			return 0, errors.Wrap(err, "failed to obtain configuration")
		}

		val, exists := config.Config["TargetAggregatorsPerCommittee"]
		if !exists {
			return 0, errors.New("failed to obtain TargetAggregatorsPerCommittee value")
		}
		targetAggregatorsPerCommittee, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return 0, errors.Wrap(err, "failed to parse TargetAggregatorsPerCommittee value")
		}
		s.targetAggregatorsPerCommittee = &targetAggregatorsPerCommittee
	}
	return *s.targetAggregatorsPerCommittee, nil
}
