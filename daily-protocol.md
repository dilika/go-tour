# Protocole quotidien — 6 h, 5 blocs cognitifs

## Structure journalière (jours profonds : lundi-vendredi)

| Bloc | Durée | Quand | Activité | Mode cognitif |
|------|-------|------|----------|---------------|
| 1 | 90 min | matin (focus max) | Amorcer (5 min) → Prédire → Lire | Profond |
| 2 | 90 min | matin | Réécrire de mémoire + écrire tests table-driven comme preuve | Production |
| 3 | 60 min | après-midi (post-déjeuner) | Approfondir le tour : exercice du chapitre (Phase 1) / Stdlib (Phase 2+) | Réception |
| 4 | 90 min | après-midi | Session recall (chapitres passés, espacé) + temps OSS (à partir Mois 3) | Mixte |
| 5 | 30 min | fin de journée | MAJ `daily-log.md` + planifier demain | Méta |

**Total** : 6 h avec variété cognitive (profond → production → réception → mixte → méta). La variété empêche l'épuisement.

**Bloc 3 dépend de la phase** :
- **Phase 1 (tour)** : exercice du chapitre (predict → tente `exercise1.go` → vérifie contre `answer1.go`). PAS de lecture stdlib — le tour d'abord, profondément. La stdlib utilise des constructs vus plus tard (interfaces, embedding, generics) ; la lire maintenant est prématuré.
- **Phase 2+** : lecture stdlib sur le même sujet (tu as alors tous les constructs du langage).

---

## Structure hebdomadaire

| Jour | Type | Heures | Contenu |
|------|------|--------|---------|
| Lun | Profond | 6 h | Blocs 1-5 |
| Mar | Profond | 6 h | Blocs 1-5 |
| Mer | Profond | 6 h | Blocs 1-5 |
| Jeu | Profond | 6 h | Blocs 1-5 |
| Ven | Profond | 6 h | Blocs 1-5 |
| Sam | Review | 4 h | Audit recall + 1 push PR OSS + plan semaine suivante |
| Dim | **REPOS** | 0 h | **PAS DE GO.** Pas d'ordinateur si possible. Obligatoire. |

**Total** : ~34 h/sem. Le repos dominical est **non négociable** — c'est l'architecture anti-burnout, pas une récompense.

---

## Rituel quotidien (Bloc 5, 30 min)

1. Remplir `daily-log.md` (1 ligne : date, phase, chapitre, blocs faits, h réelles, test 🟢/🔴, recall dû/fait).
2. Noter **1 chose apprise** + **1 chose encore floue**.
3. Planifier le Bloc 1 de demain (chapitre + fichier prediction à remplir).

---

## Rituel hebdomadaire (Sam, 4 h)

1. Audit recall : passer toutes les cartes dues cette semaine (J+1, J+2, J+7, J+30, J+90).
2. Vérifier `go build ./...` + `go vet ./...` + `go test ./...` globalement verts.
3. 1 push sur une PR OSS en cours (si Phase 2+).
4. Comparer semaine écoulée vs milestone planifié dans `roadmap.md`.
5. Si retard > 3 j : ajuster semaine suivante (**JAMAIS compresser**).
6. Écrire `weekly/week-NN.md` (3 lignes : avancée, blocage, décision).

---

## Rituel mensuel (dernier Sam du mois)

1. Vérifier critères de sortie de la phase dans `roadmap.md`.
2. Si tous ✅ : avancer de phase.
3. Si retard : dropper un objectif secondaire, NE PAS compresser.
4. Décision d'achat du bundle : fin Mois 2 (semaines 8-10).

---

## Boucle d'apprentissage Ardan-style (par chapitre)

0. **Amorcer** (5 min, Bloc 1) : ouvrir la leçon, lire UNIQUEMENT les titres d'exemples + les signatures de fonctions (1re ligne de chaque example). Fermer l'onglet. NE PAS lire les corps ni les sorties. Donne le vocabulaire minimum pour prédire informé, pas à l'aveugle. (Sujet totalement nouveau : tu peux lire aussi le 1er paragraphe d'intro. Sujet déjà vu — comme Functions pour toi : saute l'amorce, tu as déjà la base dans ton code.)
1. **Prédire** (Bloc 1, AVANT lecture) : remplir `predictions/<chap>.md` sans regarder le code. Trois niveaux selon ta base :
   - **Base forte** → sortie exacte + règle énoncée.
   - **Base faible** → la FORME (« compile / pas », « retourne N valeurs ») + une raison.
   - **Zéro base** → « je ne sais pas, mais je dirais X parce que Y ». C'EST une prédiction valide.
   Le but n'est pas d'avoir raison. Une prédiction fausse corrigée vaut 10x une lecture sans prédire (elle détruit le biais de rétrospective « je le savais »).
2. **Lire** (Bloc 1) : le chapitre du tour.
3. **Confronter** (Bloc 1) : ajouter colonnes "Chapitre dit" + "Surprise" dans `predictions/<chap>.md`.
4. **Implémenter** (Bloc 2) : recréer l'example de mémoire dans `<pkg>/exampleN/main.go`.
5. **Prouver** (Bloc 2) : écrire un test table-driven qui passe (preuve de compréhension).
6. **Approfondir** (Bloc 3) : lire le source stdlib pertinent (ex : `io` pour functions).
7. **Recall** (Bloc 4, espacé) : J+1, J+2, J+7, J+30, J+90 — fermé, sans notes.
8. **Logger** (Bloc 5) : 1 ligne dans `daily-log.md` + `LOG.md` (tracker chapitre).

---

## Règles d'or

- Une affirmation sans test est une croyance. Toujours prouver.
- Une intuition non testée n'est pas une connaissance.
- Prédire AVANT lire. Toujours. Même si on se trompe — surtout si on se trompe.
- Le repos dominical n'est pas une option. C'est l'architecture.
- Comprimer pour rattraper = bâcler = perdre la profondeur. Jamais.
- **Cadence par taille, gate par profondeur — pas par compte.** Petits chapitres : 2-3/jour OK. Chapitres moyens : 2/jour. **Keystones (1/jour max, jamais rusher) : interfaces (9 ex), composition (5 sous-articles), generics (9 sous-articles), algorithms (8 sous-articles). Chapitres à 2 exercices (pointers, error-handling) : 1-2/jour.** Avancer nécessite le **portail de profondeur** : prédictions confrontées + examples retapés + `go build`/`go test` verts + exercice tenté + 2-4 cartes recall + tu sais expliquer la règle à froid. Si tu franchis le portail en <1.5 h → avance le même jour (ne padde pas). Si l'exercice te bloque >1 h → stop, log, avance demain (ne grind pas, ne saute pas l'exercice).
- Identifiants en anglais dès maintenant — on vise Sourcegraph, on code en anglais.
