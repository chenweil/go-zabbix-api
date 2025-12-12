package zabbix

import (
	"encoding/json"
	"fmt"
)

// MFA (Multi-Factor Authentication) structures for Zabbix 7.0+

// MFAType represents the type of MFA
type MFAType int

const (
	// MFATypeTOTP Time-based One-Time Password
	MFATypeTOTP MFAType = 0
)

// MFAStatus represents MFA status
type MFAStatus int

const (
	// MFAStatusDisabled MFA is disabled
	MFAStatusDisabled MFAStatus = 0
	// MFAStatusEnabled MFA is enabled
	MFAStatusEnabled MFAStatus = 1
)

// MFA represents a multi-factor authentication configuration
type MFA struct {
	MFAID        string   `json:"mfaid,omitempty"`
	Name         string   `json:"name"`
	Type         MFAType  `json:"type,string"`
	HashFunction string   `json:"hash_function,omitempty"`
	CodeLength   int      `json:"code_length,omitempty"`
	Status       MFAStatus `json:"status,string"`
	APIAccess    string   `json:"api_access,omitempty"`
	UserGroups   MFAUserGroups `json:"user_groups,omitempty"`
}

// MFAUserGroups represents user groups for MFA
type MFAUserGroups []MFAUserGroup

// MFAUserGroup represents a user group in MFA configuration
type MFAUserGroup struct {
	UserGroupID string `json:"usrgrpid,omitempty"`
	Name        string `json:"name,omitempty"`
}

// MFAs is an array of MFA
type MFAs []MFA

// MFACreate Wrapper for mfa.create
// https://www.zabbix.com/documentation/7.0/manual/api/reference/mfa/create
func (api *API) MFACreate(mfas MFAs) error {
	if !api.versionManager.IsFeatureSupported(FeatureMFA) {
		return fmt.Errorf("MFA not supported in Zabbix version %s", api.versionManager.GetVersion())
	}

	response, err := api.CallWithError("mfa.create", mfas)
	if err != nil {
		return err
	}

	result := response.Result.(map[string]interface{})
	mfaids := result["mfaids"].([]interface{})
	for i, id := range mfaids {
		mfas[i].MFAID = id.(string)
	}
	return nil
}

// MFAGet Wrapper for mfa.get
// https://www.zabbix.com/documentation/7.0/manual/api/reference/mfa/get
func (api *API) MFAGet(params Params) (MFAs, error) {
	if !api.versionManager.IsFeatureSupported(FeatureMFA) {
		return nil, fmt.Errorf("MFA not supported in Zabbix version %s", api.versionManager.GetVersion())
	}

	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	
	var res MFAs
	err := api.CallWithErrorParse("mfa.get", params, &res)
	return res, err
}

// MFAUpdate Wrapper for mfa.update
// https://www.zabbix.com/documentation/7.0/manual/api/reference/mfa/update
func (api *API) MFAUpdate(mfas MFAs) error {
	if !api.versionManager.IsFeatureSupported(FeatureMFA) {
		return fmt.Errorf("MFA not supported in Zabbix version %s", api.versionManager.GetVersion())
	}

	_, err := api.CallWithError("mfa.update", mfas)
	return err
}

// MFADelete Wrapper for mfa.delete
// https://www.zabbix.com/documentation/7.0/manual/api/reference/mfa/delete
func (api *API) MFADelete(mfas MFAs) error {
	if !api.versionManager.IsFeatureSupported(FeatureMFA) {
		return fmt.Errorf("MFA not supported in Zabbix version %s", api.versionManager.GetVersion())
	}

	ids := make([]string, len(mfas))
	for i, mfa := range mfas {
		ids[i] = mfa.MFAID
	}

	_, err := api.CallWithError("mfa.delete", ids)
	if err == nil {
		for i := range mfas {
			mfas[i].MFAID = ""
		}
	}
	return err
}

// MFAGetByID Gets MFA by ID
func (api *API) MFAGetByID(id string) (*MFA, error) {
	mfas, err := api.MFAGet(Params{"mfaids": id})
	if err != nil {
		return nil, err
	}

	if len(mfas) != 1 {
		e := ExpectedOneResult(len(mfas))
		return nil, &e
	}
	return &mfas[0], nil
}

// UserResetTOTP Wrapper for user.resetotp
// https://www.zabbix.com/documentation/7.0/manual/api/reference/user/resetotp
func (api *API) UserResetTOTP(userIDs []string) error {
	if !api.versionManager.IsFeatureSupported(FeatureMFA) {
		return fmt.Errorf("MFA not supported in Zabbix version %s", api.versionManager.GetVersion())
	}

	params := map[string][]string{"userids": userIDs}
	_, err := api.CallWithError("user.resetotp", params)
	return err
}

// UserResetTOTPByUser Resets TOTP for a single user
func (api *API) UserResetTOTPByUser(userID string) error {
	return api.UserResetTOTP([]string{userID})
}

// User represents a user with MFA fields
type User struct {
	UserID     string    `json:"userid,omitempty"`
	Username   string    `json:"username"`
	Name       string    `json:"name,omitempty"`
	Surname    string    `json:"surname,omitempty"`
	MFAStatus  MFAStatus `json:"mfa_status,string,omitempty"`
	MFAID      string    `json:"mfaid,omitempty"`
	TOTPSecret string    `json:"totp_secret,omitempty"`
}

// Users is an array of User
type Users []User

// UsersGet Wrapper for user.get with MFA support
func (api *API) UsersGet(params Params) (Users, error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	
	// Add MFA fields if supported
	if api.versionManager.IsFeatureSupported(FeatureMFA) {
		if _, present := params["selectUsrgrps"]; !present {
			params["selectUsrgrps"] = "extend"
		}
	}
	
	var res Users
	err := api.CallWithErrorParse("user.get", params, &res)
	return res, err
}

// UserGetByID Gets user by ID
func (api *API) UserGetByID(id string) (*User, error) {
	users, err := api.UsersGet(Params{"userids": id})
	if err != nil {
		return nil, err
	}

	if len(users) != 1 {
		e := ExpectedOneResult(len(users))
		return nil, &e
	}
	return &users[0], nil
}

// UserGetByUsername Gets user by username
func (api *API) UserGetByUsername(username string) (*User, error) {
	users, err := api.UsersGet(Params{"filter": map[string]string{"username": username}})
	if err != nil {
		return nil, err
	}

	if len(users) != 1 {
		e := ExpectedOneResult(len(users))
		return nil, &e
	}
	return &users[0], nil
}