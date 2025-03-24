const NAV_LINKS = {
    home: 0,
    portfolio: 1,
    blog: 2
}

const navLinks = document.querySelectorAll("#navbar-link");

window.htmx.onLoad(() => {
    const pathName = window.location.pathname.split("/")[1]

    switch (pathName) {
        case "":
            navLinks[NAV_LINKS.home].setAttribute('data-linkactive', 'true')
            navLinks[NAV_LINKS.portfolio].setAttribute('data-linkactive', 'false')
            navLinks[NAV_LINKS.blog].setAttribute('data-linkactive', 'false')
            break
        case "portfolio":
            navLinks[NAV_LINKS.home].setAttribute('data-linkactive', 'false')
            navLinks[NAV_LINKS.portfolio].setAttribute('data-linkactive', 'true')
            navLinks[NAV_LINKS.blog].setAttribute('data-linkactive', 'false')
            break
        case "blog":
            navLinks[NAV_LINKS.home].setAttribute('data-linkactive', 'false')
            navLinks[NAV_LINKS.portfolio].setAttribute('data-linkactive', 'false')
            navLinks[NAV_LINKS.blog].setAttribute('data-linkactive', 'true')
            break
    }
})