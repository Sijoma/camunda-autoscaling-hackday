/*
Cluster Topology Management API

API for managing cluster membership and partition distribution.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zbmgmt

import (
	"encoding/json"
	"time"
)

// checks if the BrokerState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrokerState{}

// BrokerState State of a broker
type BrokerState struct {
	// The ID of a broker, starting from 0
	Id    *int32           `json:"id,omitempty"`
	State *BrokerStateCode `json:"state,omitempty"`
	// The version of the broker state
	Version *int64 `json:"version,omitempty"`
	// The time when the broker state was last updated
	LastUpdatedAt *time.Time       `json:"lastUpdatedAt,omitempty"`
	Partitions    []PartitionState `json:"partitions,omitempty"`
}

// NewBrokerState instantiates a new BrokerState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerState() *BrokerState {
	this := BrokerState{}
	return &this
}

// NewBrokerStateWithDefaults instantiates a new BrokerState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerStateWithDefaults() *BrokerState {
	this := BrokerState{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BrokerState) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerState) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BrokerState) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BrokerState) SetId(v int32) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BrokerState) GetState() BrokerStateCode {
	if o == nil || IsNil(o.State) {
		var ret BrokerStateCode
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerState) GetStateOk() (*BrokerStateCode, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BrokerState) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given BrokerStateCode and assigns it to the State field.
func (o *BrokerState) SetState(v BrokerStateCode) {
	o.State = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BrokerState) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerState) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BrokerState) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *BrokerState) SetVersion(v int64) {
	o.Version = &v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *BrokerState) GetLastUpdatedAt() time.Time {
	if o == nil || IsNil(o.LastUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerState) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedAt) {
		return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *BrokerState) HasLastUpdatedAt() bool {
	if o != nil && !IsNil(o.LastUpdatedAt) {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given time.Time and assigns it to the LastUpdatedAt field.
func (o *BrokerState) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = &v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *BrokerState) GetPartitions() []PartitionState {
	if o == nil || IsNil(o.Partitions) {
		var ret []PartitionState
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerState) GetPartitionsOk() ([]PartitionState, bool) {
	if o == nil || IsNil(o.Partitions) {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *BrokerState) HasPartitions() bool {
	if o != nil && !IsNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []PartitionState and assigns it to the Partitions field.
func (o *BrokerState) SetPartitions(v []PartitionState) {
	o.Partitions = v
}

func (o BrokerState) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrokerState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.LastUpdatedAt) {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt
	}
	if !IsNil(o.Partitions) {
		toSerialize["partitions"] = o.Partitions
	}
	return toSerialize, nil
}

type NullableBrokerState struct {
	value *BrokerState
	isSet bool
}

func (v NullableBrokerState) Get() *BrokerState {
	return v.value
}

func (v *NullableBrokerState) Set(val *BrokerState) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerState) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerState(val *BrokerState) *NullableBrokerState {
	return &NullableBrokerState{value: val, isSet: true}
}

func (v NullableBrokerState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrokerState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
