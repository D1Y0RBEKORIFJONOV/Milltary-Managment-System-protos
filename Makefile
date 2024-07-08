gen-soldiers:
	protoc -I proto proto/soldiers/soldiers.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
gen-storehouses:
	protoc -I proto proto/storehouses/storehouses.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative

gen-group:
	protoc -I proto proto/group/group.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative