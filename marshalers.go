package decimal

import (
	"fmt"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

func (d Decimal) MarshalBSONValue() (bsontype.Type, []byte, error) {
	dec128, err := primitive.ParseDecimal128(decimal.Decimal(d).String())
	if err != nil {
		return bson.TypeDecimal128, nil, err
	}

	return bson.MarshalValue(dec128)
}

func (d *Decimal) UnmarshalBSONValue(t bsontype.Type, bytes []byte) error {
	if t != bson.TypeDecimal128 {
		return nil
	}

	decimal128, _, ok := bsoncore.ReadDecimal128(bytes)
	if !ok {
		return fmt.Errorf("bsoncore.ReadDecimal128(%v) error", bytes)
	}

	d1, err := decimal.NewFromString(decimal128.String())
	if err != nil {
		return err
	}

	*d = (Decimal)(d1)

	return nil
}

func (d Decimal) MarshalBinary() (data []byte, err error) {
	return (decimal.Decimal)(d).MarshalBinary()
}

func (d *Decimal) UnmarshalBinary(data []byte) error {
	return (*decimal.Decimal)(d).UnmarshalBinary(data)
}

func (d Decimal) MarshalJSON() ([]byte, error) {
	return decimal.Decimal(d).MarshalJSON()
}

func (d *Decimal) UnmarshalJSON(b []byte) error {
	return (*decimal.Decimal)(d).UnmarshalJSON(b)
}

// TODO sql things like Scan
