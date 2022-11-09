import { imageLoader } from 'next-image-loader/image-loader'

// write self-define a custom loader
// (resolverProps: { src: string; width: number; quality?: number }) => string
imageLoader.loader = ({ src, width, quality }) =>
  `${process.env.NEXT_PUBLIC_OPTIMIZE_DOMAIN}${src}`
