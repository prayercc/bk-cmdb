package bsonx

import (
	"github.com/rentiansheng/bk_bson/bson"
	"github.com/rentiansheng/bk_bson/bson/bsoncodec"
)

// DefaultRegistry is the default bsoncodec.Registry. It contains the default codecs and the
// primitive codecs.
var DefaultRegistry = NewRegistryBuilder().Build()

// NewRegistryBuilder creates a new RegistryBuilder configured with the default encoders and
// deocders from the bsoncodec.DefaultValueEncoders and bsoncodec.DefaultValueDecoders types and the
// PrimitiveCodecs type in this package.
func NewRegistryBuilder() *bsoncodec.RegistryBuilder {
	rb := bsoncodec.NewRegistryBuilder()
	bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(rb)
	bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(rb)
	bson.PrimitiveCodecs{}.RegisterPrimitiveCodecs(rb)
	primitiveCodecs.RegisterPrimitiveCodecs(rb)
	return rb
}
