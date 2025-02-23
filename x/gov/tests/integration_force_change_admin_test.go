package keeper_test

import (
    "testing"

    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/stretchr/testify/require"
    "github.com/terra-money/lunc_gov_force_change_admin_secure/x/gov/types"
    "github.com/terra-money/lunc_gov_force_change_admin_secure/x/gov/keeper"
)

// Mock context and keeper setup
func setupTestContext() (sdk.Context, keeper.Keeper) {
    // Placeholder: Implement mock context and keeper for integration testing
    var ctx sdk.Context
    var k keeper.Keeper
    return ctx, k
}

func TestHandleMsgForceChangeAdmin(t *testing.T) {
    ctx, k := setupTestContext()

    // Valid addresses for testing
    authority := "terra1authorityaddress"
    contractAddr := "terra1contractaddress"
    newAdmin := "terra1newadminaddress"

    // Create a valid message
    msg := types.NewMsgForceChangeAdmin(authority, contractAddr, newAdmin)

    // 1. Test successful admin change
    res, err := keeper.HandleMsgForceChangeAdmin(ctx, k, msg)
    require.NoError(t, err)
    require.NotNil(t, res)

    // 2. Attempt to reuse the message (should fail due to one-time use)
    res, err = keeper.HandleMsgForceChangeAdmin(ctx, k, msg)
    require.Error(t, err)
    require.Contains(t, err.Error(), "ForceChangeAdmin can only be used once")

    // 3. Unauthorized user attempt
    unauthorizedMsg := types.NewMsgForceChangeAdmin("terra1unauthorized", contractAddr, newAdmin)
    res, err = keeper.HandleMsgForceChangeAdmin(ctx, k, unauthorizedMsg)
    require.Error(t, err)
    require.Contains(t, err.Error(), "Only governance can perform this action")

    // 4. Invalid contract address
    invalidContractMsg := types.NewMsgForceChangeAdmin(authority, "invalid_address", newAdmin)
    res, err = keeper.HandleMsgForceChangeAdmin(ctx, k, invalidContractMsg)
    require.Error(t, err)
    require.Contains(t, err.Error(), "invalid contract address")

    // 5. Invalid new admin address
    invalidNewAdminMsg := types.NewMsgForceChangeAdmin(authority, contractAddr, "invalid_address")
    res, err = keeper.HandleMsgForceChangeAdmin(ctx, k, invalidNewAdminMsg)
    require.Error(t, err)
    require.Contains(t, err.Error(), "invalid new admin address")
}
