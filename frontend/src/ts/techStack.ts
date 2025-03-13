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
        const isSmallContainer = stackElementWidth * 2 + gap > scrollContainerWidth

        if (isSmallContainer && isAtFarRight) {
            console.log("left click")
            const partialWidth = scrollContainerWidth - stackElementWidth - gap
            const scrollAmount = stackElementWidth - partialWidth / 2
            scrollContainer.scrollBy({ left: -scrollAmount - gap / 2, behavior: "smooth" })
            return
        }

        if (isAtFarRight) {
            scrollContainer.scrollBy({ left: -scrollAmount / 2 - gap, behavior: "smooth" })
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
        const scrollContainerWidth = scrollContainer.getBoundingClientRect().width
        const gap = Number(window.getComputedStyle(stackContainer).gap.slice(0, -2))

        const isAtFarLeft = scrollContainer.scrollLeft == 0
        const isSmallContainer = stackElementWidth * 2 + gap > scrollContainerWidth


        if (isSmallContainer && isAtFarLeft) {
            /*
            358
            200 8 150 -> scroll 125
            75 8 200 8 75
            */
            const partialWidth = scrollContainerWidth - stackElementWidth - gap
            const scrollAmountSmall = stackElementWidth - partialWidth / 2
            scrollContainer.scrollBy({ left: scrollAmountSmall + gap / 2, behavior: "smooth" })
            return
        }

        if (scrollContainer.scrollLeft == 0) {
            console.log("scroll amount: ", scrollAmount / 2 + gap)
            scrollContainer.scrollBy({ left: scrollAmount / 2 + gap, behavior: "smooth" })
            return
        }


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
