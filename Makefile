all: bin/lsprofiles

bin/lsprofiles: $(wildcard lsprofiles/*)
	cd lsprofiles && go build -o ../bin/lsprofiles
