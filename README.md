# gRPC Go Projects – Build Automation with Makefile

This repository contains multiple **gRPC-based Go projects** (`greet`, `calculator`, and `blog`). The build process is
automated using a `Makefile` that provides commands for generating protobuf files, compiling servers and clients,
running tests, cleaning up generated files, and displaying environment information.

---

## 📂 Project Structure

```
.
├── greet/
│   ├── proto/        # Protocol Buffers definitions for greet service
│   ├── goproto/      # Generated Go code from protobuf definitions (pb.go files)
│   ├── server/       # gRPC server implementation for greet
│   └── client/       # gRPC client implementation for greet
├── calculator/
│   ├── proto/        # Protocol Buffers definitions for calculator service
│   ├── goproto/      # Generated Go code from protobuf definitions (pb.go files)
│   ├── server/       # gRPC server implementation for calculator
│   └── client/       # gRPC client implementation for calculator
├── blog/
│   ├── proto/        # Protocol Buffers definitions for blog service
│   ├── goproto/      # Generated Go code from protobuf definitions (pb.go files)
│   ├── server/       # gRPC server implementation for blog
│   └── client/       # gRPC client implementation for blog
├── bin/              # Compiled binaries (output directory)
├── Makefile          # Automation of build, test, and clean commands
└── go.mod            # Go module definition
```

---

## ⚙️ Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/dl/) (1.18 or higher recommended)
- [Protocol Buffers Compiler (`protoc`)](https://grpc.io/docs/protoc-installation/)
- [protoc-gen-go](https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go)
- [protoc-gen-go-grpc](https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc)
- [OpenSSL](https://www.openssl.org/) (for SSL certificate generation if needed)
- Make (Linux/macOS) or PowerShell (Windows)

---

## 📜 Makefile Overview

The `Makefile` provides a set of helpful commands to:

- Generate protobuf files
- Build gRPC server and client binaries
- Run tests
- Clean up generated artifacts
- Display system and build information

---

## 🚀 Usage

### Build all projects

```sh
make all
```

This generates protobufs and builds both **server** and **client** binaries for all projects (`greet`, `calculator`,
`blog`).

### Build a specific project

```sh
make greet
make calculator
make blog
```

Each command generates protobuf files and compiles the binaries for that specific project.

### Clean generated files

```sh
make clean
```

Removes generated protobufs, SSL files, and compiled binaries.

You can also clean a specific project:

```sh
make clean_greet
make clean_calculator
make clean_blog
```

### Rebuild everything

```sh
make rebuild
```

Cleans and rebuilds all services.

### Update dependencies

```sh
make bump
```

Fetches the latest versions of Go dependencies.

### Display environment information

```sh
make about
```

Shows details such as:

- Operating system
- Shell and version
- Protobuf compiler version
- Go version
- Go package name
- OpenSSL version

### Show available commands

```sh
make help
```

Lists all available Makefile commands with descriptions.

---

## 📦 Output

After running build commands, the compiled binaries are placed inside the `bin/` directory:

```
bin/
├── greet/
│   ├── server (or server.exe on Windows)
│   └── client (or client.exe on Windows)
├── calculator/
│   ├── server
│   └── client
└── blog/
    ├── server
    └── client
```

---

## 🛠️ Example Workflow

1. Build the **greet service**:
   ```sh
   make greet
   ```
2. Run the gRPC server:
   ```sh
   ./bin/greet/server
   ```
3. Run the gRPC client in another terminal:
   ```sh
   ./bin/greet/client
   ```

---

## 🧾 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

---

## 📖 Summary

This repository provides a clean and automated setup for managing multiple gRPC services in Go. Using the provided
`Makefile`, you can easily generate protobuf files, build clients and servers, run tests, clean up builds, and check
your environment setup across **Linux, macOS, and Windows**.
