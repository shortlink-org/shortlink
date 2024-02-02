## Extension Boundary

> [!NOTE]
> The Extension Boundary in software architecture is focused on enhancing and customizing existing systems through 
> the integration of additional modules or features. It typically includes a variety of extensions, plugins, 
> and integrations, such as Chrome Extensions for browser functionality, ArgoCD for Kubernetes deployment, and OpenAI 
> integrations for AI capabilities. This boundary is key for expanding the functionality of a core application 
> without altering its base code, allowing for tailored enhancements that can adapt to specific user needs or 
> technological advancements. It's essential for flexibility and scalability in software development, 
> enabling continuous improvement and customization.

| Service               | Description          | Language/Framework | Docs                                                                         |
|-----------------------|----------------------|--------------------|------------------------------------------------------------------------------|
| chrome-extension      | Chrome extension     | JavaScript         | [docs](./chrome-extension/README.md)                                         |                                                                       
| ai-plugin             | ChatGPT plugin       | JSON               | [docs](../ui/nx-monorepo/packages/landing/public/.well-known/ai-plugin.json) |
| argocd-extension-docs | ArgoCD extension     | JavaScript         | [docs](./argocd-extension-docs/README.md)                                    |

### Docs

- [GLOSSARY.md](./GLOSSARY.md) - Ubiquitous Language of the Extension Boundary
