// Инициализация GSAP модулей из твоего старого проекта
gsap.registerPlugin(ScrollTrigger, ScrollSmoother);

ScrollSmoother.create({
	wrapper: '#smooth-wrapper',
	content: '#smooth-content',
	smooth: 1.5,
	effects: true
});

// Базовый параллакс леса через CSS переменную
window.addEventListener("scroll", e => {
	document.body.style.cssText += `--scrollTOP: ${window.scrollY}px`;
});

// Дополнительная анимация: красивое появление карточек в подземелье при скролле
gsap.fromTo('.main-article .max-w-5xl > *',
	{ opacity: 0, y: 80 },
	{
		opacity: 1,
		y: 0,
		duration: 1.2,
		stagger: 0.3,
		ease: "power2.out",
		scrollTrigger: {
			trigger: '.main-article',
			start: 'top 85%',
			toggleActions: 'play none none reverse'
		}
	}
);