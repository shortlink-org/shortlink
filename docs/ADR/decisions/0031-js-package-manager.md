# 31. JS: package manager

Date: 2024-02-21

## Status

Accepted

## Context

Our development team has been experiencing performance issues with `npm`, particularly related to slow build times and 
inefficient disk space usage. This has led to the need for exploring alternative package managers that could offer better 
performance and efficiency.

## Decision

After evaluating various options, we have decided to adopt `pnpm` as our primary package manager for JavaScript projects. 
`pnpm` offers significant improvements in terms of speed and disk space utilization by leveraging hard links 
and a shared store for packages.

### Alternatives

#### Bun

[Bun](https://bun.sh/) is a fast JS runtime, bundler, and package manager. 
It is designed to be a drop-in replacement for `npm`, but it's a very new project, so it's not as mature as `pnpm`.

## Consequences

The transition to `pnpm` is expected to result in faster build times and reduced disk space consumption 
for our JavaScript projects. This decision will necessitate changes in our development workflow, 
including updates to build scripts and CI/CD pipelines.

### Migration from **npm** to **pnpm*

For local development environments and Dockerfiles:

```shell
npm install -g pnpm
rm -rf node_modules
pnpm install
rm -rf package-lock.json
```

For Continuous Integration (CI) environments:

```shell
corepack enable
corepack prepare pnpm@latest-8 --activate
```

Additionally, we recommend updating the `package.json` scripts to enforce the use of pnpm and improve compatibility:

```shell
"scripts": {
  "build": "pnpm build",
  "preinstall": "npx only-allow pnpm"
}
```

### References

- [pnpm](https://pnpm.io/) - Official pnpm documentation and guides.
