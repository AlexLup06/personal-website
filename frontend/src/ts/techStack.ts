window.htmx.onLoad(() => {
    const leftButton = document.getElementById("left-button") as HTMLButtonElement
    const rightButton = document.getElementById("right-button") as HTMLButtonElement
    const container = document.getElementById("stack-container") as HTMLDivElement

    if (leftButton == null || rightButton == null || container == null) {
        return
    }

    const scrollAmount = 240

    leftButton.addEventListener("click", () => handleLeftClick())
    rightButton.addEventListener("click", () => handleRightClick())

    function handleLeftClick() {
        if (container.scrollLeft - 240 < 0) {
            container.scrollTo({
                top: 0,
                left: 0,
                behavior: "smooth",
            })
        } else {
            container.scrollBy({ left: -scrollAmount, behavior: "smooth" })
        }
    }

    function handleRightClick() {
        const containerWidth = container.getBoundingClientRect().width
        if (container.scrollLeft + 240 > containerWidth) {
            container.scrollTo({
                top: 0,
                left: containerWidth,
                behavior: "smooth",
            })
        } else {
            container.scrollBy({ left: scrollAmount, behavior: "smooth" })
        }
    }

})