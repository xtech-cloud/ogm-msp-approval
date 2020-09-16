.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/approval/shared.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/approval/workflow.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/approval/operator.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/approval/task.proto
