GEN_FOLDER=src/main/java
mkdir -p $GEN_FOLDER
protoc -I$GOPATH/src -I. --plugin=protoc-gen-grpc-java=/usr/local/bin/protoc-gen-grpc-java.exe --grpc-java_out=. --java_out=$GEN_FOLDER  *.proto
mvn package