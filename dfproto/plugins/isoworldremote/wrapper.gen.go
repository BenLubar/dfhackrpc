// Package isoworldremote contains protocol buffer definitions for the isoworldremote DFHack plugin.
//
// See https://dfhack.readthedocs.io/en/latest/docs/Plugins.html#isoworldremote for more details.
package isoworldremote

import (
	"github.com/BenLubar/dfhackrpc"
)

// CallGetEmbarkTile is a convenience wrapper around the isoworldremote::GetEmbarkTile RPC method.
func CallGetEmbarkTile(client *dfhackrpc.Client, request *TileRequest) (*EmbarkTile, dfhackrpc.CommandResult, error) {
	var response EmbarkTile

	rv, err := client.CallPlugin("isoworldremote", "GetEmbarkTile", request, &response)

	return &response, rv, err
}

// CallGetEmbarkInfo is a convenience wrapper around the isoworldremote::GetEmbarkInfo RPC method.
func CallGetEmbarkInfo(client *dfhackrpc.Client, request *MapRequest) (*MapReply, dfhackrpc.CommandResult, error) {
	var response MapReply

	rv, err := client.CallPlugin("isoworldremote", "GetEmbarkInfo", request, &response)

	return &response, rv, err
}

// CallGetRawNames is a convenience wrapper around the isoworldremote::GetRawNames RPC method.
func CallGetRawNames(client *dfhackrpc.Client, request *MapRequest) (*RawNames, dfhackrpc.CommandResult, error) {
	var response RawNames

	rv, err := client.CallPlugin("isoworldremote", "GetRawNames", request, &response)

	return &response, rv, err
}
