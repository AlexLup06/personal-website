import "./blogScroll"
import "./blogPage"
import "./navigation"
import "./techStack"

declare global {
    interface Window {
        htmx: any;
    }
}

window.htmx.config.allowNestedOobSwaps = false; // Disable nested OOB swaps
window.htmx.config.defaultSwapStyle = "outerHTML" // Disable nested OOB swaps
// window.htmx.config.historyCacheSize = 0

const socialsDropdown = document.getElementById("socials-dd")!;
document.addEventListener("click", function (event: any) {
    // Check if the click was outside the element
    if (!socialsDropdown.contains(event.target)) {
        (socialsDropdown.firstChild as HTMLElement).removeAttribute("open");
    }
});