import 'next'

declare global {
  namespace NodeJS {
    interface ProcessEnv {
      NEXT_PUBLIC_PIPELINE_ID: string
      NEXT_PUBLIC_CI_PIPELINE_URL: string
    }
  }
}
