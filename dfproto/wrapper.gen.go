package dfproto

import (
	"github.com/BenLubar/dfhackrpc"
	"github.com/BenLubar/dfhackrpc/dfproto/core"
)

// CallGetVersion is a convenience wrapper around the GetVersion RPC method.
func CallGetVersion(client *dfhackrpc.Client) (string, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response dfproto_core.StringMessage

	rv, err := client.Call("GetVersion", &request, &response)

	return response.GetValue(), rv, err
}

// CallGetDFVersion is a convenience wrapper around the GetDFVersion RPC method.
func CallGetDFVersion(client *dfhackrpc.Client) (string, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response dfproto_core.StringMessage

	rv, err := client.Call("GetDFVersion", &request, &response)

	return response.GetValue(), rv, err
}

// CallGetWorldInfo is a convenience wrapper around the GetWorldInfo RPC method.
func CallGetWorldInfo(client *dfhackrpc.Client) (*GetWorldInfoOut, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response GetWorldInfoOut

	rv, err := client.Call("GetWorldInfo", &request, &response)

	return &response, rv, err
}

// CallListEnums is a convenience wrapper around the ListEnums RPC method.
func CallListEnums(client *dfhackrpc.Client) (*ListEnumsOut, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response ListEnumsOut

	rv, err := client.Call("ListEnums", &request, &response)

	return &response, rv, err
}

// CallListJobSkills is a convenience wrapper around the ListJobSkills RPC method.
func CallListJobSkills(client *dfhackrpc.Client) (*ListJobSkillsOut, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response ListJobSkillsOut

	rv, err := client.Call("ListJobSkills", &request, &response)

	return &response, rv, err
}

// CallListMaterials is a convenience wrapper around the ListMaterials RPC method.
func CallListMaterials(client *dfhackrpc.Client, request *ListMaterialsIn) (*ListMaterialsOut, dfhackrpc.CommandResult, error) {
	var response ListMaterialsOut

	rv, err := client.Call("ListMaterials", request, &response)

	return &response, rv, err
}

// CallListUnits is a convenience wrapper around the ListUnits RPC method.
func CallListUnits(client *dfhackrpc.Client, request *ListUnitsIn) (*ListUnitsOut, dfhackrpc.CommandResult, error) {
	var response ListUnitsOut

	rv, err := client.Call("ListUnits", request, &response)

	return &response, rv, err
}

// CallListSquads is a convenience wrapper around the ListSquads RPC method.
func CallListSquads(client *dfhackrpc.Client, request *ListSquadsIn) (*ListSquadsOut, dfhackrpc.CommandResult, error) {
	var response ListSquadsOut

	rv, err := client.Call("ListSquads", request, &response)

	return &response, rv, err
}

// CallSetUnitLabors is a convenience wrapper around the SetUnitLabors RPC method.
func CallSetUnitLabors(client *dfhackrpc.Client, request *SetUnitLaborsIn) (dfhackrpc.CommandResult, error) {
	var response dfproto_core.EmptyMessage

	return client.Call("SetUnitLabors", request, &response)
}
