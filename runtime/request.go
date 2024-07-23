package runtime

import (
	"errors"
	"hash/fnv"

	"encoding/json"
	"net/http"

	"github.com/iancoleman/strcase"
	"github.com/wundergraph/graphql-go-tools/pkg/execution"
)

type GraphqlRequest struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
	OperationName string                 `json:"operationName"`
}

func parseRequest(r *http.Request) (*GraphqlRequest, error) {
	var req GraphqlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	if req.Variables == nil {
		req.Variables = make(map[string]interface{})
	}
	return &req, nil
}

func convertVariables(vars map[string]interface{}) execution.Variables {
	result := make(execution.Variables)
	for k, v := range vars {
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			// Handle error appropriately
			continue
		}
		// Use a hash of the key as uint64
		h := fnv.New64a()
		h.Write([]byte(k))
		key := h.Sum64()
		result[key] = jsonBytes
	}
	return result
}

// MarshalRequest marshals graphql request arguments to gRPC request message
func MarshalRequest(args, v interface{}, isCamel bool) error {
	if args == nil {
		return errors.New("Resolved params should be non-nil")
	}
	m, ok := args.(map[string]interface{}) // graphql.ResolveParams or nested object
	if !ok {
		return errors.New("Failed to type conversion of map[string]interface{}")
	}
	if isCamel {
		m = toLowerCaseKeys(m)
	}
	buf, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, &v)
}

// Convert to lower case keyname string
func toLowerCaseKeys(args map[string]interface{}) map[string]interface{} {
	lc := make(map[string]interface{})
	for k, v := range args {
		lc[strcase.ToSnake(k)] = marshal(v)
	}
	return lc
}

// marshals interface recursively
func marshal(v interface{}) interface{} {
	switch t := v.(type) {
	case map[string]interface{}:
		return toLowerCaseKeys(t)
	case []interface{}:
		ret := make([]interface{}, len(t))
		for i, si := range t {
			ret[i] = marshal(si)
		}
		return ret
	default:
		return t
	}
}
