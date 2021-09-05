
.PHONY: api
# generate api
api:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'echo "$$0" && cd "$$0" && pwd && $(MAKE) api'


.PHONY: proto
# generate proto
proto:
	find app -type d -depth 2 -print | xargs -L 1 bash -c 'echo "$$0" && cd "$$0" && pwd && $(MAKE) proto'