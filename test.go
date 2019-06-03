package main

import "fmt"

type PropertyValue struct {
	Type         string `bson:"type" json:"type"`                                     // ValueDescriptor Type of property after transformations
	ReadWrite    string `bson:"readWrite" json:"readWrite" yaml:"readWrite"`          // Read/Write Permissions set for this property
	Minimum      string `bson:"minimum" json:"minimum"`                               // Minimum value that can be get/set from this property
	Maximum      string `bson:"maximum" json:"maximum"`                               // Maximum value that can be get/set from this property
	DefaultValue string `bson:"defaultValue" json:"defaultValue" yaml:"defaultValue"` // Default value set to this property if no argument is passed
	Size         string `bson:"size" json:"size"`                                     // Size of this property in its type  (i.e. bytes for numeric types, characters for string types)
	Word         string `bson:"word" json:"word"`                                     // Word size of property used for endianness
	LSB          string `bson:"lsb" json:"lsb"`                                       // Endianness setting for a property
	Mask         string `bson:"mask" json:"mask"`                                     // Mask to be applied prior to get/set of property
	Shift        string `bson:"shift" json:"shift"`                                   // Shift to be applied after masking, prior to get/set of property
	Scale        string `bson:"scale" json:"scale"`                                   // Multiplicative factor to be applied after shifting, prior to get/set of property
	Offset       string `bson:"offset" json:"offset"`                                 // Additive factor to be applied after multiplying, prior to get/set of property
	Base         string `bson:"base" json:"base"`                                     // Base for property to be applied to, leave 0 for no power operation (i.e. base ^ property: 2 ^ 10)
	Assertion    string `bson:"assertion" json:"assertion"`                           // Required value of the property, set for checking error state.  Failing an assertion condition will mark the device with an error state
	Signed       bool   `bson:"signed" json:"signed"`                                 // Treat the property as a signed or unsigned value
	Precision    string `bson:"precision" json:"precision"`
}

type ProfileProperty struct {
	Value PropertyValue `bson:"value" json:"value"`
}

type DeviceObject struct {
	Properties ProfileProperty `bson:"properties" json:"properties" yaml:"properties"`
}

func main() {
	var b []int
	b = append(b, 1)
	b = append(b, 2)
	b = append(b, 3)
	b = append(b, 4)
	b = append(b, 5)
	var a = make([]int, 5)
	copy(a, b)
	fmt.Println(len(b), cap(b), b)
	b = b[:5+3]
	fmt.Println(len(b), cap(b), b)
	var abc string
	abc = `bson:"type" json:"type"`
	fmt.Println(abc)

	t := new(DeviceObject)
	fmt.Println(t.Properties.Value.Type)
}
