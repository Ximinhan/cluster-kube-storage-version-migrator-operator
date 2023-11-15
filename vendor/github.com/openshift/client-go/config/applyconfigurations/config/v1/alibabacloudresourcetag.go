// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// AlibabaCloudResourceTagApplyConfiguration represents an declarative configuration of the AlibabaCloudResourceTag type for use
// with apply.
type AlibabaCloudResourceTagApplyConfiguration struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// AlibabaCloudResourceTagApplyConfiguration constructs an declarative configuration of the AlibabaCloudResourceTag type for use with
// apply.
func AlibabaCloudResourceTag() *AlibabaCloudResourceTagApplyConfiguration {
	return &AlibabaCloudResourceTagApplyConfiguration{}
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *AlibabaCloudResourceTagApplyConfiguration) WithKey(value string) *AlibabaCloudResourceTagApplyConfiguration {
	b.Key = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *AlibabaCloudResourceTagApplyConfiguration) WithValue(value string) *AlibabaCloudResourceTagApplyConfiguration {
	b.Value = &value
	return b
}
