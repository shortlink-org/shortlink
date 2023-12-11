/// <reference  types="cypress" />
import './commands'

type Method = 'POST' | 'GET' | 'DELETE'

declare global {
  namespace Cypress {
    interface Chainable {
      dataCy(value: string): Chainable<Element>
      interceptRequest(method: Method): Chainable<null>
    }
  }
}
