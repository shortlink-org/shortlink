FROM gcr.io/oss-fuzz-base/base-builder-go

RUN git clone --depth 1 https://github.com/shortlink-org/shortlink.git shortlink

# Use curl-fuzzer's scripts to get latest dependencies.
# RUN $SRC/shortlink/scripts/ossfuzzdeps.sh

WORKDIR $SRC/shortlink
