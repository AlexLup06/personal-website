/******/ (() => { // webpackBootstrap
/******/ 	"use strict";
/******/ 	var __webpack_modules__ = ({

/***/ "./src/ts/blogPage.ts":
/*!****************************!*\
  !*** ./src/ts/blogPage.ts ***!
  \****************************/
/***/ (() => {


window.htmx.onLoad(function () {
    const searchAndFilter = document.getElementById("search-and-filter");
    if (searchAndFilter == null) {
        return;
    }
    const observer = new IntersectionObserver(([entry]) => {
        if (entry.intersectionRatio < 1) {
            searchAndFilter.setAttribute("data-stickyactive", "true");
        }
        else {
            searchAndFilter.setAttribute("data-stickyactive", "false");
        }
    }, { threshold: 1 });
    observer.observe(searchAndFilter);
});


/***/ }),

/***/ "./src/ts/blogScroll.ts":
/*!******************************!*\
  !*** ./src/ts/blogScroll.ts ***!
  \******************************/
/***/ (() => {


window.htmx.onLoad(function () {
    const blogSection = document.getElementById('blog-section');
    if (blogSection == null) {
        return;
    }
    const clientHeight = blogSection.clientHeight; // Visible height
    const indicatorSection = document.getElementById("indicator-section");
    const indicators = indicatorSection.childNodes;
    addHighlight(indicators[0]);
    removeHighlight(indicators[1]);
    removeHighlight(indicators[2]);
    blogSection.scrollTo({
        top: 0,
        left: 0,
        behavior: "smooth",
    });
    indicators.forEach((indicator) => {
        indicator.addEventListener("click", (event) => {
            switch (event.target.id) {
                case "indicator-1": {
                    addHighlight(indicators[0]);
                    removeHighlight(indicators[1]);
                    removeHighlight(indicators[2]);
                    blogSection.scrollTo({
                        top: 0,
                        left: 0,
                        behavior: "smooth",
                    });
                    break;
                }
                case "indicator-2": {
                    removeHighlight(indicators[0]);
                    addHighlight(indicators[1]);
                    removeHighlight(indicators[2]);
                    blogSection.scrollTo({
                        top: clientHeight,
                        left: 0,
                        behavior: "smooth",
                    });
                    break;
                }
                case "indicator-3": {
                    removeHighlight(indicators[0]);
                    removeHighlight(indicators[1]);
                    addHighlight(indicators[2]);
                    blogSection.scrollTo({
                        top: 2 * clientHeight,
                        left: 0,
                        behavior: "smooth",
                    });
                    break;
                }
            }
        });
    });
    const blogObserver = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                blogSection.addEventListener('scroll', () => handleBlogScroll(blogSection, indicators));
            }
            else {
                blogSection.removeEventListener('scroll', () => handleBlogScroll(blogSection, indicators));
            }
        });
    }, {});
    blogObserver.observe(blogSection);
});
function handleBlogScroll(blogSection, indicators) {
    const scrollTop = blogSection.scrollTop; // How far scrolled from top
    const clientHeight = blogSection.clientHeight; // Visible height
    const ratio = scrollTop / clientHeight;
    const epsilon = 0.1;
    if (!between(ratio, 0, 0 + epsilon) && !between(ratio, 1 - epsilon, 1 + epsilon) && !between(ratio, 2 - epsilon, 2)) {
        return;
    }
    for (var i = 0; i < indicators.length; i++) {
        const indicator = indicators[i];
        if (i - epsilon < scrollTop / clientHeight && scrollTop / clientHeight < i + epsilon) {
            addHighlight(indicator);
            continue;
        }
        removeHighlight(indicator);
    }
}
function addHighlight(indicator) {
    indicator.classList.add("bg-second-500");
    indicator.classList.add("h-8");
    indicator.classList.remove("bg-second-50");
    indicator.classList.remove("h-4");
}
function removeHighlight(indicator) {
    indicator.classList.remove("bg-second-500");
    indicator.classList.remove("h-8");
    indicator.classList.add("bg-second-50");
    indicator.classList.add("h-4");
}
function between(x, min, max) {
    return x >= min && x <= max;
}


/***/ }),

/***/ "./src/ts/navigation.ts":
/*!******************************!*\
  !*** ./src/ts/navigation.ts ***!
  \******************************/
/***/ (() => {


const NAV_LINKS = {
    home: 0,
    portfolio: 1,
    blog: 2
};
const navLinks = document.querySelectorAll("#navbar-link");
window.htmx.onLoad(() => {
    const pathName = window.location.pathname.split("/")[1];
    switch (pathName) {
        case "":
            navLinks[NAV_LINKS.home].setAttribute('data-linkactive', 'true');
            navLinks[NAV_LINKS.portfolio].setAttribute('data-linkactive', 'false');
            navLinks[NAV_LINKS.blog].setAttribute('data-linkactive', 'false');
            break;
        case "portfolio":
            navLinks[NAV_LINKS.home].setAttribute('data-linkactive', 'false');
            navLinks[NAV_LINKS.portfolio].setAttribute('data-linkactive', 'true');
            navLinks[NAV_LINKS.blog].setAttribute('data-linkactive', 'false');
            break;
        case "blog":
            navLinks[NAV_LINKS.home].setAttribute('data-linkactive', 'false');
            navLinks[NAV_LINKS.portfolio].setAttribute('data-linkactive', 'false');
            navLinks[NAV_LINKS.blog].setAttribute('data-linkactive', 'true');
            break;
    }
});


/***/ }),

/***/ "./src/ts/techStack.ts":
/*!*****************************!*\
  !*** ./src/ts/techStack.ts ***!
  \*****************************/
/***/ (() => {


window.htmx.onLoad(() => {
    const leftButton = document.getElementById("left-button");
    const rightButton = document.getElementById("right-button");
    const stackContainer = document.getElementById("stack-container");
    const scrollContainer = document.getElementById("scroll-stack-container");
    if (leftButton == null || rightButton == null || stackContainer == null) {
        return;
    }
    leftButton.addEventListener("click", () => handleLeftClick());
    rightButton.addEventListener("click", () => handleRightClick());
    function handleLeftClick() {
        const stackElement = stackContainer.children[0];
        const stackElementWidth = stackElement.getBoundingClientRect().width;
        const scrollContainerWidth = scrollContainer.getBoundingClientRect().width;
        const scrollAmount = stackElementWidth;
        const gap = Number(window.getComputedStyle(stackContainer).gap.slice(0, -2));
        const isAtFarRight = scrollContainer.scrollWidth - scrollContainer.clientWidth == scrollContainer.scrollLeft;
        if (isAtFarRight) {
            const fullElements = Math.floor(scrollContainerWidth / stackElementWidth);
            const partialWidthInside = scrollContainerWidth - fullElements * stackElementWidth - fullElements * gap;
            const partialWidthOutisde = stackElementWidth - partialWidthInside;
            scrollContainer.scrollBy({ left: -partialWidthOutisde, behavior: "smooth" });
            return;
        }
        if (scrollContainer.scrollLeft - scrollAmount < 0) {
            scrollContainer.scrollTo({
                top: 0,
                left: 0,
                behavior: "smooth",
            });
        }
        else {
            scrollContainer.scrollBy({ left: -scrollAmount - gap, behavior: "smooth" });
        }
    }
    function handleRightClick() {
        const stackElement = stackContainer.children[0];
        const stackElementWidth = stackElement.getBoundingClientRect().width;
        const scrollAmount = stackElementWidth;
        const gap = Number(window.getComputedStyle(stackContainer).gap.slice(0, -2));
        if (scrollContainer.scrollLeft + scrollAmount > scrollContainer.scrollWidth - scrollContainer.clientWidth) {
            scrollContainer.scrollTo({
                top: 0,
                left: scrollContainer.scrollWidth - scrollContainer.clientWidth,
                behavior: "smooth",
            });
        }
        else {
            scrollContainer.scrollBy({ left: scrollAmount + gap, behavior: "smooth" });
        }
    }
});


/***/ })

/******/ 	});
/************************************************************************/
/******/ 	// The module cache
/******/ 	var __webpack_module_cache__ = {};
/******/ 	
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/ 		// Check if module is in cache
/******/ 		var cachedModule = __webpack_module_cache__[moduleId];
/******/ 		if (cachedModule !== undefined) {
/******/ 			return cachedModule.exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = __webpack_module_cache__[moduleId] = {
/******/ 			// no module.id needed
/******/ 			// no module.loaded needed
/******/ 			exports: {}
/******/ 		};
/******/ 	
/******/ 		// Execute the module function
/******/ 		__webpack_modules__[moduleId](module, module.exports, __webpack_require__);
/******/ 	
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/ 	
/************************************************************************/
/******/ 	/* webpack/runtime/compat get default export */
/******/ 	(() => {
/******/ 		// getDefaultExport function for compatibility with non-harmony modules
/******/ 		__webpack_require__.n = (module) => {
/******/ 			var getter = module && module.__esModule ?
/******/ 				() => (module['default']) :
/******/ 				() => (module);
/******/ 			__webpack_require__.d(getter, { a: getter });
/******/ 			return getter;
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/define property getters */
/******/ 	(() => {
/******/ 		// define getter functions for harmony exports
/******/ 		__webpack_require__.d = (exports, definition) => {
/******/ 			for(var key in definition) {
/******/ 				if(__webpack_require__.o(definition, key) && !__webpack_require__.o(exports, key)) {
/******/ 					Object.defineProperty(exports, key, { enumerable: true, get: definition[key] });
/******/ 				}
/******/ 			}
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/hasOwnProperty shorthand */
/******/ 	(() => {
/******/ 		__webpack_require__.o = (obj, prop) => (Object.prototype.hasOwnProperty.call(obj, prop))
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/make namespace object */
/******/ 	(() => {
/******/ 		// define __esModule on exports
/******/ 		__webpack_require__.r = (exports) => {
/******/ 			if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 				Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 			}
/******/ 			Object.defineProperty(exports, '__esModule', { value: true });
/******/ 		};
/******/ 	})();
/******/ 	
/************************************************************************/
var __webpack_exports__ = {};
// This entry needs to be wrapped in an IIFE because it needs to be isolated against other modules in the chunk.
(() => {
/*!*********************!*\
  !*** ./src/main.ts ***!
  \*********************/
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _ts_blogScroll__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ts/blogScroll */ "./src/ts/blogScroll.ts");
/* harmony import */ var _ts_blogScroll__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_ts_blogScroll__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var _ts_blogPage__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./ts/blogPage */ "./src/ts/blogPage.ts");
/* harmony import */ var _ts_blogPage__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(_ts_blogPage__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var _ts_navigation__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./ts/navigation */ "./src/ts/navigation.ts");
/* harmony import */ var _ts_navigation__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(_ts_navigation__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var _ts_techStack__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./ts/techStack */ "./src/ts/techStack.ts");
/* harmony import */ var _ts_techStack__WEBPACK_IMPORTED_MODULE_3___default = /*#__PURE__*/__webpack_require__.n(_ts_techStack__WEBPACK_IMPORTED_MODULE_3__);




window.htmx.config.allowNestedOobSwaps = false; // Disable nested OOB swaps
window.htmx.config.defaultSwapStyle = "outerHTML"; // Disable nested OOB swaps
const socialsDropdown = document.getElementById("socials-dd");
document.addEventListener("click", function (event) {
    // Check if the click was outside the element
    if (!socialsDropdown.contains(event.target)) {
        socialsDropdown.firstChild.removeAttribute("open");
    }
});

})();

/******/ })()
;
//# sourceMappingURL=app.js.map