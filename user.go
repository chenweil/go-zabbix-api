package zabbix

import "encoding/json"

// User represent Zabbix User object
// https://www.zabbix.com/documentation/current/manual/api/reference/user/object
type User struct {
	UserID      string      `json:"userid,omitempty"`
	Username    string      `json:"username,omitempty"`
	Name        string      `json:"name,omitempty"`
	Surname     string      `json:"surname,omitempty"`
	Password    string      `json:"password,omitempty"`
	Url         string      `json:"url,omitempty"` // Max length increased to 2048 in Zabbix 6.0
	Autologout  string      `json:"autologout,omitempty"`
	Autologin   string      `json:"autologin,omitempty"`
	Theme       string      `json:"theme,omitempty"`
	Lang        string      `json:"lang,omitempty"`
	Refresh     string      `json:"refresh,omitempty"`
	RowsPerPage string      `json:"rows_per_page,omitempty"`
	Timezone    string      `json:"timezone,omitempty"`
	RoleID      string      `json:"roleid,omitempty"`
	UserGroups  []UserGroup `json:"usrgrps,omitempty"`
	Medias      []Media     `json:"medias,omitempty"`
	Alias       string      `json:"alias,omitempty"` // Deprecated in Zabbix 6.0, use username instead
	Type        string      `json:"type,omitempty"`  // Deprecated in Zabbix 6.0, use roleid instead

	// Zabbix 7.0+ MFA fields
	MFAStatus  string `json:"mfa_status,omitempty"`  // MFA status: 0 = disabled, 1 = enabled
	MFAID      string `json:"mfaid,omitempty"`       // MFA configuration ID
	TOTPSecret string `json:"totp_secret,omitempty"` // TOTP secret for MFA
}

// Users represents an array of User objects
type Users []User

// UserGroup represents a user group
type UserGroup struct {
	UsrGroupID string `json:"usrgrpid,omitempty"`
	Name       string `json:"name,omitempty"`
}

// Media represents a user media
type Media struct {
	MediaID     string   `json:"mediaid,omitempty"`
	UserID      string   `json:"userid,omitempty"`
	MediaTypeID string   `json:"mediatypeid,omitempty"`
	SendTo      []string `json:"sendto,omitempty"`
	Active      string   `json:"active,omitempty"`
	Severity    string   `json:"severity,omitempty"`
	Period      string   `json:"period,omitempty"`
}

// UserGetOptions represents parameters for user.get API call
type UserGetOptions struct {
	UserIDs                []string               `json:"userids,omitempty"`
	Filter                 map[string]interface{} `json:"filter,omitempty"`
	Search                 map[string]interface{} `json:"search,omitempty"`
	SearchWildcardsEnabled string                 `json:"searchWildcardsEnabled,omitempty"`
	Output                 string                 `json:"output,omitempty"`
	SelectUsrGrps          string                 `json:"selectUsrGrps,omitempty"`
	SelectMedias           string                 `json:"selectMedias,omitempty"`
	SortField              string                 `json:"sortfield,omitempty"`
	SortOrder              string                 `json:"sortorder,omitempty"`
	Limit                  int                    `json:"limit,omitempty"`
}

// UsersGet Wrapper for user.get
// https://www.zabbix.com/documentation/current/manual/api/reference/user/get
//
// Zabbix 6.0 Permission Notes:
// - Admin and User roles can only access limited user properties
// - Only Super Admin can access all user properties
// - Limited properties include: userid, username, name, surname, roleid, usrgrps
func (api *API) UsersGet(options UserGetOptions) (users Users, err error) {
	params := make(map[string]interface{})

	// Convert options to params
	if options.UserIDs != nil {
		params["userids"] = options.UserIDs
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
		// Default to limited fields for Zabbix 6.0 compatibility
		params["output"] = []string{"userid", "username", "name", "surname", "roleid"}
	}
	if options.SelectUsrGrps != "" {
		params["selectUsrGrps"] = options.SelectUsrGrps
	}
	if options.SelectMedias != "" {
		params["selectMedias"] = options.SelectMedias
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

	err = api.CallWithErrorParse("user.get", params, &users)
	return
}

// UsersGetById Wrapper for user.get with specific user IDs
func (api *API) UsersGetById(userIds []string) (users Users, err error) {
	options := UserGetOptions{
		UserIDs: userIds,
		Output:  "extend", // Try to get all fields, will be limited by Zabbix 6.0 permissions
	}
	return api.UsersGet(options)
}

// UserGetByUsername Wrapper for user.get with username filter
func (api *API) UserGetByUsername(username string) (users Users, err error) {
	options := UserGetOptions{
		Filter: map[string]interface{}{
			"username": username,
		},
		Output: "extend",
	}
	return api.UsersGet(options)
}

// UserCreate Wrapper for user.create
// https://www.zabbix.com/documentation/current/manual/api/reference/user/create
func (api *API) UserCreate(users Users) (result []string, err error) {
	response, err := api.CallWithError("user.create", users)
	if err != nil {
		return
	}

	var rawResult interface{}
	err = json.Unmarshal(response.Result, &rawResult)
	if err != nil {
		return
	}

	if resultArray, ok := rawResult.([]interface{}); ok {
		for _, item := range resultArray {
			if userMap, ok := item.(map[string]interface{}); ok {
				if userid, exists := userMap["userids"]; exists {
					if idArray, ok := userid.([]interface{}); ok && len(idArray) > 0 {
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

// UserUpdate Wrapper for user.update
// https://www.zabbix.com/documentation/current/manual/api/reference/user/update
func (api *API) UserUpdate(users Users) (result []string, err error) {
	response, err := api.CallWithError("user.update", users)
	if err != nil {
		return
	}

	var rawResult interface{}
	err = json.Unmarshal(response.Result, &rawResult)
	if err != nil {
		return
	}

	if resultArray, ok := rawResult.([]interface{}); ok {
		for _, item := range resultArray {
			if userMap, ok := item.(map[string]interface{}); ok {
				if userid, exists := userMap["userids"]; exists {
					if idArray, ok := userid.([]interface{}); ok && len(idArray) > 0 {
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

// UserDelete Wrapper for user.delete
// https://www.zabbix.com/documentation/current/manual/api/reference/user/delete
func (api *API) UserDelete(userIds []string) (result []string, err error) {
	response, err := api.CallWithError("user.delete", userIds)
	if err != nil {
		return
	}

	var rawResult interface{}
	err = json.Unmarshal(response.Result, &rawResult)
	if err != nil {
		return
	}

	if resultArray, ok := rawResult.([]interface{}); ok {
		for _, item := range resultArray {
			if id, ok := item.(string); ok {
				result = append(result, id)
			}
		}
	}
	return
}

// Login Wrapper for user.login (enhanced version for Zabbix 6.0 compatibility)
// This method extends the base Login method with additional error handling
func (api *API) LoginExtended(user, password string) (auth string, err error) {
	params := map[string]interface{}{
		"user":     user,
		"password": password,
	}

		response, err := api.CallWithError("user.login", params)

		if err != nil {

			return

		}

	

		err = json.Unmarshal(response.Result, &auth)

		if err != nil {

			return

		}

		api.Auth = auth

	

		return

	}
