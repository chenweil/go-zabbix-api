package zabbix

import "encoding/json"

// ItemPrototype represents Zabbix Item Prototype for LLD
type ItemPrototype struct {
	ItemID       string    `json:"itemid,omitempty"`
	Delay        string    `json:"delay"`
	HostID       string    `json:"hostid"`
	InterfaceID  string    `json:"interfaceid,omitempty"`
	Key          string    `json:"key_"`
	Name         string    `json:"name"`
	Type         ItemType  `json:"type,string"`
	ValueType    ValueType `json:"value_type,string"`
	DataType     DataType  `json:"data_type,string"`
	Delta        DeltaType `json:"delta,string"`
	Description  string    `json:"description"`
	Error        string    `json:"error,omitempty"`
	History      string    `json:"history,omitempty"`
	Trends       string    `json:"trends,omitempty"`
	TrapperHosts string    `json:"trapper_hosts,omitempty"`
	Params       string    `json:"params,omitempty"`
	UUID         string    `json:"uuid,omitempty"`

	// LLD Rule ID - required for item prototypes
	RuleID        string `json:"ruleid,omitempty"`
	DiscoveryRule *LLDRule `json:"discoveryRule,omitempty"`

	// Zabbix 6.0 uses Tags instead of Applications
	Tags Tags `json:"tags,omitempty"`

	ItemParent Hosts `json:"hosts"`

	Preprocessors Preprocessors `json:"preprocessing,omitempty"`

	// HTTP Agent Fields
	Url           string          `json:"url,omitempty"`
	RequestMethod string          `json:"request_method,omitempty"`
	PostType      string          `json:"post_type,omitempty"`
	Posts         string          `json:"posts,omitempty"`
	StatusCodes   string          `json:"status_codes,omitempty"`
	Timeout       string          `json:"timeout,omitempty"`
	VerifyHost    string          `json:"verify_host,omitempty"`
	VerifyPeer    string          `json:"verify_peer,omitempty"`
	HeadersV6     HttpHeaders     `json:"headers_v6,omitempty"`  // Zabbix 6.0 format
	HeadersV7     []HeaderField   `json:"headers_v7,omitempty"`  // Zabbix 7.0+ format

	// SNMP Fields
	SNMPOid              string `json:"snmp_oid,omitempty"`
	SNMPCommunity        string `json:"snmp_community,omitempty"`
	SNMPv3AuthPassphrase string `json:"snmpv3_authpassphrase,omitempty"`
	SNMPv3AuthProtocol   string `json:"snmpv3_authprotocol,omitempty"`
	SNMPv3ContextName    string `json:"snmpv3_contextname,omitempty"`
	SNMPv3PrivPasshrase  string `json:"snmpv3_privpassphrase,omitempty"`
	SNMPv3PrivProtocol   string `json:"snmpv3_privprotocol,omitempty"`
	SNMPv3SecurityLevel  string `json:"snmpv3_securitylevel,omitempty"`
	SNMPv3SecurityName   string `json:"snmpv3_securityname,omitempty"`

	// Dependent Fields
	MasterItemID string `json:"master_itemid,omitempty"`

	// Zabbix 7.0+ new fields
	QueryFieldsV6  map[string]string `json:"query_fields_v6,omitempty"`
	QueryFieldsV7  []HeaderField     `json:"query_fields_v7,omitempty"`
	RawQueryFields json.RawMessage   `json:"query_fields,omitempty"`

	// Browser item specific fields (Zabbix 7.0+)
	BrowserScript string `json:"browser_script,omitempty"`
	BrowserParams string `json:"browser_params,omitempty"`
}

// ItemPrototypes is an array of ItemPrototype
type ItemPrototypes []ItemPrototype

// ItemPrototypesGet Wrapper for itemprototype.get
// https://www.zabbix.com/documentation/current/manual/api/reference/itemprototype/get
func (api *API) ItemPrototypesGet(params Params) (res ItemPrototypes, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("itemprototype.get", params, &res)
	return
}

// ItemPrototypeGetByID Gets item prototype by Id only if there is exactly 1 matching item prototype
func (api *API) ItemPrototypeGetByID(id string) (res *ItemPrototype, err error) {
	items, err := api.ItemPrototypesGet(Params{"itemids": id})
	if err != nil {
		return
	}

	if len(items) == 1 {
		res = &items[0]
	} else {
		e := ExpectedOneResult(len(items))
		err = &e
	}

	return
}

// ItemPrototypesCreate Wrapper for itemprototype.create
// https://www.zabbix.com/documentation/current/manual/api/reference/itemprototype/create
func (api *API) ItemPrototypesCreate(items ItemPrototypes) (err error) {
	response, err := api.CallWithError("itemprototype.create", items)
	if err != nil {
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(response.Result, &result)
	if err != nil {
		return
	}
	itemids := result["itemids"].([]interface{})

	for i := range items {
		id := itemids[i].(string)
		items[i].ItemID = id
	}

	return
}

// ItemPrototypesUpdate Wrapper for itemprototype.update
// https://www.zabbix.com/documentation/current/manual/api/reference/itemprototype/update
func (api *API) ItemPrototypesUpdate(items ItemPrototypes) (err error) {
	_, err = api.CallWithError("itemprototype.update", items)
	return
}

// ItemPrototypesDelete Wrapper for itemprototype.delete
// Cleans ItemId in all items elements if call succeed.
// https://www.zabbix.com/documentation/current/manual/api/reference/itemprototype/delete
func (api *API) ItemPrototypesDelete(items ItemPrototypes) (err error) {
	ids := make([]string, len(items))
	for i, item := range items {
		ids[i] = item.ItemID
	}

	err = api.ItemPrototypesDeleteByIds(ids)
	if err == nil {
		for i := range items {
			items[i].ItemID = ""
		}
	}
	return
}

// ItemPrototypesDeleteByIds Wrapper for itemprototype.delete
// https://www.zabbix.com/documentation/current/manual/api/reference/itemprototype/delete
func (api *API) ItemPrototypesDeleteByIds(ids []string) (err error) {
	deleteIds, err := api.ItemPrototypesDeleteIDs(ids)
	if err != nil {
		return
	}
	l := len(deleteIds)
	if len(ids) != l {
		err = &ExpectedMore{len(ids), l}
	}
	return
}

// ItemPrototypesDeleteIDs Wrapper for itemprototype.delete
// return the id of the deleted item prototype
func (api *API) ItemPrototypesDeleteIDs(ids []string) (itemids []interface{}, err error) {
	response, err := api.CallWithError("itemprototype.delete", ids)
	if err != nil {
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(response.Result, &result)
	if err != nil {
		return
	}
	itemids1, ok := result["itemids"].([]interface{})
	if !ok {
		itemids2 := result["itemids"].(map[string]interface{})
		for _, id := range itemids2 {
			itemids = append(itemids, id)
		}
	} else {
		itemids = itemids1
	}
	return
}
