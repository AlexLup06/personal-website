window.htmx.onLoad(function () {
    const blogSection = document.getElementById('blog-section') as HTMLDivElement;
    if (blogSection == null) {
        return
    }

    const clientHeight = blogSection.clientHeight; // Visible height

    const indicatorSection = document.getElementById("indicator-section") as HTMLDivElement
    const indicators = indicatorSection.childNodes as NodeListOf<HTMLDivElement>

    indicators.forEach((indicator) => {
        indicator.addEventListener("click", (event: any) => {
            console.log(event.target.id)
            switch (event.target.id) {
                case "indicator-1": {
                    addHighlight(indicators[0])
                    removeHighlight(indicators[1])
                    removeHighlight(indicators[2])
                    blogSection.scrollTo({
                        top: 0,
                        left: 0,
                        behavior: "smooth",
                    })
                    break
                }
                case "indicator-2": {
                    removeHighlight(indicators[0])
                    addHighlight(indicators[1])
                    removeHighlight(indicators[2])
                    blogSection.scrollTo({
                        top: clientHeight,
                        left: 0,
                        behavior: "smooth",
                    })
                    break
                }
                case "indicator-3": {
                    removeHighlight(indicators[0])
                    removeHighlight(indicators[1])
                    addHighlight(indicators[2])
                    blogSection.scrollTo({
                        top: 2 * clientHeight,
                        left: 0,
                        behavior: "smooth",
                    })
                    break
                }

            }
        })
    })

    const blogObserver = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                console.log("in")
                blogSection.addEventListener('scroll', () => handleBlogScroll(blogSection, indicators));
            } else {
                console.log("out")
                blogSection.removeEventListener('scroll', () => handleBlogScroll(blogSection, indicators));
            }
        });
    }, {
    });

    blogObserver.observe(blogSection)
})


function handleBlogScroll(blogSection: HTMLDivElement, indicators: NodeListOf<HTMLDivElement>) {
    const scrollTop = blogSection.scrollTop; // How far scrolled from top
    const clientHeight = blogSection.clientHeight; // Visible height

    const ratio = scrollTop / clientHeight
    const epsilon = 0.1
    if (!between(ratio, 0, 0 + epsilon) && !between(ratio, 1 - epsilon, 1 + epsilon) && !between(ratio, 2 - epsilon, 2)) {
        return
    }

    for (var i = 0; i < indicators.length; i++) {
        const indicator = indicators[i]
        if (i - epsilon < scrollTop / clientHeight && scrollTop / clientHeight < i + epsilon) {
            addHighlight(indicator)
            continue;
        }
        removeHighlight(indicator)
    }
}

function addHighlight(indicator: HTMLDivElement) {
    indicator.classList.add("bg-second-500")
    indicator.classList.add("h-8")
    indicator.classList.remove("bg-second-25")
    indicator.classList.remove("h-4")
}

function removeHighlight(indicator: HTMLDivElement) {
    indicator.classList.remove("bg-second-500")
    indicator.classList.remove("h-8")
    indicator.classList.add("bg-second-25")
    indicator.classList.add("h-4")
}

function between(x: number, min: number, max: number) {
    return x >= min && x <= max;
}