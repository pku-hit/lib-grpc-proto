mkdir -p src/main/java
protoc --plugin=protoc-gen-grpc-java=/usr/local/bin/protoc-gen-grpc-java.exe --grpc-java_out=./src/main/java --java_out=./src/main/java  *.proto