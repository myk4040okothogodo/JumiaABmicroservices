Code Generation

Code generation takes the routine work of creating typical code off a developer’s shoulders. 
With code generation, we can spend more time on business logic instead of working on low-level API implementation logic 

In this microservice application i use code generator in the following instances
  -  To generate microservices structures and interfaces from proto files
  - To generate models for graphql and code for a Graphql service


protoc(protocol-buffers)
Proto3 (protocol buffer) is a structured data serialization method that initially allows for describing data structures in the form of messages. 
It also allows for describing a set of operations on this data in the request/response format and for compiling data into an interface. 
Along with an interface, it produces code for serialization and deserialization.
Proto3 is easier to use than other text data formats like XML because it represents data in a way similar to usual data structures instead of as a set of tags. 
Also, when using a protocol buffer, data requires less storage space thanks to the encoding algorithm and is processed quite fast because the serialized data is represented in a binary format.


1. Install Prototool

There are several ways to install Prototool. We need to install it using a binary file, so let’s download it using the following terminal command (we used Prototool version 1.8.0):

curl -sSL \
https://github.com/uber/prototool/releases/download/v1.8.0/prototool-$(uname -s)-$(uname -m) \
  -o /usr/local/bin/prototool && \
  chmod +x /usr/local/bin/prototool
This command will place a binary file in /usr/local/bin and will make it executable.

2. Install protoc

Download the protoc tool from the GitHub page. Unpack it and open the root folder. Then execute the following commands in the terminal:

./configure
Make
make check
sudo make install
sudo ldconfig
3. Install the protoc-gen-go and protoc-gen-gql plugins

Now we need to install the protoc-gen-go and protoc-gen-gql plugins to generate Go code:

go install google.golang.org/protobuf/cmd/protoc-gen-go
go install github.com/danielvladco/go-proto-gql/protoc-gen-gql
These utilities will be installed to $GOPATH/bin if the environment variables (GOPATH, PATH) are correctly configured. GOPATH must be a root folder for the binary files located in $GOPATH/bin and must be the root folder for projects. Also, GOPATH needs to be specified in PATH.

In our case, GOPATH=/home/docker/go/, the utilities protoc-gen-go and protoc-gen-gql are located in /home/docker/go/bin/, and the project is in /home/docker/go/src/.


Create a prototool.yaml file in the root project folder with the following command:

prototool config init
You will get a prototool.yaml file with the following content:

protoc:
  version: 3.8.0 // a version can vary
lint:
  group: uber2
Now let’s add more content to it by specifying the path for Go code generation:

protoc:
  version: 3.8.0
lint:
  group: uber2
  
generate:
  go_options:
    import_path: JumiaABmicroservices
  plugins:
    - name: go
      type: go
      flags: plugins=grpc
      output: gen/go 

Since we’re going to generate code, the first thing we need to do is create data structures and specify operations that will be performed on these structures while the services are working.

Further code generation will be based on these data structures and this set of operations. A code generation utility is just a tool that reads structures and operations on them and uses those structures and operations to generate server code.


Find my definition of data structures in the folder /JumiaABmicroservices/proto/*

Generate and implement microservices code
Now we generate code from proto files by executing the following command in the root folder where our prototool.yaml file is located:

prototool generate
According to the scheme we mentioned at the beginning of this article, Prototool will find all the proto files and use prototool.yaml to place service interfaces in gen/go/*.
Support models will also be there. After that, a shell for the API will be ready and we can start working on implementing the API.
