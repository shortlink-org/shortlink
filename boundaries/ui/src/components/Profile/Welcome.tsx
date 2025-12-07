interface WelcomeProps {
  nickname: string
}

export default function Welcome({ nickname }: WelcomeProps) {
  return (
    <div className="md:w-auto flex items-center content-center my-6 flex-auto bg-purple-700 dark:bg-indigo-500 text-white rounded shadow-xl px-5 w-full">
      <div className="flex flex-wrap content-center items-center">
        <div className="w-1/4 px-3 text-center hidden md:block">
          <div className="p-5 xl:px-8 md:py-5">
            <img 
              src="/assets/images/undraw_welcome_cats_thqn.svg" 
              alt="Welcome illustration" 
              className="w-full h-auto" 
            />
          </div>
        </div>
        <div className="w-full sm:w-1/2 md:w-3/4 px-3 text-left">
          <div className="p-5 xl:px-8 md:py-5">
            <h3 className="text-2xl font-bold">Welcome, {nickname}!</h3>
            <p className="text-sm text-indigo-200 mt-2">
              Manage your profile settings, update your personal information, and customize your notification preferences.
            </p>
          </div>
        </div>
      </div>
    </div>
  )
}
