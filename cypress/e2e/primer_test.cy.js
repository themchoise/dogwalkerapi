describe("Juego Piedra, Papel o Tijera", () => {
  beforeEach(() => {
    cy.visit("/rungame");
  });

  it("Debería cargar correctamente los elementos", () => {
    cy.contains("Escoge tu jugada");
    cy.get(".game-btn").should("have.length", 3);
    cy.get("#wins").should("contain", "0");
    cy.get("#losses").should("contain", "0");
  });

  it("Debería permitir jugar y mostrar el resultado", () => {
    cy.intercept("GET", "/play", {
      statusCode: 200,
      body: {
        isPlayerVictory: true,
        resultado: "Ganaste! PC eligió Tijera.",
        isdraw: false,
        jugadaPC: "Tijera",
      },
    }).as("jugadaMock");

    cy.get('[data-value="Piedra"]').click();
    cy.wait("@jugadaMock");

    cy.get("#result").should("contain", "Elegiste");
    cy.get("#computer").should("contain", "Ganaste!");
    cy.get("#gameResult").should("contain", "Victoria");
    cy.get("#wins").should("contain", "1");
  });

  it("Debería actualizar historial con jugadas", () => {
    cy.intercept("GET", "/play", {
      statusCode: 200,
      body: {
        isPlayerVictory: false,
        resultado: "Perdiste! PC eligió Papel.",
        isdraw: false,
        jugadaPC: "Papel",
      },
    }).as("jugadaMock");

    cy.get('[data-value="Piedra"]').click();
    cy.wait("@jugadaMock");

    cy.get("#history li")
      .first()
      .within(() => {
        cy.contains("Vos: Piedra");
        cy.contains("PC: Papel");
      });

    cy.get("#losses").should("contain", "1");
  });

  it("Debería mostrar empate correctamente", () => {
    cy.intercept("GET", "/play", {
      statusCode: 200,
      body: {
        isPlayerVictory: false,
        resultado: "La Pc Jugo: Papel",
        isdraw: true,
        jugadaPC: "Papel",
      },
    }).as("jugadaMock");

    cy.get('[data-value="Papel"]').click();
    cy.wait("@jugadaMock");

    cy.get("#gameResult").should("contain", "Empate");
    cy.get("#wins").should("contain", "0");
    cy.get("#losses").should("contain", "0");
  });

  it("Debería jugar 30 veces con 1 segundo de espera entre jugadas", () => {
    const jugadas = ["Piedra", "Papel", "Tijera"];

    Cypress._.times(30, (i) => {
      const jugada = jugadas[Math.floor(Math.random() * jugadas.length)];
      cy.log(`Jugando #${i + 1}: ${jugada}`);
      cy.get(`[data-value="${jugada}"]`).click();
      cy.wait(1000);
    });

    cy.get("#history li").should("have.length.at.least", 30);
  });
});
