import React from 'react';
import Grid from '@mui/material/Grid'
import Paper from '@mui/material/Paper'
import List from '@mui/material/List'

import { CatalogSearchResultListItem } from '@backstage/plugin-catalog';
import {
  catalogApiRef,
  CATALOG_FILTER_EXISTS,
} from '@backstage/plugin-catalog-react';
import { TechDocsSearchResultListItem } from '@backstage/plugin-techdocs';
import { AdrSearchResultListItem } from '@backstage-community/plugin-adr';
import { AdrDocument } from '@backstage-community/plugin-adr-common'

import { SearchType } from '@backstage/plugin-search';
import {
  DefaultResultListItem,
  SearchBar,
  SearchFilter,
  SearchResult,
  SearchPagination,
  useSearch,
} from '@backstage/plugin-search-react';
import {
  CatalogIcon,
  Content,
  DocsIcon,
  Header,
  Page,
} from '@backstage/core-components';
import { useApi } from '@backstage/core-plugin-api';

const SearchPage = () => {
  const { types } = useSearch();
  const catalogApi = useApi(catalogApiRef);

  return (
    <Page themeId="home">
      <Header title="Search" />
      <Content>
        <Grid container direction="row">
          <Grid item xs={12}>
            <Paper>
              <SearchBar />
            </Paper>
          </Grid>
          <Grid item xs={3}>
            <SearchType.Accordion
              name="Result Type"
              defaultValue="software-catalog"
              types={[
                {
                  value: 'software-catalog',
                  name: 'Software Catalog',
                  icon: <CatalogIcon />,
                },
                {
                  value: 'techdocs',
                  name: 'Documentation',
                  icon: <DocsIcon />,
                },
                {
                  value: 'adr',
                  name: 'Architecture Decision Records',
                  icon: <DocsIcon />,
                },
              ]}
            />
            <Paper>
              {types.includes('techdocs') && (
                <SearchFilter.Select
                  label="Entity"
                  name="name"
                  values={async () => {
                    // Return a list of entities which are documented.
                    const { items } = await catalogApi.getEntities({
                      fields: ['metadata.name'],
                      filter: {
                        'metadata.annotations.backstage.io/techdocs-ref':
                          CATALOG_FILTER_EXISTS,
                      },
                    });

                    const names = items.map(entity => entity.metadata.name);
                    names.sort();
                    return names;
                  }}
                />
              )}
              <SearchFilter.Select
                label="Kind"
                name="kind"
                values={['Component', 'Template']}
              />
              <SearchFilter.Checkbox
                label="Lifecycle"
                name="lifecycle"
                values={['experimental', 'production']}
              />
            </Paper>
          </Grid>
          <Grid item xs={9}>
            <SearchPagination />
            <SearchResult>
              {({ results }) => (
                <List>
                  {results.map(({ type, document, highlight, rank }) => {
                    switch (type) {
                      case 'software-catalog':
                        return (
                          <CatalogSearchResultListItem
                            key={document.location}
                            result={document}
                            highlight={highlight}
                            rank={rank}
                          />
                        );
                      case 'adr':
                        return (
                          <AdrSearchResultListItem
                            key={document.location} // @ts-ignore
                            result={document as AdrDocument}
                          />
                        );
                      case 'techdocs':
                        return (
                          <TechDocsSearchResultListItem
                            key={document.location}
                            result={document}
                            highlight={highlight}
                            rank={rank}
                          />
                        );
                      default:
                        return (
                          <DefaultResultListItem
                            key={document.location}
                            result={document}
                            highlight={highlight}
                            rank={rank}
                          />
                        );
                    }
                  })}
                </List>
              )}
            </SearchResult>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
};

export const searchPage = <SearchPage />;
