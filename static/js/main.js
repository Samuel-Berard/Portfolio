// ===== Initialisation =====
document.addEventListener('DOMContentLoaded', () => {
    initThemeToggle();
    initVisitCounter();
    initCarousel();
    initProjectFilters();
    initFormValidation();
    initScrollAnimations();
    initSmoothScroll();
});

// ===== Mode Sombre / Clair =====
function initThemeToggle() {
    const themeToggle = document.getElementById('theme-toggle');
    const body = document.body;
    
    // Charger le thème depuis localStorage
    const savedTheme = localStorage.getItem('theme');
    if (savedTheme === 'dark') {
        body.classList.add('dark-mode');
    }
    
    themeToggle.addEventListener('click', () => {
        body.classList.toggle('dark-mode');
        
        // Sauvegarder la préférence
        const isDark = body.classList.contains('dark-mode');
        localStorage.setItem('theme', isDark ? 'dark' : 'light');
        
        // Animation du bouton
        themeToggle.style.transform = 'rotate(360deg)';
        setTimeout(() => {
            themeToggle.style.transform = 'rotate(0deg)';
        }, 300);
    });
}

// ===== Compteur de Visites =====
function initVisitCounter() {
    const countElement = document.getElementById('visit-count');
    
    // Récupérer le nombre actuel de visites
    let visitCount = localStorage.getItem('visitCount');
    
    if (!visitCount) {
        visitCount = 0;
    }
    
    // Incrémenter le compteur
    visitCount = parseInt(visitCount) + 1;
    
    // Sauvegarder le nouveau compte
    localStorage.setItem('visitCount', visitCount);
    
    // Afficher avec animation
    animateCounter(countElement, 0, visitCount, 1000);
}

function animateCounter(element, start, end, duration) {
    const range = end - start;
    const increment = range / (duration / 16);
    let current = start;
    
    const timer = setInterval(() => {
        current += increment;
        if (current >= end) {
            element.textContent = end;
            clearInterval(timer);
        } else {
            element.textContent = Math.floor(current);
        }
    }, 16);
}

// ===== Carrousel de Projets =====
function initCarousel() {
    const track = document.querySelector('.carousel-track');
    const slides = Array.from(track.children);
    const nextButton = document.querySelector('.carousel-btn.next');
    const prevButton = document.querySelector('.carousel-btn.prev');
    const indicatorsContainer = document.querySelector('.carousel-indicators');
    
    let currentIndex = 0;
    let slidesPerView = getSlidesPerView();
    const totalSlides = slides.length;
    const maxIndex = Math.max(0, totalSlides - slidesPerView);
    
    // Créer les indicateurs
    createIndicators();
    
    // Event listeners pour les boutons
    nextButton.addEventListener('click', () => {
        if (currentIndex < maxIndex) {
            currentIndex++;
            updateCarousel();
        }
    });
    
    prevButton.addEventListener('click', () => {
        if (currentIndex > 0) {
            currentIndex--;
            updateCarousel();
        }
    });
    
    // Support tactile pour mobile
    let touchStartX = 0;
    let touchEndX = 0;
    
    track.addEventListener('touchstart', (e) => {
        touchStartX = e.changedTouches[0].screenX;
    });
    
    track.addEventListener('touchend', (e) => {
        touchEndX = e.changedTouches[0].screenX;
        handleSwipe();
    });
    
    function handleSwipe() {
        if (touchEndX < touchStartX - 50 && currentIndex < maxIndex) {
            currentIndex++;
            updateCarousel();
        }
        if (touchEndX > touchStartX + 50 && currentIndex > 0) {
            currentIndex--;
            updateCarousel();
        }
    }
    
    // Mettre à jour le carrousel
    function updateCarousel() {
        const slideWidth = slides[0].getBoundingClientRect().width;
        const gap = 32; // 2rem en pixels
        const moveAmount = -(slideWidth + gap) * currentIndex;
        
        track.style.transform = `translateX(${moveAmount}px)`;
        
        // Mettre à jour les indicateurs
        updateIndicators();
        
        // Mettre à jour l'état des boutons
        prevButton.style.opacity = currentIndex === 0 ? '0.5' : '1';
        prevButton.style.cursor = currentIndex === 0 ? 'not-allowed' : 'pointer';
        nextButton.style.opacity = currentIndex === maxIndex ? '0.5' : '1';
        nextButton.style.cursor = currentIndex === maxIndex ? 'not-allowed' : 'pointer';
    }
    
    // Créer les indicateurs
    function createIndicators() {
        const indicatorCount = maxIndex + 1;
        for (let i = 0; i < indicatorCount; i++) {
            const indicator = document.createElement('div');
            indicator.classList.add('indicator');
            if (i === 0) indicator.classList.add('active');
            
            indicator.addEventListener('click', () => {
                currentIndex = i;
                updateCarousel();
            });
            
            indicatorsContainer.appendChild(indicator);
        }
    }
    
    // Mettre à jour les indicateurs
    function updateIndicators() {
        const indicators = document.querySelectorAll('.indicator');
        indicators.forEach((indicator, index) => {
            indicator.classList.toggle('active', index === currentIndex);
        });
    }
    
    // Calculer le nombre de slides visibles
    function getSlidesPerView() {
        const width = window.innerWidth;
        if (width <= 768) return 1;
        if (width <= 1024) return 2;
        return 3;
    }
    
    // Recalculer lors du redimensionnement
    window.addEventListener('resize', () => {
        const newSlidesPerView = getSlidesPerView();
        if (newSlidesPerView !== slidesPerView) {
            slidesPerView = newSlidesPerView;
            currentIndex = Math.min(currentIndex, Math.max(0, totalSlides - slidesPerView));
            
            // Recréer les indicateurs
            indicatorsContainer.innerHTML = '';
            createIndicators();
            
            updateCarousel();
        }
    });
    
    // Initialiser
    updateCarousel();
}

// ===== Filtrage des Projets =====
function initProjectFilters() {
    const filterButtons = document.querySelectorAll('.filter-btn');
    const projectCards = document.querySelectorAll('.project-card');
    
    filterButtons.forEach(button => {
        button.addEventListener('click', () => {
            const filter = button.dataset.filter;
            
            // Mettre à jour les boutons actifs
            filterButtons.forEach(btn => btn.classList.remove('active'));
            button.classList.add('active');
            
            // Filtrer les projets
            projectCards.forEach(card => {
                const category = card.dataset.category;
                
                if (filter === 'all' || category === filter) {
                    card.style.display = 'flex';
                    card.style.animation = 'fadeIn 0.5s ease';
                } else {
                    card.style.display = 'none';
                }
            });
            
            // Réinitialiser le carrousel après le filtrage
            setTimeout(() => {
                const track = document.querySelector('.carousel-track');
                track.style.transform = 'translateX(0)';
            }, 100);
        });
    });
}

// ===== Validation du Formulaire en Temps Réel =====
function initFormValidation() {
    const form = document.getElementById('contact-form');
    const nameInput = document.getElementById('name');
    const emailInput = document.getElementById('email');
    const messageInput = document.getElementById('message');
    
    // Validation en temps réel
    nameInput.addEventListener('input', () => validateName());
    emailInput.addEventListener('input', () => validateEmail());
    messageInput.addEventListener('input', () => validateMessage());
    
    // Validation à la perte de focus
    nameInput.addEventListener('blur', () => validateName());
    emailInput.addEventListener('blur', () => validateEmail());
    messageInput.addEventListener('blur', () => validateMessage());
    
    // Soumission du formulaire
    form.addEventListener('submit', (e) => {
        e.preventDefault();
        
        const isNameValid = validateName();
        const isEmailValid = validateEmail();
        const isMessageValid = validateMessage();
        
        if (isNameValid && isEmailValid && isMessageValid) {
            // Simuler l'envoi (en production, le formulaire serait soumis au serveur)
            submitForm();
        }
    });
    
    function validateName() {
        const value = nameInput.value.trim();
        const errorElement = document.getElementById('name-error');
        
        if (value === '') {
            showError(nameInput, errorElement, 'Le nom est requis');
            return false;
        } else if (value.length < 2) {
            showError(nameInput, errorElement, 'Le nom doit contenir au moins 2 caractères');
            return false;
        } else {
            clearError(nameInput, errorElement);
            return true;
        }
    }
    
    function validateEmail() {
        const value = emailInput.value.trim();
        const errorElement = document.getElementById('email-error');
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        
        if (value === '') {
            showError(emailInput, errorElement, 'L\'email est requis');
            return false;
        } else if (!emailRegex.test(value)) {
            showError(emailInput, errorElement, 'Veuillez entrer un email valide');
            return false;
        } else {
            clearError(emailInput, errorElement);
            return true;
        }
    }
    
    function validateMessage() {
        const value = messageInput.value.trim();
        const errorElement = document.getElementById('message-error');
        
        if (value === '') {
            showError(messageInput, errorElement, 'Le message est requis');
            return false;
        } else if (value.length < 10) {
            showError(messageInput, errorElement, 'Le message doit contenir au moins 10 caractères');
            return false;
        } else {
            clearError(messageInput, errorElement);
            return true;
        }
    }
    
    function showError(input, errorElement, message) {
        input.classList.add('error');
        errorElement.textContent = message;
    }
    
    function clearError(input, errorElement) {
        input.classList.remove('error');
        errorElement.textContent = '';
    }
    
    function submitForm() {
        const successMessage = document.getElementById('form-success');
        
        // Afficher le message de succès
        successMessage.classList.add('show');
        
        // Réinitialiser le formulaire
        form.reset();
        
        // Cacher le message après 5 secondes
        setTimeout(() => {
            successMessage.classList.remove('show');
        }, 5000);
        
        // En production, vous soumettriez réellement le formulaire au serveur
        // form.submit();
    }
}

// ===== Animations au Scroll =====
function initScrollAnimations() {
    const observerOptions = {
        threshold: 0.1,
        rootMargin: '0px 0px -50px 0px'
    };
    
    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.classList.add('visible');
                
                // Animer les barres de compétences
                if (entry.target.classList.contains('skill-category')) {
                    animateSkillBars(entry.target);
                }
            }
        });
    }, observerOptions);
    
    // Observer tous les éléments avec animation
    const animatedElements = document.querySelectorAll('.fade-in-up, .skill-category');
    animatedElements.forEach(el => {
        el.classList.add('scroll-animate');
        observer.observe(el);
    });
}

function animateSkillBars(skillCategory) {
    const skillBars = skillCategory.querySelectorAll('.skill-progress');
    
    skillBars.forEach((bar, index) => {
        setTimeout(() => {
            const width = bar.style.width;
            bar.style.width = '0%';
            setTimeout(() => {
                bar.style.width = width;
            }, 50);
        }, index * 100);
    });
}

// ===== Smooth Scroll =====
function initSmoothScroll() {
    const navLinks = document.querySelectorAll('.nav-links a');
    
    navLinks.forEach(link => {
        link.addEventListener('click', (e) => {
            e.preventDefault();
            
            const targetId = link.getAttribute('href');
            const targetSection = document.querySelector(targetId);
            
            if (targetSection) {
                const navbarHeight = document.querySelector('.navbar').offsetHeight;
                const targetPosition = targetSection.offsetTop - navbarHeight;
                
                window.scrollTo({
                    top: targetPosition,
                    behavior: 'smooth'
                });
            }
        });
    });
    
    // Boutons du hero
    const heroButtons = document.querySelectorAll('.hero-buttons a');
    heroButtons.forEach(button => {
        if (button.getAttribute('href').startsWith('#')) {
            button.addEventListener('click', (e) => {
                e.preventDefault();
                
                const targetId = button.getAttribute('href');
                const targetSection = document.querySelector(targetId);
                
                if (targetSection) {
                    const navbarHeight = document.querySelector('.navbar').offsetHeight;
                    const targetPosition = targetSection.offsetTop - navbarHeight;
                    
                    window.scrollTo({
                        top: targetPosition,
                        behavior: 'smooth'
                    });
                }
            });
        }
    });
}

// ===== Highlight Active Nav Link =====
window.addEventListener('scroll', () => {
    const sections = document.querySelectorAll('.section');
    const navLinks = document.querySelectorAll('.nav-links a');
    
    let current = '';
    
    sections.forEach(section => {
        const sectionTop = section.offsetTop;
        const sectionHeight = section.clientHeight;
        const navbarHeight = document.querySelector('.navbar').offsetHeight;
        
        if (window.pageYOffset >= sectionTop - navbarHeight - 100) {
            current = section.getAttribute('id');
        }
    });
    
    navLinks.forEach(link => {
        link.classList.remove('active');
        if (link.getAttribute('href') === `#${current}`) {
            link.classList.add('active');
        }
    });
});

// ===== Utilitaires =====

// Fonction pour débouncer (optimiser les events de resize/scroll)
function debounce(func, wait) {
    let timeout;
    return function executedFunction(...args) {
        const later = () => {
            clearTimeout(timeout);
            func(...args);
        };
        clearTimeout(timeout);
        timeout = setTimeout(later, wait);
    };
}

// Détecter la fin du chargement des images
function onImagesLoaded(callback) {
    const images = document.querySelectorAll('img');
    let loadedCount = 0;
    const totalImages = images.length;
    
    if (totalImages === 0) {
        callback();
        return;
    }
    
    images.forEach(img => {
        if (img.complete) {
            loadedCount++;
            if (loadedCount === totalImages) {
                callback();
            }
        } else {
            img.addEventListener('load', () => {
                loadedCount++;
                if (loadedCount === totalImages) {
                    callback();
                }
            });
        }
    });
}

// Message de succès pour le paramètre URL
window.addEventListener('load', () => {
    const urlParams = new URLSearchParams(window.location.search);
    if (urlParams.get('success') === 'true') {
        const successMessage = document.getElementById('form-success');
        if (successMessage) {
            successMessage.classList.add('show');
            setTimeout(() => {
                successMessage.classList.remove('show');
            }, 5000);
        }
    }
});

// Log de démarrage
console.log('Portfolio initialized successfully! 🚀');
console.log('Features:');
console.log('✓ Dark/Light mode toggle');
console.log('✓ Visit counter');
console.log('✓ Project carousel');
console.log('✓ Project filtering');
console.log('✓ Real-time form validation');
console.log('✓ Scroll animations');
console.log('✓ Smooth scrolling');
