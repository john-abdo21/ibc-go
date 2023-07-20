package keeper_test

import (
	"encoding/hex"
	"os"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/types"
)

func (suite *KeeperTestSuite) TestQueryCode() {
	var req *types.QueryCodeRequest

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success",
			func() {
				signer := authtypes.NewModuleAddress(govtypes.ModuleName).String()
				code, err := os.ReadFile("../test_data/ics10_grandpa_cw.wasm.gz")
				suite.Require().NoError(err)
				msg := types.NewMsgStoreCode(signer, code)

				res, err := suite.chainA.GetSimApp().WasmClientKeeper.StoreCode(suite.chainA.GetContext(), msg)
				suite.Require().NoError(err)

				req = &types.QueryCodeRequest{CodeHash: hex.EncodeToString(res.Checksum)}
			},
			true,
		},
		{
			"fails with empty request",
			func() {
				req = &types.QueryCodeRequest{}
			},
			false,
		},
		{
			"fails with non-existent code hash",
			func() {
				req = &types.QueryCodeRequest{CodeHash: "test"}
			},
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			tc.malleate()

			res, err := suite.chainA.GetSimApp().WasmClientKeeper.Code(suite.chainA.GetContext(), req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().NotEmpty(res.Data)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryCodeHashes() {
	var expCodeHashes []string

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success with no code hashes",
			func() {
				expCodeHashes = []string{}
			},
			true,
		},
		{
			"success with one code hash",
			func() {
				signer := authtypes.NewModuleAddress(govtypes.ModuleName).String()
				code, err := os.ReadFile("../test_data/ics10_grandpa_cw.wasm.gz")
				suite.Require().NoError(err)
				msg := types.NewMsgStoreCode(signer, code)

				res, err := suite.chainA.GetSimApp().WasmClientKeeper.StoreCode(suite.chainA.GetContext(), msg)
				suite.Require().NoError(err)

				expCodeHashes = append(expCodeHashes, hex.EncodeToString(res.Checksum))
			},
			true,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			tc.malleate()

			req := &types.QueryCodeHashesRequest{}
			res, err := suite.chainA.GetSimApp().WasmClientKeeper.CodeHashes(suite.chainA.GetContext(), req)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().Equal(len(expCodeHashes), len(res.CodeHashes))
				suite.Require().ElementsMatch(expCodeHashes, res.CodeHashes)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}