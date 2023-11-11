# These two dependencies cause build issues and are not used by oss-fuzz:
rm -r sqlparser
rm -r parser

go mod init "github.com/shortlink-org/shortlink"
export FUZZ_ROOT="github.com/shortlink-org/shortlink"
