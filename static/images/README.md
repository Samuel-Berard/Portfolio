# Images du Portfolio

Placez ici vos images pour le portfolio :

## Images recommandées

### 1. Photo de profil
- **Nom suggéré** : `profile.jpg` ou `profile.png`
- **Dimensions recommandées** : 400x400px (carré)
- **Poids max** : 500KB
- **Emplacement dans le code** : Section Hero

### 2. Images des projets
- **Noms suggérés** : `project-1.jpg`, `project-2.jpg`, etc.
- **Dimensions recommandées** : 800x600px (ratio 4:3)
- **Poids max** : 300KB par image
- **Quantité** : 6 projets par défaut (modifiable)
- **Emplacement dans le code** : Section Projets (carrousel)

### 3. Logos/Icônes (optionnel)
- **Noms suggérés** : `logo.svg`, `favicon.ico`
- **Format** : SVG ou PNG transparent
- **Emplacement** : Navigation, favicon

## 📝 Comment ajouter vos images

1. Placez vos fichiers images dans ce dossier (`static/images/`)

2. Dans `templates/index.html`, remplacez les placeholders :

```html
<!-- Exemple pour la photo de profil -->
<div class="image-placeholder">[Image de profil ici]</div>
<!-- Devient : -->
<img src="/static/images/profile.jpg" alt="Votre nom">
```

```html
<!-- Exemple pour un projet -->
<div class="image-placeholder">[Image du projet 1 ici]</div>
<!-- Devient : -->
<img src="/static/images/project-1.jpg" alt="Nom du projet">
```

## 🎨 Optimisation des images

Pour de meilleures performances :

1. **Compression** : Utilisez des outils comme TinyPNG, ImageOptim
2. **Format** : 
   - JPG pour les photos
   - PNG pour les images avec transparence
   - WebP pour une meilleure compression (moderne)
3. **Dimensions** : Redimensionnez avant d'uploader
4. **Lazy loading** : Les images se chargeront au scroll

## 🔗 Ressources d'images gratuites

Si vous avez besoin d'images temporaires :

- **Unsplash** : https://unsplash.com/ (photos haute qualité)
- **Pexels** : https://www.pexels.com/ (photos gratuites)
- **Lorem Picsum** : https://picsum.photos/ (placeholders)
- **UI Faces** : https://uifaces.co/ (avatars)

## 📐 Template de structure

```
static/images/
├── profile.jpg              # Votre photo de profil
├── project-1.jpg            # Premier projet
├── project-2.jpg            # Deuxième projet
├── project-3.jpg            # Troisième projet
├── project-4.jpg            # Quatrième projet
├── project-5.jpg            # Cinquième projet
├── project-6.jpg            # Sixième projet
├── logo.svg                 # Logo (optionnel)
└── favicon.ico              # Favicon (optionnel)
```

---

**Note** : Ce dossier est actuellement vide. Ajoutez vos images personnelles pour compléter votre portfolio.
