var parallaxElements = document.querySelectorAll('.parallax-element');
var elementsInView = new Map();
for (var _i = 0, parallaxElements_1 = parallaxElements; _i < parallaxElements_1.length; _i++) {
    var element = parallaxElements_1[_i];
    elementsInView.set(element.id, false);
}
var observer = new IntersectionObserver(function (entries) {
    entries.forEach(function (entry) {
        if (entry.isIntersecting) {
            elementsInView.set(entry.target.id, true);
            window.addEventListener('scroll', handleScroll);
        }
        else {
            elementsInView.set(entry.target.id, false);
            window.removeEventListener('scroll', handleScroll);
        }
        var atLeastOneInView = false;
        for (var _i = 0, elementsInView_1 = elementsInView; _i < elementsInView_1.length; _i++) {
            var _a = elementsInView_1[_i], _ = _a[0], inView = _a[1];
            if (inView) {
                atLeastOneInView = true;
                break;
            }
        }
        if (atLeastOneInView) {
            window.addEventListener('scroll', handleScroll);
        }
        else {
            window.removeEventListener('scroll', handleScroll);
        }
    });
}, {});
for (var _a = 0, parallaxElements_2 = parallaxElements; _a < parallaxElements_2.length; _a++) {
    var element = parallaxElements_2[_a];
    observer.observe(element);
}
function handleScroll() {
    for (var _i = 0, elementsInView_2 = elementsInView; _i < elementsInView_2.length; _i++) {
        var _a = elementsInView_2[_i], elementId = _a[0], inView = _a[1];
        if (!inView)
            continue;
        var element = document.querySelector("#".concat(elementId));
        var minTranslateY = 0;
        var maxTranslateY = 100;
        var scrollProgress = calculateScrollProgress(element);
        var currentTranslateY = minTranslateY + (maxTranslateY - minTranslateY) * scrollProgress / 100;
        element.style.transform = "translateY(-".concat(currentTranslateY, "px)");
    }
}
function calculateScrollProgress(section) {
    var rect = section.getBoundingClientRect();
    var viewportHeight = window.innerHeight;
    var totalScrollableDistance = rect.height + viewportHeight;
    var distanceScrolled = viewportHeight - rect.top;
    var progress = Math.min(100, Math.max(0, (distanceScrolled / totalScrollableDistance) * 100));
    return progress;
}
