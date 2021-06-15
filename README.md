# gen-pfi

gen-pfi is a small program to generate an english excercise.

give it a list of en wikipedia urls, and it builds a markdown file for you

```sh
cat urls.txt | pfigen > pfi.md
```

## Example

### Input

```txt
https://en.wikipedia.org/wiki/Agile_software_development Agile software development is a useful practice to get more productive
```

### Output

```md
# PFI

## Agile software development

In software development, agile (sometimes written Agile) practices involve discovering requirements and developing solutions through the collaborative effort of self-organizing and cross-functional teams and their customer(s)/end user(s). It advocates adaptive planning, evolutionary development, early delivery, and continual improvement, and it encourages flexible responses to change.
It was popularized by the Manifesto for Agile Software Development. The values and principles espoused in this manifesto were derived from and underpin a broad range of software development frameworks, including Scrum and Kanban.
While there is much anecdotal evidence that adopting agile practices and values improves the agility of software professionals, teams and organizations, the empirical evidence is mixed and hard to find.

### Méthode agile

Vous pouvez améliorer la vérifiabilité en associant ces informations à des références à l'aide d'appels de notes.
En ingénierie logicielle, les pratiques agiles mettent en avant la collaboration entre des équipes auto-organisées et pluridisciplinaires et leurs clients. Elles s'appuient sur l'utilisation d'un cadre méthodologique léger mais suffisant centré sur l'humain et la communication. Elles préconisent une planification adaptative, un développement évolutif, une livraison précoce et une amélioration continue, et elles encouragent des réponses flexibles au changement,.
Cette approche a été popularisée à partir de 2001 par le Manifeste pour le développement agile de logiciels. Les quatre valeurs et les douze principes adoptés dans ce manifeste sont issus d'un large éventail de méthodes dont Scrum et eXtreme Programming ,. Depuis lors, les méthodes qui s'inscrivent dans la philosophie de ce manifeste sont appelées « méthodes agiles ».
Les méthodes agiles se veulent plus pragmatiques que les méthodes traditionnelles, impliquent au maximum le demandeur (client) et permettent une grande réactivité à ses demandes. Elles reposent sur un cycle de développement itératif, incrémental et adaptatif.

### Notes

Agile software development is a useful practice to get more productive

### Sources

- [https://en.wikipedia.org/wiki/Agile_software_development](https://en.wikipedia.org/wiki/Agile_software_development)
- [https://fr.wikipedia.org/wiki/M%C3%A9thode_agile](https://fr.wikipedia.org/wiki/M%C3%A9thode_agile)
```
