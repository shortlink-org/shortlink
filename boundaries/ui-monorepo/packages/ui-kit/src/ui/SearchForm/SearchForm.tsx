import SearchIcon from '@mui/icons-material/Search'
import MicIcon from '@mui/icons-material/Mic'

function SearchForm() {
  return (
    <div className="relative flex-grow">
      <button
        type="button"
        className="absolute left-2 top-1/2 transform -translate-y-1/2 p-1 text-gray-500 hover:text-gray-700"
      >
        <SearchIcon className="h-5 w-5" />
      </button>
      <input
        id="search"
        type="search"
        className="w-full pl-10 p-2 text-sm text-gray-700 form-input bg-white border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
        placeholder="Search..."
      />
      <button
        type="submit"
        className="absolute right-2 top-1/2 transform -translate-y-1/2 p-1 text-gray-500 hover:text-gray-700"
      >
        <MicIcon />
      </button>
    </div>
  )
}

export default SearchForm
