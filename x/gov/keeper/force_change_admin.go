package keeper

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
)

// ForceChangeAdmin directly changes the admin of a smart contract
func (k Keeper) ForceChangeAdmin(ctx sdk.Context, contractAddr sdk.AccAddress, newAdmin sdk.AccAddress) error {
    store := ctx.KVStore(k.storeKey)
    adminKey := append([]byte("wasm/admin/"), contractAddr.Bytes()...)

    // Update the admin in KVStore
    store.Set(adminKey, newAdmin.Bytes())

    // Emit an event for transparency
    ctx.EventManager().EmitEvent(
        sdk.NewEvent("admin_changed",
            sdk.NewAttribute("contract", contractAddr.String()),
            sdk.NewAttribute("new_admin", newAdmin.String()),
        ),
    )

    return nil
}
