//nolint:gocognit
package predicate
/*
import (
	"fmt"
	"testing"
	"encoding/json"

	"github.com/golang/mock/gomock"
	"github.com/ichiban/prolog/engine"

	. "github.com/smartystreets/goconvey/convey"

	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/okp4/okp4d/x/logic/testutil"
	"github.com/okp4/okp4d/x/logic/types"
)

type MockWasmKeeper struct {
	ChainID string
	Address sdk.AccAddress
}

func (mock MockWasmKeeper) QuerySmart(ctx sdk.Context, contractAddr sdk.AccAddress, req []byte) ([]byte, error) {
	if contractAddr != mock.Address {
		return nil, fmt.Errorf("invalid contract address")
	}

	return mock.Query(req), nil
}

func (mock MockWasmKeeper) Query(reqbz []byte) []byte {
	var req types.PrologQueryRequest 
	json.Unmarshal(reqbz, &req)
	fmt.Println(req)
	if req.PrologExtensionManifest != nil {
		resp := types.PrologExtensionManifestResponse {
			Predicates: []types.PredicateManifest{
				{
					Address: mock.Address,
					Name: "chain_id/1",
					Cost: 10,
				}
			},
		}
		respbz, _ := json.Marshal(resp)
		return respbz		
	}

	if req.RunPredicate != nil {
		switch req.RunPredicate.Name {
		case "chain_id/1":
			resp := types.PrologQueryResponse{
				RunPredicate: &types.RunPredicateResponse{
					Commands: []types.Command{
						Unify: []types.WasmTerm{
							req.RunPredicate.Args[0],
							{ Atom: mock.ChainID },
						}
					}
				}
			}
			respbz, _ := json.Marshal(resp)
			return respbz
		}
	}
	
	return nil
}
*/