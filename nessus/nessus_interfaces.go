package nessus

import (
	"io"
	"time"
)

// AuditLogInterface contains all methods for the nessus  api
type AuditLogInterface interface {
	Events(params *AuditLogEventsParams, reqEditors ...RequestEditorFn) (*[]struct {
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
	}, error)
}

// EditorInterface contains all methods for the nessus  api
type EditorInterface interface {
	Audits(arg1 EditorAuditsParamsType, arg2 int32, arg3 int32, reqEditors ...RequestEditorFn) (*EditorAuditsResponse, error)
	Details(arg1 EditorDetailsParamsType, arg2 int32, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	ListTemplates(arg1 EditorListTemplatesParamsType, reqEditors ...RequestEditorFn) (*[]struct {
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
	}, error)
	PluginDescription(arg1 int32, arg2 int32, arg3 int32, reqEditors ...RequestEditorFn) (*struct {
		Plugindescription *struct {
			Pluginattributes *map[string]interface{} "json:\"pluginattributes,omitempty\""
			Pluginfamily     *string                 "json:\"pluginfamily,omitempty\""
			Pluginid         *int                    "json:\"pluginid,omitempty\""
			Pluginname       *string                 "json:\"pluginname,omitempty\""
			Severity         *string                 "json:\"severity,omitempty\""
		} "json:\"plugindescription,omitempty\""
	}, error)
	TemplateDetails(arg1 EditorTemplateDetailsParamsType, contentType string, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
}

// IoFiltersInterface contains all methods for the nessus  api
type IoFiltersInterface interface {
	AgentsList(reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AssetsListV2WithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AssetsListV2(arg1 IoFiltersAssetsListV2JSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AssetsList(reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	CredentialsList(reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	ScanHistoryList(reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	ScanList(reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	VulnerabilitiesWorkbenchListV2WithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	VulnerabilitiesWorkbenchListV2(arg1 IoFiltersVulnerabilitiesWorkbenchListV2JSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	VulnerabilitiesWorkbenchList(reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
}

// IoPluginsInterface contains all methods for the nessus  api
type IoPluginsInterface interface {
	Details(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
		Attributes *[]struct {
			AttributeName  *string "json:\"attribute_name,omitempty\""
			AttributeValue *string "json:\"attribute_value,omitempty\""
		} "json:\"attributes,omitempty\""
		FamilyName *string "json:\"family_name,omitempty\""
		Id         *int    "json:\"id,omitempty\""
		Name       *string "json:\"name,omitempty\""
	}, error)
	FamiliesList(params *IoPluginsFamiliesListParams, reqEditors ...RequestEditorFn) (*struct {
		Families *[]struct {
			Count *int    "json:\"count,omitempty\""
			Id    *int    "json:\"id,omitempty\""
			Name  *string "json:\"name,omitempty\""
		} "json:\"families,omitempty\""
	}, error)
	FamilyDetails(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
		Id      *int    "json:\"id,omitempty\""
		Name    *string "json:\"name,omitempty\""
		Plugins *[]struct {
			Id   *int    "json:\"id,omitempty\""
			Name *string "json:\"name,omitempty\""
		} "json:\"plugins,omitempty\""
	}, error)
	List(params *IoPluginsListParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
}

// IoV2Interface contains all methods for the nessus  api
type IoV2Interface interface {
	AccessGroupsCreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AccessGroupsCreate(arg1 IoV2AccessGroupsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AccessGroupsDelete(contentType string, reqEditors ...RequestEditorFn) (*IoV2AccessGroupsDeleteResponse, error)
	AccessGroupsDetails(contentType string, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AccessGroupsEditWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AccessGroupsEdit(arg1 string, arg2 IoV2AccessGroupsEditJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AccessGroupsListFilters(reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AccessGroupsListRuleFilters(reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AccessGroupsList(params *IoV2AccessGroupsListParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
}

// AgentConfigInterface contains all methods for the nessus  api
type AgentConfigInterface interface {
	Details(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
		AutoUnlink *struct {
			Enabled    *bool  "json:\"enabled,omitempty\""
			Expiration *int32 "json:\"expiration,omitempty\""
		} "json:\"auto_unlink,omitempty\""
		SoftwareUpdate *bool "json:\"software_update,omitempty\""
	}, error)
	EditWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		AutoUnlink *struct {
			Enabled    *bool  "json:\"enabled,omitempty\""
			Expiration *int32 "json:\"expiration,omitempty\""
		} "json:\"auto_unlink,omitempty\""
		SoftwareUpdate *bool "json:\"software_update,omitempty\""
	}, error)
	Edit(arg1 int32, arg2 AgentConfigEditJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		AutoUnlink *struct {
			Enabled    *bool  "json:\"enabled,omitempty\""
			Expiration *int32 "json:\"expiration,omitempty\""
		} "json:\"auto_unlink,omitempty\""
		SoftwareUpdate *bool "json:\"software_update,omitempty\""
	}, error)
}

// AgentGroupInterface contains all methods for the nessus  api
type AgentGroupInterface interface {
	ListAgents(arg1 int32, arg2 int32, params *AgentGroupListAgentsParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	sAddAgent(arg1 int32, arg2 int32, arg3 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	sConfigureWithBody(arg1 int32, arg2 int32, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error)
	sConfigure(arg1 int32, arg2 int32, arg3 AgentGroupsConfigureJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error)
	sCreateWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		Agents               *[]string "json:\"agents,omitempty\""
		CreationDate         *int      "json:\"creation_date,omitempty\""
		Id                   *int      "json:\"id,omitempty\""
		LastModificationDate *int      "json:\"last_modification_date,omitempty\""
		Name                 *string   "json:\"name,omitempty\""
		Owner                *string   "json:\"owner,omitempty\""
		OwnerId              *string   "json:\"owner_id,omitempty\""
		OwnerName            *string   "json:\"owner_name,omitempty\""
		OwnerUuid            *string   "json:\"owner_uuid,omitempty\""
		Shared               *int      "json:\"shared,omitempty\""
		UserPermissions      *int      "json:\"user_permissions,omitempty\""
		Uuid                 *string   "json:\"uuid,omitempty\""
	}, error)
	sCreate(arg1 int32, arg2 AgentGroupsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		Agents               *[]string "json:\"agents,omitempty\""
		CreationDate         *int      "json:\"creation_date,omitempty\""
		Id                   *int      "json:\"id,omitempty\""
		LastModificationDate *int      "json:\"last_modification_date,omitempty\""
		Name                 *string   "json:\"name,omitempty\""
		Owner                *string   "json:\"owner,omitempty\""
		OwnerId              *string   "json:\"owner_id,omitempty\""
		OwnerName            *string   "json:\"owner_name,omitempty\""
		OwnerUuid            *string   "json:\"owner_uuid,omitempty\""
		Shared               *int      "json:\"shared,omitempty\""
		UserPermissions      *int      "json:\"user_permissions,omitempty\""
		Uuid                 *string   "json:\"uuid,omitempty\""
	}, error)
	sDeleteAgent(arg1 int32, arg2 int32, arg3 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	sDelete(arg1 int32, arg2 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	sDetails(arg1 int32, arg2 int32, params *AgentGroupsDetailsParams, reqEditors ...RequestEditorFn) (*struct {
		Agents               *[]string "json:\"agents,omitempty\""
		CreationDate         *int      "json:\"creation_date,omitempty\""
		Id                   *int      "json:\"id,omitempty\""
		LastModificationDate *int      "json:\"last_modification_date,omitempty\""
		Name                 *string   "json:\"name,omitempty\""
		Owner                *string   "json:\"owner,omitempty\""
		OwnerId              *string   "json:\"owner_id,omitempty\""
		OwnerName            *string   "json:\"owner_name,omitempty\""
		OwnerUuid            *string   "json:\"owner_uuid,omitempty\""
		Shared               *int      "json:\"shared,omitempty\""
		UserPermissions      *int      "json:\"user_permissions,omitempty\""
		Uuid                 *string   "json:\"uuid,omitempty\""
	}, error)
	sList(arg1 int32, reqEditors ...RequestEditorFn) (*[]struct {
		Agents               *[]string "json:\"agents,omitempty\""
		CreationDate         *int      "json:\"creation_date,omitempty\""
		Id                   *int      "json:\"id,omitempty\""
		LastModificationDate *int      "json:\"last_modification_date,omitempty\""
		Name                 *string   "json:\"name,omitempty\""
		Owner                *string   "json:\"owner,omitempty\""
		OwnerId              *string   "json:\"owner_id,omitempty\""
		OwnerName            *string   "json:\"owner_name,omitempty\""
		OwnerUuid            *string   "json:\"owner_uuid,omitempty\""
		Shared               *int      "json:\"shared,omitempty\""
		UserPermissions      *int      "json:\"user_permissions,omitempty\""
		Uuid                 *string   "json:\"uuid,omitempty\""
	}, error)
}

// IoScansInterface contains all methods for the nessus  api
type IoScansInterface interface {
	CheckAutoTargetsWithBody(params *IoScansCheckAutoTargetsParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		MatchedResourceUuids      *[]string "json:\"matched_resource_uuids,omitempty\""
		MissedTargets             *[]string "json:\"missed_targets,omitempty\""
		TotalMatchedResourceUuids *int      "json:\"total_matched_resource_uuids,omitempty\""
		TotalMissedTargets        *int      "json:\"total_missed_targets,omitempty\""
	}, error)
	CheckAutoTargets(params *IoScansCheckAutoTargetsParams, arg2 IoScansCheckAutoTargetsJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		MatchedResourceUuids      *[]string "json:\"matched_resource_uuids,omitempty\""
		MissedTargets             *[]string "json:\"missed_targets,omitempty\""
		TotalMatchedResourceUuids *int      "json:\"total_matched_resource_uuids,omitempty\""
		TotalMissedTargets        *int      "json:\"total_missed_targets,omitempty\""
	}, error)
	CredentialsConvertWithBody(arg1 int, arg2 int, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		Uuid *string "json:\"uuid,omitempty\""
	}, error)
	CredentialsConvert(arg1 int, arg2 int, arg3 IoScansCredentialsConvertJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		Uuid *string "json:\"uuid,omitempty\""
	}, error)
	RemediationCreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	RemediationCreate(arg1 IoScansRemediationCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	RemediationList(params *IoScansRemediationListParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
}

// IoV1Interface contains all methods for the nessus  api
type IoV1Interface interface {
	AccessGroupsCreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AccessGroupsCreate(arg1 IoV1AccessGroupsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AccessGroupsDelete(contentType string, reqEditors ...RequestEditorFn) (*IoV1AccessGroupsDeleteResponse, error)
	AccessGroupsDetails(contentType string, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AccessGroupsEditWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AccessGroupsEdit(arg1 string, arg2 IoV1AccessGroupsEditJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AccessGroupsListFilters(reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AccessGroupsListRuleFilters(reqEditors ...RequestEditorFn) (*struct {
		Filters *[]struct {
			Control      *string   "json:\"control,omitempty\""
			Name         *string   "json:\"name,omitempty\""
			Operators    *[]string "json:\"operators,omitempty\""
			ReadableName *string   "json:\"readable_name,omitempty\""
		} "json:\"filters,omitempty\""
	}, error)
	AccessGroupsList(params *IoV1AccessGroupsListParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
}

// ExclusionsInterface contains all methods for the nessus  api
type ExclusionsInterface interface {
	CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	Create(arg1 ExclusionsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	Delete(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	Details(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	EditWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error)
	Edit(arg1 int32, arg2 ExclusionsEditJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error)
	ImportWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error)
	Import(arg1 ExclusionsImportJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error)
	List(reqEditors ...RequestEditorFn) (*[]struct {
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
	}, error)
}

// IoScannerInterface contains all methods for the nessus  api
type IoScannerInterface interface {
	GroupsListRoutes(arg1 int32, reqEditors ...RequestEditorFn) (*[]struct {
		Route string "json:\"route\""
	}, error)
	GroupsUpdateRoutesWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error)
	GroupsUpdateRoutes(arg1 int32, arg2 IoScannerGroupsUpdateRoutesJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error)
}

// IoAgentInterface contains all methods for the nessus  api
type IoAgentInterface interface {
	BulkOperationsAddToNetworkWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	BulkOperationsAddToNetwork(arg1 IoAgentBulkOperationsAddToNetworkJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	BulkOperationsDirectiveWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	BulkOperationsDirective(arg1 IoAgentBulkOperationsDirectiveJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	BulkOperationsGroupDirectiveWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	BulkOperationsGroupDirective(arg1 int32, arg2 IoAgentBulkOperationsGroupDirectiveJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	BulkOperationsRemoveFromNetworkWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	BulkOperationsRemoveFromNetwork(arg1 IoAgentBulkOperationsRemoveFromNetworkJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
}

// ScannersInterface contains all methods for the nessus  api
type ScannersInterface interface {
	ControlScansWithBody(arg1 int32, arg2 string, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error)
	ControlScans(arg1 int32, arg2 string, arg3 ScannersControlScansJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error)
	Delete(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	Details(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
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
		UiBuild            *int    "json:\"ui_build,omitempty\""
		UiVersion          *string "json:\"ui_version,omitempty\""
		UserPermissions    *int    "json:\"user_permissions,omitempty\""
		Uuid               *string "json:\"uuid,omitempty\""
	}, error)
	EditWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	Edit(arg1 int32, arg2 ScannersEditJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	GetAwsTargets(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
		InstanceId *string "json:\"instance_id,omitempty\""
		Name       *string "json:\"name,omitempty\""
		PrivateIp  *string "json:\"private_ip,omitempty\""
		PublicIp   *string "json:\"public_ip,omitempty\""
		ScannerId  *int    "json:\"scanner_id,omitempty\""
		State      *string "json:\"state,omitempty\""
		Type       *string "json:\"type,omitempty\""
		Zone       *string "json:\"zone,omitempty\""
	}, error)
	GetScannerKey(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
		ScannerId *string "json:\"scanner_id,omitempty\""
	}, error)
	GetScans(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	List(reqEditors ...RequestEditorFn) (*[]struct {
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
		UiBuild            *int    "json:\"ui_build,omitempty\""
		UiVersion          *string "json:\"ui_version,omitempty\""
		UserPermissions    *int    "json:\"user_permissions,omitempty\""
		Uuid               *string "json:\"uuid,omitempty\""
	}, error)
	ToggleLinkStateWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error)
	ToggleLinkState(arg1 int32, arg2 ScannersToggleLinkStateJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error)
}

// ScansInterface contains all methods for the nessus  api
type ScansInterface interface {
	Attachments(arg1 string, arg2 string, params *ScansAttachmentsParams, reqEditors ...RequestEditorFn) (*ScansAttachmentsResponse, error)
	ConfigureWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	Configure(arg1 string, arg2 ScansConfigureJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	CopyWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	Copy(arg1 string, arg2 ScansCopyJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	Create(arg1 ScansCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	DeleteHistory(arg1 string, arg2 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	Delete(contentType string, reqEditors ...RequestEditorFn) (*interface{}, error)
	Details(arg1 string, params *ScansDetailsParams, reqEditors ...RequestEditorFn) (*struct {
		Comphosts *[]struct {
			Critical              *int                    "json:\"critical,omitempty\""
			High                  *int                    "json:\"high,omitempty\""
			HostId                *int                    "json:\"host_id,omitempty\""
			HostIndex             *string                 "json:\"host_index,omitempty\""
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
			HostIndex             *string                 "json:\"host_index,omitempty\""
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
			ScannerEnd      *string   "json:\"scanner_end,omitempty\""
			ScannerName     *string   "json:\"scanner_name,omitempty\""
			ScannerStart    *string   "json:\"scanner_start,omitempty\""
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
	}, error)
	ExportDownload(arg1 string, contentType string, reqEditors ...RequestEditorFn) (*ScansExportDownloadResponse, error)
	ExportRequestWithBody(arg1 string, params *ScansExportRequestParams, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		File      *string "json:\"file,omitempty\""
		TempToken *string "json:\"temp_token,omitempty\""
	}, error)
	ExportRequest(arg1 string, params *ScansExportRequestParams, arg3 ScansExportRequestJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		File      *string "json:\"file,omitempty\""
		TempToken *string "json:\"temp_token,omitempty\""
	}, error)
	ExportStatus(arg1 string, contentType string, reqEditors ...RequestEditorFn) (*struct {
		Status *string "json:\"status,omitempty\""
	}, error)
	GetLatestStatus(contentType string, reqEditors ...RequestEditorFn) (*struct {
		Status *string "json:\"status,omitempty\""
	}, error)
	HistoryDetails(arg1 string, contentType string, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	History(arg1 string, params *ScansHistoryParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	HostDetails(arg1 string, arg2 int32, params *ScansHostDetailsParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	ImportWithBody(params *ScansImportParams, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	Import(params *ScansImportParams, arg2 ScansImportJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	LaunchWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		ScanUuid *string "json:\"scan_uuid,omitempty\""
	}, error)
	Launch(arg1 string, arg2 ScansLaunchJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		ScanUuid *string "json:\"scan_uuid,omitempty\""
	}, error)
	List(params *ScansListParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	Pause(contentType string, reqEditors ...RequestEditorFn) (*interface{}, error)
	PluginOutput(arg1 string, arg2 int32, arg3 int32, params *ScansPluginOutputParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	ReadStatusWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error)
	ReadStatus(arg1 string, arg2 ScansReadStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error)
	Resume(contentType string, reqEditors ...RequestEditorFn) (*interface{}, error)
	ScheduleWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		Control   *bool   "json:\"control,omitempty\""
		Enabled   *bool   "json:\"enabled,omitempty\""
		Rrules    *string "json:\"rrules,omitempty\""
		Starttime *string "json:\"starttime,omitempty\""
		Timezone  *string "json:\"timezone,omitempty\""
	}, error)
	Schedule(arg1 string, arg2 ScansScheduleJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		Control   *bool   "json:\"control,omitempty\""
		Enabled   *bool   "json:\"enabled,omitempty\""
		Rrules    *string "json:\"rrules,omitempty\""
		Starttime *string "json:\"starttime,omitempty\""
		Timezone  *string "json:\"timezone,omitempty\""
	}, error)
	Stop(contentType string, reqEditors ...RequestEditorFn) (*interface{}, error)
	Timezones(reqEditors ...RequestEditorFn) (*[]struct {
		Name  *string "json:\"name,omitempty\""
		Value *string "json:\"value,omitempty\""
	}, error)
}

// TargetGroupsInterface contains all methods for the nessus  api
type TargetGroupsInterface interface {
	CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
		Shared               *int    "json:\"shared,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}, error)
	Create(arg1 TargetGroupsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
		Shared               *int    "json:\"shared,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}, error)
	Delete(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	Details(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
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
		Shared               *int    "json:\"shared,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}, error)
	EditWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
		Shared               *int    "json:\"shared,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}, error)
	Edit(arg1 int32, arg2 TargetGroupsEditJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
		Shared               *int    "json:\"shared,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}, error)
	List(reqEditors ...RequestEditorFn) (*[]struct {
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
		Shared               *int    "json:\"shared,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}, error)
}

// WorkbenchesInterface contains all methods for the nessus  api
type WorkbenchesInterface interface {
	AssetInfo(arg1 string, params *WorkbenchesAssetInfoParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AssetVulnerabilities(arg1 string, params *WorkbenchesAssetVulnerabilitiesParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AssetVulnerabilityInfo(arg1 string, arg2 int32, params *WorkbenchesAssetVulnerabilityInfoParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AssetVulnerabilityOutput(arg1 string, arg2 int32, params *WorkbenchesAssetVulnerabilityOutputParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AssetsActivity(contentType string, reqEditors ...RequestEditorFn) (*[]struct {
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
	}, error)
	AssetsDelete(contentType string, reqEditors ...RequestEditorFn) (*WorkbenchesAssetsDeleteResponse, error)
	AssetsVulnerabilities(params *WorkbenchesAssetsVulnerabilitiesParams, reqEditors ...RequestEditorFn) (*[]struct {
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
	}, error)
	Assets(params *WorkbenchesAssetsParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	ExportDownload(arg1 int32, reqEditors ...RequestEditorFn) (*WorkbenchesExportDownloadResponse, error)
	ExportRequest(params *WorkbenchesExportRequestParams, reqEditors ...RequestEditorFn) (*struct {
		File *int "json:\"file,omitempty\""
	}, error)
	ExportStatus(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
		Progress      *string "json:\"progress,omitempty\""
		ProgressTotal *string "json:\"progress_total,omitempty\""
		Status        *string "json:\"status,omitempty\""
	}, error)
	Vulnerabilities(params *WorkbenchesVulnerabilitiesParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	VulnerabilityInfo(arg1 int32, params *WorkbenchesVulnerabilityInfoParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	VulnerabilityOutput(arg1 int32, params *WorkbenchesVulnerabilityOutputParams, reqEditors ...RequestEditorFn) (*[]struct {
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
	}, error)
}

// ExportsVulnsInterface contains all methods for the nessus  api
type ExportsVulnsInterface interface {
	DownloadChunk(arg1 string, arg2 int32, reqEditors ...RequestEditorFn) (*ExportsVulnsDownloadChunkResponse, error)
	ExportCancel(contentType string, reqEditors ...RequestEditorFn) (*struct {
		Status *string "json:\"status,omitempty\""
	}, error)
	ExportStatusRecent(reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	ExportStatus(contentType string, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	RequestExportWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		ExportUuid *string "json:\"export_uuid,omitempty\""
	}, error)
	RequestExport(arg1 ExportsVulnsRequestExportJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		ExportUuid *string "json:\"export_uuid,omitempty\""
	}, error)
}

// FoldersInterface contains all methods for the nessus  api
type FoldersInterface interface {
	CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		Id *int "json:\"id,omitempty\""
	}, error)
	Create(arg1 FoldersCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		Id *int "json:\"id,omitempty\""
	}, error)
	Delete(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	EditWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error)
	Edit(arg1 int32, arg2 FoldersEditJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error)
	List(reqEditors ...RequestEditorFn) (*[]struct {
		Custom      *int    "json:\"custom,omitempty\""
		DefaultTag  *int    "json:\"default_tag,omitempty\""
		Id          *int    "json:\"id,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Type        *string "json:\"type,omitempty\""
		UnreadCount *int    "json:\"unread_count,omitempty\""
	}, error)
}

// TagsInterface contains all methods for the nessus  api
type TagsInterface interface {
	AssignAssetTagsWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*TagsAssignAssetTagsResponse, error)
	AssignAssetTags(arg1 TagsAssignAssetTagsJSONRequestBody, reqEditors ...RequestEditorFn) (*TagsAssignAssetTagsResponse, error)
	CreateTagCategoryWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		CreatedAt   *string "json:\"created_at,omitempty\""
		CreatedBy   *string "json:\"created_by,omitempty\""
		Description *string "json:\"description,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Reserved    *bool   "json:\"reserved,omitempty\""
		UpdatedAt   *string "json:\"updated_at,omitempty\""
		UpdatedBy   *string "json:\"updated_by,omitempty\""
		Uuid        *string "json:\"uuid,omitempty\""
	}, error)
	CreateTagCategory(arg1 TagsCreateTagCategoryJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		CreatedAt   *string "json:\"created_at,omitempty\""
		CreatedBy   *string "json:\"created_by,omitempty\""
		Description *string "json:\"description,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Reserved    *bool   "json:\"reserved,omitempty\""
		UpdatedAt   *string "json:\"updated_at,omitempty\""
		UpdatedBy   *string "json:\"updated_by,omitempty\""
		Uuid        *string "json:\"uuid,omitempty\""
	}, error)
	CreateTagValueWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	CreateTagValue(arg1 TagsCreateTagValueJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	DeleteTagCategory(contentType string, reqEditors ...RequestEditorFn) (*TagsDeleteTagCategoryResponse, error)
	DeleteTagValue(contentType string, reqEditors ...RequestEditorFn) (*TagsDeleteTagValueResponse, error)
	DeleteTagValuesBulkWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error)
	DeleteTagValuesBulk(arg1 TagsDeleteTagValuesBulkJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error)
	EditTagCategoryWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		CreatedAt   *string "json:\"created_at,omitempty\""
		CreatedBy   *string "json:\"created_by,omitempty\""
		Description *string "json:\"description,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Reserved    *bool   "json:\"reserved,omitempty\""
		UpdatedAt   *string "json:\"updated_at,omitempty\""
		UpdatedBy   *string "json:\"updated_by,omitempty\""
		Uuid        *string "json:\"uuid,omitempty\""
	}, error)
	EditTagCategory(arg1 string, arg2 TagsEditTagCategoryJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		CreatedAt   *string "json:\"created_at,omitempty\""
		CreatedBy   *string "json:\"created_by,omitempty\""
		Description *string "json:\"description,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Reserved    *bool   "json:\"reserved,omitempty\""
		UpdatedAt   *string "json:\"updated_at,omitempty\""
		UpdatedBy   *string "json:\"updated_by,omitempty\""
		Uuid        *string "json:\"uuid,omitempty\""
	}, error)
	ListAssetFilters(reqEditors ...RequestEditorFn) (*[]struct {
		Control *struct {
			ReadableRegex *string "json:\"readable_regex,omitempty\""
			Regex         *string "json:\"regex,omitempty\""
			Type          *string "json:\"type,omitempty\""
		} "json:\"control,omitempty\""
		Name         *string   "json:\"name,omitempty\""
		Operators    *[]string "json:\"operators,omitempty\""
		ReadableName *string   "json:\"readable_name,omitempty\""
	}, error)
	ListAssetTags(contentType string, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	ListTagCategories(params *TagsListTagCategoriesParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	ListTagValues(params *TagsListTagValuesParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	TagCategoryDetails(contentType string, reqEditors ...RequestEditorFn) (*struct {
		CreatedAt   *string "json:\"created_at,omitempty\""
		CreatedBy   *string "json:\"created_by,omitempty\""
		Description *string "json:\"description,omitempty\""
		Name        *string "json:\"name,omitempty\""
		Reserved    *bool   "json:\"reserved,omitempty\""
		UpdatedAt   *string "json:\"updated_at,omitempty\""
		UpdatedBy   *string "json:\"updated_by,omitempty\""
		Uuid        *string "json:\"uuid,omitempty\""
	}, error)
	TagValueDetails(contentType string, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	UpdateTagValueWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	UpdateTagValue(arg1 string, arg2 TagsUpdateTagValueJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
}

// PermissionsInterface contains all methods for the nessus  api
type PermissionsInterface interface {
	ChangeWithBody(arg1 PermissionsChangeParamsObjectType, arg2 int32, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error)
	Change(arg1 PermissionsChangeParamsObjectType, arg2 int32, arg3 PermissionsChangeJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error)
	List(arg1 PermissionsListParamsObjectType, arg2 int32, reqEditors ...RequestEditorFn) (*[]struct {
		DisplayName *string   "json:\"display_name,omitempty\""
		Id          *int      "json:\"id,omitempty\""
		Name        *string   "json:\"name,omitempty\""
		Owner       *int      "json:\"owner,omitempty\""
		Permissions *int32    "json:\"permissions,omitempty\""
		Type        *N200Type "json:\"type,omitempty\""
		Uuid        *string   "json:\"uuid,omitempty\""
	}, error)
}

// ScannerInterface contains all methods for the nessus  api
type ScannerInterface interface {
	GroupsAddScanner(arg1 int32, arg2 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	GroupsCreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
		Shared               *int    "json:\"shared,omitempty\""
		Token                *string "json:\"token,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}, error)
	GroupsCreate(arg1 ScannerGroupsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
		Shared               *int    "json:\"shared,omitempty\""
		Token                *string "json:\"token,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}, error)
	GroupsDeleteScanner(arg1 int32, arg2 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	GroupsDelete(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	GroupsDetails(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
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
		Shared               *int    "json:\"shared,omitempty\""
		Token                *string "json:\"token,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}, error)
	GroupsEditWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
		Shared               *int    "json:\"shared,omitempty\""
		Token                *string "json:\"token,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}, error)
	GroupsEdit(arg1 int32, arg2 ScannerGroupsEditJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
		Shared               *int    "json:\"shared,omitempty\""
		Token                *string "json:\"token,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}, error)
	GroupsListScanners(arg1 int32, reqEditors ...RequestEditorFn) (*[]struct {
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
		UiBuild            *int    "json:\"ui_build,omitempty\""
		UiVersion          *string "json:\"ui_version,omitempty\""
		UserPermissions    *int    "json:\"user_permissions,omitempty\""
		Uuid               *string "json:\"uuid,omitempty\""
	}, error)
	GroupsList(reqEditors ...RequestEditorFn) (*struct {
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
		Shared               *int    "json:\"shared,omitempty\""
		Token                *string "json:\"token,omitempty\""
		Type                 *string "json:\"type,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Uuid                 *string "json:\"uuid,omitempty\""
	}, error)
}

// BulkInterface contains all methods for the nessus  api
type BulkInterface interface {
	AddAgentsWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	AddAgents(arg1 string, arg2 BulkAddAgentsJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	RemoveAgentsWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	RemoveAgents(arg1 int32, arg2 BulkRemoveAgentsJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	TaskAgentGroupStatus(arg1 int32, contentType string, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	TaskAgentStatus(contentType string, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	UnlinkAgentsWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	UnlinkAgents(arg1 BulkUnlinkAgentsJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
}

// IoNetworksInterface contains all methods for the nessus  api
type IoNetworksInterface interface {
	AssetCountDetails(arg1 string, arg2 int, reqEditors ...RequestEditorFn) (*struct {
		NumAssetsNotSeen *int "json:\"numAssetsNotSeen,omitempty\""
		NumAssetstotal   *int "json:\"numAssetstotal,omitempty\""
	}, error)
}

// NetworksInterface contains all methods for the nessus  api
type NetworksInterface interface {
	AssignScannerBulkWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error)
	AssignScannerBulk(arg1 string, arg2 NetworksAssignScannerBulkJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error)
	AssignScanner(arg1 string, contentType string, reqEditors ...RequestEditorFn) (*interface{}, error)
	CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	Create(arg1 NetworksCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	Delete(contentType string, reqEditors ...RequestEditorFn) (*interface{}, error)
	Details(contentType string, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	ListAssignableScanners(contentType string, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	ListScanners(contentType string, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	List(params *NetworksListParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	UpdateWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	Update(arg1 string, arg2 NetworksUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
}

// CredentialsInterface contains all methods for the nessus  api
type CredentialsInterface interface {
	CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		Uuid *string "json:\"uuid,omitempty\""
	}, error)
	Create(arg1 CredentialsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		Uuid *string "json:\"uuid,omitempty\""
	}, error)
	Delete(contentType string, reqEditors ...RequestEditorFn) (*struct {
		Deleted *bool "json:\"deleted,omitempty\""
	}, error)
	Details(contentType string, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	FileUploadWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		Fileuploaded *string "json:\"fileuploaded,omitempty\""
	}, error)
	ListCredentialTypes(reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	List(params *CredentialsListParams, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	UpdateWithBody(arg1 string, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		Updated *bool "json:\"updated,omitempty\""
	}, error)
	Update(arg1 string, arg2 CredentialsUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		Updated *bool "json:\"updated,omitempty\""
	}, error)
}

// IoExportsComplianceInterface contains all methods for the nessus  api
type IoExportsComplianceInterface interface {
	Cancel(contentType string, reqEditors ...RequestEditorFn) (*IoExportsComplianceCancelResponse, error)
	CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		ExportUuid *string "json:\"export_uuid,omitempty\""
	}, error)
	Create(arg1 IoExportsComplianceCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		ExportUuid *string "json:\"export_uuid,omitempty\""
	}, error)
	Download(arg1 string, arg2 int32, reqEditors ...RequestEditorFn) (*[]struct {
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
	}, error)
	Status(contentType string, reqEditors ...RequestEditorFn) (*struct {
		ChunksAvailable *[]int32 "json:\"chunks_available,omitempty\""
		Status          *string  "json:\"status,omitempty\""
	}, error)
}

// VulnerabilitiesInterface contains all methods for the nessus  api
type VulnerabilitiesInterface interface {
	ImportV2WithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		JobUuid *string "json:\"job_uuid,omitempty\""
	}, error)
	ImportV2(arg1 VulnerabilitiesImportV2JSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		JobUuid *string "json:\"job_uuid,omitempty\""
	}, error)
	ImportWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error)
	Import(arg1 VulnerabilitiesImportJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error)
}

// AgentExclusionsInterface contains all methods for the nessus  api
type AgentExclusionsInterface interface {
	CreateWithBody(arg1 int32, arg2 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	Create(arg1 int32, arg2 AgentExclusionsCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	Delete(arg1 int32, arg2 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	Details(arg1 int32, arg2 int32, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	EditWithBody(arg1 int32, arg2 int32, arg3 string, body io.Reader, reqEditors ...RequestEditorFn) (*interface{}, error)
	Edit(arg1 int32, arg2 int32, arg3 AgentExclusionsEditJSONRequestBody, reqEditors ...RequestEditorFn) (*interface{}, error)
	List(arg1 int32, reqEditors ...RequestEditorFn) (*[]struct {
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
	}, error)
}

// AssetsInterface contains all methods for the nessus  api
type AssetsInterface interface {
	AssetInfo(contentType string, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	BulkDeleteWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*AssetsBulkDeleteResponse, error)
	BulkDelete(arg1 AssetsBulkDeleteJSONRequestBody, reqEditors ...RequestEditorFn) (*AssetsBulkDeleteResponse, error)
	BulkMoveWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*AssetsBulkMoveResponse, error)
	BulkMove(arg1 AssetsBulkMoveJSONRequestBody, reqEditors ...RequestEditorFn) (*AssetsBulkMoveResponse, error)
	BulkUpdateAcrWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*AssetsBulkUpdateAcrResponse, error)
	BulkUpdateAcr(arg1 AssetsBulkUpdateAcrJSONRequestBody, reqEditors ...RequestEditorFn) (*AssetsBulkUpdateAcrResponse, error)
	ImportJobInfo(contentType string, reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	ImportWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		AssetImportJobUuid *string "json:\"asset_import_job_uuid,omitempty\""
	}, error)
	Import(arg1 AssetsImportJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		AssetImportJobUuid *string "json:\"asset_import_job_uuid,omitempty\""
	}, error)
	ListAssets(reqEditors ...RequestEditorFn) (*struct {
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
	}, error)
	ListImportJobs(reqEditors ...RequestEditorFn) (*[]struct {
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
	}, error)
}

// PoliciesInterface contains all methods for the nessus  api
type PoliciesInterface interface {
	Configure(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	Copy(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	CreateWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		PolicyId   *int    "json:\"policy_id,omitempty\""
		PolicyName *string "json:\"policy_name,omitempty\""
	}, error)
	Create(arg1 PoliciesCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		PolicyId   *int    "json:\"policy_id,omitempty\""
		PolicyName *string "json:\"policy_name,omitempty\""
	}, error)
	Delete(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	Details(arg1 int32, reqEditors ...RequestEditorFn) (*struct {
		Audits      *map[string]interface{} "json:\"audits,omitempty\""
		Credentials *map[string]interface{} "json:\"credentials,omitempty\""
		Plugins     *map[string]interface{} "json:\"plugins,omitempty\""
		Scap        *map[string]interface{} "json:\"scap,omitempty\""
		Settings    *map[string]interface{} "json:\"settings,omitempty\""
		Uuid        *string                 "json:\"uuid,omitempty\""
	}, error)
	Export(arg1 int32, reqEditors ...RequestEditorFn) (*interface{}, error)
	ImportWithBody(arg1 string, body io.Reader, reqEditors ...RequestEditorFn) (*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NoTarget             *string "json:\"no_target,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		Private              *int    "json:\"private,omitempty\""
		Shared               *int    "json:\"shared,omitempty\""
		TemplateUuid         *string "json:\"template_uuid,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}, error)
	Import(arg1 PoliciesImportJSONRequestBody, reqEditors ...RequestEditorFn) (*struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NoTarget             *string "json:\"no_target,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *int    "json:\"owner_id,omitempty\""
		Private              *int    "json:\"private,omitempty\""
		Shared               *int    "json:\"shared,omitempty\""
		TemplateUuid         *string "json:\"template_uuid,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
	}, error)
	List(reqEditors ...RequestEditorFn) (*[]struct {
		CreationDate         *int    "json:\"creation_date,omitempty\""
		Description          *string "json:\"description,omitempty\""
		Id                   *int    "json:\"id,omitempty\""
		LastModificationDate *int    "json:\"last_modification_date,omitempty\""
		Name                 *string "json:\"name,omitempty\""
		NoTarget             *bool   "json:\"no_target,omitempty\""
		Owner                *string "json:\"owner,omitempty\""
		OwnerId              *string "json:\"owner_id,omitempty\""
		Shared               *int    "json:\"shared,omitempty\""
		TemplateUuid         *string "json:\"template_uuid,omitempty\""
		UserPermissions      *int    "json:\"user_permissions,omitempty\""
		Visibility           *int    "json:\"visibility,omitempty\""
	}, error)
}
