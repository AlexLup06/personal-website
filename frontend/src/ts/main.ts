import "./blogScroll"
import "./blogPage"
import "./navigation"
import "./techStack"

declare global {
    interface Window {
        htmx: any;
    }
}

const socialsDropdown = document.getElementById("socials-dd")!;
document.addEventListener("click", function (event: any) {
    // Check if the click was outside the element
    if (!socialsDropdown.contains(event.target)) {
        (socialsDropdown.firstChild as HTMLElement).removeAttribute("open");
    }
});