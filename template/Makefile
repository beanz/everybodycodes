DAY=$(notdir $(CURDIR))
YEAR=$(notdir $(patsubst %/,%,$(dir $(CURDIR))))
HERE=$(YEAR)/$(DAY)

ifneq (,$(wildcard ec.go))
ALL += ec-go
CHECK += ec-go
CLEAN += ec-go gocpu.out cover.out gomem.out
TEST += test-go
endif
ifneq (,$(wildcard ec.pl))
ALL += ec.pl
TEST += test-pl
endif
ifneq (,$(wildcard ec.zig))
ALL += ec-zig
CHECK += ec-zig-rel
CLEAN += ec-zig ec-zig-rel ec-lib.zig zig-cpu.svg
CLEAN_DIR += zig-cache
TEST += test-zig
HYPERFINE += ec-zig-rel
endif

# TODO: fix rust builds
RS_SRC=../../ec-rust/src/bin/ec-$(YEAR)-$(DAY).rs
RS_BIN=$(subst src/bin,target/debug,${RS_SRC:.rs=})
RS_REL=$(subst src/bin,target/release,${RS_SRC:.rs=})
ifneq (,$(wildcard $(RS_SRC)))
ALL += ec-rs
CHECK += ec-rs-rel
CLEAN += ec-rs
HYPERFINE += ec-rs-rel
TEST += test-rs
endif

ifneq (,$(wildcard ec.jq))
ALL += ec.jq
endif
TIMES=$(addsuffix -time,${ALL})

EXTRA=out.perf-folded perf.data perf.data.old
FG_HOME=$(HOME)/git/github.com/brendangregg/FlameGraph

all: ${ALL}

clean:
	-rm -f ${CLEAN} $(notdir $(CURDIR)).test ${EXTRA} *.err *.log *.chk
	-rm -rf ${CLEAN_DIR} nytprof nytprof.out _Inline
	-find . -name '*~' -print0 |xargs -0 rm -f

times: ${TIMES}

%-time: %
	@echo $<
	@/usr/bin/time ./$<
	@echo

ec-go: ec.go ../../lib-go/*.go
	go build -o $@

bench-go: gocpu.out

gocpu.out: ec.go
	BENCHTIME=5s ; \
	if [ -f .ec-go.benchtime ]; then BENCHTIME=$$(cat .ec-go.benchtime); fi;\
	go test $$(grep -q Bench *_test.go && echo "-run=XXX -benchmem -benchtime=$$BENCHTIME -bench .") -cpuprofile=$@ -memprofile=gomem.out

prof-go: gocpu.out
	go tool pprof -http=":9999" $<

cover-go: ec.go ec_test.go
	go test -coverprofile cover.out
	go tool cover -html=cover.out
	rm cover.out

zig-cpu.svg: ec-zig
	EC_BENCH=1 perf record -F 99 --call-graph dwarf ./$<
	perf script | $(FG_HOME)/stackcollapse-perf.pl >out.perf-folded
	$(FG_HOME)/flamegraph.pl out.perf-folded >$@
	perf report -g fractal -F+period,srcline
	-rm -f out.perf-folded perf.data perf.data.old

prof-zig: zig-cpu.svg
	google-chrome $<

test: ${TEST}

test-zig: ec-zig
	zig test ec.zig

test-go:
	go test -v

test-go-fast:
	go test -short

test-pl:
	EC_TEST=2 perl ec.pl input.txt

prof-perl:
	perl -d:NYTProf ec.pl input.txt
	nytprofhtml --open

check: ${ALL}
	@TT=$$(mktemp -d) ; \
	GOLD=$$TT/gold.log ; \
	./$< >$$GOLD ; \
	echo -n "$(HERE): " ; \
	for f in $^ ; do \
	  test -f ".$$f.slow" && continue ; \
	  if [ "$<" = "$$f" ]; then echo -n "$$f "; continue ; fi; \
	  echo -n "$$f "; \
	  ./$$f >$$TT/$$f.log 2>&1 \
	    && diff -audBbw $$GOLD $$TT/$$f.log >$$TT/$$f.diff \
	    || ( echo ; echo $$f output differs ; cat $$TT/$$f.diff); \
	done && \
	  rm -rf $$TT ; \
	echo

ec-rs: $(RS_SRC)
	( cd ../../ec-rust && \
	  cargo build --bin "$(basename $(notdir $<))" ) && \
	mv $(RS_BIN) $@

ec-rs-rel: $(RS_SRC)
	( cd ../../ec-rust && \
	  cargo build --release --bin "$(basename $(notdir $<))" ) && \
	mv $(RS_REL) $@

test-rs: $(RS_SRC)
	( cd ../../ec-rust && \
	  RUST_MIN_STACK=8388608 cargo test --bin "$(basename $(notdir $<))" )

prof-rs: $(RS_SRC)
	( cd ../../ec-rust && \
	  EC_BENCH=1 cargo flamegraph --bin "$(basename $(notdir $<))" \
		   -o ../$(YEAR)/$(DAY)/flamegraph.svg \
		   -- ../$(YEAR)/$(DAY)/input.txt ; \
		echo )

cover-rs: $(RS_SRC)
	( cd ../../ec-rust && \
	  RUSTFLAGS="-C instrument-coverage" \
		  cargo test --bin "$(basename $(notdir $<))" )
	llvm-profdata merge -sparse ../../ec-rust/default_*.profraw -o test.profdata
	llvm-cov report \
    --use-color \
		--ignore-filename-regex='/.cargo/registry' \
		--ignore-filename-regex='src/lib.rs' \
    --instr-profile=test.profdata \
    --object ../../ec-rust/target/debug/deps/ec_$(YEAR)_$(DAY)-????????????????
	llvm-cov show \
    --use-color \
		--ignore-filename-regex='/.cargo/registry' \
		--ignore-filename-regex='src/lib.rs' \
    --instr-profile=test.profdata \
    --object ../../ec-rust/target/debug/deps/ec_$(YEAR)_$(DAY)-???????????????? \
    --show-instantiations --show-line-counts-or-regions \
    --Xdemangler=rustfilt

clippy: $(RS_SRC)
	( cd ../../ec-rust && cargo clippy )

expect-pl: ec.pl
	@if [ -z "$$EC_E" ]; then echo 2>&1 "EC_E must be set"; exit 1; fi
	@while true; do perl ec.pl $${TEST:-test1.txt} | tee /dev/stderr |grep -Fq "$$EC_E" && break; sleep 1; done; perl ec.pl

expect-go: ec-go
	@if [ -z "$$EC_E" ]; then echo 2>&1 "EC_E must be set"; exit 1; fi
	@while true; do go build -o ec-go && ./ec-go $${TEST:-test1.txt} |tee /dev/stderr |grep -Fq "$$EC_E" && break; sleep 1; done; ./ec-go

template:
	@if [ ! -e Makefile ]; then ln -s ../../template/Makefile . ; fi ; \
	for f in ec.pl ec.go ec_test.go ec.zig TC.txt; do \
	  if [ ! -e $$f ]; then \
		  echo copying $$f ; cp -p ../../template/$$f . ; \
  	fi \
	done ; \
	frs="../../ec-rust/src/bin/ec-$(YEAR)-$(DAY).rs" ; \
	if [ ! -e $$frs ]; then \
	  echo copying ec-yyyy-dd.rs ; \
		sed -e's/YYYY/'"$(YEAR)"'/g;s/DD/'"$(DAY)"'/g' \
		  <../../template/ec-yyyy-dd.rs >$$frs; \
	fi
