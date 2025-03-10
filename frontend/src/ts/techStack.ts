window.htmx.onLoad(() => {
    const leftButton = document.getElementById("left-button") as HTMLButtonElement
    const rightButton = document.getElementById("right-button") as HTMLButtonElement
    const stackContainer = document.getElementById("stack-container") as HTMLDivElement
    const scrollContainer = document.getElementById("scroll-stack-container") as HTMLDivElement
    const stackElement = stackContainer.children[0] as HTMLDivElement

    if (leftButton == null || rightButton == null || stackContainer == null) {
        return
    }

    const scrollAmount = stackElement.getBoundingClientRect().width

    leftButton.addEventListener("click", () => handleLeftClick())
    rightButton.addEventListener("click", () => handleRightClick())

    function handleLeftClick() {
        if (scrollContainer.scrollLeft - scrollAmount < 0) {
            scrollContainer.scrollTo({
                top: 0,
                left: 0,
                behavior: "smooth",
            })
        } else {
            scrollContainer.scrollBy({ left: -scrollAmount, behavior: "smooth" })
        }
    }

    function handleRightClick() {
        const stackContainerWidth = stackContainer.getBoundingClientRect().width

        if (scrollContainer.scrollLeft == 0) {
            scrollContainer.scrollBy({ left: 100, behavior: "smooth" })
            return
        }

        if (scrollContainer.scrollLeft + scrollAmount > stackContainerWidth) {
            scrollContainer.scrollTo({
                top: 0,
                left: stackContainerWidth,
                behavior: "smooth",
            })
        } else {
            scrollContainer.scrollBy({ left: scrollAmount, behavior: "smooth" })
        }
    }

})