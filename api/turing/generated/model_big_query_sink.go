/*
 * Turing Minimal Openapi Spec for SDK
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// BigQuerySink struct for BigQuerySink
type BigQuerySink struct {
	Type string `json:"type"`
	Columns []string `json:"columns,omitempty"`
	SaveMode SaveMode `json:"save_mode"`
	BqConfig BigQuerySinkConfig `json:"bq_config"`
}

// NewBigQuerySink instantiates a new BigQuerySink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBigQuerySink(type_ string, saveMode SaveMode, bqConfig BigQuerySinkConfig) *BigQuerySink {
	this := BigQuerySink{}
	this.Type = type_
	this.SaveMode = saveMode
	this.BqConfig = bqConfig
	return &this
}

// NewBigQuerySinkWithDefaults instantiates a new BigQuerySink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBigQuerySinkWithDefaults() *BigQuerySink {
	this := BigQuerySink{}
	var type_ string = "BQ"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *BigQuerySink) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BigQuerySink) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BigQuerySink) SetType(v string) {
	o.Type = v
}

// GetColumns returns the Columns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BigQuerySink) GetColumns() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BigQuerySink) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *BigQuerySink) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *BigQuerySink) SetColumns(v []string) {
	o.Columns = v
}

// GetSaveMode returns the SaveMode field value
func (o *BigQuerySink) GetSaveMode() SaveMode {
	if o == nil {
		var ret SaveMode
		return ret
	}

	return o.SaveMode
}

// GetSaveModeOk returns a tuple with the SaveMode field value
// and a boolean to check if the value has been set.
func (o *BigQuerySink) GetSaveModeOk() (*SaveMode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SaveMode, true
}

// SetSaveMode sets field value
func (o *BigQuerySink) SetSaveMode(v SaveMode) {
	o.SaveMode = v
}

// GetBqConfig returns the BqConfig field value
func (o *BigQuerySink) GetBqConfig() BigQuerySinkConfig {
	if o == nil {
		var ret BigQuerySinkConfig
		return ret
	}

	return o.BqConfig
}

// GetBqConfigOk returns a tuple with the BqConfig field value
// and a boolean to check if the value has been set.
func (o *BigQuerySink) GetBqConfigOk() (*BigQuerySinkConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BqConfig, true
}

// SetBqConfig sets field value
func (o *BigQuerySink) SetBqConfig(v BigQuerySinkConfig) {
	o.BqConfig = v
}

func (o BigQuerySink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if true {
		toSerialize["save_mode"] = o.SaveMode
	}
	if true {
		toSerialize["bq_config"] = o.BqConfig
	}
	return json.Marshal(toSerialize)
}

type NullableBigQuerySink struct {
	value *BigQuerySink
	isSet bool
}

func (v NullableBigQuerySink) Get() *BigQuerySink {
	return v.value
}

func (v *NullableBigQuerySink) Set(val *BigQuerySink) {
	v.value = val
	v.isSet = true
}

func (v NullableBigQuerySink) IsSet() bool {
	return v.isSet
}

func (v *NullableBigQuerySink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBigQuerySink(val *BigQuerySink) *NullableBigQuerySink {
	return &NullableBigQuerySink{value: val, isSet: true}
}

func (v NullableBigQuerySink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBigQuerySink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


