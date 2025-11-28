package zabbix

// MediaType represent Zabbix Media Type object
// https://www.zabbix.com/documentation/current/manual/api/reference/mediatype/object
type MediaType struct {
	MediaTypeID   string            `json:"mediatypeid,omitempty"`
	Name          string            `json:"name,omitempty"`
	Type          string            `json:"type,omitempty"`
	Status        string            `json:"status,omitempty"`
	MaxSessions   string            `json:"maxsessions,omitempty"`
	MaxAttempts   string            `json:"maxattempts,omitempty"`
	AttemptInterval string          `json:"attempt_interval,omitempty"`
	Description   string            `json:"description,omitempty"`
	MessageTemplate string          `json:"message_template,omitempty"`
	Parameters    []MediaTypeParam  `json:"parameters,omitempty"`
	Provider      string            `json:"provider,omitempty"`
	Timeout       string            `json:"timeout,omitempty"`
	ProcessTags   string            `json:"process_tags,omitempty"`
	ShowEventMenu string            `json:"show_event_menu,omitempty"`
	EventMenuURL  string            `json:"event_menu_url,omitempty"`
	EventMenuName string            `json:"event_menu_name,omitempty"`
}

// MediaTypes represents an array of MediaType objects
type MediaTypes []MediaType

// MediaTypeParam represents a media type parameter
type MediaTypeParam struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// MediaTypeGetOptions represents parameters for mediatype.get API call
type MediaTypeGetOptions struct {
	MediaTypeIDs []string             `json:"mediatypeids,omitempty"`
	Filter       map[string]interface{} `json:"filter,omitempty"`
	Search       map[string]interface{} `json:"search,omitempty"`
	SearchWildcardsEnabled string      `json:"searchWildcardsEnabled,omitempty"`
	Output       string               `json:"output,omitempty"`
	SelectParams string               `json:"selectParams,omitempty"`
	SortField    string               `json:"sortfield,omitempty"`
	SortOrder    string               `json:"sortorder,omitempty"`
	Limit        int                  `json:"limit,omitempty"`
}

// MediaType constants
const (
	MediaTypeEmail        = "0"
	MediaTypeScript       = "1"
	MediaTypeSMS          = "2"
	MediaTypeJabber       = "3"
	MediaTypeEzTexting    = "4"
	MediaTypeWebhook      = "5" // Added in Zabbix 4.4+
)

// MediaType status constants
const (
	MediaTypeStatusEnabled  = "0"
	MediaTypeStatusDisabled = "1"
)

// MediaTypesGet Wrapper for mediatype.get
// https://www.zabbix.com/documentation/current/manual/api/reference/mediatype/get
//
// Zabbix 6.0 Permission Notes:
// - Admin users can only access limited media type properties
// - Only Super Admin can access all media type properties
// - Limited properties for Admin users: mediatypeid, name, type, status, maxattempts
func (api *API) MediaTypesGet(options MediaTypeGetOptions) (mediatypes MediaTypes, err error) {
	params := make(map[string]interface{})
	
	// Convert options to params
	if options.MediaTypeIDs != nil {
		params["mediatypeids"] = options.MediaTypeIDs
	}
	if options.Filter != nil {
		params["filter"] = options.Filter
	}
	if options.Search != nil {
		params["search"] = options.Search
	}
	if options.SearchWildcardsEnabled != "" {
		params["searchWildcardsEnabled"] = options.SearchWildcardsEnabled
	}
	if options.Output != "" {
		params["output"] = options.Output
	} else {
		// Default to limited fields for Zabbix 6.0 Admin user compatibility
		params["output"] = []string{"mediatypeid", "name", "type", "status", "maxattempts"}
	}
	if options.SelectParams != "" {
		params["selectParams"] = options.SelectParams
	}
	if options.SortField != "" {
		params["sortfield"] = options.SortField
	}
	if options.SortOrder != "" {
		params["sortorder"] = options.SortOrder
	}
	if options.Limit > 0 {
		params["limit"] = options.Limit
	}

	err = api.CallWithErrorParse("mediatype.get", params, &mediatypes)
	return
}

// MediaTypesGetById Wrapper for mediatype.get with specific media type IDs
func (api *API) MediaTypesGetById(mediaTypeIds []string) (mediatypes MediaTypes, err error) {
	options := MediaTypeGetOptions{
		MediaTypeIDs: mediaTypeIds,
		Output: "extend", // Try to get all fields, will be limited by Zabbix 6.0 permissions
	}
	return api.MediaTypesGet(options)
}

// MediaTypeGetByName Wrapper for mediatype.get with name filter
func (api *API) MediaTypeGetByName(name string) (mediatypes MediaTypes, err error) {
	options := MediaTypeGetOptions{
		Filter: map[string]interface{}{
			"name": name,
		},
		Output: "extend",
	}
	return api.MediaTypesGet(options)
}

// MediaTypesGetByType Wrapper for mediatype.get with type filter
func (api *API) MediaTypesGetByType(mediaType string) (mediatypes MediaTypes, err error) {
	options := MediaTypeGetOptions{
		Filter: map[string]interface{}{
			"type": mediaType,
		},
		Output: "extend",
	}
	return api.MediaTypesGet(options)
}

// MediaTypeCreate Wrapper for mediatype.create
// https://www.zabbix.com/documentation/current/manual/api/reference/mediatype/create
//
// Zabbix 6.0 Permission Notes:
// - Only Super Admin can create media types
// - Admin users will receive permission denied errors
func (api *API) MediaTypeCreate(mediatypes MediaTypes) (result []string, err error) {
	response, err := api.CallWithError("mediatype.create", mediatypes)
	if err != nil {
		return
	}

	if resultArray, ok := response.Result.([]interface{}); ok {
		for _, item := range resultArray {
			if mediaTypeMap, ok := item.(map[string]interface{}); ok {
				if mediatypeids, exists := mediaTypeMap["mediatypeids"]; exists {
					if idArray, ok := mediatypeids.([]interface{}); ok && len(idArray) > 0 {
						if id, ok := idArray[0].(string); ok {
							result = append(result, id)
						}
					}
				}
			}
		}
	}
	return
}

// MediaTypeUpdate Wrapper for mediatype.update
// https://www.zabbix.com/documentation/current/manual/api/reference/mediatype/update
//
// Zabbix 6.0 Permission Notes:
// - Only Super Admin can update media types
// - Admin users will receive permission denied errors
func (api *API) MediaTypeUpdate(mediatypes MediaTypes) (result []string, err error) {
	response, err := api.CallWithError("mediatype.update", mediatypes)
	if err != nil {
		return
	}

	if resultArray, ok := response.Result.([]interface{}); ok {
		for _, item := range resultArray {
			if mediaTypeMap, ok := item.(map[string]interface{}); ok {
				if mediatypeids, exists := mediaTypeMap["mediatypeids"]; exists {
					if idArray, ok := mediatypeids.([]interface{}); ok && len(idArray) > 0 {
						if id, ok := idArray[0].(string); ok {
							result = append(result, id)
						}
					}
				}
			}
		}
	}
	return
}

// MediaTypeDelete Wrapper for mediatype.delete
// https://www.zabbix.com/documentation/current/manual/api/reference/mediatype/delete
//
// Zabbix 6.0 Permission Notes:
// - Only Super Admin can delete media types
// - Admin users will receive permission denied errors
func (api *API) MediaTypeDelete(mediaTypeIds []string) (result []string, err error) {
	response, err := api.CallWithError("mediatype.delete", mediaTypeIds)
	if err != nil {
		return
	}

	if resultArray, ok := response.Result.([]interface{}); ok {
		for _, item := range resultArray {
			if id, ok := item.(string); ok {
				result = append(result, id)
			}
		}
	}
	return
}

// MediaTypesGetEnabled Wrapper for mediatype.get with enabled status filter
func (api *API) MediaTypesGetEnabled() (mediatypes MediaTypes, err error) {
	options := MediaTypeGetOptions{
		Filter: map[string]interface{}{
			"status": MediaTypeStatusEnabled,
		},
		Output: "extend",
	}
	return api.MediaTypesGet(options)
}

// MediaTypesGetEmail Wrapper for mediatype.get with email type filter
func (api *API) MediaTypesGetEmail() (mediatypes MediaTypes, err error) {
	options := MediaTypeGetOptions{
		Filter: map[string]interface{}{
			"type": MediaTypeEmail,
		},
		Output: "extend",
	}
	return api.MediaTypesGet(options)
}

// MediaTypesGetWebhook Wrapper for mediatype.get with webhook type filter
func (api *API) MediaTypesGetWebhook() (mediatypes MediaTypes, err error) {
	options := MediaTypeGetOptions{
		Filter: map[string]interface{}{
			"type": MediaTypeWebhook,
		},
		Output: "extend",
	}
	return api.MediaTypesGet(options)
}