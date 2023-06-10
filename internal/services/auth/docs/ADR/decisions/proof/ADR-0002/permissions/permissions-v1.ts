// @ts-nocheck
import { Namespace, Context } from "@ory/keto-namespace-types"

class User implements Namespace {}

class Link implements Namespace {
  related: {
    owners: User[]
    editors: User[]
    viewers: User[]
  }

  permits = {
    view: (ctx: Context): boolean =>
      this.related.viewers.includes(ctx.subject) ||
      this.permits.edit(ctx),

    edit: (ctx: Context): boolean =>
      this.related.editors.includes(ctx.subject) ||
      this.permits.share(ctx),

    delete: (ctx: Context): boolean =>
      this.permits.share(ctx),

    share: (ctx: Context): boolean =>
      this.related.owners.includes(ctx.subject),
  }
}
