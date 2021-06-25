package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"

	"github.com/spf13/cobra"
)

type MsgSignData struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Data    string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MsgSignData) Reset() {
	*m = MsgSignData{}
}

func (m *MsgSignData) String() string {
	return proto.CompactTextString(m)
}

func (m *MsgSignData) ProtoMessage() {
}

func CmdSignMessage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign [data]",
		Short: "sign data",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			//queryClient := types.NewQueryClient(clientCtx)
			argsData := string(args[0])
			_, err := sdk.AccAddressFromBech32(clientCtx.GetFromAddress().String())
			if err != nil {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
			}
			
			bz := codec.NewAminoCodec(codec.NewLegacyAmino()).MustMarshalJSON(&MsgSignData{Creator: clientCtx.GetFromAddress().String(), Data: argsData})


			return clientCtx.PrintString(string(sdk.MustSortJSON(bz)))
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
