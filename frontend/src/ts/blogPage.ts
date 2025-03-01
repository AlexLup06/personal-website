window.htmx.onLoad(function () {
    const searchAndFilter = document.getElementById("search-and-filter") as HTMLDivElement
    if (searchAndFilter == null) {
        return
    }

    const observer = new IntersectionObserver(
        ([entry]) => {
            if (entry.intersectionRatio < 1) {
                searchAndFilter.setAttribute("data-stickyactive", "true")
            } else {
                searchAndFilter.setAttribute("data-stickyactive", "false")
            }
        },
        { threshold: 1 }
    );


    observer.observe(searchAndFilter);

})