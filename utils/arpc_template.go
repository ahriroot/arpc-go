package utils

import (
	"fmt"
	"strings"
)

func TypeStr2GoType(type_str string) string {
	switch type_str {
	case "integer":
		return "int"
	case "int":
		return "int"
	case "int4":
		return "int4"
	case "int8":
		return "int8"
	case "int16":
		return "int16"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "string":
		return "string"
	case "bool":
		return "bool"
	case "float":
		return "float32"
	case "float32":
		return "float32"
	case "float64":
		return "float64"
	case "double":
		return "float64"
	case "json":
		return "map[string]interface{}"
	case "object":
		return "map[string]interface{}"
	default:
		return type_str
	}
}

const (
	T_INTERFACE = `
type %s interface {
%s
}
`

	T_STRUCT = `
type %s struct {
	conn *arpc_client.ArpcConn
}
`

	T_CLIENT = `
func (c *%s) %s(request *%s) (*%s, error) {
	req_bytes, err := request.Serialize()
	if err != nil {
		return nil, err
	}
	data, err := arpc_client.Handle("%s.%d", req_bytes, c.conn)
	if err != nil {
		return nil, err
	}
	response := &%s{}
	err = response.Deserialize(data)
	if err != nil {
		return nil, err
	}
	return response, nil
}
`

	T_SERVER = `    s.Register("%s.%d", func(request []byte, _ arpc_client.ArpcConn) ([]byte, error) {
		req := &%s{}
		err := req.Deserialize(request)
		if err != nil {
			return nil, err
		}
		response, err := i.%s(req)
		if err != nil {
			return nil, err
		}
		return response.Serialize()
	})`

	T_NEW_CLIENT = `
func %s(c *arpc_client.ArpcConn) %s {
	return &%s{c}
}
`
)

func GenerateParamStruct(name string, params []Param) string {
	var params_list = make([]string, 0)
	var field_list = make([]string, 0)
	var st = fmt.Sprintf("type %s struct {", name)
	var tmp = make(map[string]string)
	var max_field_length = 0
	var max_type_length = 0
	for _, param := range params {
		snake := Snake(param.Name)
		tmp[param.Name] = TypeStr2GoType(param.Type)
		field_length := len(param.Name)
		if field_length > max_field_length {
			max_field_length = field_length
		}
		type_length := len(tmp[param.Name])
		if type_length > max_type_length {
			max_type_length = type_length
		}
		params_list = append(params_list, fmt.Sprintf("%s %s", snake, tmp[param.Name]))
		field_list = append(field_list, fmt.Sprintf("b.%s = %s", param.Name, snake))
	}
	format := fmt.Sprintf("\n    %%-%ds %%-%ds `json:\"%%s\"`", max_field_length, max_type_length)
	for k, v := range tmp {
		st += fmt.Sprintf(format, k, v, Snake(k))
	}
	st += "\n}"

	var new_str = `

func (b *%s) New(%s) {
	%s
}`
	st += fmt.Sprintf(new_str, name, strings.Join(params_list, ", "), strings.Join(field_list, "\n    "))

	var serialize_str = `

func (b *%s) Serialize() ([]byte, error) {
	return json.Marshal(b)
}`
	st += fmt.Sprintf(serialize_str, name)

	var deserialize_str = `

func (b *%s) Deserialize(data []byte) error {
	return json.Unmarshal(data, b)
}`
	st += fmt.Sprintf(deserialize_str, name)
	return st
}

func GenerateProcedureStruct(name string, procedure []Procedures) string {
	var package_id = RandomString(6)

	var strs_interface []string
	var strs_client []string
	var strs_server []string

	for _, p := range procedure {
		strs_interface = append(strs_interface, fmt.Sprintf("    %s(*%s) (*%s, error)", p.Name, p.Request, p.Response))
		strs_client = append(strs_client, fmt.Sprintf(T_CLIENT, Snake(name), p.Name, p.Request, p.Response, package_id, p.Index, p.Response))
		strs_server = append(strs_server, fmt.Sprintf(T_SERVER, package_id, p.Index, p.Request, p.Name))
	}
	st := fmt.Sprintf(T_STRUCT, Snake(name))
	st += fmt.Sprintf(T_INTERFACE, name, strings.Join(strs_interface, "\n"))
	st += strings.Join(strs_client, "\n")

	var wapper = `
func %s(s *arpc_server.Server, i %s) {
%s
}
`
	st += fmt.Sprintf(wapper, "Register", name, strings.Join(strs_server, "\n"))
	return st
}

func GenerateNewClient(name string) string {
	return fmt.Sprintf(T_NEW_CLIENT, "New"+name, name, Snake(name))
}
