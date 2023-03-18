// @ts-ignore
describe('Landing: Tabs', () => {
  beforeEach(() => {
    cy.visit('/')
  })

  it('Click toggle switch', () => {
    cy.get('#full-width-tab-0').click()
    cy.get('#full-width-tab-1').click()
    cy.get('#full-width-tab-2').click()
    cy.get('#full-width-tab-3').click()
  })
})
