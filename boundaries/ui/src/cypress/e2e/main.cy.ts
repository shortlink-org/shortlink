// @ts-ignore
describe('E2E Test', () => {
  beforeEach(() => {
    cy.visit('/auth/login')
  })
  it('Click toggle switch', () => {
    // The new url should include "/about"
    cy.url().should('include', '/auth/login')
  })
})
