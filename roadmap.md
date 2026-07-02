# Roadmap — 6 mois vers Top 1% Go + Cert Ardan + Crédibilité Sourcegraph

**Budget** : ~6 h/jour × 5 jours profonds + 1 jour review + 1 jour repos = ~34 h/sem × 26 sem ≈ 900 h.
**Réalistes** (maladie, plateau, imprévus) : ~800 h productifs. Le plan absorbe le reste.

---

## Les 3 piliers

1. **Top 1% Go** — profondeur : mechanical sympathy, sémantique exacte, concurrency, generics, toolchain, stdlib comme cours, tests comme preuve.
2. **Cert Ardan Labs** — tour gratuit → Ultimate Go Bundle ($699/12 mois) → certification formelle sur `ardanlabs.training`.
3. **Crédibilité Sourcegraph** — cibles : `sourcegraph/zoekt`, `sourcegraph/conc`. Chemin : maîtriser Go → OSS Go moyen → PRs zoekt → portfolio → application.

---

## Définition mesurable de "Top 1%"

- [ ] Expliquer l'escape analysis à partir d'un snippet de code
- [ ] Lire et expliquer le source de `sync.Mutex`
- [ ] Écrire un benchmark et interpréter `-benchmem`
- [ ] Déboguer une data race avec `go test -race`
- [ ] Écrire des tests table-driven idiomatiques sans modèle
- [ ] Lire le source de `sourcegraph/zoekt` et comprendre la structure
- [ ] Avoir mergé des PRs dans des projets OSS Go
- [ ] Avoir passé la certification Ardan Labs

---

## Phase 1 — Fondations (Mois 1-2 · semaines 1-8)

**Tour Ardan gratuit — mécanique du langage**, en mode predict → read → test → recall. AUCUNE contribution OSS. Construction de l'habitude.

Chapitres du tour — **ordre canonique** (welcome → error-handling, 19 leçons ; vérifié sur le serveur) :
welcome · variables · struct-types · pointers · constants · **functions** (← tu es ici, leçon 6/40) · arrays · slices · maps · methods · **interfaces** (keystone, 9 ex) · embedding · exporting · composition (grouping, assertions, pollution, mocking, decoupling) · error-handling

Approfondissement **tour-only** (pas de stdlib — elle utilise des constructs vus plus tard : interfaces, embedding, generics. La lire maintenant est prématuré) :
- **Exercices du tour** : chaque chapitre a `exercise1.go` + `answer1.go` (certains ont 2 exercices). predict → tente → vérifie.
- **Tests table-driven** pour chaque example (preuve de compréhension).
- **Recall** : 2-4 cartes par chapitre.
- **Bonus optionnel** (pas un track parallèle) : article blog Bill Kennedy *Language Mechanics* sur le chapitre en cours (ex : chapitre pointers → escape analysis). Seulement si tu veux aller plus loin sur CE chapitre.

**Critères de sortie (tous obligatoires)** :
- [ ] Tour 100% terminé (jusqu'à error-handling)
- [ ] Chaque example a un test table-driven qui passe
- [ ] Exercices du tour tentés (jusqu'à error-handling)
- [ ] Deck recall ≥ 100 cartes
- [ ] Peux expliquer escape analysis de base (stack vs heap, quand ça escape)
- [ ] Peux écrire un test table-driven sans modèle

**Milestones** :
- **M1 (fin S4)** : tour 50%, recall 50+ cartes, 1er exercice du tour résolu sans aide
- **M2 (fin S8)** : tour 100% (mécanique), recall 100+ cartes, escape analysis solide

---

## Phase 2 — Profondeur (Mois 3-4 · semaines 9-16)

**ACHAT du bundle Ultimate Go ($699, 12 mois)**. Modules profonds + suite du tour (concurrence + generics + algorithmes) + premières PRs OSS.

Chapitres du tour — **suite** (ordre canonique, goroutines → algorithms, 21 leçons ; vérifié sur le serveur) :
goroutines · data_race · context · channels · generics (basics, underlying-types, struct-types, behavior-constraints, type-constraints, multi-type-params, slice-constraints, channels, hash-table) · algorithms (bits-seven, strings, numbers, slices, sorting, data, searches, fun)

Bundle Ultimate Go (modules profonds) :
design · concurrency deep · mechanical sympathy · escape analysis avancé · GC · profiling · benchmarking

En parallèle (maintenant que tu as tous les constructs du langage) :
- **Stdlib lu intégralement** : `io`, `strings`, `bytes`, `errors` (déplacé depuis la Phase 1), puis `sync`, `context`, `sync/atomic`.
- **Premières PRs OSS Go** (PAS Sourcegraph) : bug fixes, docs, tests sur projets Go moyens. Minimum 1 PR/mois.

**Critères de sortie** :
- [ ] Bundle 50% terminé
- [ ] 2 PRs OSS mergées
- [ ] Peux profiler un programme et expliquer les résultats
- [ ] Peux écrire un benchmark correct
- [ ] Lu intégralement : `io`, `strings`, `bytes`, `errors`, `sync`, `context`, `sync/atomic`

**Milestones** :
- **M3 (fin S12)** : bundle 30%, 1re PR OSS mergée
- **M4 (fin S16)** : bundle 50%, 2e PR OSS, benchmark + profiling maîtrisés

---

## Phase 3 — Convergence (Mois 5-6 · semaines 17-26)

**Finir le bundle. Premier PR sur `sourcegraph/zoekt`. Passer la certification Ardan. Portfolio. Application Sourcegraph.**

- Finir le bundle Ultimate Go (100%).
- Lire le source de `sourcegraph/zoekt` (moteur de recherche de code, Go-heavy).
- 1re PR sur `zoekt` (ou `conc` en plan B).
- Passer la **certification Ardan Labs** sur `ardanlabs.training`.
- Construire le **portfolio** : 3-5 PRs quality (dont 1+ sur zoekt).
- Préparer la **narrative d'application Sourcegraph** (cert + portfolio + story).

**Critères de sortie** :
- [ ] Bundle 100% terminé
- [ ] Certification Ardan Labs passée
- [ ] 3-5 PRs dans le portfolio (dont 1+ sur zoekt)
- [ ] Application Sourcegraph prête (CV + narrative + portfolio)

**Milestones** :
- **M5 (fin S20)** : bundle 75%, 1re PR zoekt
- **M6 (fin S24)** : bundle 100%, cert Ardan passée
- **M7 (fin S26)** : portfolio complet, application Sourcegraph prête

---

## Décision clé — achat du bundle

Point de décision : **fin semaine 10**. Si pas encore acheté, ne pas attendre — continuer avec articles blog + stdlib + tour jusqu'à l'achat. Pas de honte à décaler. Le bundle n'est pas un prérequis de la Phase 2, c'est un accélérateur.

---

## Contingences — toutes eventualités

| Situation | Action |
|---|---|
| Maladie courte (1-3 j) | Skip Bloc 4-5, faire 1 h recall min pour garder la chaîne. 3 j flex/mois absorbent. |
| Maladie longue (1-2 sem) | Pause. Reprise au même point, recall-only les 3 premiers jours. |
| Urgence familiale | Pause, 0 culpabilité. Reprise : recall-only 1 sem, puis normale. |
| Semaine plateau (bloqué 3+ j) | Changer modalité : 2 j de lecture stdlib seule, pas de nouveau matériel. Revenir frais. |
| Semaine démotivation | 5 j à 3 h/j maintenance (recall + lecture uniquement), puis reprise. Ne pas quitter. |
| PR rejetée | Lire TOUT le feedback, réécrire sous 48 h, resoumettre. C'est le travail. |
| PR bloquée en review | Lancer la PR suivante, ne pas se bloquer. |
| Bundle pas achetable | Continuer ressources gratuites (blog, stdlib), décaler l'achat, 0 honte. |
| Retard > 1 sem sur milestone | NE PAS compresser. Dropper un objectif secondaire (ex : 2e PR OSS) plutôt que bâcler. |
| Avance sur milestone | NE PAS avancer la phase. Aller PLUS PROFOND : plus de tests, plus de stdlib, plus de recall. |
| Chapitre trop dur (> 48 h dessus) | Consulter oracle/mentor. Ne pas grinder. |
| Chapitre trop facile (< 1 j) | Ralentir. Ajouter lecture stdlib, ÉCRIRE PLUS DE TESTS, ajouter carte recall. |
| Imposter syndrome (mois 3-4) | Regarder la croissance du deck recall + PRs mergées. Continuer. |
| Signes burnout (insomnie, anxiété) | 1 sem REPOS COMPLET obligatoire. Reprise à 4 h/j pendant 2 sem. |
| Tour server down | Redémarrer le binaire `tour`. Pas bloquant. |
| Pas d'internet (contributions OSS) | Switch en profondeur locale (lecture stdlib, exercices escape analysis). |
| Mort du PC | Backup quotidien du repo vers git remote. Roadmap/predictions/recall = sacrés. |
| Nouvelle version Go | Rester sur 1.26.3 sauf si module tour l'exige. Ne pas chaser les features. |
| Offre Sourcegraph avant M7 | NE PAS postuler avant cert + portfolio. Postuler = brûler le pont. |
| Refus Sourcegraph | Demander feedback. Rebuild 3 mois. Re-postuler avec nouveau portfolio. |

---

## Anti-patterns — refuser absolument

- Lire sans prédire
- Implémenter sans tests
- Sauter le recall
- Travailler 7 j/sem (repos dominical = architecture, pas option)
- Compresser pour rattraper
- Avancer sans franchir le portail de profondeur (prédictions confrontées + exercice tenté + tests verts + cartes recall)
- Rusher les keystones (interfaces, composition) pour tenir une cadence de 2/jour
- PR OSS avant le Mois 3 (profondeur d'abord)
- Postuler Sourcegraph avant M7
