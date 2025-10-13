# ShortLink UI - Routing Documentation

## Overview

This document describes the routing architecture for the ShortLink UI service, which is a static Next.js application served through Kubernetes Nginx Ingress and containerized Nginx.

## Architecture Components

```
Browser → Nginx Ingress Controller → Nginx Pod → Static Files
```

## Routing Flow

```
Browser: https://shortlink.best/next/add-link
    ↓
[Nginx Ingress Controller]
    • Receives: /next/add-link  
    • Regex match: /next(/|$)(.*)
    • Rewrite: /$2 → /add-link
    • Proxy to: shortlink-link-ui:8080/add-link
    • Intercepts redirects with port and removes it
    ↓
[Nginx in Pod (ui.local)]
    • Receives: /add-link
    • Tries files: /usr/share/nginx/html/add-link.html
    • Returns: HTML page
    ↓
Browser receives page ✅
```
