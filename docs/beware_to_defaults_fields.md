Default Values in Protobuf
By default, protobuf assigns default values to fields that are not explicitly set. Here are the default values for different types:

Numeric Types (int32, int64, uint32, uint64, float, double): 0
Boolean: false
String: "" (empty string)
Bytes: "" (empty bytes)
Enums: The first defined value (which must be 0 by convention)