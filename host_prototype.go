package zabbix

import "encoding/json"

// HostPrototype represents Zabbix Host Prototype for LLD
type HostPrototype struct {
	HostID     string        `json:"hostid,omitempty"`
	Host       string        `json:"host"`
	Name       string        `json:"name"`
	Status     StatusType    `json:"status,string"`
	UUID       string        `json:"uuid,omitempty"`

	// LLD Rule ID - required for host prototypes
	RuleID        string `json:"ruleid,omitempty"`
	DiscoveryRule *LLDRule `json:"discoveryRule,omitempty"`

	// Template linkage
	GroupIds         HostGroupIDs   `json:"groups,omitempty"`
	InterfaceID      string         `json:"interfaceid,omitempty"`
	Templates        TemplateIDs    `json:"templates,omitempty"`

	// Custom interfaces for the discovered host
	Interfaces HostInterfaces `json:"interfaces,omitempty"`

	// Macros for the discovered host
	Macros Macros `json:"macros,omitempty"`

	// Tags for Zabbix 6.0+
	Tags Tags `json:"tags,omitempty"`

	// Inventory mode
	InventoryMode InventoryMode `json:"inventory_mode,string"`
}

// HostPrototypes is an array of HostPrototype
type HostPrototypes []HostPrototype

// HostPrototypesGet Wrapper for hostprototype.get
// https://www.zabbix.com/documentation/current/manual/api/reference/hostprototype/get
func (api *API) HostPrototypesGet(params Params) (res HostPrototypes, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("hostprototype.get", params, &res)
	return
}

// HostPrototypeGetByID Gets host prototype by Id only if there is exactly 1 matching host prototype
func (api *API) HostPrototypeGetByID(id string) (res *HostPrototype, err error) {
	hosts, err := api.HostPrototypesGet(Params{"hostids": id})
	if err != nil {
		return
	}

	if len(hosts) == 1 {
		res = &hosts[0]
	} else {
		e := ExpectedOneResult(len(hosts))
		err = &e
	}

	return
}

// HostPrototypesCreate Wrapper for hostprototype.create
// https://www.zabbix.com/documentation/current/manual/api/reference/hostprototype/create
func (api *API) HostPrototypesCreate(hosts HostPrototypes) (err error) {
	response, err := api.CallWithError("hostprototype.create", hosts)
	if err != nil {
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(response.Result, &result)
	if err != nil {
		return
	}
	hostids := result["hostids"].([]interface{})

	for i := range hosts {
		id := hostids[i].(string)
		hosts[i].HostID = id
	}

	return
}

// HostPrototypesUpdate Wrapper for hostprototype.update
// https://www.zabbix.com/documentation/current/manual/api/reference/hostprototype/update
func (api *API) HostPrototypesUpdate(hosts HostPrototypes) (err error) {
	_, err = api.CallWithError("hostprototype.update", hosts)
	return
}

// HostPrototypesDelete Wrapper for hostprototype.delete
// Cleans HostId in all hosts elements if call succeed.
// https://www.zabbix.com/documentation/current/manual/api/reference/hostprototype/delete
func (api *API) HostPrototypesDelete(hosts HostPrototypes) (err error) {
	ids := make([]string, len(hosts))
	for i, host := range hosts {
		ids[i] = host.HostID
	}

	err = api.HostPrototypesDeleteByIds(ids)
	if err == nil {
		for i := range hosts {
			hosts[i].HostID = ""
		}
	}
	return
}

// HostPrototypesDeleteByIds Wrapper for hostprototype.delete
// https://www.zabbix.com/documentation/current/manual/api/reference/hostprototype/delete
func (api *API) HostPrototypesDeleteByIds(ids []string) (err error) {
	response, err := api.CallWithError("hostprototype.delete", ids)
	if err != nil {
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(response.Result, &result)
	if err != nil {
		return
	}
	hostids := result["hostids"].([]interface{})
	if len(ids) != len(hostids) {
		err = &ExpectedMore{len(ids), len(hostids)}
	}
	return
}
