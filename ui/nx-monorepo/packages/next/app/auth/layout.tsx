import React from 'react'

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <div className="flex h-full p-4 rotate">
      <div className="sm:max-w-xl md:max-w-3xl w-full m-auto">{children}</div>
    </div>
  )
}
