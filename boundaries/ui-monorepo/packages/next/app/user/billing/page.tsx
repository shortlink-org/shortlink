'use client'

import Discounted from 'components/Billing/Discounted'
import PaymentMethod from 'components/Billing/PaymentMethod'
import withAuthSync from 'components/Private'
import Header from 'components/Page/Header'

// <ProductJsonLd
// productName="Shortlink"
// description="Shortlink service"
// brand="Shortlink"
// color="premium"
// manufacturerLogo="https://shortlink.best/images/logo.png"
// manufacturerName="Shortlink"
// material="bright"
// slogan="Shortlink service"
// image="https://shortlink.best/images/logo.png"
// url="https://shortlink.best/"
// award="Best service"
// reviews={[
//     {
//       author: 'Jim',
//       datePublished: '2017-01-06T03:37:40Z',
//       reviewBody: 'This is my favorite product yet! Thanks Nate for the example products and reviews.',
//       name: 'So awesome!!!',
//       reviewRating: {
//         bestRating: '5',
//         ratingValue: '5',
//         worstRating: '1',
//       },
//       publisher: {
//         type: 'Organization',
//         name: 'TwoVit',
//       },
//     },
// ]}
// aggregateRating={{
//   ratingValue: '5',
//     reviewCount: '89',
// }}
// offers={[
//     {
//       price: '119.99',
//       priceCurrency: 'USD',
//       priceValidUntil: '2020-11-05',
//       itemCondition: 'https://schema.org/UsedCondition',
//       availability: 'https://schema.org/InStock',
//       url: 'https://www.example.com/executive-anvil',
//       seller: {
//         name: 'Executive Objects',
//       },
//     },
// {
//   price: '139.99',
//     priceCurrency: 'CAD',
//   priceValidUntil: '2020-09-05',
//   itemCondition: 'https://schema.org/UsedCondition',
//   availability: 'https://schema.org/InStock',
//   url: 'https://www.example.ca/executive-anvil',
//   seller: {
//   name: 'Executive Objects',
// },
// },
// ]}
// mpn="925872"
// sku="0446310786"
// gtin13="9780446310789"
// gtin8="0446310786"
//   />

function Page() {
  return (
    <>
      {/*<NextSeo title="Billing" description="Billing page for your account." />*/}

      <Header title="Billing" />

      <Discounted />

      <PaymentMethod />
    </>
  )
}

export default withAuthSync(Page)
