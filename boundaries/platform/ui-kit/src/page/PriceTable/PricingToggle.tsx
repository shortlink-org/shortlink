import React from 'react'

type PricingToggleProps = {
  isAnnual: boolean
  setIsAnnual: (isAnnual: boolean) => void
}

const PricingToggle: React.FC<PricingToggleProps> = ({
  isAnnual,
  setIsAnnual,
}) => (
  <div className="flex justify-center max-w-[14rem] m-auto mb-8 lg:mb-16">
    <div className="relative flex w-full p-1 bg-white dark:bg-slate-900 rounded-full">
      <span
        className="absolute inset-0 m-1 pointer-events-none"
        aria-hidden="true"
      >
        <span
          className={`absolute inset-0 w-1/2 bg-indigo-500 rounded-full shadow-sm shadow-indigo-950/10 transform transition-transform duration-150 ease-in-out ${
            isAnnual ? 'translate-x-0' : 'translate-x-full'
          }`}
        />
      </span>
      <button
        type="button"
        className={`relative flex-1 text-sm font-medium h-8 rounded-full focus-visible:outline-none focus-visible:ring focus-visible:ring-indigo-300 dark:focus-visible:ring-slate-600 transition-colors duration-150 ease-in-out ${
          isAnnual
            ? 'text-white'
            : 'text-slate-500 dark:text-slate-400 hover:cursor-pointer'
        }`}
        onClick={() => setIsAnnual(true)}
      >
        Yearly
        <span
          className={
            isAnnual ? 'text-indigo-200' : 'text-slate-400 dark:text-slate-500'
          }
        >
          -20%
        </span>
      </button>
      <button
        type="button"
        className={`relative flex-1 text-sm font-medium h-8 rounded-full focus-visible:outline-none focus-visible:ring focus-visible:ring-indigo-300 dark:focus-visible:ring-slate-600 transition-colors duration-150 ease-in-out ${
          isAnnual
            ? 'text-slate-500 dark:text-slate-400 hover:cursor-pointer'
            : 'text-white'
        }`}
        onClick={() => setIsAnnual(false)}
      >
        Monthly
      </button>
    </div>
  </div>
)

export default PricingToggle
