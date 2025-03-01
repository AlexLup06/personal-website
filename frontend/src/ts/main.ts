import "./parallax"
import "./slidingCards"
import "./blogScroll"
import "./blogPage"

declare global {
    interface Window {
        htmx: any;
    }
}

const navbar: HTMLDivElement = document.querySelector('#navbar')!;
const navlinks = navbar.children;
const logo = document.getElementById("logo") as HTMLDivElement;

logo.addEventListener("click", () => {
    if (window.location.pathname == "/") {
        return
    }

    const currentlyActive = document.querySelector("[data-linkactive='true']") as HTMLDivElement
    currentlyActive.setAttribute('data-linkactive', 'false')
    navlinks[0].setAttribute('data-linkactive', 'true')
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

const socialsDropdown = document.getElementById("socials-dd")!;
document.addEventListener("click", function (event: any) {
    // Check if the click was outside the element
    if (!socialsDropdown.contains(event.target)) {
        (socialsDropdown.firstChild as HTMLElement).removeAttribute("open");
    }
});