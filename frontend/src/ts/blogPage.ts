window.htmx.onLoad(function () {

    const searchAndFilter = document.getElementById("search-and-filter") as HTMLDivElement

    const observer = new IntersectionObserver(
        ([entry]) => {
            if (entry.intersectionRatio < 1) {
                // Element is stuck (not visible in normal flow)
                searchAndFilter.setAttribute("data-stickyactive", "true")
                console.log("in")
            } else {
                // Element is no longer sticky
                searchAndFilter.setAttribute("data-stickyactive", "false")
                console.log("out")
            }
        },
        { threshold: [1] } // Triggers when element is fully in or out of view
    );


    observer.observe(searchAndFilter);

})