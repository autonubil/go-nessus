package nessus

import (
	"io"
	"time"
)

// AgentConfig implementation of the AgentConfig interface
type AgentConfig struct {
	*ClientWithResponses
}

// Details calls the AgentConfig ´s function
func (c *AgentConfig) Details(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
	AutoUnlink *struct {
		Enabled    *bool  "json:\"enabled,omitempty\""
		Expiration *int32 "json:\"expiration,omitempty\""
	} "json:\"auto_unlink,omitempty\""
	SoftwareUpdate *bool "json:\"software_update,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentConfigDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AutoUnlink *struct { Enabled *bool "json:\"enabled,omitempty\""; Expiration *int32 "json:\"expiration,omitempty\"" } "json:\"auto_unlink,omitempty\""; SoftwareUpdate *bool "json:\"software_update,omitempty\"" }
	if i, ok := r.(*struct {
		AutoUnlink *struct {
			Enabled    *bool  "json:\"enabled,omitempty\""
			Expiration *int32 "json:\"expiration,omitempty\""
		} "json:\"auto_unlink,omitempty\""
		SoftwareUpdate *bool "json:\"software_update,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// EditWithBody calls the AgentConfig ´s function
func (c *AgentConfig) EditWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	AutoUnlink *struct {
		Enabled    *bool  "json:\"enabled,omitempty\""
		Expiration *int32 "json:\"expiration,omitempty\""
	} "json:\"auto_unlink,omitempty\""
	SoftwareUpdate *bool "json:\"software_update,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentConfigEditWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AutoUnlink *struct { Enabled *bool "json:\"enabled,omitempty\""; Expiration *int32 "json:\"expiration,omitempty\"" } "json:\"auto_unlink,omitempty\""; SoftwareUpdate *bool "json:\"software_update,omitempty\"" }
	if i, ok := r.(*struct {
		AutoUnlink *struct {
			Enabled    *bool  "json:\"enabled,omitempty\""
			Expiration *int32 "json:\"expiration,omitempty\""
		} "json:\"auto_unlink,omitempty\""
		SoftwareUpdate *bool "json:\"software_update,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Edit calls the AgentConfig ´s function
func (c *AgentConfig) Edit(arg1 int32, arg2 AgentConfigEditJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	AutoUnlink *struct {
		Enabled    *bool  "json:\"enabled,omitempty\""
		Expiration *int32 "json:\"expiration,omitempty\""
	} "json:\"auto_unlink,omitempty\""
	SoftwareUpdate *bool "json:\"software_update,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentConfigEditWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AutoUnlink *struct { Enabled *bool "json:\"enabled,omitempty\""; Expiration *int32 "json:\"expiration,omitempty\"" } "json:\"auto_unlink,omitempty\""; SoftwareUpdate *bool "json:\"software_update,omitempty\"" }
	if i, ok := r.(*struct {
		AutoUnlink *struct {
			Enabled    *bool  "json:\"enabled,omitempty\""
			Expiration *int32 "json:\"expiration,omitempty\""
		} "json:\"auto_unlink,omitempty\""
		SoftwareUpdate *bool "json:\"software_update,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AgentGroup implementation of the AgentGroup interface
type AgentGroup struct {
	*ClientWithResponses
}

// ListAgents calls the AgentGroup ´s function
func (c *AgentGroup) ListAgents(arg1 int32, arg2 int32, params *AgentGroupListAgentsParams, reqEditors ...RequestEditorFn) (*struct {
	Agents *[]struct {
		CoreBuild   *string "json:\"core_build,omitempty\""
		CoreVersion *string "json:\"core_version,omitempty\""
		Distro      *string "json:\"distro,omitempty\""
		Groups      *[]struct {
			Id   *int    "json:\"id,omitempty\""
			Name *string "json:\"name,omitempty\""
		} "json:\"groups,omitempty\""
		Id           *int              "json:\"id,omitempty\""
		Ip           *string           "json:\"ip,omitempty\""
		LastConnect  *int              "json:\"last_connect,omitempty\""
		LastScanned  *int              "json:\"last_scanned,omitempty\""
		LinkedOn     *int              "json:\"linked_on,omitempty\""
		Name         *string           "json:\"name,omitempty\""
		NetworkName  *string           "json:\"network_name,omitempty\""
		NetworkUuid  *string           "json:\"network_uuid,omitempty\""
		Platform     *string           "json:\"platform,omitempty\""
		PluginFeedId *string           "json:\"plugin_feed_id,omitempty\""
		Status       *N200AgentsStatus "json:\"status,omitempty\""
		Uuid         *string           "json:\"uuid,omitempty\""
	} "json:\"agents,omitempty\""
	Pagination *struct {
		Limit  *int32 "json:\"limit,omitempty\""
		Offset *int32 "json:\"offset,omitempty\""
		Sort   *[]struct {
			Name  *string "json:\"name,omitempty\""
			Order *string "json:\"order,omitempty\""
		} "json:\"sort,omitempty\""
		Total *int32 "json:\"total,omitempty\""
	} "json:\"pagination,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentGroupListAgentsWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Agents *[]struct { CoreBuild *string "json:\"core_build,omitempty\""; CoreVersion *string "json:\"core_version,omitempty\""; Distro *string "json:\"distro,omitempty\""; Groups *[]struct { Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"groups,omitempty\""; Id *int "json:\"id,omitempty\""; Ip *string "json:\"ip,omitempty\""; LastConnect *int "json:\"last_connect,omitempty\""; LastScanned *int "json:\"last_scanned,omitempty\""; LinkedOn *int "json:\"linked_on,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkName *string "json:\"network_name,omitempty\""; NetworkUuid *string "json:\"network_uuid,omitempty\""; Platform *string "json:\"platform,omitempty\""; PluginFeedId *string "json:\"plugin_feed_id,omitempty\""; Status *N200AgentsStatus "json:\"status,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"agents,omitempty\""; Pagination *struct { Limit *int32 "json:\"limit,omitempty\""; Offset *int32 "json:\"offset,omitempty\""; Sort *[]struct { Name *string "json:\"name,omitempty\""; Order *string "json:\"order,omitempty\"" } "json:\"sort,omitempty\""; Total *int32 "json:\"total,omitempty\"" } "json:\"pagination,omitempty\"" }
	if i, ok := r.(*struct {
		Agents *[]struct {
			CoreBuild   *string "json:\"core_build,omitempty\""
			CoreVersion *string "json:\"core_version,omitempty\""
			Distro      *string "json:\"distro,omitempty\""
			Groups      *[]struct {
				Id   *int    "json:\"id,omitempty\""
				Name *string "json:\"name,omitempty\""
			} "json:\"groups,omitempty\""
			Id           *int              "json:\"id,omitempty\""
			Ip           *string           "json:\"ip,omitempty\""
			LastConnect  *int              "json:\"last_connect,omitempty\""
			LastScanned  *int              "json:\"last_scanned,omitempty\""
			LinkedOn     *int              "json:\"linked_on,omitempty\""
			Name         *string           "json:\"name,omitempty\""
			NetworkName  *string           "json:\"network_name,omitempty\""
			NetworkUuid  *string           "json:\"network_uuid,omitempty\""
			Platform     *string           "json:\"platform,omitempty\""
			PluginFeedId *string           "json:\"plugin_feed_id,omitempty\""
			Status       *N200AgentsStatus "json:\"status,omitempty\""
			Uuid         *string           "json:\"uuid,omitempty\""
		} "json:\"agents,omitempty\""
		Pagination *struct {
			Limit  *int32 "json:\"limit,omitempty\""
			Offset *int32 "json:\"offset,omitempty\""
			Sort   *[]struct {
				Name  *string "json:\"name,omitempty\""
				Order *string "json:\"order,omitempty\""
			} "json:\"sort,omitempty\""
			Total *int32 "json:\"total,omitempty\""
		} "json:\"pagination,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// sAddAgent calls the AgentGroup ´s function
func (c *AgentGroup) sAddAgent(arg1 int32, arg2 int32, arg3 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentGroupsAddAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// sConfigureWithBody calls the AgentGroup ´s function
func (c *AgentGroup) sConfigureWithBody(arg1 int32, arg2 int32, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentGroupsConfigureWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// sConfigure calls the AgentGroup ´s function
func (c *AgentGroup) sConfigure(arg1 int32, arg2 int32, arg3 AgentGroupsConfigureJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentGroupsConfigureWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// sCreateWithBody calls the AgentGroup ´s function
func (c *AgentGroup) sCreateWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	Agents               *[]string "json:\"agents,omitempty\""
	CreationDate         *int      "json:\"creation_date,omitempty\""
	Id                   *int      "json:\"id,omitempty\""
	LastModificationDate *int      "json:\"last_modification_date,omitempty\""
	Name                 *string   "json:\"name,omitempty\""
	Owner                *string   "json:\"owner,omitempty\""
	OwnerId              *string   "json:\"owner_id,omitempty\""
	OwnerName            *string   "json:\"owner_name,omitempty\""
	OwnerUuid            *string   "json:\"owner_uuid,omitempty\""
	Shared               *bool     "json:\"shared,omitempty\""
	UserPermissions      *int      "json:\"user_permissions,omitempty\""
	Uuid                 *string   "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentGroupsCreateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Agents *[]string "json:\"agents,omitempty\""; CreationDate *int "json:\"creation_date,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *string "json:\"owner_id,omitempty\""; OwnerName *string "json:\"owner_name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		Agents               *[]string "json:\"agents,omitempty\""
		CreationDate         *int      "json:\"creation_date,omitempty\""
		Id                   *int      "json:\"id,omitempty\""
		LastModificationDate *int      "json:\"last_modification_date,omitempty\""
		Name                 *string   "json:\"name,omitempty\""
		Owner                *string   "json:\"owner,omitempty\""
		OwnerId              *string   "json:\"owner_id,omitempty\""
		OwnerName            *string   "json:\"owner_name,omitempty\""
		OwnerUuid            *string   "json:\"owner_uuid,omitempty\""
		Shared               *bool     "json:\"shared,omitempty\""
		UserPermissions      *int      "json:\"user_permissions,omitempty\""
		Uuid                 *string   "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// sCreate calls the AgentGroup ´s function
func (c *AgentGroup) sCreate(arg1 int32, arg2 AgentGroupsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	Agents               *[]string "json:\"agents,omitempty\""
	CreationDate         *int      "json:\"creation_date,omitempty\""
	Id                   *int      "json:\"id,omitempty\""
	LastModificationDate *int      "json:\"last_modification_date,omitempty\""
	Name                 *string   "json:\"name,omitempty\""
	Owner                *string   "json:\"owner,omitempty\""
	OwnerId              *string   "json:\"owner_id,omitempty\""
	OwnerName            *string   "json:\"owner_name,omitempty\""
	OwnerUuid            *string   "json:\"owner_uuid,omitempty\""
	Shared               *bool     "json:\"shared,omitempty\""
	UserPermissions      *int      "json:\"user_permissions,omitempty\""
	Uuid                 *string   "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentGroupsCreateWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Agents *[]string "json:\"agents,omitempty\""; CreationDate *int "json:\"creation_date,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *string "json:\"owner_id,omitempty\""; OwnerName *string "json:\"owner_name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		Agents               *[]string "json:\"agents,omitempty\""
		CreationDate         *int      "json:\"creation_date,omitempty\""
		Id                   *int      "json:\"id,omitempty\""
		LastModificationDate *int      "json:\"last_modification_date,omitempty\""
		Name                 *string   "json:\"name,omitempty\""
		Owner                *string   "json:\"owner,omitempty\""
		OwnerId              *string   "json:\"owner_id,omitempty\""
		OwnerName            *string   "json:\"owner_name,omitempty\""
		OwnerUuid            *string   "json:\"owner_uuid,omitempty\""
		Shared               *bool     "json:\"shared,omitempty\""
		UserPermissions      *int      "json:\"user_permissions,omitempty\""
		Uuid                 *string   "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// sDeleteAgent calls the AgentGroup ´s function
func (c *AgentGroup) sDeleteAgent(arg1 int32, arg2 int32, arg3 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentGroupsDeleteAgentWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// sDelete calls the AgentGroup ´s function
func (c *AgentGroup) sDelete(arg1 int32, arg2 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentGroupsDeleteWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// sDetails calls the AgentGroup ´s function
func (c *AgentGroup) sDetails(arg1 int32, arg2 int32, params *AgentGroupsDetailsParams, reqEditors ...RequestEditorFn) (*struct {
	Agents               *[]string "json:\"agents,omitempty\""
	CreationDate         *int      "json:\"creation_date,omitempty\""
	Id                   *int      "json:\"id,omitempty\""
	LastModificationDate *int      "json:\"last_modification_date,omitempty\""
	Name                 *string   "json:\"name,omitempty\""
	Owner                *string   "json:\"owner,omitempty\""
	OwnerId              *string   "json:\"owner_id,omitempty\""
	OwnerName            *string   "json:\"owner_name,omitempty\""
	OwnerUuid            *string   "json:\"owner_uuid,omitempty\""
	Shared               *bool     "json:\"shared,omitempty\""
	UserPermissions      *int      "json:\"user_permissions,omitempty\""
	Uuid                 *string   "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentGroupsDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Agents *[]string "json:\"agents,omitempty\""; CreationDate *int "json:\"creation_date,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *string "json:\"owner_id,omitempty\""; OwnerName *string "json:\"owner_name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		Agents               *[]string "json:\"agents,omitempty\""
		CreationDate         *int      "json:\"creation_date,omitempty\""
		Id                   *int      "json:\"id,omitempty\""
		LastModificationDate *int      "json:\"last_modification_date,omitempty\""
		Name                 *string   "json:\"name,omitempty\""
		Owner                *string   "json:\"owner,omitempty\""
		OwnerId              *string   "json:\"owner_id,omitempty\""
		OwnerName            *string   "json:\"owner_name,omitempty\""
		OwnerUuid            *string   "json:\"owner_uuid,omitempty\""
		Shared               *bool     "json:\"shared,omitempty\""
		UserPermissions      *int      "json:\"user_permissions,omitempty\""
		Uuid                 *string   "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// sList calls the AgentGroup ´s function
func (c *AgentGroup) sList(arg1 int32, reqEditors ...RequestEditorFn) (*[]struct {
	Agents               *[]string "json:\"agents,omitempty\""
	CreationDate         *int      "json:\"creation_date,omitempty\""
	Id                   *int      "json:\"id,omitempty\""
	LastModificationDate *int      "json:\"last_modification_date,omitempty\""
	Name                 *string   "json:\"name,omitempty\""
	Owner                *string   "json:\"owner,omitempty\""
	OwnerId              *string   "json:\"owner_id,omitempty\""
	OwnerName            *string   "json:\"owner_name,omitempty\""
	OwnerUuid            *string   "json:\"owner_uuid,omitempty\""
	Shared               *bool     "json:\"shared,omitempty\""
	UserPermissions      *int      "json:\"user_permissions,omitempty\""
	Uuid                 *string   "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentGroupsListWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { Agents *[]string "json:\"agents,omitempty\""; CreationDate *int "json:\"creation_date,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *string "json:\"owner_id,omitempty\""; OwnerName *string "json:\"owner_name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*[]struct {
		Agents               *[]string "json:\"agents,omitempty\""
		CreationDate         *int      "json:\"creation_date,omitempty\""
		Id                   *int      "json:\"id,omitempty\""
		LastModificationDate *int      "json:\"last_modification_date,omitempty\""
		Name                 *string   "json:\"name,omitempty\""
		Owner                *string   "json:\"owner,omitempty\""
		OwnerId              *string   "json:\"owner_id,omitempty\""
		OwnerName            *string   "json:\"owner_name,omitempty\""
		OwnerUuid            *string   "json:\"owner_uuid,omitempty\""
		Shared               *bool     "json:\"shared,omitempty\""
		UserPermissions      *int      "json:\"user_permissions,omitempty\""
		Uuid                 *string   "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AuditLog implementation of the AuditLog interface
type AuditLog struct {
	*ClientWithResponses
}

// Events calls the AuditLog ´s function
func (c *AuditLog) Events(params *AuditLogEventsParams, reqEditors ...RequestEditorFn) (*[]struct {
	Action *string "json:\"action,omitempty\""
	Actor  *struct {
		Id   *string "json:\"id,omitempty\""
		Name *string "json:\"name,omitempty\""
	} "json:\"actor,omitempty\""
	Crud        *string "json:\"crud,omitempty\""
	Description *string "json:\"description,omitempty\""
	Fields      *struct {
		Pair *struct {
			Key   *string "json:\"key,omitempty\""
			Value *string "json:\"value,omitempty\""
		} "json:\"pair,omitempty\""
	} "json:\"fields,omitempty\""
	Id          *string "json:\"id,omitempty\""
	IsAnonymous *bool   "json:\"is_anonymous,omitempty\""
	IsFailure   *bool   "json:\"is_failure,omitempty\""
	Received    *string "json:\"received,omitempty\""
	Target      *struct {
		Id   *string "json:\"id,omitempty\""
		Name *string "json:\"name,omitempty\""
		Type *string "json:\"type,omitempty\""
	} "json:\"target,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AuditLogEventsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { Action *string "json:\"action,omitempty\""; Actor *struct { Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"actor,omitempty\""; Crud *string "json:\"crud,omitempty\""; Description *string "json:\"description,omitempty\""; Fields *struct { Pair *struct { Key *string "json:\"key,omitempty\""; Value *string "json:\"value,omitempty\"" } "json:\"pair,omitempty\"" } "json:\"fields,omitempty\""; Id *string "json:\"id,omitempty\""; IsAnonymous *bool "json:\"is_anonymous,omitempty\""; IsFailure *bool "json:\"is_failure,omitempty\""; Received *string "json:\"received,omitempty\""; Target *struct { Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"target,omitempty\"" }
	if i, ok := r.(*[]struct {
		Action *string "json:\"action,omitempty\""
		Actor  *struct {
			Id   *string "json:\"id,omitempty\""
			Name *string "json:\"name,omitempty\""
		} "json:\"actor,omitempty\""
		Crud        *string "json:\"crud,omitempty\""
		Description *string "json:\"description,omitempty\""
		Fields      *struct {
			Pair *struct {
				Key   *string "json:\"key,omitempty\""
				Value *string "json:\"value,omitempty\""
			} "json:\"pair,omitempty\""
		} "json:\"fields,omitempty\""
		Id          *string "json:\"id,omitempty\""
		IsAnonymous *bool   "json:\"is_anonymous,omitempty\""
		IsFailure   *bool   "json:\"is_failure,omitempty\""
		Received    *string "json:\"received,omitempty\""
		Target      *struct {
			Id   *string "json:\"id,omitempty\""
			Name *string "json:\"name,omitempty\""
			Type *string "json:\"type,omitempty\""
		} "json:\"target,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Editor implementation of the Editor interface
type Editor struct {
	*ClientWithResponses
}

// Audits calls the Editor ´s function
func (c *Editor) Audits(arg1 EditorAuditsParamsType, arg2 int32, arg3 int32, reqEditors ...RequestEditorFn) (*EditorAuditsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.EditorAuditsWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *EditorAuditsResponse
	if i, ok := r.(*EditorAuditsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Details calls the Editor ´s function
func (c *Editor) Details(arg1 EditorDetailsParamsType, arg2 int32, reqEditors ...RequestEditorFn) (*struct {
	Compliance       *map[string]interface{} "json:\"compliance,omitempty\""
	Credentials      *map[string]interface{} "json:\"credentials,omitempty\""
	FilterAttributes *[]struct {
		Control *struct {
			Options        *[]map[string]interface{} "json:\"options,omitempty\""
			ReadableRegest *string                   "json:\"readable_regest,omitempty\""
			Regex          *string                   "json:\"regex,omitempty\""
			Type           *string                   "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		Name         *string                   "json:\"name,omitempty\""
		Operators    *[]map[string]interface{} "json:\"operators,omitempty\""
		ReadableName *string                   "json:\"readable_name,omitempty\""
	} "json:\"filter_attributes,omitempty\""
	IsAgent  *bool                   "json:\"is_agent,omitempty\""
	IsWas    *bool                   "json:\"is_was,omitempty\""
	Name     *string                 "json:\"name,omitempty\""
	Owner    *string                 "json:\"owner,omitempty\""
	Plugins  *map[string]interface{} "json:\"plugins,omitempty\""
	Settings *struct {
		Advanced   *map[string]interface{} "json:\"advanced,omitempty\""
		Assessment *map[string]interface{} "json:\"assessment,omitempty\""
		Basic      *map[string]interface{} "json:\"basic,omitempty\""
		Discovery  *map[string]interface{} "json:\"discovery,omitempty\""
	} "json:\"settings,omitempty\""
	Title           *string "json:\"title,omitempty\""
	UserPermissions *int    "json:\"user_permissions,omitempty\""
	Uuid            *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.EditorDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Compliance *map[string]interface {} "json:\"compliance,omitempty\""; Credentials *map[string]interface {} "json:\"credentials,omitempty\""; FilterAttributes *[]struct { Control *struct { Options *[]map[string]interface {} "json:\"options,omitempty\""; ReadableRegest *string "json:\"readable_regest,omitempty\""; Regex *string "json:\"regex,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"control,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]map[string]interface {} "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filter_attributes,omitempty\""; IsAgent *bool "json:\"is_agent,omitempty\""; IsWas *bool "json:\"is_was,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; Plugins *map[string]interface {} "json:\"plugins,omitempty\""; Settings *struct { Advanced *map[string]interface {} "json:\"advanced,omitempty\""; Assessment *map[string]interface {} "json:\"assessment,omitempty\""; Basic *map[string]interface {} "json:\"basic,omitempty\""; Discovery *map[string]interface {} "json:\"discovery,omitempty\"" } "json:\"settings,omitempty\""; Title *string "json:\"title,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		Compliance       *map[string]interface{} "json:\"compliance,omitempty\""
		Credentials      *map[string]interface{} "json:\"credentials,omitempty\""
		FilterAttributes *[]struct {
			Control *struct {
				Options        *[]map[string]interface{} "json:\"options,omitempty\""
				ReadableRegest *string                   "json:\"readable_regest,omitempty\""
				Regex          *string                   "json:\"regex,omitempty\""
				Type           *string                   "json:\"type,omitempty\""
			} "json:\"control,omitempty\""
			Name         *string                   "json:\"name,omitempty\""
			Operators    *[]map[string]interface{} "json:\"operators,omitempty\""
			ReadableName *string                   "json:\"readable_name,omitempty\""
		} "json:\"filter_attributes,omitempty\""
		IsAgent  *bool                   "json:\"is_agent,omitempty\""
		IsWas    *bool                   "json:\"is_was,omitempty\""
		Name     *string                 "json:\"name,omitempty\""
		Owner    *string                 "json:\"owner,omitempty\""
		Plugins  *map[string]interface{} "json:\"plugins,omitempty\""
		Settings *struct {
			Advanced   *map[string]interface{} "json:\"advanced,omitempty\""
			Assessment *map[string]interface{} "json:\"assessment,omitempty\""
			Basic      *map[string]interface{} "json:\"basic,omitempty\""
			Discovery  *map[string]interface{} "json:\"discovery,omitempty\""
		} "json:\"settings,omitempty\""
		Title           *string "json:\"title,omitempty\""
		UserPermissions *int    "json:\"user_permissions,omitempty\""
		Uuid            *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ListTemplates calls the Editor ´s function
func (c *Editor) ListTemplates(arg1 EditorListTemplatesParamsType, reqEditors ...RequestEditorFn) (*struct {
	Templates []struct {
		CloudOnly        *bool   "json:\"cloud_only,omitempty\""
		Desc             *string "json:\"desc,omitempty\""
		IsAgent          *bool   "json:\"is_agent,omitempty\""
		IsWas            *bool   "json:\"is_was,omitempty\""
		ManagerOnly      *bool   "json:\"manager_only,omitempty\""
		Name             *string "json:\"name,omitempty\""
		SubscriptionOnly *bool   "json:\"subscription_only,omitempty\""
		Title            *string "json:\"title,omitempty\""
		Unsupported      *bool   "json:\"unsupported,omitempty\""
		Uuid             *string "json:\"uuid,omitempty\""
	} "json:\"templates,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.EditorListTemplatesWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { CloudOnly *bool "json:\"cloud_only,omitempty\""; Desc *string "json:\"desc,omitempty\""; IsAgent *bool "json:\"is_agent,omitempty\""; IsWas *bool "json:\"is_was,omitempty\""; ManagerOnly *bool "json:\"manager_only,omitempty\""; Name *string "json:\"name,omitempty\""; SubscriptionOnly *bool "json:\"subscription_only,omitempty\""; Title *string "json:\"title,omitempty\""; Unsupported *bool "json:\"unsupported,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		Templates []struct {
			CloudOnly        *bool   "json:\"cloud_only,omitempty\""
			Desc             *string "json:\"desc,omitempty\""
			IsAgent          *bool   "json:\"is_agent,omitempty\""
			IsWas            *bool   "json:\"is_was,omitempty\""
			ManagerOnly      *bool   "json:\"manager_only,omitempty\""
			Name             *string "json:\"name,omitempty\""
			SubscriptionOnly *bool   "json:\"subscription_only,omitempty\""
			Title            *string "json:\"title,omitempty\""
			Unsupported      *bool   "json:\"unsupported,omitempty\""
			Uuid             *string "json:\"uuid,omitempty\""
		} "json:\"templates,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PluginDescription calls the Editor ´s function
func (c *Editor) PluginDescription(arg1 int32, arg2 int32, arg3 int32, reqEditors ...RequestEditorFn) (*struct {
	Plugindescription *struct {
		Pluginattributes *map[string]interface{} "json:\"pluginattributes,omitempty\""
		Pluginfamily     *string                 "json:\"pluginfamily,omitempty\""
		Pluginid         *int                    "json:\"pluginid,omitempty\""
		Pluginname       *string                 "json:\"pluginname,omitempty\""
		Severity         *string                 "json:\"severity,omitempty\""
	} "json:\"plugindescription,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.EditorPluginDescriptionWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Plugindescription *struct { Pluginattributes *map[string]interface {} "json:\"pluginattributes,omitempty\""; Pluginfamily *string "json:\"pluginfamily,omitempty\""; Pluginid *int "json:\"pluginid,omitempty\""; Pluginname *string "json:\"pluginname,omitempty\""; Severity *string "json:\"severity,omitempty\"" } "json:\"plugindescription,omitempty\"" }
	if i, ok := r.(*struct {
		Plugindescription *struct {
			Pluginattributes *map[string]interface{} "json:\"pluginattributes,omitempty\""
			Pluginfamily     *string                 "json:\"pluginfamily,omitempty\""
			Pluginid         *int                    "json:\"pluginid,omitempty\""
			Pluginname       *string                 "json:\"pluginname,omitempty\""
			Severity         *string                 "json:\"severity,omitempty\""
		} "json:\"plugindescription,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// TemplateDetails calls the Editor ´s function
func (c *Editor) TemplateDetails(arg1 EditorTemplateDetailsParamsType, contentType string, reqEditors ...RequestEditorFn) (*struct {
	Compliance       *map[string]interface{} "json:\"compliance,omitempty\""
	Credentials      *map[string]interface{} "json:\"credentials,omitempty\""
	FilterAttributes *[]struct {
		Control *struct {
			Options        *[]map[string]interface{} "json:\"options,omitempty\""
			ReadableRegest *string                   "json:\"readable_regest,omitempty\""
			Regex          *string                   "json:\"regex,omitempty\""
			Type           *string                   "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		Name         *string                   "json:\"name,omitempty\""
		Operators    *[]map[string]interface{} "json:\"operators,omitempty\""
		ReadableName *string                   "json:\"readable_name,omitempty\""
	} "json:\"filter_attributes,omitempty\""
	IsAgent  *bool                   "json:\"is_agent,omitempty\""
	IsWas    *bool                   "json:\"is_was,omitempty\""
	Name     *string                 "json:\"name,omitempty\""
	Plugins  *map[string]interface{} "json:\"plugins,omitempty\""
	Settings *struct {
		Advanced   *map[string]interface{} "json:\"advanced,omitempty\""
		Assessment *map[string]interface{} "json:\"assessment,omitempty\""
		Basic      *map[string]interface{} "json:\"basic,omitempty\""
		Discovery  *map[string]interface{} "json:\"discovery,omitempty\""
	} "json:\"settings,omitempty\""
	Title *string "json:\"title,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.EditorTemplateDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Compliance *map[string]interface {} "json:\"compliance,omitempty\""; Credentials *map[string]interface {} "json:\"credentials,omitempty\""; FilterAttributes *[]struct { Control *struct { Options *[]map[string]interface {} "json:\"options,omitempty\""; ReadableRegest *string "json:\"readable_regest,omitempty\""; Regex *string "json:\"regex,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"control,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]map[string]interface {} "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filter_attributes,omitempty\""; IsAgent *bool "json:\"is_agent,omitempty\""; IsWas *bool "json:\"is_was,omitempty\""; Name *string "json:\"name,omitempty\""; Plugins *map[string]interface {} "json:\"plugins,omitempty\""; Settings *struct { Advanced *map[string]interface {} "json:\"advanced,omitempty\""; Assessment *map[string]interface {} "json:\"assessment,omitempty\""; Basic *map[string]interface {} "json:\"basic,omitempty\""; Discovery *map[string]interface {} "json:\"discovery,omitempty\"" } "json:\"settings,omitempty\""; Title *string "json:\"title,omitempty\"" }
	if i, ok := r.(*struct {
		Compliance       *map[string]interface{} "json:\"compliance,omitempty\""
		Credentials      *map[string]interface{} "json:\"credentials,omitempty\""
		FilterAttributes *[]struct {
			Control *struct {
				Options        *[]map[string]interface{} "json:\"options,omitempty\""
				ReadableRegest *string                   "json:\"readable_regest,omitempty\""
				Regex          *string                   "json:\"regex,omitempty\""
				Type           *string                   "json:\"type,omitempty\""
			} "json:\"control,omitempty\""
			Name         *string                   "json:\"name,omitempty\""
			Operators    *[]map[string]interface{} "json:\"operators,omitempty\""
			ReadableName *string                   "json:\"readable_name,omitempty\""
		} "json:\"filter_attributes,omitempty\""
		IsAgent  *bool                   "json:\"is_agent,omitempty\""
		IsWas    *bool                   "json:\"is_was,omitempty\""
		Name     *string                 "json:\"name,omitempty\""
		Plugins  *map[string]interface{} "json:\"plugins,omitempty\""
		Settings *struct {
			Advanced   *map[string]interface{} "json:\"advanced,omitempty\""
			Assessment *map[string]interface{} "json:\"assessment,omitempty\""
			Basic      *map[string]interface{} "json:\"basic,omitempty\""
			Discovery  *map[string]interface{} "json:\"discovery,omitempty\""
		} "json:\"settings,omitempty\""
		Title *string "json:\"title,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// IoFilters implementation of the IoFilters interface
type IoFilters struct {
	*ClientWithResponses
}

// AgentsList calls the IoFilters ´s function
func (c *IoFilters) AgentsList(reqEditors ...RequestEditorFn) (*struct {
	Filters *[]struct {
		Control *struct {
			ReadableRegex *string "json:\"readable_regex,omitempty\""
			Regex         *string "json:\"regex,omitempty\""
			Type          *string "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		GroupName    *string   "json:\"group_name,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	} "json:\"filters,omitempty\""
	Sort *struct {
		MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
		SortableFields *[]string "json:\"sortable_fields,omitempty\""
	} "json:\"sort,omitempty\""
	WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoFiltersAgentsListWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Filters *[]struct { Control *struct { ReadableRegex *string "json:\"readable_regex,omitempty\""; Regex *string "json:\"regex,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"control,omitempty\""; GroupName *string "json:\"group_name,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filters,omitempty\""; Sort *struct { MaxSortFields *int "json:\"max_sort_fields,omitempty\""; SortableFields *[]string "json:\"sortable_fields,omitempty\"" } "json:\"sort,omitempty\""; WildcardFields *[]string "json:\"wildcard_fields,omitempty\"" }
	if i, ok := r.(*struct {
		Filters *[]struct {
			Control *struct {
				ReadableRegex *string "json:\"readable_regex,omitempty\""
				Regex         *string "json:\"regex,omitempty\""
				Type          *string "json:\"type,omitempty\""
			} "json:\"control,omitempty\""
			GroupName    *string   "json:\"group_name,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"filters,omitempty\""
		Sort *struct {
			MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
			SortableFields *[]string "json:\"sortable_fields,omitempty\""
		} "json:\"sort,omitempty\""
		WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AssetsListV2WithBody calls the IoFilters ´s function
func (c *IoFilters) AssetsListV2WithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	Filters *[]struct {
		Control *struct {
			ReadableRegex *string "json:\"readable_regex,omitempty\""
			Regex         *string "json:\"regex,omitempty\""
			Type          *string "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		GroupName    *string   "json:\"group_name,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	} "json:\"filters,omitempty\""
	Sort *struct {
		MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
		SortableFields *[]string "json:\"sortable_fields,omitempty\""
	} "json:\"sort,omitempty\""
	WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoFiltersAssetsListV2WithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Filters *[]struct { Control *struct { ReadableRegex *string "json:\"readable_regex,omitempty\""; Regex *string "json:\"regex,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"control,omitempty\""; GroupName *string "json:\"group_name,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filters,omitempty\""; Sort *struct { MaxSortFields *int "json:\"max_sort_fields,omitempty\""; SortableFields *[]string "json:\"sortable_fields,omitempty\"" } "json:\"sort,omitempty\""; WildcardFields *[]string "json:\"wildcard_fields,omitempty\"" }
	if i, ok := r.(*struct {
		Filters *[]struct {
			Control *struct {
				ReadableRegex *string "json:\"readable_regex,omitempty\""
				Regex         *string "json:\"regex,omitempty\""
				Type          *string "json:\"type,omitempty\""
			} "json:\"control,omitempty\""
			GroupName    *string   "json:\"group_name,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"filters,omitempty\""
		Sort *struct {
			MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
			SortableFields *[]string "json:\"sortable_fields,omitempty\""
		} "json:\"sort,omitempty\""
		WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AssetsListV2 calls the IoFilters ´s function
func (c *IoFilters) AssetsListV2(arg1 IoFiltersAssetsListV2JSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	Filters *[]struct {
		Control *struct {
			ReadableRegex *string "json:\"readable_regex,omitempty\""
			Regex         *string "json:\"regex,omitempty\""
			Type          *string "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		GroupName    *string   "json:\"group_name,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	} "json:\"filters,omitempty\""
	Sort *struct {
		MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
		SortableFields *[]string "json:\"sortable_fields,omitempty\""
	} "json:\"sort,omitempty\""
	WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoFiltersAssetsListV2WithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Filters *[]struct { Control *struct { ReadableRegex *string "json:\"readable_regex,omitempty\""; Regex *string "json:\"regex,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"control,omitempty\""; GroupName *string "json:\"group_name,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filters,omitempty\""; Sort *struct { MaxSortFields *int "json:\"max_sort_fields,omitempty\""; SortableFields *[]string "json:\"sortable_fields,omitempty\"" } "json:\"sort,omitempty\""; WildcardFields *[]string "json:\"wildcard_fields,omitempty\"" }
	if i, ok := r.(*struct {
		Filters *[]struct {
			Control *struct {
				ReadableRegex *string "json:\"readable_regex,omitempty\""
				Regex         *string "json:\"regex,omitempty\""
				Type          *string "json:\"type,omitempty\""
			} "json:\"control,omitempty\""
			GroupName    *string   "json:\"group_name,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"filters,omitempty\""
		Sort *struct {
			MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
			SortableFields *[]string "json:\"sortable_fields,omitempty\""
		} "json:\"sort,omitempty\""
		WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AssetsList calls the IoFilters ´s function
func (c *IoFilters) AssetsList(reqEditors ...RequestEditorFn) (*struct {
	Filters *[]struct {
		Control *struct {
			ReadableRegex *string "json:\"readable_regex,omitempty\""
			Regex         *string "json:\"regex,omitempty\""
			Type          *string "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		GroupName    *string   "json:\"group_name,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	} "json:\"filters,omitempty\""
	Sort *struct {
		MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
		SortableFields *[]string "json:\"sortable_fields,omitempty\""
	} "json:\"sort,omitempty\""
	WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoFiltersAssetsListWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Filters *[]struct { Control *struct { ReadableRegex *string "json:\"readable_regex,omitempty\""; Regex *string "json:\"regex,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"control,omitempty\""; GroupName *string "json:\"group_name,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filters,omitempty\""; Sort *struct { MaxSortFields *int "json:\"max_sort_fields,omitempty\""; SortableFields *[]string "json:\"sortable_fields,omitempty\"" } "json:\"sort,omitempty\""; WildcardFields *[]string "json:\"wildcard_fields,omitempty\"" }
	if i, ok := r.(*struct {
		Filters *[]struct {
			Control *struct {
				ReadableRegex *string "json:\"readable_regex,omitempty\""
				Regex         *string "json:\"regex,omitempty\""
				Type          *string "json:\"type,omitempty\""
			} "json:\"control,omitempty\""
			GroupName    *string   "json:\"group_name,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"filters,omitempty\""
		Sort *struct {
			MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
			SortableFields *[]string "json:\"sortable_fields,omitempty\""
		} "json:\"sort,omitempty\""
		WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CredentialsList calls the IoFilters ´s function
func (c *IoFilters) CredentialsList(reqEditors ...RequestEditorFn) (*struct {
	Filters *[]struct {
		Control *struct {
			ReadableRegex *string "json:\"readable_regex,omitempty\""
			Regex         *string "json:\"regex,omitempty\""
			Type          *string "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		GroupName    *string   "json:\"group_name,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	} "json:\"filters,omitempty\""
	Sort *struct {
		MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
		SortableFields *[]string "json:\"sortable_fields,omitempty\""
	} "json:\"sort,omitempty\""
	WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoFiltersCredentialsListWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Filters *[]struct { Control *struct { ReadableRegex *string "json:\"readable_regex,omitempty\""; Regex *string "json:\"regex,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"control,omitempty\""; GroupName *string "json:\"group_name,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filters,omitempty\""; Sort *struct { MaxSortFields *int "json:\"max_sort_fields,omitempty\""; SortableFields *[]string "json:\"sortable_fields,omitempty\"" } "json:\"sort,omitempty\""; WildcardFields *[]string "json:\"wildcard_fields,omitempty\"" }
	if i, ok := r.(*struct {
		Filters *[]struct {
			Control *struct {
				ReadableRegex *string "json:\"readable_regex,omitempty\""
				Regex         *string "json:\"regex,omitempty\""
				Type          *string "json:\"type,omitempty\""
			} "json:\"control,omitempty\""
			GroupName    *string   "json:\"group_name,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"filters,omitempty\""
		Sort *struct {
			MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
			SortableFields *[]string "json:\"sortable_fields,omitempty\""
		} "json:\"sort,omitempty\""
		WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ScanHistoryList calls the IoFilters ´s function
func (c *IoFilters) ScanHistoryList(reqEditors ...RequestEditorFn) (*struct {
	Filters *[]struct {
		Control *struct {
			ReadableRegex *string "json:\"readable_regex,omitempty\""
			Regex         *string "json:\"regex,omitempty\""
			Type          *string "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		GroupName    *string   "json:\"group_name,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	} "json:\"filters,omitempty\""
	Sort *struct {
		MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
		SortableFields *[]string "json:\"sortable_fields,omitempty\""
	} "json:\"sort,omitempty\""
	WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoFiltersScanHistoryListWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Filters *[]struct { Control *struct { ReadableRegex *string "json:\"readable_regex,omitempty\""; Regex *string "json:\"regex,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"control,omitempty\""; GroupName *string "json:\"group_name,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filters,omitempty\""; Sort *struct { MaxSortFields *int "json:\"max_sort_fields,omitempty\""; SortableFields *[]string "json:\"sortable_fields,omitempty\"" } "json:\"sort,omitempty\""; WildcardFields *[]string "json:\"wildcard_fields,omitempty\"" }
	if i, ok := r.(*struct {
		Filters *[]struct {
			Control *struct {
				ReadableRegex *string "json:\"readable_regex,omitempty\""
				Regex         *string "json:\"regex,omitempty\""
				Type          *string "json:\"type,omitempty\""
			} "json:\"control,omitempty\""
			GroupName    *string   "json:\"group_name,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"filters,omitempty\""
		Sort *struct {
			MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
			SortableFields *[]string "json:\"sortable_fields,omitempty\""
		} "json:\"sort,omitempty\""
		WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ScanList calls the IoFilters ´s function
func (c *IoFilters) ScanList(reqEditors ...RequestEditorFn) (*struct {
	Filters *[]struct {
		Control *struct {
			ReadableRegex *string "json:\"readable_regex,omitempty\""
			Regex         *string "json:\"regex,omitempty\""
			Type          *string "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		GroupName    *string   "json:\"group_name,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	} "json:\"filters,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoFiltersScanListWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Filters *[]struct { Control *struct { ReadableRegex *string "json:\"readable_regex,omitempty\""; Regex *string "json:\"regex,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"control,omitempty\""; GroupName *string "json:\"group_name,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filters,omitempty\"" }
	if i, ok := r.(*struct {
		Filters *[]struct {
			Control *struct {
				ReadableRegex *string "json:\"readable_regex,omitempty\""
				Regex         *string "json:\"regex,omitempty\""
				Type          *string "json:\"type,omitempty\""
			} "json:\"control,omitempty\""
			GroupName    *string   "json:\"group_name,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"filters,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// VulnerabilitiesWorkbenchListV2WithBody calls the IoFilters ´s function
func (c *IoFilters) VulnerabilitiesWorkbenchListV2WithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	Filters *[]struct {
		Control *struct {
			ReadableRegex *string "json:\"readable_regex,omitempty\""
			Regex         *string "json:\"regex,omitempty\""
			Type          *string "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		GroupName    *string   "json:\"group_name,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	} "json:\"filters,omitempty\""
	Sort *struct {
		MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
		SortableFields *[]string "json:\"sortable_fields,omitempty\""
	} "json:\"sort,omitempty\""
	WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoFiltersVulnerabilitiesWorkbenchListV2WithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Filters *[]struct { Control *struct { ReadableRegex *string "json:\"readable_regex,omitempty\""; Regex *string "json:\"regex,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"control,omitempty\""; GroupName *string "json:\"group_name,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filters,omitempty\""; Sort *struct { MaxSortFields *int "json:\"max_sort_fields,omitempty\""; SortableFields *[]string "json:\"sortable_fields,omitempty\"" } "json:\"sort,omitempty\""; WildcardFields *[]string "json:\"wildcard_fields,omitempty\"" }
	if i, ok := r.(*struct {
		Filters *[]struct {
			Control *struct {
				ReadableRegex *string "json:\"readable_regex,omitempty\""
				Regex         *string "json:\"regex,omitempty\""
				Type          *string "json:\"type,omitempty\""
			} "json:\"control,omitempty\""
			GroupName    *string   "json:\"group_name,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"filters,omitempty\""
		Sort *struct {
			MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
			SortableFields *[]string "json:\"sortable_fields,omitempty\""
		} "json:\"sort,omitempty\""
		WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// VulnerabilitiesWorkbenchListV2 calls the IoFilters ´s function
func (c *IoFilters) VulnerabilitiesWorkbenchListV2(arg1 IoFiltersVulnerabilitiesWorkbenchListV2JSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	Filters *[]struct {
		Control *struct {
			ReadableRegex *string "json:\"readable_regex,omitempty\""
			Regex         *string "json:\"regex,omitempty\""
			Type          *string "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		GroupName    *string   "json:\"group_name,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	} "json:\"filters,omitempty\""
	Sort *struct {
		MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
		SortableFields *[]string "json:\"sortable_fields,omitempty\""
	} "json:\"sort,omitempty\""
	WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoFiltersVulnerabilitiesWorkbenchListV2WithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Filters *[]struct { Control *struct { ReadableRegex *string "json:\"readable_regex,omitempty\""; Regex *string "json:\"regex,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"control,omitempty\""; GroupName *string "json:\"group_name,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filters,omitempty\""; Sort *struct { MaxSortFields *int "json:\"max_sort_fields,omitempty\""; SortableFields *[]string "json:\"sortable_fields,omitempty\"" } "json:\"sort,omitempty\""; WildcardFields *[]string "json:\"wildcard_fields,omitempty\"" }
	if i, ok := r.(*struct {
		Filters *[]struct {
			Control *struct {
				ReadableRegex *string "json:\"readable_regex,omitempty\""
				Regex         *string "json:\"regex,omitempty\""
				Type          *string "json:\"type,omitempty\""
			} "json:\"control,omitempty\""
			GroupName    *string   "json:\"group_name,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"filters,omitempty\""
		Sort *struct {
			MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
			SortableFields *[]string "json:\"sortable_fields,omitempty\""
		} "json:\"sort,omitempty\""
		WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// VulnerabilitiesWorkbenchList calls the IoFilters ´s function
func (c *IoFilters) VulnerabilitiesWorkbenchList(reqEditors ...RequestEditorFn) (*struct {
	Filters *[]struct {
		Control *struct {
			ReadableRegex *string "json:\"readable_regex,omitempty\""
			Regex         *string "json:\"regex,omitempty\""
			Type          *string "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		GroupName    *string   "json:\"group_name,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	} "json:\"filters,omitempty\""
	Sort *struct {
		MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
		SortableFields *[]string "json:\"sortable_fields,omitempty\""
	} "json:\"sort,omitempty\""
	WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoFiltersVulnerabilitiesWorkbenchListWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Filters *[]struct { Control *struct { ReadableRegex *string "json:\"readable_regex,omitempty\""; Regex *string "json:\"regex,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"control,omitempty\""; GroupName *string "json:\"group_name,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filters,omitempty\""; Sort *struct { MaxSortFields *int "json:\"max_sort_fields,omitempty\""; SortableFields *[]string "json:\"sortable_fields,omitempty\"" } "json:\"sort,omitempty\""; WildcardFields *[]string "json:\"wildcard_fields,omitempty\"" }
	if i, ok := r.(*struct {
		Filters *[]struct {
			Control *struct {
				ReadableRegex *string "json:\"readable_regex,omitempty\""
				Regex         *string "json:\"regex,omitempty\""
				Type          *string "json:\"type,omitempty\""
			} "json:\"control,omitempty\""
			GroupName    *string   "json:\"group_name,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"filters,omitempty\""
		Sort *struct {
			MaxSortFields  *int      "json:\"max_sort_fields,omitempty\""
			SortableFields *[]string "json:\"sortable_fields,omitempty\""
		} "json:\"sort,omitempty\""
		WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// IoPlugins implementation of the IoPlugins interface
type IoPlugins struct {
	*ClientWithResponses
}

// Details calls the IoPlugins ´s function
func (c *IoPlugins) Details(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
	Attributes *[]struct {
		AttributeName  *string "json:\"attribute_name,omitempty\""
		AttributeValue *string "json:\"attribute_value,omitempty\""
	} "json:\"attributes,omitempty\""
	FamilyName *string "json:\"family_name,omitempty\""
	Id         *int    "json:\"id,omitempty\""
	Name       *string "json:\"name,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoPluginsDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Attributes *[]struct { AttributeName *string "json:\"attribute_name,omitempty\""; AttributeValue *string "json:\"attribute_value,omitempty\"" } "json:\"attributes,omitempty\""; FamilyName *string "json:\"family_name,omitempty\""; Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\"" }
	if i, ok := r.(*struct {
		Attributes *[]struct {
			AttributeName  *string "json:\"attribute_name,omitempty\""
			AttributeValue *string "json:\"attribute_value,omitempty\""
		} "json:\"attributes,omitempty\""
		FamilyName *string "json:\"family_name,omitempty\""
		Id         *int    "json:\"id,omitempty\""
		Name       *string "json:\"name,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// FamiliesList calls the IoPlugins ´s function
func (c *IoPlugins) FamiliesList(params *IoPluginsFamiliesListParams, reqEditors ...RequestEditorFn) (*struct {
	Families *[]struct {
		Count *int    "json:\"count,omitempty\""
		Id    *int    "json:\"id,omitempty\""
		Name  *string "json:\"name,omitempty\""
	} "json:\"families,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoPluginsFamiliesListWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Families *[]struct { Count *int "json:\"count,omitempty\""; Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"families,omitempty\"" }
	if i, ok := r.(*struct {
		Families *[]struct {
			Count *int    "json:\"count,omitempty\""
			Id    *int    "json:\"id,omitempty\""
			Name  *string "json:\"name,omitempty\""
		} "json:\"families,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// FamilyDetails calls the IoPlugins ´s function
func (c *IoPlugins) FamilyDetails(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
	Id      *int    "json:\"id,omitempty\""
	Name    *string "json:\"name,omitempty\""
	Plugins *[]struct {
		Id   *int    "json:\"id,omitempty\""
		Name *string "json:\"name,omitempty\""
	} "json:\"plugins,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoPluginsFamilyDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Plugins *[]struct { Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"plugins,omitempty\"" }
	if i, ok := r.(*struct {
		Id      *int    "json:\"id,omitempty\""
		Name    *string "json:\"name,omitempty\""
		Plugins *[]struct {
			Id   *int    "json:\"id,omitempty\""
			Name *string "json:\"name,omitempty\""
		} "json:\"plugins,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// List calls the IoPlugins ´s function
func (c *IoPlugins) List(params *IoPluginsListParams, reqEditors ...RequestEditorFn) (*struct {
	PluginDetails *[]struct {
		Attributes *[]struct {
			Cpe                *[]string "json:\"cpe,omitempty\""
			Cve                *[]string "json:\"cve,omitempty\""
			Cvss3BaseScore     *string   "json:\"cvss3_base_score,omitempty\""
			Cvss3TemporalScore *float32  "json:\"cvss3_temporal_score,omitempty\""
			CvssBaseScore      *string   "json:\"cvss_base_score,omitempty\""
			CvssTemporalScore  *float32  "json:\"cvss_temporal_score,omitempty\""
			CvssVector         *struct {
				AccessComplexity      *string "json:\"AccessComplexity,omitempty\""
				AccessVector          *string "json:\"AccessVector,omitempty\""
				Authentication        *string "json:\"Authentication,omitempty\""
				AvailabilityImpact    *string "json:\"Availability-Impact,omitempty\""
				ConfidentialityImpact *string "json:\"Confidentiality-Impact,omitempty\""
				IntegrityImpact       *string "json:\"Integrity-Impact,omitempty\""
			} "json:\"cvss_vector,omitempty\""
			DefaultAccount             *string   "json:\"default_account,omitempty\""
			Description                *string   "json:\"description,omitempty\""
			ExploitAvailable           *bool     "json:\"exploit_available,omitempty\""
			ExploitFrameworkCanvas     *bool     "json:\"exploit_framework_canvas,omitempty\""
			ExploitFrameworkCore       *bool     "json:\"exploit_framework_core,omitempty\""
			ExploitFrameworkD2Elliot   *bool     "json:\"exploit_framework_d2_elliot,omitempty\""
			ExploitFrameworkExploithub *bool     "json:\"exploit_framework_exploithub,omitempty\""
			ExploitFrameworkMetasploit *bool     "json:\"exploit_framework_metasploit,omitempty\""
			ExploitedByMalware         *bool     "json:\"exploited_by_malware,omitempty\""
			ExploitedByNessus          *bool     "json:\"exploited_by_nessus,omitempty\""
			HasPatch                   *bool     "json:\"has_patch,omitempty\""
			InTheNews                  *bool     "json:\"in_the_news,omitempty\""
			Malware                    *bool     "json:\"malware,omitempty\""
			PatchPublicationDate       *string   "json:\"patch_publication_date,omitempty\""
			PluginModificationDate     *string   "json:\"plugin_modification_date,omitempty\""
			PluginPublicationDate      *string   "json:\"plugin_publication_date,omitempty\""
			PluginType                 *string   "json:\"plugin_type,omitempty\""
			PluginVersion              *string   "json:\"plugin_version,omitempty\""
			RiskFactor                 *string   "json:\"risk_factor,omitempty\""
			SeeAlso                    *[]string "json:\"see_also,omitempty\""
			Solution                   *string   "json:\"solution,omitempty\""
			Synopsis                   *string   "json:\"synopsis,omitempty\""
			UnsupportedByVendor        *bool     "json:\"unsupported_by_vendor,omitempty\""
			Vpr                        *struct {
				Drivers *map[string]interface{} "json:\"drivers,omitempty\""
				Score   *float32                "json:\"score,omitempty\""
				Updated *string                 "json:\"updated,omitempty\""
			} "json:\"vpr,omitempty\""
			Xref  *[]string "json:\"xref,omitempty\""
			Xrefs *[]string "json:\"xrefs,omitempty\""
		} "json:\"attributes,omitempty\""
		Id   *int    "json:\"id,omitempty\""
		Name *string "json:\"name,omitempty\""
	} "json:\"plugin_details,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoPluginsListWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { PluginDetails *[]struct { Attributes *[]struct { Cpe *[]string "json:\"cpe,omitempty\""; Cve *[]string "json:\"cve,omitempty\""; Cvss3BaseScore *string "json:\"cvss3_base_score,omitempty\""; Cvss3TemporalScore *float32 "json:\"cvss3_temporal_score,omitempty\""; CvssBaseScore *string "json:\"cvss_base_score,omitempty\""; CvssTemporalScore *float32 "json:\"cvss_temporal_score,omitempty\""; CvssVector *struct { AccessComplexity *string "json:\"AccessComplexity,omitempty\""; AccessVector *string "json:\"AccessVector,omitempty\""; Authentication *string "json:\"Authentication,omitempty\""; AvailabilityImpact *string "json:\"Availability-Impact,omitempty\""; ConfidentialityImpact *string "json:\"Confidentiality-Impact,omitempty\""; IntegrityImpact *string "json:\"Integrity-Impact,omitempty\"" } "json:\"cvss_vector,omitempty\""; DefaultAccount *string "json:\"default_account,omitempty\""; Description *string "json:\"description,omitempty\""; ExploitAvailable *bool "json:\"exploit_available,omitempty\""; ExploitFrameworkCanvas *bool "json:\"exploit_framework_canvas,omitempty\""; ExploitFrameworkCore *bool "json:\"exploit_framework_core,omitempty\""; ExploitFrameworkD2Elliot *bool "json:\"exploit_framework_d2_elliot,omitempty\""; ExploitFrameworkExploithub *bool "json:\"exploit_framework_exploithub,omitempty\""; ExploitFrameworkMetasploit *bool "json:\"exploit_framework_metasploit,omitempty\""; ExploitedByMalware *bool "json:\"exploited_by_malware,omitempty\""; ExploitedByNessus *bool "json:\"exploited_by_nessus,omitempty\""; HasPatch *bool "json:\"has_patch,omitempty\""; InTheNews *bool "json:\"in_the_news,omitempty\""; Malware *bool "json:\"malware,omitempty\""; PatchPublicationDate *string "json:\"patch_publication_date,omitempty\""; PluginModificationDate *string "json:\"plugin_modification_date,omitempty\""; PluginPublicationDate *string "json:\"plugin_publication_date,omitempty\""; PluginType *string "json:\"plugin_type,omitempty\""; PluginVersion *string "json:\"plugin_version,omitempty\""; RiskFactor *string "json:\"risk_factor,omitempty\""; SeeAlso *[]string "json:\"see_also,omitempty\""; Solution *string "json:\"solution,omitempty\""; Synopsis *string "json:\"synopsis,omitempty\""; UnsupportedByVendor *bool "json:\"unsupported_by_vendor,omitempty\""; Vpr *struct { Drivers *map[string]interface {} "json:\"drivers,omitempty\""; Score *float32 "json:\"score,omitempty\""; Updated *string "json:\"updated,omitempty\"" } "json:\"vpr,omitempty\""; Xref *[]string "json:\"xref,omitempty\""; Xrefs *[]string "json:\"xrefs,omitempty\"" } "json:\"attributes,omitempty\""; Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"plugin_details,omitempty\"" }
	if i, ok := r.(*struct {
		PluginDetails *[]struct {
			Attributes *[]struct {
				Cpe                *[]string "json:\"cpe,omitempty\""
				Cve                *[]string "json:\"cve,omitempty\""
				Cvss3BaseScore     *string   "json:\"cvss3_base_score,omitempty\""
				Cvss3TemporalScore *float32  "json:\"cvss3_temporal_score,omitempty\""
				CvssBaseScore      *string   "json:\"cvss_base_score,omitempty\""
				CvssTemporalScore  *float32  "json:\"cvss_temporal_score,omitempty\""
				CvssVector         *struct {
					AccessComplexity      *string "json:\"AccessComplexity,omitempty\""
					AccessVector          *string "json:\"AccessVector,omitempty\""
					Authentication        *string "json:\"Authentication,omitempty\""
					AvailabilityImpact    *string "json:\"Availability-Impact,omitempty\""
					ConfidentialityImpact *string "json:\"Confidentiality-Impact,omitempty\""
					IntegrityImpact       *string "json:\"Integrity-Impact,omitempty\""
				} "json:\"cvss_vector,omitempty\""
				DefaultAccount             *string   "json:\"default_account,omitempty\""
				Description                *string   "json:\"description,omitempty\""
				ExploitAvailable           *bool     "json:\"exploit_available,omitempty\""
				ExploitFrameworkCanvas     *bool     "json:\"exploit_framework_canvas,omitempty\""
				ExploitFrameworkCore       *bool     "json:\"exploit_framework_core,omitempty\""
				ExploitFrameworkD2Elliot   *bool     "json:\"exploit_framework_d2_elliot,omitempty\""
				ExploitFrameworkExploithub *bool     "json:\"exploit_framework_exploithub,omitempty\""
				ExploitFrameworkMetasploit *bool     "json:\"exploit_framework_metasploit,omitempty\""
				ExploitedByMalware         *bool     "json:\"exploited_by_malware,omitempty\""
				ExploitedByNessus          *bool     "json:\"exploited_by_nessus,omitempty\""
				HasPatch                   *bool     "json:\"has_patch,omitempty\""
				InTheNews                  *bool     "json:\"in_the_news,omitempty\""
				Malware                    *bool     "json:\"malware,omitempty\""
				PatchPublicationDate       *string   "json:\"patch_publication_date,omitempty\""
				PluginModificationDate     *string   "json:\"plugin_modification_date,omitempty\""
				PluginPublicationDate      *string   "json:\"plugin_publication_date,omitempty\""
				PluginType                 *string   "json:\"plugin_type,omitempty\""
				PluginVersion              *string   "json:\"plugin_version,omitempty\""
				RiskFactor                 *string   "json:\"risk_factor,omitempty\""
				SeeAlso                    *[]string "json:\"see_also,omitempty\""
				Solution                   *string   "json:\"solution,omitempty\""
				Synopsis                   *string   "json:\"synopsis,omitempty\""
				UnsupportedByVendor        *bool     "json:\"unsupported_by_vendor,omitempty\""
				Vpr                        *struct {
					Drivers *map[string]interface{} "json:\"drivers,omitempty\""
					Score   *float32                "json:\"score,omitempty\""
					Updated *string                 "json:\"updated,omitempty\""
				} "json:\"vpr,omitempty\""
				Xref  *[]string "json:\"xref,omitempty\""
				Xrefs *[]string "json:\"xrefs,omitempty\""
			} "json:\"attributes,omitempty\""
			Id   *int    "json:\"id,omitempty\""
			Name *string "json:\"name,omitempty\""
		} "json:\"plugin_details,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// IoV2 implementation of the IoV2 interface
type IoV2 struct {
	*ClientWithResponses
}

// AccessGroupsCreateWithBody calls the IoV2 ´s function
func (c *IoV2) AccessGroupsCreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	AccessGroupType *N200AccessGroupType "json:\"access_group_type,omitempty\""
	AllAssets       *bool                "json:\"all_assets,omitempty\""
	ContainerUuid   *string              "json:\"container_uuid,omitempty\""
	CreatedAt       *string              "json:\"created_at,omitempty\""
	CreatedByName   *string              "json:\"created_by_name,omitempty\""
	CreatedByUuid   *string              "json:\"created_by_uuid,omitempty\""
	Id              *string              "json:\"id,omitempty\""
	Name            *string              "json:\"name,omitempty\""
	Principals      *[]struct {
		Permissions   *N200PrincipalsPermissions "json:\"permissions,omitempty\""
		PrincipalId   *string                    "json:\"principal_id,omitempty\""
		PrincipalName *string                    "json:\"principal_name,omitempty\""
		Type          *N200PrincipalsType        "json:\"type,omitempty\""
	} "json:\"principals,omitempty\""
	ProcessingPercentComplete *int    "json:\"processing_percent_complete,omitempty\""
	Status                    *string "json:\"status,omitempty\""
	UpdatedAt                 *string "json:\"updated_at,omitempty\""
	UpdatedByName             *string "json:\"updated_by_name,omitempty\""
	UpdatedByUuid             *string "json:\"updated_by_uuid,omitempty\""
	Version                   *int    "json:\"version,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV2AccessGroupsCreateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AccessGroupType *N200AccessGroupType "json:\"access_group_type,omitempty\""; AllAssets *bool "json:\"all_assets,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedByName *string "json:\"created_by_name,omitempty\""; CreatedByUuid *string "json:\"created_by_uuid,omitempty\""; Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Principals *[]struct { Permissions *N200PrincipalsPermissions "json:\"permissions,omitempty\""; PrincipalId *string "json:\"principal_id,omitempty\""; PrincipalName *string "json:\"principal_name,omitempty\""; Type *N200PrincipalsType "json:\"type,omitempty\"" } "json:\"principals,omitempty\""; ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""; Status *string "json:\"status,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedByName *string "json:\"updated_by_name,omitempty\""; UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""; Version *int "json:\"version,omitempty\"" }
	if i, ok := r.(*struct {
		AccessGroupType *N200AccessGroupType "json:\"access_group_type,omitempty\""
		AllAssets       *bool                "json:\"all_assets,omitempty\""
		ContainerUuid   *string              "json:\"container_uuid,omitempty\""
		CreatedAt       *string              "json:\"created_at,omitempty\""
		CreatedByName   *string              "json:\"created_by_name,omitempty\""
		CreatedByUuid   *string              "json:\"created_by_uuid,omitempty\""
		Id              *string              "json:\"id,omitempty\""
		Name            *string              "json:\"name,omitempty\""
		Principals      *[]struct {
			Permissions   *N200PrincipalsPermissions "json:\"permissions,omitempty\""
			PrincipalId   *string                    "json:\"principal_id,omitempty\""
			PrincipalName *string                    "json:\"principal_name,omitempty\""
			Type          *N200PrincipalsType        "json:\"type,omitempty\""
		} "json:\"principals,omitempty\""
		ProcessingPercentComplete *int    "json:\"processing_percent_complete,omitempty\""
		Status                    *string "json:\"status,omitempty\""
		UpdatedAt                 *string "json:\"updated_at,omitempty\""
		UpdatedByName             *string "json:\"updated_by_name,omitempty\""
		UpdatedByUuid             *string "json:\"updated_by_uuid,omitempty\""
		Version                   *int    "json:\"version,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsCreate calls the IoV2 ´s function
func (c *IoV2) AccessGroupsCreate(arg1 IoV2AccessGroupsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	AccessGroupType *N200AccessGroupType "json:\"access_group_type,omitempty\""
	AllAssets       *bool                "json:\"all_assets,omitempty\""
	ContainerUuid   *string              "json:\"container_uuid,omitempty\""
	CreatedAt       *string              "json:\"created_at,omitempty\""
	CreatedByName   *string              "json:\"created_by_name,omitempty\""
	CreatedByUuid   *string              "json:\"created_by_uuid,omitempty\""
	Id              *string              "json:\"id,omitempty\""
	Name            *string              "json:\"name,omitempty\""
	Principals      *[]struct {
		Permissions   *N200PrincipalsPermissions "json:\"permissions,omitempty\""
		PrincipalId   *string                    "json:\"principal_id,omitempty\""
		PrincipalName *string                    "json:\"principal_name,omitempty\""
		Type          *N200PrincipalsType        "json:\"type,omitempty\""
	} "json:\"principals,omitempty\""
	ProcessingPercentComplete *int    "json:\"processing_percent_complete,omitempty\""
	Status                    *string "json:\"status,omitempty\""
	UpdatedAt                 *string "json:\"updated_at,omitempty\""
	UpdatedByName             *string "json:\"updated_by_name,omitempty\""
	UpdatedByUuid             *string "json:\"updated_by_uuid,omitempty\""
	Version                   *int    "json:\"version,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV2AccessGroupsCreateWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AccessGroupType *N200AccessGroupType "json:\"access_group_type,omitempty\""; AllAssets *bool "json:\"all_assets,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedByName *string "json:\"created_by_name,omitempty\""; CreatedByUuid *string "json:\"created_by_uuid,omitempty\""; Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Principals *[]struct { Permissions *N200PrincipalsPermissions "json:\"permissions,omitempty\""; PrincipalId *string "json:\"principal_id,omitempty\""; PrincipalName *string "json:\"principal_name,omitempty\""; Type *N200PrincipalsType "json:\"type,omitempty\"" } "json:\"principals,omitempty\""; ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""; Status *string "json:\"status,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedByName *string "json:\"updated_by_name,omitempty\""; UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""; Version *int "json:\"version,omitempty\"" }
	if i, ok := r.(*struct {
		AccessGroupType *N200AccessGroupType "json:\"access_group_type,omitempty\""
		AllAssets       *bool                "json:\"all_assets,omitempty\""
		ContainerUuid   *string              "json:\"container_uuid,omitempty\""
		CreatedAt       *string              "json:\"created_at,omitempty\""
		CreatedByName   *string              "json:\"created_by_name,omitempty\""
		CreatedByUuid   *string              "json:\"created_by_uuid,omitempty\""
		Id              *string              "json:\"id,omitempty\""
		Name            *string              "json:\"name,omitempty\""
		Principals      *[]struct {
			Permissions   *N200PrincipalsPermissions "json:\"permissions,omitempty\""
			PrincipalId   *string                    "json:\"principal_id,omitempty\""
			PrincipalName *string                    "json:\"principal_name,omitempty\""
			Type          *N200PrincipalsType        "json:\"type,omitempty\""
		} "json:\"principals,omitempty\""
		ProcessingPercentComplete *int    "json:\"processing_percent_complete,omitempty\""
		Status                    *string "json:\"status,omitempty\""
		UpdatedAt                 *string "json:\"updated_at,omitempty\""
		UpdatedByName             *string "json:\"updated_by_name,omitempty\""
		UpdatedByUuid             *string "json:\"updated_by_uuid,omitempty\""
		Version                   *int    "json:\"version,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsDelete calls the IoV2 ´s function
func (c *IoV2) AccessGroupsDelete(contentType string, reqEditors ...RequestEditorFn) (*IoV2AccessGroupsDeleteResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV2AccessGroupsDeleteWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *IoV2AccessGroupsDeleteResponse
	if i, ok := r.(*IoV2AccessGroupsDeleteResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsDetails calls the IoV2 ´s function
func (c *IoV2) AccessGroupsDetails(contentType string, reqEditors ...RequestEditorFn) (*struct {
	AccessGroupType *N200AccessGroupType "json:\"access_group_type,omitempty\""
	AllAssets       *bool                "json:\"all_assets,omitempty\""
	ContainerUuid   *string              "json:\"container_uuid,omitempty\""
	CreatedAt       *string              "json:\"created_at,omitempty\""
	CreatedByName   *string              "json:\"created_by_name,omitempty\""
	CreatedByUuid   *string              "json:\"created_by_uuid,omitempty\""
	Id              *string              "json:\"id,omitempty\""
	Name            *string              "json:\"name,omitempty\""
	Principals      *[]struct {
		Permissions   *N200PrincipalsPermissions "json:\"permissions,omitempty\""
		PrincipalId   *string                    "json:\"principal_id,omitempty\""
		PrincipalName *string                    "json:\"principal_name,omitempty\""
		Type          *N200PrincipalsType        "json:\"type,omitempty\""
	} "json:\"principals,omitempty\""
	ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""
	Rules                     *[]struct {
		Operator *string   "json:\"operator,omitempty\""
		Terms    *[]string "json:\"terms,omitempty\""
		Type     *string   "json:\"type,omitempty\""
	} "json:\"rules,omitempty\""
	Status        *string "json:\"status,omitempty\""
	UpdatedAt     *string "json:\"updated_at,omitempty\""
	UpdatedByName *string "json:\"updated_by_name,omitempty\""
	UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""
	Version       *int    "json:\"version,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV2AccessGroupsDetailsWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AccessGroupType *N200AccessGroupType "json:\"access_group_type,omitempty\""; AllAssets *bool "json:\"all_assets,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedByName *string "json:\"created_by_name,omitempty\""; CreatedByUuid *string "json:\"created_by_uuid,omitempty\""; Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Principals *[]struct { Permissions *N200PrincipalsPermissions "json:\"permissions,omitempty\""; PrincipalId *string "json:\"principal_id,omitempty\""; PrincipalName *string "json:\"principal_name,omitempty\""; Type *N200PrincipalsType "json:\"type,omitempty\"" } "json:\"principals,omitempty\""; ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""; Rules *[]struct { Operator *string "json:\"operator,omitempty\""; Terms *[]string "json:\"terms,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"rules,omitempty\""; Status *string "json:\"status,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedByName *string "json:\"updated_by_name,omitempty\""; UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""; Version *int "json:\"version,omitempty\"" }
	if i, ok := r.(*struct {
		AccessGroupType *N200AccessGroupType "json:\"access_group_type,omitempty\""
		AllAssets       *bool                "json:\"all_assets,omitempty\""
		ContainerUuid   *string              "json:\"container_uuid,omitempty\""
		CreatedAt       *string              "json:\"created_at,omitempty\""
		CreatedByName   *string              "json:\"created_by_name,omitempty\""
		CreatedByUuid   *string              "json:\"created_by_uuid,omitempty\""
		Id              *string              "json:\"id,omitempty\""
		Name            *string              "json:\"name,omitempty\""
		Principals      *[]struct {
			Permissions   *N200PrincipalsPermissions "json:\"permissions,omitempty\""
			PrincipalId   *string                    "json:\"principal_id,omitempty\""
			PrincipalName *string                    "json:\"principal_name,omitempty\""
			Type          *N200PrincipalsType        "json:\"type,omitempty\""
		} "json:\"principals,omitempty\""
		ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""
		Rules                     *[]struct {
			Operator *string   "json:\"operator,omitempty\""
			Terms    *[]string "json:\"terms,omitempty\""
			Type     *string   "json:\"type,omitempty\""
		} "json:\"rules,omitempty\""
		Status        *string "json:\"status,omitempty\""
		UpdatedAt     *string "json:\"updated_at,omitempty\""
		UpdatedByName *string "json:\"updated_by_name,omitempty\""
		UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""
		Version       *int    "json:\"version,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsEditWithBody calls the IoV2 ´s function
func (c *IoV2) AccessGroupsEditWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	AccessGroupType *N200AccessGroupType "json:\"access_group_type,omitempty\""
	AllAssets       *bool                "json:\"all_assets,omitempty\""
	ContainerUuid   *string              "json:\"container_uuid,omitempty\""
	CreatedAt       *string              "json:\"created_at,omitempty\""
	CreatedByName   *string              "json:\"created_by_name,omitempty\""
	CreatedByUuid   *string              "json:\"created_by_uuid,omitempty\""
	Id              *string              "json:\"id,omitempty\""
	Name            *string              "json:\"name,omitempty\""
	Principals      *[]struct {
		Permissions   *N200PrincipalsPermissions "json:\"permissions,omitempty\""
		PrincipalId   *string                    "json:\"principal_id,omitempty\""
		PrincipalName *string                    "json:\"principal_name,omitempty\""
		Type          *N200PrincipalsType        "json:\"type,omitempty\""
	} "json:\"principals,omitempty\""
	ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""
	Rules                     *[]struct {
		Operator *string   "json:\"operator,omitempty\""
		Terms    *[]string "json:\"terms,omitempty\""
		Type     *string   "json:\"type,omitempty\""
	} "json:\"rules,omitempty\""
	Status        *string "json:\"status,omitempty\""
	UpdatedAt     *string "json:\"updated_at,omitempty\""
	UpdatedByName *string "json:\"updated_by_name,omitempty\""
	UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""
	Version       *int    "json:\"version,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV2AccessGroupsEditWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AccessGroupType *N200AccessGroupType "json:\"access_group_type,omitempty\""; AllAssets *bool "json:\"all_assets,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedByName *string "json:\"created_by_name,omitempty\""; CreatedByUuid *string "json:\"created_by_uuid,omitempty\""; Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Principals *[]struct { Permissions *N200PrincipalsPermissions "json:\"permissions,omitempty\""; PrincipalId *string "json:\"principal_id,omitempty\""; PrincipalName *string "json:\"principal_name,omitempty\""; Type *N200PrincipalsType "json:\"type,omitempty\"" } "json:\"principals,omitempty\""; ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""; Rules *[]struct { Operator *string "json:\"operator,omitempty\""; Terms *[]string "json:\"terms,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"rules,omitempty\""; Status *string "json:\"status,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedByName *string "json:\"updated_by_name,omitempty\""; UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""; Version *int "json:\"version,omitempty\"" }
	if i, ok := r.(*struct {
		AccessGroupType *N200AccessGroupType "json:\"access_group_type,omitempty\""
		AllAssets       *bool                "json:\"all_assets,omitempty\""
		ContainerUuid   *string              "json:\"container_uuid,omitempty\""
		CreatedAt       *string              "json:\"created_at,omitempty\""
		CreatedByName   *string              "json:\"created_by_name,omitempty\""
		CreatedByUuid   *string              "json:\"created_by_uuid,omitempty\""
		Id              *string              "json:\"id,omitempty\""
		Name            *string              "json:\"name,omitempty\""
		Principals      *[]struct {
			Permissions   *N200PrincipalsPermissions "json:\"permissions,omitempty\""
			PrincipalId   *string                    "json:\"principal_id,omitempty\""
			PrincipalName *string                    "json:\"principal_name,omitempty\""
			Type          *N200PrincipalsType        "json:\"type,omitempty\""
		} "json:\"principals,omitempty\""
		ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""
		Rules                     *[]struct {
			Operator *string   "json:\"operator,omitempty\""
			Terms    *[]string "json:\"terms,omitempty\""
			Type     *string   "json:\"type,omitempty\""
		} "json:\"rules,omitempty\""
		Status        *string "json:\"status,omitempty\""
		UpdatedAt     *string "json:\"updated_at,omitempty\""
		UpdatedByName *string "json:\"updated_by_name,omitempty\""
		UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""
		Version       *int    "json:\"version,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsEdit calls the IoV2 ´s function
func (c *IoV2) AccessGroupsEdit(arg1 string, arg2 IoV2AccessGroupsEditJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	AccessGroupType *N200AccessGroupType "json:\"access_group_type,omitempty\""
	AllAssets       *bool                "json:\"all_assets,omitempty\""
	ContainerUuid   *string              "json:\"container_uuid,omitempty\""
	CreatedAt       *string              "json:\"created_at,omitempty\""
	CreatedByName   *string              "json:\"created_by_name,omitempty\""
	CreatedByUuid   *string              "json:\"created_by_uuid,omitempty\""
	Id              *string              "json:\"id,omitempty\""
	Name            *string              "json:\"name,omitempty\""
	Principals      *[]struct {
		Permissions   *N200PrincipalsPermissions "json:\"permissions,omitempty\""
		PrincipalId   *string                    "json:\"principal_id,omitempty\""
		PrincipalName *string                    "json:\"principal_name,omitempty\""
		Type          *N200PrincipalsType        "json:\"type,omitempty\""
	} "json:\"principals,omitempty\""
	ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""
	Rules                     *[]struct {
		Operator *string   "json:\"operator,omitempty\""
		Terms    *[]string "json:\"terms,omitempty\""
		Type     *string   "json:\"type,omitempty\""
	} "json:\"rules,omitempty\""
	Status        *string "json:\"status,omitempty\""
	UpdatedAt     *string "json:\"updated_at,omitempty\""
	UpdatedByName *string "json:\"updated_by_name,omitempty\""
	UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""
	Version       *int    "json:\"version,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV2AccessGroupsEditWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AccessGroupType *N200AccessGroupType "json:\"access_group_type,omitempty\""; AllAssets *bool "json:\"all_assets,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedByName *string "json:\"created_by_name,omitempty\""; CreatedByUuid *string "json:\"created_by_uuid,omitempty\""; Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Principals *[]struct { Permissions *N200PrincipalsPermissions "json:\"permissions,omitempty\""; PrincipalId *string "json:\"principal_id,omitempty\""; PrincipalName *string "json:\"principal_name,omitempty\""; Type *N200PrincipalsType "json:\"type,omitempty\"" } "json:\"principals,omitempty\""; ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""; Rules *[]struct { Operator *string "json:\"operator,omitempty\""; Terms *[]string "json:\"terms,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"rules,omitempty\""; Status *string "json:\"status,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedByName *string "json:\"updated_by_name,omitempty\""; UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""; Version *int "json:\"version,omitempty\"" }
	if i, ok := r.(*struct {
		AccessGroupType *N200AccessGroupType "json:\"access_group_type,omitempty\""
		AllAssets       *bool                "json:\"all_assets,omitempty\""
		ContainerUuid   *string              "json:\"container_uuid,omitempty\""
		CreatedAt       *string              "json:\"created_at,omitempty\""
		CreatedByName   *string              "json:\"created_by_name,omitempty\""
		CreatedByUuid   *string              "json:\"created_by_uuid,omitempty\""
		Id              *string              "json:\"id,omitempty\""
		Name            *string              "json:\"name,omitempty\""
		Principals      *[]struct {
			Permissions   *N200PrincipalsPermissions "json:\"permissions,omitempty\""
			PrincipalId   *string                    "json:\"principal_id,omitempty\""
			PrincipalName *string                    "json:\"principal_name,omitempty\""
			Type          *N200PrincipalsType        "json:\"type,omitempty\""
		} "json:\"principals,omitempty\""
		ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""
		Rules                     *[]struct {
			Operator *string   "json:\"operator,omitempty\""
			Terms    *[]string "json:\"terms,omitempty\""
			Type     *string   "json:\"type,omitempty\""
		} "json:\"rules,omitempty\""
		Status        *string "json:\"status,omitempty\""
		UpdatedAt     *string "json:\"updated_at,omitempty\""
		UpdatedByName *string "json:\"updated_by_name,omitempty\""
		UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""
		Version       *int    "json:\"version,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsListFilters calls the IoV2 ´s function
func (c *IoV2) AccessGroupsListFilters(reqEditors ...RequestEditorFn) (*struct {
	Filters *[]struct {
		Control      *string   "json:\"control,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	} "json:\"filters,omitempty\""
	Sort *[]struct {
		SortableFields *string "json:\"sortable_fields,omitempty\""
	} "json:\"sort,omitempty\""
	WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV2AccessGroupsListFiltersWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Filters *[]struct { Control *string "json:\"control,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filters,omitempty\""; Sort *[]struct { SortableFields *string "json:\"sortable_fields,omitempty\"" } "json:\"sort,omitempty\""; WildcardFields *[]string "json:\"wildcard_fields,omitempty\"" }
	if i, ok := r.(*struct {
		Filters *[]struct {
			Control      *string   "json:\"control,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"filters,omitempty\""
		Sort *[]struct {
			SortableFields *string "json:\"sortable_fields,omitempty\""
		} "json:\"sort,omitempty\""
		WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsListRuleFilters calls the IoV2 ´s function
func (c *IoV2) AccessGroupsListRuleFilters(reqEditors ...RequestEditorFn) (*struct {
	Rules *[]struct {
		Control *[]struct {
			Regex *string "json:\"regex,omitempty\""
			Type  *string "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		Placeholder  *string   "json:\"placeholder,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	} "json:\"rules,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV2AccessGroupsListRuleFiltersWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Rules *[]struct { Control *[]struct { Regex *string "json:\"regex,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"control,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; Placeholder *string "json:\"placeholder,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"rules,omitempty\"" }
	if i, ok := r.(*struct {
		Rules *[]struct {
			Control *[]struct {
				Regex *string "json:\"regex,omitempty\""
				Type  *string "json:\"type,omitempty\""
			} "json:\"control,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			Placeholder  *string   "json:\"placeholder,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"rules,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsList calls the IoV2 ´s function
func (c *IoV2) AccessGroupsList(params *IoV2AccessGroupsListParams, reqEditors ...RequestEditorFn) (*struct {
	AccessGroups *struct {
		AccessGroupType *N200AccessGroupsAccessGroupType "json:\"access_group_type,omitempty\""
		AllAssets       *bool                            "json:\"all_assets,omitempty\""
		ContainerUuid   *string                          "json:\"container_uuid,omitempty\""
		CreatedAt       *string                          "json:\"created_at,omitempty\""
		CreatedByName   *string                          "json:\"created_by_name,omitempty\""
		CreatedByUuid   *string                          "json:\"created_by_uuid,omitempty\""
		Id              *string                          "json:\"id,omitempty\""
		Name            *string                          "json:\"name,omitempty\""
		Principals      *[]struct {
			Permissions   *N200AccessGroupsPrincipalsPermissions "json:\"permissions,omitempty\""
			PrincipalId   *string                                "json:\"principal_id,omitempty\""
			PrincipalName *string                                "json:\"principal_name,omitempty\""
			Type          *N200AccessGroupsPrincipalsType        "json:\"type,omitempty\""
		} "json:\"principals,omitempty\""
		ProcessingPercentComplete *int    "json:\"processing_percent_complete,omitempty\""
		Status                    *string "json:\"status,omitempty\""
		UpdatedAt                 *string "json:\"updated_at,omitempty\""
		UpdatedByName             *string "json:\"updated_by_name,omitempty\""
		UpdatedByUuid             *string "json:\"updated_by_uuid,omitempty\""
		Version                   *int    "json:\"version,omitempty\""
	} "json:\"access_groups,omitempty\""
	Pagination *struct {
		Limit  *int32 "json:\"limit,omitempty\""
		Offset *int32 "json:\"offset,omitempty\""
		Sort   *[]struct {
			Name  *string "json:\"name,omitempty\""
			Order *string "json:\"order,omitempty\""
		} "json:\"sort,omitempty\""
		Total *int32 "json:\"total,omitempty\""
	} "json:\"pagination,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV2AccessGroupsListWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AccessGroups *struct { AccessGroupType *N200AccessGroupsAccessGroupType "json:\"access_group_type,omitempty\""; AllAssets *bool "json:\"all_assets,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedByName *string "json:\"created_by_name,omitempty\""; CreatedByUuid *string "json:\"created_by_uuid,omitempty\""; Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Principals *[]struct { Permissions *N200AccessGroupsPrincipalsPermissions "json:\"permissions,omitempty\""; PrincipalId *string "json:\"principal_id,omitempty\""; PrincipalName *string "json:\"principal_name,omitempty\""; Type *N200AccessGroupsPrincipalsType "json:\"type,omitempty\"" } "json:\"principals,omitempty\""; ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""; Status *string "json:\"status,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedByName *string "json:\"updated_by_name,omitempty\""; UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""; Version *int "json:\"version,omitempty\"" } "json:\"access_groups,omitempty\""; Pagination *struct { Limit *int32 "json:\"limit,omitempty\""; Offset *int32 "json:\"offset,omitempty\""; Sort *[]struct { Name *string "json:\"name,omitempty\""; Order *string "json:\"order,omitempty\"" } "json:\"sort,omitempty\""; Total *int32 "json:\"total,omitempty\"" } "json:\"pagination,omitempty\"" }
	if i, ok := r.(*struct {
		AccessGroups *struct {
			AccessGroupType *N200AccessGroupsAccessGroupType "json:\"access_group_type,omitempty\""
			AllAssets       *bool                            "json:\"all_assets,omitempty\""
			ContainerUuid   *string                          "json:\"container_uuid,omitempty\""
			CreatedAt       *string                          "json:\"created_at,omitempty\""
			CreatedByName   *string                          "json:\"created_by_name,omitempty\""
			CreatedByUuid   *string                          "json:\"created_by_uuid,omitempty\""
			Id              *string                          "json:\"id,omitempty\""
			Name            *string                          "json:\"name,omitempty\""
			Principals      *[]struct {
				Permissions   *N200AccessGroupsPrincipalsPermissions "json:\"permissions,omitempty\""
				PrincipalId   *string                                "json:\"principal_id,omitempty\""
				PrincipalName *string                                "json:\"principal_name,omitempty\""
				Type          *N200AccessGroupsPrincipalsType        "json:\"type,omitempty\""
			} "json:\"principals,omitempty\""
			ProcessingPercentComplete *int    "json:\"processing_percent_complete,omitempty\""
			Status                    *string "json:\"status,omitempty\""
			UpdatedAt                 *string "json:\"updated_at,omitempty\""
			UpdatedByName             *string "json:\"updated_by_name,omitempty\""
			UpdatedByUuid             *string "json:\"updated_by_uuid,omitempty\""
			Version                   *int    "json:\"version,omitempty\""
		} "json:\"access_groups,omitempty\""
		Pagination *struct {
			Limit  *int32 "json:\"limit,omitempty\""
			Offset *int32 "json:\"offset,omitempty\""
			Sort   *[]struct {
				Name  *string "json:\"name,omitempty\""
				Order *string "json:\"order,omitempty\""
			} "json:\"sort,omitempty\""
			Total *int32 "json:\"total,omitempty\""
		} "json:\"pagination,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Exclusions implementation of the Exclusions interface
type Exclusions struct {
	*ClientWithResponses
}

// CreateWithBody calls the Exclusions ´s function
func (c *Exclusions) CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Members              *string "json:\"members,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NetworkId            *string "json:\"network_id,omitempty\""
	Schedule             *struct {
		Enabled *bool   "json:\"enabled,omitempty\""
		Endtime *string "json:\"endtime,omitempty\""
		Rrules  *struct {
			Bymonthday *int    "json:\"bymonthday,omitempty\""
			Byweekday  *string "json:\"byweekday,omitempty\""
			Freq       *string "json:\"freq,omitempty\""
			Interval   *int    "json:\"interval,omitempty\""
		} "json:\"rrules,omitempty\""
		Starttime *string "json:\"starttime,omitempty\""
		Timezone  *string "json:\"timezone,omitempty\""
	} "json:\"schedule,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExclusionsCreateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreationDate *int "json:\"creation_date,omitempty\""; Description *string "json:\"description,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Members *string "json:\"members,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkId *string "json:\"network_id,omitempty\""; Schedule *struct { Enabled *bool "json:\"enabled,omitempty\""; Endtime *string "json:\"endtime,omitempty\""; Rrules *struct { Bymonthday *int "json:\"bymonthday,omitempty\""; Byweekday *string "json:\"byweekday,omitempty\""; Freq *string "json:\"freq,omitempty\""; Interval *int "json:\"interval,omitempty\"" } "json:\"rrules,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; Timezone *string "json:\"timezone,omitempty\"" } "json:\"schedule,omitempty\"" }
	if i, ok := r.(*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Members              *string "json:\"members,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NetworkId            *string "json:\"network_id,omitempty\""
		Schedule             *struct {
			Enabled *bool   "json:\"enabled,omitempty\""
			Endtime *string "json:\"endtime,omitempty\""
			Rrules  *struct {
				Bymonthday *int    "json:\"bymonthday,omitempty\""
				Byweekday  *string "json:\"byweekday,omitempty\""
				Freq       *string "json:\"freq,omitempty\""
				Interval   *int    "json:\"interval,omitempty\""
			} "json:\"rrules,omitempty\""
			Starttime *string "json:\"starttime,omitempty\""
			Timezone  *string "json:\"timezone,omitempty\""
		} "json:\"schedule,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Create calls the Exclusions ´s function
func (c *Exclusions) Create(arg1 ExclusionsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Members              *string "json:\"members,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NetworkId            *string "json:\"network_id,omitempty\""
	Schedule             *struct {
		Enabled *bool   "json:\"enabled,omitempty\""
		Endtime *string "json:\"endtime,omitempty\""
		Rrules  *struct {
			Bymonthday *int    "json:\"bymonthday,omitempty\""
			Byweekday  *string "json:\"byweekday,omitempty\""
			Freq       *string "json:\"freq,omitempty\""
			Interval   *int    "json:\"interval,omitempty\""
		} "json:\"rrules,omitempty\""
		Starttime *string "json:\"starttime,omitempty\""
		Timezone  *string "json:\"timezone,omitempty\""
	} "json:\"schedule,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExclusionsCreateWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreationDate *int "json:\"creation_date,omitempty\""; Description *string "json:\"description,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Members *string "json:\"members,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkId *string "json:\"network_id,omitempty\""; Schedule *struct { Enabled *bool "json:\"enabled,omitempty\""; Endtime *string "json:\"endtime,omitempty\""; Rrules *struct { Bymonthday *int "json:\"bymonthday,omitempty\""; Byweekday *string "json:\"byweekday,omitempty\""; Freq *string "json:\"freq,omitempty\""; Interval *int "json:\"interval,omitempty\"" } "json:\"rrules,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; Timezone *string "json:\"timezone,omitempty\"" } "json:\"schedule,omitempty\"" }
	if i, ok := r.(*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Members              *string "json:\"members,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NetworkId            *string "json:\"network_id,omitempty\""
		Schedule             *struct {
			Enabled *bool   "json:\"enabled,omitempty\""
			Endtime *string "json:\"endtime,omitempty\""
			Rrules  *struct {
				Bymonthday *int    "json:\"bymonthday,omitempty\""
				Byweekday  *string "json:\"byweekday,omitempty\""
				Freq       *string "json:\"freq,omitempty\""
				Interval   *int    "json:\"interval,omitempty\""
			} "json:\"rrules,omitempty\""
			Starttime *string "json:\"starttime,omitempty\""
			Timezone  *string "json:\"timezone,omitempty\""
		} "json:\"schedule,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Delete calls the Exclusions ´s function
func (c *Exclusions) Delete(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExclusionsDeleteWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Details calls the Exclusions ´s function
func (c *Exclusions) Details(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Members              *string "json:\"members,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NetworkId            *string "json:\"network_id,omitempty\""
	Schedule             *struct {
		Enabled *bool   "json:\"enabled,omitempty\""
		Endtime *string "json:\"endtime,omitempty\""
		Rrules  *struct {
			Bymonthday *int    "json:\"bymonthday,omitempty\""
			Byweekday  *string "json:\"byweekday,omitempty\""
			Freq       *string "json:\"freq,omitempty\""
			Interval   *int    "json:\"interval,omitempty\""
		} "json:\"rrules,omitempty\""
		Starttime *string "json:\"starttime,omitempty\""
		Timezone  *string "json:\"timezone,omitempty\""
	} "json:\"schedule,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExclusionsDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreationDate *int "json:\"creation_date,omitempty\""; Description *string "json:\"description,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Members *string "json:\"members,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkId *string "json:\"network_id,omitempty\""; Schedule *struct { Enabled *bool "json:\"enabled,omitempty\""; Endtime *string "json:\"endtime,omitempty\""; Rrules *struct { Bymonthday *int "json:\"bymonthday,omitempty\""; Byweekday *string "json:\"byweekday,omitempty\""; Freq *string "json:\"freq,omitempty\""; Interval *int "json:\"interval,omitempty\"" } "json:\"rrules,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; Timezone *string "json:\"timezone,omitempty\"" } "json:\"schedule,omitempty\"" }
	if i, ok := r.(*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Members              *string "json:\"members,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NetworkId            *string "json:\"network_id,omitempty\""
		Schedule             *struct {
			Enabled *bool   "json:\"enabled,omitempty\""
			Endtime *string "json:\"endtime,omitempty\""
			Rrules  *struct {
				Bymonthday *int    "json:\"bymonthday,omitempty\""
				Byweekday  *string "json:\"byweekday,omitempty\""
				Freq       *string "json:\"freq,omitempty\""
				Interval   *int    "json:\"interval,omitempty\""
			} "json:\"rrules,omitempty\""
			Starttime *string "json:\"starttime,omitempty\""
			Timezone  *string "json:\"timezone,omitempty\""
		} "json:\"schedule,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// EditWithBody calls the Exclusions ´s function
func (c *Exclusions) EditWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExclusionsEditWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Edit calls the Exclusions ´s function
func (c *Exclusions) Edit(arg1 int32, arg2 ExclusionsEditJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExclusionsEditWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ImportWithBody calls the Exclusions ´s function
func (c *Exclusions) ImportWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExclusionsImportWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Import calls the Exclusions ´s function
func (c *Exclusions) Import(arg1 ExclusionsImportJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExclusionsImportWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// List calls the Exclusions ´s function
func (c *Exclusions) List(reqEditors ...RequestEditorFn) (*[]struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Members              *string "json:\"members,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NetworkId            *string "json:\"network_id,omitempty\""
	Schedule             *struct {
		Enabled *bool   "json:\"enabled,omitempty\""
		Endtime *string "json:\"endtime,omitempty\""
		Rrules  *struct {
			Bymonthday *int    "json:\"bymonthday,omitempty\""
			Byweekday  *string "json:\"byweekday,omitempty\""
			Freq       *string "json:\"freq,omitempty\""
			Interval   *int    "json:\"interval,omitempty\""
		} "json:\"rrules,omitempty\""
		Starttime *string "json:\"starttime,omitempty\""
		Timezone  *string "json:\"timezone,omitempty\""
	} "json:\"schedule,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExclusionsListWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { CreationDate *int "json:\"creation_date,omitempty\""; Description *string "json:\"description,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Members *string "json:\"members,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkId *string "json:\"network_id,omitempty\""; Schedule *struct { Enabled *bool "json:\"enabled,omitempty\""; Endtime *string "json:\"endtime,omitempty\""; Rrules *struct { Bymonthday *int "json:\"bymonthday,omitempty\""; Byweekday *string "json:\"byweekday,omitempty\""; Freq *string "json:\"freq,omitempty\""; Interval *int "json:\"interval,omitempty\"" } "json:\"rrules,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; Timezone *string "json:\"timezone,omitempty\"" } "json:\"schedule,omitempty\"" }
	if i, ok := r.(*[]struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Members              *string "json:\"members,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NetworkId            *string "json:\"network_id,omitempty\""
		Schedule             *struct {
			Enabled *bool   "json:\"enabled,omitempty\""
			Endtime *string "json:\"endtime,omitempty\""
			Rrules  *struct {
				Bymonthday *int    "json:\"bymonthday,omitempty\""
				Byweekday  *string "json:\"byweekday,omitempty\""
				Freq       *string "json:\"freq,omitempty\""
				Interval   *int    "json:\"interval,omitempty\""
			} "json:\"rrules,omitempty\""
			Starttime *string "json:\"starttime,omitempty\""
			Timezone  *string "json:\"timezone,omitempty\""
		} "json:\"schedule,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// IoScanner implementation of the IoScanner interface
type IoScanner struct {
	*ClientWithResponses
}

// GroupsListRoutes calls the IoScanner ´s function
func (c *IoScanner) GroupsListRoutes(arg1 int32, reqEditors ...RequestEditorFn) (*[]struct {
	Route string "json:\"route\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoScannerGroupsListRoutesWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { Route string "json:\"route\"" }
	if i, ok := r.(*[]struct {
		Route string "json:\"route\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GroupsUpdateRoutesWithBody calls the IoScanner ´s function
func (c *IoScanner) GroupsUpdateRoutesWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoScannerGroupsUpdateRoutesWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GroupsUpdateRoutes calls the IoScanner ´s function
func (c *IoScanner) GroupsUpdateRoutes(arg1 int32, arg2 IoScannerGroupsUpdateRoutesJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoScannerGroupsUpdateRoutesWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// IoScans implementation of the IoScans interface
type IoScans struct {
	*ClientWithResponses
}

// CheckAutoTargetsWithBody calls the IoScans ´s function
func (c *IoScans) CheckAutoTargetsWithBody(params *IoScansCheckAutoTargetsParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	MatchedResourceUuids      *[]string "json:\"matched_resource_uuids,omitempty\""
	MissedTargets             *[]string "json:\"missed_targets,omitempty\""
	TotalMatchedResourceUuids *int      "json:\"total_matched_resource_uuids,omitempty\""
	TotalMissedTargets        *int      "json:\"total_missed_targets,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoScansCheckAutoTargetsWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { MatchedResourceUuids *[]string "json:\"matched_resource_uuids,omitempty\""; MissedTargets *[]string "json:\"missed_targets,omitempty\""; TotalMatchedResourceUuids *int "json:\"total_matched_resource_uuids,omitempty\""; TotalMissedTargets *int "json:\"total_missed_targets,omitempty\"" }
	if i, ok := r.(*struct {
		MatchedResourceUuids      *[]string "json:\"matched_resource_uuids,omitempty\""
		MissedTargets             *[]string "json:\"missed_targets,omitempty\""
		TotalMatchedResourceUuids *int      "json:\"total_matched_resource_uuids,omitempty\""
		TotalMissedTargets        *int      "json:\"total_missed_targets,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CheckAutoTargets calls the IoScans ´s function
func (c *IoScans) CheckAutoTargets(params *IoScansCheckAutoTargetsParams, arg2 IoScansCheckAutoTargetsJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	MatchedResourceUuids      *[]string "json:\"matched_resource_uuids,omitempty\""
	MissedTargets             *[]string "json:\"missed_targets,omitempty\""
	TotalMatchedResourceUuids *int      "json:\"total_matched_resource_uuids,omitempty\""
	TotalMissedTargets        *int      "json:\"total_missed_targets,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoScansCheckAutoTargetsWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { MatchedResourceUuids *[]string "json:\"matched_resource_uuids,omitempty\""; MissedTargets *[]string "json:\"missed_targets,omitempty\""; TotalMatchedResourceUuids *int "json:\"total_matched_resource_uuids,omitempty\""; TotalMissedTargets *int "json:\"total_missed_targets,omitempty\"" }
	if i, ok := r.(*struct {
		MatchedResourceUuids      *[]string "json:\"matched_resource_uuids,omitempty\""
		MissedTargets             *[]string "json:\"missed_targets,omitempty\""
		TotalMatchedResourceUuids *int      "json:\"total_matched_resource_uuids,omitempty\""
		TotalMissedTargets        *int      "json:\"total_missed_targets,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CredentialsConvertWithBody calls the IoScans ´s function
func (c *IoScans) CredentialsConvertWithBody(arg1 int, arg2 int, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	Uuid *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoScansCredentialsConvertWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		Uuid *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CredentialsConvert calls the IoScans ´s function
func (c *IoScans) CredentialsConvert(arg1 int, arg2 int, arg3 IoScansCredentialsConvertJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	Uuid *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoScansCredentialsConvertWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		Uuid *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RemediationCreateWithBody calls the IoScans ´s function
func (c *IoScans) RemediationCreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	ContainerId          *string "json:\"container_id,omitempty\""
	CreationDate         *int32  "json:\"creation_date,omitempty\""
	CustomTargets        *string "json:\"custom_targets,omitempty\""
	DashboardFile        *string "json:\"dashboard_file,omitempty\""
	DefaultPermissions   *int32  "json:\"default_permissions,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Emails               *string "json:\"emails,omitempty\""
	Enabled              *bool   "json:\"enabled,omitempty\""
	Id                   *int32  "json:\"id,omitempty\""
	IncludeAggregate     *bool   "json:\"include_aggregate,omitempty\""
	LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NotificationFilters  *[]struct {
		Filter  *string "json:\"filter,omitempty\""
		Quality *string "json:\"quality,omitempty\""
		Value   *string "json:\"value,omitempty\""
	} "json:\"notification_filters,omitempty\""
	Owner             *string   "json:\"owner,omitempty\""
	OwnerId           *int      "json:\"owner_id,omitempty\""
	OwnerUuid         *string   "json:\"owner_uuid,omitempty\""
	PolicyId          *int      "json:\"policy_id,omitempty\""
	Remediation       *int      "json:\"remediation,omitempty\""
	Rrules            *string   "json:\"rrules,omitempty\""
	ScanTimeWindow    *string   "json:\"scan_time_window,omitempty\""
	ScannerUuid       *string   "json:\"scanner_uuid,omitempty\""
	Shared            *bool     "json:\"shared,omitempty\""
	Sms               *string   "json:\"sms,omitempty\""
	Starttime         *string   "json:\"starttime,omitempty\""
	TagTargets        *[]string "json:\"tag_targets,omitempty\""
	TagType           *string   "json:\"tag_type,omitempty\""
	TargetNetworkUuid *string   "json:\"target_network_uuid,omitempty\""
	Timezone          *string   "json:\"timezone,omitempty\""
	Type              *string   "json:\"type,omitempty\""
	UserPermissions   *int32    "json:\"user_permissions,omitempty\""
	Uuid              *string   "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoScansRemediationCreateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ContainerId *string "json:\"container_id,omitempty\""; CreationDate *int32 "json:\"creation_date,omitempty\""; CustomTargets *string "json:\"custom_targets,omitempty\""; DashboardFile *string "json:\"dashboard_file,omitempty\""; DefaultPermissions *int32 "json:\"default_permissions,omitempty\""; Description *string "json:\"description,omitempty\""; Emails *string "json:\"emails,omitempty\""; Enabled *bool "json:\"enabled,omitempty\""; Id *int32 "json:\"id,omitempty\""; IncludeAggregate *bool "json:\"include_aggregate,omitempty\""; LastModificationDate *int32 "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NotificationFilters *[]struct { Filter *string "json:\"filter,omitempty\""; Quality *string "json:\"quality,omitempty\""; Value *string "json:\"value,omitempty\"" } "json:\"notification_filters,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; PolicyId *int "json:\"policy_id,omitempty\""; Remediation *int "json:\"remediation,omitempty\""; Rrules *string "json:\"rrules,omitempty\""; ScanTimeWindow *string "json:\"scan_time_window,omitempty\""; ScannerUuid *string "json:\"scanner_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Sms *string "json:\"sms,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; TagTargets *[]string "json:\"tag_targets,omitempty\""; TagType *string "json:\"tag_type,omitempty\""; TargetNetworkUuid *string "json:\"target_network_uuid,omitempty\""; Timezone *string "json:\"timezone,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int32 "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		ContainerId          *string "json:\"container_id,omitempty\""
		CreationDate         *int32  "json:\"creation_date,omitempty\""
		CustomTargets        *string "json:\"custom_targets,omitempty\""
		DashboardFile        *string "json:\"dashboard_file,omitempty\""
		DefaultPermissions   *int32  "json:\"default_permissions,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Emails               *string "json:\"emails,omitempty\""
		Enabled              *bool   "json:\"enabled,omitempty\""
		Id                   *int32  "json:\"id,omitempty\""
		IncludeAggregate     *bool   "json:\"include_aggregate,omitempty\""
		LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NotificationFilters  *[]struct {
			Filter  *string "json:\"filter,omitempty\""
			Quality *string "json:\"quality,omitempty\""
			Value   *string "json:\"value,omitempty\""
		} "json:\"notification_filters,omitempty\""
		Owner             *string   "json:\"owner,omitempty\""
		OwnerId           *int      "json:\"owner_id,omitempty\""
		OwnerUuid         *string   "json:\"owner_uuid,omitempty\""
		PolicyId          *int      "json:\"policy_id,omitempty\""
		Remediation       *int      "json:\"remediation,omitempty\""
		Rrules            *string   "json:\"rrules,omitempty\""
		ScanTimeWindow    *string   "json:\"scan_time_window,omitempty\""
		ScannerUuid       *string   "json:\"scanner_uuid,omitempty\""
		Shared            *bool     "json:\"shared,omitempty\""
		Sms               *string   "json:\"sms,omitempty\""
		Starttime         *string   "json:\"starttime,omitempty\""
		TagTargets        *[]string "json:\"tag_targets,omitempty\""
		TagType           *string   "json:\"tag_type,omitempty\""
		TargetNetworkUuid *string   "json:\"target_network_uuid,omitempty\""
		Timezone          *string   "json:\"timezone,omitempty\""
		Type              *string   "json:\"type,omitempty\""
		UserPermissions   *int32    "json:\"user_permissions,omitempty\""
		Uuid              *string   "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RemediationCreate calls the IoScans ´s function
func (c *IoScans) RemediationCreate(arg1 IoScansRemediationCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	ContainerId          *string "json:\"container_id,omitempty\""
	CreationDate         *int32  "json:\"creation_date,omitempty\""
	CustomTargets        *string "json:\"custom_targets,omitempty\""
	DashboardFile        *string "json:\"dashboard_file,omitempty\""
	DefaultPermissions   *int32  "json:\"default_permissions,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Emails               *string "json:\"emails,omitempty\""
	Enabled              *bool   "json:\"enabled,omitempty\""
	Id                   *int32  "json:\"id,omitempty\""
	IncludeAggregate     *bool   "json:\"include_aggregate,omitempty\""
	LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NotificationFilters  *[]struct {
		Filter  *string "json:\"filter,omitempty\""
		Quality *string "json:\"quality,omitempty\""
		Value   *string "json:\"value,omitempty\""
	} "json:\"notification_filters,omitempty\""
	Owner             *string   "json:\"owner,omitempty\""
	OwnerId           *int      "json:\"owner_id,omitempty\""
	OwnerUuid         *string   "json:\"owner_uuid,omitempty\""
	PolicyId          *int      "json:\"policy_id,omitempty\""
	Remediation       *int      "json:\"remediation,omitempty\""
	Rrules            *string   "json:\"rrules,omitempty\""
	ScanTimeWindow    *string   "json:\"scan_time_window,omitempty\""
	ScannerUuid       *string   "json:\"scanner_uuid,omitempty\""
	Shared            *bool     "json:\"shared,omitempty\""
	Sms               *string   "json:\"sms,omitempty\""
	Starttime         *string   "json:\"starttime,omitempty\""
	TagTargets        *[]string "json:\"tag_targets,omitempty\""
	TagType           *string   "json:\"tag_type,omitempty\""
	TargetNetworkUuid *string   "json:\"target_network_uuid,omitempty\""
	Timezone          *string   "json:\"timezone,omitempty\""
	Type              *string   "json:\"type,omitempty\""
	UserPermissions   *int32    "json:\"user_permissions,omitempty\""
	Uuid              *string   "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoScansRemediationCreateWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ContainerId *string "json:\"container_id,omitempty\""; CreationDate *int32 "json:\"creation_date,omitempty\""; CustomTargets *string "json:\"custom_targets,omitempty\""; DashboardFile *string "json:\"dashboard_file,omitempty\""; DefaultPermissions *int32 "json:\"default_permissions,omitempty\""; Description *string "json:\"description,omitempty\""; Emails *string "json:\"emails,omitempty\""; Enabled *bool "json:\"enabled,omitempty\""; Id *int32 "json:\"id,omitempty\""; IncludeAggregate *bool "json:\"include_aggregate,omitempty\""; LastModificationDate *int32 "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NotificationFilters *[]struct { Filter *string "json:\"filter,omitempty\""; Quality *string "json:\"quality,omitempty\""; Value *string "json:\"value,omitempty\"" } "json:\"notification_filters,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; PolicyId *int "json:\"policy_id,omitempty\""; Remediation *int "json:\"remediation,omitempty\""; Rrules *string "json:\"rrules,omitempty\""; ScanTimeWindow *string "json:\"scan_time_window,omitempty\""; ScannerUuid *string "json:\"scanner_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Sms *string "json:\"sms,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; TagTargets *[]string "json:\"tag_targets,omitempty\""; TagType *string "json:\"tag_type,omitempty\""; TargetNetworkUuid *string "json:\"target_network_uuid,omitempty\""; Timezone *string "json:\"timezone,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int32 "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		ContainerId          *string "json:\"container_id,omitempty\""
		CreationDate         *int32  "json:\"creation_date,omitempty\""
		CustomTargets        *string "json:\"custom_targets,omitempty\""
		DashboardFile        *string "json:\"dashboard_file,omitempty\""
		DefaultPermissions   *int32  "json:\"default_permissions,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Emails               *string "json:\"emails,omitempty\""
		Enabled              *bool   "json:\"enabled,omitempty\""
		Id                   *int32  "json:\"id,omitempty\""
		IncludeAggregate     *bool   "json:\"include_aggregate,omitempty\""
		LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NotificationFilters  *[]struct {
			Filter  *string "json:\"filter,omitempty\""
			Quality *string "json:\"quality,omitempty\""
			Value   *string "json:\"value,omitempty\""
		} "json:\"notification_filters,omitempty\""
		Owner             *string   "json:\"owner,omitempty\""
		OwnerId           *int      "json:\"owner_id,omitempty\""
		OwnerUuid         *string   "json:\"owner_uuid,omitempty\""
		PolicyId          *int      "json:\"policy_id,omitempty\""
		Remediation       *int      "json:\"remediation,omitempty\""
		Rrules            *string   "json:\"rrules,omitempty\""
		ScanTimeWindow    *string   "json:\"scan_time_window,omitempty\""
		ScannerUuid       *string   "json:\"scanner_uuid,omitempty\""
		Shared            *bool     "json:\"shared,omitempty\""
		Sms               *string   "json:\"sms,omitempty\""
		Starttime         *string   "json:\"starttime,omitempty\""
		TagTargets        *[]string "json:\"tag_targets,omitempty\""
		TagType           *string   "json:\"tag_type,omitempty\""
		TargetNetworkUuid *string   "json:\"target_network_uuid,omitempty\""
		Timezone          *string   "json:\"timezone,omitempty\""
		Type              *string   "json:\"type,omitempty\""
		UserPermissions   *int32    "json:\"user_permissions,omitempty\""
		Uuid              *string   "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RemediationList calls the IoScans ´s function
func (c *IoScans) RemediationList(params *IoScansRemediationListParams, reqEditors ...RequestEditorFn) (*struct {
	Pagination *struct {
		Limit  *int32 "json:\"limit,omitempty\""
		Offset *int32 "json:\"offset,omitempty\""
		Sort   *[]struct {
			Name  *string "json:\"name,omitempty\""
			Order *string "json:\"order,omitempty\""
		} "json:\"sort,omitempty\""
		Total *int32 "json:\"total,omitempty\""
	} "json:\"pagination,omitempty\""
	Scans *[]struct {
		Control              *bool   "json:\"control,omitempty\""
		CreationDate         *int32  "json:\"creation_date,omitempty\""
		Enabled              *bool   "json:\"enabled,omitempty\""
		Id                   *int32  "json:\"id,omitempty\""
		LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		Permissions          *int32  "json:\"permissions,omitempty\""
		PolicyId             *int    "json:\"policy_id,omitempty\""
		Read                 *bool   "json:\"read,omitempty\""
		Remediation          *int    "json:\"remediation,omitempty\""
		ScanCreationDate     *int32  "json:\"scan_creation_date,omitempty\""
		ScheduleUuid         *string "json:\"schedule_uuid,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Status               *string "json:\"status,omitempty\""
		TemplateUuid         *string "json:\"template_uuid,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int32  "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
		WizardUuid           *string "json:\"wizard_uuid,omitempty\""
	} "json:\"scans,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoScansRemediationListWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Pagination *struct { Limit *int32 "json:\"limit,omitempty\""; Offset *int32 "json:\"offset,omitempty\""; Sort *[]struct { Name *string "json:\"name,omitempty\""; Order *string "json:\"order,omitempty\"" } "json:\"sort,omitempty\""; Total *int32 "json:\"total,omitempty\"" } "json:\"pagination,omitempty\""; Scans *[]struct { Control *bool "json:\"control,omitempty\""; CreationDate *int32 "json:\"creation_date,omitempty\""; Enabled *bool "json:\"enabled,omitempty\""; Id *int32 "json:\"id,omitempty\""; LastModificationDate *int32 "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; Permissions *int32 "json:\"permissions,omitempty\""; PolicyId *int "json:\"policy_id,omitempty\""; Read *bool "json:\"read,omitempty\""; Remediation *int "json:\"remediation,omitempty\""; ScanCreationDate *int32 "json:\"scan_creation_date,omitempty\""; ScheduleUuid *string "json:\"schedule_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Status *string "json:\"status,omitempty\""; TemplateUuid *string "json:\"template_uuid,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int32 "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\""; WizardUuid *string "json:\"wizard_uuid,omitempty\"" } "json:\"scans,omitempty\"" }
	if i, ok := r.(*struct {
		Pagination *struct {
			Limit  *int32 "json:\"limit,omitempty\""
			Offset *int32 "json:\"offset,omitempty\""
			Sort   *[]struct {
				Name  *string "json:\"name,omitempty\""
				Order *string "json:\"order,omitempty\""
			} "json:\"sort,omitempty\""
			Total *int32 "json:\"total,omitempty\""
		} "json:\"pagination,omitempty\""
		Scans *[]struct {
			Control              *bool   "json:\"control,omitempty\""
			CreationDate         *int32  "json:\"creation_date,omitempty\""
			Enabled              *bool   "json:\"enabled,omitempty\""
			Id                   *int32  "json:\"id,omitempty\""
			LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
			Name                 *string "json:\"name,omitempty\""
			Owner                *string "json:\"owner,omitempty\""
			Permissions          *int32  "json:\"permissions,omitempty\""
			PolicyId             *int    "json:\"policy_id,omitempty\""
			Read                 *bool   "json:\"read,omitempty\""
			Remediation          *int    "json:\"remediation,omitempty\""
			ScanCreationDate     *int32  "json:\"scan_creation_date,omitempty\""
			ScheduleUuid         *string "json:\"schedule_uuid,omitempty\""
			Shared               *bool   "json:\"shared,omitempty\""
			Status               *string "json:\"status,omitempty\""
			TemplateUuid         *string "json:\"template_uuid,omitempty\""
			Type                 *string "json:\"type,omitempty\""
			UserPermissions      *int32  "json:\"user_permissions,omitempty\""
			Uuid                 *string "json:\"uuid,omitempty\""
			WizardUuid           *string "json:\"wizard_uuid,omitempty\""
		} "json:\"scans,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// IoV1 implementation of the IoV1 interface
type IoV1 struct {
	*ClientWithResponses
}

// AccessGroupsCreateWithBody calls the IoV1 ´s function
func (c *IoV1) AccessGroupsCreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	AllAssets                 *bool   "json:\"all_assets,omitempty\""
	AllUsers                  *bool   "json:\"all_users,omitempty\""
	ContainerUuid             *string "json:\"container_uuid,omitempty\""
	CreatedAt                 *string "json:\"created_at,omitempty\""
	CreatedByName             *string "json:\"created_by_name,omitempty\""
	CreatedByUuid             *string "json:\"created_by_uuid,omitempty\""
	Id                        *string "json:\"id,omitempty\""
	Name                      *string "json:\"name,omitempty\""
	ProcessingPercentComplete *int    "json:\"processing_percent_complete,omitempty\""
	Status                    *string "json:\"status,omitempty\""
	UpdatedAt                 *string "json:\"updated_at,omitempty\""
	UpdatedByName             *string "json:\"updated_by_name,omitempty\""
	UpdatedByUuid             *string "json:\"updated_by_uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV1AccessGroupsCreateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AllAssets *bool "json:\"all_assets,omitempty\""; AllUsers *bool "json:\"all_users,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedByName *string "json:\"created_by_name,omitempty\""; CreatedByUuid *string "json:\"created_by_uuid,omitempty\""; Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""; Status *string "json:\"status,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedByName *string "json:\"updated_by_name,omitempty\""; UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\"" }
	if i, ok := r.(*struct {
		AllAssets                 *bool   "json:\"all_assets,omitempty\""
		AllUsers                  *bool   "json:\"all_users,omitempty\""
		ContainerUuid             *string "json:\"container_uuid,omitempty\""
		CreatedAt                 *string "json:\"created_at,omitempty\""
		CreatedByName             *string "json:\"created_by_name,omitempty\""
		CreatedByUuid             *string "json:\"created_by_uuid,omitempty\""
		Id                        *string "json:\"id,omitempty\""
		Name                      *string "json:\"name,omitempty\""
		ProcessingPercentComplete *int    "json:\"processing_percent_complete,omitempty\""
		Status                    *string "json:\"status,omitempty\""
		UpdatedAt                 *string "json:\"updated_at,omitempty\""
		UpdatedByName             *string "json:\"updated_by_name,omitempty\""
		UpdatedByUuid             *string "json:\"updated_by_uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsCreate calls the IoV1 ´s function
func (c *IoV1) AccessGroupsCreate(arg1 IoV1AccessGroupsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	AllAssets                 *bool   "json:\"all_assets,omitempty\""
	AllUsers                  *bool   "json:\"all_users,omitempty\""
	ContainerUuid             *string "json:\"container_uuid,omitempty\""
	CreatedAt                 *string "json:\"created_at,omitempty\""
	CreatedByName             *string "json:\"created_by_name,omitempty\""
	CreatedByUuid             *string "json:\"created_by_uuid,omitempty\""
	Id                        *string "json:\"id,omitempty\""
	Name                      *string "json:\"name,omitempty\""
	ProcessingPercentComplete *int    "json:\"processing_percent_complete,omitempty\""
	Status                    *string "json:\"status,omitempty\""
	UpdatedAt                 *string "json:\"updated_at,omitempty\""
	UpdatedByName             *string "json:\"updated_by_name,omitempty\""
	UpdatedByUuid             *string "json:\"updated_by_uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV1AccessGroupsCreateWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AllAssets *bool "json:\"all_assets,omitempty\""; AllUsers *bool "json:\"all_users,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedByName *string "json:\"created_by_name,omitempty\""; CreatedByUuid *string "json:\"created_by_uuid,omitempty\""; Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""; Status *string "json:\"status,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedByName *string "json:\"updated_by_name,omitempty\""; UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\"" }
	if i, ok := r.(*struct {
		AllAssets                 *bool   "json:\"all_assets,omitempty\""
		AllUsers                  *bool   "json:\"all_users,omitempty\""
		ContainerUuid             *string "json:\"container_uuid,omitempty\""
		CreatedAt                 *string "json:\"created_at,omitempty\""
		CreatedByName             *string "json:\"created_by_name,omitempty\""
		CreatedByUuid             *string "json:\"created_by_uuid,omitempty\""
		Id                        *string "json:\"id,omitempty\""
		Name                      *string "json:\"name,omitempty\""
		ProcessingPercentComplete *int    "json:\"processing_percent_complete,omitempty\""
		Status                    *string "json:\"status,omitempty\""
		UpdatedAt                 *string "json:\"updated_at,omitempty\""
		UpdatedByName             *string "json:\"updated_by_name,omitempty\""
		UpdatedByUuid             *string "json:\"updated_by_uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsDelete calls the IoV1 ´s function
func (c *IoV1) AccessGroupsDelete(contentType string, reqEditors ...RequestEditorFn) (*IoV1AccessGroupsDeleteResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV1AccessGroupsDeleteWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *IoV1AccessGroupsDeleteResponse
	if i, ok := r.(*IoV1AccessGroupsDeleteResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsDetails calls the IoV1 ´s function
func (c *IoV1) AccessGroupsDetails(contentType string, reqEditors ...RequestEditorFn) (*struct {
	AllAssets     *bool   "json:\"all_assets,omitempty\""
	AllUsers      *bool   "json:\"all_users,omitempty\""
	ContainerUuid *string "json:\"container_uuid,omitempty\""
	CreatedAt     *string "json:\"created_at,omitempty\""
	CreatedByName *string "json:\"created_by_name,omitempty\""
	CreatedByUuid *string "json:\"created_by_uuid,omitempty\""
	Id            *string "json:\"id,omitempty\""
	Name          *string "json:\"name,omitempty\""
	Principals    *[]struct {
		PrincipalId   *string "json:\"principal_id,omitempty\""
		PrincipalName *string "json:\"principal_name,omitempty\""
		Type          *string "json:\"type,omitempty\""
	} "json:\"principals,omitempty\""
	ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""
	Rules                     *[]struct {
		Operator *string   "json:\"operator,omitempty\""
		Terms    *[]string "json:\"terms,omitempty\""
		Type     *string   "json:\"type,omitempty\""
	} "json:\"rules,omitempty\""
	Status        *string "json:\"status,omitempty\""
	UpdatedAt     *string "json:\"updated_at,omitempty\""
	UpdatedByName *string "json:\"updated_by_name,omitempty\""
	UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV1AccessGroupsDetailsWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AllAssets *bool "json:\"all_assets,omitempty\""; AllUsers *bool "json:\"all_users,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedByName *string "json:\"created_by_name,omitempty\""; CreatedByUuid *string "json:\"created_by_uuid,omitempty\""; Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Principals *[]struct { PrincipalId *string "json:\"principal_id,omitempty\""; PrincipalName *string "json:\"principal_name,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"principals,omitempty\""; ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""; Rules *[]struct { Operator *string "json:\"operator,omitempty\""; Terms *[]string "json:\"terms,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"rules,omitempty\""; Status *string "json:\"status,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedByName *string "json:\"updated_by_name,omitempty\""; UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\"" }
	if i, ok := r.(*struct {
		AllAssets     *bool   "json:\"all_assets,omitempty\""
		AllUsers      *bool   "json:\"all_users,omitempty\""
		ContainerUuid *string "json:\"container_uuid,omitempty\""
		CreatedAt     *string "json:\"created_at,omitempty\""
		CreatedByName *string "json:\"created_by_name,omitempty\""
		CreatedByUuid *string "json:\"created_by_uuid,omitempty\""
		Id            *string "json:\"id,omitempty\""
		Name          *string "json:\"name,omitempty\""
		Principals    *[]struct {
			PrincipalId   *string "json:\"principal_id,omitempty\""
			PrincipalName *string "json:\"principal_name,omitempty\""
			Type          *string "json:\"type,omitempty\""
		} "json:\"principals,omitempty\""
		ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""
		Rules                     *[]struct {
			Operator *string   "json:\"operator,omitempty\""
			Terms    *[]string "json:\"terms,omitempty\""
			Type     *string   "json:\"type,omitempty\""
		} "json:\"rules,omitempty\""
		Status        *string "json:\"status,omitempty\""
		UpdatedAt     *string "json:\"updated_at,omitempty\""
		UpdatedByName *string "json:\"updated_by_name,omitempty\""
		UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsEditWithBody calls the IoV1 ´s function
func (c *IoV1) AccessGroupsEditWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	AllAssets     *bool   "json:\"all_assets,omitempty\""
	AllUsers      *bool   "json:\"all_users,omitempty\""
	ContainerUuid *string "json:\"container_uuid,omitempty\""
	CreatedAt     *string "json:\"created_at,omitempty\""
	CreatedByName *string "json:\"created_by_name,omitempty\""
	CreatedByUuid *string "json:\"created_by_uuid,omitempty\""
	Id            *string "json:\"id,omitempty\""
	Name          *string "json:\"name,omitempty\""
	Principals    *[]struct {
		PrincipalId   *string "json:\"principal_id,omitempty\""
		PrincipalName *string "json:\"principal_name,omitempty\""
		Type          *string "json:\"type,omitempty\""
	} "json:\"principals,omitempty\""
	ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""
	Rules                     *[]struct {
		Operator *string   "json:\"operator,omitempty\""
		Terms    *[]string "json:\"terms,omitempty\""
		Type     *string   "json:\"type,omitempty\""
	} "json:\"rules,omitempty\""
	Status        *string "json:\"status,omitempty\""
	UpdatedAt     *string "json:\"updated_at,omitempty\""
	UpdatedByName *string "json:\"updated_by_name,omitempty\""
	UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV1AccessGroupsEditWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AllAssets *bool "json:\"all_assets,omitempty\""; AllUsers *bool "json:\"all_users,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedByName *string "json:\"created_by_name,omitempty\""; CreatedByUuid *string "json:\"created_by_uuid,omitempty\""; Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Principals *[]struct { PrincipalId *string "json:\"principal_id,omitempty\""; PrincipalName *string "json:\"principal_name,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"principals,omitempty\""; ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""; Rules *[]struct { Operator *string "json:\"operator,omitempty\""; Terms *[]string "json:\"terms,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"rules,omitempty\""; Status *string "json:\"status,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedByName *string "json:\"updated_by_name,omitempty\""; UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\"" }
	if i, ok := r.(*struct {
		AllAssets     *bool   "json:\"all_assets,omitempty\""
		AllUsers      *bool   "json:\"all_users,omitempty\""
		ContainerUuid *string "json:\"container_uuid,omitempty\""
		CreatedAt     *string "json:\"created_at,omitempty\""
		CreatedByName *string "json:\"created_by_name,omitempty\""
		CreatedByUuid *string "json:\"created_by_uuid,omitempty\""
		Id            *string "json:\"id,omitempty\""
		Name          *string "json:\"name,omitempty\""
		Principals    *[]struct {
			PrincipalId   *string "json:\"principal_id,omitempty\""
			PrincipalName *string "json:\"principal_name,omitempty\""
			Type          *string "json:\"type,omitempty\""
		} "json:\"principals,omitempty\""
		ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""
		Rules                     *[]struct {
			Operator *string   "json:\"operator,omitempty\""
			Terms    *[]string "json:\"terms,omitempty\""
			Type     *string   "json:\"type,omitempty\""
		} "json:\"rules,omitempty\""
		Status        *string "json:\"status,omitempty\""
		UpdatedAt     *string "json:\"updated_at,omitempty\""
		UpdatedByName *string "json:\"updated_by_name,omitempty\""
		UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsEdit calls the IoV1 ´s function
func (c *IoV1) AccessGroupsEdit(arg1 string, arg2 IoV1AccessGroupsEditJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	AllAssets     *bool   "json:\"all_assets,omitempty\""
	AllUsers      *bool   "json:\"all_users,omitempty\""
	ContainerUuid *string "json:\"container_uuid,omitempty\""
	CreatedAt     *string "json:\"created_at,omitempty\""
	CreatedByName *string "json:\"created_by_name,omitempty\""
	CreatedByUuid *string "json:\"created_by_uuid,omitempty\""
	Id            *string "json:\"id,omitempty\""
	Name          *string "json:\"name,omitempty\""
	Principals    *[]struct {
		PrincipalId   *string "json:\"principal_id,omitempty\""
		PrincipalName *string "json:\"principal_name,omitempty\""
		Type          *string "json:\"type,omitempty\""
	} "json:\"principals,omitempty\""
	ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""
	Rules                     *[]struct {
		Operator *string   "json:\"operator,omitempty\""
		Terms    *[]string "json:\"terms,omitempty\""
		Type     *string   "json:\"type,omitempty\""
	} "json:\"rules,omitempty\""
	Status        *string "json:\"status,omitempty\""
	UpdatedAt     *string "json:\"updated_at,omitempty\""
	UpdatedByName *string "json:\"updated_by_name,omitempty\""
	UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV1AccessGroupsEditWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AllAssets *bool "json:\"all_assets,omitempty\""; AllUsers *bool "json:\"all_users,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedByName *string "json:\"created_by_name,omitempty\""; CreatedByUuid *string "json:\"created_by_uuid,omitempty\""; Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Principals *[]struct { PrincipalId *string "json:\"principal_id,omitempty\""; PrincipalName *string "json:\"principal_name,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"principals,omitempty\""; ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""; Rules *[]struct { Operator *string "json:\"operator,omitempty\""; Terms *[]string "json:\"terms,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"rules,omitempty\""; Status *string "json:\"status,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedByName *string "json:\"updated_by_name,omitempty\""; UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\"" }
	if i, ok := r.(*struct {
		AllAssets     *bool   "json:\"all_assets,omitempty\""
		AllUsers      *bool   "json:\"all_users,omitempty\""
		ContainerUuid *string "json:\"container_uuid,omitempty\""
		CreatedAt     *string "json:\"created_at,omitempty\""
		CreatedByName *string "json:\"created_by_name,omitempty\""
		CreatedByUuid *string "json:\"created_by_uuid,omitempty\""
		Id            *string "json:\"id,omitempty\""
		Name          *string "json:\"name,omitempty\""
		Principals    *[]struct {
			PrincipalId   *string "json:\"principal_id,omitempty\""
			PrincipalName *string "json:\"principal_name,omitempty\""
			Type          *string "json:\"type,omitempty\""
		} "json:\"principals,omitempty\""
		ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""
		Rules                     *[]struct {
			Operator *string   "json:\"operator,omitempty\""
			Terms    *[]string "json:\"terms,omitempty\""
			Type     *string   "json:\"type,omitempty\""
		} "json:\"rules,omitempty\""
		Status        *string "json:\"status,omitempty\""
		UpdatedAt     *string "json:\"updated_at,omitempty\""
		UpdatedByName *string "json:\"updated_by_name,omitempty\""
		UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsListFilters calls the IoV1 ´s function
func (c *IoV1) AccessGroupsListFilters(reqEditors ...RequestEditorFn) (*struct {
	Filters *[]struct {
		Control      *string   "json:\"control,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	} "json:\"filters,omitempty\""
	Sort *[]struct {
		SortableFields *string "json:\"sortable_fields,omitempty\""
	} "json:\"sort,omitempty\""
	WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV1AccessGroupsListFiltersWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Filters *[]struct { Control *string "json:\"control,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filters,omitempty\""; Sort *[]struct { SortableFields *string "json:\"sortable_fields,omitempty\"" } "json:\"sort,omitempty\""; WildcardFields *[]string "json:\"wildcard_fields,omitempty\"" }
	if i, ok := r.(*struct {
		Filters *[]struct {
			Control      *string   "json:\"control,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"filters,omitempty\""
		Sort *[]struct {
			SortableFields *string "json:\"sortable_fields,omitempty\""
		} "json:\"sort,omitempty\""
		WildcardFields *[]string "json:\"wildcard_fields,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsListRuleFilters calls the IoV1 ´s function
func (c *IoV1) AccessGroupsListRuleFilters(reqEditors ...RequestEditorFn) (*struct {
	Filters *[]struct {
		Control      *string   "json:\"control,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	} "json:\"filters,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV1AccessGroupsListRuleFiltersWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Filters *[]struct { Control *string "json:\"control,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filters,omitempty\"" }
	if i, ok := r.(*struct {
		Filters *[]struct {
			Control      *string   "json:\"control,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"filters,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AccessGroupsList calls the IoV1 ´s function
func (c *IoV1) AccessGroupsList(params *IoV1AccessGroupsListParams, reqEditors ...RequestEditorFn) (*struct {
	AccessGroups *struct {
		AllAssets                 *bool   "json:\"all_assets,omitempty\""
		AllUsers                  *bool   "json:\"all_users,omitempty\""
		ContainerUuid             *string "json:\"container_uuid,omitempty\""
		CreatedAt                 *string "json:\"created_at,omitempty\""
		CreatedByName             *string "json:\"created_by_name,omitempty\""
		CreatedByUuid             *string "json:\"created_by_uuid,omitempty\""
		Id                        *string "json:\"id,omitempty\""
		Name                      *string "json:\"name,omitempty\""
		ProcessingPercentComplete *int    "json:\"processing_percent_complete,omitempty\""
		Status                    *string "json:\"status,omitempty\""
		UpdatedAt                 *string "json:\"updated_at,omitempty\""
		UpdatedByName             *string "json:\"updated_by_name,omitempty\""
		UpdatedByUuid             *string "json:\"updated_by_uuid,omitempty\""
	} "json:\"access_groups,omitempty\""
	Pagination *struct {
		Limit  *int32 "json:\"limit,omitempty\""
		Offset *int32 "json:\"offset,omitempty\""
		Sort   *[]struct {
			Name  *string "json:\"name,omitempty\""
			Order *string "json:\"order,omitempty\""
		} "json:\"sort,omitempty\""
		Total *int32 "json:\"total,omitempty\""
	} "json:\"pagination,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoV1AccessGroupsListWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AccessGroups *struct { AllAssets *bool "json:\"all_assets,omitempty\""; AllUsers *bool "json:\"all_users,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedByName *string "json:\"created_by_name,omitempty\""; CreatedByUuid *string "json:\"created_by_uuid,omitempty\""; Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; ProcessingPercentComplete *int "json:\"processing_percent_complete,omitempty\""; Status *string "json:\"status,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedByName *string "json:\"updated_by_name,omitempty\""; UpdatedByUuid *string "json:\"updated_by_uuid,omitempty\"" } "json:\"access_groups,omitempty\""; Pagination *struct { Limit *int32 "json:\"limit,omitempty\""; Offset *int32 "json:\"offset,omitempty\""; Sort *[]struct { Name *string "json:\"name,omitempty\""; Order *string "json:\"order,omitempty\"" } "json:\"sort,omitempty\""; Total *int32 "json:\"total,omitempty\"" } "json:\"pagination,omitempty\"" }
	if i, ok := r.(*struct {
		AccessGroups *struct {
			AllAssets                 *bool   "json:\"all_assets,omitempty\""
			AllUsers                  *bool   "json:\"all_users,omitempty\""
			ContainerUuid             *string "json:\"container_uuid,omitempty\""
			CreatedAt                 *string "json:\"created_at,omitempty\""
			CreatedByName             *string "json:\"created_by_name,omitempty\""
			CreatedByUuid             *string "json:\"created_by_uuid,omitempty\""
			Id                        *string "json:\"id,omitempty\""
			Name                      *string "json:\"name,omitempty\""
			ProcessingPercentComplete *int    "json:\"processing_percent_complete,omitempty\""
			Status                    *string "json:\"status,omitempty\""
			UpdatedAt                 *string "json:\"updated_at,omitempty\""
			UpdatedByName             *string "json:\"updated_by_name,omitempty\""
			UpdatedByUuid             *string "json:\"updated_by_uuid,omitempty\""
		} "json:\"access_groups,omitempty\""
		Pagination *struct {
			Limit  *int32 "json:\"limit,omitempty\""
			Offset *int32 "json:\"offset,omitempty\""
			Sort   *[]struct {
				Name  *string "json:\"name,omitempty\""
				Order *string "json:\"order,omitempty\""
			} "json:\"sort,omitempty\""
			Total *int32 "json:\"total,omitempty\""
		} "json:\"pagination,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ExportsVulns implementation of the ExportsVulns interface
type ExportsVulns struct {
	*ClientWithResponses
}

// DownloadChunk calls the ExportsVulns ´s function
func (c *ExportsVulns) DownloadChunk(arg1 string, arg2 int32, reqEditors ...RequestEditorFn) (*ExportsVulnsDownloadChunkResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExportsVulnsDownloadChunkWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *ExportsVulnsDownloadChunkResponse
	if i, ok := r.(*ExportsVulnsDownloadChunkResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ExportCancel calls the ExportsVulns ´s function
func (c *ExportsVulns) ExportCancel(contentType string, reqEditors ...RequestEditorFn) (*struct {
	Status *string "json:\"status,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExportsVulnsExportCancelWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Status *string "json:\"status,omitempty\"" }
	if i, ok := r.(*struct {
		Status *string "json:\"status,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ExportStatusRecent calls the ExportsVulns ´s function
func (c *ExportsVulns) ExportStatusRecent(reqEditors ...RequestEditorFn) (*struct {
	Exports *[]struct {
		ChunksAvailableCount *int                    "json:\"chunks_available_count,omitempty\""
		Created              *int                    "json:\"created,omitempty\""
		EmptyChunksCount     *int                    "json:\"empty_chunks_count,omitempty\""
		Filters              *map[string]interface{} "json:\"filters,omitempty\""
		FinishedChunks       *int                    "json:\"finished_chunks,omitempty\""
		NumAssetsPerChunk    *int                    "json:\"num_assets_per_chunk,omitempty\""
		Status               *string                 "json:\"status,omitempty\""
		TotalChunks          *int                    "json:\"total_chunks,omitempty\""
		Uuid                 *string                 "json:\"uuid,omitempty\""
	} "json:\"exports,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExportsVulnsExportStatusRecentWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Exports *[]struct { ChunksAvailableCount *int "json:\"chunks_available_count,omitempty\""; Created *int "json:\"created,omitempty\""; EmptyChunksCount *int "json:\"empty_chunks_count,omitempty\""; Filters *map[string]interface {} "json:\"filters,omitempty\""; FinishedChunks *int "json:\"finished_chunks,omitempty\""; NumAssetsPerChunk *int "json:\"num_assets_per_chunk,omitempty\""; Status *string "json:\"status,omitempty\""; TotalChunks *int "json:\"total_chunks,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"exports,omitempty\"" }
	if i, ok := r.(*struct {
		Exports *[]struct {
			ChunksAvailableCount *int                    "json:\"chunks_available_count,omitempty\""
			Created              *int                    "json:\"created,omitempty\""
			EmptyChunksCount     *int                    "json:\"empty_chunks_count,omitempty\""
			Filters              *map[string]interface{} "json:\"filters,omitempty\""
			FinishedChunks       *int                    "json:\"finished_chunks,omitempty\""
			NumAssetsPerChunk    *int                    "json:\"num_assets_per_chunk,omitempty\""
			Status               *string                 "json:\"status,omitempty\""
			TotalChunks          *int                    "json:\"total_chunks,omitempty\""
			Uuid                 *string                 "json:\"uuid,omitempty\""
		} "json:\"exports,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ExportStatus calls the ExportsVulns ´s function
func (c *ExportsVulns) ExportStatus(contentType string, reqEditors ...RequestEditorFn) (*struct {
	Status *struct {
		ChunksAvailable      *[]int32                "json:\"chunks_available,omitempty\""
		ChunksAvailableCount *int                    "json:\"chunks_available_count,omitempty\""
		ChunksCancelled      *[]int32                "json:\"chunks_cancelled,omitempty\""
		ChunksFailed         *[]int32                "json:\"chunks_failed,omitempty\""
		Created              *int                    "json:\"created,omitempty\""
		EmptyChunksCount     *int                    "json:\"empty_chunks_count,omitempty\""
		Filters              *map[string]interface{} "json:\"filters,omitempty\""
		FinishedChunks       *int                    "json:\"finished_chunks,omitempty\""
		NumAssetsPerChunk    *int                    "json:\"num_assets_per_chunk,omitempty\""
		Status               *string                 "json:\"status,omitempty\""
		TotalChunks          *int                    "json:\"total_chunks,omitempty\""
		Uuid                 *string                 "json:\"uuid,omitempty\""
	} "json:\"status,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExportsVulnsExportStatusWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Status *struct { ChunksAvailable *[]int32 "json:\"chunks_available,omitempty\""; ChunksAvailableCount *int "json:\"chunks_available_count,omitempty\""; ChunksCancelled *[]int32 "json:\"chunks_cancelled,omitempty\""; ChunksFailed *[]int32 "json:\"chunks_failed,omitempty\""; Created *int "json:\"created,omitempty\""; EmptyChunksCount *int "json:\"empty_chunks_count,omitempty\""; Filters *map[string]interface {} "json:\"filters,omitempty\""; FinishedChunks *int "json:\"finished_chunks,omitempty\""; NumAssetsPerChunk *int "json:\"num_assets_per_chunk,omitempty\""; Status *string "json:\"status,omitempty\""; TotalChunks *int "json:\"total_chunks,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"status,omitempty\"" }
	if i, ok := r.(*struct {
		Status *struct {
			ChunksAvailable      *[]int32                "json:\"chunks_available,omitempty\""
			ChunksAvailableCount *int                    "json:\"chunks_available_count,omitempty\""
			ChunksCancelled      *[]int32                "json:\"chunks_cancelled,omitempty\""
			ChunksFailed         *[]int32                "json:\"chunks_failed,omitempty\""
			Created              *int                    "json:\"created,omitempty\""
			EmptyChunksCount     *int                    "json:\"empty_chunks_count,omitempty\""
			Filters              *map[string]interface{} "json:\"filters,omitempty\""
			FinishedChunks       *int                    "json:\"finished_chunks,omitempty\""
			NumAssetsPerChunk    *int                    "json:\"num_assets_per_chunk,omitempty\""
			Status               *string                 "json:\"status,omitempty\""
			TotalChunks          *int                    "json:\"total_chunks,omitempty\""
			Uuid                 *string                 "json:\"uuid,omitempty\""
		} "json:\"status,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RequestExportWithBody calls the ExportsVulns ´s function
func (c *ExportsVulns) RequestExportWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	ExportUuid *string "json:\"export_uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExportsVulnsRequestExportWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ExportUuid *string "json:\"export_uuid,omitempty\"" }
	if i, ok := r.(*struct {
		ExportUuid *string "json:\"export_uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RequestExport calls the ExportsVulns ´s function
func (c *ExportsVulns) RequestExport(arg1 ExportsVulnsRequestExportJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	ExportUuid *string "json:\"export_uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ExportsVulnsRequestExportWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ExportUuid *string "json:\"export_uuid,omitempty\"" }
	if i, ok := r.(*struct {
		ExportUuid *string "json:\"export_uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Folders implementation of the Folders interface
type Folders struct {
	*ClientWithResponses
}

// CreateWithBody calls the Folders ´s function
func (c *Folders) CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	Id *int "json:\"id,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.FoldersCreateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Id *int "json:\"id,omitempty\"" }
	if i, ok := r.(*struct {
		Id *int "json:\"id,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Create calls the Folders ´s function
func (c *Folders) Create(arg1 FoldersCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	Id *int "json:\"id,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.FoldersCreateWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Id *int "json:\"id,omitempty\"" }
	if i, ok := r.(*struct {
		Id *int "json:\"id,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Delete calls the Folders ´s function
func (c *Folders) Delete(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.FoldersDeleteWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// EditWithBody calls the Folders ´s function
func (c *Folders) EditWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.FoldersEditWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Edit calls the Folders ´s function
func (c *Folders) Edit(arg1 int32, arg2 FoldersEditJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.FoldersEditWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// List calls the Folders ´s function
func (c *Folders) List(reqEditors ...RequestEditorFn) (*struct {
	Folders []struct {
		Custom      *int    "json:\"custom,omitempty\""
		DefaultTag  *int    "json:\"default_tag,omitempty\""
		Id          *int    "json:\"id,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Type        *string "json:\"type,omitempty\""
		UnreadCount *int    "json:\"unread_count,omitempty\""
	} "json:\"folders,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.FoldersListWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { Custom *int "json:\"custom,omitempty\""; DefaultTag *int "json:\"default_tag,omitempty\""; Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Type *string "json:\"type,omitempty\""; UnreadCount *int "json:\"unread_count,omitempty\"" }
	if i, ok := r.(*struct {
		Folders []struct {
			Custom      *int    "json:\"custom,omitempty\""
			DefaultTag  *int    "json:\"default_tag,omitempty\""
			Id          *int    "json:\"id,omitempty\""
			Name        *string "json:\"name,omitempty\""
			Type        *string "json:\"type,omitempty\""
			UnreadCount *int    "json:\"unread_count,omitempty\""
		} "json:\"folders,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// IoAgent implementation of the IoAgent interface
type IoAgent struct {
	*ClientWithResponses
}

// BulkOperationsAddToNetworkWithBody calls the IoAgent ´s function
func (c *IoAgent) BulkOperationsAddToNetworkWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoAgentBulkOperationsAddToNetworkWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// BulkOperationsAddToNetwork calls the IoAgent ´s function
func (c *IoAgent) BulkOperationsAddToNetwork(arg1 IoAgentBulkOperationsAddToNetworkJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoAgentBulkOperationsAddToNetworkWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// BulkOperationsDirectiveWithBody calls the IoAgent ´s function
func (c *IoAgent) BulkOperationsDirectiveWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoAgentBulkOperationsDirectiveWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// BulkOperationsDirective calls the IoAgent ´s function
func (c *IoAgent) BulkOperationsDirective(arg1 IoAgentBulkOperationsDirectiveJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoAgentBulkOperationsDirectiveWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// BulkOperationsGroupDirectiveWithBody calls the IoAgent ´s function
func (c *IoAgent) BulkOperationsGroupDirectiveWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoAgentBulkOperationsGroupDirectiveWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// BulkOperationsGroupDirective calls the IoAgent ´s function
func (c *IoAgent) BulkOperationsGroupDirective(arg1 int32, arg2 IoAgentBulkOperationsGroupDirectiveJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoAgentBulkOperationsGroupDirectiveWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// BulkOperationsRemoveFromNetworkWithBody calls the IoAgent ´s function
func (c *IoAgent) BulkOperationsRemoveFromNetworkWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoAgentBulkOperationsRemoveFromNetworkWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// BulkOperationsRemoveFromNetwork calls the IoAgent ´s function
func (c *IoAgent) BulkOperationsRemoveFromNetwork(arg1 IoAgentBulkOperationsRemoveFromNetworkJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoAgentBulkOperationsRemoveFromNetworkWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Scanners implementation of the Scanners interface
type Scanners struct {
	*ClientWithResponses
}

// ControlScansWithBody calls the Scanners ´s function
func (c *Scanners) ControlScansWithBody(arg1 int32, arg2 string, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannersControlScansWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ControlScans calls the Scanners ´s function
func (c *Scanners) ControlScans(arg1 int32, arg2 string, arg3 ScannersControlScansJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannersControlScansWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Delete calls the Scanners ´s function
func (c *Scanners) Delete(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannersDeleteWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Details calls the Scanners ´s function
func (c *Scanners) Details(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
	AwsUpdateInterval    *int      "json:\"aws_update_interval,omitempty\""
	CreationDate         *int      "json:\"creation_date,omitempty\""
	Distro               *string   "json:\"distro,omitempty\""
	EngineVersion        *string   "json:\"engine_version,omitempty\""
	Group                *bool     "json:\"group,omitempty\""
	Hostname             *string   "json:\"hostname,omitempty\""
	Id                   *int      "json:\"id,omitempty\""
	IpAddresses          *[]string "json:\"ip_addresses,omitempty\""
	Key                  *string   "json:\"key,omitempty\""
	LastConnect          *string   "json:\"last_connect,omitempty\""
	LastModificationDate *int      "json:\"last_modification_date,omitempty\""
	License              *struct {
		Agents *int "json:\"agents,omitempty\""
		Apps   *struct {
			Consec *struct {
				ActivationCode *int    "json:\"activation_code,omitempty\""
				ExpirationDate *int    "json:\"expiration_date,omitempty\""
				MaxGb          *int    "json:\"max_gb,omitempty\""
				Mode           *string "json:\"mode,omitempty\""
				Type           *string "json:\"type,omitempty\""
			} "json:\"consec,omitempty\""
			Type *string "json:\"type,omitempty\""
			Was  *struct {
				ActivationCode *int    "json:\"activation_code,omitempty\""
				ExpirationDate *int    "json:\"expiration_date,omitempty\""
				Mode           *string "json:\"mode,omitempty\""
				Type           *string "json:\"type,omitempty\""
				WebAssets      *int    "json:\"web_assets,omitempty\""
			} "json:\"was,omitempty\""
		} "json:\"apps,omitempty\""
		Ips      *int    "json:\"ips,omitempty\""
		Scanners *int    "json:\"scanners,omitempty\""
		Type     *string "json:\"type,omitempty\""
	} "json:\"license,omitempty\""
	Linked             *int    "json:\"linked,omitempty\""
	LoadedPluginSet    *string "json:\"loaded_plugin_set,omitempty\""
	Name               *string "json:\"name,omitempty\""
	NetworkName        *string "json:\"network_name,omitempty\""
	NumHosts           *int    "json:\"num_hosts,omitempty\""
	NumScans           *int    "json:\"num_scans,omitempty\""
	NumSessions        *int    "json:\"num_sessions,omitempty\""
	NumTcpSessions     *int    "json:\"num_tcp_sessions,omitempty\""
	Owner              *string "json:\"owner,omitempty\""
	OwnerId            *int    "json:\"owner_id,omitempty\""
	OwnerName          *string "json:\"owner_name,omitempty\""
	OwnerUuid          *string "json:\"owner_uuid,omitempty\""
	Platform           *string "json:\"platform,omitempty\""
	Pool               *bool   "json:\"pool,omitempty\""
	RegistrationCode   *string "json:\"registration_code,omitempty\""
	ScanCount          *int    "json:\"scan_count,omitempty\""
	Shared             *bool   "json:\"shared,omitempty\""
	Source             *string "json:\"source,omitempty\""
	Status             *string "json:\"status,omitempty\""
	SupportsRemoteLogs *bool   "json:\"supports_remote_logs,omitempty\""
	SupportsWebapp     *bool   "json:\"supports_webapp,omitempty\""
	Timestamp          *int    "json:\"timestamp,omitempty\""
	Type               *string "json:\"type,omitempty\""
	UiBuild            *string "json:\"ui_build,omitempty\""
	UiVersion          *string "json:\"ui_version,omitempty\""
	UserPermissions    *int    "json:\"user_permissions,omitempty\""
	Uuid               *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannersDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AwsUpdateInterval *int "json:\"aws_update_interval,omitempty\""; CreationDate *int "json:\"creation_date,omitempty\""; Distro *string "json:\"distro,omitempty\""; EngineVersion *string "json:\"engine_version,omitempty\""; Group *bool "json:\"group,omitempty\""; Hostname *string "json:\"hostname,omitempty\""; Id *int "json:\"id,omitempty\""; IpAddresses *[]string "json:\"ip_addresses,omitempty\""; Key *string "json:\"key,omitempty\""; LastConnect *string "json:\"last_connect,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; License *struct { Agents *int "json:\"agents,omitempty\""; Apps *struct { Consec *struct { ActivationCode *int "json:\"activation_code,omitempty\""; ExpirationDate *int "json:\"expiration_date,omitempty\""; MaxGb *int "json:\"max_gb,omitempty\""; Mode *string "json:\"mode,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"consec,omitempty\""; Type *string "json:\"type,omitempty\""; Was *struct { ActivationCode *int "json:\"activation_code,omitempty\""; ExpirationDate *int "json:\"expiration_date,omitempty\""; Mode *string "json:\"mode,omitempty\""; Type *string "json:\"type,omitempty\""; WebAssets *int "json:\"web_assets,omitempty\"" } "json:\"was,omitempty\"" } "json:\"apps,omitempty\""; Ips *int "json:\"ips,omitempty\""; Scanners *int "json:\"scanners,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"license,omitempty\""; Linked *int "json:\"linked,omitempty\""; LoadedPluginSet *string "json:\"loaded_plugin_set,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkName *string "json:\"network_name,omitempty\""; NumHosts *int "json:\"num_hosts,omitempty\""; NumScans *int "json:\"num_scans,omitempty\""; NumSessions *int "json:\"num_sessions,omitempty\""; NumTcpSessions *int "json:\"num_tcp_sessions,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; OwnerName *string "json:\"owner_name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Platform *string "json:\"platform,omitempty\""; Pool *bool "json:\"pool,omitempty\""; RegistrationCode *string "json:\"registration_code,omitempty\""; ScanCount *int "json:\"scan_count,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Source *string "json:\"source,omitempty\""; Status *string "json:\"status,omitempty\""; SupportsRemoteLogs *bool "json:\"supports_remote_logs,omitempty\""; SupportsWebapp *bool "json:\"supports_webapp,omitempty\""; Timestamp *int "json:\"timestamp,omitempty\""; Type *string "json:\"type,omitempty\""; UiBuild *int "json:\"ui_build,omitempty\""; UiVersion *string "json:\"ui_version,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		AwsUpdateInterval    *int      "json:\"aws_update_interval,omitempty\""
		CreationDate         *int      "json:\"creation_date,omitempty\""
		Distro               *string   "json:\"distro,omitempty\""
		EngineVersion        *string   "json:\"engine_version,omitempty\""
		Group                *bool     "json:\"group,omitempty\""
		Hostname             *string   "json:\"hostname,omitempty\""
		Id                   *int      "json:\"id,omitempty\""
		IpAddresses          *[]string "json:\"ip_addresses,omitempty\""
		Key                  *string   "json:\"key,omitempty\""
		LastConnect          *string   "json:\"last_connect,omitempty\""
		LastModificationDate *int      "json:\"last_modification_date,omitempty\""
		License              *struct {
			Agents *int "json:\"agents,omitempty\""
			Apps   *struct {
				Consec *struct {
					ActivationCode *int    "json:\"activation_code,omitempty\""
					ExpirationDate *int    "json:\"expiration_date,omitempty\""
					MaxGb          *int    "json:\"max_gb,omitempty\""
					Mode           *string "json:\"mode,omitempty\""
					Type           *string "json:\"type,omitempty\""
				} "json:\"consec,omitempty\""
				Type *string "json:\"type,omitempty\""
				Was  *struct {
					ActivationCode *int    "json:\"activation_code,omitempty\""
					ExpirationDate *int    "json:\"expiration_date,omitempty\""
					Mode           *string "json:\"mode,omitempty\""
					Type           *string "json:\"type,omitempty\""
					WebAssets      *int    "json:\"web_assets,omitempty\""
				} "json:\"was,omitempty\""
			} "json:\"apps,omitempty\""
			Ips      *int    "json:\"ips,omitempty\""
			Scanners *int    "json:\"scanners,omitempty\""
			Type     *string "json:\"type,omitempty\""
		} "json:\"license,omitempty\""
		Linked             *int    "json:\"linked,omitempty\""
		LoadedPluginSet    *string "json:\"loaded_plugin_set,omitempty\""
		Name               *string "json:\"name,omitempty\""
		NetworkName        *string "json:\"network_name,omitempty\""
		NumHosts           *int    "json:\"num_hosts,omitempty\""
		NumScans           *int    "json:\"num_scans,omitempty\""
		NumSessions        *int    "json:\"num_sessions,omitempty\""
		NumTcpSessions     *int    "json:\"num_tcp_sessions,omitempty\""
		Owner              *string "json:\"owner,omitempty\""
		OwnerId            *int    "json:\"owner_id,omitempty\""
		OwnerName          *string "json:\"owner_name,omitempty\""
		OwnerUuid          *string "json:\"owner_uuid,omitempty\""
		Platform           *string "json:\"platform,omitempty\""
		Pool               *bool   "json:\"pool,omitempty\""
		RegistrationCode   *string "json:\"registration_code,omitempty\""
		ScanCount          *int    "json:\"scan_count,omitempty\""
		Shared             *bool   "json:\"shared,omitempty\""
		Source             *string "json:\"source,omitempty\""
		Status             *string "json:\"status,omitempty\""
		SupportsRemoteLogs *bool   "json:\"supports_remote_logs,omitempty\""
		SupportsWebapp     *bool   "json:\"supports_webapp,omitempty\""
		Timestamp          *int    "json:\"timestamp,omitempty\""
		Type               *string "json:\"type,omitempty\""
		UiBuild            *string "json:\"ui_build,omitempty\""
		UiVersion          *string "json:\"ui_version,omitempty\""
		UserPermissions    *int    "json:\"user_permissions,omitempty\""
		Uuid               *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// EditWithBody calls the Scanners ´s function
func (c *Scanners) EditWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	ContainerUuid      *string "json:\"container_uuid,omitempty\""
	Created            *int    "json:\"created,omitempty\""
	DefaultPermissions *int    "json:\"default_permissions,omitempty\""
	Id                 *int    "json:\"id,omitempty\""
	Key                *string "json:\"key,omitempty\""
	Modified           *int    "json:\"modified,omitempty\""
	Name               *string "json:\"name,omitempty\""
	NetworkName        *string "json:\"network_name,omitempty\""
	OwnerUuid          *string "json:\"owner_uuid,omitempty\""
	Shared             *bool   "json:\"shared,omitempty\""
	Status             *string "json:\"status,omitempty\""
	UserPermissions    *int    "json:\"user_permissions,omitempty\""
	Uuid               *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannersEditWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ContainerUuid *string "json:\"container_uuid,omitempty\""; Created *int "json:\"created,omitempty\""; DefaultPermissions *int "json:\"default_permissions,omitempty\""; Id *int "json:\"id,omitempty\""; Key *string "json:\"key,omitempty\""; Modified *int "json:\"modified,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkName *string "json:\"network_name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Status *string "json:\"status,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		ContainerUuid      *string "json:\"container_uuid,omitempty\""
		Created            *int    "json:\"created,omitempty\""
		DefaultPermissions *int    "json:\"default_permissions,omitempty\""
		Id                 *int    "json:\"id,omitempty\""
		Key                *string "json:\"key,omitempty\""
		Modified           *int    "json:\"modified,omitempty\""
		Name               *string "json:\"name,omitempty\""
		NetworkName        *string "json:\"network_name,omitempty\""
		OwnerUuid          *string "json:\"owner_uuid,omitempty\""
		Shared             *bool   "json:\"shared,omitempty\""
		Status             *string "json:\"status,omitempty\""
		UserPermissions    *int    "json:\"user_permissions,omitempty\""
		Uuid               *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Edit calls the Scanners ´s function
func (c *Scanners) Edit(arg1 int32, arg2 ScannersEditJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	ContainerUuid      *string "json:\"container_uuid,omitempty\""
	Created            *int    "json:\"created,omitempty\""
	DefaultPermissions *int    "json:\"default_permissions,omitempty\""
	Id                 *int    "json:\"id,omitempty\""
	Key                *string "json:\"key,omitempty\""
	Modified           *int    "json:\"modified,omitempty\""
	Name               *string "json:\"name,omitempty\""
	NetworkName        *string "json:\"network_name,omitempty\""
	OwnerUuid          *string "json:\"owner_uuid,omitempty\""
	Shared             *bool   "json:\"shared,omitempty\""
	Status             *string "json:\"status,omitempty\""
	UserPermissions    *int    "json:\"user_permissions,omitempty\""
	Uuid               *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannersEditWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ContainerUuid *string "json:\"container_uuid,omitempty\""; Created *int "json:\"created,omitempty\""; DefaultPermissions *int "json:\"default_permissions,omitempty\""; Id *int "json:\"id,omitempty\""; Key *string "json:\"key,omitempty\""; Modified *int "json:\"modified,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkName *string "json:\"network_name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Status *string "json:\"status,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		ContainerUuid      *string "json:\"container_uuid,omitempty\""
		Created            *int    "json:\"created,omitempty\""
		DefaultPermissions *int    "json:\"default_permissions,omitempty\""
		Id                 *int    "json:\"id,omitempty\""
		Key                *string "json:\"key,omitempty\""
		Modified           *int    "json:\"modified,omitempty\""
		Name               *string "json:\"name,omitempty\""
		NetworkName        *string "json:\"network_name,omitempty\""
		OwnerUuid          *string "json:\"owner_uuid,omitempty\""
		Shared             *bool   "json:\"shared,omitempty\""
		Status             *string "json:\"status,omitempty\""
		UserPermissions    *int    "json:\"user_permissions,omitempty\""
		Uuid               *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetAwsTargets calls the Scanners ´s function
func (c *Scanners) GetAwsTargets(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
	InstanceId *string "json:\"instance_id,omitempty\""
	Name       *string "json:\"name,omitempty\""
	PrivateIp  *string "json:\"private_ip,omitempty\""
	PublicIp   *string "json:\"public_ip,omitempty\""
	ScannerId  *int    "json:\"scanner_id,omitempty\""
	State      *string "json:\"state,omitempty\""
	Type       *string "json:\"type,omitempty\""
	Zone       *string "json:\"zone,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannersGetAwsTargetsWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { InstanceId *string "json:\"instance_id,omitempty\""; Name *string "json:\"name,omitempty\""; PrivateIp *string "json:\"private_ip,omitempty\""; PublicIp *string "json:\"public_ip,omitempty\""; ScannerId *int "json:\"scanner_id,omitempty\""; State *string "json:\"state,omitempty\""; Type *string "json:\"type,omitempty\""; Zone *string "json:\"zone,omitempty\"" }
	if i, ok := r.(*struct {
		InstanceId *string "json:\"instance_id,omitempty\""
		Name       *string "json:\"name,omitempty\""
		PrivateIp  *string "json:\"private_ip,omitempty\""
		PublicIp   *string "json:\"public_ip,omitempty\""
		ScannerId  *int    "json:\"scanner_id,omitempty\""
		State      *string "json:\"state,omitempty\""
		Type       *string "json:\"type,omitempty\""
		Zone       *string "json:\"zone,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetScannerKey calls the Scanners ´s function
func (c *Scanners) GetScannerKey(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
	ScannerId *string "json:\"scanner_id,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannersGetScannerKeyWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ScannerId *string "json:\"scanner_id,omitempty\"" }
	if i, ok := r.(*struct {
		ScannerId *string "json:\"scanner_id,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetScans calls the Scanners ´s function
func (c *Scanners) GetScans(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
	Id                   *string "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NetworkId            *string "json:\"network_id,omitempty\""
	Remote               *bool   "json:\"remote,omitempty\""
	ScanId               *int    "json:\"scan_id,omitempty\""
	ScannerUuid          *string "json:\"scanner_uuid,omitempty\""
	StartTime            *int    "json:\"start_time,omitempty\""
	Status               *string "json:\"status,omitempty\""
	User                 *string "json:\"user,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannersGetScansWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Id *string "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkId *string "json:\"network_id,omitempty\""; Remote *bool "json:\"remote,omitempty\""; ScanId *int "json:\"scan_id,omitempty\""; ScannerUuid *string "json:\"scanner_uuid,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *string "json:\"status,omitempty\""; User *string "json:\"user,omitempty\"" }
	if i, ok := r.(*struct {
		Id                   *string "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NetworkId            *string "json:\"network_id,omitempty\""
		Remote               *bool   "json:\"remote,omitempty\""
		ScanId               *int    "json:\"scan_id,omitempty\""
		ScannerUuid          *string "json:\"scanner_uuid,omitempty\""
		StartTime            *int    "json:\"start_time,omitempty\""
		Status               *string "json:\"status,omitempty\""
		User                 *string "json:\"user,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// List calls the Scanners ´s function
func (c *Scanners) List(reqEditors ...RequestEditorFn) (*struct {
	Scanners []struct {
		Challenge            *string   `json:"challenge,omitempty"`
		AwsUpdateInterval    *int      "json:\"aws_update_interval,omitempty\""
		CreationDate         *int      "json:\"creation_date,omitempty\""
		Distro               *string   "json:\"distro,omitempty\""
		EngineVersion        *string   "json:\"engine_version,omitempty\""
		Group                *bool     "json:\"group,omitempty\""
		Hostname             *string   "json:\"hostname,omitempty\""
		Id                   *int      "json:\"id,omitempty\""
		IpAddresses          *[]string "json:\"ip_addresses,omitempty\""
		Key                  *string   "json:\"key,omitempty\""
		LastConnect          *string   "json:\"last_connect,omitempty\""
		LastModificationDate *int      "json:\"last_modification_date,omitempty\""
		License              *struct {
			Agents *int "json:\"agents,omitempty\""
			Apps   *struct {
				Consec *struct {
					ActivationCode *int    "json:\"activation_code,omitempty\""
					ExpirationDate *int    "json:\"expiration_date,omitempty\""
					MaxGb          *int    "json:\"max_gb,omitempty\""
					Mode           *string "json:\"mode,omitempty\""
					Type           *string "json:\"type,omitempty\""
				} "json:\"consec,omitempty\""
				Type *string "json:\"type,omitempty\""
				Was  *struct {
					ActivationCode *int    "json:\"activation_code,omitempty\""
					ExpirationDate *int    "json:\"expiration_date,omitempty\""
					Mode           *string "json:\"mode,omitempty\""
					Type           *string "json:\"type,omitempty\""
					WebAssets      *int    "json:\"web_assets,omitempty\""
				} "json:\"was,omitempty\""
			} "json:\"apps,omitempty\""
			Ips      *int    "json:\"ips,omitempty\""
			Scanners *int    "json:\"scanners,omitempty\""
			Type     *string "json:\"type,omitempty\""
		} "json:\"license,omitempty\""
		Linked             *int    "json:\"linked,omitempty\""
		LoadedPluginSet    *string "json:\"loaded_plugin_set,omitempty\""
		Name               *string "json:\"name,omitempty\""
		NetworkName        *string "json:\"network_name,omitempty\""
		NumHosts           *int    "json:\"num_hosts,omitempty\""
		NumScans           *int    "json:\"num_scans,omitempty\""
		NumSessions        *int    "json:\"num_sessions,omitempty\""
		NumTcpSessions     *int    "json:\"num_tcp_sessions,omitempty\""
		Owner              *string "json:\"owner,omitempty\""
		OwnerId            *int    "json:\"owner_id,omitempty\""
		OwnerName          *string "json:\"owner_name,omitempty\""
		OwnerUuid          *string "json:\"owner_uuid,omitempty\""
		Platform           *string "json:\"platform,omitempty\""
		Pool               *bool   "json:\"pool,omitempty\""
		RegistrationCode   *string "json:\"registration_code,omitempty\""
		ScanCount          *int    "json:\"scan_count,omitempty\""
		Shared             *int    "json:\"shared,omitempty\""
		Source             *string "json:\"source,omitempty\""
		Status             *string "json:\"status,omitempty\""
		SupportsRemoteLogs *bool   "json:\"supports_remote_logs,omitempty\""
		SupportsWebapp     *bool   "json:\"supports_webapp,omitempty\""
		Timestamp          *int    "json:\"timestamp,omitempty\""
		Type               *string "json:\"type,omitempty\""
		UiBuild            *string "json:\"ui_build,omitempty\""
		UiVersion          *string "json:\"ui_version,omitempty\""
		UserPermissions    *int    "json:\"user_permissions,omitempty\""
		Uuid               *string "json:\"uuid,omitempty\""
	} "json:\"scanners,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannersListWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { AwsUpdateInterval *int "json:\"aws_update_interval,omitempty\""; CreationDate *int "json:\"creation_date,omitempty\""; Distro *string "json:\"distro,omitempty\""; EngineVersion *string "json:\"engine_version,omitempty\""; Group *bool "json:\"group,omitempty\""; Hostname *string "json:\"hostname,omitempty\""; Id *int "json:\"id,omitempty\""; IpAddresses *[]string "json:\"ip_addresses,omitempty\""; Key *string "json:\"key,omitempty\""; LastConnect *string "json:\"last_connect,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; License *struct { Agents *int "json:\"agents,omitempty\""; Apps *struct { Consec *struct { ActivationCode *int "json:\"activation_code,omitempty\""; ExpirationDate *int "json:\"expiration_date,omitempty\""; MaxGb *int "json:\"max_gb,omitempty\""; Mode *string "json:\"mode,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"consec,omitempty\""; Type *string "json:\"type,omitempty\""; Was *struct { ActivationCode *int "json:\"activation_code,omitempty\""; ExpirationDate *int "json:\"expiration_date,omitempty\""; Mode *string "json:\"mode,omitempty\""; Type *string "json:\"type,omitempty\""; WebAssets *int "json:\"web_assets,omitempty\"" } "json:\"was,omitempty\"" } "json:\"apps,omitempty\""; Ips *int "json:\"ips,omitempty\""; Scanners *int "json:\"scanners,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"license,omitempty\""; Linked *int "json:\"linked,omitempty\""; LoadedPluginSet *string "json:\"loaded_plugin_set,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkName *string "json:\"network_name,omitempty\""; NumHosts *int "json:\"num_hosts,omitempty\""; NumScans *int "json:\"num_scans,omitempty\""; NumSessions *int "json:\"num_sessions,omitempty\""; NumTcpSessions *int "json:\"num_tcp_sessions,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; OwnerName *string "json:\"owner_name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Platform *string "json:\"platform,omitempty\""; Pool *bool "json:\"pool,omitempty\""; RegistrationCode *string "json:\"registration_code,omitempty\""; ScanCount *int "json:\"scan_count,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Source *string "json:\"source,omitempty\""; Status *string "json:\"status,omitempty\""; SupportsRemoteLogs *bool "json:\"supports_remote_logs,omitempty\""; SupportsWebapp *bool "json:\"supports_webapp,omitempty\""; Timestamp *int "json:\"timestamp,omitempty\""; Type *string "json:\"type,omitempty\""; UiBuild *int "json:\"ui_build,omitempty\""; UiVersion *string "json:\"ui_version,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		Scanners []struct {
			Challenge            *string   `json:"challenge,omitempty"`
			AwsUpdateInterval    *int      "json:\"aws_update_interval,omitempty\""
			CreationDate         *int      "json:\"creation_date,omitempty\""
			Distro               *string   "json:\"distro,omitempty\""
			EngineVersion        *string   "json:\"engine_version,omitempty\""
			Group                *bool     "json:\"group,omitempty\""
			Hostname             *string   "json:\"hostname,omitempty\""
			Id                   *int      "json:\"id,omitempty\""
			IpAddresses          *[]string "json:\"ip_addresses,omitempty\""
			Key                  *string   "json:\"key,omitempty\""
			LastConnect          *string   "json:\"last_connect,omitempty\""
			LastModificationDate *int      "json:\"last_modification_date,omitempty\""
			License              *struct {
				Agents *int "json:\"agents,omitempty\""
				Apps   *struct {
					Consec *struct {
						ActivationCode *int    "json:\"activation_code,omitempty\""
						ExpirationDate *int    "json:\"expiration_date,omitempty\""
						MaxGb          *int    "json:\"max_gb,omitempty\""
						Mode           *string "json:\"mode,omitempty\""
						Type           *string "json:\"type,omitempty\""
					} "json:\"consec,omitempty\""
					Type *string "json:\"type,omitempty\""
					Was  *struct {
						ActivationCode *int    "json:\"activation_code,omitempty\""
						ExpirationDate *int    "json:\"expiration_date,omitempty\""
						Mode           *string "json:\"mode,omitempty\""
						Type           *string "json:\"type,omitempty\""
						WebAssets      *int    "json:\"web_assets,omitempty\""
					} "json:\"was,omitempty\""
				} "json:\"apps,omitempty\""
				Ips      *int    "json:\"ips,omitempty\""
				Scanners *int    "json:\"scanners,omitempty\""
				Type     *string "json:\"type,omitempty\""
			} "json:\"license,omitempty\""
			Linked             *int    "json:\"linked,omitempty\""
			LoadedPluginSet    *string "json:\"loaded_plugin_set,omitempty\""
			Name               *string "json:\"name,omitempty\""
			NetworkName        *string "json:\"network_name,omitempty\""
			NumHosts           *int    "json:\"num_hosts,omitempty\""
			NumScans           *int    "json:\"num_scans,omitempty\""
			NumSessions        *int    "json:\"num_sessions,omitempty\""
			NumTcpSessions     *int    "json:\"num_tcp_sessions,omitempty\""
			Owner              *string "json:\"owner,omitempty\""
			OwnerId            *int    "json:\"owner_id,omitempty\""
			OwnerName          *string "json:\"owner_name,omitempty\""
			OwnerUuid          *string "json:\"owner_uuid,omitempty\""
			Platform           *string "json:\"platform,omitempty\""
			Pool               *bool   "json:\"pool,omitempty\""
			RegistrationCode   *string "json:\"registration_code,omitempty\""
			ScanCount          *int    "json:\"scan_count,omitempty\""
			Shared             *int    "json:\"shared,omitempty\""
			Source             *string "json:\"source,omitempty\""
			Status             *string "json:\"status,omitempty\""
			SupportsRemoteLogs *bool   "json:\"supports_remote_logs,omitempty\""
			SupportsWebapp     *bool   "json:\"supports_webapp,omitempty\""
			Timestamp          *int    "json:\"timestamp,omitempty\""
			Type               *string "json:\"type,omitempty\""
			UiBuild            *string "json:\"ui_build,omitempty\""
			UiVersion          *string "json:\"ui_version,omitempty\""
			UserPermissions    *int    "json:\"user_permissions,omitempty\""
			Uuid               *string "json:\"uuid,omitempty\""
		} "json:\"scanners,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ToggleLinkStateWithBody calls the Scanners ´s function
func (c *Scanners) ToggleLinkStateWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannersToggleLinkStateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ToggleLinkState calls the Scanners ´s function
func (c *Scanners) ToggleLinkState(arg1 int32, arg2 ScannersToggleLinkStateJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannersToggleLinkStateWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Scans implementation of the Scans interface
type Scans struct {
	*ClientWithResponses
}

// Attachments calls the Scans ´s function
func (c *Scans) Attachments(arg1 string, arg2 string, params *ScansAttachmentsParams, reqEditors ...RequestEditorFn) (*ScansAttachmentsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansAttachmentsWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *ScansAttachmentsResponse
	if i, ok := r.(*ScansAttachmentsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ConfigureWithBody calls the Scans ´s function
func (c *Scans) ConfigureWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	ContainerId          *string "json:\"container_id,omitempty\""
	CreationDate         *int32  "json:\"creation_date,omitempty\""
	CustomTargets        *string "json:\"custom_targets,omitempty\""
	DashboardFile        *string "json:\"dashboard_file,omitempty\""
	DefaultPermissions   *int32  "json:\"default_permissions,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Emails               *string "json:\"emails,omitempty\""
	Enabled              *bool   "json:\"enabled,omitempty\""
	Id                   *int32  "json:\"id,omitempty\""
	IncludeAggregate     *bool   "json:\"include_aggregate,omitempty\""
	LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NotificationFilters  *[]struct {
		Filter  *string "json:\"filter,omitempty\""
		Quality *string "json:\"quality,omitempty\""
		Value   *string "json:\"value,omitempty\""
	} "json:\"notification_filters,omitempty\""
	Owner             *string   "json:\"owner,omitempty\""
	OwnerId           *int      "json:\"owner_id,omitempty\""
	OwnerUuid         *string   "json:\"owner_uuid,omitempty\""
	PolicyId          *int      "json:\"policy_id,omitempty\""
	Remediation       *int      "json:\"remediation,omitempty\""
	Rrules            *string   "json:\"rrules,omitempty\""
	ScanTimeWindow    *string   "json:\"scan_time_window,omitempty\""
	ScannerUuid       *string   "json:\"scanner_uuid,omitempty\""
	Shared            *bool     "json:\"shared,omitempty\""
	Sms               *string   "json:\"sms,omitempty\""
	Starttime         *string   "json:\"starttime,omitempty\""
	TagTargets        *[]string "json:\"tag_targets,omitempty\""
	TagType           *string   "json:\"tag_type,omitempty\""
	TargetNetworkUuid *string   "json:\"target_network_uuid,omitempty\""
	Timezone          *string   "json:\"timezone,omitempty\""
	Type              *string   "json:\"type,omitempty\""
	UserPermissions   *int32    "json:\"user_permissions,omitempty\""
	Uuid              *string   "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansConfigureWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ContainerId *string "json:\"container_id,omitempty\""; CreationDate *int32 "json:\"creation_date,omitempty\""; CustomTargets *string "json:\"custom_targets,omitempty\""; DashboardFile *string "json:\"dashboard_file,omitempty\""; DefaultPermissions *int32 "json:\"default_permissions,omitempty\""; Description *string "json:\"description,omitempty\""; Emails *string "json:\"emails,omitempty\""; Enabled *bool "json:\"enabled,omitempty\""; Id *int32 "json:\"id,omitempty\""; IncludeAggregate *bool "json:\"include_aggregate,omitempty\""; LastModificationDate *int32 "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NotificationFilters *[]struct { Filter *string "json:\"filter,omitempty\""; Quality *string "json:\"quality,omitempty\""; Value *string "json:\"value,omitempty\"" } "json:\"notification_filters,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; PolicyId *int "json:\"policy_id,omitempty\""; Remediation *int "json:\"remediation,omitempty\""; Rrules *string "json:\"rrules,omitempty\""; ScanTimeWindow *string "json:\"scan_time_window,omitempty\""; ScannerUuid *string "json:\"scanner_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Sms *string "json:\"sms,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; TagTargets *[]string "json:\"tag_targets,omitempty\""; TagType *string "json:\"tag_type,omitempty\""; TargetNetworkUuid *string "json:\"target_network_uuid,omitempty\""; Timezone *string "json:\"timezone,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int32 "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		ContainerId          *string "json:\"container_id,omitempty\""
		CreationDate         *int32  "json:\"creation_date,omitempty\""
		CustomTargets        *string "json:\"custom_targets,omitempty\""
		DashboardFile        *string "json:\"dashboard_file,omitempty\""
		DefaultPermissions   *int32  "json:\"default_permissions,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Emails               *string "json:\"emails,omitempty\""
		Enabled              *bool   "json:\"enabled,omitempty\""
		Id                   *int32  "json:\"id,omitempty\""
		IncludeAggregate     *bool   "json:\"include_aggregate,omitempty\""
		LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NotificationFilters  *[]struct {
			Filter  *string "json:\"filter,omitempty\""
			Quality *string "json:\"quality,omitempty\""
			Value   *string "json:\"value,omitempty\""
		} "json:\"notification_filters,omitempty\""
		Owner             *string   "json:\"owner,omitempty\""
		OwnerId           *int      "json:\"owner_id,omitempty\""
		OwnerUuid         *string   "json:\"owner_uuid,omitempty\""
		PolicyId          *int      "json:\"policy_id,omitempty\""
		Remediation       *int      "json:\"remediation,omitempty\""
		Rrules            *string   "json:\"rrules,omitempty\""
		ScanTimeWindow    *string   "json:\"scan_time_window,omitempty\""
		ScannerUuid       *string   "json:\"scanner_uuid,omitempty\""
		Shared            *bool     "json:\"shared,omitempty\""
		Sms               *string   "json:\"sms,omitempty\""
		Starttime         *string   "json:\"starttime,omitempty\""
		TagTargets        *[]string "json:\"tag_targets,omitempty\""
		TagType           *string   "json:\"tag_type,omitempty\""
		TargetNetworkUuid *string   "json:\"target_network_uuid,omitempty\""
		Timezone          *string   "json:\"timezone,omitempty\""
		Type              *string   "json:\"type,omitempty\""
		UserPermissions   *int32    "json:\"user_permissions,omitempty\""
		Uuid              *string   "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Configure calls the Scans ´s function
func (c *Scans) Configure(arg1 string, arg2 ScansConfigureJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	ContainerId          *string "json:\"container_id,omitempty\""
	CreationDate         *int32  "json:\"creation_date,omitempty\""
	CustomTargets        *string "json:\"custom_targets,omitempty\""
	DashboardFile        *string "json:\"dashboard_file,omitempty\""
	DefaultPermissions   *int32  "json:\"default_permissions,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Emails               *string "json:\"emails,omitempty\""
	Enabled              *bool   "json:\"enabled,omitempty\""
	Id                   *int32  "json:\"id,omitempty\""
	IncludeAggregate     *bool   "json:\"include_aggregate,omitempty\""
	LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NotificationFilters  *[]struct {
		Filter  *string "json:\"filter,omitempty\""
		Quality *string "json:\"quality,omitempty\""
		Value   *string "json:\"value,omitempty\""
	} "json:\"notification_filters,omitempty\""
	Owner             *string   "json:\"owner,omitempty\""
	OwnerId           *int      "json:\"owner_id,omitempty\""
	OwnerUuid         *string   "json:\"owner_uuid,omitempty\""
	PolicyId          *int      "json:\"policy_id,omitempty\""
	Remediation       *int      "json:\"remediation,omitempty\""
	Rrules            *string   "json:\"rrules,omitempty\""
	ScanTimeWindow    *string   "json:\"scan_time_window,omitempty\""
	ScannerUuid       *string   "json:\"scanner_uuid,omitempty\""
	Shared            *bool     "json:\"shared,omitempty\""
	Sms               *string   "json:\"sms,omitempty\""
	Starttime         *string   "json:\"starttime,omitempty\""
	TagTargets        *[]string "json:\"tag_targets,omitempty\""
	TagType           *string   "json:\"tag_type,omitempty\""
	TargetNetworkUuid *string   "json:\"target_network_uuid,omitempty\""
	Timezone          *string   "json:\"timezone,omitempty\""
	Type              *string   "json:\"type,omitempty\""
	UserPermissions   *int32    "json:\"user_permissions,omitempty\""
	Uuid              *string   "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansConfigureWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ContainerId *string "json:\"container_id,omitempty\""; CreationDate *int32 "json:\"creation_date,omitempty\""; CustomTargets *string "json:\"custom_targets,omitempty\""; DashboardFile *string "json:\"dashboard_file,omitempty\""; DefaultPermissions *int32 "json:\"default_permissions,omitempty\""; Description *string "json:\"description,omitempty\""; Emails *string "json:\"emails,omitempty\""; Enabled *bool "json:\"enabled,omitempty\""; Id *int32 "json:\"id,omitempty\""; IncludeAggregate *bool "json:\"include_aggregate,omitempty\""; LastModificationDate *int32 "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NotificationFilters *[]struct { Filter *string "json:\"filter,omitempty\""; Quality *string "json:\"quality,omitempty\""; Value *string "json:\"value,omitempty\"" } "json:\"notification_filters,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; PolicyId *int "json:\"policy_id,omitempty\""; Remediation *int "json:\"remediation,omitempty\""; Rrules *string "json:\"rrules,omitempty\""; ScanTimeWindow *string "json:\"scan_time_window,omitempty\""; ScannerUuid *string "json:\"scanner_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Sms *string "json:\"sms,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; TagTargets *[]string "json:\"tag_targets,omitempty\""; TagType *string "json:\"tag_type,omitempty\""; TargetNetworkUuid *string "json:\"target_network_uuid,omitempty\""; Timezone *string "json:\"timezone,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int32 "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		ContainerId          *string "json:\"container_id,omitempty\""
		CreationDate         *int32  "json:\"creation_date,omitempty\""
		CustomTargets        *string "json:\"custom_targets,omitempty\""
		DashboardFile        *string "json:\"dashboard_file,omitempty\""
		DefaultPermissions   *int32  "json:\"default_permissions,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Emails               *string "json:\"emails,omitempty\""
		Enabled              *bool   "json:\"enabled,omitempty\""
		Id                   *int32  "json:\"id,omitempty\""
		IncludeAggregate     *bool   "json:\"include_aggregate,omitempty\""
		LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NotificationFilters  *[]struct {
			Filter  *string "json:\"filter,omitempty\""
			Quality *string "json:\"quality,omitempty\""
			Value   *string "json:\"value,omitempty\""
		} "json:\"notification_filters,omitempty\""
		Owner             *string   "json:\"owner,omitempty\""
		OwnerId           *int      "json:\"owner_id,omitempty\""
		OwnerUuid         *string   "json:\"owner_uuid,omitempty\""
		PolicyId          *int      "json:\"policy_id,omitempty\""
		Remediation       *int      "json:\"remediation,omitempty\""
		Rrules            *string   "json:\"rrules,omitempty\""
		ScanTimeWindow    *string   "json:\"scan_time_window,omitempty\""
		ScannerUuid       *string   "json:\"scanner_uuid,omitempty\""
		Shared            *bool     "json:\"shared,omitempty\""
		Sms               *string   "json:\"sms,omitempty\""
		Starttime         *string   "json:\"starttime,omitempty\""
		TagTargets        *[]string "json:\"tag_targets,omitempty\""
		TagType           *string   "json:\"tag_type,omitempty\""
		TargetNetworkUuid *string   "json:\"target_network_uuid,omitempty\""
		Timezone          *string   "json:\"timezone,omitempty\""
		Type              *string   "json:\"type,omitempty\""
		UserPermissions   *int32    "json:\"user_permissions,omitempty\""
		Uuid              *string   "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CopyWithBody calls the Scans ´s function
func (c *Scans) CopyWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	Control              *bool   "json:\"control,omitempty\""
	CreationDate         *int    "json:\"creation_date,omitempty\""
	Enabled              *bool   "json:\"enabled,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	Read                 *bool   "json:\"read,omitempty\""
	Rrules               *string "json:\"rrules,omitempty\""
	ScheduleUuid         *string "json:\"schedule_uuid,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Starttime            *string "json:\"starttime,omitempty\""
	Status               *string "json:\"status,omitempty\""
	Timezone             *string "json:\"timezone,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
	Uuid                 *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansCopyWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Control *bool "json:\"control,omitempty\""; CreationDate *int "json:\"creation_date,omitempty\""; Enabled *bool "json:\"enabled,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int32 "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; Read *bool "json:\"read,omitempty\""; Rrules *string "json:\"rrules,omitempty\""; ScheduleUuid *string "json:\"schedule_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; Status *string "json:\"status,omitempty\""; Timezone *string "json:\"timezone,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		Control              *bool   "json:\"control,omitempty\""
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Enabled              *bool   "json:\"enabled,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		Read                 *bool   "json:\"read,omitempty\""
		Rrules               *string "json:\"rrules,omitempty\""
		ScheduleUuid         *string "json:\"schedule_uuid,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Starttime            *string "json:\"starttime,omitempty\""
		Status               *string "json:\"status,omitempty\""
		Timezone             *string "json:\"timezone,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Copy calls the Scans ´s function
func (c *Scans) Copy(arg1 string, arg2 ScansCopyJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	Control              *bool   "json:\"control,omitempty\""
	CreationDate         *int    "json:\"creation_date,omitempty\""
	Enabled              *bool   "json:\"enabled,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	Read                 *bool   "json:\"read,omitempty\""
	Rrules               *string "json:\"rrules,omitempty\""
	ScheduleUuid         *string "json:\"schedule_uuid,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Starttime            *string "json:\"starttime,omitempty\""
	Status               *string "json:\"status,omitempty\""
	Timezone             *string "json:\"timezone,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
	Uuid                 *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansCopyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Control *bool "json:\"control,omitempty\""; CreationDate *int "json:\"creation_date,omitempty\""; Enabled *bool "json:\"enabled,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int32 "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; Read *bool "json:\"read,omitempty\""; Rrules *string "json:\"rrules,omitempty\""; ScheduleUuid *string "json:\"schedule_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; Status *string "json:\"status,omitempty\""; Timezone *string "json:\"timezone,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		Control              *bool   "json:\"control,omitempty\""
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Enabled              *bool   "json:\"enabled,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		Read                 *bool   "json:\"read,omitempty\""
		Rrules               *string "json:\"rrules,omitempty\""
		ScheduleUuid         *string "json:\"schedule_uuid,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Starttime            *string "json:\"starttime,omitempty\""
		Status               *string "json:\"status,omitempty\""
		Timezone             *string "json:\"timezone,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CreateWithBody calls the Scans ´s function
func (c *Scans) CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	ContainerId          *string "json:\"container_id,omitempty\""
	CreationDate         *int32  "json:\"creation_date,omitempty\""
	CustomTargets        *string "json:\"custom_targets,omitempty\""
	DashboardFile        *string "json:\"dashboard_file,omitempty\""
	DefaultPermissions   *int32  "json:\"default_permissions,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Emails               *string "json:\"emails,omitempty\""
	Enabled              *bool   "json:\"enabled,omitempty\""
	Id                   *int32  "json:\"id,omitempty\""
	IncludeAggregate     *bool   "json:\"include_aggregate,omitempty\""
	LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NotificationFilters  *[]struct {
		Filter  *string "json:\"filter,omitempty\""
		Quality *string "json:\"quality,omitempty\""
		Value   *string "json:\"value,omitempty\""
	} "json:\"notification_filters,omitempty\""
	Owner             *string   "json:\"owner,omitempty\""
	OwnerId           *int      "json:\"owner_id,omitempty\""
	OwnerUuid         *string   "json:\"owner_uuid,omitempty\""
	PolicyId          *int      "json:\"policy_id,omitempty\""
	Remediation       *int      "json:\"remediation,omitempty\""
	Rrules            *string   "json:\"rrules,omitempty\""
	ScanTimeWindow    *string   "json:\"scan_time_window,omitempty\""
	ScannerUuid       *string   "json:\"scanner_uuid,omitempty\""
	Shared            *bool     "json:\"shared,omitempty\""
	Sms               *string   "json:\"sms,omitempty\""
	Starttime         *string   "json:\"starttime,omitempty\""
	TagTargets        *[]string "json:\"tag_targets,omitempty\""
	TagType           *string   "json:\"tag_type,omitempty\""
	TargetNetworkUuid *string   "json:\"target_network_uuid,omitempty\""
	Timezone          *string   "json:\"timezone,omitempty\""
	Type              *string   "json:\"type,omitempty\""
	UserPermissions   *int32    "json:\"user_permissions,omitempty\""
	Uuid              *string   "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansCreateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ContainerId *string "json:\"container_id,omitempty\""; CreationDate *int32 "json:\"creation_date,omitempty\""; CustomTargets *string "json:\"custom_targets,omitempty\""; DashboardFile *string "json:\"dashboard_file,omitempty\""; DefaultPermissions *int32 "json:\"default_permissions,omitempty\""; Description *string "json:\"description,omitempty\""; Emails *string "json:\"emails,omitempty\""; Enabled *bool "json:\"enabled,omitempty\""; Id *int32 "json:\"id,omitempty\""; IncludeAggregate *bool "json:\"include_aggregate,omitempty\""; LastModificationDate *int32 "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NotificationFilters *[]struct { Filter *string "json:\"filter,omitempty\""; Quality *string "json:\"quality,omitempty\""; Value *string "json:\"value,omitempty\"" } "json:\"notification_filters,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; PolicyId *int "json:\"policy_id,omitempty\""; Remediation *int "json:\"remediation,omitempty\""; Rrules *string "json:\"rrules,omitempty\""; ScanTimeWindow *string "json:\"scan_time_window,omitempty\""; ScannerUuid *string "json:\"scanner_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Sms *string "json:\"sms,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; TagTargets *[]string "json:\"tag_targets,omitempty\""; TagType *string "json:\"tag_type,omitempty\""; TargetNetworkUuid *string "json:\"target_network_uuid,omitempty\""; Timezone *string "json:\"timezone,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int32 "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		ContainerId          *string "json:\"container_id,omitempty\""
		CreationDate         *int32  "json:\"creation_date,omitempty\""
		CustomTargets        *string "json:\"custom_targets,omitempty\""
		DashboardFile        *string "json:\"dashboard_file,omitempty\""
		DefaultPermissions   *int32  "json:\"default_permissions,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Emails               *string "json:\"emails,omitempty\""
		Enabled              *bool   "json:\"enabled,omitempty\""
		Id                   *int32  "json:\"id,omitempty\""
		IncludeAggregate     *bool   "json:\"include_aggregate,omitempty\""
		LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NotificationFilters  *[]struct {
			Filter  *string "json:\"filter,omitempty\""
			Quality *string "json:\"quality,omitempty\""
			Value   *string "json:\"value,omitempty\""
		} "json:\"notification_filters,omitempty\""
		Owner             *string   "json:\"owner,omitempty\""
		OwnerId           *int      "json:\"owner_id,omitempty\""
		OwnerUuid         *string   "json:\"owner_uuid,omitempty\""
		PolicyId          *int      "json:\"policy_id,omitempty\""
		Remediation       *int      "json:\"remediation,omitempty\""
		Rrules            *string   "json:\"rrules,omitempty\""
		ScanTimeWindow    *string   "json:\"scan_time_window,omitempty\""
		ScannerUuid       *string   "json:\"scanner_uuid,omitempty\""
		Shared            *bool     "json:\"shared,omitempty\""
		Sms               *string   "json:\"sms,omitempty\""
		Starttime         *string   "json:\"starttime,omitempty\""
		TagTargets        *[]string "json:\"tag_targets,omitempty\""
		TagType           *string   "json:\"tag_type,omitempty\""
		TargetNetworkUuid *string   "json:\"target_network_uuid,omitempty\""
		Timezone          *string   "json:\"timezone,omitempty\""
		Type              *string   "json:\"type,omitempty\""
		UserPermissions   *int32    "json:\"user_permissions,omitempty\""
		Uuid              *string   "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Create calls the Scans ´s function
func (c *Scans) Create(arg1 ScansCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	ContainerId          *string "json:\"container_id,omitempty\""
	CreationDate         *int32  "json:\"creation_date,omitempty\""
	CustomTargets        *string "json:\"custom_targets,omitempty\""
	DashboardFile        *string "json:\"dashboard_file,omitempty\""
	DefaultPermissions   *int32  "json:\"default_permissions,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Emails               *string "json:\"emails,omitempty\""
	Enabled              *bool   "json:\"enabled,omitempty\""
	Id                   *int32  "json:\"id,omitempty\""
	IncludeAggregate     *bool   "json:\"include_aggregate,omitempty\""
	LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NotificationFilters  *[]struct {
		Filter  *string "json:\"filter,omitempty\""
		Quality *string "json:\"quality,omitempty\""
		Value   *string "json:\"value,omitempty\""
	} "json:\"notification_filters,omitempty\""
	Owner             *string   "json:\"owner,omitempty\""
	OwnerId           *int      "json:\"owner_id,omitempty\""
	OwnerUuid         *string   "json:\"owner_uuid,omitempty\""
	PolicyId          *int      "json:\"policy_id,omitempty\""
	Remediation       *int      "json:\"remediation,omitempty\""
	Rrules            *string   "json:\"rrules,omitempty\""
	ScanTimeWindow    *string   "json:\"scan_time_window,omitempty\""
	ScannerUuid       *string   "json:\"scanner_uuid,omitempty\""
	Shared            *bool     "json:\"shared,omitempty\""
	Sms               *string   "json:\"sms,omitempty\""
	Starttime         *string   "json:\"starttime,omitempty\""
	TagTargets        *[]string "json:\"tag_targets,omitempty\""
	TagType           *string   "json:\"tag_type,omitempty\""
	TargetNetworkUuid *string   "json:\"target_network_uuid,omitempty\""
	Timezone          *string   "json:\"timezone,omitempty\""
	Type              *string   "json:\"type,omitempty\""
	UserPermissions   *int32    "json:\"user_permissions,omitempty\""
	Uuid              *string   "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansCreateWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ContainerId *string "json:\"container_id,omitempty\""; CreationDate *int32 "json:\"creation_date,omitempty\""; CustomTargets *string "json:\"custom_targets,omitempty\""; DashboardFile *string "json:\"dashboard_file,omitempty\""; DefaultPermissions *int32 "json:\"default_permissions,omitempty\""; Description *string "json:\"description,omitempty\""; Emails *string "json:\"emails,omitempty\""; Enabled *bool "json:\"enabled,omitempty\""; Id *int32 "json:\"id,omitempty\""; IncludeAggregate *bool "json:\"include_aggregate,omitempty\""; LastModificationDate *int32 "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NotificationFilters *[]struct { Filter *string "json:\"filter,omitempty\""; Quality *string "json:\"quality,omitempty\""; Value *string "json:\"value,omitempty\"" } "json:\"notification_filters,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; PolicyId *int "json:\"policy_id,omitempty\""; Remediation *int "json:\"remediation,omitempty\""; Rrules *string "json:\"rrules,omitempty\""; ScanTimeWindow *string "json:\"scan_time_window,omitempty\""; ScannerUuid *string "json:\"scanner_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Sms *string "json:\"sms,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; TagTargets *[]string "json:\"tag_targets,omitempty\""; TagType *string "json:\"tag_type,omitempty\""; TargetNetworkUuid *string "json:\"target_network_uuid,omitempty\""; Timezone *string "json:\"timezone,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int32 "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		ContainerId          *string "json:\"container_id,omitempty\""
		CreationDate         *int32  "json:\"creation_date,omitempty\""
		CustomTargets        *string "json:\"custom_targets,omitempty\""
		DashboardFile        *string "json:\"dashboard_file,omitempty\""
		DefaultPermissions   *int32  "json:\"default_permissions,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Emails               *string "json:\"emails,omitempty\""
		Enabled              *bool   "json:\"enabled,omitempty\""
		Id                   *int32  "json:\"id,omitempty\""
		IncludeAggregate     *bool   "json:\"include_aggregate,omitempty\""
		LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NotificationFilters  *[]struct {
			Filter  *string "json:\"filter,omitempty\""
			Quality *string "json:\"quality,omitempty\""
			Value   *string "json:\"value,omitempty\""
		} "json:\"notification_filters,omitempty\""
		Owner             *string   "json:\"owner,omitempty\""
		OwnerId           *int      "json:\"owner_id,omitempty\""
		OwnerUuid         *string   "json:\"owner_uuid,omitempty\""
		PolicyId          *int      "json:\"policy_id,omitempty\""
		Remediation       *int      "json:\"remediation,omitempty\""
		Rrules            *string   "json:\"rrules,omitempty\""
		ScanTimeWindow    *string   "json:\"scan_time_window,omitempty\""
		ScannerUuid       *string   "json:\"scanner_uuid,omitempty\""
		Shared            *bool     "json:\"shared,omitempty\""
		Sms               *string   "json:\"sms,omitempty\""
		Starttime         *string   "json:\"starttime,omitempty\""
		TagTargets        *[]string "json:\"tag_targets,omitempty\""
		TagType           *string   "json:\"tag_type,omitempty\""
		TargetNetworkUuid *string   "json:\"target_network_uuid,omitempty\""
		Timezone          *string   "json:\"timezone,omitempty\""
		Type              *string   "json:\"type,omitempty\""
		UserPermissions   *int32    "json:\"user_permissions,omitempty\""
		Uuid              *string   "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DeleteHistory calls the Scans ´s function
func (c *Scans) DeleteHistory(arg1 string, arg2 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansDeleteHistoryWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Delete calls the Scans ´s function
func (c *Scans) Delete(contentType string, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansDeleteWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Details calls the Scans ´s function
func (c *Scans) Details(arg1 string, params *ScansDetailsParams, reqEditors ...RequestEditorFn) (*struct {
	Comphosts *[]struct {
		Critical              *int                    "json:\"critical,omitempty\""
		High                  *int                    "json:\"high,omitempty\""
		HostId                *int                    "json:\"host_id,omitempty\""
		HostIndex             *int                    "json:\"host_index,omitempty\""
		Hostname              *string                 "json:\"hostname,omitempty\""
		Info                  *int                    "json:\"info,omitempty\""
		Low                   *int                    "json:\"low,omitempty\""
		Medium                *int                    "json:\"medium,omitempty\""
		Numchecksconsidered   *int                    "json:\"numchecksconsidered,omitempty\""
		Progress              *string                 "json:\"progress,omitempty\""
		Scanprogresscurrent   *int                    "json:\"scanprogresscurrent,omitempty\""
		Scanprogresstotal     *int                    "json:\"scanprogresstotal,omitempty\""
		Score                 *int                    "json:\"score,omitempty\""
		Severitycount         *map[string]interface{} "json:\"severitycount,omitempty\""
		Totalchecksconsidered *int                    "json:\"totalchecksconsidered,omitempty\""
	} "json:\"comphosts,omitempty\""
	Compliance *[]struct {
		Count         *int    "json:\"count,omitempty\""
		PluginFamily  *string "json:\"plugin_family,omitempty\""
		PluginId      *int    "json:\"plugin_id,omitempty\""
		PluginName    *string "json:\"plugin_name,omitempty\""
		Severity      *int    "json:\"severity,omitempty\""
		SeverityIndex *int    "json:\"severity_index,omitempty\""
		VulnIndex     *int    "json:\"vuln_index,omitempty\""
	} "json:\"compliance,omitempty\""
	Filters *[]struct {
		Control *struct {
			Options        *[]string "json:\"options,omitempty\""
			ReadableRegest *string   "json:\"readable_regest,omitempty\""
			Regex          *string   "json:\"regex,omitempty\""
			Type           *string   "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	} "json:\"filters,omitempty\""
	History *[]struct {
		AltTargetsUsed       *bool   "json:\"alt_targets_used,omitempty\""
		CreationDate         *int    "json:\"creation_date,omitempty\""
		HistoryId            *int    "json:\"history_id,omitempty\""
		IsArchived           *bool   "json:\"is_archived,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		Scheduler            *int    "json:\"scheduler,omitempty\""
		Status               *string "json:\"status,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	} "json:\"history,omitempty\""
	Hosts *[]struct {
		Critical              *int                    "json:\"critical,omitempty\""
		High                  *int                    "json:\"high,omitempty\""
		HostId                *int                    "json:\"host_id,omitempty\""
		HostIndex             *int                    "json:\"host_index,omitempty\""
		Hostname              *string                 "json:\"hostname,omitempty\""
		Info                  *int                    "json:\"info,omitempty\""
		Low                   *int                    "json:\"low,omitempty\""
		Medium                *int                    "json:\"medium,omitempty\""
		Numchecksconsidered   *int                    "json:\"numchecksconsidered,omitempty\""
		Progress              *string                 "json:\"progress,omitempty\""
		Scanprogresscurrent   *int                    "json:\"scanprogresscurrent,omitempty\""
		Scanprogresstotal     *int                    "json:\"scanprogresstotal,omitempty\""
		Score                 *int                    "json:\"score,omitempty\""
		Severitycount         *map[string]interface{} "json:\"severitycount,omitempty\""
		Totalchecksconsidered *int                    "json:\"totalchecksconsidered,omitempty\""
	} "json:\"hosts,omitempty\""
	Info *struct {
		Acls *[]struct {
			DisplayName *string "json:\"display_name,omitempty\""
			Id          *int    "json:\"id,omitempty\""
			Name        *string "json:\"name,omitempty\""
			Owner       *int    "json:\"owner,omitempty\""
			Permissions *int    "json:\"permissions,omitempty\""
			Type        *string "json:\"type,omitempty\""
		} "json:\"acls,omitempty\""
		AltTargetsUsed  *bool     "json:\"alt_targets_used,omitempty\""
		Control         *bool     "json:\"control,omitempty\""
		EditAllowed     *bool     "json:\"edit_allowed,omitempty\""
		FolderId        *int32    "json:\"folder_id,omitempty\""
		Hasaudittrail   *bool     "json:\"hasaudittrail,omitempty\""
		Haskb           *bool     "json:\"haskb,omitempty\""
		Hostcount       *int32    "json:\"hostcount,omitempty\""
		IsArchived      *bool     "json:\"is_archived,omitempty\""
		Name            *string   "json:\"name,omitempty\""
		NoTarget        *bool     "json:\"no_target,omitempty\""
		ObjectId        *int32    "json:\"object_id,omitempty\""
		Owner           *string   "json:\"owner,omitempty\""
		PciCanUpload    *bool     "json:\"pci-can-upload,omitempty\""
		Policy          *string   "json:\"policy,omitempty\""
		ScanEnd         *int      "json:\"scan_end,omitempty\""
		ScanStart       *int32    "json:\"scan_start,omitempty\""
		ScanType        *string   "json:\"scan_type,omitempty\""
		ScannerEnd      *int      "json:\"scanner_end,omitempty\""
		ScannerName     *string   "json:\"scanner_name,omitempty\""
		ScannerStart    *int      "json:\"scanner_start,omitempty\""
		ScheduleUuid    *string   "json:\"schedule_uuid,omitempty\""
		Shared          *bool     "json:\"shared,omitempty\""
		Status          *string   "json:\"status,omitempty\""
		TagTargets      *[]string "json:\"tag_targets,omitempty\""
		Targets         *string   "json:\"targets,omitempty\""
		Timestamp       *int      "json:\"timestamp,omitempty\""
		UserPermissions *int32    "json:\"user_permissions,omitempty\""
		Uuid            *string   "json:\"uuid,omitempty\""
	} "json:\"info,omitempty\""
	Notes *[]struct {
		Message  *string "json:\"message,omitempty\""
		Severity *int    "json:\"severity,omitempty\""
		Title    *string "json:\"title,omitempty\""
	} "json:\"notes,omitempty\""
	Remediations    *map[string]interface{} "json:\"remediations,omitempty\""
	Vulnerabilities *[]struct {
		Count         *int    "json:\"count,omitempty\""
		PluginFamily  *string "json:\"plugin_family,omitempty\""
		PluginId      *int    "json:\"plugin_id,omitempty\""
		PluginName    *string "json:\"plugin_name,omitempty\""
		Severity      *int    "json:\"severity,omitempty\""
		SeverityIndex *int    "json:\"severity_index,omitempty\""
		VulnIndex     *int    "json:\"vuln_index,omitempty\""
	} "json:\"vulnerabilities,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Comphosts *[]struct { Critical *int "json:\"critical,omitempty\""; High *int "json:\"high,omitempty\""; HostId *int "json:\"host_id,omitempty\""; HostIndex *int "json:\"host_index,omitempty\""; Hostname *string "json:\"hostname,omitempty\""; Info *int "json:\"info,omitempty\""; Low *int "json:\"low,omitempty\""; Medium *int "json:\"medium,omitempty\""; Numchecksconsidered *int "json:\"numchecksconsidered,omitempty\""; Progress *string "json:\"progress,omitempty\""; Scanprogresscurrent *int "json:\"scanprogresscurrent,omitempty\""; Scanprogresstotal *int "json:\"scanprogresstotal,omitempty\""; Score *int "json:\"score,omitempty\""; Severitycount *map[string]interface {} "json:\"severitycount,omitempty\""; Totalchecksconsidered *int "json:\"totalchecksconsidered,omitempty\"" } "json:\"comphosts,omitempty\""; Compliance *[]struct { Count *int "json:\"count,omitempty\""; PluginFamily *string "json:\"plugin_family,omitempty\""; PluginId *int "json:\"plugin_id,omitempty\""; PluginName *string "json:\"plugin_name,omitempty\""; Severity *int "json:\"severity,omitempty\""; SeverityIndex *int "json:\"severity_index,omitempty\""; VulnIndex *int "json:\"vuln_index,omitempty\"" } "json:\"compliance,omitempty\""; Filters *[]struct { Control *struct { Options *[]string "json:\"options,omitempty\""; ReadableRegest *string "json:\"readable_regest,omitempty\""; Regex *string "json:\"regex,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"control,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" } "json:\"filters,omitempty\""; History *[]struct { AltTargetsUsed *bool "json:\"alt_targets_used,omitempty\""; CreationDate *int "json:\"creation_date,omitempty\""; HistoryId *int "json:\"history_id,omitempty\""; IsArchived *bool "json:\"is_archived,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; Scheduler *int "json:\"scheduler,omitempty\""; Status *string "json:\"status,omitempty\""; Type *string "json:\"type,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"history,omitempty\""; Hosts *[]struct { Critical *int "json:\"critical,omitempty\""; High *int "json:\"high,omitempty\""; HostId *int "json:\"host_id,omitempty\""; HostIndex *int "json:\"host_index,omitempty\""; Hostname *string "json:\"hostname,omitempty\""; Info *int "json:\"info,omitempty\""; Low *int "json:\"low,omitempty\""; Medium *int "json:\"medium,omitempty\""; Numchecksconsidered *int "json:\"numchecksconsidered,omitempty\""; Progress *string "json:\"progress,omitempty\""; Scanprogresscurrent *int "json:\"scanprogresscurrent,omitempty\""; Scanprogresstotal *int "json:\"scanprogresstotal,omitempty\""; Score *int "json:\"score,omitempty\""; Severitycount *map[string]interface {} "json:\"severitycount,omitempty\""; Totalchecksconsidered *int "json:\"totalchecksconsidered,omitempty\"" } "json:\"hosts,omitempty\""; Info *struct { Acls *[]struct { DisplayName *string "json:\"display_name,omitempty\""; Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *int "json:\"owner,omitempty\""; Permissions *int "json:\"permissions,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"acls,omitempty\""; AltTargetsUsed *bool "json:\"alt_targets_used,omitempty\""; Control *bool "json:\"control,omitempty\""; EditAllowed *bool "json:\"edit_allowed,omitempty\""; FolderId *int32 "json:\"folder_id,omitempty\""; Hasaudittrail *bool "json:\"hasaudittrail,omitempty\""; Haskb *bool "json:\"haskb,omitempty\""; Hostcount *int32 "json:\"hostcount,omitempty\""; IsArchived *bool "json:\"is_archived,omitempty\""; Name *string "json:\"name,omitempty\""; NoTarget *bool "json:\"no_target,omitempty\""; ObjectId *int32 "json:\"object_id,omitempty\""; Owner *string "json:\"owner,omitempty\""; PciCanUpload *bool "json:\"pci-can-upload,omitempty\""; Policy *string "json:\"policy,omitempty\""; ScanEnd *int "json:\"scan_end,omitempty\""; ScanStart *int32 "json:\"scan_start,omitempty\""; ScanType *string "json:\"scan_type,omitempty\""; ScannerEnd      *int "json:\"scanner_end,omitempty\""; ScannerName *string "json:\"scanner_name,omitempty\""; ScannerStart *int "json:\"scanner_start,omitempty\""; ScheduleUuid *string "json:\"schedule_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Status *string "json:\"status,omitempty\""; TagTargets *[]string "json:\"tag_targets,omitempty\""; Targets *string "json:\"targets,omitempty\""; Timestamp *int "json:\"timestamp,omitempty\""; UserPermissions *int32 "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"info,omitempty\""; Notes *[]struct { Message *string "json:\"message,omitempty\""; Severity *int "json:\"severity,omitempty\""; Title *string "json:\"title,omitempty\"" } "json:\"notes,omitempty\""; Remediations *map[string]interface {} "json:\"remediations,omitempty\""; Vulnerabilities *[]struct { Count *int "json:\"count,omitempty\""; PluginFamily *string "json:\"plugin_family,omitempty\""; PluginId *int "json:\"plugin_id,omitempty\""; PluginName *string "json:\"plugin_name,omitempty\""; Severity *int "json:\"severity,omitempty\""; SeverityIndex *int "json:\"severity_index,omitempty\""; VulnIndex *int "json:\"vuln_index,omitempty\"" } "json:\"vulnerabilities,omitempty\"" }
	if i, ok := r.(*struct {
		Comphosts *[]struct {
			Critical              *int                    "json:\"critical,omitempty\""
			High                  *int                    "json:\"high,omitempty\""
			HostId                *int                    "json:\"host_id,omitempty\""
			HostIndex             *int                    "json:\"host_index,omitempty\""
			Hostname              *string                 "json:\"hostname,omitempty\""
			Info                  *int                    "json:\"info,omitempty\""
			Low                   *int                    "json:\"low,omitempty\""
			Medium                *int                    "json:\"medium,omitempty\""
			Numchecksconsidered   *int                    "json:\"numchecksconsidered,omitempty\""
			Progress              *string                 "json:\"progress,omitempty\""
			Scanprogresscurrent   *int                    "json:\"scanprogresscurrent,omitempty\""
			Scanprogresstotal     *int                    "json:\"scanprogresstotal,omitempty\""
			Score                 *int                    "json:\"score,omitempty\""
			Severitycount         *map[string]interface{} "json:\"severitycount,omitempty\""
			Totalchecksconsidered *int                    "json:\"totalchecksconsidered,omitempty\""
		} "json:\"comphosts,omitempty\""
		Compliance *[]struct {
			Count         *int    "json:\"count,omitempty\""
			PluginFamily  *string "json:\"plugin_family,omitempty\""
			PluginId      *int    "json:\"plugin_id,omitempty\""
			PluginName    *string "json:\"plugin_name,omitempty\""
			Severity      *int    "json:\"severity,omitempty\""
			SeverityIndex *int    "json:\"severity_index,omitempty\""
			VulnIndex     *int    "json:\"vuln_index,omitempty\""
		} "json:\"compliance,omitempty\""
		Filters *[]struct {
			Control *struct {
				Options        *[]string "json:\"options,omitempty\""
				ReadableRegest *string   "json:\"readable_regest,omitempty\""
				Regex          *string   "json:\"regex,omitempty\""
				Type           *string   "json:\"type,omitempty\""
			} "json:\"control,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"filters,omitempty\""
		History *[]struct {
			AltTargetsUsed       *bool   "json:\"alt_targets_used,omitempty\""
			CreationDate         *int    "json:\"creation_date,omitempty\""
			HistoryId            *int    "json:\"history_id,omitempty\""
			IsArchived           *bool   "json:\"is_archived,omitempty\""
			LastModificationDate *int    "json:\"last_modification_date,omitempty\""
			OwnerId              *int    "json:\"owner_id,omitempty\""
			Scheduler            *int    "json:\"scheduler,omitempty\""
			Status               *string "json:\"status,omitempty\""
			Type                 *string "json:\"type,omitempty\""
			Uuid                 *string "json:\"uuid,omitempty\""
		} "json:\"history,omitempty\""
		Hosts *[]struct {
			Critical              *int                    "json:\"critical,omitempty\""
			High                  *int                    "json:\"high,omitempty\""
			HostId                *int                    "json:\"host_id,omitempty\""
			HostIndex             *int                    "json:\"host_index,omitempty\""
			Hostname              *string                 "json:\"hostname,omitempty\""
			Info                  *int                    "json:\"info,omitempty\""
			Low                   *int                    "json:\"low,omitempty\""
			Medium                *int                    "json:\"medium,omitempty\""
			Numchecksconsidered   *int                    "json:\"numchecksconsidered,omitempty\""
			Progress              *string                 "json:\"progress,omitempty\""
			Scanprogresscurrent   *int                    "json:\"scanprogresscurrent,omitempty\""
			Scanprogresstotal     *int                    "json:\"scanprogresstotal,omitempty\""
			Score                 *int                    "json:\"score,omitempty\""
			Severitycount         *map[string]interface{} "json:\"severitycount,omitempty\""
			Totalchecksconsidered *int                    "json:\"totalchecksconsidered,omitempty\""
		} "json:\"hosts,omitempty\""
		Info *struct {
			Acls *[]struct {
				DisplayName *string "json:\"display_name,omitempty\""
				Id          *int    "json:\"id,omitempty\""
				Name        *string "json:\"name,omitempty\""
				Owner       *int    "json:\"owner,omitempty\""
				Permissions *int    "json:\"permissions,omitempty\""
				Type        *string "json:\"type,omitempty\""
			} "json:\"acls,omitempty\""
			AltTargetsUsed  *bool     "json:\"alt_targets_used,omitempty\""
			Control         *bool     "json:\"control,omitempty\""
			EditAllowed     *bool     "json:\"edit_allowed,omitempty\""
			FolderId        *int32    "json:\"folder_id,omitempty\""
			Hasaudittrail   *bool     "json:\"hasaudittrail,omitempty\""
			Haskb           *bool     "json:\"haskb,omitempty\""
			Hostcount       *int32    "json:\"hostcount,omitempty\""
			IsArchived      *bool     "json:\"is_archived,omitempty\""
			Name            *string   "json:\"name,omitempty\""
			NoTarget        *bool     "json:\"no_target,omitempty\""
			ObjectId        *int32    "json:\"object_id,omitempty\""
			Owner           *string   "json:\"owner,omitempty\""
			PciCanUpload    *bool     "json:\"pci-can-upload,omitempty\""
			Policy          *string   "json:\"policy,omitempty\""
			ScanEnd         *int      "json:\"scan_end,omitempty\""
			ScanStart       *int32    "json:\"scan_start,omitempty\""
			ScanType        *string   "json:\"scan_type,omitempty\""
			ScannerEnd      *int      "json:\"scanner_end,omitempty\""
			ScannerName     *string   "json:\"scanner_name,omitempty\""
			ScannerStart    *int      "json:\"scanner_start,omitempty\""
			ScheduleUuid    *string   "json:\"schedule_uuid,omitempty\""
			Shared          *bool     "json:\"shared,omitempty\""
			Status          *string   "json:\"status,omitempty\""
			TagTargets      *[]string "json:\"tag_targets,omitempty\""
			Targets         *string   "json:\"targets,omitempty\""
			Timestamp       *int      "json:\"timestamp,omitempty\""
			UserPermissions *int32    "json:\"user_permissions,omitempty\""
			Uuid            *string   "json:\"uuid,omitempty\""
		} "json:\"info,omitempty\""
		Notes *[]struct {
			Message  *string "json:\"message,omitempty\""
			Severity *int    "json:\"severity,omitempty\""
			Title    *string "json:\"title,omitempty\""
		} "json:\"notes,omitempty\""
		Remediations    *map[string]interface{} "json:\"remediations,omitempty\""
		Vulnerabilities *[]struct {
			Count         *int    "json:\"count,omitempty\""
			PluginFamily  *string "json:\"plugin_family,omitempty\""
			PluginId      *int    "json:\"plugin_id,omitempty\""
			PluginName    *string "json:\"plugin_name,omitempty\""
			Severity      *int    "json:\"severity,omitempty\""
			SeverityIndex *int    "json:\"severity_index,omitempty\""
			VulnIndex     *int    "json:\"vuln_index,omitempty\""
		} "json:\"vulnerabilities,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ExportDownload calls the Scans ´s function
func (c *Scans) ExportDownload(arg1 string, contentType string, reqEditors ...RequestEditorFn) (*ScansExportDownloadResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansExportDownloadWithResponse(c.ClientInterface.(*Client).ctx, arg1, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *ScansExportDownloadResponse
	if i, ok := r.(*ScansExportDownloadResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ExportRequestWithBody calls the Scans ´s function
func (c *Scans) ExportRequestWithBody(arg1 string, params *ScansExportRequestParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	File      *string "json:\"file,omitempty\""
	TempToken *string "json:\"temp_token,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansExportRequestWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { File *string "json:\"file,omitempty\""; TempToken *string "json:\"temp_token,omitempty\"" }
	if i, ok := r.(*struct {
		File      *string "json:\"file,omitempty\""
		TempToken *string "json:\"temp_token,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ExportRequest calls the Scans ´s function
func (c *Scans) ExportRequest(arg1 string, params *ScansExportRequestParams, arg3 ScansExportRequestJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	File      *string "json:\"file,omitempty\""
	TempToken *string "json:\"temp_token,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansExportRequestWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, arg3, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { File *string "json:\"file,omitempty\""; TempToken *string "json:\"temp_token,omitempty\"" }
	if i, ok := r.(*struct {
		File      *string "json:\"file,omitempty\""
		TempToken *string "json:\"temp_token,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ExportStatus calls the Scans ´s function
func (c *Scans) ExportStatus(arg1 string, contentType string, reqEditors ...RequestEditorFn) (*struct {
	Status *string "json:\"status,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansExportStatusWithResponse(c.ClientInterface.(*Client).ctx, arg1, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Status *string "json:\"status,omitempty\"" }
	if i, ok := r.(*struct {
		Status *string "json:\"status,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GetLatestStatus calls the Scans ´s function
func (c *Scans) GetLatestStatus(contentType string, reqEditors ...RequestEditorFn) (*struct {
	Status *string "json:\"status,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansGetLatestStatusWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Status *string "json:\"status,omitempty\"" }
	if i, ok := r.(*struct {
		Status *string "json:\"status,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// HistoryDetails calls the Scans ´s function
func (c *Scans) HistoryDetails(arg1 string, contentType string, reqEditors ...RequestEditorFn) (*struct {
	IsArchived   *bool   "json:\"is_archived,omitempty\""
	Name         *string "json:\"name,omitempty\""
	ObjectId     *int32  "json:\"object_id,omitempty\""
	Owner        *string "json:\"owner,omitempty\""
	OwnerId      *int    "json:\"owner_id,omitempty\""
	OwnerUuid    *int    "json:\"owner_uuid,omitempty\""
	ScanEnd      *int    "json:\"scan_end,omitempty\""
	ScanStart    *int32  "json:\"scan_start,omitempty\""
	ScanType     *string "json:\"scan_type,omitempty\""
	ScheduleUuid *string "json:\"schedule_uuid,omitempty\""
	Status       *string "json:\"status,omitempty\""
	Targets      *string "json:\"targets,omitempty\""
	Uuid         *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansHistoryDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { IsArchived *bool "json:\"is_archived,omitempty\""; Name *string "json:\"name,omitempty\""; ObjectId *int32 "json:\"object_id,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; OwnerUuid *int "json:\"owner_uuid,omitempty\""; ScanEnd *int "json:\"scan_end,omitempty\""; ScanStart *int32 "json:\"scan_start,omitempty\""; ScanType *string "json:\"scan_type,omitempty\""; ScheduleUuid *string "json:\"schedule_uuid,omitempty\""; Status *string "json:\"status,omitempty\""; Targets *string "json:\"targets,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		IsArchived   *bool   "json:\"is_archived,omitempty\""
		Name         *string "json:\"name,omitempty\""
		ObjectId     *int32  "json:\"object_id,omitempty\""
		Owner        *string "json:\"owner,omitempty\""
		OwnerId      *int    "json:\"owner_id,omitempty\""
		OwnerUuid    *int    "json:\"owner_uuid,omitempty\""
		ScanEnd      *int    "json:\"scan_end,omitempty\""
		ScanStart    *int32  "json:\"scan_start,omitempty\""
		ScanType     *string "json:\"scan_type,omitempty\""
		ScheduleUuid *string "json:\"schedule_uuid,omitempty\""
		Status       *string "json:\"status,omitempty\""
		Targets      *string "json:\"targets,omitempty\""
		Uuid         *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// History calls the Scans ´s function
func (c *Scans) History(arg1 string, params *ScansHistoryParams, reqEditors ...RequestEditorFn) (*struct {
	History *[]struct {
		Id         *int    "json:\"id,omitempty\""
		IsArchived *bool   "json:\"is_archived,omitempty\""
		ScanUuid   *string "json:\"scan_uuid,omitempty\""
		Status     *string "json:\"status,omitempty\""
		Targets    *struct {
			Custom  *bool "json:\"custom,omitempty\""
			Default *bool "json:\"default,omitempty\""
		} "json:\"targets,omitempty\""
		TimeEnd    *int    "json:\"time_end,omitempty\""
		TimeStart  *int    "json:\"time_start,omitempty\""
		Visibility *string "json:\"visibility,omitempty\""
	} "json:\"history,omitempty\""
	Pagination *struct {
		Limit  *int32 "json:\"limit,omitempty\""
		Offset *int32 "json:\"offset,omitempty\""
		Sort   *[]struct {
			Name  *string "json:\"name,omitempty\""
			Order *string "json:\"order,omitempty\""
		} "json:\"sort,omitempty\""
		Total *int32 "json:\"total,omitempty\""
	} "json:\"pagination,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansHistoryWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { History *[]struct { Id *int "json:\"id,omitempty\""; IsArchived *bool "json:\"is_archived,omitempty\""; ScanUuid *string "json:\"scan_uuid,omitempty\""; Status *string "json:\"status,omitempty\""; Targets *struct { Custom *bool "json:\"custom,omitempty\""; Default *bool "json:\"default,omitempty\"" } "json:\"targets,omitempty\""; TimeEnd *int "json:\"time_end,omitempty\""; TimeStart *int "json:\"time_start,omitempty\""; Visibility *string "json:\"visibility,omitempty\"" } "json:\"history,omitempty\""; Pagination *struct { Limit *int32 "json:\"limit,omitempty\""; Offset *int32 "json:\"offset,omitempty\""; Sort *[]struct { Name *string "json:\"name,omitempty\""; Order *string "json:\"order,omitempty\"" } "json:\"sort,omitempty\""; Total *int32 "json:\"total,omitempty\"" } "json:\"pagination,omitempty\"" }
	if i, ok := r.(*struct {
		History *[]struct {
			Id         *int    "json:\"id,omitempty\""
			IsArchived *bool   "json:\"is_archived,omitempty\""
			ScanUuid   *string "json:\"scan_uuid,omitempty\""
			Status     *string "json:\"status,omitempty\""
			Targets    *struct {
				Custom  *bool "json:\"custom,omitempty\""
				Default *bool "json:\"default,omitempty\""
			} "json:\"targets,omitempty\""
			TimeEnd    *int    "json:\"time_end,omitempty\""
			TimeStart  *int    "json:\"time_start,omitempty\""
			Visibility *string "json:\"visibility,omitempty\""
		} "json:\"history,omitempty\""
		Pagination *struct {
			Limit  *int32 "json:\"limit,omitempty\""
			Offset *int32 "json:\"offset,omitempty\""
			Sort   *[]struct {
				Name  *string "json:\"name,omitempty\""
				Order *string "json:\"order,omitempty\""
			} "json:\"sort,omitempty\""
			Total *int32 "json:\"total,omitempty\""
		} "json:\"pagination,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// HostDetails calls the Scans ´s function
func (c *Scans) HostDetails(arg1 string, arg2 int32, params *ScansHostDetailsParams, reqEditors ...RequestEditorFn) (*struct {
	Compliance *[]struct {
		Count         *int    "json:\"count,omitempty\""
		HostId        *int    "json:\"host_id,omitempty\""
		Hostname      *string "json:\"hostname,omitempty\""
		PluginFamily  *string "json:\"plugin_family,omitempty\""
		PluginId      *int    "json:\"plugin_id,omitempty\""
		PluginName    *string "json:\"plugin_name,omitempty\""
		Severity      *int    "json:\"severity,omitempty\""
		SeverityIndex *int    "json:\"severity_index,omitempty\""
	} "json:\"compliance,omitempty\""
	Info            *map[string]interface{} "json:\"info,omitempty\""
	Vulnerabilities *[]struct {
		Count         *int    "json:\"count,omitempty\""
		HostId        *int    "json:\"host_id,omitempty\""
		Hostname      *string "json:\"hostname,omitempty\""
		PluginFamily  *string "json:\"plugin_family,omitempty\""
		PluginId      *int    "json:\"plugin_id,omitempty\""
		PluginName    *string "json:\"plugin_name,omitempty\""
		Severity      *int    "json:\"severity,omitempty\""
		SeverityIndex *int    "json:\"severity_index,omitempty\""
		VulnIndex     *int    "json:\"vuln_index,omitempty\""
	} "json:\"vulnerabilities,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansHostDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Compliance *[]struct { Count *int "json:\"count,omitempty\""; HostId *int "json:\"host_id,omitempty\""; Hostname *string "json:\"hostname,omitempty\""; PluginFamily *string "json:\"plugin_family,omitempty\""; PluginId *int "json:\"plugin_id,omitempty\""; PluginName *string "json:\"plugin_name,omitempty\""; Severity *int "json:\"severity,omitempty\""; SeverityIndex *int "json:\"severity_index,omitempty\"" } "json:\"compliance,omitempty\""; Info *map[string]interface {} "json:\"info,omitempty\""; Vulnerabilities *[]struct { Count *int "json:\"count,omitempty\""; HostId *int "json:\"host_id,omitempty\""; Hostname *string "json:\"hostname,omitempty\""; PluginFamily *string "json:\"plugin_family,omitempty\""; PluginId *int "json:\"plugin_id,omitempty\""; PluginName *string "json:\"plugin_name,omitempty\""; Severity *int "json:\"severity,omitempty\""; SeverityIndex *int "json:\"severity_index,omitempty\""; VulnIndex *int "json:\"vuln_index,omitempty\"" } "json:\"vulnerabilities,omitempty\"" }
	if i, ok := r.(*struct {
		Compliance *[]struct {
			Count         *int    "json:\"count,omitempty\""
			HostId        *int    "json:\"host_id,omitempty\""
			Hostname      *string "json:\"hostname,omitempty\""
			PluginFamily  *string "json:\"plugin_family,omitempty\""
			PluginId      *int    "json:\"plugin_id,omitempty\""
			PluginName    *string "json:\"plugin_name,omitempty\""
			Severity      *int    "json:\"severity,omitempty\""
			SeverityIndex *int    "json:\"severity_index,omitempty\""
		} "json:\"compliance,omitempty\""
		Info            *map[string]interface{} "json:\"info,omitempty\""
		Vulnerabilities *[]struct {
			Count         *int    "json:\"count,omitempty\""
			HostId        *int    "json:\"host_id,omitempty\""
			Hostname      *string "json:\"hostname,omitempty\""
			PluginFamily  *string "json:\"plugin_family,omitempty\""
			PluginId      *int    "json:\"plugin_id,omitempty\""
			PluginName    *string "json:\"plugin_name,omitempty\""
			Severity      *int    "json:\"severity,omitempty\""
			SeverityIndex *int    "json:\"severity_index,omitempty\""
			VulnIndex     *int    "json:\"vuln_index,omitempty\""
		} "json:\"vulnerabilities,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ImportWithBody calls the Scans ´s function
func (c *Scans) ImportWithBody(params *ScansImportParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	Control              *bool   "json:\"control,omitempty\""
	CreationDate         *int    "json:\"creation_date,omitempty\""
	Enabled              *bool   "json:\"enabled,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	Read                 *bool   "json:\"read,omitempty\""
	Rrules               *string "json:\"rrules,omitempty\""
	ScheduleUuid         *string "json:\"schedule_uuid,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Starttime            *string "json:\"starttime,omitempty\""
	Status               *string "json:\"status,omitempty\""
	Timezone             *string "json:\"timezone,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
	Uuid                 *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansImportWithBodyWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Control *bool "json:\"control,omitempty\""; CreationDate *int "json:\"creation_date,omitempty\""; Enabled *bool "json:\"enabled,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int32 "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; Read *bool "json:\"read,omitempty\""; Rrules *string "json:\"rrules,omitempty\""; ScheduleUuid *string "json:\"schedule_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; Status *string "json:\"status,omitempty\""; Timezone *string "json:\"timezone,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		Control              *bool   "json:\"control,omitempty\""
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Enabled              *bool   "json:\"enabled,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		Read                 *bool   "json:\"read,omitempty\""
		Rrules               *string "json:\"rrules,omitempty\""
		ScheduleUuid         *string "json:\"schedule_uuid,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Starttime            *string "json:\"starttime,omitempty\""
		Status               *string "json:\"status,omitempty\""
		Timezone             *string "json:\"timezone,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Import calls the Scans ´s function
func (c *Scans) Import(params *ScansImportParams, arg2 ScansImportJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	Control              *bool   "json:\"control,omitempty\""
	CreationDate         *int    "json:\"creation_date,omitempty\""
	Enabled              *bool   "json:\"enabled,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	Read                 *bool   "json:\"read,omitempty\""
	Rrules               *string "json:\"rrules,omitempty\""
	ScheduleUuid         *string "json:\"schedule_uuid,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Starttime            *string "json:\"starttime,omitempty\""
	Status               *string "json:\"status,omitempty\""
	Timezone             *string "json:\"timezone,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
	Uuid                 *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansImportWithResponse(c.ClientInterface.(*Client).ctx, params, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Control *bool "json:\"control,omitempty\""; CreationDate *int "json:\"creation_date,omitempty\""; Enabled *bool "json:\"enabled,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int32 "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; Read *bool "json:\"read,omitempty\""; Rrules *string "json:\"rrules,omitempty\""; ScheduleUuid *string "json:\"schedule_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; Status *string "json:\"status,omitempty\""; Timezone *string "json:\"timezone,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		Control              *bool   "json:\"control,omitempty\""
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Enabled              *bool   "json:\"enabled,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		Read                 *bool   "json:\"read,omitempty\""
		Rrules               *string "json:\"rrules,omitempty\""
		ScheduleUuid         *string "json:\"schedule_uuid,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Starttime            *string "json:\"starttime,omitempty\""
		Status               *string "json:\"status,omitempty\""
		Timezone             *string "json:\"timezone,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// LaunchWithBody calls the Scans ´s function
func (c *Scans) LaunchWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	ScanUuid *string "json:\"scan_uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansLaunchWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ScanUuid *string "json:\"scan_uuid,omitempty\"" }
	if i, ok := r.(*struct {
		ScanUuid *string "json:\"scan_uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Launch calls the Scans ´s function
func (c *Scans) Launch(arg1 string, arg2 ScansLaunchJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	ScanUuid *string "json:\"scan_uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansLaunchWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ScanUuid *string "json:\"scan_uuid,omitempty\"" }
	if i, ok := r.(*struct {
		ScanUuid *string "json:\"scan_uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// List calls the Scans ´s function
func (c *Scans) List(params *ScansListParams, reqEditors ...RequestEditorFn) (*struct {
	Folders *[]struct {
		Custom      *int    "json:\"custom,omitempty\""
		DefaultTag  *int    "json:\"default_tag,omitempty\""
		Id          *int    "json:\"id,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Type        *string "json:\"type,omitempty\""
		UnreadCount *int    "json:\"unread_count,omitempty\""
	} "json:\"folders,omitempty\""
	Scans *[]struct {
		Control              *bool   "json:\"control,omitempty\""
		CreationDate         *int32  "json:\"creation_date,omitempty\""
		Enabled              *bool   "json:\"enabled,omitempty\""
		Id                   *int32  "json:\"id,omitempty\""
		LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
		Legacy               *bool   "json:\"legacy,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		Permissions          *int32  "json:\"permissions,omitempty\""
		PolicyId             *int    "json:\"policy_id,omitempty\""
		Read                 *bool   "json:\"read,omitempty\""
		Rrules               *string "json:\"rrules,omitempty\""
		ScheduleUuid         *string "json:\"schedule_uuid,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Starttime            *string "json:\"starttime,omitempty\""
		Status               *string "json:\"status,omitempty\""
		Timezone             *string "json:\"timezone,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int32  "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
		WizardUuid           *string "json:\"wizard_uuid,omitempty\""
	} "json:\"scans,omitempty\""
	Timestamp *int32 "json:\"timestamp,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansListWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Folders *[]struct { Custom *int "json:\"custom,omitempty\""; DefaultTag *int "json:\"default_tag,omitempty\""; Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Type *string "json:\"type,omitempty\""; UnreadCount *int "json:\"unread_count,omitempty\"" } "json:\"folders,omitempty\""; Scans *[]struct { Control *bool "json:\"control,omitempty\""; CreationDate *int32 "json:\"creation_date,omitempty\""; Enabled *bool "json:\"enabled,omitempty\""; Id *int32 "json:\"id,omitempty\""; LastModificationDate *int32 "json:\"last_modification_date,omitempty\""; Legacy *bool "json:\"legacy,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; Permissions *int32 "json:\"permissions,omitempty\""; PolicyId *int "json:\"policy_id,omitempty\""; Read *bool "json:\"read,omitempty\""; Rrules *string "json:\"rrules,omitempty\""; ScheduleUuid *string "json:\"schedule_uuid,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; Status *string "json:\"status,omitempty\""; Timezone *string "json:\"timezone,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int32 "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\""; WizardUuid *string "json:\"wizard_uuid,omitempty\"" } "json:\"scans,omitempty\""; Timestamp *int32 "json:\"timestamp,omitempty\"" }
	if i, ok := r.(*struct {
		Folders *[]struct {
			Custom      *int    "json:\"custom,omitempty\""
			DefaultTag  *int    "json:\"default_tag,omitempty\""
			Id          *int    "json:\"id,omitempty\""
			Name        *string "json:\"name,omitempty\""
			Type        *string "json:\"type,omitempty\""
			UnreadCount *int    "json:\"unread_count,omitempty\""
		} "json:\"folders,omitempty\""
		Scans *[]struct {
			Control              *bool   "json:\"control,omitempty\""
			CreationDate         *int32  "json:\"creation_date,omitempty\""
			Enabled              *bool   "json:\"enabled,omitempty\""
			Id                   *int32  "json:\"id,omitempty\""
			LastModificationDate *int32  "json:\"last_modification_date,omitempty\""
			Legacy               *bool   "json:\"legacy,omitempty\""
			Name                 *string "json:\"name,omitempty\""
			Owner                *string "json:\"owner,omitempty\""
			Permissions          *int32  "json:\"permissions,omitempty\""
			PolicyId             *int    "json:\"policy_id,omitempty\""
			Read                 *bool   "json:\"read,omitempty\""
			Rrules               *string "json:\"rrules,omitempty\""
			ScheduleUuid         *string "json:\"schedule_uuid,omitempty\""
			Shared               *bool   "json:\"shared,omitempty\""
			Starttime            *string "json:\"starttime,omitempty\""
			Status               *string "json:\"status,omitempty\""
			Timezone             *string "json:\"timezone,omitempty\""
			Type                 *string "json:\"type,omitempty\""
			UserPermissions      *int32  "json:\"user_permissions,omitempty\""
			Uuid                 *string "json:\"uuid,omitempty\""
			WizardUuid           *string "json:\"wizard_uuid,omitempty\""
		} "json:\"scans,omitempty\""
		Timestamp *int32 "json:\"timestamp,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Pause calls the Scans ´s function
func (c *Scans) Pause(contentType string, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansPauseWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// PluginOutput calls the Scans ´s function
func (c *Scans) PluginOutput(arg1 string, arg2 int32, arg3 int32, params *ScansPluginOutputParams, reqEditors ...RequestEditorFn) (*struct {
	Info *struct {
		HostFqdn   *string "json:\"host-fqdn,omitempty\""
		HostIp     *string "json:\"host-ip,omitempty\""
		HostUuid   *string "json:\"host-uuid,omitempty\""
		HostEnd    *string "json:\"host_end,omitempty\""
		HostFqdn2  *string "json:\"host_fqdn,omitempty\""
		HostStart  *string "json:\"host_start,omitempty\""
		MacAddress *string "json:\"mac-address,omitempty\""
	} "json:\"info,omitempty\""
	Output *[]struct {
		CustomDescription *string                 "json:\"custom_description,omitempty\""
		HasAttachment     *int                    "json:\"has_attachment,omitempty\""
		Hosts             *string                 "json:\"hosts,omitempty\""
		PluginOutput      *string                 "json:\"plugin_output,omitempty\""
		Ports             *map[string]interface{} "json:\"ports,omitempty\""
		Severity          *int                    "json:\"severity,omitempty\""
	} "json:\"output,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansPluginOutputWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Info *struct { HostFqdn *string "json:\"host-fqdn,omitempty\""; HostIp *string "json:\"host-ip,omitempty\""; HostUuid *string "json:\"host-uuid,omitempty\""; HostEnd *string "json:\"host_end,omitempty\""; HostFqdn2 *string "json:\"host_fqdn,omitempty\""; HostStart *string "json:\"host_start,omitempty\""; MacAddress *string "json:\"mac-address,omitempty\"" } "json:\"info,omitempty\""; Output *[]struct { CustomDescription *string "json:\"custom_description,omitempty\""; HasAttachment *int "json:\"has_attachment,omitempty\""; Hosts *string "json:\"hosts,omitempty\""; PluginOutput *string "json:\"plugin_output,omitempty\""; Ports *map[string]interface {} "json:\"ports,omitempty\""; Severity *int "json:\"severity,omitempty\"" } "json:\"output,omitempty\"" }
	if i, ok := r.(*struct {
		Info *struct {
			HostFqdn   *string "json:\"host-fqdn,omitempty\""
			HostIp     *string "json:\"host-ip,omitempty\""
			HostUuid   *string "json:\"host-uuid,omitempty\""
			HostEnd    *string "json:\"host_end,omitempty\""
			HostFqdn2  *string "json:\"host_fqdn,omitempty\""
			HostStart  *string "json:\"host_start,omitempty\""
			MacAddress *string "json:\"mac-address,omitempty\""
		} "json:\"info,omitempty\""
		Output *[]struct {
			CustomDescription *string                 "json:\"custom_description,omitempty\""
			HasAttachment     *int                    "json:\"has_attachment,omitempty\""
			Hosts             *string                 "json:\"hosts,omitempty\""
			PluginOutput      *string                 "json:\"plugin_output,omitempty\""
			Ports             *map[string]interface{} "json:\"ports,omitempty\""
			Severity          *int                    "json:\"severity,omitempty\""
		} "json:\"output,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ReadStatusWithBody calls the Scans ´s function
func (c *Scans) ReadStatusWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansReadStatusWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ReadStatus calls the Scans ´s function
func (c *Scans) ReadStatus(arg1 string, arg2 ScansReadStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansReadStatusWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Resume calls the Scans ´s function
func (c *Scans) Resume(contentType string, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansResumeWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ScheduleWithBody calls the Scans ´s function
func (c *Scans) ScheduleWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	Control   *bool   "json:\"control,omitempty\""
	Enabled   *bool   "json:\"enabled,omitempty\""
	Rrules    *string "json:\"rrules,omitempty\""
	Starttime *string "json:\"starttime,omitempty\""
	Timezone  *string "json:\"timezone,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansScheduleWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Control *bool "json:\"control,omitempty\""; Enabled *bool "json:\"enabled,omitempty\""; Rrules *string "json:\"rrules,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; Timezone *string "json:\"timezone,omitempty\"" }
	if i, ok := r.(*struct {
		Control   *bool   "json:\"control,omitempty\""
		Enabled   *bool   "json:\"enabled,omitempty\""
		Rrules    *string "json:\"rrules,omitempty\""
		Starttime *string "json:\"starttime,omitempty\""
		Timezone  *string "json:\"timezone,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Schedule calls the Scans ´s function
func (c *Scans) Schedule(arg1 string, arg2 ScansScheduleJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	Control   *bool   "json:\"control,omitempty\""
	Enabled   *bool   "json:\"enabled,omitempty\""
	Rrules    *string "json:\"rrules,omitempty\""
	Starttime *string "json:\"starttime,omitempty\""
	Timezone  *string "json:\"timezone,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansScheduleWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Control *bool "json:\"control,omitempty\""; Enabled *bool "json:\"enabled,omitempty\""; Rrules *string "json:\"rrules,omitempty\""; Starttime *string "json:\"starttime,omitempty\""; Timezone *string "json:\"timezone,omitempty\"" }
	if i, ok := r.(*struct {
		Control   *bool   "json:\"control,omitempty\""
		Enabled   *bool   "json:\"enabled,omitempty\""
		Rrules    *string "json:\"rrules,omitempty\""
		Starttime *string "json:\"starttime,omitempty\""
		Timezone  *string "json:\"timezone,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Stop calls the Scans ´s function
func (c *Scans) Stop(contentType string, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansStopWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Timezones calls the Scans ´s function
func (c *Scans) Timezones(reqEditors ...RequestEditorFn) (*[]struct {
	Name  *string "json:\"name,omitempty\""
	Value *string "json:\"value,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScansTimezonesWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { Name *string "json:\"name,omitempty\""; Value *string "json:\"value,omitempty\"" }
	if i, ok := r.(*[]struct {
		Name  *string "json:\"name,omitempty\""
		Value *string "json:\"value,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// TargetGroups implementation of the TargetGroups interface
type TargetGroups struct {
	*ClientWithResponses
}

// CreateWithBody calls the TargetGroups ´s function
func (c *TargetGroups) CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	Acls *[]struct {
		DisplayName *string       "json:\"display_name,omitempty\""
		Id          *int          "json:\"id,omitempty\""
		Name        *string       "json:\"name,omitempty\""
		Owner       *int          "json:\"owner,omitempty\""
		Permissions *int32        "json:\"permissions,omitempty\""
		Type        *N200AclsType "json:\"type,omitempty\""
		Uuid        *string       "json:\"uuid,omitempty\""
	} "json:\"acls,omitempty\""
	DefaultGroup         *bool   "json:\"default_group,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Members              *string "json:\"members,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	OwnerId              *int    "json:\"owner_id,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TargetGroupsCreateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Acls *[]struct { DisplayName *string "json:\"display_name,omitempty\""; Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *int "json:\"owner,omitempty\""; Permissions *int32 "json:\"permissions,omitempty\""; Type *N200AclsType "json:\"type,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"acls,omitempty\""; DefaultGroup *bool "json:\"default_group,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Members *string "json:\"members,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\"" }
	if i, ok := r.(*struct {
		Acls *[]struct {
			DisplayName *string       "json:\"display_name,omitempty\""
			Id          *int          "json:\"id,omitempty\""
			Name        *string       "json:\"name,omitempty\""
			Owner       *int          "json:\"owner,omitempty\""
			Permissions *int32        "json:\"permissions,omitempty\""
			Type        *N200AclsType "json:\"type,omitempty\""
			Uuid        *string       "json:\"uuid,omitempty\""
		} "json:\"acls,omitempty\""
		DefaultGroup         *bool   "json:\"default_group,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Members              *string "json:\"members,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Create calls the TargetGroups ´s function
func (c *TargetGroups) Create(arg1 TargetGroupsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	Acls *[]struct {
		DisplayName *string       "json:\"display_name,omitempty\""
		Id          *int          "json:\"id,omitempty\""
		Name        *string       "json:\"name,omitempty\""
		Owner       *int          "json:\"owner,omitempty\""
		Permissions *int32        "json:\"permissions,omitempty\""
		Type        *N200AclsType "json:\"type,omitempty\""
		Uuid        *string       "json:\"uuid,omitempty\""
	} "json:\"acls,omitempty\""
	DefaultGroup         *bool   "json:\"default_group,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Members              *string "json:\"members,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	OwnerId              *int    "json:\"owner_id,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TargetGroupsCreateWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Acls *[]struct { DisplayName *string "json:\"display_name,omitempty\""; Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *int "json:\"owner,omitempty\""; Permissions *int32 "json:\"permissions,omitempty\""; Type *N200AclsType "json:\"type,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"acls,omitempty\""; DefaultGroup *bool "json:\"default_group,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Members *string "json:\"members,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\"" }
	if i, ok := r.(*struct {
		Acls *[]struct {
			DisplayName *string       "json:\"display_name,omitempty\""
			Id          *int          "json:\"id,omitempty\""
			Name        *string       "json:\"name,omitempty\""
			Owner       *int          "json:\"owner,omitempty\""
			Permissions *int32        "json:\"permissions,omitempty\""
			Type        *N200AclsType "json:\"type,omitempty\""
			Uuid        *string       "json:\"uuid,omitempty\""
		} "json:\"acls,omitempty\""
		DefaultGroup         *bool   "json:\"default_group,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Members              *string "json:\"members,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Delete calls the TargetGroups ´s function
func (c *TargetGroups) Delete(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TargetGroupsDeleteWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Details calls the TargetGroups ´s function
func (c *TargetGroups) Details(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
	Acls *[]struct {
		DisplayName *string       "json:\"display_name,omitempty\""
		Id          *int          "json:\"id,omitempty\""
		Name        *string       "json:\"name,omitempty\""
		Owner       *int          "json:\"owner,omitempty\""
		Permissions *int32        "json:\"permissions,omitempty\""
		Type        *N200AclsType "json:\"type,omitempty\""
		Uuid        *string       "json:\"uuid,omitempty\""
	} "json:\"acls,omitempty\""
	DefaultGroup         *bool   "json:\"default_group,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Members              *string "json:\"members,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	OwnerId              *int    "json:\"owner_id,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TargetGroupsDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Acls *[]struct { DisplayName *string "json:\"display_name,omitempty\""; Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *int "json:\"owner,omitempty\""; Permissions *int32 "json:\"permissions,omitempty\""; Type *N200AclsType "json:\"type,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"acls,omitempty\""; DefaultGroup *bool "json:\"default_group,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Members *string "json:\"members,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\"" }
	if i, ok := r.(*struct {
		Acls *[]struct {
			DisplayName *string       "json:\"display_name,omitempty\""
			Id          *int          "json:\"id,omitempty\""
			Name        *string       "json:\"name,omitempty\""
			Owner       *int          "json:\"owner,omitempty\""
			Permissions *int32        "json:\"permissions,omitempty\""
			Type        *N200AclsType "json:\"type,omitempty\""
			Uuid        *string       "json:\"uuid,omitempty\""
		} "json:\"acls,omitempty\""
		DefaultGroup         *bool   "json:\"default_group,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Members              *string "json:\"members,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// EditWithBody calls the TargetGroups ´s function
func (c *TargetGroups) EditWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	Acls *[]struct {
		DisplayName *string       "json:\"display_name,omitempty\""
		Id          *int          "json:\"id,omitempty\""
		Name        *string       "json:\"name,omitempty\""
		Owner       *int          "json:\"owner,omitempty\""
		Permissions *int32        "json:\"permissions,omitempty\""
		Type        *N200AclsType "json:\"type,omitempty\""
		Uuid        *string       "json:\"uuid,omitempty\""
	} "json:\"acls,omitempty\""
	DefaultGroup         *bool   "json:\"default_group,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Members              *string "json:\"members,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	OwnerId              *int    "json:\"owner_id,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TargetGroupsEditWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Acls *[]struct { DisplayName *string "json:\"display_name,omitempty\""; Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *int "json:\"owner,omitempty\""; Permissions *int32 "json:\"permissions,omitempty\""; Type *N200AclsType "json:\"type,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"acls,omitempty\""; DefaultGroup *bool "json:\"default_group,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Members *string "json:\"members,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\"" }
	if i, ok := r.(*struct {
		Acls *[]struct {
			DisplayName *string       "json:\"display_name,omitempty\""
			Id          *int          "json:\"id,omitempty\""
			Name        *string       "json:\"name,omitempty\""
			Owner       *int          "json:\"owner,omitempty\""
			Permissions *int32        "json:\"permissions,omitempty\""
			Type        *N200AclsType "json:\"type,omitempty\""
			Uuid        *string       "json:\"uuid,omitempty\""
		} "json:\"acls,omitempty\""
		DefaultGroup         *bool   "json:\"default_group,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Members              *string "json:\"members,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Edit calls the TargetGroups ´s function
func (c *TargetGroups) Edit(arg1 int32, arg2 TargetGroupsEditJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	Acls *[]struct {
		DisplayName *string       "json:\"display_name,omitempty\""
		Id          *int          "json:\"id,omitempty\""
		Name        *string       "json:\"name,omitempty\""
		Owner       *int          "json:\"owner,omitempty\""
		Permissions *int32        "json:\"permissions,omitempty\""
		Type        *N200AclsType "json:\"type,omitempty\""
		Uuid        *string       "json:\"uuid,omitempty\""
	} "json:\"acls,omitempty\""
	DefaultGroup         *bool   "json:\"default_group,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Members              *string "json:\"members,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	OwnerId              *int    "json:\"owner_id,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TargetGroupsEditWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Acls *[]struct { DisplayName *string "json:\"display_name,omitempty\""; Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *int "json:\"owner,omitempty\""; Permissions *int32 "json:\"permissions,omitempty\""; Type *N200AclsType "json:\"type,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"acls,omitempty\""; DefaultGroup *bool "json:\"default_group,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Members *string "json:\"members,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\"" }
	if i, ok := r.(*struct {
		Acls *[]struct {
			DisplayName *string       "json:\"display_name,omitempty\""
			Id          *int          "json:\"id,omitempty\""
			Name        *string       "json:\"name,omitempty\""
			Owner       *int          "json:\"owner,omitempty\""
			Permissions *int32        "json:\"permissions,omitempty\""
			Type        *N200AclsType "json:\"type,omitempty\""
			Uuid        *string       "json:\"uuid,omitempty\""
		} "json:\"acls,omitempty\""
		DefaultGroup         *bool   "json:\"default_group,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Members              *string "json:\"members,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// List calls the TargetGroups ´s function
func (c *TargetGroups) List(reqEditors ...RequestEditorFn) (*[]struct {
	Acls *[]struct {
		DisplayName *string       "json:\"display_name,omitempty\""
		Id          *int          "json:\"id,omitempty\""
		Name        *string       "json:\"name,omitempty\""
		Owner       *int          "json:\"owner,omitempty\""
		Permissions *int32        "json:\"permissions,omitempty\""
		Type        *N200AclsType "json:\"type,omitempty\""
		Uuid        *string       "json:\"uuid,omitempty\""
	} "json:\"acls,omitempty\""
	DefaultGroup         *bool   "json:\"default_group,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Members              *string "json:\"members,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	OwnerId              *int    "json:\"owner_id,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TargetGroupsListWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { Acls *[]struct { DisplayName *string "json:\"display_name,omitempty\""; Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *int "json:\"owner,omitempty\""; Permissions *int32 "json:\"permissions,omitempty\""; Type *N200AclsType "json:\"type,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"acls,omitempty\""; DefaultGroup *bool "json:\"default_group,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Members *string "json:\"members,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\"" }
	if i, ok := r.(*[]struct {
		Acls *[]struct {
			DisplayName *string       "json:\"display_name,omitempty\""
			Id          *int          "json:\"id,omitempty\""
			Name        *string       "json:\"name,omitempty\""
			Owner       *int          "json:\"owner,omitempty\""
			Permissions *int32        "json:\"permissions,omitempty\""
			Type        *N200AclsType "json:\"type,omitempty\""
			Uuid        *string       "json:\"uuid,omitempty\""
		} "json:\"acls,omitempty\""
		DefaultGroup         *bool   "json:\"default_group,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Members              *string "json:\"members,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Workbenches implementation of the Workbenches interface
type Workbenches struct {
	*ClientWithResponses
}

// AssetInfo calls the Workbenches ´s function
func (c *Workbenches) AssetInfo(arg1 string, params *WorkbenchesAssetInfoParams, reqEditors ...RequestEditorFn) (*struct {
	Info *struct {
		AcrDrivers *[]struct {
			DriverName  *string   "json:\"driver_name,omitempty\""
			DriverValue *[]string "json:\"driver_value,omitempty\""
		} "json:\"acr_drivers,omitempty\""
		AcrScore                  *int                    "json:\"acr_score,omitempty\""
		AgentName                 *[]string               "json:\"agent_name,omitempty\""
		AwsAvailabilityZone       *[]string               "json:\"aws_availability_zone,omitempty\""
		AwsEc2InstanceAmiId       *[]string               "json:\"aws_ec2_instance_ami_id,omitempty\""
		AwsEc2InstanceGroupName   *[]string               "json:\"aws_ec2_instance_group_name,omitempty\""
		AwsEc2InstanceId          *[]string               "json:\"aws_ec2_instance_id,omitempty\""
		AwsEc2InstanceStateName   *[]string               "json:\"aws_ec2_instance_state_name,omitempty\""
		AwsEc2InstanceType        *[]string               "json:\"aws_ec2_instance_type,omitempty\""
		AwsEc2Name                *[]string               "json:\"aws_ec2_name,omitempty\""
		AwsEc2ProductCode         *[]string               "json:\"aws_ec2_product_code,omitempty\""
		AwsOwnerId                *[]string               "json:\"aws_owner_id,omitempty\""
		AwsRegion                 *[]string               "json:\"aws_region,omitempty\""
		AwsSubnetId               *[]string               "json:\"aws_subnet_id,omitempty\""
		AwsVpcId                  *[]string               "json:\"aws_vpc_id,omitempty\""
		AzureResourceId           *[]string               "json:\"azure_resource_id,omitempty\""
		AzureVmId                 *[]string               "json:\"azure_vm_id,omitempty\""
		BiosUuid                  *[]string               "json:\"bios_uuid,omitempty\""
		Counts                    *map[string]interface{} "json:\"counts,omitempty\""
		CreatedAt                 *string                 "json:\"created_at,omitempty\""
		ExposureScore             *int                    "json:\"exposure_score,omitempty\""
		FirstSeen                 *string                 "json:\"first_seen,omitempty\""
		Fqdn                      *[]string               "json:\"fqdn,omitempty\""
		GcpInstanceId             *[]string               "json:\"gcp_instance_id,omitempty\""
		GcpProjectId              *[]string               "json:\"gcp_project_id,omitempty\""
		GcpZone                   *[]string               "json:\"gcp_zone,omitempty\""
		HasAgent                  *bool                   "json:\"has_agent,omitempty\""
		Hostname                  *[]string               "json:\"hostname,omitempty\""
		Id                        *string                 "json:\"id,omitempty\""
		InstalledSoftware         *[]string               "json:\"installed_software,omitempty\""
		Ipv4                      *[]string               "json:\"ipv4,omitempty\""
		Ipv6                      *[]string               "json:\"ipv6,omitempty\""
		LastAuthenticatedScanDate *string                 "json:\"last_authenticated_scan_date,omitempty\""
		LastLicensedScanDate      *string                 "json:\"last_licensed_scan_date,omitempty\""
		LastScanTarget            *string                 "json:\"last_scan_target,omitempty\""
		LastSeen                  *string                 "json:\"last_seen,omitempty\""
		MacAddress                *[]string               "json:\"mac_address,omitempty\""
		McafeeEpoAgentGuid        *[]string               "json:\"mcafee_epo_agent_guid,omitempty\""
		McafeeEpoGuid             *[]string               "json:\"mcafee_epo_guid,omitempty\""
		NetbiosName               *[]string               "json:\"netbios_name,omitempty\""
		OperatingSystem           *[]string               "json:\"operating_system,omitempty\""
		QualysAssetId             *[]string               "json:\"qualys_asset_id,omitempty\""
		QualysHostId              *[]string               "json:\"qualys_host_id,omitempty\""
		ScanFrequency             *[]struct {
			Frequency *int  "json:\"frequency,omitempty\""
			Interval  *int  "json:\"interval,omitempty\""
			Licensed  *bool "json:\"licensed,omitempty\""
		} "json:\"scan_frequency,omitempty\""
		ServicenowSysid *[]string "json:\"servicenow_sysid,omitempty\""
		Sources         *[]struct {
			FirstSeen *string "json:\"first_seen,omitempty\""
			LastSeen  *string "json:\"last_seen,omitempty\""
			Name      *string "json:\"name,omitempty\""
		} "json:\"sources,omitempty\""
		SshFingerprint *[]string "json:\"ssh_fingerprint,omitempty\""
		SystemType     *[]string "json:\"system_type,omitempty\""
		Tags           *[]struct {
			AddedAt  *string "json:\"added_at,omitempty\""
			AddedBy  *string "json:\"added_by,omitempty\""
			Source   *string "json:\"source,omitempty\""
			TagKey   *string "json:\"tag_key,omitempty\""
			TagUuid  *string "json:\"tag_uuid,omitempty\""
			TagValue *string "json:\"tag_value,omitempty\""
		} "json:\"tags,omitempty\""
		TenableUuid *[]string "json:\"tenable_uuid,omitempty\""
		TimeEnd     *string   "json:\"time_end,omitempty\""
		TimeStart   *string   "json:\"time_start,omitempty\""
		UpdatedAt   *string   "json:\"updated_at,omitempty\""
		Uuid        *string   "json:\"uuid,omitempty\""
	} "json:\"info,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.WorkbenchesAssetInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Info *struct { AcrDrivers *[]struct { DriverName *string "json:\"driver_name,omitempty\""; DriverValue *[]string "json:\"driver_value,omitempty\"" } "json:\"acr_drivers,omitempty\""; AcrScore *int "json:\"acr_score,omitempty\""; AgentName *[]string "json:\"agent_name,omitempty\""; AwsAvailabilityZone *[]string "json:\"aws_availability_zone,omitempty\""; AwsEc2InstanceAmiId *[]string "json:\"aws_ec2_instance_ami_id,omitempty\""; AwsEc2InstanceGroupName *[]string "json:\"aws_ec2_instance_group_name,omitempty\""; AwsEc2InstanceId *[]string "json:\"aws_ec2_instance_id,omitempty\""; AwsEc2InstanceStateName *[]string "json:\"aws_ec2_instance_state_name,omitempty\""; AwsEc2InstanceType *[]string "json:\"aws_ec2_instance_type,omitempty\""; AwsEc2Name *[]string "json:\"aws_ec2_name,omitempty\""; AwsEc2ProductCode *[]string "json:\"aws_ec2_product_code,omitempty\""; AwsOwnerId *[]string "json:\"aws_owner_id,omitempty\""; AwsRegion *[]string "json:\"aws_region,omitempty\""; AwsSubnetId *[]string "json:\"aws_subnet_id,omitempty\""; AwsVpcId *[]string "json:\"aws_vpc_id,omitempty\""; AzureResourceId *[]string "json:\"azure_resource_id,omitempty\""; AzureVmId *[]string "json:\"azure_vm_id,omitempty\""; BiosUuid *[]string "json:\"bios_uuid,omitempty\""; Counts *map[string]interface {} "json:\"counts,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; ExposureScore *int "json:\"exposure_score,omitempty\""; FirstSeen *string "json:\"first_seen,omitempty\""; Fqdn *[]string "json:\"fqdn,omitempty\""; GcpInstanceId *[]string "json:\"gcp_instance_id,omitempty\""; GcpProjectId *[]string "json:\"gcp_project_id,omitempty\""; GcpZone *[]string "json:\"gcp_zone,omitempty\""; HasAgent *bool "json:\"has_agent,omitempty\""; Hostname *[]string "json:\"hostname,omitempty\""; Id *string "json:\"id,omitempty\""; InstalledSoftware *[]string "json:\"installed_software,omitempty\""; Ipv4 *[]string "json:\"ipv4,omitempty\""; Ipv6 *[]string "json:\"ipv6,omitempty\""; LastAuthenticatedScanDate *string "json:\"last_authenticated_scan_date,omitempty\""; LastLicensedScanDate *string "json:\"last_licensed_scan_date,omitempty\""; LastScanTarget *string "json:\"last_scan_target,omitempty\""; LastSeen *string "json:\"last_seen,omitempty\""; MacAddress *[]string "json:\"mac_address,omitempty\""; McafeeEpoAgentGuid *[]string "json:\"mcafee_epo_agent_guid,omitempty\""; McafeeEpoGuid *[]string "json:\"mcafee_epo_guid,omitempty\""; NetbiosName *[]string "json:\"netbios_name,omitempty\""; OperatingSystem *[]string "json:\"operating_system,omitempty\""; QualysAssetId *[]string "json:\"qualys_asset_id,omitempty\""; QualysHostId *[]string "json:\"qualys_host_id,omitempty\""; ScanFrequency *[]struct { Frequency *int "json:\"frequency,omitempty\""; Interval *int "json:\"interval,omitempty\""; Licensed *bool "json:\"licensed,omitempty\"" } "json:\"scan_frequency,omitempty\""; ServicenowSysid *[]string "json:\"servicenow_sysid,omitempty\""; Sources *[]struct { FirstSeen *string "json:\"first_seen,omitempty\""; LastSeen *string "json:\"last_seen,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"sources,omitempty\""; SshFingerprint *[]string "json:\"ssh_fingerprint,omitempty\""; SystemType *[]string "json:\"system_type,omitempty\""; Tags *[]struct { AddedAt *string "json:\"added_at,omitempty\""; AddedBy *string "json:\"added_by,omitempty\""; Source *string "json:\"source,omitempty\""; TagKey *string "json:\"tag_key,omitempty\""; TagUuid *string "json:\"tag_uuid,omitempty\""; TagValue *string "json:\"tag_value,omitempty\"" } "json:\"tags,omitempty\""; TenableUuid *[]string "json:\"tenable_uuid,omitempty\""; TimeEnd *string "json:\"time_end,omitempty\""; TimeStart *string "json:\"time_start,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"info,omitempty\"" }
	if i, ok := r.(*struct {
		Info *struct {
			AcrDrivers *[]struct {
				DriverName  *string   "json:\"driver_name,omitempty\""
				DriverValue *[]string "json:\"driver_value,omitempty\""
			} "json:\"acr_drivers,omitempty\""
			AcrScore                  *int                    "json:\"acr_score,omitempty\""
			AgentName                 *[]string               "json:\"agent_name,omitempty\""
			AwsAvailabilityZone       *[]string               "json:\"aws_availability_zone,omitempty\""
			AwsEc2InstanceAmiId       *[]string               "json:\"aws_ec2_instance_ami_id,omitempty\""
			AwsEc2InstanceGroupName   *[]string               "json:\"aws_ec2_instance_group_name,omitempty\""
			AwsEc2InstanceId          *[]string               "json:\"aws_ec2_instance_id,omitempty\""
			AwsEc2InstanceStateName   *[]string               "json:\"aws_ec2_instance_state_name,omitempty\""
			AwsEc2InstanceType        *[]string               "json:\"aws_ec2_instance_type,omitempty\""
			AwsEc2Name                *[]string               "json:\"aws_ec2_name,omitempty\""
			AwsEc2ProductCode         *[]string               "json:\"aws_ec2_product_code,omitempty\""
			AwsOwnerId                *[]string               "json:\"aws_owner_id,omitempty\""
			AwsRegion                 *[]string               "json:\"aws_region,omitempty\""
			AwsSubnetId               *[]string               "json:\"aws_subnet_id,omitempty\""
			AwsVpcId                  *[]string               "json:\"aws_vpc_id,omitempty\""
			AzureResourceId           *[]string               "json:\"azure_resource_id,omitempty\""
			AzureVmId                 *[]string               "json:\"azure_vm_id,omitempty\""
			BiosUuid                  *[]string               "json:\"bios_uuid,omitempty\""
			Counts                    *map[string]interface{} "json:\"counts,omitempty\""
			CreatedAt                 *string                 "json:\"created_at,omitempty\""
			ExposureScore             *int                    "json:\"exposure_score,omitempty\""
			FirstSeen                 *string                 "json:\"first_seen,omitempty\""
			Fqdn                      *[]string               "json:\"fqdn,omitempty\""
			GcpInstanceId             *[]string               "json:\"gcp_instance_id,omitempty\""
			GcpProjectId              *[]string               "json:\"gcp_project_id,omitempty\""
			GcpZone                   *[]string               "json:\"gcp_zone,omitempty\""
			HasAgent                  *bool                   "json:\"has_agent,omitempty\""
			Hostname                  *[]string               "json:\"hostname,omitempty\""
			Id                        *string                 "json:\"id,omitempty\""
			InstalledSoftware         *[]string               "json:\"installed_software,omitempty\""
			Ipv4                      *[]string               "json:\"ipv4,omitempty\""
			Ipv6                      *[]string               "json:\"ipv6,omitempty\""
			LastAuthenticatedScanDate *string                 "json:\"last_authenticated_scan_date,omitempty\""
			LastLicensedScanDate      *string                 "json:\"last_licensed_scan_date,omitempty\""
			LastScanTarget            *string                 "json:\"last_scan_target,omitempty\""
			LastSeen                  *string                 "json:\"last_seen,omitempty\""
			MacAddress                *[]string               "json:\"mac_address,omitempty\""
			McafeeEpoAgentGuid        *[]string               "json:\"mcafee_epo_agent_guid,omitempty\""
			McafeeEpoGuid             *[]string               "json:\"mcafee_epo_guid,omitempty\""
			NetbiosName               *[]string               "json:\"netbios_name,omitempty\""
			OperatingSystem           *[]string               "json:\"operating_system,omitempty\""
			QualysAssetId             *[]string               "json:\"qualys_asset_id,omitempty\""
			QualysHostId              *[]string               "json:\"qualys_host_id,omitempty\""
			ScanFrequency             *[]struct {
				Frequency *int  "json:\"frequency,omitempty\""
				Interval  *int  "json:\"interval,omitempty\""
				Licensed  *bool "json:\"licensed,omitempty\""
			} "json:\"scan_frequency,omitempty\""
			ServicenowSysid *[]string "json:\"servicenow_sysid,omitempty\""
			Sources         *[]struct {
				FirstSeen *string "json:\"first_seen,omitempty\""
				LastSeen  *string "json:\"last_seen,omitempty\""
				Name      *string "json:\"name,omitempty\""
			} "json:\"sources,omitempty\""
			SshFingerprint *[]string "json:\"ssh_fingerprint,omitempty\""
			SystemType     *[]string "json:\"system_type,omitempty\""
			Tags           *[]struct {
				AddedAt  *string "json:\"added_at,omitempty\""
				AddedBy  *string "json:\"added_by,omitempty\""
				Source   *string "json:\"source,omitempty\""
				TagKey   *string "json:\"tag_key,omitempty\""
				TagUuid  *string "json:\"tag_uuid,omitempty\""
				TagValue *string "json:\"tag_value,omitempty\""
			} "json:\"tags,omitempty\""
			TenableUuid *[]string "json:\"tenable_uuid,omitempty\""
			TimeEnd     *string   "json:\"time_end,omitempty\""
			TimeStart   *string   "json:\"time_start,omitempty\""
			UpdatedAt   *string   "json:\"updated_at,omitempty\""
			Uuid        *string   "json:\"uuid,omitempty\""
		} "json:\"info,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AssetVulnerabilities calls the Workbenches ´s function
func (c *Workbenches) AssetVulnerabilities(arg1 string, params *WorkbenchesAssetVulnerabilitiesParams, reqEditors ...RequestEditorFn) (*struct {
	TotalAssetCount         *int "json:\"total_asset_count,omitempty\""
	TotalVulnerabilityCount *int "json:\"total_vulnerability_count,omitempty\""
	Vulnerabilities         *[]struct {
		AcceptedCount    *int32 "json:\"accepted_count,omitempty\""
		Count            *int   "json:\"count,omitempty\""
		CountsBySeverity *[]struct {
			Count *int32 "json:\"count,omitempty\""
			Value *int   "json:\"value,omitempty\""
		} "json:\"counts by severity,omitempty\""
		PluginFamily       *string  "json:\"plugin_family,omitempty\""
		PluginId           *int     "json:\"plugin_id,omitempty\""
		PluginName         *string  "json:\"plugin_name,omitempty\""
		RecastedCount      *int32   "json:\"recasted_count,omitempty\""
		Severity           *int32   "json:\"severity,omitempty\""
		VprScore           *float32 "json:\"vpr_score,omitempty\""
		VulnerabilityState *string  "json:\"vulnerability_state,omitempty\""
	} "json:\"vulnerabilities,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.WorkbenchesAssetVulnerabilitiesWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { TotalAssetCount *int "json:\"total_asset_count,omitempty\""; TotalVulnerabilityCount *int "json:\"total_vulnerability_count,omitempty\""; Vulnerabilities *[]struct { AcceptedCount *int32 "json:\"accepted_count,omitempty\""; Count *int "json:\"count,omitempty\""; CountsBySeverity *[]struct { Count *int32 "json:\"count,omitempty\""; Value *int "json:\"value,omitempty\"" } "json:\"counts by severity,omitempty\""; PluginFamily *string "json:\"plugin_family,omitempty\""; PluginId *int "json:\"plugin_id,omitempty\""; PluginName *string "json:\"plugin_name,omitempty\""; RecastedCount *int32 "json:\"recasted_count,omitempty\""; Severity *int32 "json:\"severity,omitempty\""; VprScore *float32 "json:\"vpr_score,omitempty\""; VulnerabilityState *string "json:\"vulnerability_state,omitempty\"" } "json:\"vulnerabilities,omitempty\"" }
	if i, ok := r.(*struct {
		TotalAssetCount         *int "json:\"total_asset_count,omitempty\""
		TotalVulnerabilityCount *int "json:\"total_vulnerability_count,omitempty\""
		Vulnerabilities         *[]struct {
			AcceptedCount    *int32 "json:\"accepted_count,omitempty\""
			Count            *int   "json:\"count,omitempty\""
			CountsBySeverity *[]struct {
				Count *int32 "json:\"count,omitempty\""
				Value *int   "json:\"value,omitempty\""
			} "json:\"counts by severity,omitempty\""
			PluginFamily       *string  "json:\"plugin_family,omitempty\""
			PluginId           *int     "json:\"plugin_id,omitempty\""
			PluginName         *string  "json:\"plugin_name,omitempty\""
			RecastedCount      *int32   "json:\"recasted_count,omitempty\""
			Severity           *int32   "json:\"severity,omitempty\""
			VprScore           *float32 "json:\"vpr_score,omitempty\""
			VulnerabilityState *string  "json:\"vulnerability_state,omitempty\""
		} "json:\"vulnerabilities,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AssetVulnerabilityInfo calls the Workbenches ´s function
func (c *Workbenches) AssetVulnerabilityInfo(arg1 string, arg2 int32, params *WorkbenchesAssetVulnerabilityInfoParams, reqEditors ...RequestEditorFn) (*struct {
	Info *struct {
		Count       *int32  "json:\"count,omitempty\""
		Description *string "json:\"description,omitempty\""
		Discovery   *struct {
			SeenFirst *string "json:\"seen_first,omitempty\""
			SeenLast  *string "json:\"seen_last,omitempty\""
		} "json:\"discovery,omitempty\""
		PluginDetails *struct {
			Family           *string "json:\"family,omitempty\""
			ModificationDate *string "json:\"modification_date,omitempty\""
			Name             *string "json:\"name,omitempty\""
			PublicationDate  *string "json:\"publication_date,omitempty\""
			Severity         *int32  "json:\"severity,omitempty\""
			Type             *string "json:\"type,omitempty\""
			Version          *string "json:\"version,omitempty\""
		} "json:\"plugin_details,omitempty\""
		ReferenceInformation *[]struct {
			Name   *string   "json:\"name,omitempty\""
			Url    *string   "json:\"url,omitempty\""
			Values *[]string "json:\"values,omitempty\""
		} "json:\"reference_information,omitempty\""
		RiskInformation *struct {
			Cvss3BaseScore      *string "json:\"cvss3_base_score,omitempty\""
			Cvss3TemporalScore  *string "json:\"cvss3_temporal_score,omitempty\""
			Cvss3TemporalVector *string "json:\"cvss3_temporal_vector,omitempty\""
			Cvss3Vector         *string "json:\"cvss3_vector,omitempty\""
			CvssBaseScore       *string "json:\"cvss_base_score,omitempty\""
			CvssTemporalScore   *string "json:\"cvss_temporal_score,omitempty\""
			CvssTemporalVector  *string "json:\"cvss_temporal_vector,omitempty\""
			CvssVector          *string "json:\"cvss_vector,omitempty\""
			RiskFactor          *string "json:\"risk_factor,omitempty\""
			StigSeverity        *string "json:\"stig_severity,omitempty\""
		} "json:\"risk_information,omitempty\""
		SeeAlso  *[]string "json:\"see_also,omitempty\""
		Severity *int      "json:\"severity,omitempty\""
		Solution *string   "json:\"solution,omitempty\""
		Synopsis *string   "json:\"synopsis,omitempty\""
		Vpr      *struct {
			Drivers *map[string]interface{} "json:\"drivers,omitempty\""
			Score   *float32                "json:\"score,omitempty\""
			Updated *string                 "json:\"updated,omitempty\""
		} "json:\"vpr,omitempty\""
		VulnCount                *int32 "json:\"vuln_count,omitempty\""
		VulnerabilityInformation *struct {
			AssetInventory    *string "json:\"asset_inventory,omitempty\""
			Cpe               *string "json:\"cpe,omitempty\""
			DefaultAccount    *string "json:\"default_account,omitempty\""
			ExploitAvailable  *bool   "json:\"exploit_available,omitempty\""
			ExploitFrameworks *[]struct {
				Exploits *[]struct {
					Name *string "json:\"name,omitempty\""
					Url  *string "json:\"url,omitempty\""
				} "json:\"exploits,omitempty\""
				Name *string "json:\"name,omitempty\""
			} "json:\"exploit_frameworks,omitempty\""
			ExploitabilityEase           *string "json:\"exploitability_ease,omitempty\""
			ExploitedByMalware           *bool   "json:\"exploited_by_malware,omitempty\""
			ExploitedByNessus            *bool   "json:\"exploited_by_nessus,omitempty\""
			InTheNews                    *bool   "json:\"in_the_news,omitempty\""
			Malware                      *string "json:\"malware,omitempty\""
			PatchPublicationDate         *string "json:\"patch_publication_date,omitempty\""
			UnsupportedByVendor          *bool   "json:\"unsupported_by_vendor,omitempty\""
			VulnerabilityPublicationDate *string "json:\"vulnerability_publication_date,omitempty\""
		} "json:\"vulnerability_information,omitempty\""
	} "json:\"info,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.WorkbenchesAssetVulnerabilityInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Info *struct { Count *int32 "json:\"count,omitempty\""; Description *string "json:\"description,omitempty\""; Discovery *struct { SeenFirst *string "json:\"seen_first,omitempty\""; SeenLast *string "json:\"seen_last,omitempty\"" } "json:\"discovery,omitempty\""; PluginDetails *struct { Family *string "json:\"family,omitempty\""; ModificationDate *string "json:\"modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; PublicationDate *string "json:\"publication_date,omitempty\""; Severity *int32 "json:\"severity,omitempty\""; Type *string "json:\"type,omitempty\""; Version *string "json:\"version,omitempty\"" } "json:\"plugin_details,omitempty\""; ReferenceInformation *[]struct { Name *string "json:\"name,omitempty\""; Url *string "json:\"url,omitempty\""; Values *[]string "json:\"values,omitempty\"" } "json:\"reference_information,omitempty\""; RiskInformation *struct { Cvss3BaseScore *string "json:\"cvss3_base_score,omitempty\""; Cvss3TemporalScore *string "json:\"cvss3_temporal_score,omitempty\""; Cvss3TemporalVector *string "json:\"cvss3_temporal_vector,omitempty\""; Cvss3Vector *string "json:\"cvss3_vector,omitempty\""; CvssBaseScore *string "json:\"cvss_base_score,omitempty\""; CvssTemporalScore *string "json:\"cvss_temporal_score,omitempty\""; CvssTemporalVector *string "json:\"cvss_temporal_vector,omitempty\""; CvssVector *string "json:\"cvss_vector,omitempty\""; RiskFactor *string "json:\"risk_factor,omitempty\""; StigSeverity *string "json:\"stig_severity,omitempty\"" } "json:\"risk_information,omitempty\""; SeeAlso *[]string "json:\"see_also,omitempty\""; Severity *int "json:\"severity,omitempty\""; Solution *string "json:\"solution,omitempty\""; Synopsis *string "json:\"synopsis,omitempty\""; Vpr *struct { Drivers *map[string]interface {} "json:\"drivers,omitempty\""; Score *float32 "json:\"score,omitempty\""; Updated *string "json:\"updated,omitempty\"" } "json:\"vpr,omitempty\""; VulnCount *int32 "json:\"vuln_count,omitempty\""; VulnerabilityInformation *struct { AssetInventory *string "json:\"asset_inventory,omitempty\""; Cpe *string "json:\"cpe,omitempty\""; DefaultAccount *string "json:\"default_account,omitempty\""; ExploitAvailable *bool "json:\"exploit_available,omitempty\""; ExploitFrameworks *[]struct { Exploits *[]struct { Name *string "json:\"name,omitempty\""; Url *string "json:\"url,omitempty\"" } "json:\"exploits,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"exploit_frameworks,omitempty\""; ExploitabilityEase *string "json:\"exploitability_ease,omitempty\""; ExploitedByMalware *bool "json:\"exploited_by_malware,omitempty\""; ExploitedByNessus *bool "json:\"exploited_by_nessus,omitempty\""; InTheNews *bool "json:\"in_the_news,omitempty\""; Malware *string "json:\"malware,omitempty\""; PatchPublicationDate *string "json:\"patch_publication_date,omitempty\""; UnsupportedByVendor *bool "json:\"unsupported_by_vendor,omitempty\""; VulnerabilityPublicationDate *string "json:\"vulnerability_publication_date,omitempty\"" } "json:\"vulnerability_information,omitempty\"" } "json:\"info,omitempty\"" }
	if i, ok := r.(*struct {
		Info *struct {
			Count       *int32  "json:\"count,omitempty\""
			Description *string "json:\"description,omitempty\""
			Discovery   *struct {
				SeenFirst *string "json:\"seen_first,omitempty\""
				SeenLast  *string "json:\"seen_last,omitempty\""
			} "json:\"discovery,omitempty\""
			PluginDetails *struct {
				Family           *string "json:\"family,omitempty\""
				ModificationDate *string "json:\"modification_date,omitempty\""
				Name             *string "json:\"name,omitempty\""
				PublicationDate  *string "json:\"publication_date,omitempty\""
				Severity         *int32  "json:\"severity,omitempty\""
				Type             *string "json:\"type,omitempty\""
				Version          *string "json:\"version,omitempty\""
			} "json:\"plugin_details,omitempty\""
			ReferenceInformation *[]struct {
				Name   *string   "json:\"name,omitempty\""
				Url    *string   "json:\"url,omitempty\""
				Values *[]string "json:\"values,omitempty\""
			} "json:\"reference_information,omitempty\""
			RiskInformation *struct {
				Cvss3BaseScore      *string "json:\"cvss3_base_score,omitempty\""
				Cvss3TemporalScore  *string "json:\"cvss3_temporal_score,omitempty\""
				Cvss3TemporalVector *string "json:\"cvss3_temporal_vector,omitempty\""
				Cvss3Vector         *string "json:\"cvss3_vector,omitempty\""
				CvssBaseScore       *string "json:\"cvss_base_score,omitempty\""
				CvssTemporalScore   *string "json:\"cvss_temporal_score,omitempty\""
				CvssTemporalVector  *string "json:\"cvss_temporal_vector,omitempty\""
				CvssVector          *string "json:\"cvss_vector,omitempty\""
				RiskFactor          *string "json:\"risk_factor,omitempty\""
				StigSeverity        *string "json:\"stig_severity,omitempty\""
			} "json:\"risk_information,omitempty\""
			SeeAlso  *[]string "json:\"see_also,omitempty\""
			Severity *int      "json:\"severity,omitempty\""
			Solution *string   "json:\"solution,omitempty\""
			Synopsis *string   "json:\"synopsis,omitempty\""
			Vpr      *struct {
				Drivers *map[string]interface{} "json:\"drivers,omitempty\""
				Score   *float32                "json:\"score,omitempty\""
				Updated *string                 "json:\"updated,omitempty\""
			} "json:\"vpr,omitempty\""
			VulnCount                *int32 "json:\"vuln_count,omitempty\""
			VulnerabilityInformation *struct {
				AssetInventory    *string "json:\"asset_inventory,omitempty\""
				Cpe               *string "json:\"cpe,omitempty\""
				DefaultAccount    *string "json:\"default_account,omitempty\""
				ExploitAvailable  *bool   "json:\"exploit_available,omitempty\""
				ExploitFrameworks *[]struct {
					Exploits *[]struct {
						Name *string "json:\"name,omitempty\""
						Url  *string "json:\"url,omitempty\""
					} "json:\"exploits,omitempty\""
					Name *string "json:\"name,omitempty\""
				} "json:\"exploit_frameworks,omitempty\""
				ExploitabilityEase           *string "json:\"exploitability_ease,omitempty\""
				ExploitedByMalware           *bool   "json:\"exploited_by_malware,omitempty\""
				ExploitedByNessus            *bool   "json:\"exploited_by_nessus,omitempty\""
				InTheNews                    *bool   "json:\"in_the_news,omitempty\""
				Malware                      *string "json:\"malware,omitempty\""
				PatchPublicationDate         *string "json:\"patch_publication_date,omitempty\""
				UnsupportedByVendor          *bool   "json:\"unsupported_by_vendor,omitempty\""
				VulnerabilityPublicationDate *string "json:\"vulnerability_publication_date,omitempty\""
			} "json:\"vulnerability_information,omitempty\""
		} "json:\"info,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AssetVulnerabilityOutput calls the Workbenches ´s function
func (c *Workbenches) AssetVulnerabilityOutput(arg1 string, arg2 int32, params *WorkbenchesAssetVulnerabilityOutputParams, reqEditors ...RequestEditorFn) (*struct {
	Outputs *[]struct {
		PluginOutput *string "json:\"plugin_output,omitempty\""
		States       *[]struct {
			Name    *string "json:\"name,omitempty\""
			Results *[]struct {
				ApplicationProtocol *string "json:\"application_protocol,omitempty\""
				Assets              *[]struct {
					FirstSeen   *time.Time "json:\"first_seen,omitempty\""
					Fqdn        *string    "json:\"fqdn,omitempty\""
					Hostname    *string    "json:\"hostname,omitempty\""
					Id          *string    "json:\"id,omitempty\""
					Ipv4        *string    "json:\"ipv4,omitempty\""
					LastSeen    *time.Time "json:\"last_seen,omitempty\""
					NetbiosName *string    "json:\"netbios_name,omitempty\""
					Uuid        *string    "json:\"uuid,omitempty\""
				} "json:\"assets,omitempty\""
				Port              *int    "json:\"port,omitempty\""
				Severity          *int    "json:\"severity,omitempty\""
				TransportProtocol *string "json:\"transport_protocol,omitempty\""
			} "json:\"results,omitempty\""
		} "json:\"states,omitempty\""
	} "json:\"outputs,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.WorkbenchesAssetVulnerabilityOutputWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Outputs *[]struct { PluginOutput *string "json:\"plugin_output,omitempty\""; States *[]struct { Name *string "json:\"name,omitempty\""; Results *[]struct { ApplicationProtocol *string "json:\"application_protocol,omitempty\""; Assets *[]struct { FirstSeen *time.Time "json:\"first_seen,omitempty\""; Fqdn *string "json:\"fqdn,omitempty\""; Hostname *string "json:\"hostname,omitempty\""; Id *string "json:\"id,omitempty\""; Ipv4 *string "json:\"ipv4,omitempty\""; LastSeen *time.Time "json:\"last_seen,omitempty\""; NetbiosName *string "json:\"netbios_name,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"assets,omitempty\""; Port *int "json:\"port,omitempty\""; Severity *int "json:\"severity,omitempty\""; TransportProtocol *string "json:\"transport_protocol,omitempty\"" } "json:\"results,omitempty\"" } "json:\"states,omitempty\"" } "json:\"outputs,omitempty\"" }
	if i, ok := r.(*struct {
		Outputs *[]struct {
			PluginOutput *string "json:\"plugin_output,omitempty\""
			States       *[]struct {
				Name    *string "json:\"name,omitempty\""
				Results *[]struct {
					ApplicationProtocol *string "json:\"application_protocol,omitempty\""
					Assets              *[]struct {
						FirstSeen   *time.Time "json:\"first_seen,omitempty\""
						Fqdn        *string    "json:\"fqdn,omitempty\""
						Hostname    *string    "json:\"hostname,omitempty\""
						Id          *string    "json:\"id,omitempty\""
						Ipv4        *string    "json:\"ipv4,omitempty\""
						LastSeen    *time.Time "json:\"last_seen,omitempty\""
						NetbiosName *string    "json:\"netbios_name,omitempty\""
						Uuid        *string    "json:\"uuid,omitempty\""
					} "json:\"assets,omitempty\""
					Port              *int    "json:\"port,omitempty\""
					Severity          *int    "json:\"severity,omitempty\""
					TransportProtocol *string "json:\"transport_protocol,omitempty\""
				} "json:\"results,omitempty\""
			} "json:\"states,omitempty\""
		} "json:\"outputs,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AssetsActivity calls the Workbenches ´s function
func (c *Workbenches) AssetsActivity(contentType string, reqEditors ...RequestEditorFn) (*[]struct {
	Details *struct {
		AssetId                   *string                 "json:\"assetId,omitempty\""
		ContainerId               *string                 "json:\"containerId,omitempty\""
		CreatedAt                 *int                    "json:\"createdAt,omitempty\""
		DeletedAt                 *int                    "json:\"deletedAt,omitempty\""
		DeletedBy                 *string                 "json:\"deletedBy,omitempty\""
		FirstScanTime             *int                    "json:\"firstScanTime,omitempty\""
		HasAgent                  *bool                   "json:\"hasAgent,omitempty\""
		HasPluginResults          *bool                   "json:\"hasPluginResults,omitempty\""
		LastAuthenticatedScanTime *int                    "json:\"lastAuthenticatedScanTime,omitempty\""
		LastLicensedScanTime      *int                    "json:\"lastLicensedScanTime,omitempty\""
		LastLicensedScanTimeV2    *int                    "json:\"lastLicensedScanTimeV2,omitempty\""
		LastScanTime              *int                    "json:\"lastScanTime,omitempty\""
		Properties                *map[string]interface{} "json:\"properties,omitempty\""
		Sources                   *[]struct {
			FirstSeen *string "json:\"firstSeen,omitempty\""
			LastSeen  *string "json:\"lastSeen,omitempty\""
			Name      *string "json:\"name,omitempty\""
		} "json:\"sources,omitempty\""
		TerminatedAt *int    "json:\"terminatedAt,omitempty\""
		TerminatedBy *string "json:\"terminatedBy,omitempty\""
		UpdatedAt    *int    "json:\"updatedAt,omitempty\""
	} "json:\"details,omitempty\""
	ScanId     *string "json:\"scan_id,omitempty\""
	ScheduleId *string "json:\"schedule_id,omitempty\""
	Source     *string "json:\"source,omitempty\""
	Timestamp  *int    "json:\"timestamp,omitempty\""
	Type       *string "json:\"type,omitempty\""
	Updates    *[]struct {
		Method   *string "json:\"method,omitempty\""
		Property *string "json:\"property,omitempty\""
		Value    *string "json:\"value,omitempty\""
	} "json:\"updates,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.WorkbenchesAssetsActivityWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { Details *struct { AssetId *string "json:\"assetId,omitempty\""; ContainerId *string "json:\"containerId,omitempty\""; CreatedAt *int "json:\"createdAt,omitempty\""; DeletedAt *int "json:\"deletedAt,omitempty\""; DeletedBy *string "json:\"deletedBy,omitempty\""; FirstScanTime *int "json:\"firstScanTime,omitempty\""; HasAgent *bool "json:\"hasAgent,omitempty\""; HasPluginResults *bool "json:\"hasPluginResults,omitempty\""; LastAuthenticatedScanTime *int "json:\"lastAuthenticatedScanTime,omitempty\""; LastLicensedScanTime *int "json:\"lastLicensedScanTime,omitempty\""; LastLicensedScanTimeV2 *int "json:\"lastLicensedScanTimeV2,omitempty\""; LastScanTime *int "json:\"lastScanTime,omitempty\""; Properties *map[string]interface {} "json:\"properties,omitempty\""; Sources *[]struct { FirstSeen *string "json:\"firstSeen,omitempty\""; LastSeen *string "json:\"lastSeen,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"sources,omitempty\""; TerminatedAt *int "json:\"terminatedAt,omitempty\""; TerminatedBy *string "json:\"terminatedBy,omitempty\""; UpdatedAt *int "json:\"updatedAt,omitempty\"" } "json:\"details,omitempty\""; ScanId *string "json:\"scan_id,omitempty\""; ScheduleId *string "json:\"schedule_id,omitempty\""; Source *string "json:\"source,omitempty\""; Timestamp *int "json:\"timestamp,omitempty\""; Type *string "json:\"type,omitempty\""; Updates *[]struct { Method *string "json:\"method,omitempty\""; Property *string "json:\"property,omitempty\""; Value *string "json:\"value,omitempty\"" } "json:\"updates,omitempty\"" }
	if i, ok := r.(*[]struct {
		Details *struct {
			AssetId                   *string                 "json:\"assetId,omitempty\""
			ContainerId               *string                 "json:\"containerId,omitempty\""
			CreatedAt                 *int                    "json:\"createdAt,omitempty\""
			DeletedAt                 *int                    "json:\"deletedAt,omitempty\""
			DeletedBy                 *string                 "json:\"deletedBy,omitempty\""
			FirstScanTime             *int                    "json:\"firstScanTime,omitempty\""
			HasAgent                  *bool                   "json:\"hasAgent,omitempty\""
			HasPluginResults          *bool                   "json:\"hasPluginResults,omitempty\""
			LastAuthenticatedScanTime *int                    "json:\"lastAuthenticatedScanTime,omitempty\""
			LastLicensedScanTime      *int                    "json:\"lastLicensedScanTime,omitempty\""
			LastLicensedScanTimeV2    *int                    "json:\"lastLicensedScanTimeV2,omitempty\""
			LastScanTime              *int                    "json:\"lastScanTime,omitempty\""
			Properties                *map[string]interface{} "json:\"properties,omitempty\""
			Sources                   *[]struct {
				FirstSeen *string "json:\"firstSeen,omitempty\""
				LastSeen  *string "json:\"lastSeen,omitempty\""
				Name      *string "json:\"name,omitempty\""
			} "json:\"sources,omitempty\""
			TerminatedAt *int    "json:\"terminatedAt,omitempty\""
			TerminatedBy *string "json:\"terminatedBy,omitempty\""
			UpdatedAt    *int    "json:\"updatedAt,omitempty\""
		} "json:\"details,omitempty\""
		ScanId     *string "json:\"scan_id,omitempty\""
		ScheduleId *string "json:\"schedule_id,omitempty\""
		Source     *string "json:\"source,omitempty\""
		Timestamp  *int    "json:\"timestamp,omitempty\""
		Type       *string "json:\"type,omitempty\""
		Updates    *[]struct {
			Method   *string "json:\"method,omitempty\""
			Property *string "json:\"property,omitempty\""
			Value    *string "json:\"value,omitempty\""
		} "json:\"updates,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AssetsDelete calls the Workbenches ´s function
func (c *Workbenches) AssetsDelete(contentType string, reqEditors ...RequestEditorFn) (*WorkbenchesAssetsDeleteResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.WorkbenchesAssetsDeleteWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *WorkbenchesAssetsDeleteResponse
	if i, ok := r.(*WorkbenchesAssetsDeleteResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AssetsVulnerabilities calls the Workbenches ´s function
func (c *Workbenches) AssetsVulnerabilities(params *WorkbenchesAssetsVulnerabilitiesParams, reqEditors ...RequestEditorFn) (*[]struct {
	AgentName   *[]string "json:\"agent_name,omitempty\""
	Fqdn        *[]string "json:\"fqdn,omitempty\""
	Id          *string   "json:\"id,omitempty\""
	Ipv4        *[]string "json:\"ipv4,omitempty\""
	Ipv6        *[]string "json:\"ipv6,omitempty\""
	LastSeen    *string   "json:\"last_seen,omitempty\""
	NetbiosName *[]string "json:\"netbios_name,omitempty\""
	Severities  *[]struct {
		Count *int    "json:\"count,omitempty\""
		Level *int    "json:\"level,omitempty\""
		Name  *string "json:\"name,omitempty\""
	} "json:\"severities,omitempty\""
	Total *int "json:\"total,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.WorkbenchesAssetsVulnerabilitiesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { AgentName *[]string "json:\"agent_name,omitempty\""; Fqdn *[]string "json:\"fqdn,omitempty\""; Id *string "json:\"id,omitempty\""; Ipv4 *[]string "json:\"ipv4,omitempty\""; Ipv6 *[]string "json:\"ipv6,omitempty\""; LastSeen *string "json:\"last_seen,omitempty\""; NetbiosName *[]string "json:\"netbios_name,omitempty\""; Severities *[]struct { Count *int "json:\"count,omitempty\""; Level *int "json:\"level,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"severities,omitempty\""; Total *int "json:\"total,omitempty\"" }
	if i, ok := r.(*[]struct {
		AgentName   *[]string "json:\"agent_name,omitempty\""
		Fqdn        *[]string "json:\"fqdn,omitempty\""
		Id          *string   "json:\"id,omitempty\""
		Ipv4        *[]string "json:\"ipv4,omitempty\""
		Ipv6        *[]string "json:\"ipv6,omitempty\""
		LastSeen    *string   "json:\"last_seen,omitempty\""
		NetbiosName *[]string "json:\"netbios_name,omitempty\""
		Severities  *[]struct {
			Count *int    "json:\"count,omitempty\""
			Level *int    "json:\"level,omitempty\""
			Name  *string "json:\"name,omitempty\""
		} "json:\"severities,omitempty\""
		Total *int "json:\"total,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Assets calls the Workbenches ´s function
func (c *Workbenches) Assets(params *WorkbenchesAssetsParams, reqEditors ...RequestEditorFn) (*struct {
	Assets *[]struct {
		AcrDrivers *[]struct {
			DriverName  *string   "json:\"driver_name,omitempty\""
			DriverValue *[]string "json:\"driver_value,omitempty\""
		} "json:\"acr_drivers,omitempty\""
		AcrScore        *int      "json:\"acr_score,omitempty\""
		AgentName       *[]string "json:\"agent_name,omitempty\""
		AwsEc2Name      *[]string "json:\"aws_ec2_name,omitempty\""
		BigfixAssetId   *[]string "json:\"bigfix_asset_id,omitempty\""
		ExposureScore   *int      "json:\"exposure_score,omitempty\""
		Fqdn            *[]string "json:\"fqdn,omitempty\""
		HasAgent        *bool     "json:\"has_agent,omitempty\""
		Id              *string   "json:\"id,omitempty\""
		Ipv4            *[]string "json:\"ipv4,omitempty\""
		Ipv6            *[]string "json:\"ipv6,omitempty\""
		LastScanTarget  *string   "json:\"last_scan_target,omitempty\""
		LastSeen        *string   "json:\"last_seen,omitempty\""
		MacAddress      *[]string "json:\"mac_address,omitempty\""
		NetbiosName     *[]string "json:\"netbios_name,omitempty\""
		OperatingSystem *[]string "json:\"operating_system,omitempty\""
		ScanFrequency   *[]struct {
			Frequency *int  "json:\"frequency,omitempty\""
			Interval  *int  "json:\"interval,omitempty\""
			Licensed  *bool "json:\"licensed,omitempty\""
		} "json:\"scan_frequency,omitempty\""
		Sources *[]struct {
			FirstSeen *string "json:\"first_seen,omitempty\""
			LastSeen  *string "json:\"last_seen,omitempty\""
			Name      *string "json:\"name,omitempty\""
		} "json:\"sources,omitempty\""
	} "json:\"assets,omitempty\""
	Total *int "json:\"total,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.WorkbenchesAssetsWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Assets *[]struct { AcrDrivers *[]struct { DriverName *string "json:\"driver_name,omitempty\""; DriverValue *[]string "json:\"driver_value,omitempty\"" } "json:\"acr_drivers,omitempty\""; AcrScore *int "json:\"acr_score,omitempty\""; AgentName *[]string "json:\"agent_name,omitempty\""; AwsEc2Name *[]string "json:\"aws_ec2_name,omitempty\""; BigfixAssetId *[]string "json:\"bigfix_asset_id,omitempty\""; ExposureScore *int "json:\"exposure_score,omitempty\""; Fqdn *[]string "json:\"fqdn,omitempty\""; HasAgent *bool "json:\"has_agent,omitempty\""; Id *string "json:\"id,omitempty\""; Ipv4 *[]string "json:\"ipv4,omitempty\""; Ipv6 *[]string "json:\"ipv6,omitempty\""; LastScanTarget *string "json:\"last_scan_target,omitempty\""; LastSeen *string "json:\"last_seen,omitempty\""; MacAddress *[]string "json:\"mac_address,omitempty\""; NetbiosName *[]string "json:\"netbios_name,omitempty\""; OperatingSystem *[]string "json:\"operating_system,omitempty\""; ScanFrequency *[]struct { Frequency *int "json:\"frequency,omitempty\""; Interval *int "json:\"interval,omitempty\""; Licensed *bool "json:\"licensed,omitempty\"" } "json:\"scan_frequency,omitempty\""; Sources *[]struct { FirstSeen *string "json:\"first_seen,omitempty\""; LastSeen *string "json:\"last_seen,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"sources,omitempty\"" } "json:\"assets,omitempty\""; Total *int "json:\"total,omitempty\"" }
	if i, ok := r.(*struct {
		Assets *[]struct {
			AcrDrivers *[]struct {
				DriverName  *string   "json:\"driver_name,omitempty\""
				DriverValue *[]string "json:\"driver_value,omitempty\""
			} "json:\"acr_drivers,omitempty\""
			AcrScore        *int      "json:\"acr_score,omitempty\""
			AgentName       *[]string "json:\"agent_name,omitempty\""
			AwsEc2Name      *[]string "json:\"aws_ec2_name,omitempty\""
			BigfixAssetId   *[]string "json:\"bigfix_asset_id,omitempty\""
			ExposureScore   *int      "json:\"exposure_score,omitempty\""
			Fqdn            *[]string "json:\"fqdn,omitempty\""
			HasAgent        *bool     "json:\"has_agent,omitempty\""
			Id              *string   "json:\"id,omitempty\""
			Ipv4            *[]string "json:\"ipv4,omitempty\""
			Ipv6            *[]string "json:\"ipv6,omitempty\""
			LastScanTarget  *string   "json:\"last_scan_target,omitempty\""
			LastSeen        *string   "json:\"last_seen,omitempty\""
			MacAddress      *[]string "json:\"mac_address,omitempty\""
			NetbiosName     *[]string "json:\"netbios_name,omitempty\""
			OperatingSystem *[]string "json:\"operating_system,omitempty\""
			ScanFrequency   *[]struct {
				Frequency *int  "json:\"frequency,omitempty\""
				Interval  *int  "json:\"interval,omitempty\""
				Licensed  *bool "json:\"licensed,omitempty\""
			} "json:\"scan_frequency,omitempty\""
			Sources *[]struct {
				FirstSeen *string "json:\"first_seen,omitempty\""
				LastSeen  *string "json:\"last_seen,omitempty\""
				Name      *string "json:\"name,omitempty\""
			} "json:\"sources,omitempty\""
		} "json:\"assets,omitempty\""
		Total *int "json:\"total,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ExportDownload calls the Workbenches ´s function
func (c *Workbenches) ExportDownload(arg1 int32, reqEditors ...RequestEditorFn) (*WorkbenchesExportDownloadResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.WorkbenchesExportDownloadWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *WorkbenchesExportDownloadResponse
	if i, ok := r.(*WorkbenchesExportDownloadResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ExportRequest calls the Workbenches ´s function
func (c *Workbenches) ExportRequest(params *WorkbenchesExportRequestParams, reqEditors ...RequestEditorFn) (*struct {
	File *int "json:\"file,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.WorkbenchesExportRequestWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { File *int "json:\"file,omitempty\"" }
	if i, ok := r.(*struct {
		File *int "json:\"file,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ExportStatus calls the Workbenches ´s function
func (c *Workbenches) ExportStatus(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
	Progress      *string "json:\"progress,omitempty\""
	ProgressTotal *string "json:\"progress_total,omitempty\""
	Status        *string "json:\"status,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.WorkbenchesExportStatusWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Progress *string "json:\"progress,omitempty\""; ProgressTotal *string "json:\"progress_total,omitempty\""; Status *string "json:\"status,omitempty\"" }
	if i, ok := r.(*struct {
		Progress      *string "json:\"progress,omitempty\""
		ProgressTotal *string "json:\"progress_total,omitempty\""
		Status        *string "json:\"status,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Vulnerabilities calls the Workbenches ´s function
func (c *Workbenches) Vulnerabilities(params *WorkbenchesVulnerabilitiesParams, reqEditors ...RequestEditorFn) (*struct {
	TotalAssetCount         *int "json:\"total_asset_count,omitempty\""
	TotalVulnerabilityCount *int "json:\"total_vulnerability_count,omitempty\""
	Vulnerabilities         *[]struct {
		AcceptedCount    *int32 "json:\"accepted_count,omitempty\""
		Count            *int   "json:\"count,omitempty\""
		CountsBySeverity *[]struct {
			Count *int32 "json:\"count,omitempty\""
			Value *int   "json:\"value,omitempty\""
		} "json:\"counts by severity,omitempty\""
		PluginFamily       *string  "json:\"plugin_family,omitempty\""
		PluginId           *int     "json:\"plugin_id,omitempty\""
		PluginName         *string  "json:\"plugin_name,omitempty\""
		RecastedCount      *int32   "json:\"recasted_count,omitempty\""
		Severity           *int32   "json:\"severity,omitempty\""
		VprScore           *float32 "json:\"vpr_score,omitempty\""
		VulnerabilityState *string  "json:\"vulnerability_state,omitempty\""
	} "json:\"vulnerabilities,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.WorkbenchesVulnerabilitiesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { TotalAssetCount *int "json:\"total_asset_count,omitempty\""; TotalVulnerabilityCount *int "json:\"total_vulnerability_count,omitempty\""; Vulnerabilities *[]struct { AcceptedCount *int32 "json:\"accepted_count,omitempty\""; Count *int "json:\"count,omitempty\""; CountsBySeverity *[]struct { Count *int32 "json:\"count,omitempty\""; Value *int "json:\"value,omitempty\"" } "json:\"counts by severity,omitempty\""; PluginFamily *string "json:\"plugin_family,omitempty\""; PluginId *int "json:\"plugin_id,omitempty\""; PluginName *string "json:\"plugin_name,omitempty\""; RecastedCount *int32 "json:\"recasted_count,omitempty\""; Severity *int32 "json:\"severity,omitempty\""; VprScore *float32 "json:\"vpr_score,omitempty\""; VulnerabilityState *string "json:\"vulnerability_state,omitempty\"" } "json:\"vulnerabilities,omitempty\"" }
	if i, ok := r.(*struct {
		TotalAssetCount         *int "json:\"total_asset_count,omitempty\""
		TotalVulnerabilityCount *int "json:\"total_vulnerability_count,omitempty\""
		Vulnerabilities         *[]struct {
			AcceptedCount    *int32 "json:\"accepted_count,omitempty\""
			Count            *int   "json:\"count,omitempty\""
			CountsBySeverity *[]struct {
				Count *int32 "json:\"count,omitempty\""
				Value *int   "json:\"value,omitempty\""
			} "json:\"counts by severity,omitempty\""
			PluginFamily       *string  "json:\"plugin_family,omitempty\""
			PluginId           *int     "json:\"plugin_id,omitempty\""
			PluginName         *string  "json:\"plugin_name,omitempty\""
			RecastedCount      *int32   "json:\"recasted_count,omitempty\""
			Severity           *int32   "json:\"severity,omitempty\""
			VprScore           *float32 "json:\"vpr_score,omitempty\""
			VulnerabilityState *string  "json:\"vulnerability_state,omitempty\""
		} "json:\"vulnerabilities,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// VulnerabilityInfo calls the Workbenches ´s function
func (c *Workbenches) VulnerabilityInfo(arg1 int32, params *WorkbenchesVulnerabilityInfoParams, reqEditors ...RequestEditorFn) (*struct {
	Count       *int32  "json:\"count,omitempty\""
	Description *string "json:\"description,omitempty\""
	Discovery   *struct {
		SeenFirst *string "json:\"seen_first,omitempty\""
		SeenLast  *string "json:\"seen_last,omitempty\""
	} "json:\"discovery,omitempty\""
	PluginDetails *struct {
		Family           *string "json:\"family,omitempty\""
		ModificationDate *string "json:\"modification_date,omitempty\""
		Name             *string "json:\"name,omitempty\""
		PublicationDate  *string "json:\"publication_date,omitempty\""
		Severity         *int32  "json:\"severity,omitempty\""
		Type             *string "json:\"type,omitempty\""
		Version          *string "json:\"version,omitempty\""
	} "json:\"plugin_details,omitempty\""
	ReferenceInformation *[]struct {
		Name   *string   "json:\"name,omitempty\""
		Url    *string   "json:\"url,omitempty\""
		Values *[]string "json:\"values,omitempty\""
	} "json:\"reference_information,omitempty\""
	RiskInformation *struct {
		Cvss3BaseScore      *string "json:\"cvss3_base_score,omitempty\""
		Cvss3TemporalScore  *string "json:\"cvss3_temporal_score,omitempty\""
		Cvss3TemporalVector *string "json:\"cvss3_temporal_vector,omitempty\""
		Cvss3Vector         *string "json:\"cvss3_vector,omitempty\""
		CvssBaseScore       *string "json:\"cvss_base_score,omitempty\""
		CvssTemporalScore   *string "json:\"cvss_temporal_score,omitempty\""
		CvssTemporalVector  *string "json:\"cvss_temporal_vector,omitempty\""
		CvssVector          *string "json:\"cvss_vector,omitempty\""
		RiskFactor          *string "json:\"risk_factor,omitempty\""
		StigSeverity        *string "json:\"stig_severity,omitempty\""
	} "json:\"risk_information,omitempty\""
	SeeAlso  *[]string "json:\"see_also,omitempty\""
	Severity *int      "json:\"severity,omitempty\""
	Solution *string   "json:\"solution,omitempty\""
	Synopsis *string   "json:\"synopsis,omitempty\""
	Vpr      *struct {
		Drivers *map[string]interface{} "json:\"drivers,omitempty\""
		Score   *float32                "json:\"score,omitempty\""
		Updated *string                 "json:\"updated,omitempty\""
	} "json:\"vpr,omitempty\""
	VulnCount                *int32 "json:\"vuln_count,omitempty\""
	VulnerabilityInformation *struct {
		AssetInventory    *string "json:\"asset_inventory,omitempty\""
		Cpe               *string "json:\"cpe,omitempty\""
		DefaultAccount    *string "json:\"default_account,omitempty\""
		ExploitAvailable  *bool   "json:\"exploit_available,omitempty\""
		ExploitFrameworks *[]struct {
			Exploits *[]struct {
				Name *string "json:\"name,omitempty\""
				Url  *string "json:\"url,omitempty\""
			} "json:\"exploits,omitempty\""
			Name *string "json:\"name,omitempty\""
		} "json:\"exploit_frameworks,omitempty\""
		ExploitabilityEase           *string "json:\"exploitability_ease,omitempty\""
		ExploitedByMalware           *bool   "json:\"exploited_by_malware,omitempty\""
		ExploitedByNessus            *bool   "json:\"exploited_by_nessus,omitempty\""
		InTheNews                    *bool   "json:\"in_the_news,omitempty\""
		Malware                      *string "json:\"malware,omitempty\""
		PatchPublicationDate         *string "json:\"patch_publication_date,omitempty\""
		UnsupportedByVendor          *bool   "json:\"unsupported_by_vendor,omitempty\""
		VulnerabilityPublicationDate *string "json:\"vulnerability_publication_date,omitempty\""
	} "json:\"vulnerability_information,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.WorkbenchesVulnerabilityInfoWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Count *int32 "json:\"count,omitempty\""; Description *string "json:\"description,omitempty\""; Discovery *struct { SeenFirst *string "json:\"seen_first,omitempty\""; SeenLast *string "json:\"seen_last,omitempty\"" } "json:\"discovery,omitempty\""; PluginDetails *struct { Family *string "json:\"family,omitempty\""; ModificationDate *string "json:\"modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; PublicationDate *string "json:\"publication_date,omitempty\""; Severity *int32 "json:\"severity,omitempty\""; Type *string "json:\"type,omitempty\""; Version *string "json:\"version,omitempty\"" } "json:\"plugin_details,omitempty\""; ReferenceInformation *[]struct { Name *string "json:\"name,omitempty\""; Url *string "json:\"url,omitempty\""; Values *[]string "json:\"values,omitempty\"" } "json:\"reference_information,omitempty\""; RiskInformation *struct { Cvss3BaseScore *string "json:\"cvss3_base_score,omitempty\""; Cvss3TemporalScore *string "json:\"cvss3_temporal_score,omitempty\""; Cvss3TemporalVector *string "json:\"cvss3_temporal_vector,omitempty\""; Cvss3Vector *string "json:\"cvss3_vector,omitempty\""; CvssBaseScore *string "json:\"cvss_base_score,omitempty\""; CvssTemporalScore *string "json:\"cvss_temporal_score,omitempty\""; CvssTemporalVector *string "json:\"cvss_temporal_vector,omitempty\""; CvssVector *string "json:\"cvss_vector,omitempty\""; RiskFactor *string "json:\"risk_factor,omitempty\""; StigSeverity *string "json:\"stig_severity,omitempty\"" } "json:\"risk_information,omitempty\""; SeeAlso *[]string "json:\"see_also,omitempty\""; Severity *int "json:\"severity,omitempty\""; Solution *string "json:\"solution,omitempty\""; Synopsis *string "json:\"synopsis,omitempty\""; Vpr *struct { Drivers *map[string]interface {} "json:\"drivers,omitempty\""; Score *float32 "json:\"score,omitempty\""; Updated *string "json:\"updated,omitempty\"" } "json:\"vpr,omitempty\""; VulnCount *int32 "json:\"vuln_count,omitempty\""; VulnerabilityInformation *struct { AssetInventory *string "json:\"asset_inventory,omitempty\""; Cpe *string "json:\"cpe,omitempty\""; DefaultAccount *string "json:\"default_account,omitempty\""; ExploitAvailable *bool "json:\"exploit_available,omitempty\""; ExploitFrameworks *[]struct { Exploits *[]struct { Name *string "json:\"name,omitempty\""; Url *string "json:\"url,omitempty\"" } "json:\"exploits,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"exploit_frameworks,omitempty\""; ExploitabilityEase *string "json:\"exploitability_ease,omitempty\""; ExploitedByMalware *bool "json:\"exploited_by_malware,omitempty\""; ExploitedByNessus *bool "json:\"exploited_by_nessus,omitempty\""; InTheNews *bool "json:\"in_the_news,omitempty\""; Malware *string "json:\"malware,omitempty\""; PatchPublicationDate *string "json:\"patch_publication_date,omitempty\""; UnsupportedByVendor *bool "json:\"unsupported_by_vendor,omitempty\""; VulnerabilityPublicationDate *string "json:\"vulnerability_publication_date,omitempty\"" } "json:\"vulnerability_information,omitempty\"" }
	if i, ok := r.(*struct {
		Count       *int32  "json:\"count,omitempty\""
		Description *string "json:\"description,omitempty\""
		Discovery   *struct {
			SeenFirst *string "json:\"seen_first,omitempty\""
			SeenLast  *string "json:\"seen_last,omitempty\""
		} "json:\"discovery,omitempty\""
		PluginDetails *struct {
			Family           *string "json:\"family,omitempty\""
			ModificationDate *string "json:\"modification_date,omitempty\""
			Name             *string "json:\"name,omitempty\""
			PublicationDate  *string "json:\"publication_date,omitempty\""
			Severity         *int32  "json:\"severity,omitempty\""
			Type             *string "json:\"type,omitempty\""
			Version          *string "json:\"version,omitempty\""
		} "json:\"plugin_details,omitempty\""
		ReferenceInformation *[]struct {
			Name   *string   "json:\"name,omitempty\""
			Url    *string   "json:\"url,omitempty\""
			Values *[]string "json:\"values,omitempty\""
		} "json:\"reference_information,omitempty\""
		RiskInformation *struct {
			Cvss3BaseScore      *string "json:\"cvss3_base_score,omitempty\""
			Cvss3TemporalScore  *string "json:\"cvss3_temporal_score,omitempty\""
			Cvss3TemporalVector *string "json:\"cvss3_temporal_vector,omitempty\""
			Cvss3Vector         *string "json:\"cvss3_vector,omitempty\""
			CvssBaseScore       *string "json:\"cvss_base_score,omitempty\""
			CvssTemporalScore   *string "json:\"cvss_temporal_score,omitempty\""
			CvssTemporalVector  *string "json:\"cvss_temporal_vector,omitempty\""
			CvssVector          *string "json:\"cvss_vector,omitempty\""
			RiskFactor          *string "json:\"risk_factor,omitempty\""
			StigSeverity        *string "json:\"stig_severity,omitempty\""
		} "json:\"risk_information,omitempty\""
		SeeAlso  *[]string "json:\"see_also,omitempty\""
		Severity *int      "json:\"severity,omitempty\""
		Solution *string   "json:\"solution,omitempty\""
		Synopsis *string   "json:\"synopsis,omitempty\""
		Vpr      *struct {
			Drivers *map[string]interface{} "json:\"drivers,omitempty\""
			Score   *float32                "json:\"score,omitempty\""
			Updated *string                 "json:\"updated,omitempty\""
		} "json:\"vpr,omitempty\""
		VulnCount                *int32 "json:\"vuln_count,omitempty\""
		VulnerabilityInformation *struct {
			AssetInventory    *string "json:\"asset_inventory,omitempty\""
			Cpe               *string "json:\"cpe,omitempty\""
			DefaultAccount    *string "json:\"default_account,omitempty\""
			ExploitAvailable  *bool   "json:\"exploit_available,omitempty\""
			ExploitFrameworks *[]struct {
				Exploits *[]struct {
					Name *string "json:\"name,omitempty\""
					Url  *string "json:\"url,omitempty\""
				} "json:\"exploits,omitempty\""
				Name *string "json:\"name,omitempty\""
			} "json:\"exploit_frameworks,omitempty\""
			ExploitabilityEase           *string "json:\"exploitability_ease,omitempty\""
			ExploitedByMalware           *bool   "json:\"exploited_by_malware,omitempty\""
			ExploitedByNessus            *bool   "json:\"exploited_by_nessus,omitempty\""
			InTheNews                    *bool   "json:\"in_the_news,omitempty\""
			Malware                      *string "json:\"malware,omitempty\""
			PatchPublicationDate         *string "json:\"patch_publication_date,omitempty\""
			UnsupportedByVendor          *bool   "json:\"unsupported_by_vendor,omitempty\""
			VulnerabilityPublicationDate *string "json:\"vulnerability_publication_date,omitempty\""
		} "json:\"vulnerability_information,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// VulnerabilityOutput calls the Workbenches ´s function
func (c *Workbenches) VulnerabilityOutput(arg1 int32, params *WorkbenchesVulnerabilityOutputParams, reqEditors ...RequestEditorFn) (*[]struct {
	PluginOutput *string "json:\"plugin_output,omitempty\""
	States       *[]struct {
		Name    *string "json:\"name,omitempty\""
		Results *[]struct {
			ApplicationProtocol *string "json:\"application_protocol,omitempty\""
			Assets              *[]struct {
				FirstSeen   *time.Time "json:\"first_seen,omitempty\""
				Fqdn        *string    "json:\"fqdn,omitempty\""
				Hostname    *string    "json:\"hostname,omitempty\""
				Id          *string    "json:\"id,omitempty\""
				Ipv4        *string    "json:\"ipv4,omitempty\""
				LastSeen    *time.Time "json:\"last_seen,omitempty\""
				NetbiosName *string    "json:\"netbios_name,omitempty\""
				Uuid        *string    "json:\"uuid,omitempty\""
			} "json:\"assets,omitempty\""
			Port              *int    "json:\"port,omitempty\""
			Severity          *int    "json:\"severity,omitempty\""
			TransportProtocol *string "json:\"transport_protocol,omitempty\""
		} "json:\"results,omitempty\""
	} "json:\"states,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.WorkbenchesVulnerabilityOutputWithResponse(c.ClientInterface.(*Client).ctx, arg1, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { PluginOutput *string "json:\"plugin_output,omitempty\""; States *[]struct { Name *string "json:\"name,omitempty\""; Results *[]struct { ApplicationProtocol *string "json:\"application_protocol,omitempty\""; Assets *[]struct { FirstSeen *time.Time "json:\"first_seen,omitempty\""; Fqdn *string "json:\"fqdn,omitempty\""; Hostname *string "json:\"hostname,omitempty\""; Id *string "json:\"id,omitempty\""; Ipv4 *string "json:\"ipv4,omitempty\""; LastSeen *time.Time "json:\"last_seen,omitempty\""; NetbiosName *string "json:\"netbios_name,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"assets,omitempty\""; Port *int "json:\"port,omitempty\""; Severity *int "json:\"severity,omitempty\""; TransportProtocol *string "json:\"transport_protocol,omitempty\"" } "json:\"results,omitempty\"" } "json:\"states,omitempty\"" }
	if i, ok := r.(*[]struct {
		PluginOutput *string "json:\"plugin_output,omitempty\""
		States       *[]struct {
			Name    *string "json:\"name,omitempty\""
			Results *[]struct {
				ApplicationProtocol *string "json:\"application_protocol,omitempty\""
				Assets              *[]struct {
					FirstSeen   *time.Time "json:\"first_seen,omitempty\""
					Fqdn        *string    "json:\"fqdn,omitempty\""
					Hostname    *string    "json:\"hostname,omitempty\""
					Id          *string    "json:\"id,omitempty\""
					Ipv4        *string    "json:\"ipv4,omitempty\""
					LastSeen    *time.Time "json:\"last_seen,omitempty\""
					NetbiosName *string    "json:\"netbios_name,omitempty\""
					Uuid        *string    "json:\"uuid,omitempty\""
				} "json:\"assets,omitempty\""
				Port              *int    "json:\"port,omitempty\""
				Severity          *int    "json:\"severity,omitempty\""
				TransportProtocol *string "json:\"transport_protocol,omitempty\""
			} "json:\"results,omitempty\""
		} "json:\"states,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Permissions implementation of the Permissions interface
type Permissions struct {
	*ClientWithResponses
}

// ChangeWithBody calls the Permissions ´s function
func (c *Permissions) ChangeWithBody(arg1 PermissionsChangeParamsObjectType, arg2 int32, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.PermissionsChangeWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Change calls the Permissions ´s function
func (c *Permissions) Change(arg1 PermissionsChangeParamsObjectType, arg2 int32, arg3 PermissionsChangeJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.PermissionsChangeWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// List calls the Permissions ´s function
func (c *Permissions) List(arg1 PermissionsListParamsObjectType, arg2 int32, reqEditors ...RequestEditorFn) (*[]struct {
	DisplayName *string   "json:\"display_name,omitempty\""
	Id          *int      "json:\"id,omitempty\""
	Name        *string   "json:\"name,omitempty\""
	Owner       *int      "json:\"owner,omitempty\""
	Permissions *int32    "json:\"permissions,omitempty\""
	Type        *N200Type "json:\"type,omitempty\""
	Uuid        *string   "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.PermissionsListWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { DisplayName *string "json:\"display_name,omitempty\""; Id *int "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Owner *int "json:\"owner,omitempty\""; Permissions *int32 "json:\"permissions,omitempty\""; Type *N200Type "json:\"type,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*[]struct {
		DisplayName *string   "json:\"display_name,omitempty\""
		Id          *int      "json:\"id,omitempty\""
		Name        *string   "json:\"name,omitempty\""
		Owner       *int      "json:\"owner,omitempty\""
		Permissions *int32    "json:\"permissions,omitempty\""
		Type        *N200Type "json:\"type,omitempty\""
		Uuid        *string   "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Scanner implementation of the Scanner interface
type Scanner struct {
	*ClientWithResponses
}

// GroupsAddScanner calls the Scanner ´s function
func (c *Scanner) GroupsAddScanner(arg1 int32, arg2 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannerGroupsAddScannerWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GroupsCreateWithBody calls the Scanner ´s function
func (c *Scanner) GroupsCreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	DefaultPermissions   *int    "json:\"default_permissions,omitempty\""
	Flag                 *string "json:\"flag,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NetworkName          *string "json:\"network_name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	OwnerId              *int    "json:\"owner_id,omitempty\""
	ScanCount            *int    "json:\"scan_count,omitempty\""
	ScannerCount         *string "json:\"scanner_count,omitempty\""
	ScannerId            *int    "json:\"scanner_id,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Token                *string "json:\"token,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
	Uuid                 *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannerGroupsCreateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreationDate *int "json:\"creation_date,omitempty\""; DefaultPermissions *int "json:\"default_permissions,omitempty\""; Flag *string "json:\"flag,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkName *string "json:\"network_name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; ScanCount *int "json:\"scan_count,omitempty\""; ScannerCount *string "json:\"scanner_count,omitempty\""; ScannerId *int "json:\"scanner_id,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Token *string "json:\"token,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		DefaultPermissions   *int    "json:\"default_permissions,omitempty\""
		Flag                 *string "json:\"flag,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NetworkName          *string "json:\"network_name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		ScanCount            *int    "json:\"scan_count,omitempty\""
		ScannerCount         *string "json:\"scanner_count,omitempty\""
		ScannerId            *int    "json:\"scanner_id,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Token                *string "json:\"token,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GroupsCreate calls the Scanner ´s function
func (c *Scanner) GroupsCreate(arg1 ScannerGroupsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	DefaultPermissions   *int    "json:\"default_permissions,omitempty\""
	Flag                 *string "json:\"flag,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NetworkName          *string "json:\"network_name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	OwnerId              *int    "json:\"owner_id,omitempty\""
	ScanCount            *int    "json:\"scan_count,omitempty\""
	ScannerCount         *string "json:\"scanner_count,omitempty\""
	ScannerId            *int    "json:\"scanner_id,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Token                *string "json:\"token,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
	Uuid                 *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannerGroupsCreateWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreationDate *int "json:\"creation_date,omitempty\""; DefaultPermissions *int "json:\"default_permissions,omitempty\""; Flag *string "json:\"flag,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkName *string "json:\"network_name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; ScanCount *int "json:\"scan_count,omitempty\""; ScannerCount *string "json:\"scanner_count,omitempty\""; ScannerId *int "json:\"scanner_id,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Token *string "json:\"token,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		DefaultPermissions   *int    "json:\"default_permissions,omitempty\""
		Flag                 *string "json:\"flag,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NetworkName          *string "json:\"network_name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		ScanCount            *int    "json:\"scan_count,omitempty\""
		ScannerCount         *string "json:\"scanner_count,omitempty\""
		ScannerId            *int    "json:\"scanner_id,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Token                *string "json:\"token,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GroupsDeleteScanner calls the Scanner ´s function
func (c *Scanner) GroupsDeleteScanner(arg1 int32, arg2 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannerGroupsDeleteScannerWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GroupsDelete calls the Scanner ´s function
func (c *Scanner) GroupsDelete(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannerGroupsDeleteWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GroupsDetails calls the Scanner ´s function
func (c *Scanner) GroupsDetails(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	DefaultPermissions   *int    "json:\"default_permissions,omitempty\""
	Flag                 *string "json:\"flag,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NetworkName          *string "json:\"network_name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	OwnerId              *int    "json:\"owner_id,omitempty\""
	ScanCount            *int    "json:\"scan_count,omitempty\""
	ScannerCount         *string "json:\"scanner_count,omitempty\""
	ScannerId            *int    "json:\"scanner_id,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Token                *string "json:\"token,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
	Uuid                 *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannerGroupsDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreationDate *int "json:\"creation_date,omitempty\""; DefaultPermissions *int "json:\"default_permissions,omitempty\""; Flag *string "json:\"flag,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkName *string "json:\"network_name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; ScanCount *int "json:\"scan_count,omitempty\""; ScannerCount *string "json:\"scanner_count,omitempty\""; ScannerId *int "json:\"scanner_id,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Token *string "json:\"token,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		DefaultPermissions   *int    "json:\"default_permissions,omitempty\""
		Flag                 *string "json:\"flag,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NetworkName          *string "json:\"network_name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		ScanCount            *int    "json:\"scan_count,omitempty\""
		ScannerCount         *string "json:\"scanner_count,omitempty\""
		ScannerId            *int    "json:\"scanner_id,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Token                *string "json:\"token,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GroupsEditWithBody calls the Scanner ´s function
func (c *Scanner) GroupsEditWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	DefaultPermissions   *int    "json:\"default_permissions,omitempty\""
	Flag                 *string "json:\"flag,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NetworkName          *string "json:\"network_name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	OwnerId              *int    "json:\"owner_id,omitempty\""
	ScanCount            *int    "json:\"scan_count,omitempty\""
	ScannerCount         *string "json:\"scanner_count,omitempty\""
	ScannerId            *int    "json:\"scanner_id,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Token                *string "json:\"token,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
	Uuid                 *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannerGroupsEditWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreationDate *int "json:\"creation_date,omitempty\""; DefaultPermissions *int "json:\"default_permissions,omitempty\""; Flag *string "json:\"flag,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkName *string "json:\"network_name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; ScanCount *int "json:\"scan_count,omitempty\""; ScannerCount *string "json:\"scanner_count,omitempty\""; ScannerId *int "json:\"scanner_id,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Token *string "json:\"token,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		DefaultPermissions   *int    "json:\"default_permissions,omitempty\""
		Flag                 *string "json:\"flag,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NetworkName          *string "json:\"network_name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		ScanCount            *int    "json:\"scan_count,omitempty\""
		ScannerCount         *string "json:\"scanner_count,omitempty\""
		ScannerId            *int    "json:\"scanner_id,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Token                *string "json:\"token,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GroupsEdit calls the Scanner ´s function
func (c *Scanner) GroupsEdit(arg1 int32, arg2 ScannerGroupsEditJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	DefaultPermissions   *int    "json:\"default_permissions,omitempty\""
	Flag                 *string "json:\"flag,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NetworkName          *string "json:\"network_name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	OwnerId              *int    "json:\"owner_id,omitempty\""
	ScanCount            *int    "json:\"scan_count,omitempty\""
	ScannerCount         *string "json:\"scanner_count,omitempty\""
	ScannerId            *int    "json:\"scanner_id,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Token                *string "json:\"token,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
	Uuid                 *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannerGroupsEditWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreationDate *int "json:\"creation_date,omitempty\""; DefaultPermissions *int "json:\"default_permissions,omitempty\""; Flag *string "json:\"flag,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkName *string "json:\"network_name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; ScanCount *int "json:\"scan_count,omitempty\""; ScannerCount *string "json:\"scanner_count,omitempty\""; ScannerId *int "json:\"scanner_id,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Token *string "json:\"token,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		DefaultPermissions   *int    "json:\"default_permissions,omitempty\""
		Flag                 *string "json:\"flag,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NetworkName          *string "json:\"network_name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		ScanCount            *int    "json:\"scan_count,omitempty\""
		ScannerCount         *string "json:\"scanner_count,omitempty\""
		ScannerId            *int    "json:\"scanner_id,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Token                *string "json:\"token,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GroupsListScanners calls the Scanner ´s function
func (c *Scanner) GroupsListScanners(arg1 int32, reqEditors ...RequestEditorFn) (*[]struct {
	AwsUpdateInterval    *int      "json:\"aws_update_interval,omitempty\""
	CreationDate         *int      "json:\"creation_date,omitempty\""
	Distro               *string   "json:\"distro,omitempty\""
	EngineVersion        *string   "json:\"engine_version,omitempty\""
	Group                *bool     "json:\"group,omitempty\""
	Hostname             *string   "json:\"hostname,omitempty\""
	Id                   *int      "json:\"id,omitempty\""
	IpAddresses          *[]string "json:\"ip_addresses,omitempty\""
	Key                  *string   "json:\"key,omitempty\""
	LastConnect          *string   "json:\"last_connect,omitempty\""
	LastModificationDate *int      "json:\"last_modification_date,omitempty\""
	License              *struct {
		Agents *int "json:\"agents,omitempty\""
		Apps   *struct {
			Consec *struct {
				ActivationCode *int    "json:\"activation_code,omitempty\""
				ExpirationDate *int    "json:\"expiration_date,omitempty\""
				MaxGb          *int    "json:\"max_gb,omitempty\""
				Mode           *string "json:\"mode,omitempty\""
				Type           *string "json:\"type,omitempty\""
			} "json:\"consec,omitempty\""
			Type *string "json:\"type,omitempty\""
			Was  *struct {
				ActivationCode *int    "json:\"activation_code,omitempty\""
				ExpirationDate *int    "json:\"expiration_date,omitempty\""
				Mode           *string "json:\"mode,omitempty\""
				Type           *string "json:\"type,omitempty\""
				WebAssets      *int    "json:\"web_assets,omitempty\""
			} "json:\"was,omitempty\""
		} "json:\"apps,omitempty\""
		Ips      *int    "json:\"ips,omitempty\""
		Scanners *int    "json:\"scanners,omitempty\""
		Type     *string "json:\"type,omitempty\""
	} "json:\"license,omitempty\""
	Linked             *int    "json:\"linked,omitempty\""
	LoadedPluginSet    *string "json:\"loaded_plugin_set,omitempty\""
	Name               *string "json:\"name,omitempty\""
	NetworkName        *string "json:\"network_name,omitempty\""
	NumHosts           *int    "json:\"num_hosts,omitempty\""
	NumScans           *int    "json:\"num_scans,omitempty\""
	NumSessions        *int    "json:\"num_sessions,omitempty\""
	NumTcpSessions     *int    "json:\"num_tcp_sessions,omitempty\""
	Owner              *string "json:\"owner,omitempty\""
	OwnerId            *int    "json:\"owner_id,omitempty\""
	OwnerName          *string "json:\"owner_name,omitempty\""
	OwnerUuid          *string "json:\"owner_uuid,omitempty\""
	Platform           *string "json:\"platform,omitempty\""
	Pool               *bool   "json:\"pool,omitempty\""
	RegistrationCode   *string "json:\"registration_code,omitempty\""
	ScanCount          *int    "json:\"scan_count,omitempty\""
	Shared             *bool   "json:\"shared,omitempty\""
	Source             *string "json:\"source,omitempty\""
	Status             *string "json:\"status,omitempty\""
	SupportsRemoteLogs *bool   "json:\"supports_remote_logs,omitempty\""
	SupportsWebapp     *bool   "json:\"supports_webapp,omitempty\""
	Timestamp          *int    "json:\"timestamp,omitempty\""
	Type               *string "json:\"type,omitempty\""
	UiBuild            *string "json:\"ui_build,omitempty\""
	UiVersion          *string "json:\"ui_version,omitempty\""
	UserPermissions    *int    "json:\"user_permissions,omitempty\""
	Uuid               *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannerGroupsListScannersWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { AwsUpdateInterval *int "json:\"aws_update_interval,omitempty\""; CreationDate *int "json:\"creation_date,omitempty\""; Distro *string "json:\"distro,omitempty\""; EngineVersion *string "json:\"engine_version,omitempty\""; Group *bool "json:\"group,omitempty\""; Hostname *string "json:\"hostname,omitempty\""; Id *int "json:\"id,omitempty\""; IpAddresses *[]string "json:\"ip_addresses,omitempty\""; Key *string "json:\"key,omitempty\""; LastConnect *string "json:\"last_connect,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; License *struct { Agents *int "json:\"agents,omitempty\""; Apps *struct { Consec *struct { ActivationCode *int "json:\"activation_code,omitempty\""; ExpirationDate *int "json:\"expiration_date,omitempty\""; MaxGb *int "json:\"max_gb,omitempty\""; Mode *string "json:\"mode,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"consec,omitempty\""; Type *string "json:\"type,omitempty\""; Was *struct { ActivationCode *int "json:\"activation_code,omitempty\""; ExpirationDate *int "json:\"expiration_date,omitempty\""; Mode *string "json:\"mode,omitempty\""; Type *string "json:\"type,omitempty\""; WebAssets *int "json:\"web_assets,omitempty\"" } "json:\"was,omitempty\"" } "json:\"apps,omitempty\""; Ips *int "json:\"ips,omitempty\""; Scanners *int "json:\"scanners,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"license,omitempty\""; Linked *int "json:\"linked,omitempty\""; LoadedPluginSet *string "json:\"loaded_plugin_set,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkName *string "json:\"network_name,omitempty\""; NumHosts *int "json:\"num_hosts,omitempty\""; NumScans *int "json:\"num_scans,omitempty\""; NumSessions *int "json:\"num_sessions,omitempty\""; NumTcpSessions *int "json:\"num_tcp_sessions,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; OwnerName *string "json:\"owner_name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Platform *string "json:\"platform,omitempty\""; Pool *bool "json:\"pool,omitempty\""; RegistrationCode *string "json:\"registration_code,omitempty\""; ScanCount *int "json:\"scan_count,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Source *string "json:\"source,omitempty\""; Status *string "json:\"status,omitempty\""; SupportsRemoteLogs *bool "json:\"supports_remote_logs,omitempty\""; SupportsWebapp *bool "json:\"supports_webapp,omitempty\""; Timestamp *int "json:\"timestamp,omitempty\""; Type *string "json:\"type,omitempty\""; UiBuild *int "json:\"ui_build,omitempty\""; UiVersion *string "json:\"ui_version,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*[]struct {
		AwsUpdateInterval    *int      "json:\"aws_update_interval,omitempty\""
		CreationDate         *int      "json:\"creation_date,omitempty\""
		Distro               *string   "json:\"distro,omitempty\""
		EngineVersion        *string   "json:\"engine_version,omitempty\""
		Group                *bool     "json:\"group,omitempty\""
		Hostname             *string   "json:\"hostname,omitempty\""
		Id                   *int      "json:\"id,omitempty\""
		IpAddresses          *[]string "json:\"ip_addresses,omitempty\""
		Key                  *string   "json:\"key,omitempty\""
		LastConnect          *string   "json:\"last_connect,omitempty\""
		LastModificationDate *int      "json:\"last_modification_date,omitempty\""
		License              *struct {
			Agents *int "json:\"agents,omitempty\""
			Apps   *struct {
				Consec *struct {
					ActivationCode *int    "json:\"activation_code,omitempty\""
					ExpirationDate *int    "json:\"expiration_date,omitempty\""
					MaxGb          *int    "json:\"max_gb,omitempty\""
					Mode           *string "json:\"mode,omitempty\""
					Type           *string "json:\"type,omitempty\""
				} "json:\"consec,omitempty\""
				Type *string "json:\"type,omitempty\""
				Was  *struct {
					ActivationCode *int    "json:\"activation_code,omitempty\""
					ExpirationDate *int    "json:\"expiration_date,omitempty\""
					Mode           *string "json:\"mode,omitempty\""
					Type           *string "json:\"type,omitempty\""
					WebAssets      *int    "json:\"web_assets,omitempty\""
				} "json:\"was,omitempty\""
			} "json:\"apps,omitempty\""
			Ips      *int    "json:\"ips,omitempty\""
			Scanners *int    "json:\"scanners,omitempty\""
			Type     *string "json:\"type,omitempty\""
		} "json:\"license,omitempty\""
		Linked             *int    "json:\"linked,omitempty\""
		LoadedPluginSet    *string "json:\"loaded_plugin_set,omitempty\""
		Name               *string "json:\"name,omitempty\""
		NetworkName        *string "json:\"network_name,omitempty\""
		NumHosts           *int    "json:\"num_hosts,omitempty\""
		NumScans           *int    "json:\"num_scans,omitempty\""
		NumSessions        *int    "json:\"num_sessions,omitempty\""
		NumTcpSessions     *int    "json:\"num_tcp_sessions,omitempty\""
		Owner              *string "json:\"owner,omitempty\""
		OwnerId            *int    "json:\"owner_id,omitempty\""
		OwnerName          *string "json:\"owner_name,omitempty\""
		OwnerUuid          *string "json:\"owner_uuid,omitempty\""
		Platform           *string "json:\"platform,omitempty\""
		Pool               *bool   "json:\"pool,omitempty\""
		RegistrationCode   *string "json:\"registration_code,omitempty\""
		ScanCount          *int    "json:\"scan_count,omitempty\""
		Shared             *bool   "json:\"shared,omitempty\""
		Source             *string "json:\"source,omitempty\""
		Status             *string "json:\"status,omitempty\""
		SupportsRemoteLogs *bool   "json:\"supports_remote_logs,omitempty\""
		SupportsWebapp     *bool   "json:\"supports_webapp,omitempty\""
		Timestamp          *int    "json:\"timestamp,omitempty\""
		Type               *string "json:\"type,omitempty\""
		UiBuild            *string "json:\"ui_build,omitempty\""
		UiVersion          *string "json:\"ui_version,omitempty\""
		UserPermissions    *int    "json:\"user_permissions,omitempty\""
		Uuid               *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// GroupsList calls the Scanner ´s function
func (c *Scanner) GroupsList(reqEditors ...RequestEditorFn) (*struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	DefaultPermissions   *int    "json:\"default_permissions,omitempty\""
	Flag                 *string "json:\"flag,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NetworkName          *string "json:\"network_name,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	OwnerId              *int    "json:\"owner_id,omitempty\""
	ScanCount            *int    "json:\"scan_count,omitempty\""
	ScannerCount         *string "json:\"scanner_count,omitempty\""
	ScannerId            *int    "json:\"scanner_id,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	Token                *string "json:\"token,omitempty\""
	Type                 *string "json:\"type,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
	Uuid                 *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.ScannerGroupsListWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreationDate *int "json:\"creation_date,omitempty\""; DefaultPermissions *int "json:\"default_permissions,omitempty\""; Flag *string "json:\"flag,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NetworkName *string "json:\"network_name,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; ScanCount *int "json:\"scan_count,omitempty\""; ScannerCount *string "json:\"scanner_count,omitempty\""; ScannerId *int "json:\"scanner_id,omitempty\""; Shared *bool "json:\"shared,omitempty\""; Token *string "json:\"token,omitempty\""; Type *string "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		DefaultPermissions   *int    "json:\"default_permissions,omitempty\""
		Flag                 *string "json:\"flag,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NetworkName          *string "json:\"network_name,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		ScanCount            *int    "json:\"scan_count,omitempty\""
		ScannerCount         *string "json:\"scanner_count,omitempty\""
		ScannerId            *int    "json:\"scanner_id,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		Token                *string "json:\"token,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Tags implementation of the Tags interface
type Tags struct {
	*ClientWithResponses
}

// AssignAssetTagsWithBody calls the Tags ´s function
func (c *Tags) AssignAssetTagsWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*TagsAssignAssetTagsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsAssignAssetTagsWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *TagsAssignAssetTagsResponse
	if i, ok := r.(*TagsAssignAssetTagsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AssignAssetTags calls the Tags ´s function
func (c *Tags) AssignAssetTags(arg1 TagsAssignAssetTagsJSONRequestBody, reqEditors ...RequestEditorFn) (*TagsAssignAssetTagsResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsAssignAssetTagsWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *TagsAssignAssetTagsResponse
	if i, ok := r.(*TagsAssignAssetTagsResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CreateTagCategoryWithBody calls the Tags ´s function
func (c *Tags) CreateTagCategoryWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	CreatedAt   *string "json:\"created_at,omitempty\""
	CreatedBy   *string "json:\"created_by,omitempty\""
	Description *string "json:\"description,omitempty\""
	Name        *string "json:\"name,omitempty\""
	Reserved    *bool   "json:\"reserved,omitempty\""
	UpdatedAt   *string "json:\"updated_at,omitempty\""
	UpdatedBy   *string "json:\"updated_by,omitempty\""
	Uuid        *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsCreateTagCategoryWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreatedAt *string "json:\"created_at,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; Description *string "json:\"description,omitempty\""; Name *string "json:\"name,omitempty\""; Reserved *bool "json:\"reserved,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedBy *string "json:\"updated_by,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		CreatedAt   *string "json:\"created_at,omitempty\""
		CreatedBy   *string "json:\"created_by,omitempty\""
		Description *string "json:\"description,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Reserved    *bool   "json:\"reserved,omitempty\""
		UpdatedAt   *string "json:\"updated_at,omitempty\""
		UpdatedBy   *string "json:\"updated_by,omitempty\""
		Uuid        *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CreateTagCategory calls the Tags ´s function
func (c *Tags) CreateTagCategory(arg1 TagsCreateTagCategoryJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	CreatedAt   *string "json:\"created_at,omitempty\""
	CreatedBy   *string "json:\"created_by,omitempty\""
	Description *string "json:\"description,omitempty\""
	Name        *string "json:\"name,omitempty\""
	Reserved    *bool   "json:\"reserved,omitempty\""
	UpdatedAt   *string "json:\"updated_at,omitempty\""
	UpdatedBy   *string "json:\"updated_by,omitempty\""
	Uuid        *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsCreateTagCategoryWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreatedAt *string "json:\"created_at,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; Description *string "json:\"description,omitempty\""; Name *string "json:\"name,omitempty\""; Reserved *bool "json:\"reserved,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedBy *string "json:\"updated_by,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		CreatedAt   *string "json:\"created_at,omitempty\""
		CreatedBy   *string "json:\"created_by,omitempty\""
		Description *string "json:\"description,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Reserved    *bool   "json:\"reserved,omitempty\""
		UpdatedAt   *string "json:\"updated_at,omitempty\""
		UpdatedBy   *string "json:\"updated_by,omitempty\""
		Uuid        *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CreateTagValueWithBody calls the Tags ´s function
func (c *Tags) CreateTagValueWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	AccessControl *struct {
		AllUsersPermissions      *N200AccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""
		CurrentDomainPermissions *[]struct {
			Id          *string                                        "json:\"id,omitempty\""
			Name        *string                                        "json:\"name,omitempty\""
			Permissions *[]string                                      "json:\"permissions,omitempty\""
			Type        *N200AccessControlCurrentDomainPermissionsType "json:\"type,omitempty\""
		} "json:\"current_domain_permissions,omitempty\""
		CurrentUserPermissions   *N200AccessControlCurrentUserPermissions   "json:\"current_user_permissions,omitempty\""
		DefinedDomainPermissions *N200AccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""
		Version                  *int64                                     "json:\"version,omitempty\""
	} "json:\"access_control,omitempty\""
	CategoryDescription *string "json:\"category_description,omitempty\""
	CategoryName        *string "json:\"category_name,omitempty\""
	CategoryUuid        *string "json:\"category_uuid,omitempty\""
	CreatedAt           *string "json:\"created_at,omitempty\""
	CreatedBy           *string "json:\"created_by,omitempty\""
	Description         *string "json:\"description,omitempty\""
	Filters             *struct {
		Asset *string "json:\"asset,omitempty\""
	} "json:\"filters,omitempty\""
	Type      *string "json:\"type,omitempty\""
	UpdatedAt *string "json:\"updated_at,omitempty\""
	UpdatedBy *string "json:\"updated_by,omitempty\""
	Uuid      *string "json:\"uuid,omitempty\""
	Value     *string "json:\"value,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsCreateTagValueWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AccessControl *struct { AllUsersPermissions *N200AccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""; CurrentDomainPermissions *[]struct { Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Permissions *[]string "json:\"permissions,omitempty\""; Type *N200AccessControlCurrentDomainPermissionsType "json:\"type,omitempty\"" } "json:\"current_domain_permissions,omitempty\""; CurrentUserPermissions *N200AccessControlCurrentUserPermissions "json:\"current_user_permissions,omitempty\""; DefinedDomainPermissions *N200AccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""; Version *int64 "json:\"version,omitempty\"" } "json:\"access_control,omitempty\""; CategoryDescription *string "json:\"category_description,omitempty\""; CategoryName *string "json:\"category_name,omitempty\""; CategoryUuid *string "json:\"category_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; Description *string "json:\"description,omitempty\""; Filters *struct { Asset *string "json:\"asset,omitempty\"" } "json:\"filters,omitempty\""; Type *string "json:\"type,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedBy *string "json:\"updated_by,omitempty\""; Uuid *string "json:\"uuid,omitempty\""; Value *string "json:\"value,omitempty\"" }
	if i, ok := r.(*struct {
		AccessControl *struct {
			AllUsersPermissions      *N200AccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""
			CurrentDomainPermissions *[]struct {
				Id          *string                                        "json:\"id,omitempty\""
				Name        *string                                        "json:\"name,omitempty\""
				Permissions *[]string                                      "json:\"permissions,omitempty\""
				Type        *N200AccessControlCurrentDomainPermissionsType "json:\"type,omitempty\""
			} "json:\"current_domain_permissions,omitempty\""
			CurrentUserPermissions   *N200AccessControlCurrentUserPermissions   "json:\"current_user_permissions,omitempty\""
			DefinedDomainPermissions *N200AccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""
			Version                  *int64                                     "json:\"version,omitempty\""
		} "json:\"access_control,omitempty\""
		CategoryDescription *string "json:\"category_description,omitempty\""
		CategoryName        *string "json:\"category_name,omitempty\""
		CategoryUuid        *string "json:\"category_uuid,omitempty\""
		CreatedAt           *string "json:\"created_at,omitempty\""
		CreatedBy           *string "json:\"created_by,omitempty\""
		Description         *string "json:\"description,omitempty\""
		Filters             *struct {
			Asset *string "json:\"asset,omitempty\""
		} "json:\"filters,omitempty\""
		Type      *string "json:\"type,omitempty\""
		UpdatedAt *string "json:\"updated_at,omitempty\""
		UpdatedBy *string "json:\"updated_by,omitempty\""
		Uuid      *string "json:\"uuid,omitempty\""
		Value     *string "json:\"value,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CreateTagValue calls the Tags ´s function
func (c *Tags) CreateTagValue(arg1 TagsCreateTagValueJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	AccessControl *struct {
		AllUsersPermissions      *N200AccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""
		CurrentDomainPermissions *[]struct {
			Id          *string                                        "json:\"id,omitempty\""
			Name        *string                                        "json:\"name,omitempty\""
			Permissions *[]string                                      "json:\"permissions,omitempty\""
			Type        *N200AccessControlCurrentDomainPermissionsType "json:\"type,omitempty\""
		} "json:\"current_domain_permissions,omitempty\""
		CurrentUserPermissions   *N200AccessControlCurrentUserPermissions   "json:\"current_user_permissions,omitempty\""
		DefinedDomainPermissions *N200AccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""
		Version                  *int64                                     "json:\"version,omitempty\""
	} "json:\"access_control,omitempty\""
	CategoryDescription *string "json:\"category_description,omitempty\""
	CategoryName        *string "json:\"category_name,omitempty\""
	CategoryUuid        *string "json:\"category_uuid,omitempty\""
	CreatedAt           *string "json:\"created_at,omitempty\""
	CreatedBy           *string "json:\"created_by,omitempty\""
	Description         *string "json:\"description,omitempty\""
	Filters             *struct {
		Asset *string "json:\"asset,omitempty\""
	} "json:\"filters,omitempty\""
	Type      *string "json:\"type,omitempty\""
	UpdatedAt *string "json:\"updated_at,omitempty\""
	UpdatedBy *string "json:\"updated_by,omitempty\""
	Uuid      *string "json:\"uuid,omitempty\""
	Value     *string "json:\"value,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsCreateTagValueWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AccessControl *struct { AllUsersPermissions *N200AccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""; CurrentDomainPermissions *[]struct { Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Permissions *[]string "json:\"permissions,omitempty\""; Type *N200AccessControlCurrentDomainPermissionsType "json:\"type,omitempty\"" } "json:\"current_domain_permissions,omitempty\""; CurrentUserPermissions *N200AccessControlCurrentUserPermissions "json:\"current_user_permissions,omitempty\""; DefinedDomainPermissions *N200AccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""; Version *int64 "json:\"version,omitempty\"" } "json:\"access_control,omitempty\""; CategoryDescription *string "json:\"category_description,omitempty\""; CategoryName *string "json:\"category_name,omitempty\""; CategoryUuid *string "json:\"category_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; Description *string "json:\"description,omitempty\""; Filters *struct { Asset *string "json:\"asset,omitempty\"" } "json:\"filters,omitempty\""; Type *string "json:\"type,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedBy *string "json:\"updated_by,omitempty\""; Uuid *string "json:\"uuid,omitempty\""; Value *string "json:\"value,omitempty\"" }
	if i, ok := r.(*struct {
		AccessControl *struct {
			AllUsersPermissions      *N200AccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""
			CurrentDomainPermissions *[]struct {
				Id          *string                                        "json:\"id,omitempty\""
				Name        *string                                        "json:\"name,omitempty\""
				Permissions *[]string                                      "json:\"permissions,omitempty\""
				Type        *N200AccessControlCurrentDomainPermissionsType "json:\"type,omitempty\""
			} "json:\"current_domain_permissions,omitempty\""
			CurrentUserPermissions   *N200AccessControlCurrentUserPermissions   "json:\"current_user_permissions,omitempty\""
			DefinedDomainPermissions *N200AccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""
			Version                  *int64                                     "json:\"version,omitempty\""
		} "json:\"access_control,omitempty\""
		CategoryDescription *string "json:\"category_description,omitempty\""
		CategoryName        *string "json:\"category_name,omitempty\""
		CategoryUuid        *string "json:\"category_uuid,omitempty\""
		CreatedAt           *string "json:\"created_at,omitempty\""
		CreatedBy           *string "json:\"created_by,omitempty\""
		Description         *string "json:\"description,omitempty\""
		Filters             *struct {
			Asset *string "json:\"asset,omitempty\""
		} "json:\"filters,omitempty\""
		Type      *string "json:\"type,omitempty\""
		UpdatedAt *string "json:\"updated_at,omitempty\""
		UpdatedBy *string "json:\"updated_by,omitempty\""
		Uuid      *string "json:\"uuid,omitempty\""
		Value     *string "json:\"value,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DeleteTagCategory calls the Tags ´s function
func (c *Tags) DeleteTagCategory(contentType string, reqEditors ...RequestEditorFn) (*TagsDeleteTagCategoryResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsDeleteTagCategoryWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *TagsDeleteTagCategoryResponse
	if i, ok := r.(*TagsDeleteTagCategoryResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DeleteTagValue calls the Tags ´s function
func (c *Tags) DeleteTagValue(contentType string, reqEditors ...RequestEditorFn) (*TagsDeleteTagValueResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsDeleteTagValueWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *TagsDeleteTagValueResponse
	if i, ok := r.(*TagsDeleteTagValueResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DeleteTagValuesBulkWithBody calls the Tags ´s function
func (c *Tags) DeleteTagValuesBulkWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsDeleteTagValuesBulkWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// DeleteTagValuesBulk calls the Tags ´s function
func (c *Tags) DeleteTagValuesBulk(arg1 TagsDeleteTagValuesBulkJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsDeleteTagValuesBulkWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// EditTagCategoryWithBody calls the Tags ´s function
func (c *Tags) EditTagCategoryWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	CreatedAt   *string "json:\"created_at,omitempty\""
	CreatedBy   *string "json:\"created_by,omitempty\""
	Description *string "json:\"description,omitempty\""
	Name        *string "json:\"name,omitempty\""
	Reserved    *bool   "json:\"reserved,omitempty\""
	UpdatedAt   *string "json:\"updated_at,omitempty\""
	UpdatedBy   *string "json:\"updated_by,omitempty\""
	Uuid        *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsEditTagCategoryWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreatedAt *string "json:\"created_at,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; Description *string "json:\"description,omitempty\""; Name *string "json:\"name,omitempty\""; Reserved *bool "json:\"reserved,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedBy *string "json:\"updated_by,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		CreatedAt   *string "json:\"created_at,omitempty\""
		CreatedBy   *string "json:\"created_by,omitempty\""
		Description *string "json:\"description,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Reserved    *bool   "json:\"reserved,omitempty\""
		UpdatedAt   *string "json:\"updated_at,omitempty\""
		UpdatedBy   *string "json:\"updated_by,omitempty\""
		Uuid        *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// EditTagCategory calls the Tags ´s function
func (c *Tags) EditTagCategory(arg1 string, arg2 TagsEditTagCategoryJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	CreatedAt   *string "json:\"created_at,omitempty\""
	CreatedBy   *string "json:\"created_by,omitempty\""
	Description *string "json:\"description,omitempty\""
	Name        *string "json:\"name,omitempty\""
	Reserved    *bool   "json:\"reserved,omitempty\""
	UpdatedAt   *string "json:\"updated_at,omitempty\""
	UpdatedBy   *string "json:\"updated_by,omitempty\""
	Uuid        *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsEditTagCategoryWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreatedAt *string "json:\"created_at,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; Description *string "json:\"description,omitempty\""; Name *string "json:\"name,omitempty\""; Reserved *bool "json:\"reserved,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedBy *string "json:\"updated_by,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		CreatedAt   *string "json:\"created_at,omitempty\""
		CreatedBy   *string "json:\"created_by,omitempty\""
		Description *string "json:\"description,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Reserved    *bool   "json:\"reserved,omitempty\""
		UpdatedAt   *string "json:\"updated_at,omitempty\""
		UpdatedBy   *string "json:\"updated_by,omitempty\""
		Uuid        *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ListAssetFilters calls the Tags ´s function
func (c *Tags) ListAssetFilters(reqEditors ...RequestEditorFn) (*[]struct {
	Control *struct {
		ReadableRegex *string "json:\"readable_regex,omitempty\""
		Regex         *string "json:\"regex,omitempty\""
		Type          *string "json:\"type,omitempty\""
	} "json:\"control,omitempty\""
	Name         *string   "json:\"name,omitempty\""
	Operators    *[]string "json:\"operators,omitempty\""
	ReadableName *string   "json:\"readable_name,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsListAssetFiltersWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { Control *struct { ReadableRegex *string "json:\"readable_regex,omitempty\""; Regex *string "json:\"regex,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"control,omitempty\""; Name *string "json:\"name,omitempty\""; Operators *[]string "json:\"operators,omitempty\""; ReadableName *string "json:\"readable_name,omitempty\"" }
	if i, ok := r.(*[]struct {
		Control *struct {
			ReadableRegex *string "json:\"readable_regex,omitempty\""
			Regex         *string "json:\"regex,omitempty\""
			Type          *string "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ListAssetTags calls the Tags ´s function
func (c *Tags) ListAssetTags(contentType string, reqEditors ...RequestEditorFn) (*struct {
	Tags *[]struct {
		AssetUuid    *string "json:\"asset_uuid,omitempty\""
		CategoryName *string "json:\"category_name,omitempty\""
		CategoryUuid *string "json:\"category_uuid,omitempty\""
		CreatedAt    *string "json:\"created_at,omitempty\""
		CreatedBy    *string "json:\"created_by,omitempty\""
		Source       *string "json:\"source,omitempty\""
		Value        *string "json:\"value,omitempty\""
		ValueUuid    *string "json:\"value_uuid,omitempty\""
	} "json:\"tags,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsListAssetTagsWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Tags *[]struct { AssetUuid *string "json:\"asset_uuid,omitempty\""; CategoryName *string "json:\"category_name,omitempty\""; CategoryUuid *string "json:\"category_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; Source *string "json:\"source,omitempty\""; Value *string "json:\"value,omitempty\""; ValueUuid *string "json:\"value_uuid,omitempty\"" } "json:\"tags,omitempty\"" }
	if i, ok := r.(*struct {
		Tags *[]struct {
			AssetUuid    *string "json:\"asset_uuid,omitempty\""
			CategoryName *string "json:\"category_name,omitempty\""
			CategoryUuid *string "json:\"category_uuid,omitempty\""
			CreatedAt    *string "json:\"created_at,omitempty\""
			CreatedBy    *string "json:\"created_by,omitempty\""
			Source       *string "json:\"source,omitempty\""
			Value        *string "json:\"value,omitempty\""
			ValueUuid    *string "json:\"value_uuid,omitempty\""
		} "json:\"tags,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ListTagCategories calls the Tags ´s function
func (c *Tags) ListTagCategories(params *TagsListTagCategoriesParams, reqEditors ...RequestEditorFn) (*struct {
	Categories *[]struct {
		CreatedAt   *string "json:\"created_at,omitempty\""
		CreatedBy   *string "json:\"created_by,omitempty\""
		Description *string "json:\"description,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Reserved    *bool   "json:\"reserved,omitempty\""
		UpdatedAt   *string "json:\"updated_at,omitempty\""
		UpdatedBy   *string "json:\"updated_by,omitempty\""
		Uuid        *string "json:\"uuid,omitempty\""
	} "json:\"categories,omitempty\""
	Pagination *struct {
		Limit  *int32 "json:\"limit,omitempty\""
		Offset *int32 "json:\"offset,omitempty\""
		Sort   *[]struct {
			Name  *string "json:\"name,omitempty\""
			Order *string "json:\"order,omitempty\""
		} "json:\"sort,omitempty\""
		Total *int32 "json:\"total,omitempty\""
	} "json:\"pagination,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsListTagCategoriesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Categories *[]struct { CreatedAt *string "json:\"created_at,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; Description *string "json:\"description,omitempty\""; Name *string "json:\"name,omitempty\""; Reserved *bool "json:\"reserved,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedBy *string "json:\"updated_by,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"categories,omitempty\""; Pagination *struct { Limit *int32 "json:\"limit,omitempty\""; Offset *int32 "json:\"offset,omitempty\""; Sort *[]struct { Name *string "json:\"name,omitempty\""; Order *string "json:\"order,omitempty\"" } "json:\"sort,omitempty\""; Total *int32 "json:\"total,omitempty\"" } "json:\"pagination,omitempty\"" }
	if i, ok := r.(*struct {
		Categories *[]struct {
			CreatedAt   *string "json:\"created_at,omitempty\""
			CreatedBy   *string "json:\"created_by,omitempty\""
			Description *string "json:\"description,omitempty\""
			Name        *string "json:\"name,omitempty\""
			Reserved    *bool   "json:\"reserved,omitempty\""
			UpdatedAt   *string "json:\"updated_at,omitempty\""
			UpdatedBy   *string "json:\"updated_by,omitempty\""
			Uuid        *string "json:\"uuid,omitempty\""
		} "json:\"categories,omitempty\""
		Pagination *struct {
			Limit  *int32 "json:\"limit,omitempty\""
			Offset *int32 "json:\"offset,omitempty\""
			Sort   *[]struct {
				Name  *string "json:\"name,omitempty\""
				Order *string "json:\"order,omitempty\""
			} "json:\"sort,omitempty\""
			Total *int32 "json:\"total,omitempty\""
		} "json:\"pagination,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ListTagValues calls the Tags ´s function
func (c *Tags) ListTagValues(params *TagsListTagValuesParams, reqEditors ...RequestEditorFn) (*struct {
	Pagination *struct {
		Limit  *int32 "json:\"limit,omitempty\""
		Offset *int32 "json:\"offset,omitempty\""
		Sort   *[]struct {
			Name  *string "json:\"name,omitempty\""
			Order *string "json:\"order,omitempty\""
		} "json:\"sort,omitempty\""
		Total *int32 "json:\"total,omitempty\""
	} "json:\"pagination,omitempty\""
	Values *[]struct {
		AccessControl *struct {
			AllUsersPermissions      *N200ValuesAccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""
			CurrentDomainPermissions *[]struct {
				Id          *string                                              "json:\"id,omitempty\""
				Name        *string                                              "json:\"name,omitempty\""
				Permissions *[]string                                            "json:\"permissions,omitempty\""
				Type        *N200ValuesAccessControlCurrentDomainPermissionsType "json:\"type,omitempty\""
			} "json:\"current_domain_permissions,omitempty\""
			CurrentUserPermissions   *N200ValuesAccessControlCurrentUserPermissions   "json:\"current_user_permissions,omitempty\""
			DefinedDomainPermissions *N200ValuesAccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""
			Version                  *int64                                           "json:\"version,omitempty\""
		} "json:\"access_control,omitempty\""
		CategoryDescription *string "json:\"category_description,omitempty\""
		CategoryName        *string "json:\"category_name,omitempty\""
		CategoryUuid        *string "json:\"category_uuid,omitempty\""
		CreatedAt           *string "json:\"created_at,omitempty\""
		CreatedBy           *string "json:\"created_by,omitempty\""
		Description         *string "json:\"description,omitempty\""
		Type                *string "json:\"type,omitempty\""
		UpdatedAt           *string "json:\"updated_at,omitempty\""
		UpdatedBy           *string "json:\"updated_by,omitempty\""
		Uuid                *string "json:\"uuid,omitempty\""
		Value               *string "json:\"value,omitempty\""
	} "json:\"values,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsListTagValuesWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Pagination *struct { Limit *int32 "json:\"limit,omitempty\""; Offset *int32 "json:\"offset,omitempty\""; Sort *[]struct { Name *string "json:\"name,omitempty\""; Order *string "json:\"order,omitempty\"" } "json:\"sort,omitempty\""; Total *int32 "json:\"total,omitempty\"" } "json:\"pagination,omitempty\""; Values *[]struct { AccessControl *struct { AllUsersPermissions *N200ValuesAccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""; CurrentDomainPermissions *[]struct { Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Permissions *[]string "json:\"permissions,omitempty\""; Type *N200ValuesAccessControlCurrentDomainPermissionsType "json:\"type,omitempty\"" } "json:\"current_domain_permissions,omitempty\""; CurrentUserPermissions *N200ValuesAccessControlCurrentUserPermissions "json:\"current_user_permissions,omitempty\""; DefinedDomainPermissions *N200ValuesAccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""; Version *int64 "json:\"version,omitempty\"" } "json:\"access_control,omitempty\""; CategoryDescription *string "json:\"category_description,omitempty\""; CategoryName *string "json:\"category_name,omitempty\""; CategoryUuid *string "json:\"category_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; Description *string "json:\"description,omitempty\""; Type *string "json:\"type,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedBy *string "json:\"updated_by,omitempty\""; Uuid *string "json:\"uuid,omitempty\""; Value *string "json:\"value,omitempty\"" } "json:\"values,omitempty\"" }
	if i, ok := r.(*struct {
		Pagination *struct {
			Limit  *int32 "json:\"limit,omitempty\""
			Offset *int32 "json:\"offset,omitempty\""
			Sort   *[]struct {
				Name  *string "json:\"name,omitempty\""
				Order *string "json:\"order,omitempty\""
			} "json:\"sort,omitempty\""
			Total *int32 "json:\"total,omitempty\""
		} "json:\"pagination,omitempty\""
		Values *[]struct {
			AccessControl *struct {
				AllUsersPermissions      *N200ValuesAccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""
				CurrentDomainPermissions *[]struct {
					Id          *string                                              "json:\"id,omitempty\""
					Name        *string                                              "json:\"name,omitempty\""
					Permissions *[]string                                            "json:\"permissions,omitempty\""
					Type        *N200ValuesAccessControlCurrentDomainPermissionsType "json:\"type,omitempty\""
				} "json:\"current_domain_permissions,omitempty\""
				CurrentUserPermissions   *N200ValuesAccessControlCurrentUserPermissions   "json:\"current_user_permissions,omitempty\""
				DefinedDomainPermissions *N200ValuesAccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""
				Version                  *int64                                           "json:\"version,omitempty\""
			} "json:\"access_control,omitempty\""
			CategoryDescription *string "json:\"category_description,omitempty\""
			CategoryName        *string "json:\"category_name,omitempty\""
			CategoryUuid        *string "json:\"category_uuid,omitempty\""
			CreatedAt           *string "json:\"created_at,omitempty\""
			CreatedBy           *string "json:\"created_by,omitempty\""
			Description         *string "json:\"description,omitempty\""
			Type                *string "json:\"type,omitempty\""
			UpdatedAt           *string "json:\"updated_at,omitempty\""
			UpdatedBy           *string "json:\"updated_by,omitempty\""
			Uuid                *string "json:\"uuid,omitempty\""
			Value               *string "json:\"value,omitempty\""
		} "json:\"values,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// TagCategoryDetails calls the Tags ´s function
func (c *Tags) TagCategoryDetails(contentType string, reqEditors ...RequestEditorFn) (*struct {
	CreatedAt   *string "json:\"created_at,omitempty\""
	CreatedBy   *string "json:\"created_by,omitempty\""
	Description *string "json:\"description,omitempty\""
	Name        *string "json:\"name,omitempty\""
	Reserved    *bool   "json:\"reserved,omitempty\""
	UpdatedAt   *string "json:\"updated_at,omitempty\""
	UpdatedBy   *string "json:\"updated_by,omitempty\""
	Uuid        *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsTagCategoryDetailsWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreatedAt *string "json:\"created_at,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; Description *string "json:\"description,omitempty\""; Name *string "json:\"name,omitempty\""; Reserved *bool "json:\"reserved,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedBy *string "json:\"updated_by,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		CreatedAt   *string "json:\"created_at,omitempty\""
		CreatedBy   *string "json:\"created_by,omitempty\""
		Description *string "json:\"description,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Reserved    *bool   "json:\"reserved,omitempty\""
		UpdatedAt   *string "json:\"updated_at,omitempty\""
		UpdatedBy   *string "json:\"updated_by,omitempty\""
		Uuid        *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// TagValueDetails calls the Tags ´s function
func (c *Tags) TagValueDetails(contentType string, reqEditors ...RequestEditorFn) (*struct {
	AccessControl *struct {
		AllUsersPermissions      *N200AccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""
		CurrentDomainPermissions *[]struct {
			Id          *string                                        "json:\"id,omitempty\""
			Name        *string                                        "json:\"name,omitempty\""
			Permissions *[]string                                      "json:\"permissions,omitempty\""
			Type        *N200AccessControlCurrentDomainPermissionsType "json:\"type,omitempty\""
		} "json:\"current_domain_permissions,omitempty\""
		CurrentUserPermissions   *N200AccessControlCurrentUserPermissions   "json:\"current_user_permissions,omitempty\""
		DefinedDomainPermissions *N200AccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""
		Version                  *int64                                     "json:\"version,omitempty\""
	} "json:\"access_control,omitempty\""
	CategoryDescription *string "json:\"category_description,omitempty\""
	CategoryName        *string "json:\"category_name,omitempty\""
	CategoryUuid        *string "json:\"category_uuid,omitempty\""
	CreatedAt           *string "json:\"created_at,omitempty\""
	CreatedBy           *string "json:\"created_by,omitempty\""
	Description         *string "json:\"description,omitempty\""
	Filters             *struct {
		Asset *string "json:\"asset,omitempty\""
	} "json:\"filters,omitempty\""
	Type      *string "json:\"type,omitempty\""
	UpdatedAt *string "json:\"updated_at,omitempty\""
	UpdatedBy *string "json:\"updated_by,omitempty\""
	Uuid      *string "json:\"uuid,omitempty\""
	Value     *string "json:\"value,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsTagValueDetailsWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AccessControl *struct { AllUsersPermissions *N200AccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""; CurrentDomainPermissions *[]struct { Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Permissions *[]string "json:\"permissions,omitempty\""; Type *N200AccessControlCurrentDomainPermissionsType "json:\"type,omitempty\"" } "json:\"current_domain_permissions,omitempty\""; CurrentUserPermissions *N200AccessControlCurrentUserPermissions "json:\"current_user_permissions,omitempty\""; DefinedDomainPermissions *N200AccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""; Version *int64 "json:\"version,omitempty\"" } "json:\"access_control,omitempty\""; CategoryDescription *string "json:\"category_description,omitempty\""; CategoryName *string "json:\"category_name,omitempty\""; CategoryUuid *string "json:\"category_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; Description *string "json:\"description,omitempty\""; Filters *struct { Asset *string "json:\"asset,omitempty\"" } "json:\"filters,omitempty\""; Type *string "json:\"type,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedBy *string "json:\"updated_by,omitempty\""; Uuid *string "json:\"uuid,omitempty\""; Value *string "json:\"value,omitempty\"" }
	if i, ok := r.(*struct {
		AccessControl *struct {
			AllUsersPermissions      *N200AccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""
			CurrentDomainPermissions *[]struct {
				Id          *string                                        "json:\"id,omitempty\""
				Name        *string                                        "json:\"name,omitempty\""
				Permissions *[]string                                      "json:\"permissions,omitempty\""
				Type        *N200AccessControlCurrentDomainPermissionsType "json:\"type,omitempty\""
			} "json:\"current_domain_permissions,omitempty\""
			CurrentUserPermissions   *N200AccessControlCurrentUserPermissions   "json:\"current_user_permissions,omitempty\""
			DefinedDomainPermissions *N200AccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""
			Version                  *int64                                     "json:\"version,omitempty\""
		} "json:\"access_control,omitempty\""
		CategoryDescription *string "json:\"category_description,omitempty\""
		CategoryName        *string "json:\"category_name,omitempty\""
		CategoryUuid        *string "json:\"category_uuid,omitempty\""
		CreatedAt           *string "json:\"created_at,omitempty\""
		CreatedBy           *string "json:\"created_by,omitempty\""
		Description         *string "json:\"description,omitempty\""
		Filters             *struct {
			Asset *string "json:\"asset,omitempty\""
		} "json:\"filters,omitempty\""
		Type      *string "json:\"type,omitempty\""
		UpdatedAt *string "json:\"updated_at,omitempty\""
		UpdatedBy *string "json:\"updated_by,omitempty\""
		Uuid      *string "json:\"uuid,omitempty\""
		Value     *string "json:\"value,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UpdateTagValueWithBody calls the Tags ´s function
func (c *Tags) UpdateTagValueWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	AccessControl *struct {
		AllUsersPermissions      *N200AccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""
		CurrentDomainPermissions *[]struct {
			Id          *string                                        "json:\"id,omitempty\""
			Name        *string                                        "json:\"name,omitempty\""
			Permissions *[]string                                      "json:\"permissions,omitempty\""
			Type        *N200AccessControlCurrentDomainPermissionsType "json:\"type,omitempty\""
		} "json:\"current_domain_permissions,omitempty\""
		CurrentUserPermissions   *N200AccessControlCurrentUserPermissions   "json:\"current_user_permissions,omitempty\""
		DefinedDomainPermissions *N200AccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""
		Version                  *int64                                     "json:\"version,omitempty\""
	} "json:\"access_control,omitempty\""
	CategoryDescription *string "json:\"category_description,omitempty\""
	CategoryName        *string "json:\"category_name,omitempty\""
	CategoryUuid        *string "json:\"category_uuid,omitempty\""
	CreatedAt           *string "json:\"created_at,omitempty\""
	CreatedBy           *string "json:\"created_by,omitempty\""
	Description         *string "json:\"description,omitempty\""
	Filters             *struct {
		Asset *string "json:\"asset,omitempty\""
	} "json:\"filters,omitempty\""
	Type      *string "json:\"type,omitempty\""
	UpdatedAt *string "json:\"updated_at,omitempty\""
	UpdatedBy *string "json:\"updated_by,omitempty\""
	Uuid      *string "json:\"uuid,omitempty\""
	Value     *string "json:\"value,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsUpdateTagValueWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AccessControl *struct { AllUsersPermissions *N200AccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""; CurrentDomainPermissions *[]struct { Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Permissions *[]string "json:\"permissions,omitempty\""; Type *N200AccessControlCurrentDomainPermissionsType "json:\"type,omitempty\"" } "json:\"current_domain_permissions,omitempty\""; CurrentUserPermissions *N200AccessControlCurrentUserPermissions "json:\"current_user_permissions,omitempty\""; DefinedDomainPermissions *N200AccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""; Version *int64 "json:\"version,omitempty\"" } "json:\"access_control,omitempty\""; CategoryDescription *string "json:\"category_description,omitempty\""; CategoryName *string "json:\"category_name,omitempty\""; CategoryUuid *string "json:\"category_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; Description *string "json:\"description,omitempty\""; Filters *struct { Asset *string "json:\"asset,omitempty\"" } "json:\"filters,omitempty\""; Type *string "json:\"type,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedBy *string "json:\"updated_by,omitempty\""; Uuid *string "json:\"uuid,omitempty\""; Value *string "json:\"value,omitempty\"" }
	if i, ok := r.(*struct {
		AccessControl *struct {
			AllUsersPermissions      *N200AccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""
			CurrentDomainPermissions *[]struct {
				Id          *string                                        "json:\"id,omitempty\""
				Name        *string                                        "json:\"name,omitempty\""
				Permissions *[]string                                      "json:\"permissions,omitempty\""
				Type        *N200AccessControlCurrentDomainPermissionsType "json:\"type,omitempty\""
			} "json:\"current_domain_permissions,omitempty\""
			CurrentUserPermissions   *N200AccessControlCurrentUserPermissions   "json:\"current_user_permissions,omitempty\""
			DefinedDomainPermissions *N200AccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""
			Version                  *int64                                     "json:\"version,omitempty\""
		} "json:\"access_control,omitempty\""
		CategoryDescription *string "json:\"category_description,omitempty\""
		CategoryName        *string "json:\"category_name,omitempty\""
		CategoryUuid        *string "json:\"category_uuid,omitempty\""
		CreatedAt           *string "json:\"created_at,omitempty\""
		CreatedBy           *string "json:\"created_by,omitempty\""
		Description         *string "json:\"description,omitempty\""
		Filters             *struct {
			Asset *string "json:\"asset,omitempty\""
		} "json:\"filters,omitempty\""
		Type      *string "json:\"type,omitempty\""
		UpdatedAt *string "json:\"updated_at,omitempty\""
		UpdatedBy *string "json:\"updated_by,omitempty\""
		Uuid      *string "json:\"uuid,omitempty\""
		Value     *string "json:\"value,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UpdateTagValue calls the Tags ´s function
func (c *Tags) UpdateTagValue(arg1 string, arg2 TagsUpdateTagValueJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	AccessControl *struct {
		AllUsersPermissions      *N200AccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""
		CurrentDomainPermissions *[]struct {
			Id          *string                                        "json:\"id,omitempty\""
			Name        *string                                        "json:\"name,omitempty\""
			Permissions *[]string                                      "json:\"permissions,omitempty\""
			Type        *N200AccessControlCurrentDomainPermissionsType "json:\"type,omitempty\""
		} "json:\"current_domain_permissions,omitempty\""
		CurrentUserPermissions   *N200AccessControlCurrentUserPermissions   "json:\"current_user_permissions,omitempty\""
		DefinedDomainPermissions *N200AccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""
		Version                  *int64                                     "json:\"version,omitempty\""
	} "json:\"access_control,omitempty\""
	CategoryDescription *string "json:\"category_description,omitempty\""
	CategoryName        *string "json:\"category_name,omitempty\""
	CategoryUuid        *string "json:\"category_uuid,omitempty\""
	CreatedAt           *string "json:\"created_at,omitempty\""
	CreatedBy           *string "json:\"created_by,omitempty\""
	Description         *string "json:\"description,omitempty\""
	Filters             *struct {
		Asset *string "json:\"asset,omitempty\""
	} "json:\"filters,omitempty\""
	Type      *string "json:\"type,omitempty\""
	UpdatedAt *string "json:\"updated_at,omitempty\""
	UpdatedBy *string "json:\"updated_by,omitempty\""
	Uuid      *string "json:\"uuid,omitempty\""
	Value     *string "json:\"value,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.TagsUpdateTagValueWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AccessControl *struct { AllUsersPermissions *N200AccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""; CurrentDomainPermissions *[]struct { Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Permissions *[]string "json:\"permissions,omitempty\""; Type *N200AccessControlCurrentDomainPermissionsType "json:\"type,omitempty\"" } "json:\"current_domain_permissions,omitempty\""; CurrentUserPermissions *N200AccessControlCurrentUserPermissions "json:\"current_user_permissions,omitempty\""; DefinedDomainPermissions *N200AccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""; Version *int64 "json:\"version,omitempty\"" } "json:\"access_control,omitempty\""; CategoryDescription *string "json:\"category_description,omitempty\""; CategoryName *string "json:\"category_name,omitempty\""; CategoryUuid *string "json:\"category_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; Description *string "json:\"description,omitempty\""; Filters *struct { Asset *string "json:\"asset,omitempty\"" } "json:\"filters,omitempty\""; Type *string "json:\"type,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\""; UpdatedBy *string "json:\"updated_by,omitempty\""; Uuid *string "json:\"uuid,omitempty\""; Value *string "json:\"value,omitempty\"" }
	if i, ok := r.(*struct {
		AccessControl *struct {
			AllUsersPermissions      *N200AccessControlAllUsersPermissions "json:\"all_users_permissions,omitempty\""
			CurrentDomainPermissions *[]struct {
				Id          *string                                        "json:\"id,omitempty\""
				Name        *string                                        "json:\"name,omitempty\""
				Permissions *[]string                                      "json:\"permissions,omitempty\""
				Type        *N200AccessControlCurrentDomainPermissionsType "json:\"type,omitempty\""
			} "json:\"current_domain_permissions,omitempty\""
			CurrentUserPermissions   *N200AccessControlCurrentUserPermissions   "json:\"current_user_permissions,omitempty\""
			DefinedDomainPermissions *N200AccessControlDefinedDomainPermissions "json:\"defined_domain_permissions,omitempty\""
			Version                  *int64                                     "json:\"version,omitempty\""
		} "json:\"access_control,omitempty\""
		CategoryDescription *string "json:\"category_description,omitempty\""
		CategoryName        *string "json:\"category_name,omitempty\""
		CategoryUuid        *string "json:\"category_uuid,omitempty\""
		CreatedAt           *string "json:\"created_at,omitempty\""
		CreatedBy           *string "json:\"created_by,omitempty\""
		Description         *string "json:\"description,omitempty\""
		Filters             *struct {
			Asset *string "json:\"asset,omitempty\""
		} "json:\"filters,omitempty\""
		Type      *string "json:\"type,omitempty\""
		UpdatedAt *string "json:\"updated_at,omitempty\""
		UpdatedBy *string "json:\"updated_by,omitempty\""
		Uuid      *string "json:\"uuid,omitempty\""
		Value     *string "json:\"value,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Bulk implementation of the Bulk interface
type Bulk struct {
	*ClientWithResponses
}

// AddAgentsWithBody calls the Bulk ´s function
func (c *Bulk) AddAgentsWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.BulkAddAgentsWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AddAgents calls the Bulk ´s function
func (c *Bulk) AddAgents(arg1 string, arg2 BulkAddAgentsJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.BulkAddAgentsWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RemoveAgentsWithBody calls the Bulk ´s function
func (c *Bulk) RemoveAgentsWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.BulkRemoveAgentsWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// RemoveAgents calls the Bulk ´s function
func (c *Bulk) RemoveAgents(arg1 int32, arg2 BulkRemoveAgentsJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.BulkRemoveAgentsWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// TaskAgentGroupStatus calls the Bulk ´s function
func (c *Bulk) TaskAgentGroupStatus(arg1 int32, contentType string, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.BulkTaskAgentGroupStatusWithResponse(c.ClientInterface.(*Client).ctx, arg1, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// TaskAgentStatus calls the Bulk ´s function
func (c *Bulk) TaskAgentStatus(contentType string, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.BulkTaskAgentStatusWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UnlinkAgentsWithBody calls the Bulk ´s function
func (c *Bulk) UnlinkAgentsWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.BulkUnlinkAgentsWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UnlinkAgents calls the Bulk ´s function
func (c *Bulk) UnlinkAgents(arg1 BulkUnlinkAgentsJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
	ContainerUuid           *string     "json:\"container_uuid,omitempty\""
	EndTime                 *int        "json:\"end_time,omitempty\""
	LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
	Message                 *string     "json:\"message,omitempty\""
	StartTime               *int        "json:\"start_time,omitempty\""
	Status                  *N200Status "json:\"status,omitempty\""
	TaskId                  *string     "json:\"task_id,omitempty\""
	TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
	TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.BulkUnlinkAgentsWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CompletionPercentage *int "json:\"completion_percentage,omitempty\""; ContainerUuid *string "json:\"container_uuid,omitempty\""; EndTime *int "json:\"end_time,omitempty\""; LastUpdateTime *int "json:\"last_update_time,omitempty\""; Message *string "json:\"message,omitempty\""; StartTime *int "json:\"start_time,omitempty\""; Status *N200Status "json:\"status,omitempty\""; TaskId *string "json:\"task_id,omitempty\""; TotalWorkUnits *int "json:\"total_work_units,omitempty\""; TotalWorkUnitsCompleted *int "json:\"total_work_units_completed,omitempty\"" }
	if i, ok := r.(*struct {
		CompletionPercentage    *int        "json:\"completion_percentage,omitempty\""
		ContainerUuid           *string     "json:\"container_uuid,omitempty\""
		EndTime                 *int        "json:\"end_time,omitempty\""
		LastUpdateTime          *int        "json:\"last_update_time,omitempty\""
		Message                 *string     "json:\"message,omitempty\""
		StartTime               *int        "json:\"start_time,omitempty\""
		Status                  *N200Status "json:\"status,omitempty\""
		TaskId                  *string     "json:\"task_id,omitempty\""
		TotalWorkUnits          *int        "json:\"total_work_units,omitempty\""
		TotalWorkUnitsCompleted *int        "json:\"total_work_units_completed,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Credentials implementation of the Credentials interface
type Credentials struct {
	*ClientWithResponses
}

// CreateWithBody calls the Credentials ´s function
func (c *Credentials) CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	Uuid *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CredentialsCreateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		Uuid *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Create calls the Credentials ´s function
func (c *Credentials) Create(arg1 CredentialsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	Uuid *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CredentialsCreateWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		Uuid *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Delete calls the Credentials ´s function
func (c *Credentials) Delete(contentType string, reqEditors ...RequestEditorFn) (*struct {
	Deleted *bool "json:\"deleted,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CredentialsDeleteWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Deleted *bool "json:\"deleted,omitempty\"" }
	if i, ok := r.(*struct {
		Deleted *bool "json:\"deleted,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Details calls the Credentials ´s function
func (c *Credentials) Details(contentType string, reqEditors ...RequestEditorFn) (*struct {
	AdHoc    *bool "json:\"ad_hoc,omitempty\""
	Category *struct {
		Id   *string "json:\"id,omitempty\""
		Name *string "json:\"name,omitempty\""
	} "json:\"category,omitempty\""
	Description *string "json:\"description,omitempty\""
	Name        *string "json:\"name,omitempty\""
	Permissions *[]struct {
		GranteeUuid *string              "json:\"grantee_uuid,omitempty\""
		Name        *string              "json:\"name,omitempty\""
		Permissions *int                 "json:\"permissions,omitempty\""
		Type        *N200PermissionsType "json:\"type,omitempty\""
	} "json:\"permissions,omitempty\""
	Settings *map[string]interface{} "json:\"settings,omitempty\""
	Type     *struct {
		Id   *string "json:\"id,omitempty\""
		Name *string "json:\"name,omitempty\""
	} "json:\"type,omitempty\""
	UserPermissions *int "json:\"user_permissions,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CredentialsDetailsWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AdHoc *bool "json:\"ad_hoc,omitempty\""; Category *struct { Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"category,omitempty\""; Description *string "json:\"description,omitempty\""; Name *string "json:\"name,omitempty\""; Permissions *[]struct { GranteeUuid *string "json:\"grantee_uuid,omitempty\""; Name *string "json:\"name,omitempty\""; Permissions *int "json:\"permissions,omitempty\""; Type *N200PermissionsType "json:\"type,omitempty\"" } "json:\"permissions,omitempty\""; Settings *map[string]interface {} "json:\"settings,omitempty\""; Type *struct { Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\"" }
	if i, ok := r.(*struct {
		AdHoc    *bool "json:\"ad_hoc,omitempty\""
		Category *struct {
			Id   *string "json:\"id,omitempty\""
			Name *string "json:\"name,omitempty\""
		} "json:\"category,omitempty\""
		Description *string "json:\"description,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Permissions *[]struct {
			GranteeUuid *string              "json:\"grantee_uuid,omitempty\""
			Name        *string              "json:\"name,omitempty\""
			Permissions *int                 "json:\"permissions,omitempty\""
			Type        *N200PermissionsType "json:\"type,omitempty\""
		} "json:\"permissions,omitempty\""
		Settings *map[string]interface{} "json:\"settings,omitempty\""
		Type     *struct {
			Id   *string "json:\"id,omitempty\""
			Name *string "json:\"name,omitempty\""
		} "json:\"type,omitempty\""
		UserPermissions *int "json:\"user_permissions,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// FileUploadWithBody calls the Credentials ´s function
func (c *Credentials) FileUploadWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	Fileuploaded *string "json:\"fileuploaded,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CredentialsFileUploadWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Fileuploaded *string "json:\"fileuploaded,omitempty\"" }
	if i, ok := r.(*struct {
		Fileuploaded *string "json:\"fileuploaded,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ListCredentialTypes calls the Credentials ´s function
func (c *Credentials) ListCredentialTypes(reqEditors ...RequestEditorFn) (*struct {
	Credentials *[]struct {
		Category      *string "json:\"category,omitempty\""
		DefaultExpand *bool   "json:\"default_expand,omitempty\""
		Id            *string "json:\"id,omitempty\""
		Types         *[]struct {
			Configuration *[]struct {
				AltIds  *string "json:\"alt_ids,omitempty\""
				Default *string "json:\"default,omitempty\""
				Id      *string "json:\"id,omitempty\""
				Name    *string "json:\"name,omitempty\""
				Options *[]struct {
					Id     *string "json:\"id,omitempty\""
					Inputs *[]struct {
						Callback        *string "json:\"callback,omitempty\""
						DefaultRowCount *int    "json:\"default-row-count,omitempty\""
						HideValues      *bool   "json:\"hide-values,omitempty\""
						Hint            *string "json:\"hint,omitempty\""
						Id              *string "json:\"id,omitempty\""
						Name            *string "json:\"name,omitempty\""
						Placeholder     *string "json:\"placeholder,omitempty\""
						Regex           *string "json:\"regex,omitempty\""
						Required        *bool   "json:\"required,omitempty\""
						Type            *string "json:\"type,omitempty\""
					} "json:\"inputs,omitempty\""
					Name *string "json:\"name,omitempty\""
				} "json:\"options,omitempty\""
				Placeholder *string   "json:\"placeholder,omitempty\""
				Preferences *[]string "json:\"preferences,omitempty\""
				Required    *bool     "json:\"required,omitempty\""
				Type        *string   "json:\"type,omitempty\""
			} "json:\"configuration,omitempty\""
			ExpandSettings *bool   "json:\"expand_settings,omitempty\""
			Id             *string "json:\"id,omitempty\""
			Max            *int    "json:\"max,omitempty\""
			Name           *string "json:\"name,omitempty\""
		} "json:\"types,omitempty\""
	} "json:\"credentials,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CredentialsListCredentialTypesWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Credentials *[]struct { Category *string "json:\"category,omitempty\""; DefaultExpand *bool "json:\"default_expand,omitempty\""; Id *string "json:\"id,omitempty\""; Types *[]struct { Configuration *[]struct { AltIds *string "json:\"alt_ids,omitempty\""; Default *string "json:\"default,omitempty\""; Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Options *[]struct { Id *string "json:\"id,omitempty\""; Inputs *[]struct { Callback *string "json:\"callback,omitempty\""; DefaultRowCount *int "json:\"default-row-count,omitempty\""; HideValues *bool "json:\"hide-values,omitempty\""; Hint *string "json:\"hint,omitempty\""; Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\""; Placeholder *string "json:\"placeholder,omitempty\""; Regex *string "json:\"regex,omitempty\""; Required *bool "json:\"required,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"inputs,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"options,omitempty\""; Placeholder *string "json:\"placeholder,omitempty\""; Preferences *[]string "json:\"preferences,omitempty\""; Required *bool "json:\"required,omitempty\""; Type *string "json:\"type,omitempty\"" } "json:\"configuration,omitempty\""; ExpandSettings *bool "json:\"expand_settings,omitempty\""; Id *string "json:\"id,omitempty\""; Max *int "json:\"max,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"types,omitempty\"" } "json:\"credentials,omitempty\"" }
	if i, ok := r.(*struct {
		Credentials *[]struct {
			Category      *string "json:\"category,omitempty\""
			DefaultExpand *bool   "json:\"default_expand,omitempty\""
			Id            *string "json:\"id,omitempty\""
			Types         *[]struct {
				Configuration *[]struct {
					AltIds  *string "json:\"alt_ids,omitempty\""
					Default *string "json:\"default,omitempty\""
					Id      *string "json:\"id,omitempty\""
					Name    *string "json:\"name,omitempty\""
					Options *[]struct {
						Id     *string "json:\"id,omitempty\""
						Inputs *[]struct {
							Callback        *string "json:\"callback,omitempty\""
							DefaultRowCount *int    "json:\"default-row-count,omitempty\""
							HideValues      *bool   "json:\"hide-values,omitempty\""
							Hint            *string "json:\"hint,omitempty\""
							Id              *string "json:\"id,omitempty\""
							Name            *string "json:\"name,omitempty\""
							Placeholder     *string "json:\"placeholder,omitempty\""
							Regex           *string "json:\"regex,omitempty\""
							Required        *bool   "json:\"required,omitempty\""
							Type            *string "json:\"type,omitempty\""
						} "json:\"inputs,omitempty\""
						Name *string "json:\"name,omitempty\""
					} "json:\"options,omitempty\""
					Placeholder *string   "json:\"placeholder,omitempty\""
					Preferences *[]string "json:\"preferences,omitempty\""
					Required    *bool     "json:\"required,omitempty\""
					Type        *string   "json:\"type,omitempty\""
				} "json:\"configuration,omitempty\""
				ExpandSettings *bool   "json:\"expand_settings,omitempty\""
				Id             *string "json:\"id,omitempty\""
				Max            *int    "json:\"max,omitempty\""
				Name           *string "json:\"name,omitempty\""
			} "json:\"types,omitempty\""
		} "json:\"credentials,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// List calls the Credentials ´s function
func (c *Credentials) List(params *CredentialsListParams, reqEditors ...RequestEditorFn) (*struct {
	Credentials *[]struct {
		Category *struct {
			Id   *string "json:\"id,omitempty\""
			Name *string "json:\"name,omitempty\""
		} "json:\"category,omitempty\""
		CreatedBy *struct {
			DisplayName *string "json:\"display_name,omitempty\""
			Id          *int    "json:\"id,omitempty\""
		} "json:\"created_by,omitempty\""
		CreatedDate *string "json:\"created_date,omitempty\""
		Description *string "json:\"description,omitempty\""
		LastUsedBy  *struct {
			DisplayName *string "json:\"display_name,omitempty\""
			Id          *int    "json:\"id,omitempty\""
		} "json:\"last_used_by,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Permissions *int    "json:\"permissions,omitempty\""
		Type        *struct {
			Id   *string "json:\"id,omitempty\""
			Name *string "json:\"name,omitempty\""
		} "json:\"type,omitempty\""
		UserPermissions *int    "json:\"user_permissions,omitempty\""
		Uuid            *string "json:\"uuid,omitempty\""
	} "json:\"credentials,omitempty\""
	Pagination *struct {
		Limit  *int32 "json:\"limit,omitempty\""
		Offset *int32 "json:\"offset,omitempty\""
		Sort   *[]struct {
			Name  *string "json:\"name,omitempty\""
			Order *string "json:\"order,omitempty\""
		} "json:\"sort,omitempty\""
		Total *int32 "json:\"total,omitempty\""
	} "json:\"pagination,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CredentialsListWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Credentials *[]struct { Category *struct { Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"category,omitempty\""; CreatedBy *struct { DisplayName *string "json:\"display_name,omitempty\""; Id *int "json:\"id,omitempty\"" } "json:\"created_by,omitempty\""; CreatedDate *string "json:\"created_date,omitempty\""; Description *string "json:\"description,omitempty\""; LastUsedBy *struct { DisplayName *string "json:\"display_name,omitempty\""; Id *int "json:\"id,omitempty\"" } "json:\"last_used_by,omitempty\""; Name *string "json:\"name,omitempty\""; Permissions *int "json:\"permissions,omitempty\""; Type *struct { Id *string "json:\"id,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"type,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"credentials,omitempty\""; Pagination *struct { Limit *int32 "json:\"limit,omitempty\""; Offset *int32 "json:\"offset,omitempty\""; Sort *[]struct { Name *string "json:\"name,omitempty\""; Order *string "json:\"order,omitempty\"" } "json:\"sort,omitempty\""; Total *int32 "json:\"total,omitempty\"" } "json:\"pagination,omitempty\"" }
	if i, ok := r.(*struct {
		Credentials *[]struct {
			Category *struct {
				Id   *string "json:\"id,omitempty\""
				Name *string "json:\"name,omitempty\""
			} "json:\"category,omitempty\""
			CreatedBy *struct {
				DisplayName *string "json:\"display_name,omitempty\""
				Id          *int    "json:\"id,omitempty\""
			} "json:\"created_by,omitempty\""
			CreatedDate *string "json:\"created_date,omitempty\""
			Description *string "json:\"description,omitempty\""
			LastUsedBy  *struct {
				DisplayName *string "json:\"display_name,omitempty\""
				Id          *int    "json:\"id,omitempty\""
			} "json:\"last_used_by,omitempty\""
			Name        *string "json:\"name,omitempty\""
			Permissions *int    "json:\"permissions,omitempty\""
			Type        *struct {
				Id   *string "json:\"id,omitempty\""
				Name *string "json:\"name,omitempty\""
			} "json:\"type,omitempty\""
			UserPermissions *int    "json:\"user_permissions,omitempty\""
			Uuid            *string "json:\"uuid,omitempty\""
		} "json:\"credentials,omitempty\""
		Pagination *struct {
			Limit  *int32 "json:\"limit,omitempty\""
			Offset *int32 "json:\"offset,omitempty\""
			Sort   *[]struct {
				Name  *string "json:\"name,omitempty\""
				Order *string "json:\"order,omitempty\""
			} "json:\"sort,omitempty\""
			Total *int32 "json:\"total,omitempty\""
		} "json:\"pagination,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UpdateWithBody calls the Credentials ´s function
func (c *Credentials) UpdateWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	Updated *bool "json:\"updated,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CredentialsUpdateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Updated *bool "json:\"updated,omitempty\"" }
	if i, ok := r.(*struct {
		Updated *bool "json:\"updated,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Update calls the Credentials ´s function
func (c *Credentials) Update(arg1 string, arg2 CredentialsUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	Updated *bool "json:\"updated,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.CredentialsUpdateWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Updated *bool "json:\"updated,omitempty\"" }
	if i, ok := r.(*struct {
		Updated *bool "json:\"updated,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// IoExportsCompliance implementation of the IoExportsCompliance interface
type IoExportsCompliance struct {
	*ClientWithResponses
}

// Cancel calls the IoExportsCompliance ´s function
func (c *IoExportsCompliance) Cancel(contentType string, reqEditors ...RequestEditorFn) (*IoExportsComplianceCancelResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoExportsComplianceCancelWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *IoExportsComplianceCancelResponse
	if i, ok := r.(*IoExportsComplianceCancelResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CreateWithBody calls the IoExportsCompliance ´s function
func (c *IoExportsCompliance) CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	ExportUuid *string "json:\"export_uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoExportsComplianceCreateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ExportUuid *string "json:\"export_uuid,omitempty\"" }
	if i, ok := r.(*struct {
		ExportUuid *string "json:\"export_uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Create calls the IoExportsCompliance ´s function
func (c *IoExportsCompliance) Create(arg1 IoExportsComplianceCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	ExportUuid *string "json:\"export_uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoExportsComplianceCreateWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ExportUuid *string "json:\"export_uuid,omitempty\"" }
	if i, ok := r.(*struct {
		ExportUuid *string "json:\"export_uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Download calls the IoExportsCompliance ´s function
func (c *IoExportsCompliance) Download(arg1 string, arg2 int32, reqEditors ...RequestEditorFn) (*[]struct {
	ActualValue   *string "json:\"actual_value,omitempty\""
	AssetUuid     *string "json:\"asset_uuid,omitempty\""
	AuditFile     *string "json:\"audit_file,omitempty\""
	CheckError    *string "json:\"check_error,omitempty\""
	CheckId       *string "json:\"check_id,omitempty\""
	CheckInfo     *string "json:\"check_info,omitempty\""
	CheckName     *string "json:\"check_name,omitempty\""
	DbType        *string "json:\"db_type,omitempty\""
	ExpectedValue *string "json:\"expected_value,omitempty\""
	FirstSeen     *int64  "json:\"first_seen,omitempty\""
	LastSeen      *int64  "json:\"last_seen,omitempty\""
	PluginId      *int    "json:\"plugin_id,omitempty\""
	ProfileName   *string "json:\"profile_name,omitempty\""
	Reference     *[]struct {
		Control   *string "json:\"control,omitempty\""
		Framework *string "json:\"framework,omitempty\""
	} "json:\"reference,omitempty\""
	SeeAlso  *string "json:\"see_also,omitempty\""
	Solution *string "json:\"solution,omitempty\""
	Status   *string "json:\"status,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoExportsComplianceDownloadWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { ActualValue *string "json:\"actual_value,omitempty\""; AssetUuid *string "json:\"asset_uuid,omitempty\""; AuditFile *string "json:\"audit_file,omitempty\""; CheckError *string "json:\"check_error,omitempty\""; CheckId *string "json:\"check_id,omitempty\""; CheckInfo *string "json:\"check_info,omitempty\""; CheckName *string "json:\"check_name,omitempty\""; DbType *string "json:\"db_type,omitempty\""; ExpectedValue *string "json:\"expected_value,omitempty\""; FirstSeen *int64 "json:\"first_seen,omitempty\""; LastSeen *int64 "json:\"last_seen,omitempty\""; PluginId *int "json:\"plugin_id,omitempty\""; ProfileName *string "json:\"profile_name,omitempty\""; Reference *[]struct { Control *string "json:\"control,omitempty\""; Framework *string "json:\"framework,omitempty\"" } "json:\"reference,omitempty\""; SeeAlso *string "json:\"see_also,omitempty\""; Solution *string "json:\"solution,omitempty\""; Status *string "json:\"status,omitempty\"" }
	if i, ok := r.(*[]struct {
		ActualValue   *string "json:\"actual_value,omitempty\""
		AssetUuid     *string "json:\"asset_uuid,omitempty\""
		AuditFile     *string "json:\"audit_file,omitempty\""
		CheckError    *string "json:\"check_error,omitempty\""
		CheckId       *string "json:\"check_id,omitempty\""
		CheckInfo     *string "json:\"check_info,omitempty\""
		CheckName     *string "json:\"check_name,omitempty\""
		DbType        *string "json:\"db_type,omitempty\""
		ExpectedValue *string "json:\"expected_value,omitempty\""
		FirstSeen     *int64  "json:\"first_seen,omitempty\""
		LastSeen      *int64  "json:\"last_seen,omitempty\""
		PluginId      *int    "json:\"plugin_id,omitempty\""
		ProfileName   *string "json:\"profile_name,omitempty\""
		Reference     *[]struct {
			Control   *string "json:\"control,omitempty\""
			Framework *string "json:\"framework,omitempty\""
		} "json:\"reference,omitempty\""
		SeeAlso  *string "json:\"see_also,omitempty\""
		Solution *string "json:\"solution,omitempty\""
		Status   *string "json:\"status,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Status calls the IoExportsCompliance ´s function
func (c *IoExportsCompliance) Status(contentType string, reqEditors ...RequestEditorFn) (*struct {
	ChunksAvailable *[]int32 "json:\"chunks_available,omitempty\""
	Status          *string  "json:\"status,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoExportsComplianceStatusWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { ChunksAvailable *[]int32 "json:\"chunks_available,omitempty\""; Status *string "json:\"status,omitempty\"" }
	if i, ok := r.(*struct {
		ChunksAvailable *[]int32 "json:\"chunks_available,omitempty\""
		Status          *string  "json:\"status,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// IoNetworks implementation of the IoNetworks interface
type IoNetworks struct {
	*ClientWithResponses
}

// AssetCountDetails calls the IoNetworks ´s function
func (c *IoNetworks) AssetCountDetails(arg1 string, arg2 int, reqEditors ...RequestEditorFn) (*struct {
	NumAssetsNotSeen *int "json:\"numAssetsNotSeen,omitempty\""
	NumAssetstotal   *int "json:\"numAssetstotal,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.IoNetworksAssetCountDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { NumAssetsNotSeen *int "json:\"numAssetsNotSeen,omitempty\""; NumAssetstotal *int "json:\"numAssetstotal,omitempty\"" }
	if i, ok := r.(*struct {
		NumAssetsNotSeen *int "json:\"numAssetsNotSeen,omitempty\""
		NumAssetstotal   *int "json:\"numAssetstotal,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Networks implementation of the Networks interface
type Networks struct {
	*ClientWithResponses
}

// AssignScannerBulkWithBody calls the Networks ´s function
func (c *Networks) AssignScannerBulkWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.NetworksAssignScannerBulkWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AssignScannerBulk calls the Networks ´s function
func (c *Networks) AssignScannerBulk(arg1 string, arg2 NetworksAssignScannerBulkJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.NetworksAssignScannerBulkWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AssignScanner calls the Networks ´s function
func (c *Networks) AssignScanner(arg1 string, contentType string, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.NetworksAssignScannerWithResponse(c.ClientInterface.(*Client).ctx, arg1, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CreateWithBody calls the Networks ´s function
func (c *Networks) CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	AssetsTtlDays     *int    "json:\"assets_ttl_days,omitempty\""
	Created           *int    "json:\"created,omitempty\""
	CreatedBy         *string "json:\"created_by,omitempty\""
	CreatedInSeconds  *int    "json:\"created_in_seconds,omitempty\""
	Deleted           *int    "json:\"deleted,omitempty\""
	DeletedBy         *string "json:\"deleted_by,omitempty\""
	Description       *string "json:\"description,omitempty\""
	IsDefault         *bool   "json:\"is_default,omitempty\""
	Modified          *int    "json:\"modified,omitempty\""
	ModifiedBy        *string "json:\"modified_by,omitempty\""
	ModifiedInSeconds *int    "json:\"modified_in_seconds,omitempty\""
	Name              *string "json:\"name,omitempty\""
	OwnerUuid         *string "json:\"owner_uuid,omitempty\""
	Uuid              *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.NetworksCreateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AssetsTtlDays *int "json:\"assets_ttl_days,omitempty\""; Created *int "json:\"created,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; CreatedInSeconds *int "json:\"created_in_seconds,omitempty\""; Deleted *int "json:\"deleted,omitempty\""; DeletedBy *string "json:\"deleted_by,omitempty\""; Description *string "json:\"description,omitempty\""; IsDefault *bool "json:\"is_default,omitempty\""; Modified *int "json:\"modified,omitempty\""; ModifiedBy *string "json:\"modified_by,omitempty\""; ModifiedInSeconds *int "json:\"modified_in_seconds,omitempty\""; Name *string "json:\"name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		AssetsTtlDays     *int    "json:\"assets_ttl_days,omitempty\""
		Created           *int    "json:\"created,omitempty\""
		CreatedBy         *string "json:\"created_by,omitempty\""
		CreatedInSeconds  *int    "json:\"created_in_seconds,omitempty\""
		Deleted           *int    "json:\"deleted,omitempty\""
		DeletedBy         *string "json:\"deleted_by,omitempty\""
		Description       *string "json:\"description,omitempty\""
		IsDefault         *bool   "json:\"is_default,omitempty\""
		Modified          *int    "json:\"modified,omitempty\""
		ModifiedBy        *string "json:\"modified_by,omitempty\""
		ModifiedInSeconds *int    "json:\"modified_in_seconds,omitempty\""
		Name              *string "json:\"name,omitempty\""
		OwnerUuid         *string "json:\"owner_uuid,omitempty\""
		Uuid              *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Create calls the Networks ´s function
func (c *Networks) Create(arg1 NetworksCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	AssetsTtlDays     *int    "json:\"assets_ttl_days,omitempty\""
	Created           *int    "json:\"created,omitempty\""
	CreatedBy         *string "json:\"created_by,omitempty\""
	CreatedInSeconds  *int    "json:\"created_in_seconds,omitempty\""
	Deleted           *int    "json:\"deleted,omitempty\""
	DeletedBy         *string "json:\"deleted_by,omitempty\""
	Description       *string "json:\"description,omitempty\""
	IsDefault         *bool   "json:\"is_default,omitempty\""
	Modified          *int    "json:\"modified,omitempty\""
	ModifiedBy        *string "json:\"modified_by,omitempty\""
	ModifiedInSeconds *int    "json:\"modified_in_seconds,omitempty\""
	Name              *string "json:\"name,omitempty\""
	OwnerUuid         *string "json:\"owner_uuid,omitempty\""
	Uuid              *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.NetworksCreateWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AssetsTtlDays *int "json:\"assets_ttl_days,omitempty\""; Created *int "json:\"created,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; CreatedInSeconds *int "json:\"created_in_seconds,omitempty\""; Deleted *int "json:\"deleted,omitempty\""; DeletedBy *string "json:\"deleted_by,omitempty\""; Description *string "json:\"description,omitempty\""; IsDefault *bool "json:\"is_default,omitempty\""; Modified *int "json:\"modified,omitempty\""; ModifiedBy *string "json:\"modified_by,omitempty\""; ModifiedInSeconds *int "json:\"modified_in_seconds,omitempty\""; Name *string "json:\"name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		AssetsTtlDays     *int    "json:\"assets_ttl_days,omitempty\""
		Created           *int    "json:\"created,omitempty\""
		CreatedBy         *string "json:\"created_by,omitempty\""
		CreatedInSeconds  *int    "json:\"created_in_seconds,omitempty\""
		Deleted           *int    "json:\"deleted,omitempty\""
		DeletedBy         *string "json:\"deleted_by,omitempty\""
		Description       *string "json:\"description,omitempty\""
		IsDefault         *bool   "json:\"is_default,omitempty\""
		Modified          *int    "json:\"modified,omitempty\""
		ModifiedBy        *string "json:\"modified_by,omitempty\""
		ModifiedInSeconds *int    "json:\"modified_in_seconds,omitempty\""
		Name              *string "json:\"name,omitempty\""
		OwnerUuid         *string "json:\"owner_uuid,omitempty\""
		Uuid              *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Delete calls the Networks ´s function
func (c *Networks) Delete(contentType string, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.NetworksDeleteWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Details calls the Networks ´s function
func (c *Networks) Details(contentType string, reqEditors ...RequestEditorFn) (*struct {
	AssetsTtlDays     *int    "json:\"assets_ttl_days,omitempty\""
	Created           *int    "json:\"created,omitempty\""
	CreatedBy         *string "json:\"created_by,omitempty\""
	CreatedInSeconds  *int    "json:\"created_in_seconds,omitempty\""
	Deleted           *int    "json:\"deleted,omitempty\""
	DeletedBy         *string "json:\"deleted_by,omitempty\""
	Description       *string "json:\"description,omitempty\""
	IsDefault         *bool   "json:\"is_default,omitempty\""
	Modified          *int    "json:\"modified,omitempty\""
	ModifiedBy        *string "json:\"modified_by,omitempty\""
	ModifiedInSeconds *int    "json:\"modified_in_seconds,omitempty\""
	Name              *string "json:\"name,omitempty\""
	OwnerUuid         *string "json:\"owner_uuid,omitempty\""
	Uuid              *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.NetworksDetailsWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AssetsTtlDays *int "json:\"assets_ttl_days,omitempty\""; Created *int "json:\"created,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; CreatedInSeconds *int "json:\"created_in_seconds,omitempty\""; Deleted *int "json:\"deleted,omitempty\""; DeletedBy *string "json:\"deleted_by,omitempty\""; Description *string "json:\"description,omitempty\""; IsDefault *bool "json:\"is_default,omitempty\""; Modified *int "json:\"modified,omitempty\""; ModifiedBy *string "json:\"modified_by,omitempty\""; ModifiedInSeconds *int "json:\"modified_in_seconds,omitempty\""; Name *string "json:\"name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		AssetsTtlDays     *int    "json:\"assets_ttl_days,omitempty\""
		Created           *int    "json:\"created,omitempty\""
		CreatedBy         *string "json:\"created_by,omitempty\""
		CreatedInSeconds  *int    "json:\"created_in_seconds,omitempty\""
		Deleted           *int    "json:\"deleted,omitempty\""
		DeletedBy         *string "json:\"deleted_by,omitempty\""
		Description       *string "json:\"description,omitempty\""
		IsDefault         *bool   "json:\"is_default,omitempty\""
		Modified          *int    "json:\"modified,omitempty\""
		ModifiedBy        *string "json:\"modified_by,omitempty\""
		ModifiedInSeconds *int    "json:\"modified_in_seconds,omitempty\""
		Name              *string "json:\"name,omitempty\""
		OwnerUuid         *string "json:\"owner_uuid,omitempty\""
		Uuid              *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ListAssignableScanners calls the Networks ´s function
func (c *Networks) ListAssignableScanners(contentType string, reqEditors ...RequestEditorFn) (*struct {
	Scanners *[]struct {
		CreationDate         *int32                  "json:\"creation_date,omitempty\""
		Distro               *string                 "json:\"distro,omitempty\""
		EngineBuild          *string                 "json:\"engine_build,omitempty\""
		EngineVersion        *string                 "json:\"engine_version,omitempty\""
		Group                *bool                   "json:\"group,omitempty\""
		Id                   *int32                  "json:\"id,omitempty\""
		Key                  *string                 "json:\"key,omitempty\""
		LastConnect          *int32                  "json:\"last_connect,omitempty\""
		LastModificationDate *int                    "json:\"last_modification_date,omitempty\""
		Linked               *int                    "json:\"linked,omitempty\""
		LoadedPluginSet      *string                 "json:\"loaded_plugin_set,omitempty\""
		Name                 *string                 "json:\"name,omitempty\""
		NumHosts             *int                    "json:\"num_hosts,omitempty\""
		NumScans             *int                    "json:\"num_scans,omitempty\""
		NumSessions          *int                    "json:\"num_sessions,omitempty\""
		NumTcpSessions       *int                    "json:\"num_tcp_sessions,omitempty\""
		Owner                *string                 "json:\"owner,omitempty\""
		OwnerId              *int                    "json:\"owner_id,omitempty\""
		OwnerName            *string                 "json:\"owner_name,omitempty\""
		OwnerUuid            *string                 "json:\"owner_uuid,omitempty\""
		Platform             *string                 "json:\"platform,omitempty\""
		Pool                 *bool                   "json:\"pool,omitempty\""
		RemoteUuid           *string                 "json:\"remote_uuid,omitempty\""
		ReportFrequency      *int                    "json:\"report_frequency,omitempty\""
		ScanCount            *int                    "json:\"scan_count,omitempty\""
		Settings             *map[string]interface{} "json:\"settings,omitempty\""
		Source               *string                 "json:\"source,omitempty\""
		Status               *string                 "json:\"status,omitempty\""
		SupportsRemoteLogs   *bool                   "json:\"supports_remote_logs,omitempty\""
		Timestamp            *int                    "json:\"timestamp,omitempty\""
		Type                 *string                 "json:\"type,omitempty\""
		Uuid                 *string                 "json:\"uuid,omitempty\""
	} "json:\"scanners,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.NetworksListAssignableScannersWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Scanners *[]struct { CreationDate *int32 "json:\"creation_date,omitempty\""; Distro *string "json:\"distro,omitempty\""; EngineBuild *string "json:\"engine_build,omitempty\""; EngineVersion *string "json:\"engine_version,omitempty\""; Group *bool "json:\"group,omitempty\""; Id *int32 "json:\"id,omitempty\""; Key *string "json:\"key,omitempty\""; LastConnect *int32 "json:\"last_connect,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Linked *int "json:\"linked,omitempty\""; LoadedPluginSet *string "json:\"loaded_plugin_set,omitempty\""; Name *string "json:\"name,omitempty\""; NumHosts *int "json:\"num_hosts,omitempty\""; NumScans *int "json:\"num_scans,omitempty\""; NumSessions *int "json:\"num_sessions,omitempty\""; NumTcpSessions *int "json:\"num_tcp_sessions,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; OwnerName *string "json:\"owner_name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Platform *string "json:\"platform,omitempty\""; Pool *bool "json:\"pool,omitempty\""; RemoteUuid *string "json:\"remote_uuid,omitempty\""; ReportFrequency *int "json:\"report_frequency,omitempty\""; ScanCount *int "json:\"scan_count,omitempty\""; Settings *map[string]interface {} "json:\"settings,omitempty\""; Source *string "json:\"source,omitempty\""; Status *string "json:\"status,omitempty\""; SupportsRemoteLogs *bool "json:\"supports_remote_logs,omitempty\""; Timestamp *int "json:\"timestamp,omitempty\""; Type *string "json:\"type,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"scanners,omitempty\"" }
	if i, ok := r.(*struct {
		Scanners *[]struct {
			CreationDate         *int32                  "json:\"creation_date,omitempty\""
			Distro               *string                 "json:\"distro,omitempty\""
			EngineBuild          *string                 "json:\"engine_build,omitempty\""
			EngineVersion        *string                 "json:\"engine_version,omitempty\""
			Group                *bool                   "json:\"group,omitempty\""
			Id                   *int32                  "json:\"id,omitempty\""
			Key                  *string                 "json:\"key,omitempty\""
			LastConnect          *int32                  "json:\"last_connect,omitempty\""
			LastModificationDate *int                    "json:\"last_modification_date,omitempty\""
			Linked               *int                    "json:\"linked,omitempty\""
			LoadedPluginSet      *string                 "json:\"loaded_plugin_set,omitempty\""
			Name                 *string                 "json:\"name,omitempty\""
			NumHosts             *int                    "json:\"num_hosts,omitempty\""
			NumScans             *int                    "json:\"num_scans,omitempty\""
			NumSessions          *int                    "json:\"num_sessions,omitempty\""
			NumTcpSessions       *int                    "json:\"num_tcp_sessions,omitempty\""
			Owner                *string                 "json:\"owner,omitempty\""
			OwnerId              *int                    "json:\"owner_id,omitempty\""
			OwnerName            *string                 "json:\"owner_name,omitempty\""
			OwnerUuid            *string                 "json:\"owner_uuid,omitempty\""
			Platform             *string                 "json:\"platform,omitempty\""
			Pool                 *bool                   "json:\"pool,omitempty\""
			RemoteUuid           *string                 "json:\"remote_uuid,omitempty\""
			ReportFrequency      *int                    "json:\"report_frequency,omitempty\""
			ScanCount            *int                    "json:\"scan_count,omitempty\""
			Settings             *map[string]interface{} "json:\"settings,omitempty\""
			Source               *string                 "json:\"source,omitempty\""
			Status               *string                 "json:\"status,omitempty\""
			SupportsRemoteLogs   *bool                   "json:\"supports_remote_logs,omitempty\""
			Timestamp            *int                    "json:\"timestamp,omitempty\""
			Type                 *string                 "json:\"type,omitempty\""
			Uuid                 *string                 "json:\"uuid,omitempty\""
		} "json:\"scanners,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ListScanners calls the Networks ´s function
func (c *Networks) ListScanners(contentType string, reqEditors ...RequestEditorFn) (*struct {
	Scanners *[]struct {
		CreationDate         *int32                  "json:\"creation_date,omitempty\""
		Distro               *string                 "json:\"distro,omitempty\""
		EngineBuild          *string                 "json:\"engine_build,omitempty\""
		EngineVersion        *string                 "json:\"engine_version,omitempty\""
		Group                *bool                   "json:\"group,omitempty\""
		Id                   *int32                  "json:\"id,omitempty\""
		Key                  *string                 "json:\"key,omitempty\""
		LastConnect          *int32                  "json:\"last_connect,omitempty\""
		LastModificationDate *int                    "json:\"last_modification_date,omitempty\""
		Linked               *int                    "json:\"linked,omitempty\""
		LoadedPluginSet      *string                 "json:\"loaded_plugin_set,omitempty\""
		Name                 *string                 "json:\"name,omitempty\""
		NumHosts             *int                    "json:\"num_hosts,omitempty\""
		NumScans             *int                    "json:\"num_scans,omitempty\""
		NumSessions          *int                    "json:\"num_sessions,omitempty\""
		NumTcpSessions       *int                    "json:\"num_tcp_sessions,omitempty\""
		Owner                *string                 "json:\"owner,omitempty\""
		OwnerId              *int                    "json:\"owner_id,omitempty\""
		OwnerName            *string                 "json:\"owner_name,omitempty\""
		OwnerUuid            *string                 "json:\"owner_uuid,omitempty\""
		Platform             *string                 "json:\"platform,omitempty\""
		Pool                 *bool                   "json:\"pool,omitempty\""
		RemoteUuid           *string                 "json:\"remote_uuid,omitempty\""
		ReportFrequency      *int                    "json:\"report_frequency,omitempty\""
		ScanCount            *int                    "json:\"scan_count,omitempty\""
		Settings             *map[string]interface{} "json:\"settings,omitempty\""
		Source               *string                 "json:\"source,omitempty\""
		Status               *string                 "json:\"status,omitempty\""
		SupportsRemoteLogs   *bool                   "json:\"supports_remote_logs,omitempty\""
		Timestamp            *int                    "json:\"timestamp,omitempty\""
		Type                 *string                 "json:\"type,omitempty\""
		Uuid                 *string                 "json:\"uuid,omitempty\""
	} "json:\"scanners,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.NetworksListScannersWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Scanners *[]struct { CreationDate *int32 "json:\"creation_date,omitempty\""; Distro *string "json:\"distro,omitempty\""; EngineBuild *string "json:\"engine_build,omitempty\""; EngineVersion *string "json:\"engine_version,omitempty\""; Group *bool "json:\"group,omitempty\""; Id *int32 "json:\"id,omitempty\""; Key *string "json:\"key,omitempty\""; LastConnect *int32 "json:\"last_connect,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Linked *int "json:\"linked,omitempty\""; LoadedPluginSet *string "json:\"loaded_plugin_set,omitempty\""; Name *string "json:\"name,omitempty\""; NumHosts *int "json:\"num_hosts,omitempty\""; NumScans *int "json:\"num_scans,omitempty\""; NumSessions *int "json:\"num_sessions,omitempty\""; NumTcpSessions *int "json:\"num_tcp_sessions,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; OwnerName *string "json:\"owner_name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Platform *string "json:\"platform,omitempty\""; Pool *bool "json:\"pool,omitempty\""; RemoteUuid *string "json:\"remote_uuid,omitempty\""; ReportFrequency *int "json:\"report_frequency,omitempty\""; ScanCount *int "json:\"scan_count,omitempty\""; Settings *map[string]interface {} "json:\"settings,omitempty\""; Source *string "json:\"source,omitempty\""; Status *string "json:\"status,omitempty\""; SupportsRemoteLogs *bool "json:\"supports_remote_logs,omitempty\""; Timestamp *int "json:\"timestamp,omitempty\""; Type *string "json:\"type,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"scanners,omitempty\"" }
	if i, ok := r.(*struct {
		Scanners *[]struct {
			CreationDate         *int32                  "json:\"creation_date,omitempty\""
			Distro               *string                 "json:\"distro,omitempty\""
			EngineBuild          *string                 "json:\"engine_build,omitempty\""
			EngineVersion        *string                 "json:\"engine_version,omitempty\""
			Group                *bool                   "json:\"group,omitempty\""
			Id                   *int32                  "json:\"id,omitempty\""
			Key                  *string                 "json:\"key,omitempty\""
			LastConnect          *int32                  "json:\"last_connect,omitempty\""
			LastModificationDate *int                    "json:\"last_modification_date,omitempty\""
			Linked               *int                    "json:\"linked,omitempty\""
			LoadedPluginSet      *string                 "json:\"loaded_plugin_set,omitempty\""
			Name                 *string                 "json:\"name,omitempty\""
			NumHosts             *int                    "json:\"num_hosts,omitempty\""
			NumScans             *int                    "json:\"num_scans,omitempty\""
			NumSessions          *int                    "json:\"num_sessions,omitempty\""
			NumTcpSessions       *int                    "json:\"num_tcp_sessions,omitempty\""
			Owner                *string                 "json:\"owner,omitempty\""
			OwnerId              *int                    "json:\"owner_id,omitempty\""
			OwnerName            *string                 "json:\"owner_name,omitempty\""
			OwnerUuid            *string                 "json:\"owner_uuid,omitempty\""
			Platform             *string                 "json:\"platform,omitempty\""
			Pool                 *bool                   "json:\"pool,omitempty\""
			RemoteUuid           *string                 "json:\"remote_uuid,omitempty\""
			ReportFrequency      *int                    "json:\"report_frequency,omitempty\""
			ScanCount            *int                    "json:\"scan_count,omitempty\""
			Settings             *map[string]interface{} "json:\"settings,omitempty\""
			Source               *string                 "json:\"source,omitempty\""
			Status               *string                 "json:\"status,omitempty\""
			SupportsRemoteLogs   *bool                   "json:\"supports_remote_logs,omitempty\""
			Timestamp            *int                    "json:\"timestamp,omitempty\""
			Type                 *string                 "json:\"type,omitempty\""
			Uuid                 *string                 "json:\"uuid,omitempty\""
		} "json:\"scanners,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// List calls the Networks ´s function
func (c *Networks) List(params *NetworksListParams, reqEditors ...RequestEditorFn) (*struct {
	Networks *[]struct {
		AssetsTtlDays     *int    "json:\"assets_ttl_days,omitempty\""
		Created           *int    "json:\"created,omitempty\""
		CreatedBy         *string "json:\"created_by,omitempty\""
		CreatedInSeconds  *int    "json:\"created_in_seconds,omitempty\""
		Deleted           *int    "json:\"deleted,omitempty\""
		DeletedBy         *string "json:\"deleted_by,omitempty\""
		Description       *string "json:\"description,omitempty\""
		IsDefault         *bool   "json:\"is_default,omitempty\""
		Modified          *int    "json:\"modified,omitempty\""
		ModifiedBy        *string "json:\"modified_by,omitempty\""
		ModifiedInSeconds *int    "json:\"modified_in_seconds,omitempty\""
		Name              *string "json:\"name,omitempty\""
		OwnerUuid         *string "json:\"owner_uuid,omitempty\""
		Uuid              *string "json:\"uuid,omitempty\""
	} "json:\"networks,omitempty\""
	Pagination *[]struct {
		Limit  *int32 "json:\"limit,omitempty\""
		Offset *int32 "json:\"offset,omitempty\""
		Sort   *[]struct {
			Name  *string "json:\"name,omitempty\""
			Order *string "json:\"order,omitempty\""
		} "json:\"sort,omitempty\""
		Total *int32 "json:\"total,omitempty\""
	} "json:\"pagination,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.NetworksListWithResponse(c.ClientInterface.(*Client).ctx, params, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Networks *[]struct { AssetsTtlDays *int "json:\"assets_ttl_days,omitempty\""; Created *int "json:\"created,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; CreatedInSeconds *int "json:\"created_in_seconds,omitempty\""; Deleted *int "json:\"deleted,omitempty\""; DeletedBy *string "json:\"deleted_by,omitempty\""; Description *string "json:\"description,omitempty\""; IsDefault *bool "json:\"is_default,omitempty\""; Modified *int "json:\"modified,omitempty\""; ModifiedBy *string "json:\"modified_by,omitempty\""; ModifiedInSeconds *int "json:\"modified_in_seconds,omitempty\""; Name *string "json:\"name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" } "json:\"networks,omitempty\""; Pagination *[]struct { Limit *int32 "json:\"limit,omitempty\""; Offset *int32 "json:\"offset,omitempty\""; Sort *[]struct { Name *string "json:\"name,omitempty\""; Order *string "json:\"order,omitempty\"" } "json:\"sort,omitempty\""; Total *int32 "json:\"total,omitempty\"" } "json:\"pagination,omitempty\"" }
	if i, ok := r.(*struct {
		Networks *[]struct {
			AssetsTtlDays     *int    "json:\"assets_ttl_days,omitempty\""
			Created           *int    "json:\"created,omitempty\""
			CreatedBy         *string "json:\"created_by,omitempty\""
			CreatedInSeconds  *int    "json:\"created_in_seconds,omitempty\""
			Deleted           *int    "json:\"deleted,omitempty\""
			DeletedBy         *string "json:\"deleted_by,omitempty\""
			Description       *string "json:\"description,omitempty\""
			IsDefault         *bool   "json:\"is_default,omitempty\""
			Modified          *int    "json:\"modified,omitempty\""
			ModifiedBy        *string "json:\"modified_by,omitempty\""
			ModifiedInSeconds *int    "json:\"modified_in_seconds,omitempty\""
			Name              *string "json:\"name,omitempty\""
			OwnerUuid         *string "json:\"owner_uuid,omitempty\""
			Uuid              *string "json:\"uuid,omitempty\""
		} "json:\"networks,omitempty\""
		Pagination *[]struct {
			Limit  *int32 "json:\"limit,omitempty\""
			Offset *int32 "json:\"offset,omitempty\""
			Sort   *[]struct {
				Name  *string "json:\"name,omitempty\""
				Order *string "json:\"order,omitempty\""
			} "json:\"sort,omitempty\""
			Total *int32 "json:\"total,omitempty\""
		} "json:\"pagination,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// UpdateWithBody calls the Networks ´s function
func (c *Networks) UpdateWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	AssetsTtlDays     *int    "json:\"assets_ttl_days,omitempty\""
	Created           *int    "json:\"created,omitempty\""
	CreatedBy         *string "json:\"created_by,omitempty\""
	CreatedInSeconds  *int    "json:\"created_in_seconds,omitempty\""
	Deleted           *int    "json:\"deleted,omitempty\""
	DeletedBy         *string "json:\"deleted_by,omitempty\""
	Description       *string "json:\"description,omitempty\""
	IsDefault         *bool   "json:\"is_default,omitempty\""
	Modified          *int    "json:\"modified,omitempty\""
	ModifiedBy        *string "json:\"modified_by,omitempty\""
	ModifiedInSeconds *int    "json:\"modified_in_seconds,omitempty\""
	Name              *string "json:\"name,omitempty\""
	OwnerUuid         *string "json:\"owner_uuid,omitempty\""
	Uuid              *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.NetworksUpdateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AssetsTtlDays *int "json:\"assets_ttl_days,omitempty\""; Created *int "json:\"created,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; CreatedInSeconds *int "json:\"created_in_seconds,omitempty\""; Deleted *int "json:\"deleted,omitempty\""; DeletedBy *string "json:\"deleted_by,omitempty\""; Description *string "json:\"description,omitempty\""; IsDefault *bool "json:\"is_default,omitempty\""; Modified *int "json:\"modified,omitempty\""; ModifiedBy *string "json:\"modified_by,omitempty\""; ModifiedInSeconds *int "json:\"modified_in_seconds,omitempty\""; Name *string "json:\"name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		AssetsTtlDays     *int    "json:\"assets_ttl_days,omitempty\""
		Created           *int    "json:\"created,omitempty\""
		CreatedBy         *string "json:\"created_by,omitempty\""
		CreatedInSeconds  *int    "json:\"created_in_seconds,omitempty\""
		Deleted           *int    "json:\"deleted,omitempty\""
		DeletedBy         *string "json:\"deleted_by,omitempty\""
		Description       *string "json:\"description,omitempty\""
		IsDefault         *bool   "json:\"is_default,omitempty\""
		Modified          *int    "json:\"modified,omitempty\""
		ModifiedBy        *string "json:\"modified_by,omitempty\""
		ModifiedInSeconds *int    "json:\"modified_in_seconds,omitempty\""
		Name              *string "json:\"name,omitempty\""
		OwnerUuid         *string "json:\"owner_uuid,omitempty\""
		Uuid              *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Update calls the Networks ´s function
func (c *Networks) Update(arg1 string, arg2 NetworksUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	AssetsTtlDays     *int    "json:\"assets_ttl_days,omitempty\""
	Created           *int    "json:\"created,omitempty\""
	CreatedBy         *string "json:\"created_by,omitempty\""
	CreatedInSeconds  *int    "json:\"created_in_seconds,omitempty\""
	Deleted           *int    "json:\"deleted,omitempty\""
	DeletedBy         *string "json:\"deleted_by,omitempty\""
	Description       *string "json:\"description,omitempty\""
	IsDefault         *bool   "json:\"is_default,omitempty\""
	Modified          *int    "json:\"modified,omitempty\""
	ModifiedBy        *string "json:\"modified_by,omitempty\""
	ModifiedInSeconds *int    "json:\"modified_in_seconds,omitempty\""
	Name              *string "json:\"name,omitempty\""
	OwnerUuid         *string "json:\"owner_uuid,omitempty\""
	Uuid              *string "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.NetworksUpdateWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AssetsTtlDays *int "json:\"assets_ttl_days,omitempty\""; Created *int "json:\"created,omitempty\""; CreatedBy *string "json:\"created_by,omitempty\""; CreatedInSeconds *int "json:\"created_in_seconds,omitempty\""; Deleted *int "json:\"deleted,omitempty\""; DeletedBy *string "json:\"deleted_by,omitempty\""; Description *string "json:\"description,omitempty\""; IsDefault *bool "json:\"is_default,omitempty\""; Modified *int "json:\"modified,omitempty\""; ModifiedBy *string "json:\"modified_by,omitempty\""; ModifiedInSeconds *int "json:\"modified_in_seconds,omitempty\""; Name *string "json:\"name,omitempty\""; OwnerUuid *string "json:\"owner_uuid,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		AssetsTtlDays     *int    "json:\"assets_ttl_days,omitempty\""
		Created           *int    "json:\"created,omitempty\""
		CreatedBy         *string "json:\"created_by,omitempty\""
		CreatedInSeconds  *int    "json:\"created_in_seconds,omitempty\""
		Deleted           *int    "json:\"deleted,omitempty\""
		DeletedBy         *string "json:\"deleted_by,omitempty\""
		Description       *string "json:\"description,omitempty\""
		IsDefault         *bool   "json:\"is_default,omitempty\""
		Modified          *int    "json:\"modified,omitempty\""
		ModifiedBy        *string "json:\"modified_by,omitempty\""
		ModifiedInSeconds *int    "json:\"modified_in_seconds,omitempty\""
		Name              *string "json:\"name,omitempty\""
		OwnerUuid         *string "json:\"owner_uuid,omitempty\""
		Uuid              *string "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// AgentExclusions implementation of the AgentExclusions interface
type AgentExclusions struct {
	*ClientWithResponses
}

// CreateWithBody calls the AgentExclusions ´s function
func (c *AgentExclusions) CreateWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	Schedule             *struct {
		Enabled *bool   "json:\"enabled,omitempty\""
		Endtime *string "json:\"endtime,omitempty\""
		Rrules  struct {
			Bymonthday *int                   "json:\"bymonthday,omitempty\""
			Byweekday  *string                "json:\"byweekday,omitempty\""
			Freq       N200ScheduleRrulesFreq "json:\"freq\""
			Interval   *int                   "json:\"interval,omitempty\""
		} "json:\"rrules\""
		Starttime string "json:\"starttime\""
		Timezone  string "json:\"timezone\""
	} "json:\"schedule,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentExclusionsCreateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreationDate *int "json:\"creation_date,omitempty\""; Description *string "json:\"description,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; Schedule *struct { Enabled *bool "json:\"enabled,omitempty\""; Endtime *string "json:\"endtime,omitempty\""; Rrules struct { Bymonthday *int "json:\"bymonthday,omitempty\""; Byweekday *string "json:\"byweekday,omitempty\""; Freq N200ScheduleRrulesFreq "json:\"freq\""; Interval *int "json:\"interval,omitempty\"" } "json:\"rrules\""; Starttime string "json:\"starttime\""; Timezone string "json:\"timezone\"" } "json:\"schedule,omitempty\"" }
	if i, ok := r.(*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Schedule             *struct {
			Enabled *bool   "json:\"enabled,omitempty\""
			Endtime *string "json:\"endtime,omitempty\""
			Rrules  struct {
				Bymonthday *int                   "json:\"bymonthday,omitempty\""
				Byweekday  *string                "json:\"byweekday,omitempty\""
				Freq       N200ScheduleRrulesFreq "json:\"freq\""
				Interval   *int                   "json:\"interval,omitempty\""
			} "json:\"rrules\""
			Starttime string "json:\"starttime\""
			Timezone  string "json:\"timezone\""
		} "json:\"schedule,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Create calls the AgentExclusions ´s function
func (c *AgentExclusions) Create(arg1 int32, arg2 AgentExclusionsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	Schedule             *struct {
		Enabled *bool   "json:\"enabled,omitempty\""
		Endtime *string "json:\"endtime,omitempty\""
		Rrules  struct {
			Bymonthday *int                   "json:\"bymonthday,omitempty\""
			Byweekday  *string                "json:\"byweekday,omitempty\""
			Freq       N200ScheduleRrulesFreq "json:\"freq\""
			Interval   *int                   "json:\"interval,omitempty\""
		} "json:\"rrules\""
		Starttime string "json:\"starttime\""
		Timezone  string "json:\"timezone\""
	} "json:\"schedule,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentExclusionsCreateWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreationDate *int "json:\"creation_date,omitempty\""; Description *string "json:\"description,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; Schedule *struct { Enabled *bool "json:\"enabled,omitempty\""; Endtime *string "json:\"endtime,omitempty\""; Rrules struct { Bymonthday *int "json:\"bymonthday,omitempty\""; Byweekday *string "json:\"byweekday,omitempty\""; Freq N200ScheduleRrulesFreq "json:\"freq\""; Interval *int "json:\"interval,omitempty\"" } "json:\"rrules\""; Starttime string "json:\"starttime\""; Timezone string "json:\"timezone\"" } "json:\"schedule,omitempty\"" }
	if i, ok := r.(*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Schedule             *struct {
			Enabled *bool   "json:\"enabled,omitempty\""
			Endtime *string "json:\"endtime,omitempty\""
			Rrules  struct {
				Bymonthday *int                   "json:\"bymonthday,omitempty\""
				Byweekday  *string                "json:\"byweekday,omitempty\""
				Freq       N200ScheduleRrulesFreq "json:\"freq\""
				Interval   *int                   "json:\"interval,omitempty\""
			} "json:\"rrules\""
			Starttime string "json:\"starttime\""
			Timezone  string "json:\"timezone\""
		} "json:\"schedule,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Delete calls the AgentExclusions ´s function
func (c *AgentExclusions) Delete(arg1 int32, arg2 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentExclusionsDeleteWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Details calls the AgentExclusions ´s function
func (c *AgentExclusions) Details(arg1 int32, arg2 int32, reqEditors ...RequestEditorFn) (*struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	Schedule             *struct {
		Enabled *bool   "json:\"enabled,omitempty\""
		Endtime *string "json:\"endtime,omitempty\""
		Rrules  struct {
			Bymonthday *int                   "json:\"bymonthday,omitempty\""
			Byweekday  *string                "json:\"byweekday,omitempty\""
			Freq       N200ScheduleRrulesFreq "json:\"freq\""
			Interval   *int                   "json:\"interval,omitempty\""
		} "json:\"rrules\""
		Starttime string "json:\"starttime\""
		Timezone  string "json:\"timezone\""
	} "json:\"schedule,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentExclusionsDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreationDate *int "json:\"creation_date,omitempty\""; Description *string "json:\"description,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; Schedule *struct { Enabled *bool "json:\"enabled,omitempty\""; Endtime *string "json:\"endtime,omitempty\""; Rrules struct { Bymonthday *int "json:\"bymonthday,omitempty\""; Byweekday *string "json:\"byweekday,omitempty\""; Freq N200ScheduleRrulesFreq "json:\"freq\""; Interval *int "json:\"interval,omitempty\"" } "json:\"rrules\""; Starttime string "json:\"starttime\""; Timezone string "json:\"timezone\"" } "json:\"schedule,omitempty\"" }
	if i, ok := r.(*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Schedule             *struct {
			Enabled *bool   "json:\"enabled,omitempty\""
			Endtime *string "json:\"endtime,omitempty\""
			Rrules  struct {
				Bymonthday *int                   "json:\"bymonthday,omitempty\""
				Byweekday  *string                "json:\"byweekday,omitempty\""
				Freq       N200ScheduleRrulesFreq "json:\"freq\""
				Interval   *int                   "json:\"interval,omitempty\""
			} "json:\"rrules\""
			Starttime string "json:\"starttime\""
			Timezone  string "json:\"timezone\""
		} "json:\"schedule,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// EditWithBody calls the AgentExclusions ´s function
func (c *AgentExclusions) EditWithBody(arg1 int32, arg2 int32, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentExclusionsEditWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Edit calls the AgentExclusions ´s function
func (c *AgentExclusions) Edit(arg1 int32, arg2 int32, arg3 AgentExclusionsEditJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentExclusionsEditWithResponse(c.ClientInterface.(*Client).ctx, arg1, arg2, arg3, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// List calls the AgentExclusions ´s function
func (c *AgentExclusions) List(arg1 int32, reqEditors ...RequestEditorFn) (*[]struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	Schedule             *struct {
		Enabled *bool   "json:\"enabled,omitempty\""
		Endtime *string "json:\"endtime,omitempty\""
		Rrules  struct {
			Bymonthday *int                   "json:\"bymonthday,omitempty\""
			Byweekday  *string                "json:\"byweekday,omitempty\""
			Freq       N200ScheduleRrulesFreq "json:\"freq\""
			Interval   *int                   "json:\"interval,omitempty\""
		} "json:\"rrules\""
		Starttime string "json:\"starttime\""
		Timezone  string "json:\"timezone\""
	} "json:\"schedule,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AgentExclusionsListWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { CreationDate *int "json:\"creation_date,omitempty\""; Description *string "json:\"description,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; Schedule *struct { Enabled *bool "json:\"enabled,omitempty\""; Endtime *string "json:\"endtime,omitempty\""; Rrules struct { Bymonthday *int "json:\"bymonthday,omitempty\""; Byweekday *string "json:\"byweekday,omitempty\""; Freq N200ScheduleRrulesFreq "json:\"freq\""; Interval *int "json:\"interval,omitempty\"" } "json:\"rrules\""; Starttime string "json:\"starttime\""; Timezone string "json:\"timezone\"" } "json:\"schedule,omitempty\"" }
	if i, ok := r.(*[]struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		Schedule             *struct {
			Enabled *bool   "json:\"enabled,omitempty\""
			Endtime *string "json:\"endtime,omitempty\""
			Rrules  struct {
				Bymonthday *int                   "json:\"bymonthday,omitempty\""
				Byweekday  *string                "json:\"byweekday,omitempty\""
				Freq       N200ScheduleRrulesFreq "json:\"freq\""
				Interval   *int                   "json:\"interval,omitempty\""
			} "json:\"rrules\""
			Starttime string "json:\"starttime\""
			Timezone  string "json:\"timezone\""
		} "json:\"schedule,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Assets implementation of the Assets interface
type Assets struct {
	*ClientWithResponses
}

// AssetInfo calls the Assets ´s function
func (c *Assets) AssetInfo(contentType string, reqEditors ...RequestEditorFn) (*struct {
	AcrDrivers *[]struct {
		DriverName  *string   "json:\"driver_name,omitempty\""
		DriverValue *[]string "json:\"driver_value,omitempty\""
	} "json:\"acr_drivers,omitempty\""
	AcrScore                  *int      "json:\"acr_score,omitempty\""
	AgentName                 *[]string "json:\"agent_name,omitempty\""
	AwsAvailabilityZone       *[]string "json:\"aws_availability_zone,omitempty\""
	AwsEc2InstanceAmiId       *[]string "json:\"aws_ec2_instance_ami_id,omitempty\""
	AwsEc2InstanceGroupName   *[]string "json:\"aws_ec2_instance_group_name,omitempty\""
	AwsEc2InstanceId          *[]string "json:\"aws_ec2_instance_id,omitempty\""
	AwsEc2InstanceStateName   *[]string "json:\"aws_ec2_instance_state_name,omitempty\""
	AwsEc2InstanceType        *[]string "json:\"aws_ec2_instance_type,omitempty\""
	AwsEc2Name                *[]string "json:\"aws_ec2_name,omitempty\""
	AwsEc2ProductCode         *[]string "json:\"aws_ec2_product_code,omitempty\""
	AwsOwnerId                *[]string "json:\"aws_owner_id,omitempty\""
	AwsRegion                 *[]string "json:\"aws_region,omitempty\""
	AwsSubnetId               *[]string "json:\"aws_subnet_id,omitempty\""
	AwsVpcId                  *[]string "json:\"aws_vpc_id,omitempty\""
	AzureResourceId           *[]string "json:\"azure_resource_id,omitempty\""
	AzureVmId                 *[]string "json:\"azure_vm_id,omitempty\""
	BiosUuid                  *[]string "json:\"bios_uuid,omitempty\""
	CreatedAt                 *string   "json:\"created_at,omitempty\""
	ExposureScore             *int      "json:\"exposure_score,omitempty\""
	FirstSeen                 *string   "json:\"first_seen,omitempty\""
	Fqdn                      *[]string "json:\"fqdn,omitempty\""
	GcpInstanceId             *[]string "json:\"gcp_instance_id,omitempty\""
	GcpProjectId              *[]string "json:\"gcp_project_id,omitempty\""
	GcpZone                   *[]string "json:\"gcp_zone,omitempty\""
	HasAgent                  *bool     "json:\"has_agent,omitempty\""
	Hostname                  *[]string "json:\"hostname,omitempty\""
	Id                        *string   "json:\"id,omitempty\""
	InstalledSoftware         *[]string "json:\"installed_software,omitempty\""
	Ipv4                      *[]string "json:\"ipv4,omitempty\""
	Ipv6                      *[]string "json:\"ipv6,omitempty\""
	LastAuthenticatedScanDate *string   "json:\"last_authenticated_scan_date,omitempty\""
	LastLicensedScanDate      *string   "json:\"last_licensed_scan_date,omitempty\""
	LastScanTarget            *string   "json:\"last_scan_target,omitempty\""
	LastSeen                  *string   "json:\"last_seen,omitempty\""
	MacAddress                *[]string "json:\"mac_address,omitempty\""
	McafeeEpoAgentGuid        *[]string "json:\"mcafee_epo_agent_guid,omitempty\""
	McafeeEpoGuid             *[]string "json:\"mcafee_epo_guid,omitempty\""
	NetbiosName               *[]string "json:\"netbios_name,omitempty\""
	NetworkId                 *[]string "json:\"network_id,omitempty\""
	OperatingSystem           *[]string "json:\"operating_system,omitempty\""
	QualysAssetId             *[]string "json:\"qualys_asset_id,omitempty\""
	QualysHostId              *[]string "json:\"qualys_host_id,omitempty\""
	ScanFrequency             *[]struct {
		Frequency *int  "json:\"frequency,omitempty\""
		Interval  *int  "json:\"interval,omitempty\""
		Licensed  *bool "json:\"licensed,omitempty\""
	} "json:\"scan_frequency,omitempty\""
	ServicenowSysid *[]string "json:\"servicenow_sysid,omitempty\""
	Sources         *[][]struct {
		FirstSeen *string "json:\"first_seen,omitempty\""
		LastSeen  *string "json:\"last_seen,omitempty\""
		Name      *string "json:\"name,omitempty\""
	} "json:\"sources,omitempty\""
	SshFingerprint *[]string "json:\"ssh_fingerprint,omitempty\""
	SystemType     *[]string "json:\"system_type,omitempty\""
	Tags           *[]struct {
		AddedAt  *string "json:\"added_at,omitempty\""
		AddedBy  *string "json:\"added_by,omitempty\""
		TagKey   *string "json:\"tag_key,omitempty\""
		TagUuid  *string "json:\"tag_uuid,omitempty\""
		TagValue *string "json:\"tag_value,omitempty\""
	} "json:\"tags,omitempty\""
	TenableUuid *[]string "json:\"tenable_uuid,omitempty\""
	UpdatedAt   *string   "json:\"updated_at,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AssetsAssetInfoWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AcrDrivers *[]struct { DriverName *string "json:\"driver_name,omitempty\""; DriverValue *[]string "json:\"driver_value,omitempty\"" } "json:\"acr_drivers,omitempty\""; AcrScore *int "json:\"acr_score,omitempty\""; AgentName *[]string "json:\"agent_name,omitempty\""; AwsAvailabilityZone *[]string "json:\"aws_availability_zone,omitempty\""; AwsEc2InstanceAmiId *[]string "json:\"aws_ec2_instance_ami_id,omitempty\""; AwsEc2InstanceGroupName *[]string "json:\"aws_ec2_instance_group_name,omitempty\""; AwsEc2InstanceId *[]string "json:\"aws_ec2_instance_id,omitempty\""; AwsEc2InstanceStateName *[]string "json:\"aws_ec2_instance_state_name,omitempty\""; AwsEc2InstanceType *[]string "json:\"aws_ec2_instance_type,omitempty\""; AwsEc2Name *[]string "json:\"aws_ec2_name,omitempty\""; AwsEc2ProductCode *[]string "json:\"aws_ec2_product_code,omitempty\""; AwsOwnerId *[]string "json:\"aws_owner_id,omitempty\""; AwsRegion *[]string "json:\"aws_region,omitempty\""; AwsSubnetId *[]string "json:\"aws_subnet_id,omitempty\""; AwsVpcId *[]string "json:\"aws_vpc_id,omitempty\""; AzureResourceId *[]string "json:\"azure_resource_id,omitempty\""; AzureVmId *[]string "json:\"azure_vm_id,omitempty\""; BiosUuid *[]string "json:\"bios_uuid,omitempty\""; CreatedAt *string "json:\"created_at,omitempty\""; ExposureScore *int "json:\"exposure_score,omitempty\""; FirstSeen *string "json:\"first_seen,omitempty\""; Fqdn *[]string "json:\"fqdn,omitempty\""; GcpInstanceId *[]string "json:\"gcp_instance_id,omitempty\""; GcpProjectId *[]string "json:\"gcp_project_id,omitempty\""; GcpZone *[]string "json:\"gcp_zone,omitempty\""; HasAgent *bool "json:\"has_agent,omitempty\""; Hostname *[]string "json:\"hostname,omitempty\""; Id *string "json:\"id,omitempty\""; InstalledSoftware *[]string "json:\"installed_software,omitempty\""; Ipv4 *[]string "json:\"ipv4,omitempty\""; Ipv6 *[]string "json:\"ipv6,omitempty\""; LastAuthenticatedScanDate *string "json:\"last_authenticated_scan_date,omitempty\""; LastLicensedScanDate *string "json:\"last_licensed_scan_date,omitempty\""; LastScanTarget *string "json:\"last_scan_target,omitempty\""; LastSeen *string "json:\"last_seen,omitempty\""; MacAddress *[]string "json:\"mac_address,omitempty\""; McafeeEpoAgentGuid *[]string "json:\"mcafee_epo_agent_guid,omitempty\""; McafeeEpoGuid *[]string "json:\"mcafee_epo_guid,omitempty\""; NetbiosName *[]string "json:\"netbios_name,omitempty\""; NetworkId *[]string "json:\"network_id,omitempty\""; OperatingSystem *[]string "json:\"operating_system,omitempty\""; QualysAssetId *[]string "json:\"qualys_asset_id,omitempty\""; QualysHostId *[]string "json:\"qualys_host_id,omitempty\""; ScanFrequency *[]struct { Frequency *int "json:\"frequency,omitempty\""; Interval *int "json:\"interval,omitempty\""; Licensed *bool "json:\"licensed,omitempty\"" } "json:\"scan_frequency,omitempty\""; ServicenowSysid *[]string "json:\"servicenow_sysid,omitempty\""; Sources *[][]struct { FirstSeen *string "json:\"first_seen,omitempty\""; LastSeen *string "json:\"last_seen,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"sources,omitempty\""; SshFingerprint *[]string "json:\"ssh_fingerprint,omitempty\""; SystemType *[]string "json:\"system_type,omitempty\""; Tags *[]struct { AddedAt *string "json:\"added_at,omitempty\""; AddedBy *string "json:\"added_by,omitempty\""; TagKey *string "json:\"tag_key,omitempty\""; TagUuid *string "json:\"tag_uuid,omitempty\""; TagValue *string "json:\"tag_value,omitempty\"" } "json:\"tags,omitempty\""; TenableUuid *[]string "json:\"tenable_uuid,omitempty\""; UpdatedAt *string "json:\"updated_at,omitempty\"" }
	if i, ok := r.(*struct {
		AcrDrivers *[]struct {
			DriverName  *string   "json:\"driver_name,omitempty\""
			DriverValue *[]string "json:\"driver_value,omitempty\""
		} "json:\"acr_drivers,omitempty\""
		AcrScore                  *int      "json:\"acr_score,omitempty\""
		AgentName                 *[]string "json:\"agent_name,omitempty\""
		AwsAvailabilityZone       *[]string "json:\"aws_availability_zone,omitempty\""
		AwsEc2InstanceAmiId       *[]string "json:\"aws_ec2_instance_ami_id,omitempty\""
		AwsEc2InstanceGroupName   *[]string "json:\"aws_ec2_instance_group_name,omitempty\""
		AwsEc2InstanceId          *[]string "json:\"aws_ec2_instance_id,omitempty\""
		AwsEc2InstanceStateName   *[]string "json:\"aws_ec2_instance_state_name,omitempty\""
		AwsEc2InstanceType        *[]string "json:\"aws_ec2_instance_type,omitempty\""
		AwsEc2Name                *[]string "json:\"aws_ec2_name,omitempty\""
		AwsEc2ProductCode         *[]string "json:\"aws_ec2_product_code,omitempty\""
		AwsOwnerId                *[]string "json:\"aws_owner_id,omitempty\""
		AwsRegion                 *[]string "json:\"aws_region,omitempty\""
		AwsSubnetId               *[]string "json:\"aws_subnet_id,omitempty\""
		AwsVpcId                  *[]string "json:\"aws_vpc_id,omitempty\""
		AzureResourceId           *[]string "json:\"azure_resource_id,omitempty\""
		AzureVmId                 *[]string "json:\"azure_vm_id,omitempty\""
		BiosUuid                  *[]string "json:\"bios_uuid,omitempty\""
		CreatedAt                 *string   "json:\"created_at,omitempty\""
		ExposureScore             *int      "json:\"exposure_score,omitempty\""
		FirstSeen                 *string   "json:\"first_seen,omitempty\""
		Fqdn                      *[]string "json:\"fqdn,omitempty\""
		GcpInstanceId             *[]string "json:\"gcp_instance_id,omitempty\""
		GcpProjectId              *[]string "json:\"gcp_project_id,omitempty\""
		GcpZone                   *[]string "json:\"gcp_zone,omitempty\""
		HasAgent                  *bool     "json:\"has_agent,omitempty\""
		Hostname                  *[]string "json:\"hostname,omitempty\""
		Id                        *string   "json:\"id,omitempty\""
		InstalledSoftware         *[]string "json:\"installed_software,omitempty\""
		Ipv4                      *[]string "json:\"ipv4,omitempty\""
		Ipv6                      *[]string "json:\"ipv6,omitempty\""
		LastAuthenticatedScanDate *string   "json:\"last_authenticated_scan_date,omitempty\""
		LastLicensedScanDate      *string   "json:\"last_licensed_scan_date,omitempty\""
		LastScanTarget            *string   "json:\"last_scan_target,omitempty\""
		LastSeen                  *string   "json:\"last_seen,omitempty\""
		MacAddress                *[]string "json:\"mac_address,omitempty\""
		McafeeEpoAgentGuid        *[]string "json:\"mcafee_epo_agent_guid,omitempty\""
		McafeeEpoGuid             *[]string "json:\"mcafee_epo_guid,omitempty\""
		NetbiosName               *[]string "json:\"netbios_name,omitempty\""
		NetworkId                 *[]string "json:\"network_id,omitempty\""
		OperatingSystem           *[]string "json:\"operating_system,omitempty\""
		QualysAssetId             *[]string "json:\"qualys_asset_id,omitempty\""
		QualysHostId              *[]string "json:\"qualys_host_id,omitempty\""
		ScanFrequency             *[]struct {
			Frequency *int  "json:\"frequency,omitempty\""
			Interval  *int  "json:\"interval,omitempty\""
			Licensed  *bool "json:\"licensed,omitempty\""
		} "json:\"scan_frequency,omitempty\""
		ServicenowSysid *[]string "json:\"servicenow_sysid,omitempty\""
		Sources         *[][]struct {
			FirstSeen *string "json:\"first_seen,omitempty\""
			LastSeen  *string "json:\"last_seen,omitempty\""
			Name      *string "json:\"name,omitempty\""
		} "json:\"sources,omitempty\""
		SshFingerprint *[]string "json:\"ssh_fingerprint,omitempty\""
		SystemType     *[]string "json:\"system_type,omitempty\""
		Tags           *[]struct {
			AddedAt  *string "json:\"added_at,omitempty\""
			AddedBy  *string "json:\"added_by,omitempty\""
			TagKey   *string "json:\"tag_key,omitempty\""
			TagUuid  *string "json:\"tag_uuid,omitempty\""
			TagValue *string "json:\"tag_value,omitempty\""
		} "json:\"tags,omitempty\""
		TenableUuid *[]string "json:\"tenable_uuid,omitempty\""
		UpdatedAt   *string   "json:\"updated_at,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// BulkDeleteWithBody calls the Assets ´s function
func (c *Assets) BulkDeleteWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*AssetsBulkDeleteResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AssetsBulkDeleteWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *AssetsBulkDeleteResponse
	if i, ok := r.(*AssetsBulkDeleteResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// BulkDelete calls the Assets ´s function
func (c *Assets) BulkDelete(arg1 AssetsBulkDeleteJSONRequestBody, reqEditors ...RequestEditorFn) (*AssetsBulkDeleteResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AssetsBulkDeleteWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *AssetsBulkDeleteResponse
	if i, ok := r.(*AssetsBulkDeleteResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// BulkMoveWithBody calls the Assets ´s function
func (c *Assets) BulkMoveWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*AssetsBulkMoveResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AssetsBulkMoveWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *AssetsBulkMoveResponse
	if i, ok := r.(*AssetsBulkMoveResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// BulkMove calls the Assets ´s function
func (c *Assets) BulkMove(arg1 AssetsBulkMoveJSONRequestBody, reqEditors ...RequestEditorFn) (*AssetsBulkMoveResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AssetsBulkMoveWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *AssetsBulkMoveResponse
	if i, ok := r.(*AssetsBulkMoveResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// BulkUpdateAcrWithBody calls the Assets ´s function
func (c *Assets) BulkUpdateAcrWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*AssetsBulkUpdateAcrResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AssetsBulkUpdateAcrWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *AssetsBulkUpdateAcrResponse
	if i, ok := r.(*AssetsBulkUpdateAcrResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// BulkUpdateAcr calls the Assets ´s function
func (c *Assets) BulkUpdateAcr(arg1 AssetsBulkUpdateAcrJSONRequestBody, reqEditors ...RequestEditorFn) (*AssetsBulkUpdateAcrResponse, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AssetsBulkUpdateAcrWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *AssetsBulkUpdateAcrResponse
	if i, ok := r.(*AssetsBulkUpdateAcrResponse); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ImportJobInfo calls the Assets ´s function
func (c *Assets) ImportJobInfo(contentType string, reqEditors ...RequestEditorFn) (*struct {
	Batches        *int32  "json:\"batches,omitempty\""
	ContainerId    *string "json:\"container_id,omitempty\""
	EndTime        *int32  "json:\"end_time,omitempty\""
	FailedAssets   *int    "json:\"failed_assets,omitempty\""
	JobId          *string "json:\"job_id,omitempty\""
	LastUpdateTime *int32  "json:\"last_update_time,omitempty\""
	Source         *string "json:\"source,omitempty\""
	StartTime      *int32  "json:\"start_time,omitempty\""
	Status         *string "json:\"status,omitempty\""
	StatusMessage  *string "json:\"status_message,omitempty\""
	UploadedAssets *int32  "json:\"uploaded_assets,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AssetsImportJobInfoWithResponse(c.ClientInterface.(*Client).ctx, contentType, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Batches *int32 "json:\"batches,omitempty\""; ContainerId *string "json:\"container_id,omitempty\""; EndTime *int32 "json:\"end_time,omitempty\""; FailedAssets *int "json:\"failed_assets,omitempty\""; JobId *string "json:\"job_id,omitempty\""; LastUpdateTime *int32 "json:\"last_update_time,omitempty\""; Source *string "json:\"source,omitempty\""; StartTime *int32 "json:\"start_time,omitempty\""; Status *string "json:\"status,omitempty\""; StatusMessage *string "json:\"status_message,omitempty\""; UploadedAssets *int32 "json:\"uploaded_assets,omitempty\"" }
	if i, ok := r.(*struct {
		Batches        *int32  "json:\"batches,omitempty\""
		ContainerId    *string "json:\"container_id,omitempty\""
		EndTime        *int32  "json:\"end_time,omitempty\""
		FailedAssets   *int    "json:\"failed_assets,omitempty\""
		JobId          *string "json:\"job_id,omitempty\""
		LastUpdateTime *int32  "json:\"last_update_time,omitempty\""
		Source         *string "json:\"source,omitempty\""
		StartTime      *int32  "json:\"start_time,omitempty\""
		Status         *string "json:\"status,omitempty\""
		StatusMessage  *string "json:\"status_message,omitempty\""
		UploadedAssets *int32  "json:\"uploaded_assets,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ImportWithBody calls the Assets ´s function
func (c *Assets) ImportWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	AssetImportJobUuid *string "json:\"asset_import_job_uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AssetsImportWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AssetImportJobUuid *string "json:\"asset_import_job_uuid,omitempty\"" }
	if i, ok := r.(*struct {
		AssetImportJobUuid *string "json:\"asset_import_job_uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Import calls the Assets ´s function
func (c *Assets) Import(arg1 AssetsImportJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	AssetImportJobUuid *string "json:\"asset_import_job_uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AssetsImportWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { AssetImportJobUuid *string "json:\"asset_import_job_uuid,omitempty\"" }
	if i, ok := r.(*struct {
		AssetImportJobUuid *string "json:\"asset_import_job_uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ListAssets calls the Assets ´s function
func (c *Assets) ListAssets(reqEditors ...RequestEditorFn) (*struct {
	Assets *[]struct {
		AcrDrivers *[]struct {
			DriverName  *string   "json:\"driver_name,omitempty\""
			DriverValue *[]string "json:\"driver_value,omitempty\""
		} "json:\"acr_drivers,omitempty\""
		AcrScore        *int      "json:\"acr_score,omitempty\""
		AgentName       *[]string "json:\"agent_name,omitempty\""
		AwsEc2Name      *[]string "json:\"aws_ec2_name,omitempty\""
		ExposureScore   *int      "json:\"exposure_score,omitempty\""
		Fqdn            *[]string "json:\"fqdn,omitempty\""
		HasAgent        *bool     "json:\"has_agent,omitempty\""
		Id              *string   "json:\"id,omitempty\""
		Ipv4            *[]string "json:\"ipv4,omitempty\""
		Ipv6            *[]string "json:\"ipv6,omitempty\""
		LastScanTarget  *string   "json:\"last_scan_target,omitempty\""
		LastSeen        *string   "json:\"last_seen,omitempty\""
		MacAddress      *[]string "json:\"mac_address,omitempty\""
		NetbiosName     *[]string "json:\"netbios_name,omitempty\""
		OperatingSystem *[]string "json:\"operating_system,omitempty\""
		ScanFrequency   *[]struct {
			Frequency *int  "json:\"frequency,omitempty\""
			Interval  *int  "json:\"interval,omitempty\""
			Licensed  *bool "json:\"licensed,omitempty\""
		} "json:\"scan_frequency,omitempty\""
		Sources *[][]struct {
			FirstSeen *string "json:\"first_seen,omitempty\""
			LastSeen  *string "json:\"last_seen,omitempty\""
			Name      *string "json:\"name,omitempty\""
		} "json:\"sources,omitempty\""
	} "json:\"assets,omitempty\""
	Total *int "json:\"total,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AssetsListAssetsWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Assets *[]struct { AcrDrivers *[]struct { DriverName *string "json:\"driver_name,omitempty\""; DriverValue *[]string "json:\"driver_value,omitempty\"" } "json:\"acr_drivers,omitempty\""; AcrScore *int "json:\"acr_score,omitempty\""; AgentName *[]string "json:\"agent_name,omitempty\""; AwsEc2Name *[]string "json:\"aws_ec2_name,omitempty\""; ExposureScore *int "json:\"exposure_score,omitempty\""; Fqdn *[]string "json:\"fqdn,omitempty\""; HasAgent *bool "json:\"has_agent,omitempty\""; Id *string "json:\"id,omitempty\""; Ipv4 *[]string "json:\"ipv4,omitempty\""; Ipv6 *[]string "json:\"ipv6,omitempty\""; LastScanTarget *string "json:\"last_scan_target,omitempty\""; LastSeen *string "json:\"last_seen,omitempty\""; MacAddress *[]string "json:\"mac_address,omitempty\""; NetbiosName *[]string "json:\"netbios_name,omitempty\""; OperatingSystem *[]string "json:\"operating_system,omitempty\""; ScanFrequency *[]struct { Frequency *int "json:\"frequency,omitempty\""; Interval *int "json:\"interval,omitempty\""; Licensed *bool "json:\"licensed,omitempty\"" } "json:\"scan_frequency,omitempty\""; Sources *[][]struct { FirstSeen *string "json:\"first_seen,omitempty\""; LastSeen *string "json:\"last_seen,omitempty\""; Name *string "json:\"name,omitempty\"" } "json:\"sources,omitempty\"" } "json:\"assets,omitempty\""; Total *int "json:\"total,omitempty\"" }
	if i, ok := r.(*struct {
		Assets *[]struct {
			AcrDrivers *[]struct {
				DriverName  *string   "json:\"driver_name,omitempty\""
				DriverValue *[]string "json:\"driver_value,omitempty\""
			} "json:\"acr_drivers,omitempty\""
			AcrScore        *int      "json:\"acr_score,omitempty\""
			AgentName       *[]string "json:\"agent_name,omitempty\""
			AwsEc2Name      *[]string "json:\"aws_ec2_name,omitempty\""
			ExposureScore   *int      "json:\"exposure_score,omitempty\""
			Fqdn            *[]string "json:\"fqdn,omitempty\""
			HasAgent        *bool     "json:\"has_agent,omitempty\""
			Id              *string   "json:\"id,omitempty\""
			Ipv4            *[]string "json:\"ipv4,omitempty\""
			Ipv6            *[]string "json:\"ipv6,omitempty\""
			LastScanTarget  *string   "json:\"last_scan_target,omitempty\""
			LastSeen        *string   "json:\"last_seen,omitempty\""
			MacAddress      *[]string "json:\"mac_address,omitempty\""
			NetbiosName     *[]string "json:\"netbios_name,omitempty\""
			OperatingSystem *[]string "json:\"operating_system,omitempty\""
			ScanFrequency   *[]struct {
				Frequency *int  "json:\"frequency,omitempty\""
				Interval  *int  "json:\"interval,omitempty\""
				Licensed  *bool "json:\"licensed,omitempty\""
			} "json:\"scan_frequency,omitempty\""
			Sources *[][]struct {
				FirstSeen *string "json:\"first_seen,omitempty\""
				LastSeen  *string "json:\"last_seen,omitempty\""
				Name      *string "json:\"name,omitempty\""
			} "json:\"sources,omitempty\""
		} "json:\"assets,omitempty\""
		Total *int "json:\"total,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ListImportJobs calls the Assets ´s function
func (c *Assets) ListImportJobs(reqEditors ...RequestEditorFn) (*[]struct {
	Batches        *int32  "json:\"batches,omitempty\""
	ContainerId    *string "json:\"container_id,omitempty\""
	EndTime        *int32  "json:\"end_time,omitempty\""
	FailedAssets   *int    "json:\"failed_assets,omitempty\""
	JobId          *string "json:\"job_id,omitempty\""
	LastUpdateTime *int32  "json:\"last_update_time,omitempty\""
	Source         *string "json:\"source,omitempty\""
	StartTime      *int32  "json:\"start_time,omitempty\""
	Status         *string "json:\"status,omitempty\""
	StatusMessage  *string "json:\"status_message,omitempty\""
	UploadedAssets *int32  "json:\"uploaded_assets,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.AssetsListImportJobsWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { Batches *int32 "json:\"batches,omitempty\""; ContainerId *string "json:\"container_id,omitempty\""; EndTime *int32 "json:\"end_time,omitempty\""; FailedAssets *int "json:\"failed_assets,omitempty\""; JobId *string "json:\"job_id,omitempty\""; LastUpdateTime *int32 "json:\"last_update_time,omitempty\""; Source *string "json:\"source,omitempty\""; StartTime *int32 "json:\"start_time,omitempty\""; Status *string "json:\"status,omitempty\""; StatusMessage *string "json:\"status_message,omitempty\""; UploadedAssets *int32 "json:\"uploaded_assets,omitempty\"" }
	if i, ok := r.(*[]struct {
		Batches        *int32  "json:\"batches,omitempty\""
		ContainerId    *string "json:\"container_id,omitempty\""
		EndTime        *int32  "json:\"end_time,omitempty\""
		FailedAssets   *int    "json:\"failed_assets,omitempty\""
		JobId          *string "json:\"job_id,omitempty\""
		LastUpdateTime *int32  "json:\"last_update_time,omitempty\""
		Source         *string "json:\"source,omitempty\""
		StartTime      *int32  "json:\"start_time,omitempty\""
		Status         *string "json:\"status,omitempty\""
		StatusMessage  *string "json:\"status_message,omitempty\""
		UploadedAssets *int32  "json:\"uploaded_assets,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Vulnerabilities implementation of the Vulnerabilities interface
type Vulnerabilities struct {
	*ClientWithResponses
}

// ImportV2WithBody calls the Vulnerabilities ´s function
func (c *Vulnerabilities) ImportV2WithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	JobUuid *string "json:\"job_uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.VulnerabilitiesImportV2WithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { JobUuid *string "json:\"job_uuid,omitempty\"" }
	if i, ok := r.(*struct {
		JobUuid *string "json:\"job_uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ImportV2 calls the Vulnerabilities ´s function
func (c *Vulnerabilities) ImportV2(arg1 VulnerabilitiesImportV2JSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	JobUuid *string "json:\"job_uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.VulnerabilitiesImportV2WithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { JobUuid *string "json:\"job_uuid,omitempty\"" }
	if i, ok := r.(*struct {
		JobUuid *string "json:\"job_uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ImportWithBody calls the Vulnerabilities ´s function
func (c *Vulnerabilities) ImportWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.VulnerabilitiesImportWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Import calls the Vulnerabilities ´s function
func (c *Vulnerabilities) Import(arg1 VulnerabilitiesImportJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.VulnerabilitiesImportWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Policies implementation of the Policies interface
type Policies struct {
	*ClientWithResponses
}

// Configure calls the Policies ´s function
func (c *Policies) Configure(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.PoliciesConfigureWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Copy calls the Policies ´s function
func (c *Policies) Copy(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.PoliciesCopyWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// CreateWithBody calls the Policies ´s function
func (c *Policies) CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	PolicyId   *int    "json:\"policy_id,omitempty\""
	PolicyName *string "json:\"policy_name,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.PoliciesCreateWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { PolicyId *int "json:\"policy_id,omitempty\""; PolicyName *string "json:\"policy_name,omitempty\"" }
	if i, ok := r.(*struct {
		PolicyId   *int    "json:\"policy_id,omitempty\""
		PolicyName *string "json:\"policy_name,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Create calls the Policies ´s function
func (c *Policies) Create(arg1 PoliciesCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	PolicyId   *int    "json:\"policy_id,omitempty\""
	PolicyName *string "json:\"policy_name,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.PoliciesCreateWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { PolicyId *int "json:\"policy_id,omitempty\""; PolicyName *string "json:\"policy_name,omitempty\"" }
	if i, ok := r.(*struct {
		PolicyId   *int    "json:\"policy_id,omitempty\""
		PolicyName *string "json:\"policy_name,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Delete calls the Policies ´s function
func (c *Policies) Delete(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.PoliciesDeleteWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Details calls the Policies ´s function
func (c *Policies) Details(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
	Audits      *map[string]interface{} "json:\"audits,omitempty\""
	Credentials *map[string]interface{} "json:\"credentials,omitempty\""
	Plugins     *map[string]interface{} "json:\"plugins,omitempty\""
	Scap        *map[string]interface{} "json:\"scap,omitempty\""
	Settings    *map[string]interface{} "json:\"settings,omitempty\""
	Uuid        *string                 "json:\"uuid,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.PoliciesDetailsWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { Audits *map[string]interface {} "json:\"audits,omitempty\""; Credentials *map[string]interface {} "json:\"credentials,omitempty\""; Plugins *map[string]interface {} "json:\"plugins,omitempty\""; Scap *map[string]interface {} "json:\"scap,omitempty\""; Settings *map[string]interface {} "json:\"settings,omitempty\""; Uuid *string "json:\"uuid,omitempty\"" }
	if i, ok := r.(*struct {
		Audits      *map[string]interface{} "json:\"audits,omitempty\""
		Credentials *map[string]interface{} "json:\"credentials,omitempty\""
		Plugins     *map[string]interface{} "json:\"plugins,omitempty\""
		Scap        *map[string]interface{} "json:\"scap,omitempty\""
		Settings    *map[string]interface{} "json:\"settings,omitempty\""
		Uuid        *string                 "json:\"uuid,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Export calls the Policies ´s function
func (c *Policies) Export(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.PoliciesExportWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *interface {}
	if i, ok := r.(*interface{}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// ImportWithBody calls the Policies ´s function
func (c *Policies) ImportWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NoTarget             *string "json:\"no_target,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	OwnerId              *int    "json:\"owner_id,omitempty\""
	Private              *int    "json:\"private,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	TemplateUuid         *string "json:\"template_uuid,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.PoliciesImportWithBodyWithResponse(c.ClientInterface.(*Client).ctx, arg1, body, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreationDate *int "json:\"creation_date,omitempty\""; Description *string "json:\"description,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NoTarget *string "json:\"no_target,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; Private *int "json:\"private,omitempty\""; Shared *bool "json:\"shared,omitempty\""; TemplateUuid *string "json:\"template_uuid,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\"" }
	if i, ok := r.(*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NoTarget             *string "json:\"no_target,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		Private              *int    "json:\"private,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		TemplateUuid         *string "json:\"template_uuid,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// Import calls the Policies ´s function
func (c *Policies) Import(arg1 PoliciesImportJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NoTarget             *string "json:\"no_target,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	OwnerId              *int    "json:\"owner_id,omitempty\""
	Private              *int    "json:\"private,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	TemplateUuid         *string "json:\"template_uuid,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.PoliciesImportWithResponse(c.ClientInterface.(*Client).ctx, arg1, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *struct { CreationDate *int "json:\"creation_date,omitempty\""; Description *string "json:\"description,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NoTarget *string "json:\"no_target,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *int "json:\"owner_id,omitempty\""; Private *int "json:\"private,omitempty\""; Shared *bool "json:\"shared,omitempty\""; TemplateUuid *string "json:\"template_uuid,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\"" }
	if i, ok := r.(*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NoTarget             *string "json:\"no_target,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		Private              *int    "json:\"private,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		TemplateUuid         *string "json:\"template_uuid,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}

// List calls the Policies ´s function
func (c *Policies) List(reqEditors ...RequestEditorFn) (*[]struct {
	CreationDate         *int    "json:\"creation_date,omitempty\""
	Description          *string "json:\"description,omitempty\""
	Id                   *int    "json:\"id,omitempty\""
	LastModificationDate *int    "json:\"last_modification_date,omitempty\""
	Name                 *string "json:\"name,omitempty\""
	NoTarget             *bool   "json:\"no_target,omitempty\""
	Owner                *string "json:\"owner,omitempty\""
	OwnerId              *string "json:\"owner_id,omitempty\""
	Shared               *bool   "json:\"shared,omitempty\""
	TemplateUuid         *string "json:\"template_uuid,omitempty\""
	UserPermissions      *int    "json:\"user_permissions,omitempty\""
	Visibility           *int    "json:\"visibility,omitempty\""
}, error) {
	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.PoliciesListWithResponse(c.ClientInterface.(*Client).ctx, reqEditors...))
	if err != nil {
		return nil, err
	}
	// convert to *[]struct { CreationDate *int "json:\"creation_date,omitempty\""; Description *string "json:\"description,omitempty\""; Id *int "json:\"id,omitempty\""; LastModificationDate *int "json:\"last_modification_date,omitempty\""; Name *string "json:\"name,omitempty\""; NoTarget *bool "json:\"no_target,omitempty\""; Owner *string "json:\"owner,omitempty\""; OwnerId *string "json:\"owner_id,omitempty\""; Shared *bool "json:\"shared,omitempty\""; TemplateUuid *string "json:\"template_uuid,omitempty\""; UserPermissions *int "json:\"user_permissions,omitempty\""; Visibility *int "json:\"visibility,omitempty\"" }
	if i, ok := r.(*[]struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NoTarget             *bool   "json:\"no_target,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *string "json:\"owner_id,omitempty\""
		Shared               *bool   "json:\"shared,omitempty\""
		TemplateUuid         *string "json:\"template_uuid,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Visibility           *int    "json:\"visibility,omitempty\""
	}); ok {
		return i, nil
	}

	// cannot convert, return nil
	return nil, nil
}
