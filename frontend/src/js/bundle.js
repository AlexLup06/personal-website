/*
 * ATTENTION: The "eval" devtool has been used (maybe by default in mode: "development").
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
/******/ (() => { // webpackBootstrap
/******/ 	"use strict";
/******/ 	var __webpack_modules__ = ({

/***/ "./src/ts/parallax.ts":
/*!****************************!*\
  !*** ./src/ts/parallax.ts ***!
  \****************************/
/***/ (() => {

eval("\nconst parallaxElements = document.querySelectorAll('.parallax-element');\nconst elementsInView = new Map();\nfor (const element of parallaxElements) {\n    elementsInView.set(element.id, false);\n}\nconst observer = new IntersectionObserver((entries) => {\n    entries.forEach(entry => {\n        if (entry.isIntersecting) {\n            elementsInView.set(entry.target.id, true);\n            window.addEventListener('scroll', handleScroll);\n        }\n        else {\n            elementsInView.set(entry.target.id, false);\n            window.removeEventListener('scroll', handleScroll);\n        }\n        var atLeastOneInView = false;\n        for (const [_, inView] of elementsInView) {\n            if (inView) {\n                atLeastOneInView = true;\n                break;\n            }\n        }\n        if (atLeastOneInView) {\n            window.addEventListener('scroll', handleScroll);\n        }\n        else {\n            window.removeEventListener('scroll', handleScroll);\n        }\n    });\n}, {});\nfor (const element of parallaxElements) {\n    observer.observe(element);\n}\nfunction handleScroll() {\n    for (const [elementId, inView] of elementsInView) {\n        if (!inView)\n            continue;\n        const element = document.querySelector(`#${elementId}`);\n        const minTranslateY = 0;\n        const maxTranslateY = 100;\n        const scrollProgress = calculateScrollProgress(element);\n        const currentTranslateY = minTranslateY + (maxTranslateY - minTranslateY) * scrollProgress / 100;\n        element.style.transform = `translateY(-${currentTranslateY}px)`;\n    }\n}\nfunction calculateScrollProgress(section) {\n    const rect = section.getBoundingClientRect();\n    const viewportHeight = window.innerHeight;\n    const totalScrollableDistance = rect.height + viewportHeight;\n    const distanceScrolled = viewportHeight - rect.top;\n    const progress = Math.min(100, Math.max(0, (distanceScrolled / totalScrollableDistance) * 100));\n    return progress;\n}\n\n\n//# sourceURL=webpack://presonal-website/./src/ts/parallax.ts?");

/***/ }),

/***/ "./src/ts/test.ts":
/*!************************!*\
  !*** ./src/ts/test.ts ***!
  \************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var _parallax_ts__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./parallax.ts */ \"./src/ts/parallax.ts\");\n/* harmony import */ var _parallax_ts__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_parallax_ts__WEBPACK_IMPORTED_MODULE_0__);\n\nconst navbar = document.querySelector('#navbar');\nconst navlinks = navbar.children;\nfor (var i = 0; i < navlinks.length; i++) {\n    const navlink = navlinks[i];\n    navlink.addEventListener('click', function () {\n        const pathAttr = navlink.getAttribute('data-path');\n        const pathName = window.location.pathname;\n        if (pathAttr === pathName) {\n            return;\n        }\n        const currentlyActive = document.querySelector('[data-active=\"true\"]');\n        currentlyActive.setAttribute('data-active', 'false');\n        navlink.setAttribute('data-active', 'true');\n    });\n}\nconst socialsDropdown = document.getElementById(\"socials-dd\");\ndocument.addEventListener(\"click\", function (event) {\n    // Check if the click was outside the element\n    if (!socialsDropdown.contains(event.target)) {\n        socialsDropdown.firstChild.removeAttribute(\"open\");\n    }\n});\n\n\n//# sourceURL=webpack://presonal-website/./src/ts/test.ts?");

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
/******/ 	
/******/ 	// startup
/******/ 	// Load entry module and return exports
/******/ 	// This entry module can't be inlined because the eval devtool is used.
/******/ 	var __webpack_exports__ = __webpack_require__("./src/ts/test.ts");
/******/ 	
/******/ })()
;