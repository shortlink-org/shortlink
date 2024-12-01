import { liteClient as algoliasearch } from 'algoliasearch/lite';
import { InstantSearch, SearchBox, RefinementList } from 'react-instantsearch';

const searchClient = algoliasearch('YourApplicationID', 'YourSearchOnlyAPIKey')

function SearchForm() {
  return (
    <InstantSearch
      searchClient={searchClient}
      future={{
        preserveSharedStateOnUnmount: true,
      }}
      indexName="instant_search"
      insights
    >
      <SearchBox
        placeholder="Product, brand, color, …"
        classNames={{
          root: 'p-3 shadow-sm',
          form: 'relative',
          input: 'block w-full pl-9 pr-3 py-2 bg-white border border-slate-300 placeholder-slate-400 focus:outline-none focus:border-sky-500 focus:ring-sky-500 rounded-md focus:ring-1',
          submitIcon: 'absolute top-0 left-0 bottom-0 w-6',
        }}
      />
      <RefinementList
        attribute="categories"
      />
    </InstantSearch>
  )
}

export default SearchForm
