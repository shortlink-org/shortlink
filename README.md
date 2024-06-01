<div align="center">

# shortlink 

ShortLink is an open-source educational project that provides a pretty user interface and respects GDPR. 

The goal of the project is to demonstrate the practical application of microservices architecture.

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/shortlink)](https://artifacthub.io/packages/search?repo=shortlink)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/shortlink-org/shortlink)](https://pkg.go.dev/mod/github.com/shortlink-org/shortlink)
[![codecov](https://codecov.io/gh/shortlink-org/shortlink/branch/main/graph/badge.svg?token=Wxz5bI4QzF)](https://codecov.io/gh/shortlink-org/shortlink)
[![Go Report Card](https://goreportcard.com/badge/github.com/shortlink-org/shortlink)](https://goreportcard.com/report/github.com/shortlink-org/shortlink)
[![Releases](https://img.shields.io/github/release-pre/shortlink-org/shortlink.svg)](https://github.com/shortlink-org/shortlink/releases)
[![LICENSE](https://img.shields.io/github/license/shortlink-org/shortlink.svg)](https://github.com/shortlink-org/shortlink/blob/main/LICENSE)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3510/badge)](https://bestpractices.coreinfrastructure.org/projects/3510)
[![StackShare](http://img.shields.io/badge/tech-stack-0690fa.svg?style=flat)](https://stackshare.io/shortlink-org/shortlink)
[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink.svg?type=shield)](https://app.fossa.com/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink?ref=badge_shield)
[![DeepSource](https://app.deepsource.com/gh/shortlink-org/shortlink.svg/?label=active+issues&show_trend=true&token=DL-zlqtnyx6CvlHCroG0Jdx5)](https://app.deepsource.com/gh/shortlink-org/shortlink/)

<hr />

<div style="align-items: center; display: flex;">
  <a href="https://www.producthunt.com/posts/shortlink-2?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-shortlink&#0045;2" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=374140&theme=light" alt="ShortLink - Get&#0032;ready&#0032;to&#0032;share&#0032;your&#0032;links&#0032;with&#0032;ease&#0033; | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>
  <img height="100px" src="https://slsa.dev/images/SLSA-Badge-full-level1.svg" alt="SLSA">
</div>

</div>
<hr />

### High Level Architecture 🚀

The project covers the entire process - from identifying Bounded Contexts to implementing microservices using
cutting-edge technologies and best practices.  

![shortlink-architecture](./docs/shortlink-architecture.png)
_Please [star ⭐](https://github.com/shortlink-org/shortlink/stargazers) the repo if you want us to continue developing and improving ShortLink! 😀_

### Boundaries

> [!TIP]
> 
> Our project follows Domain-Driven Design (DDD) principles, organizing code into distinct domains for clarity and easier updates.

| Bounded Context       | Description              | Type subdomain | Docs                                        |
|-----------------------|--------------------------|----------------|---------------------------------------------|
| API Gateway           | Gateway for all services | Supporting     | [docs](./boundaries/api/README.md)          |
| Auth Boundary         | Auth services            | Generic        | [docs](./boundaries/auth/README.md)         |
| Billing Boundary      | Payment services         | Generic        | [docs](./boundaries/billing/README.md)      |
| Chat Boundary         | Chat services            | Supporting     | [docs](./boundaries/chat/README.md)         |
| Delivery Boundary     | Delivery services        | Supporting     | [docs](./boundaries/delivery/README.md)     |
| DS Boundary           | Data Science services    | Supporting     | [docs](./boundaries/ds/README.md)           |
| Link Boundary         | Link services            | Core           | [docs](./boundaries/link/README.md)         |
| Marketing Boundary    | Marketing services       | Supporting     | [docs](./boundaries/marketing/README.md)    |
| Notification Boundary | Notification services    | Generic        | [docs](./boundaries/notification/README.md) |
| Platform Boundary     | Platform services        | Supporting     | [docs](./boundaries/platform/README.md)     |
| Search Boundary       | Search services          | Supporting     | [docs](./boundaries/search/README.md)       |
| Shop Boundary         | Shop services            | Supporting     | [docs](./boundaries/shop/README.md)         |
| ShortDB Boundary      | ShortDB services         | Supporting     | [docs](./boundaries/shortdb/README.md)      |
| UI Boundary           | UI services              | Supporting     | [docs](./boundaries/ui-monorepo/README.md)  |

> #### Contributing
>
> - [Getting Started](./CONTRIBUTING.md#getting-started)

### Architecture decision records (ADR)

> [!IMPORTANT]
> An architecture decision record (ADR) is a document that captures an important architecture decision 
made along with its context and consequences.
>
>+ [Docs ADR](https://github.com/joelparkerhenderson/architecture-decision-record)
>
> **Decisions:**
>  + [main decisions](./docs/ADR/README.md)
>  + [ops decisions](./ops/docs/ADR/README.md)
>
> Also, each boundary context and service has its own ADR. You can find them in the relevant sections.

### License

> [!WARNING]
> 
> This project includes dependencies licensed under the GNU Lesser General Public License (LGPL). 
> Users must comply with LGPL terms when using or modifying these dependencies. 
> For detailed information on each LGPL library used in this project, please refer to the respective license documentation 
> included with each library. For comprehensive license compliance information, including dependencies and their licenses, 
> you can read more details in our FOSSA report.

[![FOSSA Status](https://app.fossa.com/api/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink.svg?type=large)](https://app.fossa.com/projects/custom%2B396%2Fgithub.com%2Fshortlink-org%2Fshortlink?ref=badge_large)

[mergify]: https://mergify.io

[mergify-status]: https://img.shields.io/endpoint.svg?url=https://dashboard.mergify.io/badges/shortlink-org/shortlink&style=flat
