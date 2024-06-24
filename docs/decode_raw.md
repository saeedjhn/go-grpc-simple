**************************** decode_raw *********************************
Run command line:
> cat file.bin | protoc --decode_raw
(for example: cat simple.bin | protoc --decode_raw)
golang_with_protocol_buffer git:(master) cat simple.bin | protoc --decode_raw
1: 52
3: "Leader"
4: "\001\002\003\004"

*************************** decode **************************************
Run command line:
> cat file.bin | protoc --decode=MESSAGE_TYPE file.proto
(for example: cat simple.bin | protoc --decode=Simple simple.proto)
golang_with_protocol_buffer git:(master) cat simple.bin | protoc --decode=simple.Simple proto/simple.proto
id: 52
name: "Leader"
sample_lists: 1
sample_lists: 2
sample_lists: 3
sample_lists: 4

decode toFile
> cat simple.bin | protoc --decode=simple.Simple proto/simple.proto > simple.txt

*************************** encode **************************************
Run command line:
> cat file.txt | protoc --encode=MESSAGE_TYPE file.proto > file.bin
(for example: cat simple.txt | protoc --encode=Simple simple.proto > simple.bin)

simple.txt:
id: 52
name: "Leader"
sample_lists: 1
sample_lists: 2
sample_lists: 3
sample_lists: 4
> cat simple.txt | protoc --encode=simple.Simple proto/simple.proto > simple.bin
> diff file1 file2