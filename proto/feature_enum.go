// Code generated by "enumer -type Feature -trimprefix Feature -output feature_enum.go"; DO NOT EDIT.

package proto

import (
	"fmt"
	"strings"
)

const _FeatureName = "TempTablesBlockInfoTimezoneQuotaKeyInClientInfoDisplayNameVersionPatchServerLogsColumnDefaultsMetadataClientWriteInfoSettingsSerializedAsStringsInterServerSecretOpenTelemetryXForwardedForInClientInfoRefererInClientInfoDistributedDepthQueryStartTimeProfileEventsParallelReplicasCustomSerializationQuotaKeyParametersServerQueryTimeInProgressJSONStrings"
const _FeatureLowerName = "temptablesblockinfotimezonequotakeyinclientinfodisplaynameversionpatchserverlogscolumndefaultsmetadataclientwriteinfosettingsserializedasstringsinterserversecretopentelemetryxforwardedforinclientinforefererinclientinfodistributeddepthquerystarttimeprofileeventsparallelreplicascustomserializationquotakeyparametersserverquerytimeinprogressjsonstrings"

var _FeatureMap = map[Feature]string{
	50264: _FeatureName[0:10],
	51903: _FeatureName[10:19],
	54058: _FeatureName[19:27],
	54060: _FeatureName[27:47],
	54372: _FeatureName[47:58],
	54401: _FeatureName[58:70],
	54406: _FeatureName[70:80],
	54410: _FeatureName[80:102],
	54420: _FeatureName[102:117],
	54429: _FeatureName[117:144],
	54441: _FeatureName[144:161],
	54442: _FeatureName[161:174],
	54443: _FeatureName[174:199],
	54447: _FeatureName[199:218],
	54448: _FeatureName[218:234],
	54449: _FeatureName[234:248],
	54451: _FeatureName[248:261],
	54453: _FeatureName[261:277],
	54454: _FeatureName[277:296],
	54458: _FeatureName[296:304],
	54459: _FeatureName[304:314],
	54460: _FeatureName[314:339],
	54473: _FeatureName[339:350],
}

func (i Feature) String() string {
	if str, ok := _FeatureMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Feature(%d)", i)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _FeatureNoOp() {
	var x [1]struct{}
	_ = x[FeatureTempTables-(50264)]
	_ = x[FeatureBlockInfo-(51903)]
	_ = x[FeatureTimezone-(54058)]
	_ = x[FeatureQuotaKeyInClientInfo-(54060)]
	_ = x[FeatureDisplayName-(54372)]
	_ = x[FeatureVersionPatch-(54401)]
	_ = x[FeatureServerLogs-(54406)]
	_ = x[FeatureColumnDefaultsMetadata-(54410)]
	_ = x[FeatureClientWriteInfo-(54420)]
	_ = x[FeatureSettingsSerializedAsStrings-(54429)]
	_ = x[FeatureInterServerSecret-(54441)]
	_ = x[FeatureOpenTelemetry-(54442)]
	_ = x[FeatureXForwardedForInClientInfo-(54443)]
	_ = x[FeatureRefererInClientInfo-(54447)]
	_ = x[FeatureDistributedDepth-(54448)]
	_ = x[FeatureQueryStartTime-(54449)]
	_ = x[FeatureProfileEvents-(54451)]
	_ = x[FeatureParallelReplicas-(54453)]
	_ = x[FeatureCustomSerialization-(54454)]
	_ = x[FeatureQuotaKey-(54458)]
	_ = x[FeatureParameters-(54459)]
	_ = x[FeatureServerQueryTimeInProgress-(54460)]
	_ = x[FeatureJSONStrings-(54473)]
}

var _FeatureValues = []Feature{FeatureTempTables, FeatureBlockInfo, FeatureTimezone, FeatureQuotaKeyInClientInfo, FeatureDisplayName, FeatureVersionPatch, FeatureServerLogs, FeatureColumnDefaultsMetadata, FeatureClientWriteInfo, FeatureSettingsSerializedAsStrings, FeatureInterServerSecret, FeatureOpenTelemetry, FeatureXForwardedForInClientInfo, FeatureRefererInClientInfo, FeatureDistributedDepth, FeatureQueryStartTime, FeatureProfileEvents, FeatureParallelReplicas, FeatureCustomSerialization, FeatureQuotaKey, FeatureParameters, FeatureServerQueryTimeInProgress, FeatureJSONStrings}

var _FeatureNameToValueMap = map[string]Feature{
	_FeatureName[0:10]:         FeatureTempTables,
	_FeatureLowerName[0:10]:    FeatureTempTables,
	_FeatureName[10:19]:        FeatureBlockInfo,
	_FeatureLowerName[10:19]:   FeatureBlockInfo,
	_FeatureName[19:27]:        FeatureTimezone,
	_FeatureLowerName[19:27]:   FeatureTimezone,
	_FeatureName[27:47]:        FeatureQuotaKeyInClientInfo,
	_FeatureLowerName[27:47]:   FeatureQuotaKeyInClientInfo,
	_FeatureName[47:58]:        FeatureDisplayName,
	_FeatureLowerName[47:58]:   FeatureDisplayName,
	_FeatureName[58:70]:        FeatureVersionPatch,
	_FeatureLowerName[58:70]:   FeatureVersionPatch,
	_FeatureName[70:80]:        FeatureServerLogs,
	_FeatureLowerName[70:80]:   FeatureServerLogs,
	_FeatureName[80:102]:       FeatureColumnDefaultsMetadata,
	_FeatureLowerName[80:102]:  FeatureColumnDefaultsMetadata,
	_FeatureName[102:117]:      FeatureClientWriteInfo,
	_FeatureLowerName[102:117]: FeatureClientWriteInfo,
	_FeatureName[117:144]:      FeatureSettingsSerializedAsStrings,
	_FeatureLowerName[117:144]: FeatureSettingsSerializedAsStrings,
	_FeatureName[144:161]:      FeatureInterServerSecret,
	_FeatureLowerName[144:161]: FeatureInterServerSecret,
	_FeatureName[161:174]:      FeatureOpenTelemetry,
	_FeatureLowerName[161:174]: FeatureOpenTelemetry,
	_FeatureName[174:199]:      FeatureXForwardedForInClientInfo,
	_FeatureLowerName[174:199]: FeatureXForwardedForInClientInfo,
	_FeatureName[199:218]:      FeatureRefererInClientInfo,
	_FeatureLowerName[199:218]: FeatureRefererInClientInfo,
	_FeatureName[218:234]:      FeatureDistributedDepth,
	_FeatureLowerName[218:234]: FeatureDistributedDepth,
	_FeatureName[234:248]:      FeatureQueryStartTime,
	_FeatureLowerName[234:248]: FeatureQueryStartTime,
	_FeatureName[248:261]:      FeatureProfileEvents,
	_FeatureLowerName[248:261]: FeatureProfileEvents,
	_FeatureName[261:277]:      FeatureParallelReplicas,
	_FeatureLowerName[261:277]: FeatureParallelReplicas,
	_FeatureName[277:296]:      FeatureCustomSerialization,
	_FeatureLowerName[277:296]: FeatureCustomSerialization,
	_FeatureName[296:304]:      FeatureQuotaKey,
	_FeatureLowerName[296:304]: FeatureQuotaKey,
	_FeatureName[304:314]:      FeatureParameters,
	_FeatureLowerName[304:314]: FeatureParameters,
	_FeatureName[314:339]:      FeatureServerQueryTimeInProgress,
	_FeatureLowerName[314:339]: FeatureServerQueryTimeInProgress,
	_FeatureName[339:350]:      FeatureJSONStrings,
	_FeatureLowerName[339:350]: FeatureJSONStrings,
}

var _FeatureNames = []string{
	_FeatureName[0:10],
	_FeatureName[10:19],
	_FeatureName[19:27],
	_FeatureName[27:47],
	_FeatureName[47:58],
	_FeatureName[58:70],
	_FeatureName[70:80],
	_FeatureName[80:102],
	_FeatureName[102:117],
	_FeatureName[117:144],
	_FeatureName[144:161],
	_FeatureName[161:174],
	_FeatureName[174:199],
	_FeatureName[199:218],
	_FeatureName[218:234],
	_FeatureName[234:248],
	_FeatureName[248:261],
	_FeatureName[261:277],
	_FeatureName[277:296],
	_FeatureName[296:304],
	_FeatureName[304:314],
	_FeatureName[314:339],
	_FeatureName[339:350],
}

// FeatureString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FeatureString(s string) (Feature, error) {
	if val, ok := _FeatureNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _FeatureNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Feature values", s)
}

// FeatureValues returns all values of the enum
func FeatureValues() []Feature {
	return _FeatureValues
}

// FeatureStrings returns a slice of all String values of the enum
func FeatureStrings() []string {
	strs := make([]string, len(_FeatureNames))
	copy(strs, _FeatureNames)
	return strs
}

// IsAFeature returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Feature) IsAFeature() bool {
	_, ok := _FeatureMap[i]
	return ok
}
