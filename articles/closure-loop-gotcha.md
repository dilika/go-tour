# Le piège que tout développeur Go devrait connaître : la closure dans une boucle

```go
var fns []func()
for i := 0; i < 3; i++ {
    fns = append(fns, func() { fmt.Println(i) })
}
for _, f := range fns {
    f()
}
```

Qu'affiche ce code ?

La majorité des développeurs Go expérimentés répondent sans hésiter : **`3 3 3`**. C'était la réponse correcte pendant plus de dix ans. C'est aussi, aujourd'hui, une réponse fausse.

## La version de Go décide

Sur Go 1.21 et antérieur, ce code affiche :

```
3
3
3
```

Sur Go 1.22 et supérieur, il affiche :

```
0
1
2
```

Même code. Comportement différent. Et ce n'est pas un bug — c'est un changement volontaire de la spécification du langage.

## Pourquoi `3 3 3` pendant dix ans

Pendant toute la première décennie de Go, la variable de boucle `i` était **une seule variable**, déclarée une fois, réutilisée à chaque itération. Toutes les closures créées dans la boucle capturaient **la même référence** vers cette variable unique.

Au moment où la boucle se termine, `i` vaut `3`. Les closures ne sont exécutées qu'après, dans le second `for`. À cet instant, elles lisent toutes la même `i` — qui vaut `3`. D'où `3 3 3`.

Ce piège était si classique qu'il figurait dans quasiment tous les onboarding Go. La parade universelle était explicite :

```go
for i := 0; i < 3; i++ {
    i := i // copie locale, une nouvelle variable par itération
    fns = append(fns, func() { fmt.Println(i) })
}
// → 0 1 2
```

`i := i` — la ligne la plus étrange du langage, répétée par générations de développeurs pour contourner une spécification contestable.

## Ce que Go 1.22 a changé

En 2024, Go 1.22 a modifié la sémantique des boucles `for` : **chaque itération crée une nouvelle variable**. La copie `i := i` n'est plus nécessaire. La closure capture maintenant une `i` distincte à chaque tour, qui conserve sa valeur du moment de l'itération.

`0 1 2` devient le comportement par défaut — et le comportement intuitif que la plupart des débutants attendaient depuis le début.

C'est l'un des rares cas où Go a cassé une compatibilité sémantique au nom de la correction. Les migrations ne sont pas triviales pour les codebases qui dépendaient (volontairement ou non) de l'ancien comportement.

## Comment j'ai découvert que je ne savais pas

Je prédisais `1 2 3`. Ni l'ancien piège (`3 3 3`), ni le nouveau comportement (`0 1 2`).

J'avais une compréhension partielle : je savais que les closures capturent, je savais qu'il y avait une histoire de version, mais je n'avais jamais vérifié. Mon modèle mental était une superposition floue de deux comportements, et je n'avais jamais exécuté le code.

La méthode que j'applique désormais pour apprendre Go : **prédire avant de lire, puis exécuter, puis confronter**.

1. **Prédire** : écrire ce que je crois que le code fait, sans regarder la solution.
2. **Exécuter** : lancer le code réel, observer la sortie.
3. **Confronter** : comparer prédiction et réalité, noter l'écart précisément.

L'étape 1 est la plus importante. Une prédiction exposée ne peut plus se cacher. Elle ne peut pas glisser vers "ah oui, je savais ça" — c'est écrit, c'est daté, c'est faux, et c'est irréfutable.

`1 2 3` écrit noir sur blanc m'a forcé à admettre que je ne comprenais pas, là où `3 3 3` m'aurait laissé croire que je maîtrisais le sujet. Mieux vaut une prédiction fausse précise qu'une intuition vague correcte.

## La leçon plus large

Le piège de la closure en boucle n'est pas anecdotique. Il illustre trois choses qui définissent la maîtrise de Go :

1. **La capture se fait par référence, pas par valeur.** Une closure ne copie pas la variable ; elle référence son emplacement. Tant que vous ne savez pas où vit cette variable — stack ou heap, unique ou par itération — vous ne prédisez pas, vous devinez.

2. **La spécification évolue.** Go 1.22 a changé une sémantique en place. Le code que vous écrivez aujourd'hui ne se comporte pas comme le code que vous avez lu dans un tutoriel de 2019. Relire la spec est un réflexe de senior, pas un aveu de débutant.

3. **Une intuition non testée n'est pas une connaissance.** Le fossé entre "je crois que je sais" et "je sais parce que je l'ai exécuté" est exactement là que se nichent les bugs de production.

## Pour aller plus loin

- Le changement exact : [Go 1.22 release notes — loop variable scoping](https://go.dev/doc/go1.22)
- La proposal originelle : [issue #60078](https://github.com/golang/go/issues/60078)
- Vérifiez vous-même : copiez le snippet en haut de cet article, lancez `go run` sur Go 1.21 puis Go 1.22. Quatre lignes de code, deux décennies d'histoire.

---

Je documente mon apprentissage de Go (Ardan Labs Tour → certification → crédibilité open source) prédiction par prédiction. Pas de copie, pas de survol. Chaque chapitre, je me trompe en public, je corrige, j'ancrage.

Si vous apprenez Go ou voulez devenir sérieux sur le langage, essayez la méthode : prédisez avant d'exécuter. Le premier moment où votre prédiction est fausse est le moment où vous commencez réellement à apprendre.
