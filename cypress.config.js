const { defineConfig } = require("cypress");

module.exports = defineConfig({
  e2e: {
    supportFile: false,
    baseUrl: "http://localhost:8080",
    specPattern: "cypress/e2e/**/*.cy.js",
  },

  setupNodeEvents(on, config) {
    config.env.sharedSecret =
      process.env.NODE_ENV === "qa" ? "hoop brick tort" : "sushi cup lemon";

    return config;
  },
});
