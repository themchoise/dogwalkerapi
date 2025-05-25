describe("Mi primera prueba", () => {
  it("Visita una página y verifica el título", () => {
    cy.visit("http://localhost:8080/hello");
    cy.title().should("include", "Example Domain");
  });
});
