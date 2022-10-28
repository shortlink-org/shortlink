# SBOMs

### Example

```shell
$> docker run -it --rm \
  -v $PWD:/repository \
  -v "$(pwd)/out:/out" 
  spdx/spdx-sbom-generator \
    -p /repository/pkg/shortdb-operator 
    -o /docs/bom
```

### Docs

- [The ultimate guide to SBOMs](https://about.gitlab.com/blog/2022/10/25/the-ultimate-guide-to-sboms/)
