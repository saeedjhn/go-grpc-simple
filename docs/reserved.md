Example of Using reserved in a .proto File
Here is an example demonstrating how to use the reserved keyword in a .proto file:

proto
Copy code
syntax = "proto3";

package example;

message MyMessage {
int32 id = 1;
string name = 2;

// Reserving the field number 3 and the name "email"
reserved 3;
reserved "email";
}
Explanation:
Reserving Field Numbers:

The reserved 3; statement indicates that the field number 3 should not be reused for any new fields in the MyMessage message. This prevents any new fields from accidentally using this field number, which could cause issues if older versions of the message are still in use.
Reserving Field Names:

The reserved "email"; statement indicates that the field name email should not be reused. This helps avoid confusion and potential conflicts in the message definition.
Extended Example with Multiple Reserved Fields
You can reserve multiple field numbers and names at once:

proto
Copy code
syntax = "proto3";

package example;

message MyMessage {
int32 id = 1;
string name = 2;

// Reserving multiple field numbers
reserved 3, 4, 5;

// Reserving multiple field names
reserved "email", "phone", "address";
}
Combining Reserved Field Numbers and Names
You can combine both reserved field numbers and names in one reserved statement:

proto
Copy code
syntax = "proto3";

package example;

message MyMessage {
int32 id = 1;
string name = 2;

// Reserving field numbers and names
reserved 3, 4, 5, "email", "phone", "address";
}
Reasons to Use reserved
Preventing Accidental Reuse:

Reserving field numbers and names helps ensure that they are not accidentally reused in the future, which can cause compatibility issues.
Maintaining Compatibility:

By reserving fields that have been removed or deprecated, you help maintain backward and forward compatibility in your protocol buffers. This is particularly important in large projects with multiple versions of message definitions.
Clear Documentation:

Using the reserved keyword serves as a form of documentation for other developers, indicating that certain field numbers or names are intentionally not to be reused.
Summary
The reserved keyword in protobuf is a useful tool for maintaining the integrity and compatibility of your message definitions. By reserving field numbers and names that are no longer in use, you can prevent accidental conflicts and ensure that your protocol buffer messages remain consistent and reliable over time.





