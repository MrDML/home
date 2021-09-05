

.PHONY: init
# generate api
init:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'echo "$$0" && cd "$$0" && pwd && $(MAKE) init'

.PHONY: api
# generate api
api:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'echo "$$0" && cd "$$0" && pwd && $(MAKE) api'


.PHONY: proto
# generate proto
proto:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'echo "$$0" && cd "$$0" && pwd && $(MAKE) proto'


#.PHONY: init
## init env
#init:
##	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
##	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
##	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
##	go get -u github.com/google/wire/cmd/wire
##	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
##	go get -u github.com/go-kratos/gin
##	go get -u github.com/gin-gonic/gin
##	go get github.com/afocus/captcha
##	go get github.com/go-redis/redis/v8
#	# 短信核心库
#	go get -u github.com/aliyun/alibaba-cloud-sdk-go/sdk
#	# 短信API
#	go get github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi