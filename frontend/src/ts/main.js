"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
require("./parallax");
require("./slidingCards");
var navbar = document.querySelector('#navbar');
var navlinks = navbar.children;
var _loop_1 = function () {
    var navlink = navlinks[i];
    navlink.addEventListener('click', function () {
        var pathAttr = navlink.getAttribute('data-path');
        var pathName = window.location.pathname;
        if (pathAttr === pathName) {
            return;
        }
        var currentlyActive = document.querySelector('[data-active="true"]');
        currentlyActive.setAttribute('data-active', 'false');
        navlink.setAttribute('data-active', 'true');
    });
};
for (var i = 0; i < navlinks.length; i++) {
    _loop_1();
}
var socialsDropdown = document.getElementById("socials-dd");
document.addEventListener("click", function (event) {
    // Check if the click was outside the element
    if (!socialsDropdown.contains(event.target)) {
        socialsDropdown.firstChild.removeAttribute("open");
    }
});
