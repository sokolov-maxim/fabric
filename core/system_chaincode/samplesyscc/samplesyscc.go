/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package samplesyscc

import (
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SampleSysCC example simple Chaincode implementation
type SampleSysCC struct {
}

// Init initializes the sample system chaincode by storing the key and value
// arguments passed in as parameters
func (t *SampleSysCC) Init(stub shim.ChaincodeStubInterface) ([]byte, error) {
	//as system chaincodes do not take part in consensus and are part of the system,
	//best practice to do nothing (or very little) in Init.

	return nil, nil
}

// Invoke gets the supplied key and if it exists, updates the key with the newly
// supplied value.
func (t *SampleSysCC) Invoke(stub shim.ChaincodeStubInterface) ([]byte, error) {
	f, args := stub.GetFunctionAndParameters()

	switch f {
	case "putval":
		if len(args) != 2 {
			return nil, errors.New("need 2 args (key and a value)")
		}

		// Initialize the chaincode
		key := args[0]
		val := args[1]

		_, err := stub.GetState(key)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get val for " + key + "\"}"
			return nil, errors.New(jsonResp)
		}

		// Write the state to the ledger
		err = stub.PutState(key, []byte(val))
		return nil, err
	case "getval":
		var err error

		if len(args) != 1 {
			return nil, errors.New("Incorrect number of arguments. Expecting key to query")
		}

		key := args[0]

		// Get the state from the ledger
		valbytes, err := stub.GetState(key)
		if err != nil {
			jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
			return nil, errors.New(jsonResp)
		}

		if valbytes == nil {
			jsonResp := "{\"Error\":\"Nil val for " + key + "\"}"
			return nil, errors.New(jsonResp)
		}

		return valbytes, nil
	default:
		jsonResp := "{\"Error\":\"Unknown functon " + f + "\"}"
		return nil, errors.New(jsonResp)
	}
}
