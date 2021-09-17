//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package additionalpropsgroup

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

type CatAPTrue struct {
	PetAPTrue
	Friendly *bool `json:"friendly,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type CatAPTrue.
func (c CatAPTrue) MarshalJSON() ([]byte, error) {
	objectMap := c.PetAPTrue.marshalInternal()
	populate(objectMap, "friendly", c.Friendly)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CatAPTrue.
func (c *CatAPTrue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "friendly":
			err = unpopulate(val, &c.Friendly)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return c.PetAPTrue.unmarshalInternal(rawMsg)
}

// Implements the error and azcore.HTTPResponse interfaces.
type Error struct {
	raw     string
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
// The contents of the error text are not contractual and subject to change.
func (e Error) Error() string {
	return e.raw
}

type PetAPInProperties struct {
	// REQUIRED
	ID *int32 `json:"id,omitempty"`

	// Dictionary of
	AdditionalProperties map[string]*float32 `json:"additionalProperties,omitempty"`
	Name                 *string             `json:"name,omitempty"`

	// READ-ONLY
	Status *bool `json:"status,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type PetAPInProperties.
func (p PetAPInProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalProperties", p.AdditionalProperties)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "status", p.Status)
	return json.Marshal(objectMap)
}

type PetAPInPropertiesWithAPString struct {
	// REQUIRED
	ID *int32 `json:"id,omitempty"`

	// REQUIRED
	ODataLocation *string `json:"@odata.location,omitempty"`

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]*string

	// Dictionary of
	AdditionalProperties1 map[string]*float32 `json:"additionalProperties,omitempty"`
	Name                  *string             `json:"name,omitempty"`

	// READ-ONLY
	Status *bool `json:"status,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type PetAPInPropertiesWithAPString.
func (p PetAPInPropertiesWithAPString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalProperties", p.AdditionalProperties1)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "@odata.location", p.ODataLocation)
	populate(objectMap, "status", p.Status)
	if p.AdditionalProperties != nil {
		for key, val := range p.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetAPInPropertiesWithAPString.
func (p *PetAPInPropertiesWithAPString) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalProperties":
			err = unpopulate(val, &p.AdditionalProperties1)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &p.Name)
			delete(rawMsg, key)
		case "@odata.location":
			err = unpopulate(val, &p.ODataLocation)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &p.Status)
			delete(rawMsg, key)
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = map[string]*string{}
			}
			if val != nil {
				var aux string
				err = json.Unmarshal(val, &aux)
				p.AdditionalProperties[key] = &aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type PetAPObject struct {
	// REQUIRED
	ID *int32 `json:"id,omitempty"`

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]interface{}
	Name                 *string `json:"name,omitempty"`

	// READ-ONLY
	Status *bool `json:"status,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type PetAPObject.
func (p PetAPObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "status", p.Status)
	if p.AdditionalProperties != nil {
		for key, val := range p.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetAPObject.
func (p *PetAPObject) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &p.Name)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &p.Status)
			delete(rawMsg, key)
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				p.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type PetAPString struct {
	// REQUIRED
	ID *int32 `json:"id,omitempty"`

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]*string
	Name                 *string `json:"name,omitempty"`

	// READ-ONLY
	Status *bool `json:"status,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type PetAPString.
func (p PetAPString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "status", p.Status)
	if p.AdditionalProperties != nil {
		for key, val := range p.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetAPString.
func (p *PetAPString) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &p.Name)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &p.Status)
			delete(rawMsg, key)
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = map[string]*string{}
			}
			if val != nil {
				var aux string
				err = json.Unmarshal(val, &aux)
				p.AdditionalProperties[key] = &aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type PetAPTrue struct {
	// REQUIRED
	ID *int32 `json:"id,omitempty"`

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]interface{}
	Name                 *string `json:"name,omitempty"`

	// READ-ONLY
	Status *bool `json:"status,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type PetAPTrue.
func (p PetAPTrue) MarshalJSON() ([]byte, error) {
	objectMap := p.marshalInternal()
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetAPTrue.
func (p *PetAPTrue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return p.unmarshalInternal(rawMsg)
}

func (p PetAPTrue) marshalInternal() map[string]interface{} {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", p.ID)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "status", p.Status)
	if p.AdditionalProperties != nil {
		for key, val := range p.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return objectMap
}

func (p *PetAPTrue) unmarshalInternal(rawMsg map[string]json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, &p.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &p.Name)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, &p.Status)
			delete(rawMsg, key)
		default:
			if p.AdditionalProperties == nil {
				p.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				p.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// PetsCreateAPInPropertiesOptions contains the optional parameters for the Pets.CreateAPInProperties method.
type PetsCreateAPInPropertiesOptions struct {
	// placeholder for future optional parameters
}

// PetsCreateAPInPropertiesWithAPStringOptions contains the optional parameters for the Pets.CreateAPInPropertiesWithAPString method.
type PetsCreateAPInPropertiesWithAPStringOptions struct {
	// placeholder for future optional parameters
}

// PetsCreateAPObjectOptions contains the optional parameters for the Pets.CreateAPObject method.
type PetsCreateAPObjectOptions struct {
	// placeholder for future optional parameters
}

// PetsCreateAPStringOptions contains the optional parameters for the Pets.CreateAPString method.
type PetsCreateAPStringOptions struct {
	// placeholder for future optional parameters
}

// PetsCreateAPTrueOptions contains the optional parameters for the Pets.CreateAPTrue method.
type PetsCreateAPTrueOptions struct {
	// placeholder for future optional parameters
}

// PetsCreateCatAPTrueOptions contains the optional parameters for the Pets.CreateCatAPTrue method.
type PetsCreateCatAPTrueOptions struct {
	// placeholder for future optional parameters
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
