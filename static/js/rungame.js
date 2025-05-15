document.addEventListener("DOMContentLoaded", () => {
  const buttons = document.querySelectorAll(".game-btn");
  const result = document.getElementById("result");

  const request = async (choice = "") => {
    try {
      const res = await fetch("/play", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          jugada: choice,
        },
      });

      const data = await res.json();
      alert(`Resultado: ${data.resultado}`);
    } catch (err) {
      console.error("Error:", err);
    }
  };

  buttons.forEach((button) => {
    button.addEventListener("click", () => {
      const choice = button.getAttribute("data-value");
      result.textContent = `Elegiste: ${choice.toUpperCase()}`;
      console.log("Opción seleccionada:", choice);
      request(choice);
    });
  });
});
