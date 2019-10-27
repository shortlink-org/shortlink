package schema

// Use `go generate` to pack all *.graphql files under this directory (and sub-directories) into
// a binary format.
//
//go:generate go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...

import "bytes"

// GetRootSchema - string reads the .graphql schema files from the generated _bindata.go file, concatenating the
// files together into one string.
//
// If this method complains about not finding functions AssetNames() or MustAsset(),
// run `go generate` against this package to generate the functions.
func GetRootSchema() string { // nolint unused
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b := MustAsset(name)
		if _, err := buf.Write(b); err != nil {
			panic(err)
		}

		// Add a newline if the file does not end in a newline.
		if len(b) > 0 && b[len(b)-1] != '\n' {
			if err := buf.WriteByte('\n'); err != nil {
				panic(err)
			}
		}
	}

	return buf.String()
}
