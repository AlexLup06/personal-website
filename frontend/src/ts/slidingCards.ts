window.htmx.onLoad(function () {
    const lifeSection = document.getElementById('life-section') as HTMLElement;
    if (lifeSection == null) {
        return
    }
    const lifeSectionCards = lifeSection.children as HTMLCollectionOf<HTMLDivElement>;

    const lifeCardObserver = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                window.addEventListener('scroll', handleLifeCardScroll);
            } else {
                window.removeEventListener('scroll', handleLifeCardScroll);
            }
        });
    }, {
    });

    lifeCardObserver.observe(lifeSection);

    function handleLifeCardScroll() {
        const scrolled = window.scrollY + window.innerHeight - lifeSection.offsetTop;
        const currentIndex = Math.floor(scrolled / window.innerHeight)
        const ratioInView = scrolled / window.innerHeight - currentIndex;

        // console.log(scrolled, ratioInView, currentIndex);

        if (currentIndex < 0) {
            return;
        }

        if (currentIndex > 0) {
            const prevCard = lifeSectionCards[currentIndex - 1];
            prevCard.style.transform = `scale(${2 - ratioInView})`;
            if (ratioInView > 0.2) {
                prevCard.style.opacity = `${1.2 - ratioInView}`;
            } else {
                prevCard.style.opacity = '1';
            }
        }

        if (currentIndex < lifeSectionCards.length) {
            const currentCard = lifeSectionCards[currentIndex];
            currentCard.style.transform = `scale(${1 + ratioInView})`;
            currentCard.style.opacity = `${1.5 * ratioInView}`;
        }
    }
})