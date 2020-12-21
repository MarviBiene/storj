// +build js,wasm
// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/zeebo/errs"

	console "storj.io/storj/satellite/console/consolewasm"
)

func main() {
	js.Global().Set("generateAccessGrant", generateAccessGrant())
	js.Global().Set("setAPIKeyPermission", setAPIKeyPermission())
	js.Global().Set("newPermission", newPermission())
	<-make(chan bool)
}

// generateAccessGrant creates a new access grant with the provided api key and encryption passphrase.
func generateAccessGrant() js.Func {
	return js.FuncOf(responseHandler(func(this js.Value, args []js.Value) (interface{}, error) {
		if len(args) < 4 {
			return nil, errs.New("not enough arguments. Need 4, but only %d supplied. The order of arguments are: satellite Node URL, API key, encryption passphrase, and project ID.", len(args))
		}
		satelliteNodeURL := args[0].String()
		apiKey := args[1].String()
		encryptionPassphrase := args[2].String()
		projectSalt := args[3].String()

		access, err := console.GenAccessGrant(satelliteNodeURL,
			apiKey,
			encryptionPassphrase,
			projectSalt,
		)
		if err != nil {
			return nil, err
		}

		return access, nil
	}))
}

// setAPIKeyPermission creates a new api key with specific permissions.
func setAPIKeyPermission() js.Func {
	return js.FuncOf(responseHandler(func(this js.Value, args []js.Value) (interface{}, error) {
		if len(args) < 3 {
			return nil, errs.New("not enough arguments. Need 3, but only %d supplied. The order of arguments are: API key, bucket names, and permission object.", len(args))
		}
		apiKey := args[0].String()

		// convert array of bucket names to go []string type
		buckets := args[1]
		if ok := buckets.InstanceOf(js.Global().Get("Array")); !ok {
			return nil, errs.New("invalid data type. Expect Array, Got %s", buckets.Type().String())
		}
		bucketNames, err := parseArrayOfStrings(buckets)
		if err != nil {
			return nil, err
		}

		// convert js permission to go permission type
		permissionJS := args[2]
		if permissionJS.Type() != js.TypeObject {
			return nil, errs.New("invalid argument type. Expect %s, Got %s", js.TypeObject.String(), permissionJS.Type().String())
		}
		permission, err := parsePermission(permissionJS)
		if err != nil {
			return nil, err
		}

		restrictedKey, err := console.SetPermission(apiKey, bucketNames, permission)
		if err != nil {
			return nil, err
		}

		return restrictedKey.Serialize(), nil
	}))
}

// newPermission creates a new permission object.
func newPermission() js.Func {
	return js.FuncOf(responseHandler(func(this js.Value, args []js.Value) (interface{}, error) {
		p, err := json.Marshal(console.Permission{})
		if err != nil {
			return nil, err
		}

		var jsObj map[string]interface{}
		if err = json.Unmarshal(p, &jsObj); err != nil {
			return nil, err
		}
		return jsObj, nil
	}))
}

func parsePermission(arg js.Value) (console.Permission, error) {
	var permission console.Permission

	// convert javascript object to a json string
	jsJSON := js.Global().Get("JSON")
	p := jsJSON.Call("stringify", arg)

	if err := json.Unmarshal([]byte(p.String()), &permission); err != nil {
		return permission, err
	}

	return permission, nil
}

func parseArrayOfStrings(arg js.Value) ([]string, error) {
	data := make([]string, arg.Length())
	for i := 0; i < arg.Length(); i++ {
		data[i] = arg.Index(i).String()
	}

	return data, nil
}

type result struct {
	value interface{}
	err   error
}

func (r result) ToJS() map[string]interface{} {
	var errMsg string
	if r.err != nil {
		errMsg = r.err.Error()
	}
	return map[string]interface{}{
		"value": js.ValueOf(r.value),
		"error": errMsg,
	}
}

func responseHandler(fn func(this js.Value, args []js.Value) (value interface{}, err error)) func(js.Value, []js.Value) interface{} {
	return func(this js.Value, args []js.Value) interface{} {
		value, err := fn(this, args)
		return result{value, err}.ToJS()
	}
}
