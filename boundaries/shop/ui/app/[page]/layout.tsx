import Footer from 'components/layout/footer';

// DOCS: https://nextjs.org/docs/app/api-reference/file-conventions/route-segment-config#experimental_ppr
export const experimental_ppr = true

// DOCS: https://nextjs.org/docs/app/api-reference/file-conventions/route-segment-config#dynamic
export const dynamic = 'force-dynamic'

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <>
      <div className="w-full">
        <div className="mx-8 max-w-2xl py-20 sm:mx-auto">{children}</div>
      </div>
      <Footer />
    </>
  );
}
