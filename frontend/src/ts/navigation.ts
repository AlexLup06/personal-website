const NAV_LINKS = {
    home: 0,
    portfolio: 1,
    blog: 2
}
window.htmx.onLoad(function () {
    const navbar = document.getElementById('navbar') as HTMLDivElement;
    const footer = document.getElementById('footer-navlinks') as HTMLDivElement;
    const logo = document.getElementById("logo") as HTMLDivElement;
    if (navbar == null || footer == null || logo == null) {
        return
    }

    const navlinks = navbar.children;
    const footerLinks = footer.children

    logo.addEventListener("click", () => {
        if (window.location.pathname == "/") {
            return
        }

        const currentlyActive = document.querySelector("[data-linkactive='true']") as HTMLDivElement
        currentlyActive.setAttribute('data-linkactive', 'false')
        navlinks[NAV_LINKS.home].setAttribute('data-linkactive', 'true')
    })

    for (var i = 0; i < navlinks.length; i++) {
        const navlink = navlinks[i];
        navlink.addEventListener('click', function () {
            const pathAttr = navlink.getAttribute('data-path')
            const pathName = window.location.pathname

            if (pathAttr === pathName) {
                return
            }

            const currentlyActive = document.querySelector('[data-linkactive="true"]')!
            currentlyActive.setAttribute('data-linkactive', 'false')

            navlink.setAttribute('data-linkactive', 'true')
        });
    }

    for (var i = 0; i < footerLinks.length; i++) {
        const footerLink = footerLinks[i];
        footerLink.addEventListener('click', function () {

            const pathAttr = footerLink.getAttribute('data-footerpath')!
            const pathName = window.location.pathname

            if (pathAttr === pathName) {
                return
            }

            switch (pathAttr) {
                case "/":
                    navlinks[NAV_LINKS.home].setAttribute('data-linkactive', 'true')
                    navlinks[NAV_LINKS.portfolio].setAttribute('data-linkactive', 'false')
                    navlinks[NAV_LINKS.blog].setAttribute('data-linkactive', 'false')
                    break
                case "/portfolio":
                    navlinks[NAV_LINKS.home].setAttribute('data-linkactive', 'false')
                    navlinks[NAV_LINKS.portfolio].setAttribute('data-linkactive', 'true')
                    navlinks[NAV_LINKS.blog].setAttribute('data-linkactive', 'false')
                    break
                case "/blog":
                    navlinks[NAV_LINKS.home].setAttribute('data-linkactive', 'false')
                    navlinks[NAV_LINKS.portfolio].setAttribute('data-linkactive', 'false')
                    navlinks[NAV_LINKS.blog].setAttribute('data-linkactive', 'true')
                    break
            }
            window.scrollTo({
                top: 0,
                left: 0,
                behavior: "smooth",
            })
        });
    }
})