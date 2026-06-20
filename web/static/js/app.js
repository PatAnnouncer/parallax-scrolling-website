// Базовый параллакс через CSS переменную
window.addEventListener("scroll", e => {
	document.body.style.cssText += `--scrollTOP: ${this.scrollY}px`
});
