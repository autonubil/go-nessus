## Generate Client

https://github.com/OpenAPITools/openapi-generator#1---installation

download httpshttps://developer.tenable.com/reference#read-the-docs

```
go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

remove "`scanner_id`" paramter from `/scanners/null/agent-groups/{group_id}/agents/_bulk/add`, `/scanners/null/agent-groups/{group_id}/agents/_bulk/directive`, `/scanners/null/agent-groups/{group_id}/agents/_bulk/remove`, `/scanners/null/agent-groups/{group_id}/agents/_bulk/{task_uuid}`, `/scanners/null/agents/_bulk/directive`, `/scanners/null/agents/_bulk/unlink`, `/scanners/null/agents/_bulk/`

rm -rf ./nessus-generated/\* ; oapi-codegen -generate types -package api ./TIO-API-Vulnerability-Management.json > ./api/types.go ; oapi-codegen -generate client -package api ./TIO-API-Vulnerability-Management.json > ./api/client.go

```


## Generate implementation

replace "HostFqdn *string `json:"host_fqdn,omitempty"`" with "HostFqdn2 *string `json:"host_fqdn,omitempty"`  

use "test" classes in `nessus_gen_code_test.go`

Fix "List Results" (described as array, but acutally returned as object with array as value to propery named alike the type (folders: []folder))

```
# curl -LO https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/bin/utils/openapi-generator-cli.sh
# chmod +x openapi-generator-cli.sh
# rm -rf ./wazuh-generated/\* ; ./openapi-generator-cli.sh generate -g go -i https://raw.githubusercontent.com/wazuh/wazuh/v4.2.4/api/api/spec/spec.yaml --http-user-agent "go-wazuh" -o ./# wazuh-generated --package-name wazuh --api-package api --model-package models --skip-validate-spec
```
