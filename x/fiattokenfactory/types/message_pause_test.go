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

package types

import (
	"testing"

	"github.com/circlefin/noble-fiattokenfactory/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgPause_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgPause
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgPause{
				From: "invalid_address",
			},
			err: ErrInvalidAddress,
		},
		{
			name: "happy path",
			msg: MsgPause{
				From: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
