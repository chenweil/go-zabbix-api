package zabbix

import (
	"encoding/json"
	"fmt"
)

type (
	// ItemType type of the item
	ItemType int
	// ValueType type of information of the item
	ValueType int
	// DataType data type of the item
	DataType int
	// DeltaType value that will be stored
	DeltaType int
)

const (
	// Different item type, see :
	// - "type" in https://www.zabbix.com/documentation/3.2/manual/api/reference/item/object
	// - https://www.zabbix.com/documentation/3.2/manual/config/items/itemtypes

	// ZabbixAgent type
	ZabbixAgent ItemType = 0
	// SNMPv1Agent type
	SNMPv1Agent ItemType = 1
	// ZabbixTrapper type
	ZabbixTrapper ItemType = 2
	// SimpleCheck type
	SimpleCheck ItemType = 3
	// SNMPv2Agent type
	SNMPv2Agent ItemType = 4
	// ZabbixInternal type
	ZabbixInternal ItemType = 5
	// SNMPv3Agent type
	SNMPv3Agent ItemType = 6
	// ZabbixAgentActive type
	ZabbixAgentActive ItemType = 7
	// ZabbixAggregate type
	ZabbixAggregate ItemType = 8
	// WebItem type
	WebItem ItemType = 9
	// ExternalCheck type
	ExternalCheck ItemType = 10
	// DatabaseMonitor type
	DatabaseMonitor ItemType = 11
	//IPMIAgent type
	IPMIAgent ItemType = 12
	// SSHAgent type
	SSHAgent ItemType = 13
	// TELNETAgent type
	TELNETAgent ItemType = 14
	// Calculated type
	Calculated ItemType = 15
	// JMXAgent type
	JMXAgent ItemType = 16
	// SNMPTrap type
	SNMPTrap ItemType = 17
	// Dependent type
	Dependent ItemType = 18
	// HTTPAgent type
	HTTPAgent ItemType = 19
	SNMPAgent ItemType = 20
	// Browser type (Zabbix 7.0+)
	Browser ItemType = 22
)

const (
	// HTTP Request Method constants for HTTP Agent items
	// Enhanced for Zabbix 6.0 with additional HTTP methods

	// HTTP GET method
	HTTPMethodGET = "0"
	// HTTP POST method
	HTTPMethodPOST = "1"
	// HTTP PUT method
	HTTPMethodPUT = "2"
	// HTTP HEAD method (Added in Zabbix 6.0)
	HTTPMethodHEAD = "3"
	// HTTP PATCH method (Added in Zabbix 6.0)
	HTTPMethodPATCH = "4"
	// HTTP DELETE method
	HTTPMethodDELETE = "5"
	// HTTP OPTIONS method (Added in Zabbix 6.0)
	HTTPMethodOPTIONS = "6"
	// HTTP TRACE method (Added in Zabbix 6.0)
	HTTPMethodTRACE = "7"
	// HTTP CONNECT method (Added in Zabbix 6.0)
	HTTPMethodCONNECT = "8"
)

const (
	// Type of information of the item
	// see "value_type" in https://www.zabbix.com/documentation/3.2/manual/api/reference/item/object

	// Numeric float (default)
	NumericFloat ValueType = 0
	// Character
	Character ValueType = 1
	// Log
	Log ValueType = 2
	// Numeric unsigned
	NumericUnsigned ValueType = 3
	// Text
	Text ValueType = 4
	// Binary (not supported)
	Binary ValueType = 5
	// Zabbix 6.0: Calculated text (新增)
	CalculatedText ValueType = 5
	// Zabbix 6.0: Calculated log (新增)
	CalculatedLog ValueType = 6
	// Zabbix 6.0: Calculated character (新增)
	CalculatedChar ValueType = 7
)

const (
	// Data type of the item
	// see "data_type" in https://www.zabbix.com/documentation/3.2/manual/api/reference/item/object

	// Decimal data (default)
	Decimal DataType = 0
	// Octal data
	Octal DataType = 1
	// Hexadecimal data
	Hexadecimal DataType = 2
	// Boolean data
	Boolean DataType = 3
)

const (
	// Value that will be stored
	// see "delta" in https://www.zabbix.com/documentation/3.2/manual/api/reference/item/object

	// AsIs as is (default)
	AsIs DeltaType = 0
	// Speed speed per second
	Speed DeltaType = 1
	// Delta simple change
	Delta DeltaType = 2
)

// Item represent Zabbix item object
// https://www.zabbix.com/documentation/3.2/manual/api/reference/item/object
type Item struct {
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

	// Zabbix 6.0 uses Tags instead of Applications
	Tags Tags `json:"tags,omitempty"`

	ItemParent Hosts `json:"hosts"`

	Preprocessors Preprocessors `json:"preprocessing,omitempty"`

	// HTTP Agent Fields
	Url           string `json:"url,omitempty"`
	RequestMethod string `json:"request_method,omitempty"`
	PostType      string `json:"post_type,omitempty"`
	Posts         string `json:"posts,omitempty"`
	StatusCodes   string `json:"status_codes,omitempty"`
	Timeout       string `json:"timeout,omitempty"`
	VerifyHost    string `json:"verify_host,omitempty"`
	VerifyPeer    string `json:"verify_peer,omitempty"`
	HeadersV6     HttpHeaders   `json:"headers_v6,omitempty"`  // Zabbix 6.0 format
	HeadersV7     []HeaderField `json:"headers_v7,omitempty"`  // Zabbix 7.0+ format

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

	// Prototype
	RuleID        string   `json:"ruleid,omitempty"`
	DiscoveryRule *LLDRule `json:"discoveryRule,omitEmpty"`

	// Zabbix 7.0+ new fields
	QueryFieldsV6  map[string]string `json:"query_fields_v6,omitempty"`
	QueryFieldsV7  []HeaderField     `json:"query_fields_v7,omitempty"`
	RawQueryFields json.RawMessage   `json:"query_fields,omitempty"`

	// Browser item specific fields (Zabbix 7.0+)
	BrowserScript string `json:"browser_script,omitempty"`
	BrowserParams string `json:"browser_params,omitempty"`
}

// HeaderField represents HTTP header field for Zabbix 7.0+ format
type HeaderField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// HttpHeaders represents HTTP headers for Zabbix 6.0 format
type HttpHeaders map[string]string

// Items is an array of Item
type Items []Item

// BrowserItem represents a Browser monitoring item (Zabbix 7.0+)
type BrowserItem struct {
	Item
}

type Preprocessors []Preprocessor

type Preprocessor struct {
	Type               string `json:"type,omitempty"`
	Params             string `json:"params,omitempty"`
	ErrorHandler       string `json:"error_handler,omitempty"`
	ErrorHandlerParams string `json:"error_handler_params,omitempty"`
}

// ItemsGet Wrapper for item.get
// https://www.zabbix.com/documentation/3.2/manual/api/reference/item/get
func (api *API) ItemsGet(params Params) (res Items, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	err = api.CallWithErrorParse("item.get", params, &res)
	return
}

// ItemsGetByHostIds Gets items by host Ids.
func (api *API) ItemsGetByHostIds(ids []string) (res Items, err error) {
	return api.ItemsGet(Params{"hostids": ids})
}

// ItemsGetByHosts Gets items by hosts.
func (api *API) ItemsGetByHosts(hosts Hosts) (res Items, err error) {
	ids := make([]string, len(hosts))
	for i, host := range hosts {
		ids[i] = host.HostID
	}
	return api.ItemsGetByHostIds(ids)
}

// ItemGetByID Gets item by Id only if there is exactly 1 matching item.
func (api *API) ItemGetByID(id string) (res *Item, err error) {
	items, err := api.ItemsGet(Params{"itemids": id})
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

// ItemsCreate Wrapper for item.create
// https://www.zabbix.com/documentation/3.2/manual/api/reference/item/create
//
// Zabbix 6.0 Enhancement: For HTTP Agent items, interfaceid is no longer required
func (api *API) ItemsCreate(items Items) (err error) {
	// Use adapter pattern for multi-version support
	if api.itemAdapter != nil {
		return api.itemAdapter.CreateItems(items)
	}

	// Fallback to original implementation
	response, err := api.CallWithError("item.create", items)
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

// ItemsUpdate Wrapper for item.update
// https://www.zabbix.com/documentation/3.2/manual/api/reference/item/update
func (api *API) ItemsUpdate(items Items) (err error) {
	// Use adapter pattern for multi-version support
	if api.itemAdapter != nil {
		return api.itemAdapter.UpdateItems(items)
	}

	// Fallback to original implementation
	_, err = api.CallWithError("item.update", items)
	return
}

// ItemsDelete Wrapper for item.delete
// https://www.zabbix.com/documentation/3.2/manual/api/reference/item/delete
func (api *API) ItemsDelete(items Items) (err error) {
	ids := make([]string, len(items))
	for i, item := range items {
		ids[i] = item.ItemID
	}
	_, err = api.CallWithError("item.delete", ids)
	return
}

// ItemsDeleteByIds Wrapper for item.delete
// https://www.zabbix.com/documentation/3.2/manual/api/reference/item/delete
func (api *API) ItemsDeleteByIds(ids []string) (err error) {
	// Use adapter pattern for multi-version support
	if api.itemAdapter != nil {
		return api.itemAdapter.DeleteItems(ids)
	}

	// Fallback to original implementation
	_, err = api.CallWithError("item.delete", ids)
	return
}

// Zabbix6ItemAdapter implements ItemAdapter for Zabbix 6.0
type Zabbix6ItemAdapter struct {
	api *API
}

func (adapter *Zabbix6ItemAdapter) CreateItems(items Items) error {
	// Prepare items for Zabbix 6.0 format
	for i := range items {
		item := &items[i]

		// Convert HeadersV7 to HeadersV6 if needed
		if len(item.HeadersV7) > 0 {
			item.HeadersV6 = make(HttpHeaders)
			for _, header := range item.HeadersV7 {
				item.HeadersV6[header.Name] = header.Value
			}
			item.HeadersV7 = nil
		}

		// Convert QueryFieldsV7 to QueryFieldsV6 if needed
		if len(item.QueryFieldsV7) > 0 {
			item.QueryFieldsV6 = make(map[string]string)
			for _, field := range item.QueryFieldsV7 {
				item.QueryFieldsV6[field.Name] = field.Value
			}
			item.QueryFieldsV7 = nil
		}

		// Clear Zabbix 7.0 specific fields
		item.BrowserScript = ""
		item.BrowserParams = ""
	}

	return adapter.api.createItemsLegacy(items)
}

func (adapter *Zabbix6ItemAdapter) GetItems(params Params) (Items, error) {
	return adapter.api.ItemsGet(params)
}

func (adapter *Zabbix6ItemAdapter) UpdateItems(items Items) error {
	// Prepare items for Zabbix 6.0 format
	for i := range items {
		item := &items[i]

		// Convert HeadersV7 to HeadersV6 if needed
		if len(item.HeadersV7) > 0 {
			item.HeadersV6 = make(HttpHeaders)
			for _, header := range item.HeadersV7 {
				item.HeadersV6[header.Name] = header.Value
			}
			item.HeadersV7 = nil
		}

		// Convert QueryFieldsV7 to QueryFieldsV6 if needed
		if len(item.QueryFieldsV7) > 0 {
			item.QueryFieldsV6 = make(map[string]string)
			for _, field := range item.QueryFieldsV7 {
				item.QueryFieldsV6[field.Name] = field.Value
			}
			item.QueryFieldsV7 = nil
		}

		// Clear Zabbix 7.0 specific fields
		item.BrowserScript = ""
		item.BrowserParams = ""
	}

	return adapter.api.updateItemsLegacy(items)
}

func (adapter *Zabbix6ItemAdapter) DeleteItems(itemIds []string) error {
	return adapter.api.ItemsDeleteByIds(itemIds)
}

// Zabbix7ItemAdapter implements ItemAdapter for Zabbix 7.0+
type Zabbix7ItemAdapter struct {
	api *API
}

func (adapter *Zabbix7ItemAdapter) CreateItems(items Items) error {
	// Prepare items for Zabbix 7.0 format
	for i := range items {
		item := &items[i]

		// Convert HeadersV6 to HeadersV7 if needed
		if len(item.HeadersV6) > 0 {
			item.HeadersV7 = make([]HeaderField, 0, len(item.HeadersV6))
			for name, value := range item.HeadersV6 {
				item.HeadersV7 = append(item.HeadersV7, HeaderField{
					Name:  name,
					Value: value,
				})
			}
			item.HeadersV6 = nil
		}

		// Convert QueryFieldsV6 to QueryFieldsV7 if needed
		if len(item.QueryFieldsV6) > 0 {
			item.QueryFieldsV7 = make([]HeaderField, 0, len(item.QueryFieldsV6))
			for name, value := range item.QueryFieldsV6 {
				item.QueryFieldsV7 = append(item.QueryFieldsV7, HeaderField{
					Name:  name,
					Value: value,
				})
			}
			item.QueryFieldsV6 = nil
		}
	}

	return adapter.api.createItemsLegacy(items)
}

func (adapter *Zabbix7ItemAdapter) GetItems(params Params) (Items, error) {
	return adapter.api.ItemsGet(params)
}

func (adapter *Zabbix7ItemAdapter) UpdateItems(items Items) error {
	// Prepare items for Zabbix 7.0 format
	for i := range items {
		item := &items[i]

		// Convert HeadersV6 to HeadersV7 if needed
		if len(item.HeadersV6) > 0 {
			item.HeadersV7 = make([]HeaderField, 0, len(item.HeadersV6))
			for name, value := range item.HeadersV6 {
				item.HeadersV7 = append(item.HeadersV7, HeaderField{
					Name:  name,
					Value: value,
				})
			}
			item.HeadersV6 = nil
		}

		// Convert QueryFieldsV6 to QueryFieldsV7 if needed
		if len(item.QueryFieldsV6) > 0 {
			item.QueryFieldsV7 = make([]HeaderField, 0, len(item.QueryFieldsV6))
			for name, value := range item.QueryFieldsV6 {
				item.QueryFieldsV7 = append(item.QueryFieldsV7, HeaderField{
					Name:  name,
					Value: value,
				})
			}
			item.QueryFieldsV6 = nil
		}
	}

	return adapter.api.updateItemsLegacy(items)
}

func (adapter *Zabbix7ItemAdapter) DeleteItems(itemIds []string) error {
	return adapter.api.ItemsDeleteByIds(itemIds)
}

// Legacy methods for fallback
func (api *API) createItemsLegacy(items Items) error {
	response, err := api.CallWithError("item.create", items)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	err = json.Unmarshal(response.Result, &result)
	if err != nil {
		return err
	}
	itemids := result["itemids"].([]interface{})

	for i := range items {
		id := itemids[i].(string)
		items[i].ItemID = id
	}

	return nil
}

func (api *API) updateItemsLegacy(items Items) error {
	_, err := api.CallWithError("item.update", items)
	return err
}

// Proxy interface methods for backward compatibility
func (api *API) CreateItems(items Items) error {
	return api.ItemsCreate(items)
}

func (api *API) GetItems(params Params) (Items, error) {
	return api.ItemsGet(params)
}

func (api *API) UpdateItems(items Items) error {
	return api.ItemsUpdate(items)
}

// ValidateBrowserItem validates a Browser Item configuration
func ValidateBrowserItem(item BrowserItem) error {
	if item.Type != Browser {
		return fmt.Errorf("item type must be Browser (22)")
	}
	if item.BrowserScript == "" {
		return fmt.Errorf("browser_script is required for Browser items")
	}
	return nil
}

// ValidateItemForVersion validates an item against a specific Zabbix version
func ValidateItemForVersion(item Item, version string) error {
	// Parse version
	major := 0
	fmt.Sscanf(version, "%d", &major)
	
	// Browser items only supported in Zabbix 7.0+
	if item.Type == Browser && major < 7 {
		return fmt.Errorf("Browser items are only supported in Zabbix 7.0+, current version: %s", version)
	}
	
	return nil
}

// ConvertHeadersToV7 converts Zabbix 6.0 format headers to 7.0 format
func ConvertHeadersToV7(headersV6 HttpHeaders) []HeaderField {
	var headersV7 []HeaderField
	for name, value := range headersV6 {
		headersV7 = append(headersV7, HeaderField{
			Name:  name,
			Value: value,
		})
	}
	return headersV7
}

// ConvertHeadersToV6 converts Zabbix 7.0 format headers to 6.0 format
func ConvertHeadersToV6(headersV7 []HeaderField) HttpHeaders {
	headersV6 := make(HttpHeaders)
	for _, header := range headersV7 {
		headersV6[header.Name] = header.Value
	}
	return headersV6
}
