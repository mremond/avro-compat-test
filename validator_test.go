package validator

import (
	"bytes"
	"testing"

	s1 "github.com/mremond/avro-compat-test/schema1"
	s2 "github.com/mremond/avro-compat-test/schema2"
)

func TestV2ToV1(t *testing.T) {
	message := s2.Schema2{
		RequestId: unionNullString2("az-123"),
		ProductId:      123,
		ProductName:  "Macbook Pro",
		Brand:  "Apple",
	}
	var buf bytes.Buffer
	message.Serialize(&buf)

	messageS1, err := s1.DeserializeSchema1(&buf)
	if err != nil {
		t.Error("Cannot deserialize object2 with schema1:", err)
		return
	}

	if messageS1.ProductName != message.ProductName {
		t.Error("Fields messed up on decode. ProductName:", messageS1.ProductName)
	}
}

func TestV1ToV2(t *testing.T) {
	message := s1.Schema1{
		RequestId: unionNullString("az-123"),
		ProductId:      123,
		ProductName:  "Macbook Pro",
	}
	var buf bytes.Buffer
	message.Serialize(&buf)

	messageS2, err := s2.DeserializeSchema2(&buf)
	if err != nil {
		t.Error("Cannot deserialize object1 with schema2:", err)
		return
	}

	t.Log("Brand:", messageS2.Brand)
}

// TODO Should we patch gogen-avro to generate this method ?
func unionNullString(s string) s1.UnionNullString {
	return s1.UnionNullString{
		String:    s,
		UnionType: s1.UnionNullStringTypeEnumString,
	}
}

// TODO Should we patch gogen-avro to generate this method ?
func unionNullString2(s string) s2.UnionNullString {
	return s2.UnionNullString{
		String:    s,
		UnionType: s2.UnionNullStringTypeEnumString,
	}
}
