all: bin/lsprofiles

bin/lsprofiles: $(wildcard cmd/lsprofiles/*)
	cd cmd/lsprofiles && go build -o ../bin/lsprofiles
