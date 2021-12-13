// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/golang/protobuf/proto"
	. "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"
	"go.pantheon.tech/stonework/proto/bfd"
)

////////// type-safe key-value pair with metadata //////////

type BfdKVWithMetadata struct {
	Key      string
	Value    *bfd.BFD
	Metadata interface{}
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type BfdDescriptor struct {
	Name                 string
	KeySelector          KeySelector
	ValueTypeName        string
	KeyLabel             func(key string) string
	ValueComparator      func(key string, oldValue, newValue *bfd.BFD) bool
	NBKeyPrefix          string
	WithMetadata         bool
	MetadataMapFactory   MetadataMapFactory
	Validate             func(key string, value *bfd.BFD) error
	Create               func(key string, value *bfd.BFD) (metadata interface{}, err error)
	Delete               func(key string, value *bfd.BFD, metadata interface{}) error
	Update               func(key string, oldValue, newValue *bfd.BFD, oldMetadata interface{}) (newMetadata interface{}, err error)
	UpdateWithRecreate   func(key string, oldValue, newValue *bfd.BFD, metadata interface{}) bool
	Retrieve             func(correlate []BfdKVWithMetadata) ([]BfdKVWithMetadata, error)
	IsRetriableFailure   func(err error) bool
	DerivedValues        func(key string, value *bfd.BFD) []KeyValuePair
	Dependencies         func(key string, value *bfd.BFD) []Dependency
	RetrieveDependencies []string /* descriptor name */
}

////////// Descriptor adapter //////////

type BfdDescriptorAdapter struct {
	descriptor *BfdDescriptor
}

func NewBfdDescriptor(typedDescriptor *BfdDescriptor) *KVDescriptor {
	adapter := &BfdDescriptorAdapter{descriptor: typedDescriptor}
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

func (da *BfdDescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castBfdValue(key, oldValue)
	typedNewValue, err2 := castBfdValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *BfdDescriptorAdapter) Validate(key string, value proto.Message) (err error) {
	typedValue, err := castBfdValue(key, value)
	if err != nil {
		return err
	}
	return da.descriptor.Validate(key, typedValue)
}

func (da *BfdDescriptorAdapter) Create(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castBfdValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Create(key, typedValue)
}

func (da *BfdDescriptorAdapter) Update(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castBfdValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castBfdValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castBfdMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Update(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *BfdDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castBfdValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castBfdMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *BfdDescriptorAdapter) UpdateWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castBfdValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castBfdValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castBfdMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.UpdateWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *BfdDescriptorAdapter) Retrieve(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []BfdKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castBfdValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castBfdMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			BfdKVWithMetadata{
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

func (da *BfdDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castBfdValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *BfdDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castBfdValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

////////// Helper methods //////////

func castBfdValue(key string, value proto.Message) (*bfd.BFD, error) {
	typedValue, ok := value.(*bfd.BFD)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castBfdMetadata(key string, metadata Metadata) (interface{}, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(interface{})
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}
