package api

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

type Function struct {
	Controller string
	Name       string
	CallName   string
	Method     reflect.Method
	Doc        string
	// ResultType    reflect.Type
	ResultTypeDef string
}

func NewFunction(controller string, name string, method reflect.Method) *Function {
	callName := name

	f := &Function{
		Controller: controller,
		Name:       name,
		CallName:   callName,
		Method:     method,
	}

	x := f.Method.Type
	t := x.Out(0)

	if okFld, ok := t.Elem().FieldByName("JSON200"); ok {
		y := okFld.Type.Elem()
		if y.Kind() == reflect.Struct {
			if dataFld, ok := y.FieldByName("Data"); ok {
				t = dataFld.Type

			} else {
				t = okFld.Type
			}
		} else {
			t = okFld.Type
		}
	}
	// re := regexp.MustCompile(`"(json|yaml).*?\\""`) re.ReplaceAllString(t.String(), "")

	st := strings.ReplaceAll(t.String(), "nessus.", "")

	f.ResultTypeDef = st

	return f
}

func (f *Function) Declaration() {
	fmt.Printf("  %s(", f.CallName)
	x := f.Method.Type
	numIn := x.NumIn()
	for i := 0; i < numIn; i++ {
		inV := x.In(i)
		name := fmt.Sprintf("arg%d", i)

		if strings.HasSuffix(inV.String(), "Params") {
			name = "params"
		} else if inV.String() == "context.Context" {
			name = "ctx"
			continue
		} else if inV.String() == "io.Reader" {
			name = "body"
		} else if strings.HasPrefix(inV.String(), "nessus.") {
			name = strings.ToLower(string(inV.String()[6])) + inV.String()[7:]
		} else if inV.String() == "string" && i == numIn-2 {
			name = "contentType"
		}
		if strings.HasPrefix(name, "policyId") {
			name = "policyID" + name[8:]
		} else if strings.HasPrefix(name, "scanId") {
			name = "scanID" + name[6:]
		}

		if i < numIn-1 {
			fmt.Printf("%s %s", name, strings.ReplaceAll(inV.String(), "wazuh.", ""))
			fmt.Printf(", ")
		} else {
			fmt.Printf("reqEditors ...RequestEditorFn")
		}
	}
	fmt.Printf(")")
	fmt.Printf(" (%s, error)", f.ResultTypeDef)
}

func (f *Function) Call() string {
	var sb strings.Builder
	x := f.Method.Type
	numIn := x.NumIn()
	for i := 0; i < numIn; i++ {
		inV := x.In(i)
		name := fmt.Sprintf("arg%d", i)
		if strings.HasSuffix(inV.String(), "Params") {
			name = "params"
		} else if inV.String() == "context.Context" {
			name = "c.ClientInterface.(*Client).ctx"
		} else if inV.String() == "io.Reader" {
			name = "body"
		} else if strings.HasPrefix(inV.String(), "wazuh.") {
			name = strings.ToLower(string(inV.String()[6])) + inV.String()[7:]
		} else if inV.String() == "string" && i == numIn-2 {
			name = "contentType"
		}

		if i == numIn-1 {
			name = "reqEditors"
		} else if strings.HasPrefix(name, "policyId") {
			name = "policyID" + name[8:]
		} else if strings.HasPrefix(name, "scanId") {
			name = "scanID" + name[6:]
		}

		sb.WriteString(name)
		if i < numIn-1 {
			sb.WriteString(", ")
		} else {
			sb.WriteString("...")
		}
	}
	return sb.String()
}

type FunctionList []*Function
type Controllers map[string]FunctionList

func TestGenCode(t *testing.T) {
	ct := reflect.TypeOf((*ClientWithResponsesInterface)(nil)).Elem()
	mc := ct.NumMethod()
	r := regexp.MustCompile(`^(IoV1|IoV2|Assets|Agents|Vulnerabilities|AuditLog|IoExportsCompliance|Credentials|Editor|Exclusions|IoFilters|Folders|Networks|IoNetworks|Permissions|IoPlugins|Policies|Scanners|Scanner|IoScanner|Bulk|IoAgent|AgentGroup|AgentConfig|AgentExclusions|Scans|IoScans|Tags|TargetGroups|ExportsVulns|Workbenches)(.*)WithResponse$`)
	controllers := Controllers{}
	for mi := 0; mi < mc; mi++ {
		m := ct.Method(mi)
		sm := r.FindStringSubmatch(m.Name)
		if len(sm) < 2 {
			continue
		}
		var controllerFunctions FunctionList
		var ok bool
		if controllerFunctions, ok = controllers[sm[1]]; !ok {
			controllerFunctions = make(FunctionList, 0)
			controllers[sm[1]] = controllerFunctions
		}
		controllers[sm[1]] = append(controllerFunctions, NewFunction(sm[1], sm[2], m))
	}
	for controllerName, controller := range controllers {
		fmt.Printf("// %sInterface contains all methods for the nessus controller api\ntype %sInterface interface {\n", controllerName, controllerName)
		for _, function := range controller {
			function.Declaration()
			fmt.Printf("\n")
		}
		fmt.Printf("}\n\n")
	}

	for controllerName, controller := range controllers {
		fmt.Printf("// %s implementation of the %s interface \ntype %s struct {\n	*ClientWithResponses\n}\n\n", controllerName, controllerName, controllerName)

		for _, function := range controller {
			fmt.Printf("// %s calls the %s controller´s function\nfunc (c *%s) ", function.CallName, controllerName, controllerName)
			function.Declaration()
			fmt.Printf("{\n")
			fmt.Printf(`	if c.ClientInterface.(*Client).token == "" {
		err := c.Authenticate()
		if err != nil {
			return nil, err
		}
	}
	r, err := c.evaluateResponse(c.ClientWithResponses.%s%sWithResponse(%s))
	if err != nil {
		return nil, err
	}
	// convert to %s
	if i, ok := r.(%s); ok {
		return i, nil
	}

    // cannot convert, return nil
	return nil, nil
}

`, function.Controller, function.Name, function.Call(), function.ResultTypeDef, function.ResultTypeDef)
		}
	}

	fmt.Println("/*")
	for controllerName := range controllers {
		fmt.Printf("	%s %sInterface\n", controllerName, controllerName)
	}
	fmt.Println()
	for controllerName := range controllers {
		fmt.Printf("	%s:  &%s{clientWithResponses},\n", controllerName, controllerName)
	}
	fmt.Println()
	fmt.Println("*/")
}
