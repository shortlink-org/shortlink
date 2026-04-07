# ShortLink UI - Routing Documentation

## Overview

This document describes the routing architecture for the ShortLink UI service: a **Next.js standalone** app (Node) fronted by Kubernetes Ingress.

## Architecture Components

```
Browser → Ingress → link-ui Pod (Next.js on :8080)
```

## Routing Flow

```
Browser: https://shortlink.best/next/add-link
    ↓
[Nginx Ingress Controller]
    • Receives: /next/add-link  
    • Regex match: /next(/|$)(.*)
    • Rewrite: /$2 → /add-link
    • Proxy to: shortlink-link-ui:8080/next/add-link (strip prefix per ingress rules)
    ↓
[Next.js server in Pod]
    • Serves App Router routes under basePath `/next`
    ↓
Browser receives page ✅
```
