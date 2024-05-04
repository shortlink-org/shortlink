## ShortLink plugin for Istio

Go 1.21 support build to WebAssembly. So we can write a plugin for Istio using Go without tinygo.

### References

- [WASI support in Go](https://go.dev/blog/wasi)
- [Istio: wasm-plugin](https://istio.io/latest/docs/reference/config/proxy_extensions/wasm-plugin/)
- [proxy-wasm-go-sdk](https://github.com/tetratelabs/proxy-wasm-go-sdk)
- Tutorial:
  - [Writing an Istio WASM Plugin in Go for migrating 100s of services to new auth strategy (Part 1)](https://zendesk.engineering/writing-an-istio-wasm-plugin-in-go-for-migrating-100s-of-services-to-new-auth-strategy-part-1-cd551e1455d7)
    - [github](https://github.com/henders/writing-an-envoy-wasm-plugin)
  - [Example using with EnvoyFilter](https://gitverse.ru/kozlov.a.e/wasm-gofunc/content/master/conf/envoy_filter.yml)
