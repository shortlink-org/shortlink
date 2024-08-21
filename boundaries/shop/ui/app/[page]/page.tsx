import type { Metadata } from 'next';

import Prose from 'components/prose';
import type { Page } from 'lib/shopify/types';
import { notFound } from 'next/navigation';

export async function generateMetadata({
  params
}: {
  params: { page: string };
}): Promise<Metadata> {
  // const page: Page = await getPage(params.page);
  const page: Page = {
    id: "1",
    title: "Sample Title",
    handle: "sample-handle",
    body: "Sample Body",
    bodySummary: "Sample Body Summary",
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
    seo: {
      title: "Sample SEO Title",
      description: "Sample SEO Description"
    }
  };

  if (!page) return notFound();

  return {
    title: page.seo?.title || page.title,
    description: page.seo?.description || page.bodySummary,
    openGraph: {
      publishedTime: page.createdAt,
      modifiedTime: page.updatedAt,
      type: 'article'
    }
  };
}

export default async function Page({ params }: { params: { page: string } }) {
  // const page = await getPage(params.page);
  const page: Page = {
    id: "1",
    title: "Sample Title",
    handle: "sample-handle",
    body: "Sample Body",
    bodySummary: "Sample Body Summary",
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
    seo: {
      title: "Sample SEO Title",
      description: "Sample SEO Description"
    }
  };

  if (!page) return notFound();

  return (
    <>
      <h1 className="mb-8 text-5xl font-bold">{page.title}</h1>
      <Prose className="mb-8" html={page.body as string} />
      <p className="text-sm italic">
        {`This document was last updated on ${new Intl.DateTimeFormat(undefined, {
          year: 'numeric',
          month: 'long',
          day: 'numeric'
        }).format(new Date(page.updatedAt))}.`}
      </p>
    </>
  );
}
