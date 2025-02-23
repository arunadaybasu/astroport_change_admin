package types

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgForceChangeAdmin defines the message for forcefully changing contract admin
type MsgForceChangeAdmin struct {
    Authority    string `json:"authority" yaml:"authority"`         // Bech32 address of the authority
    ContractAddr string `json:"contract_addr" yaml:"contract_addr"` // Bech32 contract address
    NewAdmin     string `json:"new_admin" yaml:"new_admin"`         // Bech32 new admin address
}

// NewMsgForceChangeAdmin creates a new MsgForceChangeAdmin message
func NewMsgForceChangeAdmin(authority, contractAddr, newAdmin string) *MsgForceChangeAdmin {
    return &MsgForceChangeAdmin{
        Authority:    authority,
        ContractAddr: contractAddr,
        NewAdmin:     newAdmin,
    }
}

// Route implements sdk.Msg and returns the message route
func (msg MsgForceChangeAdmin) Route() string { return "gov" }

// Type implements sdk.Msg and returns the message type
func (msg MsgForceChangeAdmin) Type() string { return "force_change_admin" }

// ValidateBasic performs basic validation on the message fields
func (msg MsgForceChangeAdmin) ValidateBasic() error {
    if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
        return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid authority address")
    }
    if _, err := sdk.AccAddressFromBech32(msg.ContractAddr); err != nil {
        return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid contract address")
    }
    if _, err := sdk.AccAddressFromBech32(msg.NewAdmin); err != nil {
        return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid new admin address")
    }
    return nil
}

// GetSigners defines whose signature is required for the message
func (msg MsgForceChangeAdmin) GetSigners() []sdk.AccAddress {
    addr, err := sdk.AccAddressFromBech32(msg.Authority)
    if err != nil {
        panic(err)
    }
    return []sdk.AccAddress{addr}
}
