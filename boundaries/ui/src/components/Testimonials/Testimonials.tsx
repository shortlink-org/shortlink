'use client'

const testimonials = [
  {
    name: 'Viktor Login',
    handle: '@batazor',
    avatar: 'https://www.tailwind-kit.com/images/person/1.jpg',
    text: 'Shortlink has transformed how I manage my URLs. The analytics are incredibly detailed and the interface is intuitive.',
    rating: 5,
  },
  {
    name: 'Alex Chen',
    handle: '@alexc',
    avatar: 'https://www.tailwind-kit.com/images/person/2.jpg',
    text: 'Finally a URL shortener that respects privacy and gives you full control. Open source for the win!',
    rating: 5,
  },
  {
    name: 'Sarah Miller',
    handle: '@sarahm',
    avatar: 'https://www.tailwind-kit.com/images/person/3.jpg',
    text: 'The API is fantastic for automation. I integrated it into my workflow in minutes.',
    rating: 5,
  },
]

function StarIcon({ className }: { className?: string }) {
  return (
    <svg className={className} fill="currentColor" viewBox="0 0 20 20">
      <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
    </svg>
  )
}

function QuoteIcon({ className }: { className?: string }) {
  return (
    <svg className={className} fill="currentColor" viewBox="0 0 24 24">
      <path d="M14.017 21v-7.391c0-5.704 3.731-9.57 8.983-10.609l.995 2.151c-2.432.917-3.995 3.638-3.995 5.849h4v10h-9.983zm-14.017 0v-7.391c0-5.704 3.748-9.57 9-10.609l.996 2.151c-2.433.917-3.996 3.638-3.996 5.849h3.983v10h-9.983z" />
    </svg>
  )
}

export default function Testimonials() {
  return (
    <section className="py-16 lg:py-24">
      <div className="text-center mb-12">
        <span className="text-xs uppercase tracking-widest font-semibold text-indigo-600 dark:text-indigo-400">Testimonials</span>
        <h2 className="text-3xl lg:text-4xl font-bold text-gray-900 dark:text-white mt-2">Loved by developers</h2>
        <p className="text-gray-600 dark:text-gray-400 mt-3 max-w-2xl mx-auto">See what our users have to say about Shortlink</p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        {testimonials.map((testimonial, index) => (
          <div
            key={index}
            className="relative overflow-visible bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-2xl p-6 shadow-sm hover:shadow-xl hover:border-indigo-300 dark:hover:border-indigo-600 transition-all duration-300 hover:-translate-y-1"
          >
            {/* Quote Icon */}
            <div className="absolute -top-4 left-6">
              <div className="w-10 h-10 rounded-full bg-gradient-to-br from-indigo-500 to-purple-600 flex items-center justify-center shadow-lg">
                <QuoteIcon className="w-4 h-4 text-white" />
              </div>
            </div>

            {/* Rating */}
            <div className="flex gap-0.5 mt-4 mb-4">
              {[...Array(testimonial.rating)].map((_, i) => (
                <StarIcon key={i} className="w-5 h-5 text-yellow-400" />
              ))}
            </div>

            {/* Testimonial Text */}
            <p className="text-gray-700 dark:text-gray-300 mb-6 leading-relaxed">"{testimonial.text}"</p>

            {/* Author */}
            <div className="flex items-center gap-3 pt-4 border-t border-gray-100 dark:border-gray-700">
              <img
                src={testimonial.avatar}
                alt={testimonial.name}
                className="w-12 h-12 rounded-full ring-2 ring-indigo-100 dark:ring-indigo-900 object-cover"
              />
              <div>
                <p className="font-semibold text-gray-900 dark:text-white text-sm">{testimonial.name}</p>
                <p className="text-xs text-indigo-600 dark:text-indigo-400">{testimonial.handle}</p>
              </div>
            </div>
          </div>
        ))}
      </div>
    </section>
  )
}
