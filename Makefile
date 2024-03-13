PRODUCT := filemagego
ALIAS := fm

PATH  := $(PATH):$(HOME)/go/bin
URL := https://api.filemage.io/abb5d69a-f0c9-4a49-9b25-e419ddfe613f
#GO_POST_PROCESS_FILE := "go doc ." 

GOOPT := --go_opt=Mproto/$(PRODUCT).proto=github.com/elgranjero/$(PRODUCT)

properties := --enable-post-process-file --git-user-id=elgranjero --git-repo-id=$(PRODUCT) --minimal-update \
	--additional-properties=generateInterfaces=true,packageName=$(PRODUCT),goImportAlias=$(ALIAS),withGoCodegenComment=true

build: schema.json
	openapi-generator generate -i schema.json -g go -t author $(properties)

schema.json:
	wget -q -O - $(URL) | jq -r . > schema.json

doc:
	mkdir -p docs
	go doc . > docs/funcs.txt
	openapi-generator generate -i schema.json -g html2 $(properties)

clean:
	rm -rf *.go go.* api docs test index.html git_push.sh *.md

client: schema.json
	mkdir -p ./$(PRODUCT)
	swagger generate client -f schema.json -t ./$(PRODUCT)-client

all: clean build doc
