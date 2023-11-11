FROM gcr.io/oss-fuzz-base/base-builder-go

RUN git clone --depth 1 https://github.com/shortlink-org/shortlink.git shortlink

WORKDIR $SRC/shortlink
COPY .clusterfuzzlite/build.sh $SRC/
