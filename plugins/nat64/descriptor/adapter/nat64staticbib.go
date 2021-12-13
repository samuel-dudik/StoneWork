// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/golang/protobuf/proto"
	. "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"
	"go.pantheon.tech/stonework/proto/nat64"
)

////////// type-safe key-value pair with metadata //////////

type NAT64StaticBIBKVWithMetadata struct {
	Key      string
	Value    *nat64.Nat64StaticBIB
	Metadata interface{}
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type NAT64StaticBIBDescriptor struct {
	Name                 string
	KeySelector          KeySelector
	ValueTypeName        string
	KeyLabel             func(key string) string
	ValueComparator      func(key string, oldValue, newValue *nat64.Nat64StaticBIB) bool
	NBKeyPrefix          string
	WithMetadata         bool
	MetadataMapFactory   MetadataMapFactory
	Validate             func(key string, value *nat64.Nat64StaticBIB) error
	Create               func(key string, value *nat64.Nat64StaticBIB) (metadata interface{}, err error)
	Delete               func(key string, value *nat64.Nat64StaticBIB, metadata interface{}) error
	Update               func(key string, oldValue, newValue *nat64.Nat64StaticBIB, oldMetadata interface{}) (newMetadata interface{}, err error)
	UpdateWithRecreate   func(key string, oldValue, newValue *nat64.Nat64StaticBIB, metadata interface{}) bool
	Retrieve             func(correlate []NAT64StaticBIBKVWithMetadata) ([]NAT64StaticBIBKVWithMetadata, error)
	IsRetriableFailure   func(err error) bool
	DerivedValues        func(key string, value *nat64.Nat64StaticBIB) []KeyValuePair
	Dependencies         func(key string, value *nat64.Nat64StaticBIB) []Dependency
	RetrieveDependencies []string /* descriptor name */
}

////////// Descriptor adapter //////////

type NAT64StaticBIBDescriptorAdapter struct {
	descriptor *NAT64StaticBIBDescriptor
}

func NewNAT64StaticBIBDescriptor(typedDescriptor *NAT64StaticBIBDescriptor) *KVDescriptor {
	adapter := &NAT64StaticBIBDescriptorAdapter{descriptor: typedDescriptor}
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

func (da *NAT64StaticBIBDescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castNAT64StaticBIBValue(key, oldValue)
	typedNewValue, err2 := castNAT64StaticBIBValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *NAT64StaticBIBDescriptorAdapter) Validate(key string, value proto.Message) (err error) {
	typedValue, err := castNAT64StaticBIBValue(key, value)
	if err != nil {
		return err
	}
	return da.descriptor.Validate(key, typedValue)
}

func (da *NAT64StaticBIBDescriptorAdapter) Create(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castNAT64StaticBIBValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Create(key, typedValue)
}

func (da *NAT64StaticBIBDescriptorAdapter) Update(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castNAT64StaticBIBValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castNAT64StaticBIBValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castNAT64StaticBIBMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Update(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *NAT64StaticBIBDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castNAT64StaticBIBValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castNAT64StaticBIBMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *NAT64StaticBIBDescriptorAdapter) UpdateWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castNAT64StaticBIBValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castNAT64StaticBIBValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castNAT64StaticBIBMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.UpdateWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *NAT64StaticBIBDescriptorAdapter) Retrieve(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []NAT64StaticBIBKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castNAT64StaticBIBValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castNAT64StaticBIBMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			NAT64StaticBIBKVWithMetadata{
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

func (da *NAT64StaticBIBDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castNAT64StaticBIBValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *NAT64StaticBIBDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castNAT64StaticBIBValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

////////// Helper methods //////////

func castNAT64StaticBIBValue(key string, value proto.Message) (*nat64.Nat64StaticBIB, error) {
	typedValue, ok := value.(*nat64.Nat64StaticBIB)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castNAT64StaticBIBMetadata(key string, metadata Metadata) (interface{}, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(interface{})
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}
