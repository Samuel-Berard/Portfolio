# 🚀 Guide de Démarrage Rapide

## Étape 1 : Vérifier les prérequis

Assurez-vous que Go est installé sur votre système :

```bash
go version
```

Si Go n'est pas installé, téléchargez-le depuis : https://golang.org/dl/

## Étape 2 : Lancer le serveur

Dans le terminal, naviguez vers le dossier du projet et lancez :

```powershell
# PowerShell (Windows)
cd "c:\Users\samba\Desktop\Y boost\Portfolio"
go run main.go
```

```bash
# Bash (Linux/Mac)
cd /path/to/Portfolio
go run main.go
```

Vous devriez voir :
```
2025/11/26 [heure] Serveur démarré sur http://localhost:8080
```

## Étape 3 : Ouvrir dans le navigateur

Ouvrez votre navigateur et allez à :
```
http://localhost:8080
```

## Étape 4 : Tester les fonctionnalités

### ✅ Checklist de test

1. **Navigation**
   - [ ] Cliquer sur les liens de navigation (À propos, Compétences, etc.)
   - [ ] Vérifier le scroll fluide vers les sections

2. **Mode Sombre/Clair**
   - [ ] Cliquer sur le bouton ☀️/🌙 en haut à droite
   - [ ] Vérifier que le thème change
   - [ ] Recharger la page → le thème doit être conservé

3. **Compteur de visites**
   - [ ] Vérifier que le compteur s'affiche
   - [ ] Recharger la page → le compteur augmente

4. **Carrousel de projets**
   - [ ] Cliquer sur les boutons ‹ et ›
   - [ ] Cliquer sur les indicateurs (points en bas)
   - [ ] Sur mobile : swiper gauche/droite

5. **Filtrage des projets**
   - [ ] Cliquer sur les boutons de filtres
   - [ ] Vérifier que les projets se filtrent

6. **Formulaire de contact**
   - [ ] Essayer de soumettre le formulaire vide → erreurs affichées
   - [ ] Entrer un email invalide → erreur affichée
   - [ ] Remplir correctement → message de succès

7. **Animations**
   - [ ] Scroller la page → les sections apparaissent progressivement
   - [ ] Section Compétences → barres qui se remplissent

8. **Responsive**
   - [ ] Redimensionner la fenêtre
   - [ ] Tester sur mobile (F12 → mode mobile)

## Étape 5 : Personnaliser le contenu

### 5.1 Modifier les textes

Ouvrez `templates/index.html` et remplacez tous les placeholders `[...]` :

**Exemples de placeholders à remplacer :**
- `[Entrez votre nom/logo ici]`
- `[Entrez votre nom complet ici]`
- `[Entrez votre titre/profession ici]`
- `[Entrez une brève description de vous ici]`
- etc.

### 5.2 Ajouter vos images

1. Placez vos images dans `static/images/`
2. Dans `templates/index.html`, remplacez :
```html
<div class="image-placeholder">[Image de profil ici]</div>
```
par :
```html
<img src="/static/images/votre-photo.jpg" alt="Votre nom">
```

### 5.3 Personnaliser les couleurs

Dans `static/css/style.css`, modifiez les variables CSS :

```css
:root {
    --primary-color: #6366f1;      /* Couleur principale */
    --secondary-color: #8b5cf6;    /* Couleur secondaire */
    --accent-color: #ec4899;       /* Couleur d'accent */
}
```

**Suggestions de palettes :**
- **Bleu** : `#3b82f6`, `#2563eb`, `#1d4ed8`
- **Vert** : `#10b981`, `#059669`, `#047857`
- **Rouge** : `#ef4444`, `#dc2626`, `#b91c1c`
- **Orange** : `#f97316`, `#ea580c`, `#c2410c`

### 5.4 Modifier les projets

Dans `templates/index.html`, section Projets :

```html
<div class="project-card" data-category="web">
    <div class="project-image">
        <img src="/static/images/projet-1.jpg" alt="Mon projet">
    </div>
    <div class="project-content">
        <h3>Nom de mon projet</h3>
        <p>Description de mon projet...</p>
        <div class="project-tags">
            <span class="tag">HTML</span>
            <span class="tag">CSS</span>
            <span class="tag">JavaScript</span>
        </div>
        <div class="project-links">
            <a href="https://mon-projet.com" class="project-link">Voir le projet</a>
            <a href="https://github.com/mon-compte/projet" class="project-link">Code source</a>
        </div>
    </div>
</div>
```

## Étape 6 : Compiler pour la production (optionnel)

Pour créer un exécutable :

```bash
# Windows
go build -o portfolio.exe main.go

# Linux/Mac
go build -o portfolio main.go
```

Ensuite, lancez simplement :
```bash
./portfolio        # Linux/Mac
portfolio.exe      # Windows
```

## 🆘 Dépannage

### Le serveur ne démarre pas

**Erreur : "port already in use"**
- Un autre programme utilise le port 8080
- Solution : Arrêtez l'autre programme ou changez le port dans `main.go` :
```go
log.Fatal(http.ListenAndServe(":3000", nil))  // Utilisez le port 3000
```

**Erreur : "go: command not found"**
- Go n'est pas installé ou pas dans le PATH
- Solution : Installez Go depuis https://golang.org/dl/

### Les styles ne s'appliquent pas

- Vérifiez que le dossier `static/css/style.css` existe
- Ouvrez la console du navigateur (F12) → onglet Network
- Vérifiez que `style.css` se charge sans erreur 404

### Le JavaScript ne fonctionne pas

- Ouvrez la console du navigateur (F12) → onglet Console
- Vérifiez les erreurs JavaScript
- Vérifiez que `static/js/main.js` existe

### Les images ne s'affichent pas

- Vérifiez le chemin : `/static/images/nom-image.jpg`
- Vérifiez que les images sont bien dans le dossier `static/images/`
- Vérifiez les permissions de lecture sur les fichiers

## 📚 Ressources supplémentaires

- **Documentation Go** : https://golang.org/doc/
- **MDN Web Docs** : https://developer.mozilla.org/
- **CSS Tricks** : https://css-tricks.com/
- **JavaScript.info** : https://javascript.info/

## 🎯 Prochaines étapes

1. ✅ Personnaliser tout le contenu
2. ✅ Ajouter vos vraies images
3. ✅ Tester sur différents navigateurs
4. ✅ Tester sur mobile
5. ✅ Optimiser les images
6. ✅ Déployer en ligne (Netlify, Vercel, etc.)

## 💡 Conseils

- **Sauvegardez régulièrement** votre travail
- **Testez après chaque modification** importante
- **Utilisez Git** pour versionner votre code
- **Demandez des retours** à vos amis/collègues

---

**Besoin d'aide ?** Consultez `PROMPTS.md` pour plus de détails techniques !

Bon développement ! 🚀
