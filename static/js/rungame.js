document.addEventListener("DOMContentLoaded", () => {
  const buttons = document.querySelectorAll(".game-btn");
  const result = document.getElementById("result");
  const computer = document.getElementById("computer");
  const gameResult = document.getElementById("gameResult");

  const historyList = document.getElementById("history");
  const winsSpan = document.getElementById("wins");
  const lossesSpan = document.getElementById("losses");

  let wins = 0;
  let losses = 0;

  const request = async (choice = "") => {
    try {
      const res = await fetch("/play", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          jugada: choice,
        },
      });

      return await res.json();
    } catch (err) {
      console.error("Error:", err);
    }
  };

  buttons.forEach((button) => {
    button.addEventListener("click", () => {
      const choice = button.getAttribute("data-value");
      result.textContent = `Elegiste: ${choice.toUpperCase()}`;
      console.log("Opción seleccionada:", choice);

      request(choice).then((res) => {
        if (!res) return;

        const { isPlayerVictory, resultado, isdraw, jugadaPC } = res;
        console.log(isdraw);

        computer.innerHTML = resultado;
        gameResult.innerHTML = isdraw
          ? "Empate"
          : isPlayerVictory
          ? "Victoria"
          : "Derrota";

        if (isPlayerVictory && !isdraw) {
          wins++;
          winsSpan.textContent = wins;
        }
        if (!isPlayerVictory && !isdraw) {
          losses++;
          lossesSpan.textContent = losses;
        }

        const li = document.createElement("li");
        li.classList.add(
          "list-group-item",
          "d-flex",
          "justify-content-between"
        );
        li.innerHTML = `
          <span><strong>Vos:</strong> ${choice}</span>
          <span><strong>PC:</strong> ${jugadaPC || "?"}</span>
        `;
        historyList.prepend(li);
      });
    });
  });
});
