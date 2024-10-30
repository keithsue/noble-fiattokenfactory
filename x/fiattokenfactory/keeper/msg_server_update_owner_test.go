// Copyright 2024 Circle Internet Group, Inc.  All rights reserved.
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
//
// SPDX-License-Identifier: Apache-2.0

package keeper_test

import (
	"testing"

	testkeeper "github.com/circlefin/noble-fiattokenfactory/testutil/keeper"
	"github.com/circlefin/noble-fiattokenfactory/testutil/sample"
	"github.com/circlefin/noble-fiattokenfactory/x/fiattokenfactory/keeper"
	"github.com/circlefin/noble-fiattokenfactory/x/fiattokenfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestUpdateOwner_OwnerNotSet(t *testing.T) {
	ftf, ctx := testkeeper.FiatTokenfactoryKeeper()
	msgServer := keeper.NewMsgServerImpl(ftf)

	_, err := msgServer.UpdateOwner(sdk.WrapSDKContext(ctx), &types.MsgUpdateOwner{})
	require.ErrorIs(t, err, types.ErrUserNotFound)
	require.ErrorContains(t, err, "owner is not set")
}

func TestUpdateOwner_FromAddressIsNotOwner(t *testing.T) {
	owner := sample.TestAccount()
	ftf, ctx := testkeeper.FiatTokenfactoryKeeper()
	msgServer := keeper.NewMsgServerImpl(ftf)
	ftf.SetOwner(ctx, types.Owner{Address: owner.Address})

	_, err := msgServer.UpdateOwner(sdk.WrapSDKContext(ctx), &types.MsgUpdateOwner{From: sample.AccAddress()})
	require.ErrorIs(t, err, types.ErrUnauthorized)
	require.ErrorContains(t, err, "you are not the owner")
}

func TestUpdateOwner_AddressAlreadyPrivileged(t *testing.T) {
	owner := sample.TestAccount()
	newOwner := sample.TestAccount()
	ftf, ctx := testkeeper.FiatTokenfactoryKeeper()
	msgServer := keeper.NewMsgServerImpl(ftf)
	ftf.SetOwner(ctx, types.Owner{Address: owner.Address})
	ftf.SetMasterMinter(ctx, types.MasterMinter{Address: newOwner.Address})

	_, err := msgServer.UpdateOwner(sdk.WrapSDKContext(ctx), &types.MsgUpdateOwner{From: owner.Address, Address: newOwner.Address})
	require.ErrorIs(t, err, types.ErrAlreadyPrivileged)
}

func TestUpdateOwner_Success(t *testing.T) {
	owner := sample.TestAccount()
	newOwner := sample.TestAccount()
	ftf, ctx := testkeeper.FiatTokenfactoryKeeper()
	msgServer := keeper.NewMsgServerImpl(ftf)
	ftf.SetOwner(ctx, types.Owner{Address: owner.Address})

	res, err := msgServer.UpdateOwner(sdk.WrapSDKContext(ctx), &types.MsgUpdateOwner{From: owner.Address, Address: newOwner.Address})
	require.NoError(t, err)
	require.Equal(t, &types.MsgUpdateOwnerResponse{}, res)
}