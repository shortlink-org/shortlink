# ShortLink

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-init.md) - Adopting Expo Monorepo with Nx for ShortLink Service Client

### Getting Started

| Action        | Command                                                                                                                                                                                                     |
|---------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Start the App | `nx serve shortlink` - Access at [http://localhost:4200/](http://localhost:4200/).                                                                                                                          |
| Generate Code | `nx list` for available plugins. For specific generators, use `nx list <plugin-name>`.                                                                                                                      |
| Running Tasks | Use `nx <target> <project> <...options>` to execute tasks.                                                                                                                                                  |
| Deployment    | Build with `nx build shortlink`. Artifacts in `dist/` directory.                                                                                                                                            |
| CI Setup      | Nx provides local caching (see `nx.json`). Enhance CI with [remote caching](https://nx.dev/core-features/share-your-cache) and [task distribution](https://nx.dev/core-features/distribute-task-execution). |

---

<a alt="Nx logo" href="https://nx.dev" target="_blank" rel="noreferrer"><img src="https://raw.githubusercontent.com/nrwl/nx/master/images/nx-logo.png" width="30"></a>
Powered by [Nx](https://nx.dev) - a smart and extensible build system.
