// @ts-ignore
describe('Landing: Tabs', () => {
  beforeEach(() => {
    cy.visit('/')
  })

  it('Menu: macbook-16', () => {
    cy.viewport('macbook-16') // 1536 x 960

    cy.screenshot()
    cy.get('#full-width-tab-0').click()

    cy.screenshot()
    cy.get('#full-width-tab-1').click()

    cy.screenshot()
    cy.get('#full-width-tab-2').click()

    cy.screenshot()
    cy.get('#full-width-tab-3').click()

    cy.screenshot()
  })
})
