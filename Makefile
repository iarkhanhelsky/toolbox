BIN_DIR=$(realpath bin/)

all: bin/lsprofiles bin/lshosts

bin/lsprofiles: $(wildcard cmd/lsprofiles/*)
	cd cmd/lsprofiles && go build -o $(BIN_DIR)/lsprofiles

bin/lshosts: $(wildcard cmd/lshosts/*)
	cd cmd/lshosts && go build -o $(BIN_DIR)/lshosts