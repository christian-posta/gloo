package azure

import (
	"errors"
	"regexp"

	"github.com/gogo/protobuf/types"
	"github.com/solo-io/gloo-api/pkg/api/types/v1"
	"github.com/solo-io/gloo/pkg/protoutil"

	multierror "github.com/hashicorp/go-multierror"
)

type UpstreamSpec struct {
	FunctionAppName string `json:"function_app_name"`
	SecretRef       string `json:"secret_ref"`
}

func EncodeUpstreamSpec(spec UpstreamSpec) *types.Struct {
	v1Spec, err := protoutil.MarshalStruct(spec)
	if err != nil {
		panic(err)
	}
	return v1Spec
}

func DecodeUpstreamSpec(generic v1.UpstreamSpec) (*UpstreamSpec, error) {
	s := new(UpstreamSpec)
	if err := protoutil.UnmarshalStruct(generic, s); err != nil {
		return nil, err
	}
	return s, s.Validate()
}

func (s *UpstreamSpec) Validate() error {
	if IsValidFunctionAppName(s.FunctionAppName) {
		return nil
	} else {
		return errors.New("invalid function app name")
	}
}

type FunctionSpec struct {
	FunctionName string `json:"function_name"`
	AuthLevel    string `json:"auth_level"`
}

func EncodeFunctionSpec(spec FunctionSpec) *types.Struct {
	v1Spec, err := protoutil.MarshalStruct(spec)
	if err != nil {
		panic(err)
	}
	return v1Spec
}

func DecodeFunctionSpec(generic v1.FunctionSpec) (*FunctionSpec, error) {
	s := new(FunctionSpec)
	if err := protoutil.UnmarshalStruct(generic, s); err != nil {
		return nil, err
	}
	return s, s.Validate()
}

func (s *FunctionSpec) Validate() error {
	var result error

	if !IsValidFunctionName(s.FunctionName) {
		result = multierror.Append(result, errors.New("invalid function name"))
	}
	if !IsValidAuthLevel(s.AuthLevel) {
		result = multierror.Append(result, errors.New("invalid authentication level"))
	}

	return result
}

func IsValidFunctionAppName(functionAppName string) bool {
	// Valid characters are `a-z`, `0-9`, and `-`.
	return regexp.MustCompile("^[[:alpha:][:digit:]-]+$").MatchString(functionAppName)
}

func IsValidFunctionName(functionName string) bool {
	// The name must be unique within a Function App. It must start with a letter and can contain
	// letters, numbers (0-9), dashes ("-"), and underscores ("_").
	return regexp.MustCompile("^[[:alpha:]][[:alpha:][:digit:]_-]*$").MatchString(functionName)
}

func IsValidAuthLevel(authLevel string) bool {
	// The authorization level can be one of the following values:
	// * "anonymous" - No API key is required.
	// * "function" - A function-specific API key is required. This is the default value if none is
	//                provided.
	// * "admin" - The master key is required.
	for _, str := range []string{"anonymous", "function", "admin"} {
		if authLevel == str {
			return true
		}
	}
	return false
}
