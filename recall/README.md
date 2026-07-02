# recall/

Rappels espacés. C'est ici que la connaissance se consolide — c'est la **preuve** que tu sais. Le reste est croyance.

## Protocole de répétition espacée

Pour chaque carte, rappeler aux intervalles (fermé, sans notes, sans code) :

- **J+1** : fermé. Si fail → retour J+0 (relire + refaire la carte).
- **J+2** : fermé. Si fail → retour J+1.
- **J+7** : fermé. Si fail → retour J+1.
- **J+30** : fermé. Si fail → retour J+7.
- **J+90** : fermé. Si fail → retour J+30.

## Format d'une carte

Fichier : `recall/<topic>-NN.md` (ex : `recall/functions-01.md`)

```
## Question
<question courte, ex : "Quand une valeur escape-t-elle vers le heap ?">

## Réponse (rappel sans notes)
<écrire de mémoire, PUIS vérifier — ne triche pas>

## Vérification
<ce que dit le source / le chapitre — 2-3 lignes>

## Historique
- J+1 (date) : ✅/❌
- J+2 (date) : ✅/❌
- J+7 (date) : ✅/❌
- J+30 (date) : ✅/❌
- J+90 (date) : ✅/❌
```

## Règle

Si tu bloques à un recall → tu n'avais pas compris → retour au test + recréer la carte.
Le recall est la **preuve** que tu sais. Le reste est croyance.

## Quand créer une carte

À la fin de chaque chapitre (Bloc 2 ou 4), créer 2-4 cartes :
- une sur la règle exacte (ex : `:=` redeclaration)
- une sur un piège (ex : loop var en Go 1.22+)
- une sur un lien stdlib (ex : `io.Writer` vs `Write`)

Cible : **100 cartes fin Phase 1**, 200 fin Phase 2, 300 fin Phase 3.
