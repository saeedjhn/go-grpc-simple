syntax = "proto3";

package example;

message ExampleRequest {
    int32 id = 1;
}

message ExampleResponse {
    int32 int_value = 1;
    float float_value = 2;
    string string_value = 3;
    bytes bytes_value = 4;
    repeated int32 int_list = 5;
    NestedMessage nested_message = 6;
    ExampleEnum example_enum = 7;
}

message NestedMessage {
    string nested_string = 1;
    int64 nested_int = 2;
}

enum ExampleEnum {
    UNKNOWN = 0;
    FIRST_OPTION = 1;
    SECOND_OPTION = 2;
}

service ExampleService {
    rpc GetExampleData (ExampleRequest) returns (ExampleResponse);
}
