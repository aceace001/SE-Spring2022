const register = Cypress.env("");

describe("renders the home page",() => {
    beforeEach(() => {
        cy.visit("/");
        
    });
    it('renders correctly',() =>{
        cy.get("#container").should("exist");
    });
    it("allow user to register", () => {
        cy.visit("/");
        cy.findByText('Register').click();
        cy.get('#name').type('aaa');
        cy.get('#email').type('aabb@gmail.com');
        cy.get('#password').type('2233');
        cy.findByRole('button',{name: /Sign Up/i}).click();
        cy.get('#email').type('aabb@gmail.com');
        cy.get('#password').type('2233');
        cy.findByRole('button',{name: /Login/i}).click()
        cy.findByText("Post Page").should("exist");

    });
    it("allows the email and password fields to be used", () => {
        cy.visit("/")
        cy.get('#email').type('hhhh@gmail.com')
        cy.get('#password').type('12345')
        cy.findByRole('button',{name: /Login/i}).click()
        cy.findByText("Post Page").should("exist");
    })
});