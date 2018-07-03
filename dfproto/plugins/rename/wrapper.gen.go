// Package rename contains protocol buffer definitions for the rename DFHack plugin.
//
// See https://dfhack.readthedocs.io/en/latest/docs/Plugins.html#rename for more details.
package rename

import (
	"github.com/BenLubar/dfhackrpc"
	"github.com/BenLubar/dfhackrpc/dfproto/core"
)

// CallRenameSquad is a convenience wrapper around the rename::RenameSquad RPC method.
func CallRenameSquad(client *dfhackrpc.Client, request *RenameSquadIn) (dfhackrpc.CommandResult, error) {
	var response dfproto_core.EmptyMessage

	return client.CallPlugin("rename", "RenameSquad", request, &response)
}

// CallRenameUnit is a convenience wrapper around the rename::RenameUnit RPC method.
func CallRenameUnit(client *dfhackrpc.Client, request *RenameUnitIn) (dfhackrpc.CommandResult, error) {
	var response dfproto_core.EmptyMessage

	return client.CallPlugin("rename", "RenameUnit", request, &response)
}

// CallRenameBuilding is a convenience wrapper around the rename::RenameBuilding RPC method.
func CallRenameBuilding(client *dfhackrpc.Client, request *RenameBuildingIn) (dfhackrpc.CommandResult, error) {
	var response dfproto_core.EmptyMessage

	return client.CallPlugin("rename", "RenameBuilding", request, &response)
}
