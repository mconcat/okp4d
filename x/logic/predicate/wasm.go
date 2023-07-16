package predicate

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ichiban/prolog/engine"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okp4/okp4d/x/logic/types"
	"github.com/okp4/okp4d/x/logic/util"
)

func NewWasmExtension(contractAddress sdk.AccAddress, name string) any {
	parts := strings.Split(name, "/")
	if len(parts) != 2 {
		return nil
	}

	arity := parts[1]
	switch arity {
	case "1":
		return func(vm *engine.VM, arg0 engine.Term, cont engine.Cont, env *engine.Env) *engine.Promise {
			return CosmWasmQuery(name, vm, contractAddress, []engine.Term{arg0}, cont, env)
		}
	case "2":
		return func(vm *engine.VM, arg0, arg1 engine.Term, cont engine.Cont, env *engine.Env) *engine.Promise {
			return CosmWasmQuery(name, vm, contractAddress, []engine.Term{arg0, arg1}, cont, env)
		}
	case "3":
		return func(vm *engine.VM, arg0, arg1, arg2 engine.Term, cont engine.Cont, env *engine.Env) *engine.Promise {
			return CosmWasmQuery(name, vm, contractAddress, []engine.Term{arg0, arg1, arg2}, cont, env)
		}
	default:
		return nil
	}
}

func solvePredicate(ctx sdk.Context, vm *engine.VM, wasm types.WasmKeeper, contractAddr sdk.AccAddress, predicateName string, termArgs []engine.Term, cont engine.Cont, env *engine.Env) (func(context.Context) *engine.Promise, error) {
	args := make([]types.WasmTerm, len(termArgs))
	for i, arg := range termArgs {
		switch arg := arg.(type) {
		case engine.Atom:
			atomstr := arg.String()
			args[i] = types.WasmTerm{Atom: &atomstr}
		case engine.Variable:
			varid := int64(arg)
			args[i] = types.WasmTerm{Var: &varid}
		}
	}

	msg := types.PrologQueryRequest{
		RunPredicate: &types.RunPredicateRequest{
			Name: predicateName,
			Args: args,
		},
	}
	bz, err := json.Marshal(msg)

	resbz, err := wasm.QuerySmart(ctx, contractAddr, bz)
	if err != nil {
		return nil, err
	}

	var res types.PrologQueryResponse
	err = json.Unmarshal(resbz, &res)
	if err != nil {
		return nil, err
	}

	return func(ctx context.Context) *engine.Promise {
		xs := make([]engine.Term, len(res.RunPredicate.Commands))
		ys := make([]engine.Term, len(res.RunPredicate.Commands))
		for i, command := range res.RunPredicate.Commands {
			xs[i] = command.Unify[0].ToTerm()
			ys[i] = command.Unify[1].ToTerm()
		}
		return engine.Unify(
			vm,
			Tuple(xs...),
			Tuple(ys...),
			cont,
			env,
		)
	}, nil
}

func CosmWasmQuery(predicate string, vm *engine.VM, contractAddress sdk.AccAddress, args []engine.Term, cont engine.Cont, env *engine.Env) *engine.Promise {
	return engine.Delay(func(ctx context.Context) *engine.Promise {
		sdkContext, err := util.UnwrapSDKContext(ctx)
		if err != nil {
			return engine.Error(err)
		}
		wasmKeeper := sdkContext.Value(types.CosmWasmKeeperContextKey).(types.WasmKeeper)

		unification, err := solvePredicate(sdkContext, vm, wasmKeeper, contractAddress, predicate, args, cont, env)
		if err != nil {
			return engine.Error(fmt.Errorf("%s: %w", predicate, err))
		}

		return engine.Delay(unification)
	})
}
