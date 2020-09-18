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

	client "github.com/attestantio/go-eth2-client"
	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
)

// Validators provides the validators, with their balance and status, for a given state.
// stateID can be a slot number or state root, or one of the special values "genesis", "head", "justified" or "finalized".
// validators is a list of validators to restrict the returned values.  If no validators are supplied no filter will be applied.
func (s *Service) ValidatorBalances(ctx context.Context, stateID string, validators []client.ValidatorIDProvider) (map[uint64]uint64, error) {
	beaconChainClient := ethpb.NewBeaconChainClient(s.conn)
	if beaconChainClient == nil {
		return nil, errors.New("failed to obtain beacon chain client")
	}

	validatorBalancesReq := &ethpb.ListValidatorBalancesRequest{
		PageSize: s.maxPageSize,
	}

	epoch, err := s.epochFromStateID(ctx, stateID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to obtain epoch from state ID")
	}
	if epoch == 0 {
		log.Trace().Msg("Fetching genesis validator balances")
		validatorBalancesReq.QueryFilter = &ethpb.ListValidatorBalancesRequest_Genesis{Genesis: true}
	} else {
		log.Trace().Uint64("epoch", epoch).Msg("Fetching epoch validator balances")
		validatorBalancesReq.QueryFilter = &ethpb.ListValidatorBalancesRequest_Epoch{Epoch: epoch}
	}

	res := make(map[uint64]uint64)

	pageToken := ""
	for i := int32(0); ; i += s.maxPageSize {
		log.Trace().Msg("Calling ListValidators()")
		validatorBalancesReq.PageToken = pageToken
		validatorBalancesResp, err := beaconChainClient.ListValidatorBalances(ctx, validatorBalancesReq)
		if err != nil {
			return nil, errors.Wrap(err, "failed to obtain validator balances")
		}
		if len(validatorBalancesResp.Balances) == 0 {
			break
		}

		for _, entry := range validatorBalancesResp.Balances {
			res[entry.Index] = entry.Balance
		}

		if validatorBalancesResp.NextPageToken == "" {
			break
		}
		pageToken = validatorBalancesResp.NextPageToken
	}

	return res, nil
}
