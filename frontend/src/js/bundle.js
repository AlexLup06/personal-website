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

/***/ "./src/ts/test.ts":
/*!************************!*\
  !*** ./src/ts/test.ts ***!
  \************************/
/***/ (() => {

eval("\nconst navbar = document.querySelector('#navbar');\nconst navlinks = navbar.children;\nfor (var i = 0; i < navlinks.length; i++) {\n    const navlink = navlinks[i];\n    navlink.addEventListener('click', function () {\n        const pathAttr = navlink.getAttribute('data-path');\n        const pathName = window.location.pathname;\n        if (pathAttr === pathName) {\n            return;\n        }\n        const currentlyActive = document.querySelector('[data-active=\"true\"]');\n        currentlyActive.setAttribute('data-active', 'false');\n        navlink.setAttribute('data-active', 'true');\n    });\n}\n\n\n//# sourceURL=webpack://presonal-website/./src/ts/test.ts?");

/***/ })

/******/ 	});
/************************************************************************/
/******/ 	
/******/ 	// startup
/******/ 	// Load entry module and return exports
/******/ 	// This entry module can't be inlined because the eval devtool is used.
/******/ 	var __webpack_exports__ = {};
/******/ 	__webpack_modules__["./src/ts/test.ts"]();
/******/ 	
/******/ })()
;