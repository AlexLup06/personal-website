import "./ts/blogScroll"
import "./ts/blogPage"
import "./ts/navigation"
import "./ts/techStack"

declare global {
    interface Window {
        htmx: any;
    }
}

window.htmx.config.allowNestedOobSwaps = false; // Disable nested OOB swaps
window.htmx.config.defaultSwapStyle = "outerHTML" // Disable nested OOB swaps

const socialsDropdown = document.getElementById("socials-dd")!;
document.addEventListener("click", function (event: any) {
    // Check if the click was outside the element
    if (!socialsDropdown.contains(event.target)) {
        (socialsDropdown.firstChild as HTMLElement).removeAttribute("open");
    }
});