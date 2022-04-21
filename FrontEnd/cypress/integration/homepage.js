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
        cy.findByText('Register').click({force: true});
        cy.get('#name').type('aaa');
        cy.get('#email').type('aabb@gmail.com');
        cy.get('#password').type('2233');
        cy.findByRole('button',{name: /Sign Up/i}).click();
        

    });
    it("allows the email and password fields to be used", () => {
        cy.visit("/")
        cy.get('#email')
        cy.get('#password')
        cy.findByRole('button',{name: /Login/i}).click({force:true});
    });
    it("allows the users to type in the message page",()=>{
        cy.visit("/messages")
        cy.get(".ChatInput").type("aabbcc").type('{enter}')
    });
    it("allows the users to type username and tweet in the post page",()=>{
        cy.visit("/posts")
        cy.get('#first').type('hhh')
        cy.get('#last').type("This is a tweet")
    });
    it("allows the users to click on the Post Now button",()=>{
        cy.visit("/posts")
        cy.findByRole('button',{name: /Post Now!/i}).click();
    });
    it("allows the users to hoverover the chatbot avatar", ()=>{
        cy.visit("/")
        cy.get('#avatar').invoke('show')
    });
    it("allows the user to hoverover on the chatbot avatar and sometging become visible", ()=>{
        cy.visit("/")
        cy.get('#avatar').trigger('mouseover')
        cy.get('.transition-3').should('be.visible');
    });
    it("allows the user to click the chatbot button and email form visible", ()=>{
        cy.visit("/")
        cy.findByRole('button',{name: /Clickme/i}).click();

    });
    it("allows the user to enter email in the email form",()=>{
        cy.visit("/")
        cy.findByRole('button',{name: /Clickme/i}).click();
        cy.get('.semail').type('bob@mail.com');
    });
    it("allows send the email after type and show the chat window", ()=>{
        cy.visit("/")
        cy.findByRole('button',{name: /Clickme/i}).click();
        cy.get('.semail').type('bob@mail.com').type('{enter}');
    });
    it("allows send the message in the chat", ()=>{
        cy.visit("/")
        cy.findByRole('button',{name: /Clickme/i}).click();
        cy.get('.semail').type('bob@mail.com').type('{enter}');
        cy.get('.quill').type('hello').type('{enter}').type('{enter}');
    });
    it("allows to see the chat with customer service in the support page", ()=>{
        cy.visit("/support")
        //The chat is in charge of by the other API and it is out of my control
        //So, it is very hard and not recommended to do the cypress test
    });
    it("allows the users to search for frineds", ()=>{
        cy.visit("/search")
        cy.get('.searchterm').type('Abby')
        cy.contains("Abby")
    });
    it("allows the users to see all friends", ()=>{
        cy.visit("/users")
        cy.findByRole('button',{name: /Clear all/i}).click();
        cy.contains('0')
        cy.findByRole('button', {name: /Get all/i}).click();
        cy.contains('300')
    });
});
