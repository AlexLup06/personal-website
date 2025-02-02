var lifeSection = document.getElementById('life-section');
var lifeSectionCards = lifeSection.children;
var lifeCardObserver = new IntersectionObserver(function (entries) {
    entries.forEach(function (entry) {
        if (entry.isIntersecting) {
            window.addEventListener('scroll', handleLifeCardScroll);
        }
        else {
            window.removeEventListener('scroll', handleLifeCardScroll);
        }
    });
}, {});
lifeCardObserver.observe(lifeSection);
function handleLifeCardScroll() {
    var scrolled = window.scrollY + window.innerHeight - lifeSection.offsetTop;
    var currentIndex = Math.floor(scrolled / window.innerHeight);
    var ratioInView = scrolled / window.innerHeight - currentIndex;
    // console.log(scrolled, ratioInView, currentIndex);
    if (currentIndex < 0) {
        return;
    }
    if (currentIndex > 0) {
        var prevCard = lifeSectionCards[currentIndex - 1];
        prevCard.style.transform = "scale(".concat(2 - ratioInView, ")");
        if (ratioInView > 0.2) {
            prevCard.style.opacity = "".concat(1.2 - ratioInView);
        }
        else {
            prevCard.style.opacity = '1';
        }
    }
    if (currentIndex < lifeSectionCards.length) {
        var currentCard = lifeSectionCards[currentIndex];
        currentCard.style.transform = "scale(".concat(1 + ratioInView, ")");
        currentCard.style.opacity = "".concat(1.5 * ratioInView);
    }
}
