export GOPATH := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

#####################

#SOURCES_BASE = src/types/types.go src/readline/readline.go \
#	       src/reader/reader.go src/printer/printer.go \
#	       src/env/env.go src/core/core.go

#####################

SRCS = step0_repl.go 
BINS = $(SRCS:%.go=%)

#####################

all: $(BINS)

dist: mal

mal: $(word $(words $(BINS)),$(BINS))
	cp $< $@

define dep_template
$(1): $(SOURCES_BASE) src/$(1)/$(1).go
	go build $$@
endef

$(foreach b,$(BINS),$(eval $(call dep_template,$(b))))

clean:
	rm -f $(BINS) mal
