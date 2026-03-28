/******/ (() => { // webpackBootstrap
/******/ 	"use strict";
/******/ 	var __webpack_modules__ = ({

/***/ "./src/navigation.ts"
/*!***************************!*\
  !*** ./src/navigation.ts ***!
  \***************************/
() {


const NAV_LINKS = {
    home: 0,
    portfolio: 1,
    blog: 2
};
window.htmx.onLoad(() => {
    const navLinks = document.querySelectorAll("#navbar-link");
    const pathName = window.location.pathname.split("/")[1];
    switch (pathName) {
        case "":
            navLinks[NAV_LINKS.home].setAttribute('data-linkactive', 'true');
            navLinks[NAV_LINKS.portfolio].setAttribute('data-linkactive', 'false');
            // navLinks[NAV_LINKS.blog].setAttribute('data-linkactive', 'false')
            break;
        case "portfolio":
            navLinks[NAV_LINKS.home].setAttribute('data-linkactive', 'false');
            navLinks[NAV_LINKS.portfolio].setAttribute('data-linkactive', 'true');
            // navLinks[NAV_LINKS.blog].setAttribute('data-linkactive', 'false')
            break;
        case "blog":
            navLinks[NAV_LINKS.home].setAttribute('data-linkactive', 'false');
            navLinks[NAV_LINKS.portfolio].setAttribute('data-linkactive', 'false');
            // navLinks[NAV_LINKS.blog].setAttribute('data-linkactive', 'true')
            break;
    }
});


/***/ },

/***/ "./src/techStack.ts"
/*!**************************!*\
  !*** ./src/techStack.ts ***!
  \**************************/
() {


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


/***/ }

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
/******/ 		if (!(moduleId in __webpack_modules__)) {
/******/ 			delete __webpack_module_cache__[moduleId];
/******/ 			var e = new Error("Cannot find module '" + moduleId + "'");
/******/ 			e.code = 'MODULE_NOT_FOUND';
/******/ 			throw e;
/******/ 		}
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
/* harmony import */ var _navigation__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./navigation */ "./src/navigation.ts");
/* harmony import */ var _navigation__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_navigation__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var _techStack__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./techStack */ "./src/techStack.ts");
/* harmony import */ var _techStack__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(_techStack__WEBPACK_IMPORTED_MODULE_1__);


window.htmx.config.allowNestedOobSwaps = false; // Disable nested OOB swaps
window.htmx.config.defaultSwapStyle = "outerHTML"; // Disable nested OOB swaps
document.body.addEventListener("htmx:beforeSwap", function (evt) {
    // Allow 422 and 400 responses to swap
    // We treat these as form validation errors
    if (evt.detail.xhr.status === 422 ||
        evt.detail.xhr.status === 400 ||
        evt.detail.xhr.status === 429) {
        evt.detail.shouldSwap = true;
        evt.detail.isError = false;
    }
});
document.body.addEventListener("htmx:afterSettle", function (event) {
    if (event.detail.target && event.detail.target.id === "body-section") {
        window.scrollTo({ top: 0, behavior: "smooth" });
    }
});

})();

/******/ })()
;
//# sourceMappingURL=app.js.map