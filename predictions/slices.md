# Mes predictons sur le chapitre des arrays( Sans regarder le code )

## Example1: Declare and Length
  - predictons: La déclaration d'une slice se fait de la meme maniere que celle d'un array, a la creation on peut specifier ou pas la length de la sclice
  - Raisons:  C'est juste l'intuition que j'ai sans avoir regarder le code et je serai attentif a la raison en lisant les codes.


## Example2: Reference Types
  - predictons: La references type ici est le fait de declarer une slice comme etant un pointer.
  - Raisons: La reference fait allusion au pointer contrairement au value


## Example3: Appending slices
  - predictons: Les slices ont une methode interne qui permet de d'ajouter des élements a la slice vide ou contenant déja des elements.
  - Raisons: d'ou la signification de append et je le suppose parceque ça n'a pas été le cas avec les arrays.


## Example4: Taking slices of slices
  - Predictions:  Je vois ça comme une matrice de slice
  - Raisons:  Pour ne pas raconter du n'importe quoi je prefere ne rien dire.
  - Surprise:

## Example5: Slices and References
  - Predictions:  L'utilisation des slices avec reference
  - Raisons:  Rien a dire.


## Example6: Strings and slices
  - Predictions:  L'utilsation de string avec les slices.
  - Raisons: Rien a dire.


## Example7: Variadic functions
  - Predictions: Une variadic fonction permet de prendre en parametre un nombre illimité de valeur.
  - Raisons:  Rien a dire.

  
## Example8: Range mechanics
  - Predictions: Range mechanics sont les differentes manieres de parcourir une slices.
  - Raisons: Rien a dire.


## Example9: Efficient Traversals
  - Predictions: C'est une bonne maniere de parcourir une slice.
  - Raisons: Rien a dire.


## Example10: Three index slicing
  - Predictions: Aucune idée.
  - Raisons:  Aucune idée.



# Confrontation:
## Example1: Declare and Length
  - predictons: La déclaration d'une slice se fait de la meme maniere que celle d'un array, a la creation on peut specifier ou pas la length de la sclice
  - Chapitre dit: Il y a quatre manieres de delcarer une slice qui sont:
      - var nil:
      `code 
      var slices []string;
      `
      - With initials values:
      `code 
      var slices := string[]{"A", "B", "C", "C"}
      `
      - empty string:
      `code 
      slices := []string{}
      `
      - make with length with capacity: 
      `code
       slices := make([]string, 4, 5)
      `
      - make with only length:
        `code 
      slices := make([]string, 4)
      `

  - Surprise: Surpris de voir la declaration avec une methode.


## Example2: Reference Types
  - predictons: La references type ici est le fait de declarer une slice comme etant un pointer.
  - Chapitre dit : Une slice possède un tableau sous-jacent contigu avec un pas previsible, la capacité et la length peuvent etre different
  - Surprise: Surpris de voir que le slices n'est pas un tableau c'est une structure a trois mots qui pointent vers tableau


## Example3: Appending slices
  - predictons: Les slices ont une methode interne qui permet de d'ajouter des élements a la slice vide ou contenant déja des elements.
  - Chapitre dit : Le chapitre s'accorde avec ce que j'avais predis.
  - Surprise: Un nouveau tableau est construit si le append est appélé sur un tableau dont la capacité est atteinte.


## Example4: Taking slices of slices
  - Predictions:  Je vois ça comme une matrice de slice
  - Chapitre dit: Une slice de slice permet de creer une vue sur une slice donnée, lorsqu'on fait un append et que la capacité de la sous slice est atteinte alors la slice de slice modifie le tableau partagé par les deux slice, pour eviter cela la methode copy est plus specialisé pour cela
  - Surprise: surpris de la petite nuance concernant le tableau sous-jacent modifier par l'operation append et ce selon la capacité atteinte 

## Example5: Slices and References
  - Predictions:  L'utilisation des slices avec reference
  - Chapitre dit : Les slices sont de types references. Quand tu passes une slice a une autre variable ou que tu le passe en parametre a une fonction c'est seulement le header (3 mots: pointer vers le tableau sous-jacent + capacité + longueur), les deux slices pointent vers le meme tableau sous-jacent. Par conséquent modifier une slice via un index sur l'une des slices modifie la donnée pour toute les slices qui partage ce tableau.
  - Surprise: La modification (s[i] = x) est visible partout, par contre append ne modifie pas forcement les autres slices (ça depend si la capacité permet de rallouer ou non)


## Example6: Strings and slices:
  - Predictions:  L'utilsation de string avec les slices.
  - Chapitre dit : on peut convertir une string en []byte ou []rune.
          - []byte(str) ---> creer une copie des donnée(car []byte est mutable).
          - string([]byte) ---> créer une nouvelle string à partir des bytes.
          Les string sont immuable en GO si on fait for range sur une string on obtient des runes(et non des bytes)
  - Surprise: La conversion []byte(str) fait une copie des donnée. Il n'y a pas de partage de memoire comme avec les slices entre elles.


## Example7: Variadic functions
  - Predictions: Une variadic fonction permet de prendre en parametre un nombre illimité de valeur.
  - Chapitre dit : Une fonction variadic se declare en mettant ...Type en dernier parametre. A l'interieur de la fonction, ce parametre est traité comme une slice de type.
      on peut appeler la fonction de deux maniere:
       - En passant plusieurs argument separé par des virgules.
       - En passant une slice existante en le precedent de ... pour la deballer.
  - Surprise: les parametre variadic sont toujours une slices a l'interieur de la fonction meme si on passse zero argumant

  
## Example8: Range mechanics
  - Predictions: Range mechanics sont les differentes manieres de parcourir une slices.
  - Chapitre dit: L'instruction for range retourne a chaque iteration:
    - l'index
    - Une copie de la valeur de l'elements
    Modifier une valeur ne modifie pas la valeur dans la slice original, par contre elle peut modifie si c'est un pointeur ou une valeur de reference.
  - Surprise: Range fait une copie de la valeur. C'est important a comprendre quand on manipule un struct ou des types volumineux.


## Example9: Efficient Traversals
  - Predictions: C'est une bonne maniere de parcourir une slice.
  - Chapitre dit : Le chapitre compare deux façons de parcourir une slice et explique laquel est la plus efficace selon le context:
    - for i := 0; i < len(s); i++ ---> accès a index.
    - for i, v := range s ----> plus idiomatique
    Quand les elements de la slice sont gros(gros struct), utiliser l'index (s[i]) est souvent plus performant que range, car range copie la valleur a chaque iteration.
  - Surprise: range n'est toujours la solution la plus performante.Il faut parfois privilégier l'accès par index quand on veut eviter des copie unitiles.


## Example10: Three index slicing
  - Predictions: Aucune idée.
  - Chapitre dit: Go permet de fair une slicing a trois index:
     `code
      slice := original(low:hight:max)
     `
     -low: l'index de depart.
     -hight: l'index de fin(exlus)
     -max: limite de la capacité de la nouvelle slice
     Cela permet de controler explicitement la capacité de la sous slice( cap = max - low).
  - Surprise: C'est la solution au probleme soulevé dans l'exemple 4.
    En limitant la capacité avec le troisieme index, on empeche que append sur la sous slice ne n'ecrase les donnée dans le tableau partagé au dela de hight. C'est une technique tres important pour eviter des bug subtils.



