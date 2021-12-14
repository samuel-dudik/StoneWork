// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/golang/protobuf/proto"
	. "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"
	"go.pantheon.tech/stonework/plugins/abx/abxidx"
	"go.pantheon.tech/stonework/proto/abx"
)

////////// type-safe key-value pair with metadata //////////

type ABXKVWithMetadata struct {
	Key      string
	Value    *vpp_abx.ABX
	Metadata *abxidx.ABXMetadata
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type ABXDescriptor struct {
	Name                 string
	KeySelector          KeySelector
	ValueTypeName        string
	KeyLabel             func(key string) string
	ValueComparator      func(key string, oldValue, newValue *vpp_abx.ABX) bool
	NBKeyPrefix          string
	WithMetadata         bool
	MetadataMapFactory   MetadataMapFactory
	Validate             func(key string, value *vpp_abx.ABX) error
	Create               func(key string, value *vpp_abx.ABX) (metadata *abxidx.ABXMetadata, err error)
	Delete               func(key string, value *vpp_abx.ABX, metadata *abxidx.ABXMetadata) error
	Update               func(key string, oldValue, newValue *vpp_abx.ABX, oldMetadata *abxidx.ABXMetadata) (newMetadata *abxidx.ABXMetadata, err error)
	UpdateWithRecreate   func(key string, oldValue, newValue *vpp_abx.ABX, metadata *abxidx.ABXMetadata) bool
	Retrieve             func(correlate []ABXKVWithMetadata) ([]ABXKVWithMetadata, error)
	IsRetriableFailure   func(err error) bool
	DerivedValues        func(key string, value *vpp_abx.ABX) []KeyValuePair
	Dependencies         func(key string, value *vpp_abx.ABX) []Dependency
	RetrieveDependencies []string /* descriptor name */
}

////////// Descriptor adapter //////////

type ABXDescriptorAdapter struct {
	descriptor *ABXDescriptor
}

func NewABXDescriptor(typedDescriptor *ABXDescriptor) *KVDescriptor {
	adapter := &ABXDescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:                 typedDescriptor.Name,
		KeySelector:          typedDescriptor.KeySelector,
		ValueTypeName:        typedDescriptor.ValueTypeName,
		KeyLabel:             typedDescriptor.KeyLabel,
		NBKeyPrefix:          typedDescriptor.NBKeyPrefix,
		WithMetadata:         typedDescriptor.WithMetadata,
		MetadataMapFactory:   typedDescriptor.MetadataMapFactory,
		IsRetriableFailure:   typedDescriptor.IsRetriableFailure,
		RetrieveDependencies: typedDescriptor.RetrieveDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Validate != nil {
		descriptor.Validate = adapter.Validate
	}
	if typedDescriptor.Create != nil {
		descriptor.Create = adapter.Create
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.UpdateWithRecreate != nil {
		descriptor.UpdateWithRecreate = adapter.UpdateWithRecreate
	}
	if typedDescriptor.Retrieve != nil {
		descriptor.Retrieve = adapter.Retrieve
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	return descriptor
}

func (da *ABXDescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castABXValue(key, oldValue)
	typedNewValue, err2 := castABXValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *ABXDescriptorAdapter) Validate(key string, value proto.Message) (err error) {
	typedValue, err := castABXValue(key, value)
	if err != nil {
		return err
	}
	return da.descriptor.Validate(key, typedValue)
}

func (da *ABXDescriptorAdapter) Create(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castABXValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Create(key, typedValue)
}

func (da *ABXDescriptorAdapter) Update(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castABXValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castABXValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castABXMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Update(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *ABXDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castABXValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castABXMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *ABXDescriptorAdapter) UpdateWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castABXValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castABXValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castABXMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.UpdateWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *ABXDescriptorAdapter) Retrieve(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []ABXKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castABXValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castABXMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			ABXKVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedValues, err := da.descriptor.Retrieve(correlateWithType)
	if err != nil {
		return nil, err
	}
	var values []KVWithMetadata
	for _, typedKVWithMetadata := range typedValues {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		values = append(values, kvWithMetadata)
	}
	return values, err
}

func (da *ABXDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castABXValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *ABXDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castABXValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

////////// Helper methods //////////

func castABXValue(key string, value proto.Message) (*vpp_abx.ABX, error) {
	typedValue, ok := value.(*vpp_abx.ABX)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castABXMetadata(key string, metadata Metadata) (*abxidx.ABXMetadata, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(*abxidx.ABXMetadata)
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}