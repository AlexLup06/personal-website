window.htmx.onLoad(() => {
    const leftButton = document.getElementById("left-button") as HTMLButtonElement
    const rightButton = document.getElementById("right-button") as HTMLButtonElement

    const stackContainer = document.getElementById("stack-container") as HTMLDivElement

    const scrollAmount = 240

    leftButton.addEventListener("click", () => stackContainer.scrollBy({ left: -scrollAmount, behavior: "smooth" }))
    rightButton.addEventListener("click", () => stackContainer.scrollBy({ left: scrollAmount, behavior: "smooth" }))

})