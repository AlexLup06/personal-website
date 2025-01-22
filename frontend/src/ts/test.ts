const navbar: HTMLDivElement = document.querySelector('#navbar')!;
const navlinks = navbar.children;

for (var i = 0; i < navlinks.length; i++) {
    const navlink = navlinks[i];
    navlink.addEventListener('click', function () {
        const pathAttr = navlink.getAttribute('data-path')
        const pathName = window.location.pathname

        if (pathAttr === pathName) {
            return
        }

        const currentlyActive = document.querySelector('[data-active="true"]')!
        currentlyActive.setAttribute('data-active', 'false')

        navlink.setAttribute('data-active', 'true')
    });
}