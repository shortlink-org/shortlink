## Istio extension

<img width='200' height='200' src="./docs/public/logo.svg">

> [!NOTE]
> It's plugins for Istio that are written in Go and compiled to WebAssembly.

| Name                                         | Description                                                   | Language        |
|----------------------------------------------|---------------------------------------------------------------|-----------------|
| [shortlink-go](./shortlink_go/README.md)     | A simple example of an Istio extension that shortens the URL. | Go (TinyGo)     |
| [shortlink-rust](./shortlink_rust/README.md) | A simple example of an Istio extension that shortens the URL. | Rust (Wasmtime) |                              

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-contract.md) - Contract for Istio extension.
