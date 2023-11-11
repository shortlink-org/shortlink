FROM gcr.io/oss-fuzz-base/base-builder

RUN git clone --depth 1 https://github.com/shortlink-org/shortlink.git /src/shortlink

# Use curl-fuzzer's scripts to get latest dependencies.
# RUN $SRC/shortlink/scripts/ossfuzzdeps.sh

WORKDIR $SRC/shortlink
COPY .clusterfuzzlite/build.sh $SRC/
