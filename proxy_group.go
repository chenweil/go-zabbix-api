package zabbix

import (
	"encoding/json"
	"fmt"
)

// Proxy Group structures for Zabbix 7.0+

// ProxyGroup represents a proxy group configuration
type ProxyGroup struct {
	ProxyGroupID string       `json:"proxy_groupid,omitempty"`
	Name         string       `json:"name"`
	Description  string       `json:"description,omitempty"`
	ProxyState   ProxyState   `json:"proxy_state,string"`
	ProxyCount   string       `json:"proxy_count,omitempty"`
	Discovery    ProxyDiscovery `json:"discovery,omitempty"`
}

// ProxyState represents the state of proxy group
type ProxyState int

const (
	// ProxyStateAny Any proxy state
	ProxyStateAny ProxyState = 0
	// ProxyStateOnline Proxy is online
	ProxyStateOnline ProxyState = 1
	// ProxyStateOffline Proxy is offline
	ProxyStateOffline ProxyState = 2
	// ProxyStateUnknown Proxy state is unknown
	ProxyStateUnknown ProxyState = 3
)

// ProxyDiscovery represents proxy discovery configuration
type ProxyDiscovery struct {
	ProxyDiscoveryRuleID string `json:"proxy_discovery_ruleid,omitempty"`
	Name                 string `json:"name,omitempty"`
	ProxyGroupID         string `json:"proxy_groupid,omitempty"`
}

// ProxyGroups is an array of ProxyGroup
type ProxyGroups []ProxyGroup

// ProxyGroupCreate Wrapper for proxygroup.create
// https://www.zabbix.com/documentation/7.0/manual/api/reference/proxygroup/create
func (api *API) ProxyGroupCreate(proxyGroups ProxyGroups) error {
	if !api.versionManager.IsFeatureSupported(FeatureProxyGroup) {
		return fmt.Errorf("Proxy Group not supported in Zabbix version %s", api.versionManager.serverVersion)
	}

	response, err := api.CallWithError("proxygroup.create", proxyGroups)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	err = json.Unmarshal(response.Result, &result)
	if err != nil {
		return err
	}
	proxyGroupIDs := result["proxy_groupids"].([]interface{})
	for i, id := range proxyGroupIDs {
		proxyGroups[i].ProxyGroupID = id.(string)
	}
	return nil
}

// ProxyGroupGet Wrapper for proxygroup.get
// https://www.zabbix.com/documentation/7.0/manual/api/reference/proxygroup/get
func (api *API) ProxyGroupGet(params Params) (ProxyGroups, error) {
	if !api.versionManager.IsFeatureSupported(FeatureProxyGroup) {
		return nil, fmt.Errorf("Proxy Group not supported in Zabbix version %s", api.versionManager.serverVersion)
	}

	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	
	var res ProxyGroups
	err := api.CallWithErrorParse("proxygroup.get", params, &res)
	return res, err
}

// ProxyGroupUpdate Wrapper for proxygroup.update
// https://www.zabbix.com/documentation/7.0/manual/api/reference/proxygroup/update
func (api *API) ProxyGroupUpdate(proxyGroups ProxyGroups) error {
	if !api.versionManager.IsFeatureSupported(FeatureProxyGroup) {
		return fmt.Errorf("Proxy Group not supported in Zabbix version %s", api.versionManager.serverVersion)
	}

	_, err := api.CallWithError("proxygroup.update", proxyGroups)
	return err
}

// ProxyGroupDelete Wrapper for proxygroup.delete
// https://www.zabbix.com/documentation/7.0/manual/api/reference/proxygroup/delete
func (api *API) ProxyGroupDelete(proxyGroups ProxyGroups) error {
	if !api.versionManager.IsFeatureSupported(FeatureProxyGroup) {
		return fmt.Errorf("Proxy Group not supported in Zabbix version %s", api.versionManager.serverVersion)
	}

	ids := make([]string, len(proxyGroups))
	for i, pg := range proxyGroups {
		ids[i] = pg.ProxyGroupID
	}

	_, err := api.CallWithError("proxygroup.delete", ids)
	if err == nil {
		for i := range proxyGroups {
			proxyGroups[i].ProxyGroupID = ""
		}
	}
	return err
}

// ProxyGroupGetByID Gets proxy group by ID
func (api *API) ProxyGroupGetByID(id string) (*ProxyGroup, error) {
	groups, err := api.ProxyGroupGet(Params{"proxy_groupids": id})
	if err != nil {
		return nil, err
	}

	if len(groups) != 1 {
		e := ExpectedOneResult(len(groups))
		return nil, &e
	}
	return &groups[0], nil
}

// ProxyGroupGetByName Gets proxy group by name
func (api *API) ProxyGroupGetByName(name string) (*ProxyGroup, error) {
	groups, err := api.ProxyGroupGet(Params{"filter": map[string]string{"name": name}})
	if err != nil {
		return nil, err
	}

	if len(groups) != 1 {
		e := ExpectedOneResult(len(groups))
		return nil, &e
	}
	return &groups[0], nil
}

// Enhanced Proxy structure for Zabbix 7.0+
type Proxy7 struct {
	ProxyID      string       `json:"proxyid,omitempty"`
	Host         string       `json:"host"`
	Name         string       `json:"name"`
	Status       ProxyStatus  `json:"status,string"`
	Description  string       `json:"description,omitempty"`
	TLSConnect   TLSConnect   `json:"tls_connect,string"`
	TLSAccept    TLSAccept    `json:"tls_accept,string"`
	TLSIssuer    string       `json:"tls_issuer,omitempty"`
	TLSSubject   string       `json:"tls_subject,omitempty"`
	TLSPSKIdentity string     `json:"tls_psk_identity,omitempty"`
	TLSPSK       string       `json:"tls_psk,omitempty"`
	
	// Zabbix 7.0+ fields
	ProxyGroupID string       `json:"proxy_groupid,omitempty"`
	Address      string       `json:"address,omitempty"`
	Port         string       `json:"port,omitempty"`
	
	// Custom timeouts for Zabbix 7.0+
	CustomTimeouts CustomTimeouts `json:"custom_timeouts,omitempty"`
	
	// Browser timeout for Zabbix 7.0+
	TimeoutBrowser string `json:"timeout_browser,omitempty"`
}

// CustomTimeouts represents custom timeout configuration
type CustomTimeouts struct {
	TimeoutAgent     string `json:"timeout_agent,omitempty"`
	TimeoutSimpleCheck string `json:"timeout_simple_check,omitempty"`
	TimeoutSNMPTrapper string `json:"timeout_snmp_trapper,omitempty"`
	TimeoutExternalCheck string `json:"timeout_external_check,omitempty"`
	TimeoutDBMonitor string `json:"timeout_db_monitor,omitempty"`
	TimeoutHTTPAgent string `json:"timeout_http_agent,omitempty"`
	TimeoutSSHAgent string `json:"timeout_ssh_agent,omitempty"`
	TimeoutTELNETAgent string `json:"timeout_telnet_agent,omitempty"`
	TimeoutJMXAgent string `json:"timeout_jmx_agent,omitempty"`
	TimeoutIPMIAgent string `json:"timeout_ipmi_agent,omitempty"`
}

// TLSConnect represents TLS connection mode
type TLSConnect int

const (
	// TLSConnectNoEncryption No encryption
	TLSConnectNoEncryption TLSConnect = 0
	// TLSConnectPSK PSK encryption
	TLSConnectPSK TLSConnect = 1
	// TLSConnectCertificate Certificate encryption
	TLSConnectCertificate TLSConnect = 2
)

// TLSAccept represents TLS accepted modes
type TLSAccept int

const (
	// TLSAcceptNoEncryption Accept connections without encryption
	TLSAcceptNoEncryption TLSAccept = 0
	// TLSAcceptPSK Accept connections with PSK
	TLSAcceptPSK TLSAccept = 1
	// TLSAcceptCertificate Accept connections with certificate
	TLSAcceptCertificate TLSAccept = 2
)

// ProxyStatus represents proxy status
type ProxyStatus int

const (
	// ProxyStatusEnabled Proxy is enabled
	ProxyStatusEnabled ProxyStatus = 5
	// ProxyStatusDisabled Proxy is disabled
	ProxyStatusDisabled ProxyStatus = 6
)

// Proxies7 is an array of Proxy7
type Proxies7 []Proxy7

// Proxy7Create Wrapper for proxy.create with Zabbix 7.0+ features
func (api *API) Proxy7Create(proxies Proxies7) error {
	if !api.versionManager.IsFeatureSupported(FeatureProxyGroup) {
		return fmt.Errorf("Enhanced proxy features not supported in Zabbix version %s", api.versionManager.serverVersion)
	}

	response, err := api.CallWithError("proxy.create", proxies)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	err = json.Unmarshal(response.Result, &result)
	if err != nil {
		return err
	}
	proxyIDs := result["proxyids"].([]interface{})
	for i, id := range proxyIDs {
		proxies[i].ProxyID = id.(string)
	}
	return nil
}

// Proxy7Get Wrapper for proxy.get with Zabbix 7.0+ features
func (api *API) Proxy7Get(params Params) (Proxies7, error) {
	if !api.versionManager.IsFeatureSupported(FeatureProxyGroup) {
		return nil, fmt.Errorf("Enhanced proxy features not supported in Zabbix version %s", api.versionManager.serverVersion)
	}

	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	
	var res Proxies7
	err := api.CallWithErrorParse("proxy.get", params, &res)
	return res, err
}

// Proxy7Update Wrapper for proxy.update with Zabbix 7.0+ features
func (api *API) Proxy7Update(proxies Proxies7) error {
	if !api.versionManager.IsFeatureSupported(FeatureProxyGroup) {
		return fmt.Errorf("Enhanced proxy features not supported in Zabbix version %s", api.versionManager.serverVersion)
	}

	_, err := api.CallWithError("proxy.update", proxies)
	return err
}

// Proxy7GetByID Gets proxy by ID
func (api *API) Proxy7GetByID(id string) (*Proxy7, error) {
	proxies, err := api.Proxy7Get(Params{"proxyids": id})
	if err != nil {
		return nil, err
	}

	if len(proxies) != 1 {
		e := ExpectedOneResult(len(proxies))
		return nil, &e
	}
	return &proxies[0], nil
}