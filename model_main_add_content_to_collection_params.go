/*
Estuary API

This is the API for the Estuary application.

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MainAddContentToCollectionParams struct for MainAddContentToCollectionParams
type MainAddContentToCollectionParams struct {
	Cids []string `json:"cids,omitempty"`
	Coluuid *string `json:"coluuid,omitempty"`
	Contents []int32 `json:"contents,omitempty"`
}

// NewMainAddContentToCollectionParams instantiates a new MainAddContentToCollectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMainAddContentToCollectionParams() *MainAddContentToCollectionParams {
	this := MainAddContentToCollectionParams{}
	return &this
}

// NewMainAddContentToCollectionParamsWithDefaults instantiates a new MainAddContentToCollectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMainAddContentToCollectionParamsWithDefaults() *MainAddContentToCollectionParams {
	this := MainAddContentToCollectionParams{}
	return &this
}

// GetCids returns the Cids field value if set, zero value otherwise.
func (o *MainAddContentToCollectionParams) GetCids() []string {
	if o == nil || o.Cids == nil {
		var ret []string
		return ret
	}
	return o.Cids
}

// GetCidsOk returns a tuple with the Cids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainAddContentToCollectionParams) GetCidsOk() ([]string, bool) {
	if o == nil || o.Cids == nil {
		return nil, false
	}
	return o.Cids, true
}

// HasCids returns a boolean if a field has been set.
func (o *MainAddContentToCollectionParams) HasCids() bool {
	if o != nil && o.Cids != nil {
		return true
	}

	return false
}

// SetCids gets a reference to the given []string and assigns it to the Cids field.
func (o *MainAddContentToCollectionParams) SetCids(v []string) {
	o.Cids = v
}

// GetColuuid returns the Coluuid field value if set, zero value otherwise.
func (o *MainAddContentToCollectionParams) GetColuuid() string {
	if o == nil || o.Coluuid == nil {
		var ret string
		return ret
	}
	return *o.Coluuid
}

// GetColuuidOk returns a tuple with the Coluuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainAddContentToCollectionParams) GetColuuidOk() (*string, bool) {
	if o == nil || o.Coluuid == nil {
		return nil, false
	}
	return o.Coluuid, true
}

// HasColuuid returns a boolean if a field has been set.
func (o *MainAddContentToCollectionParams) HasColuuid() bool {
	if o != nil && o.Coluuid != nil {
		return true
	}

	return false
}

// SetColuuid gets a reference to the given string and assigns it to the Coluuid field.
func (o *MainAddContentToCollectionParams) SetColuuid(v string) {
	o.Coluuid = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *MainAddContentToCollectionParams) GetContents() []int32 {
	if o == nil || o.Contents == nil {
		var ret []int32
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainAddContentToCollectionParams) GetContentsOk() ([]int32, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *MainAddContentToCollectionParams) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given []int32 and assigns it to the Contents field.
func (o *MainAddContentToCollectionParams) SetContents(v []int32) {
	o.Contents = v
}

func (o MainAddContentToCollectionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cids != nil {
		toSerialize["cids"] = o.Cids
	}
	if o.Coluuid != nil {
		toSerialize["coluuid"] = o.Coluuid
	}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}
	return json.Marshal(toSerialize)
}

type NullableMainAddContentToCollectionParams struct {
	value *MainAddContentToCollectionParams
	isSet bool
}

func (v NullableMainAddContentToCollectionParams) Get() *MainAddContentToCollectionParams {
	return v.value
}

func (v *NullableMainAddContentToCollectionParams) Set(val *MainAddContentToCollectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMainAddContentToCollectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMainAddContentToCollectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMainAddContentToCollectionParams(val *MainAddContentToCollectionParams) *NullableMainAddContentToCollectionParams {
	return &NullableMainAddContentToCollectionParams{value: val, isSet: true}
}

func (v NullableMainAddContentToCollectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMainAddContentToCollectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


