package keeper

import (
    "fmt"

    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
    "github.com/terra-money/lunc_gov_force_change_admin_secure/x/gov/types"
)

var KeyForceChangeAdminUsed = []byte("force_change_admin_used")

// HandleMsgForceChangeAdmin processes the MsgForceChangeAdmin message securely
func HandleMsgForceChangeAdmin(ctx sdk.Context, keeper Keeper, msg *types.MsgForceChangeAdmin) (*sdk.Result, error) {
    store := ctx.KVStore(keeper.storeKey)

    // Check if ForceChangeAdmin has already been used
    if store.Has(KeyForceChangeAdminUsed) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "ForceChangeAdmin can only be used once")
    }

    // Verify sender is governance authority
    if !keeper.IsGovernanceAuthority(ctx, msg.Authority) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Only governance can perform this action")
    }

    // Validate addresses
    if msg.NewAdmin == "" {
        return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "new admin address cannot be empty")
    }

    if !keeper.ContractExists(ctx, msg.ContractAddr) {
        return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "contract does not exist")
    }

    // Perform the admin change
    err := keeper.ForceChangeAdmin(ctx, msg.ContractAddr, msg.NewAdmin)
    if err != nil {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Failed to change admin")
    }

    // Mark as used to ensure one-time use with block height metadata
    store.Set(KeyForceChangeAdminUsed, []byte(fmt.Sprintf("used_at_height_%d", ctx.BlockHeight())))

    // Emit event for transparency
    ctx.EventManager().EmitEvent(
        sdk.NewEvent("force_admin_change",
            sdk.NewAttribute("contract", msg.ContractAddr),
            sdk.NewAttribute("new_admin", msg.NewAdmin),
            sdk.NewAttribute("executed_at_height", fmt.Sprintf("%d", ctx.BlockHeight())),
        ),
    )

    // Log the action
    logger := ctx.Logger().With("module", "force_change_admin")
    logger.Info("Admin change executed", "contract", msg.ContractAddr, "new_admin", msg.NewAdmin, "executed_by", msg.Authority)

    return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
