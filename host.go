package zabbix

import (
	"encoding/json"
)

type (
	// AvailableType (readonly) Availability of Zabbix agent
	// see "available" in: https://www.zabbix.com/documentation/3.2/manual/api/reference/host/object
	AvailableType int

	// StatusType Status and function of the host.
	// see "status" in:	https://www.zabbix.com/documentation/3.2/manual/api/reference/host/object
	StatusType int

	// InventoryMode Inventory mode (from 2.2)
	InventoryMode int

	// MonitoredBy type for Zabbix 7.0+
	MonitoredBy int
)

const (
	// Unknown (default)
	Unknown AvailableType = 0
	// Available host is available
	Available AvailableType = 1
	// Unavailable host is unavailable
	Unavailable AvailableType = 2
)

const (
	// Monitored monitored host(default)
	Monitored StatusType = 0
	// Unmonitored unmonitored host
	Unmonitored StatusType = 1
)

const (
	// Disabled inventory mode (default)
	Disabled InventoryMode = -1
	// Manual inventory mode
	Manual InventoryMode = 0
	// Automatic inventory mode
	Automatic InventoryMode = 1
)

const (
	// Monitored by Zabbix server (default)
	MonitoredByServer MonitoredBy = 0
	// Monitored by Zabbix proxy
	MonitoredByProxy MonitoredBy = 1
	// Monitored by proxy group (Zabbix 7.0+)
	MonitoredByProxyGroup MonitoredBy = 2
)

// Host represent Zabbix host object
// https://www.zabbix.com/documentation/3.2/manual/api/reference/host/object
type Host struct {
	HostID     string        `json:"hostid,omitempty"`
	Host       string        `json:"host"`
	Available  AvailableType `json:"available,string"`
	Error      string        `json:"error"`
	Name       string        `json:"name"`
	Status     StatusType    `json:"status,string"`
	UUID       string        `json:"uuid,omitempty"`
	UserMacros Macros        `json:"macros,omitempty"`

	RawInventory  json.RawMessage `json:"inventory,omitempty"`
	Inventory     Inventory       `json:"-"`
	InventoryMode InventoryMode   `json:"inventory_mode,string"`

	// Zabbix 6.0 Tags support
	Tags Tags `json:"tags,omitempty"`

	// Fields below used only when creating hosts
	GroupIds         HostGroupIDs   `json:"groups,omitempty"`
	Interfaces       HostInterfaces `json:"interfaces,omitempty"`
	TemplateIDs      TemplateIDs    `json:"templates,omitempty"`
	TemplateIDsClear TemplateIDs    `json:"templates_clear,omitempty"`
	// templates are read back from this one
	ParentTemplateIDs TemplateIDs `json:"parentTemplates,omitempty"`
	
	// Multi-version proxy support
	ProxyHostID string      `json:"proxy_hostid,omitempty"`  // Zabbix 6.0 format
	ProxyID     string      `json:"proxyid,omitempty"`      // Zabbix 7.0 format
	MonitoredBy MonitoredBy `json:"monitored_by,omitempty"` // Zabbix 7.0+ required field
}

// Hosts is an array of Host
type Hosts []Host

// Tag structure for Zabbix 6.0 compatibility (reused from trigger.go)
type Tag struct {
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

type Tags []Tag

// HostsGet Wrapper for host.get
// https://www.zabbix.com/documentation/3.2/manual/api/reference/host/get
func (api *API) HostsGet(params Params) (res Hosts, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("host.get", params, &res)
	return
}

// HostsGetByHostGroupIds Gets hosts by host group Ids.
func (api *API) HostsGetByHostGroupIds(ids []string) (res Hosts, err error) {
	return api.HostsGet(Params{"groupids": ids})
}

// HostsGetByHostGroups Gets hosts by host groups.
func (api *API) HostsGetByHostGroups(hostGroups HostGroups) (res Hosts, err error) {
	ids := make([]string, len(hostGroups))
	for i, id := range hostGroups {
		ids[i] = id.GroupID
	}
	return api.HostsGetByHostGroupIds(ids)
}

// HostGetByID Gets host by Id only if there is exactly 1 matching host.
func (api *API) HostGetByID(id string) (res *Host, err error) {
	hosts, err := api.HostsGet(Params{"hostids": id})
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

// HostsCreate Wrapper for host.create
// https://www.zabbix.com/documentation/3.2/manual/api/reference/host/create
func (api *API) HostsCreate(hosts Hosts) (err error) {
	response, err := api.CallWithError("host.create", hosts)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	hostids := result["hostids"].([]interface{})

	for i := range hosts {
		id := hostids[i].(string)
		hosts[i].HostID = id
	}

	return
}

// HostsUpdate Wrapper for host.update
// https://www.zabbix.com/documentation/3.2/manual/api/reference/host/update
func (api *API) HostsUpdate(hosts Hosts) (err error) {
	_, err = api.CallWithError("host.update", hosts)
	return
}

// HostsDelete Wrapper for host.delete
// https://www.zabbix.com/documentation/3.2/manual/api/reference/host/delete
func (api *API) HostsDelete(hosts Hosts) (err error) {
	ids := make([]string, len(hosts))
	for i, host := range hosts {
		ids[i] = host.HostID
	}
	_, err = api.CallWithError("host.delete", ids)
	return
}

// HostsDeleteByIds Wrapper for host.delete
// https://www.zabbix.com/documentation/3.2/manual/api/reference/host/delete
func (api *API) HostsDeleteByIds(ids []string) (err error) {
	_, err = api.CallWithError("host.delete", ids)
	return
}

// Zabbix6HostAdapter implements HostAdapter for Zabbix 6.0
type Zabbix6HostAdapter struct {
	api *API
}

func (adapter *Zabbix6HostAdapter) CreateHosts(hosts Hosts) error {
	// Prepare hosts for Zabbix 6.0 format
	for i := range hosts {
		host := &hosts[i]
		
		// Convert Zabbix 7.0 proxy fields to Zabbix 6.0 format if needed
		if host.ProxyID != "" && host.ProxyHostID == "" {
			host.ProxyHostID = host.ProxyID
		}
		
		// Clear Zabbix 7.0 specific fields
		host.MonitoredBy = 0 // Default to server monitoring
	}
	
	return adapter.api.HostsCreate(hosts)
}

func (adapter *Zabbix6HostAdapter) GetHosts(params Params) (Hosts, error) {
	return adapter.api.HostsGet(params)
}

func (adapter *Zabbix6HostAdapter) UpdateHosts(hosts Hosts) error {
	// Prepare hosts for Zabbix 6.0 format
	for i := range hosts {
		host := &hosts[i]
		
		// Convert Zabbix 7.0 proxy fields to Zabbix 6.0 format if needed
		if host.ProxyID != "" && host.ProxyHostID == "" {
			host.ProxyHostID = host.ProxyID
		}
		
		// Clear Zabbix 7.0 specific fields
		host.MonitoredBy = 0 // Default to server monitoring
	}
	
	return adapter.api.HostsUpdate(hosts)
}

func (adapter *Zabbix6HostAdapter) DeleteHosts(hostIds []string) error {
	return adapter.api.HostsDeleteByIds(hostIds)
}

// Zabbix7HostAdapter implements HostAdapter for Zabbix 7.0+
type Zabbix7HostAdapter struct {
	api *API
}

func (adapter *Zabbix7HostAdapter) CreateHosts(hosts Hosts) error {
	// Prepare hosts for Zabbix 7.0 format
	for i := range hosts {
		host := &hosts[i]
		
		// Convert Zabbix 6.0 proxy fields to Zabbix 7.0 format if needed
		if host.ProxyHostID != "" && host.ProxyID == "" {
			host.ProxyID = host.ProxyHostID
		}
		
		// Set monitored_by if proxyid is specified
		if host.ProxyID != "" && host.MonitoredBy == 0 {
			host.MonitoredBy = MonitoredByProxy
		}
	}
	
	return adapter.api.HostsCreate(hosts)
}

func (adapter *Zabbix7HostAdapter) GetHosts(params Params) (Hosts, error) {
	return adapter.api.HostsGet(params)
}

func (adapter *Zabbix7HostAdapter) UpdateHosts(hosts Hosts) error {
	// Prepare hosts for Zabbix 7.0 format
	for i := range hosts {
		host := &hosts[i]
		
		// Convert Zabbix 6.0 proxy fields to Zabbix 7.0 format if needed
		if host.ProxyHostID != "" && host.ProxyID == "" {
			host.ProxyID = host.ProxyHostID
		}
		
		// Set monitored_by if proxyid is specified
		if host.ProxyID != "" && host.MonitoredBy == 0 {
			host.MonitoredBy = MonitoredByProxy
		}
	}
	
	return adapter.api.HostsUpdate(hosts)
}

func (adapter *Zabbix7HostAdapter) DeleteHosts(hostIds []string) error {
	return adapter.api.HostsDeleteByIds(hostIds)
}