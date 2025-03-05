const NAV_LINKS = {
    home: 0,
    portfolio: 1,
    blog: 2
}

const navLinks = document.querySelectorAll("#navbar-link");
const footerLinks = document.querySelectorAll("#footer-links");
const logo = document.getElementById("logo") as HTMLDivElement;


logo.addEventListener("click", () => {
    if (window.location.pathname == "/") {
        return
    }

    const currentlyActive = document.querySelector("[data-linkactive='true']") as HTMLDivElement
    currentlyActive.setAttribute('data-linkactive', 'false')
    navLinks[NAV_LINKS.home].setAttribute('data-linkactive', 'true')
})

for (var i = 0; i < navLinks.length; i++) {
    const navlink = navLinks[i];
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
                navLinks[NAV_LINKS.home].setAttribute('data-linkactive', 'true')
                navLinks[NAV_LINKS.portfolio].setAttribute('data-linkactive', 'false')
                navLinks[NAV_LINKS.blog].setAttribute('data-linkactive', 'false')
                break
            case "/portfolio":
                navLinks[NAV_LINKS.home].setAttribute('data-linkactive', 'false')
                navLinks[NAV_LINKS.portfolio].setAttribute('data-linkactive', 'true')
                navLinks[NAV_LINKS.blog].setAttribute('data-linkactive', 'false')
                break
            case "/blog":
                navLinks[NAV_LINKS.home].setAttribute('data-linkactive', 'false')
                navLinks[NAV_LINKS.portfolio].setAttribute('data-linkactive', 'false')
                navLinks[NAV_LINKS.blog].setAttribute('data-linkactive', 'true')
                break
        }
    });
}