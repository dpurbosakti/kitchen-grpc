run-orders:
	@go run services/orders/*.go

run-kitchen:
	@go run services/kitchen/*.go

## gen-order: generate go code from proto
gen-order:
	rm -f services/common/genproto/orders/*.go
	@protoc \
	--proto_path=protobuf \
	"protobuf/orders.proto" \
	--go_out=services/common/genproto/orders \
	--go_opt=paths=source_relative \
	--go-grpc_out=services/common/genproto/orders \
	--go-grpc_opt=paths=source_relative

# after running this command, before u call the RPC, try tp check the package and service, then use package <package_name> and service <service_name> 
## evans: connect to evans
evans:
	@evans --host localhost --port 9000 -r