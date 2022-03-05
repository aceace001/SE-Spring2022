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
        cy.get('#first').type('aaa');
        cy.get('#last').type('bbb');
        cy.get('#email').type('aabb@gmail.com');
        cy.get('#password').type('2233');
        cy.findByRole('button',{name: /Sign Up/i}).click();
        cy.get('#email').type('aabb@gmail.com');
        cy.get('#password').type('2233');
        cy.findByRole('button',{name: /Login/i}).click()
        cy.findByText("Wrong username or password").should("exist");

    });
    it("allows the email and password fields to be used", () => {
        cy.visit("/")
        cy.get('#email').type('hhhh@gmail.com')
        cy.get('#password').type('12345')
        cy.findByRole('button',{name: /Login/i}).click()
        cy.findByText("Wrong username or password").should("exist");
    });
    it("allows the users type in the chat box", () =>{
        cy.visit("/messages")
        cy.get('#chatcontent').type('cool!')

    });
    it("allows the users to click send  button after the type", () =>{
        cy.visit("/messages")
        cy.get('#chatcontent').type('check click send button this time')
        cy.findByRole('button',{name: /Send/i}).click();
    });
});