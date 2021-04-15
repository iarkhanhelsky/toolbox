BIN_DIR=$(realpath bin/)

all: bin/lsprofiles

bin/lsprofiles: $(wildcard cmd/lsprofiles/*)
	cd cmd/lsprofiles && go build -o $(BIN_DIR)/lsprofiles
