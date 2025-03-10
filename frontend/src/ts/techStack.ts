window.htmx.onLoad(() => {
    const leftButton = document.getElementById("left-button") as HTMLButtonElement
    const rightButton = document.getElementById("right-button") as HTMLButtonElement
    const stackContainer = document.getElementById("stack-container") as HTMLDivElement
    const scrollContainer = document.getElementById("scroll-stack-container") as HTMLDivElement
    const stackElement = stackContainer.children[0] as HTMLDivElement

    if (leftButton == null || rightButton == null || stackContainer == null) {
        return
    }

    const stackElementWidth = stackElement.getBoundingClientRect().width
    const scrollAmount = stackElementWidth
    const gap = Number(window.getComputedStyle(stackContainer).gap.slice(0, -2))

    console.log(gap)

    if (!gap) {
        return
    }


    leftButton.addEventListener("click", () => handleLeftClick())
    rightButton.addEventListener("click", () => handleRightClick())

    function handleLeftClick() {
        const stackContainerWidth = stackContainer.getBoundingClientRect().width
        const scrollContainerWidth = scrollContainer.getBoundingClientRect().width

        if (stackContainerWidth - scrollContainerWidth == scrollContainer.scrollLeft) {
            scrollContainer.scrollBy({ left: -stackElementWidth / 2 - gap, behavior: "smooth" })
            return
        }

        if (scrollContainer.scrollLeft - scrollAmount < 0) {
            scrollContainer.scrollTo({
                top: 0,
                left: 0,
                behavior: "smooth",
            })
        } else {
            scrollContainer.scrollBy({ left: -scrollAmount - gap, behavior: "smooth" })
        }
    }

    function handleRightClick() {
        const stackContainerWidth = stackContainer.getBoundingClientRect().width

        if (scrollContainer.scrollLeft == 0) {
            scrollContainer.scrollBy({ left: stackElementWidth / 2 + gap, behavior: "smooth" })
            return
        }



        if (scrollContainer.scrollLeft + scrollAmount > stackContainerWidth) {
            scrollContainer.scrollTo({
                top: 0,
                left: stackContainerWidth,
                behavior: "smooth",
            })
        } else {
            scrollContainer.scrollBy({ left: scrollAmount + gap, behavior: "smooth" })
        }

    }

})