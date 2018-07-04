// Package remotefortressreader contains protocol buffer definitions for the RemoteFortressReader DFHack plugin.
//
// See https://dfhack.readthedocs.io/en/latest/docs/Plugins.html#remotefortressreader for more details.
package remotefortressreader

import (
	"github.com/BenLubar/dfhackrpc"
	"github.com/BenLubar/dfhackrpc/dfproto/core"
)

// CallGetMaterialList is a convenience wrapper around the RemoteFortressReader::GetMaterialList RPC method.
func CallGetMaterialList(client *dfhackrpc.Client) ([]*MaterialDefinition, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response MaterialList

	rv, err := client.CallPlugin("RemoteFortressReader", "GetMaterialList", &request, &response)

	return response.GetMaterialList(), rv, err
}

// CallGetGrowthList is a convenience wrapper around the RemoteFortressReader::GetGrowthList RPC method.
func CallGetGrowthList(client *dfhackrpc.Client) ([]*MaterialDefinition, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response MaterialList

	rv, err := client.CallPlugin("RemoteFortressReader", "GetGrowthList", &request, &response)

	return response.GetMaterialList(), rv, err
}

// CallGetBlockList is a convenience wrapper around the RemoteFortressReader::GetBlockList RPC method.
func CallGetBlockList(client *dfhackrpc.Client, request *BlockRequest) (*BlockList, dfhackrpc.CommandResult, error) {
	var response BlockList

	rv, err := client.CallPlugin("RemoteFortressReader", "GetBlockList", request, &response)

	return &response, rv, err
}

// CallCheckHashes is a convenience wrapper around the RemoteFortressReader::CheckHashes RPC method.
func CallCheckHashes(client *dfhackrpc.Client) (dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response dfproto_core.EmptyMessage

	return client.CallPlugin("RemoteFortressReader", "CheckHashes", &request, &response)
}

// CallGetTiletypeList is a convenience wrapper around the RemoteFortressReader::GetTiletypeList RPC method.
func CallGetTiletypeList(client *dfhackrpc.Client) ([]*Tiletype, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response TiletypeList

	rv, err := client.CallPlugin("RemoteFortressReader", "GetTiletypeList", &request, &response)

	return response.GetTiletypeList(), rv, err
}

// CallGetPlantList is a convenience wrapper around the RemoteFortressReader::GetPlantList RPC method.
func CallGetPlantList(client *dfhackrpc.Client, request *BlockRequest) ([]*PlantDef, dfhackrpc.CommandResult, error) {
	var response PlantList

	rv, err := client.CallPlugin("RemoteFortressReader", "GetPlantList", request, &response)

	return response.GetPlantList(), rv, err
}

// CallGetUnitList is a convenience wrapper around the RemoteFortressReader::GetUnitList RPC method.
func CallGetUnitList(client *dfhackrpc.Client) ([]*UnitDefinition, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response UnitList

	rv, err := client.CallPlugin("RemoteFortressReader", "GetUnitList", &request, &response)

	return response.GetCreatureList(), rv, err
}

// CallGetUnitListInside is a convenience wrapper around the RemoteFortressReader::GetUnitListInside RPC method.
func CallGetUnitListInside(client *dfhackrpc.Client, request *BlockRequest) ([]*UnitDefinition, dfhackrpc.CommandResult, error) {
	var response UnitList

	rv, err := client.CallPlugin("RemoteFortressReader", "GetUnitListInside", request, &response)

	return response.GetCreatureList(), rv, err
}

// CallGetViewInfo is a convenience wrapper around the RemoteFortressReader::GetViewInfo RPC method.
func CallGetViewInfo(client *dfhackrpc.Client) (*ViewInfo, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response ViewInfo

	rv, err := client.CallPlugin("RemoteFortressReader", "GetViewInfo", &request, &response)

	return &response, rv, err
}

// CallGetMapInfo is a convenience wrapper around the RemoteFortressReader::GetMapInfo RPC method.
func CallGetMapInfo(client *dfhackrpc.Client) (*MapInfo, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response MapInfo

	rv, err := client.CallPlugin("RemoteFortressReader", "GetMapInfo", &request, &response)

	return &response, rv, err
}

// CallResetMapHashes is a convenience wrapper around the RemoteFortressReader::ResetMapHashes RPC method.
func CallResetMapHashes(client *dfhackrpc.Client) (dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response dfproto_core.EmptyMessage

	return client.CallPlugin("RemoteFortressReader", "ResetMapHashes", &request, &response)
}

// CallGetItemList is a convenience wrapper around the RemoteFortressReader::GetItemList RPC method.
func CallGetItemList(client *dfhackrpc.Client) ([]*MaterialDefinition, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response MaterialList

	rv, err := client.CallPlugin("RemoteFortressReader", "GetItemList", &request, &response)

	return response.GetMaterialList(), rv, err
}

// CallGetBuildingDefList is a convenience wrapper around the RemoteFortressReader::GetBuildingDefList RPC method.
func CallGetBuildingDefList(client *dfhackrpc.Client) ([]*BuildingDefinition, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response BuildingList

	rv, err := client.CallPlugin("RemoteFortressReader", "GetBuildingDefList", &request, &response)

	return response.GetBuildingList(), rv, err
}

// CallGetWorldMap is a convenience wrapper around the RemoteFortressReader::GetWorldMap RPC method.
func CallGetWorldMap(client *dfhackrpc.Client) (*WorldMap, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response WorldMap

	rv, err := client.CallPlugin("RemoteFortressReader", "GetWorldMap", &request, &response)

	return &response, rv, err
}

// CallGetWorldMapNew is a convenience wrapper around the RemoteFortressReader::GetWorldMapNew RPC method.
func CallGetWorldMapNew(client *dfhackrpc.Client) (*WorldMap, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response WorldMap

	rv, err := client.CallPlugin("RemoteFortressReader", "GetWorldMapNew", &request, &response)

	return &response, rv, err
}

// CallGetRegionMaps is a convenience wrapper around the RemoteFortressReader::GetRegionMaps RPC method.
func CallGetRegionMaps(client *dfhackrpc.Client) (*RegionMaps, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response RegionMaps

	rv, err := client.CallPlugin("RemoteFortressReader", "GetRegionMaps", &request, &response)

	return &response, rv, err
}

// CallGetRegionMapsNew is a convenience wrapper around the RemoteFortressReader::GetRegionMapsNew RPC method.
func CallGetRegionMapsNew(client *dfhackrpc.Client) (*RegionMaps, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response RegionMaps

	rv, err := client.CallPlugin("RemoteFortressReader", "GetRegionMapsNew", &request, &response)

	return &response, rv, err
}

// CallGetCreatureRaws is a convenience wrapper around the RemoteFortressReader::GetCreatureRaws RPC method.
func CallGetCreatureRaws(client *dfhackrpc.Client) ([]*CreatureRaw, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response CreatureRawList

	rv, err := client.CallPlugin("RemoteFortressReader", "GetCreatureRaws", &request, &response)

	return response.GetCreatureRaws(), rv, err
}

// CallGetPartialCreatureRaws is a convenience wrapper around the RemoteFortressReader::GetPartialCreatureRaws RPC method.
func CallGetPartialCreatureRaws(client *dfhackrpc.Client, request *ListRequest) ([]*CreatureRaw, dfhackrpc.CommandResult, error) {
	var response CreatureRawList

	rv, err := client.CallPlugin("RemoteFortressReader", "GetPartialCreatureRaws", request, &response)

	return response.GetCreatureRaws(), rv, err
}

// CallGetWorldMapCenter is a convenience wrapper around the RemoteFortressReader::GetWorldMapCenter RPC method.
func CallGetWorldMapCenter(client *dfhackrpc.Client) (*WorldMap, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response WorldMap

	rv, err := client.CallPlugin("RemoteFortressReader", "GetWorldMapCenter", &request, &response)

	return &response, rv, err
}

// CallGetPlantRaws is a convenience wrapper around the RemoteFortressReader::GetPlantRaws RPC method.
func CallGetPlantRaws(client *dfhackrpc.Client) ([]*PlantRaw, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response PlantRawList

	rv, err := client.CallPlugin("RemoteFortressReader", "GetPlantRaws", &request, &response)

	return response.GetPlantRaws(), rv, err
}

// CallGetPartialPlantRaws is a convenience wrapper around the RemoteFortressReader::GetPartialPlantRaws RPC method.
func CallGetPartialPlantRaws(client *dfhackrpc.Client, request *ListRequest) ([]*PlantRaw, dfhackrpc.CommandResult, error) {
	var response PlantRawList

	rv, err := client.CallPlugin("RemoteFortressReader", "GetPartialPlantRaws", request, &response)

	return response.GetPlantRaws(), rv, err
}

// CallCopyScreen is a convenience wrapper around the RemoteFortressReader::CopyScreen RPC method.
func CallCopyScreen(client *dfhackrpc.Client) (*ScreenCapture, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response ScreenCapture

	rv, err := client.CallPlugin("RemoteFortressReader", "CopyScreen", &request, &response)

	return &response, rv, err
}

// CallPassKeyboardEvent is a convenience wrapper around the RemoteFortressReader::PassKeyboardEvent RPC method.
func CallPassKeyboardEvent(client *dfhackrpc.Client, request *KeyboardEvent) (dfhackrpc.CommandResult, error) {
	var response dfproto_core.EmptyMessage

	return client.CallPlugin("RemoteFortressReader", "PassKeyboardEvent", request, &response)
}

// CallSendDigCommand is a convenience wrapper around the RemoteFortressReader::SendDigCommand RPC method.
func CallSendDigCommand(client *dfhackrpc.Client, request *DigCommand) (dfhackrpc.CommandResult, error) {
	var response dfproto_core.EmptyMessage

	return client.CallPlugin("RemoteFortressReader", "SendDigCommand", request, &response)
}

// CallSetPauseState is a convenience wrapper around the RemoteFortressReader::SetPauseState RPC method.
func CallSetPauseState(client *dfhackrpc.Client, value bool) (dfhackrpc.CommandResult, error) {
	var request SingleBool
	request.Value = &value

	var response dfproto_core.EmptyMessage

	return client.CallPlugin("RemoteFortressReader", "SetPauseState", &request, &response)
}

// CallGetPauseState is a convenience wrapper around the RemoteFortressReader::GetPauseState RPC method.
func CallGetPauseState(client *dfhackrpc.Client) (bool, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response SingleBool

	rv, err := client.CallPlugin("RemoteFortressReader", "GetPauseState", &request, &response)

	return response.GetValue(), rv, err
}

// CallGetVersionInfo is a convenience wrapper around the RemoteFortressReader::GetVersionInfo RPC method.
func CallGetVersionInfo(client *dfhackrpc.Client) (*VersionInfo, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response VersionInfo

	rv, err := client.CallPlugin("RemoteFortressReader", "GetVersionInfo", &request, &response)

	return &response, rv, err
}

// CallGetReports is a convenience wrapper around the RemoteFortressReader::GetReports RPC method.
func CallGetReports(client *dfhackrpc.Client) ([]*Report, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response Status

	rv, err := client.CallPlugin("RemoteFortressReader", "GetReports", &request, &response)

	return response.GetReports(), rv, err
}

// CallMoveCommand is a convenience wrapper around the RemoteFortressReader::MoveCommand RPC method.
func CallMoveCommand(client *dfhackrpc.Client, request *MoveCommandParams) (dfhackrpc.CommandResult, error) {
	var response dfproto_core.EmptyMessage

	return client.CallPlugin("RemoteFortressReader", "MoveCommand", request, &response)
}

// CallJumpCommand is a convenience wrapper around the RemoteFortressReader::JumpCommand RPC method.
func CallJumpCommand(client *dfhackrpc.Client, request *MoveCommandParams) (dfhackrpc.CommandResult, error) {
	var response dfproto_core.EmptyMessage

	return client.CallPlugin("RemoteFortressReader", "JumpCommand", request, &response)
}

// CallMenuQuery is a convenience wrapper around the RemoteFortressReader::MenuQuery RPC method.
func CallMenuQuery(client *dfhackrpc.Client) (*MenuContents, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response MenuContents

	rv, err := client.CallPlugin("RemoteFortressReader", "MenuQuery", &request, &response)

	return &response, rv, err
}

// CallMovementSelectCommand is a convenience wrapper around the RemoteFortressReader::MovementSelectCommand RPC method.
func CallMovementSelectCommand(client *dfhackrpc.Client, value int32) (dfhackrpc.CommandResult, error) {
	var request dfproto_core.IntMessage
	request.Value = &value

	var response dfproto_core.EmptyMessage

	return client.CallPlugin("RemoteFortressReader", "MovementSelectCommand", &request, &response)
}

// CallMiscMoveCommand is a convenience wrapper around the RemoteFortressReader::MiscMoveCommand RPC method.
func CallMiscMoveCommand(client *dfhackrpc.Client, request *MiscMoveParams) (dfhackrpc.CommandResult, error) {
	var response dfproto_core.EmptyMessage

	return client.CallPlugin("RemoteFortressReader", "MiscMoveCommand", request, &response)
}

// CallGetLanguage is a convenience wrapper around the RemoteFortressReader::GetLanguage RPC method.
func CallGetLanguage(client *dfhackrpc.Client) (*Language, dfhackrpc.CommandResult, error) {
	var request dfproto_core.EmptyMessage

	var response Language

	rv, err := client.CallPlugin("RemoteFortressReader", "GetLanguage", &request, &response)

	return &response, rv, err
}
