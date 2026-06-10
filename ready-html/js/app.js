window.addEventListener("scroll", e => {
	document.body.style.cssText += `--scrollTOP: ${this.scrollY}px`
});

// Инициализация параллакса
gsap.registerPlugin(ScrollTrigger, ScrollSmoother);
ScrollSmoother.create({
	wrapper: '#smooth-wrapper',
	content: '#smooth-content',
	smooth: 1.5,
	effects: true
});

// Получение статуса серверов от Go-бэкенда
async function fetchServerStatus() {
	try {
		const response = await fetch('/api/status');
		const data = await response.json();

		const mcElement = document.getElementById('mc-status');
		const dndElement = document.getElementById('dnd-status');

		// Обновляем статус Minecraft
		if (data.minecraft) {
			mcElement.textContent = "Онлайн";
			mcElement.style.color = "#55ff55"; // Зеленый
		} else {
			mcElement.textContent = "Оффлайн";
			mcElement.style.color = "#ff5555"; // Красный
		}

		// Обновляем статус D&D
		if (data.dnd) {
			dndElement.textContent = "Онлайн";
			dndElement.style.color = "#55ff55";
		} else {
			dndElement.textContent = "Оффлайн";
			dndElement.style.color = "#ff5555";
		}
	} catch (error) {
		console.error("Ошибка при получении статуса серверов:", error);
		document.getElementById('mc-status').textContent = "Ошибка API";
		document.getElementById('dnd-status').textContent = "Ошибка API";
	}
}

// Запускаем проверку при загрузке страницы
document.addEventListener("DOMContentLoaded", fetchServerStatus);