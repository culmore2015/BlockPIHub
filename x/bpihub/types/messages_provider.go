package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateProvider{}

func NewMsgCreateProvider(creator string, name string, details string, website string) *MsgCreateProvider {
	return &MsgCreateProvider{
		Creator: creator,
		Name:    name,
		Details: details,
		Website: website,
	}
}

func (msg *MsgCreateProvider) Route() string {
	return RouterKey
}

func (msg *MsgCreateProvider) Type() string {
	return "CreateProvider"
}

func (msg *MsgCreateProvider) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateProvider) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateProvider) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateProvider{}

func NewMsgUpdateProvider(creator string, id uint64, name string, details string, website string) *MsgUpdateProvider {
	return &MsgUpdateProvider{
		Id:      id,
		Creator: creator,
		Name:    name,
		Details: details,
		Website: website,
	}
}

func (msg *MsgUpdateProvider) Route() string {
	return RouterKey
}

func (msg *MsgUpdateProvider) Type() string {
	return "UpdateProvider"
}

func (msg *MsgUpdateProvider) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateProvider) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateProvider) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateProvider{}

func NewMsgDeleteProvider(creator string, id uint64) *MsgDeleteProvider {
	return &MsgDeleteProvider{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteProvider) Route() string {
	return RouterKey
}

func (msg *MsgDeleteProvider) Type() string {
	return "DeleteProvider"
}

func (msg *MsgDeleteProvider) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteProvider) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteProvider) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
