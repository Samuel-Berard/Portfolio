# Documentation des Prompts - Portfolio Web

Ce document contient tous les prompts utilisés pour créer ce site web Portfolio/CV et les résultats obtenus.

---

## 📋 Table des matières

1. [Structure de base](#1-structure-de-base)
2. [Template HTML](#2-template-html)
3. [CSS Moderne](#3-css-moderne)
4. [JavaScript Interactif](#4-javascript-interactif)
5. [Résumé des fonctionnalités](#résumé-des-fonctionnalités)

---

## 1. Structure de base

### Prompt 1.1 : Serveur Go et structure de dossiers

**Prompt utilisé :**
```
Créer un serveur web en Go avec :
- Un fichier main.go qui sert des templates HTML
- Route principale "/" qui affiche index.html
- Route "/contact" pour gérer les soumissions de formulaire
- Servir les fichiers statiques depuis un dossier "static"
- Structure de dossiers : templates/, static/css/, static/js/, static/images/
```

**Résultat obtenu :**
- ✅ Fichier `main.go` créé avec serveur HTTP
- ✅ Gestion des routes et templates
- ✅ Validation côté serveur pour le formulaire de contact
- ✅ Structure de dossiers complète créée

**Fichiers créés :**
- `main.go`
- `templates/` (dossier)
- `static/css/` (dossier)
- `static/js/` (dossier)
- `static/images/` (dossier)

---

## 2. Template HTML

### Prompt 2.1 : Page HTML avec sections complètes

**Prompt utilisé :**
```
Créer un template HTML5 avec :
- Navigation avec liens vers sections et bouton de thème
- Section Hero avec titre, description, boutons CTA et placeholder d'image
- Compteur de visites
- Section "À propos" avec texte et statistiques
- Section "Compétences" avec catégories et barres de progression
- Section "Projets" avec :
  - Boutons de filtrage par catégorie
  - Carrousel de cartes de projets (6 projets)
  - Indicateurs de navigation
- Section "Contact" avec :
  - Informations de contact
  - Liens réseaux sociaux
  - Formulaire avec validation (nom, email, message)
- Footer

Tous les contenus doivent être des placeholders entre crochets.
```

**Résultat obtenu :**
- ✅ Structure HTML sémantique complète
- ✅ Navigation responsive avec bouton de thème
- ✅ Hero section avec deux colonnes
- ✅ Compteur de visites intégré
- ✅ Sections À propos, Compétences, Projets, Contact
- ✅ Carrousel avec 6 projets et attributs data-category
- ✅ Formulaire avec champs et messages d'erreur
- ✅ Tous les contenus sont des placeholders

**Fichier créé :**
- `templates/index.html`

---

## 3. CSS Moderne

### Prompt 3.1 : Design moderne et responsive

**Prompt utilisé :**
```
Créer un fichier CSS moderne avec :
- Variables CSS pour couleurs, espacements, typographies
- Mode sombre avec classe .dark-mode sur body
- Design moderne avec gradients et ombres
- Navigation sticky avec styles
- Hero section avec layout grid deux colonnes
- Styles pour toutes les sections
- Compétences affichées en badges (sans barres de progression)
- Carrousel de projets stylisé avec boutons de navigation
- Formulaire moderne avec validation visuelle
- Animations : fadeIn, fadeInUp, slideInRight
- Design entièrement responsive :
  - Desktop (>1024px)
  - Tablette (768px-1024px)
  - Mobile (<768px)
  - Petit mobile (<480px)
```

**Résultat obtenu :**
- ✅ Variables CSS organisées pour thèmes clair/sombre
- ✅ Palette de couleurs moderne (violet/indigo)
- ✅ Layout responsive avec grid et flexbox
- ✅ Animations CSS fluides
- ✅ Transitions douces (0.3s)
- ✅ Styles de formulaire avec états (focus, error)
- ✅ Media queries pour tous les breakpoints
- ✅ Carrousel stylisé avec indicateurs
- ✅ Badges de compétences avec hover effects
- ✅ Ombres et effets modernes

**Fichier créé :**
- `static/css/style.css`

---

## 4. JavaScript Interactif

### Prompt 4.1 : Toutes les fonctionnalités JavaScript

**Prompt utilisé :**
```
Créer un fichier JavaScript complet avec :

1. Mode sombre/clair :
   - Toggle avec bouton
   - Sauvegarde dans localStorage
   - Animation de transition

2. Compteur de visites :
   - Utiliser localStorage
   - Incrémenter à chaque visite
   - Animation de comptage

3. Carrousel de projets :
   - Navigation prev/next
   - Indicateurs cliquables
   - Support tactile (swipe)
   - Responsive (3 slides desktop, 2 tablette, 1 mobile)
   - Mise à jour automatique au resize

4. Filtrage des projets :
   - Filtrer par catégorie (data-attribute)
   - Bouton actif visuel
   - Animation fadeIn sur les cartes

5. Validation formulaire en temps réel :
   - Validation à la saisie et au blur
   - Messages d'erreur spécifiques
   - Validation finale à la soumission
   - Message de succès

6. Animations au scroll :
   - Intersection Observer
   - Animation des sections au scroll
   - Fade in progressif

7. Smooth scroll :
   - Navigation fluide vers sections
   - Prise en compte de la hauteur du navbar

Inclure des logs console pour confirmation du chargement.
```

**Résultat obtenu :**
- ✅ Mode sombre persistant avec localStorage
- ✅ Compteur de visites avec animation
- ✅ Carrousel entièrement fonctionnel
- ✅ Support tactile pour mobile
- ✅ Filtrage dynamique des projets
- ✅ Validation en temps réel complète
- ✅ Animations au scroll avec IntersectionObserver
- ✅ Smooth scroll avec offset du navbar
- ✅ Gestion des événements optimisée
- ✅ Code bien commenté et organisé

**Fichier créé :**
- `static/js/main.js`

---

## Résumé des fonctionnalités

### ✅ Fonctionnalités de base (Partie 1)
- [x] Structure HTML avec sections : À propos, Compétences, Projets, Contact
- [x] CSS moderne avec layout, couleurs, typographies
- [x] Génération par prompts uniquement

### ✅ Interactivité (Partie 2)
- [x] Formulaire de contact fonctionnel (avec validation serveur Go)
- [x] Carrousel de projets en JavaScript
- [x] Bouton mode sombre/clair avec persistance
- [x] Animations d'entrée (fade-in, slide, etc.)

### ✅ Fonctionnalités avancées (Partie 3)
- [x] Système de filtrage des projets par catégorie
- [x] Validation du formulaire en temps réel
- [x] Compteur de visites via localStorage
- [x] Design entièrement responsive (mobile/tablette/desktop)

### ✅ Personnalisation & optimisation (Partie 4)
- [x] Design graphique moderne avec gradients
- [x] Palette de couleurs cohérente
- [x] Animations fluides et performantes
- [x] Code CSS organisé avec variables
- [x] JavaScript modulaire et commenté

---

## 🚀 Comment lancer le projet

### Prérequis
- Go 1.16+ installé

### Installation et lancement

1. Naviguer vers le dossier du projet :
```bash
cd "c:\Users\samba\Desktop\Y boost\Portfolio"
```

2. Lancer le serveur Go :
```bash
go run main.go
```

3. Ouvrir le navigateur à l'adresse :
```
http://localhost:8080
```

### Structure finale du projet

```
Portfolio/
├── main.go                    # Serveur Go
├── templates/
│   └── index.html            # Template principal
├── static/
│   ├── css/
│   │   └── style.css        # Styles complets
│   ├── js/
│   │   └── main.js          # JavaScript interactif
│   └── images/              # Placez vos images ici
├── PROMPTS.md               # Ce fichier
└── README.md                # Documentation du projet
```

---

## 📝 Personnalisation du contenu

Pour personnaliser le site avec votre contenu :

1. **Textes** : Remplacer tous les placeholders `[...]` dans `templates/index.html`
2. **Images** : Ajouter vos images dans `static/images/` et mettre à jour les placeholders
3. **Couleurs** : Modifier les variables CSS dans `:root` dans `style.css`
4. **Projets** : Ajouter/modifier les cartes de projets dans la section projects

---

## 🎨 Personnalisation du design

### Changer les couleurs principales

Dans `static/css/style.css`, modifier les variables :

```css
:root {
    --primary-color: #6366f1;      /* Couleur principale */
    --secondary-color: #8b5cf6;    /* Couleur secondaire */
    --accent-color: #ec4899;       /* Couleur d'accent */
}
```

### Changer les polices

```css
:root {
    --font-main: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto;
    --font-heading: 'Georgia', serif;
}
```

### Ajuster les espacements

```css
:root {
    --spacing-xs: 0.5rem;
    --spacing-sm: 1rem;
    --spacing-md: 2rem;
    --spacing-lg: 4rem;
    --spacing-xl: 6rem;
}
```

---

## 🔧 Fonctionnalités détaillées

### 1. Mode Sombre/Clair
- **Stockage** : localStorage clé `theme`
- **Classes** : `.dark-mode` sur `<body>`
- **Transition** : Automatique avec CSS transitions

### 2. Compteur de visites
- **Stockage** : localStorage clé `visitCount`
- **Animation** : Comptage progressif de 0 au nombre actuel
- **Persistance** : Conservé entre les sessions

### 3. Carrousel
- **Navigation** : Boutons prev/next + indicateurs
- **Tactile** : Support swipe gauche/droite
- **Responsive** : 
  - Desktop: 3 slides
  - Tablette: 2 slides
  - Mobile: 1 slide

### 4. Filtrage projets
- **Méthode** : Attributs `data-category` sur les cartes
- **Catégories** : all, web, mobile, design (personnalisables)
- **Animation** : Fade in lors de l'affichage

### 5. Validation formulaire
- **Champs validés** :
  - Nom : minimum 2 caractères
  - Email : format valide
  - Message : minimum 10 caractères
- **Validation** : En temps réel + au blur + à la soumission
- **Feedback** : Messages d'erreur spécifiques sous chaque champ

### 6. Animations scroll
- **Méthode** : IntersectionObserver API
- **Éléments animés** : Tous les `.fade-in-up`
- **Barres de compétences** : Animation au scroll

### 7. Smooth scroll
- **Méthode** : `scrollTo()` avec `behavior: 'smooth'`
- **Offset** : Prend en compte la hauteur du navbar sticky

---

## 🎯 Checklist de personnalisation

Avant de publier votre portfolio, remplacez tous les placeholders :

### Hero Section
- [ ] [Entrez votre nom/logo ici]
- [ ] [Entrez votre nom complet ici]
- [ ] [Entrez votre titre/profession ici]
- [ ] [Entrez une brève description de vous ici]
- [ ] [Image de profil ici]

### Section À propos
- [ ] [Entrez un titre de présentation ici]
- [ ] 3 paragraphes de présentation
- [ ] 3 statistiques personnelles

### Section Compétences
- [ ] 3 catégories de compétences
- [ ] 9 compétences avec niveaux (3 par catégorie)

### Section Projets
- [ ] 6 projets avec :
  - [ ] Images
  - [ ] Titres
  - [ ] Descriptions
  - [ ] Technologies utilisées
  - [ ] Liens (projet + code source)

### Section Contact
- [ ] Message d'invitation
- [ ] Email
- [ ] Téléphone
- [ ] Localisation
- [ ] 4 liens réseaux sociaux

### Footer
- [ ] Votre nom

---

## 📚 Technologies utilisées

- **Backend** : Go (Golang)
- **Frontend** : HTML5, CSS3, JavaScript (ES6+)
- **APIs** : 
  - IntersectionObserver (animations)
  - LocalStorage (compteur + thème)
  - Touch Events (carousel mobile)

---

## 🌐 Compatibilité navigateurs

- ✅ Chrome/Edge (90+)
- ✅ Firefox (88+)
- ✅ Safari (14+)
- ✅ Mobile browsers (iOS Safari, Chrome Mobile)

---

## 📱 Responsive Breakpoints

- **Desktop** : > 1024px (3 colonnes projets)
- **Tablette** : 768px - 1024px (2 colonnes projets)
- **Mobile** : < 768px (1 colonne, navigation adaptée)
- **Petit mobile** : < 480px (textes réduits)

---

## 🎓 Apprentissages du projet

Ce projet vous a permis de maîtriser :

1. **Go (Backend)**
   - Création de serveur HTTP
   - Templating HTML
   - Gestion des routes
   - Validation de formulaires

2. **HTML5**
   - Structure sémantique
   - Attributs data-*
   - Formulaires accessibles
   - SEO-friendly markup

3. **CSS3**
   - Variables CSS
   - Grid & Flexbox
   - Animations & Transitions
   - Media queries
   - Mode sombre

4. **JavaScript**
   - DOM Manipulation
   - Event Listeners
   - LocalStorage
   - IntersectionObserver
   - Touch Events
   - Form Validation
   - Carousel Logic

5. **UX/UI**
   - Design responsive
   - Animations fluides
   - Feedback utilisateur
   - Navigation intuitive

---

## 🚧 Améliorations possibles

Pour aller plus loin :

1. **Backend**
   - [ ] Envoyer des emails via le formulaire (SMTP)
   - [ ] Base de données pour stocker les messages
   - [ ] API REST pour les projets

2. **Frontend**
   - [ ] Progressive Web App (PWA)
   - [ ] Lazy loading des images
   - [ ] Optimisation des performances (Lighthouse)
   - [ ] Internationalisation (i18n)

3. **SEO**
   - [ ] Meta tags OpenGraph
   - [ ] Sitemap.xml
   - [ ] Schema.org markup

4. **Analytics**
   - [ ] Google Analytics
   - [ ] Heatmaps

---

## 📄 Licence

Ce projet est libre de droits. Vous pouvez l'utiliser, le modifier et le distribuer librement.

---

## ✨ Conclusion

Vous avez maintenant un portfolio complet, moderne et fonctionnel ! 

Toutes les fonctionnalités demandées ont été implémentées :
- ✅ Structure HTML complète
- ✅ CSS moderne et responsive
- ✅ Interactivité JavaScript avancée
- ✅ Mode sombre/clair
- ✅ Carrousel de projets
- ✅ Filtrage par catégorie
- ✅ Validation en temps réel
- ✅ Compteur de visites
- ✅ Animations fluides

**Prochaine étape** : Personnalisez le contenu avec vos informations et publiez votre portfolio !

Bon développement ! 🚀

---

## 5. Améliorations et Personnalisations

### Prompt 5.1 : Suppression des barres de progression

**Prompt utilisé :**
```
peut tu retirer les barre de progression des skill car je sais aps quel porucentage metrre vu que c'est assez subjectif
```

**Besoin identifié :**
L'utilisateur trouve que les pourcentages de compétences sont trop subjectifs et souhaite une présentation plus objective.

**Résultat obtenu :**
- ✅ Suppression des barres de progression (`skill-bar`, `skill-progress`)
- ✅ Remplacement par un système de badges (`skill-badge`)
- ✅ Mise à jour du HTML pour retirer les éléments de progression
- ✅ Mise à jour du CSS pour styliser les badges uniquement
- ✅ Mise à jour du JavaScript pour retirer l'animation des barres

**Fichiers modifiés :**
- `templates/index.html`
- `static/css/style.css`
- `static/js/main.js`

---

### Prompt 5.2 : Aide pour remplir les compétences

**Prompt utilisé :**
```
peut tu m'aider a remplir mes competence stp
```

**Contexte fourni par l'utilisateur :**
- Technologies : Golang, Rust, Java, HTML, CSS, JavaScript, Python, C#, React, Node.js, Git, Docker, PostgreSQL, MongoDB

**Résultat obtenu :**
- ✅ Organisation des compétences en 3 catégories :
  - **Backend & Systèmes** : Golang, Rust, Java, Python, C#
  - **Frontend & Web** : HTML, CSS, JavaScript, React, Node.js
  - **Outils & DevOps** : Git, Docker, PostgreSQL, MongoDB
- ✅ Intégration dans le template HTML

**Fichiers modifiés :**
- `templates/index.html`

---

### Prompt 5.3 : Alignement du logo dans la navigation

**Prompt utilisé :**
```
mon logo et [Entrez votre nom/logo ici] ne sont pas sur la meme ligne
```

**Problème identifié :**
Le logo (image) et le texte dans la navigation n'étaient pas alignés horizontalement.

**Résultat obtenu :**
- ✅ Ajout de `display: flex` et `align-items: center` à `.logo`
- ✅ Ajout de `gap: 0.5rem` pour l'espacement entre l'image et le texte
- ✅ Définition de `height: 1.5rem` pour l'image du logo

**Fichiers modifiés :**
- `static/css/style.css`

---

### Prompt 5.4 : Ajustement des images de projets

**Prompt utilisé :**
```
peux tu faire en sorte que les image ne soient pas crop mais bien visible
```

**Problème identifié :**
Les images de projets étaient recadrées (cropped) avec `object-fit: cover`, ce qui masquait des parties importantes des images.

**Résultat obtenu :**
- ✅ Changement de `object-fit: cover` à `object-fit: contain`
- ✅ Ajout de `object-position: center` pour centrer l'image
- ✅ Les images sont maintenant entièrement visibles sans recadrage

**Fichiers modifiés :**
- `static/css/style.css`

---

### Prompt 5.5 : Correction de l'apparence des projets filtrés

**Prompt utilisé :**
```
quand je les filtre par categorie les projets change d'apparence
```

**Problème identifié :**
Le JavaScript appliquait `display: 'flex'` aux projets filtrés, ce qui changeait leur apparence par rapport à l'affichage initial.

**Résultat obtenu :**
- ✅ Remplacement de `display: 'flex'` par `display: ''` (chaîne vide)
- ✅ Permet au CSS d'appliquer l'affichage par défaut
- ✅ Préserve l'apparence d'origine des cartes de projet

**Fichiers modifiés :**
- `static/js/main.js`

---

### Prompt 5.6 : Filtrage multi-catégories

**Prompt utilisé :**
```
est il possible que les projet ai plusieurs categorie
```

**Besoin identifié :**
Permettre à un projet d'appartenir à plusieurs catégories simultanément.

**Résultat obtenu :**
- ✅ Modification du JavaScript pour supporter les catégories multiples
- ✅ Utilisation de `split(' ')` pour séparer les catégories dans `data-category`
- ✅ Vérification avec `includes()` pour le filtrage
- ✅ Exemple d'utilisation : `data-category="web mobile"` pour un projet qui appartient aux deux catégories

**Fichiers modifiés :**
- `static/js/main.js`

---

### Prompt 5.7 : Renommage des catégories

**Prompt utilisé :**
```
peut tu changer le nom des filtre : web -> golang mobile -> web design -> java
```

**Résultat obtenu :**
- ✅ Mise à jour de tous les attributs `data-filter` dans les boutons
- ✅ Mise à jour de tous les attributs `data-category` dans les cartes de projet
- ✅ Nouvelles catégories : `golang`, `web`, `java`

**Fichiers modifiés :**
- `templates/index.html`

---

### Prompt 5.8 : Intégration des icônes de réseaux sociaux

**Prompt utilisé :**
```
je voudrait que tu integre les icone des reseau sociaux : Linkedin, github, twitter et mon site personnel en forme de personne
```

**Résultat obtenu :**
- ✅ Intégration d'icônes SVG inline pour :
  - LinkedIn (logo officiel)
  - GitHub (logo officiel)
  - Twitter (logo officiel)
  - Portfolio personnel (icône de personne)
- ✅ Styles CSS pour les icônes (taille, couleurs, hover effects)
- ✅ Remplacement des placeholders textuels par des icônes visuelles

**Fichiers modifiés :**
- `templates/index.html`
- `static/css/style.css`

---

### Prompt 5.9 : Simplification des boutons de projets

**Prompt utilisé :**
```
peut tu faire en sorte qu'il n'y ai plus qu'un bouton avec l'icone github et ecrit voir le code source
```

**Résultat obtenu :**
- ✅ Suppression du bouton "Voir le projet"
- ✅ Remplacement par un seul bouton "Voir le code source"
- ✅ Ajout de l'icône GitHub SVG dans le bouton
- ✅ Nouveau style avec classe `.github-link` (gradient violet/rose)
- ✅ Effets hover modernes

**Fichiers modifiés :**
- `templates/index.html`
- `static/css/style.css`

---

### Prompt 5.10 : Remplissage de la section À propos

**Prompt utilisé :**
```
j'aimerait que tu remplisse la section a propos pour moi je te donne mes info [...]
Je suis actuellement en terminale et je passe mon bac en 2025. J'ai rejoins Ynov l'année prochaine [...]
Je travaille actuellement chez decathlon [...]
j'aime jouer au jeux video les voiture la technologie et le sport
```

**Contexte fourni :**
- Actuellement en Terminale, Bac 2025
- Rejoindra Ynov l'année prochaine en Informatique
- Travaille chez Décathlon
- Passions : jeux vidéo, voitures, technologie, sport

**Résultat obtenu :**
- ✅ Rédaction d'un texte naturel et authentique en 3 paragraphes :
  1. Présentation et parcours scolaire
  2. Expérience professionnelle chez Décathlon
  3. Passions et hobbies
- ✅ Mise à jour des statistiques pertinentes
- ✅ Style d'écriture personnel et engageant

**Fichiers modifiés :**
- `templates/index.html`

---

### Prompt 5.11 : Enregistrement des formulaires en fichier texte

**Prompt utilisé :**
```
peut tu faire en sorte que le formulaire une fois rempli et envoyé s'enrgistre sous sous format txt ?
```

**Besoin identifié :**
Sauvegarder les soumissions du formulaire de contact dans des fichiers texte pour consultation ultérieure.

**Première implémentation (serveur) :**
- ✅ Ajout des imports nécessaires : `fmt`, `os`, `time`
- ✅ Création d'une fonction `saveContactToFile(name, email, message string) error`
- ✅ Création automatique d'un dossier `contacts/`
- ✅ Génération d'un fichier unique par soumission : `contact_YYYY-MM-DD_HH-MM-SS_NomPrenom.txt`
- ✅ Format structuré avec date, nom, email et message

**Fichiers modifiés :**
- `main.go`

---

### Prompt 5.12 : Téléchargement d'une copie après envoi

**Prompt utilisé :**
```
telecharger une copie après l'envoi st(p)
```

**Besoin identifié :**
En plus de l'enregistrement serveur, proposer au navigateur de télécharger une copie du fichier.

**Résultat obtenu :**
- ✅ Modification de la fonction `saveContactToFile` pour retourner le nom de fichier et le contenu
- ✅ Ajout des headers HTTP pour forcer le téléchargement
- ✅ `Content-Disposition: attachment` pour déclencher le téléchargement
- ✅ Le fichier est enregistré sur le serveur ET téléchargé dans le navigateur

**Fichiers modifiés :**
- `main.go`

---

### Prompt 5.13 : Téléchargement local uniquement

**Prompt utilisé :**
```
ça me propose rien finalement fais moi juste un telkechargementr en locla stp pas de stockazge serveur
```

**Problème identifié :**
Le téléchargement serveur ne fonctionnait pas comme prévu. L'utilisateur préfère un téléchargement 100% côté client sans passer par le serveur.

**Résultat obtenu :**
- ✅ Suppression de l'attribut `action` et `method` du formulaire HTML
- ✅ Modification complète de la fonction `submitForm()` en JavaScript
- ✅ Création du contenu du fichier directement dans le navigateur
- ✅ Utilisation de l'API Blob pour créer le fichier
- ✅ Utilisation de `URL.createObjectURL()` pour générer un lien de téléchargement temporaire
- ✅ Déclenchement automatique du téléchargement via `a.click()`
- ✅ Nettoyage de l'URL temporaire avec `URL.revokeObjectURL()`
- ✅ Format de fichier : `contact_YYYY-MM-DDTHH-MM-SS_NomPrenom.txt`
- ✅ Format du contenu identique (date française, structure avec séparateurs)

**Avantages de cette solution :**
- ❌ Aucun stockage serveur
- ✅ Téléchargement instantané et automatique
- ✅ Fonctionne même hors ligne
- ✅ Plus simple et plus rapide
- ✅ Pas de dépendance au backend Go

**Fichiers modifiés :**
- `templates/index.html` (suppression de `action="/contact" method="POST"`)
- `static/js/main.js` (réécriture complète de `submitForm()`)

---

## 🔄 Historique des modifications

| Date | Modification | Fichiers affectés |
|------|-------------|-------------------|
| Initial | Création du portfolio complet | Tous les fichiers |
| Nov 2025 | Suppression barres de progression | HTML, CSS, JS |
| Nov 2025 | Remplissage des compétences | HTML |
| Nov 2025 | Alignement logo navigation | CSS |
| Nov 2025 | Images projets en mode contain | CSS |
| Nov 2025 | Correction filtrage projets | JS |
| Nov 2025 | Support multi-catégories | JS |
| Nov 2025 | Renommage catégories | HTML |
| Nov 2025 | Intégration icônes sociales | HTML, CSS |
| Nov 2025 | Bouton GitHub unique | HTML, CSS |
| Nov 2025 | Remplissage section À propos | HTML |
| Nov 2025 | Enregistrement formulaires (serveur) | main.go |
| Nov 2025 | Téléchargement copie (serveur) | main.go |
| Nov 2025 | **Téléchargement local (client)** | HTML, JS |

---

## 💡 Notes importantes

### Gestion des formulaires
Le système utilise maintenant un téléchargement **100% côté client** :
- Le formulaire ne fait plus appel au serveur Go
- Le fichier texte est généré directement dans le navigateur avec l'API Blob
- Le téléchargement se déclenche automatiquement après validation
- Format du fichier : `contact_YYYY-MM-DDTHH-MM-SS_NomPrenom.txt`
- Le contenu est structuré avec date française, nom, email et message

### Évolutions demandées par l'utilisateur
1. **Objectivité** : Retrait des pourcentages de compétences jugés subjectifs
2. **Multi-catégories** : Projets pouvant appartenir à plusieurs catégories
3. **Personnalisation** : Contenu authentique reflétant le parcours réel de l'utilisateur
4. **Simplicité visuelle** : Un seul bouton par projet, icônes claires
5. **Persistance des données** : ~~Sauvegarde des messages de contact~~ → Téléchargement local instantané

---

**Document mis à jour le : 28 novembre 2025**
