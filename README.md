# 🎨 Portfolio Personnel - Go + HTML/CSS/JavaScript

Site web portfolio/CV moderne et interactif créé avec Go, HTML5, CSS3 et JavaScript.

## 🔗 Démo en ligne

**URL du projet :** [https://portfolio-berard-samuel.osc-fr1.scalingo.io](https://portfolio-berard-samuel.osc-fr1.scalingo.io)

## 🌟 Fonctionnalités

### ✨ Interface Utilisateur
- **Design moderne** avec gradients et animations fluides
- **Mode sombre/clair** avec persistance localStorage
- **Navigation sticky** avec liens smooth scroll
- **Entièrement responsive** (Desktop, Tablette, Mobile)

### 🎯 Sections
1. **Hero** - Présentation avec CTA
2. **À propos** - Biographie et statistiques
3. **Compétences** - Badges de compétences organisés par catégories
4. **Projets** - Carrousel interactif avec filtrage par catégorie
5. **Contact** - Formulaire avec validation en temps réel

### 🚀 Fonctionnalités Interactives
- ✅ **Carrousel de projets** avec navigation (boutons + indicateurs + swipe mobile)
- ✅ **Filtrage des projets** par catégorie avec animations
- ✅ **Formulaire de contact** avec validation temps réel et côté serveur
- ✅ **Compteur de visites** avec localStorage
- ✅ **Animations au scroll** avec IntersectionObserver
- ✅ **Mode sombre/clair** avec toggle et persistance

## 🛠️ Technologies

- **Backend** : Go (Golang)
- **Frontend** : HTML5, CSS3, JavaScript (ES6+)
- **APIs** : IntersectionObserver, LocalStorage, Touch Events

## 📦 Installation

### Prérequis
- Go 1.16 ou supérieur

### Lancement

1. Cloner ou télécharger le projet

2. Naviguer vers le dossier :
```bash
cd Portfolio
```

3. Lancer le serveur :
```bash
go run main.go
```

4. Ouvrir dans le navigateur :
```
http://localhost:8080
```

## 📁 Structure du projet

```
Portfolio/
├── main.go                    # Serveur Go
├── templates/
│   └── index.html            # Template HTML principal
├── static/
│   ├── css/
│   │   └── style.css        # Styles complets
│   ├── js/
│   │   └── main.js          # JavaScript interactif
│   └── images/              # Images du portfolio
├── PROMPTS.md               # Documentation des prompts utilisés
└── README.md                # Ce fichier
```

## 🎨 Personnalisation

### 1. Contenu
Remplacez tous les placeholders `[...]` dans `templates/index.html` :
- Nom, titre, description
- Compétences et niveaux
- Projets et descriptions
- Informations de contact

### 2. Images
Ajoutez vos images dans `static/images/` et remplacez les placeholders.

### 3. Couleurs
Modifiez les variables CSS dans `static/css/style.css` :
```css
:root {
    --primary-color: #6366f1;
    --secondary-color: #8b5cf6;
    --accent-color: #ec4899;
}
```

### 4. Police
```css
:root {
    --font-main: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto;
    --font-heading: 'Georgia', serif;
}
```

## 🎯 Fonctionnalités détaillées

### Mode Sombre/Clair
- Toggle via bouton dans la navigation
- Préférence sauvegardée dans localStorage
- Transition fluide entre les thèmes

### Carrousel de Projets
- Navigation avec boutons prev/next
- Indicateurs cliquables
- Support tactile (swipe) sur mobile
- Responsive : 3 slides (desktop), 2 (tablette), 1 (mobile)

### Filtrage des Projets
- Filtres par catégorie (web, mobile, design, etc.)
- Animation fade-in des cartes
- Bouton actif visuellement

### Validation Formulaire
- Validation en temps réel pendant la saisie
- Messages d'erreur spécifiques par champ
- Validation finale avant soumission
- Message de succès après envoi

### Compteur de Visites
- Incrémente à chaque visite de page
- Sauvegardé dans localStorage
- Animation de comptage

### Animations au Scroll
- Apparition progressive des sections
- Animation des barres de compétences
- Utilisation d'IntersectionObserver pour les performances

## 📱 Responsive Design

- **Desktop** (>1024px) : Layout complet, 3 projets visibles
- **Tablette** (768px-1024px) : Layout adapté, 2 projets visibles
- **Mobile** (<768px) : Layout vertical, 1 projet visible, navigation simplifiée
- **Petit mobile** (<480px) : Textes optimisés, espacements réduits

## 🌐 Compatibilité

- Chrome/Edge 90+
- Firefox 88+
- Safari 14+
- Mobile browsers (iOS Safari, Chrome Mobile)

## 📚 Documentation

Consultez `PROMPTS.md` pour :
- Liste complète des prompts utilisés
- Résultats obtenus pour chaque prompt
- Détails techniques de chaque fonctionnalité
- Guide de personnalisation approfondi

## 🎓 Compétences démontrées

Ce projet démontre la maîtrise de :
- **Go** : Serveur HTTP, routing, templating
- **HTML5** : Sémantique, accessibilité, SEO
- **CSS3** : Variables, Grid/Flexbox, animations, responsive
- **JavaScript** : DOM, events, LocalStorage, APIs modernes
- **UX/UI** : Design moderne, animations fluides, feedback utilisateur

## 🚧 Améliorations futures

- [ ] Envoyer des emails via le formulaire (SMTP)
- [ ] Base de données pour stocker les messages
- [ ] Progressive Web App (PWA)
- [ ] Lazy loading des images
- [ ] Google Analytics
- [ ] Internationalisation (i18n)

## 📄 Licence

Ce projet est libre de droits. Vous pouvez l'utiliser, le modifier et le distribuer librement.

## 🤝 Contribution

Les suggestions et améliorations sont les bienvenues !

---

**Créé avec ❤️ en utilisant Go et des technologies web modernes**