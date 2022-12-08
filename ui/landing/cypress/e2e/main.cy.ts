describe('E2E Test', () => {
    beforeEach(() => {
        cy.visit('/')
    })
    it('Click toggle switch', () => {
        cy.get('.toggle__handler').click()
    })
})