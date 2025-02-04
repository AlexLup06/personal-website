window.htmx.onLoad(function () {
    const parallaxElements = document.querySelectorAll('.parallax-element');
    const elementsInView = new Map<string, boolean>();

    for (const element of parallaxElements) {
        elementsInView.set(element.id, false);
    }

    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                elementsInView.set(entry.target.id, true);
                window.addEventListener('scroll', handleScroll);
            } else {
                elementsInView.set(entry.target.id, false);
                window.removeEventListener('scroll', handleScroll);
            }

            var atLeastOneInView = false
            for (const [_, inView] of elementsInView) {
                if (inView) {
                    atLeastOneInView = true
                    break
                }
            }
            if (atLeastOneInView) {
                window.addEventListener('scroll', handleScroll);
            } else {
                window.removeEventListener('scroll', handleScroll);
            }

        });
    }, {
    });

    for (const element of parallaxElements) {
        observer.observe(element);
    }

    function handleScroll() {
        for (const [elementId, inView] of elementsInView) {
            if (!inView) continue;

            const element = document.querySelector(`#${elementId}`) as HTMLDivElement;
            const minTranslateY = 0;
            const maxTranslateY = 100;
            const scrollProgress = calculateScrollProgress(element);
            const currentTranslateY = minTranslateY + (maxTranslateY - minTranslateY) * scrollProgress / 100;
            element.style.transform = `translateY(-${currentTranslateY}px)`;
        }
    }


    function calculateScrollProgress(section: HTMLDivElement) {
        const rect = section.getBoundingClientRect();
        const viewportHeight = window.innerHeight;

        const totalScrollableDistance = rect.height + viewportHeight;
        const distanceScrolled = viewportHeight - rect.top;

        const progress = Math.min(
            100,
            Math.max(0, (distanceScrolled / totalScrollableDistance) * 100)
        );

        return progress;
    }
})