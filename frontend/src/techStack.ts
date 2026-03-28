window.htmx.onLoad(() => {
    const leftButton = document.getElementById("left-button") as HTMLButtonElement
    const rightButton = document.getElementById("right-button") as HTMLButtonElement
    const stackContainer = document.getElementById("stack-container") as HTMLDivElement
    const scrollContainer = document.getElementById("scroll-stack-container") as HTMLDivElement

    if (leftButton == null || rightButton == null || stackContainer == null) {
        return
    }


    leftButton.addEventListener("click", () => handleLeftClick())
    rightButton.addEventListener("click", () => handleRightClick())

    function handleLeftClick() {
        const stackElement = stackContainer.children[0] as HTMLDivElement
        const stackElementWidth = stackElement.getBoundingClientRect().width
        const scrollContainerWidth = scrollContainer.getBoundingClientRect().width

        const scrollAmount = stackElementWidth
        const gap = Number(window.getComputedStyle(stackContainer).gap.slice(0, -2))

        const isAtFarRight = scrollContainer.scrollWidth - scrollContainer.clientWidth == scrollContainer.scrollLeft

        if (isAtFarRight) {
            const fullElements = Math.floor(scrollContainerWidth / stackElementWidth)
            const partialWidthInside = scrollContainerWidth - fullElements * stackElementWidth - fullElements * gap
            const partialWidthOutisde = stackElementWidth - partialWidthInside
            scrollContainer.scrollBy({ left: -partialWidthOutisde, behavior: "smooth" })
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
        const stackElement = stackContainer.children[0] as HTMLDivElement
        const stackElementWidth = stackElement.getBoundingClientRect().width
        const scrollAmount = stackElementWidth
        const gap = Number(window.getComputedStyle(stackContainer).gap.slice(0, -2))

        if (scrollContainer.scrollLeft + scrollAmount > scrollContainer.scrollWidth - scrollContainer.clientWidth) {
            scrollContainer.scrollTo({
                top: 0,
                left: scrollContainer.scrollWidth - scrollContainer.clientWidth,
                behavior: "smooth",
            })
        } else {
            scrollContainer.scrollBy({ left: scrollAmount + gap, behavior: "smooth" })
        }

    }
})
