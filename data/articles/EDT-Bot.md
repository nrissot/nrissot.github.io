# EDT Bot FR

*Été 2024 ~ Projet Personnel ~ Python, parsing de fichier .ics*

EDT Bot (<span title="also known as">a.k.a.</span> trainwreck.py) est un Bot discord créé en binome avec Dany Dudiot pour donner un accès facilité aux étudiants de L3 Info à leur emploi du temps

Je suis responsable de la majorité de la partie gestion des emplois du temps, et Dany de la majoritée de la partie bot discord, pour de plus amples détails sur sa partie du projet, vous pouvez aller consulter la <a href="https://danydudiot.github.io/fr/projet/edtbot.html">page du projet sur son site</a>

### Pourquoi ?

Depuis la L2, les emplois du temps sont la partie de mon experience académique qui m'agace le plus, et ce pour deux raisons :

1. **Accés** : Bien que nous soyons sensé disposer d'emplois du temps personnalisé, ceux ci peuvent mettre plusieurs mois avant d'être disponibles, et avant qu'ils le soit, l'utilisateur doit partir a la pêche pour troiver son emploi du temps de filiere à travers des menu successifs pour trouver l'emploi du temps correct parmi plusieurs autres aux noms proches, voir identiques
2. **Exactitude** : Lorsque l'utilisateur arrive enfin a son emploi du temps de filiere, il est alors submergé de tous les cours de tous les groupes de TD/TP de sa promo, voire même les groupes de TD/TP des *autres filières* si il a des cours en commun avec ces autres filieres... (De plus, à l'heure ou je rédige ces lignes, les emplois du temps personnalisés sont certes arrivés, mais sont faux, il y manque 1-2 cours par semaine ...)

ex : En L3, les étudiants d'informatiques sont répartis en 2 filières : INGÉ et MIAGE. Ces deux filieres partagent un certains nombres de cours en communs (Framework Web, Réseaux...) or nous nous sommes aperçus que sur l'emploi du temps des MIAGEs figurait des cours labellés : "TP3 - Réseaux", problème étant que ce cours est un cours destiné aux étudiants de la filière INGÉ, le TP correspondant pour les MIAGEs etant "TP3 Miage - Réseaux".

Au début de l'année, pour consulter son emploi du temps, un étudiant doit :

1. se connecter à l'ENT
2. cliquer sur l'application Emploi du temps dans l'ENT
3. revenir en arriere et cliquer une seconde fois car l'app ne charge pas du premier coups
4. cliquer sur le dossier "filieres" puis "UFR-ST" puis "LICENCE INFORMATIQUE" puis selectionner parmis les 13 possibilitées lequel est le sien en n'oubliant evidemment pas de cliquer sur un autre avant, car sinon il ne charge pas :)
5. lire avec attention chaque cours de la journée pour essayer de trouver les siens

### Quelles solutions :

Heureusement la fac propose quelques solutions pour faciliter la tâches des étudiants, comme par exemple un flux RSS pour s'abonner aux mises a jours de l'emploi du temps, ou encore l'export d'un fichier .ics contenant les cours pour une plage horaire donnée.

Malheureusement, le flux RSS ne fonctionne pas, et les tentatives d'envoi de mails aux différents services de la fac nous laisse a croire que cette fonctionnalitée n'a soit jamais été activée, soit est cassée depuis tellement longtemps que plus personne ne se souvient de son existence ...

Les fichiers .ics (ICalendar) en revanche fonctionnent plutôt bien, bien que le fichier téléchargé ne puisse pas evoluer dans le temps, ma solution en L2 etait de telecharger le fichier pour la semaine à venir le dimanche soir, de l'importer dans Google Calendar, de passer une dizaine de minutes a retirer les cours superflux, et d'esperer que les emplois du temps ne changent pas au cours de la semaine.

Cette organisation etant techniquement fonctionnelle, j'ai alors décidé de suivre la devise de tout informaticien : 

> "Pourquoi passer 10 minutes a faire quelque chose de répétitif quand je peut passer 2h a coder un script qui le fera a ma place".

### Première Intention :

L'Objectif initial du projet Trainwreck (nom legerement sarcastique decrivant l'état des emploi du temps de la fac https://www.dictionary.com/browse/train-wreck) etait de creer un script python simple qui, en CLI, permettrait de nettoyer un fichier .ics de tous les cours qui ne me concernait pas :

```py
python trainwreck.py clean EDT.ics
```

### Fichier ICS

Comme dit ci dessus, les emplois du temps sont exportables au format .ics ou ICalendar, un format de calendrier et d'emploi du temps créé en 1998 par l'Internet Engineering Task Force (IETF) avec le standard RFC 2445. 

Ce format est représenté par un fichier en plain text, composé d'objets multilignes, et d'attribut sur une ligne. Chaque calendrier est un objet multiligne de type `VCALENDAR`, contenant des propriétées et des objets `VEVENT` qui representent chaqu'un des evenements (ici, des cours) du calendrier.

Ces `VEVENT` sont eux même composé d'un certain nombre de parametres, comme les timestamps de début et de fin (`DTSTART` et `DTEND`) en utc-0, une salle `LOCATION`, un `UID` etc

```ics
BEGIN:VCALENDAR
METHOD:REQUEST
PRODID:-//ADE/version 6.0
VERSION:2.0
CALSCALE:GREGORIAN
BEGIN:VEVENT
DTSTAMP:20240930T080214Z
DTSTART:20241127T123000Z
DTEND:20241127T153000Z
SUMMARY:Pro. Imp. Pro. Ori. Obj - TP2
LOCATION:E02
DESCRIPTION:\n\nGr TPB\nPro. Imp. Pro. Ori. Obj\nPro. Imp. Pro. Ori. Obj\
 nPro. Imp. Pro. Ori. Obj\nL3 INFO - INGENIERIE\nL3 INFORMAT-UPEX MINERVE
 \nL3 INFO-UPEX MINERVE - FC\nDURAND-LOSE\n(Exporté le:30/09/2024 10:02)\
 n
UID:ADE60323032342d323032352d32323835342d392d30
CREATED:19700101T000000Z
LAST-MODIFIED:20240930T080214Z
SEQUENCE:2141444862
END:VEVENT
...
END:VCALENDAR
```
extrait d'un des ICalendar téléchargé.

> Remarque intéressante, dans le standard RFC 2445, il est recommandé que les lignes d'un fichier fassent 75 octet de longueur, les lignes plus longues doivent être envoyées a ligne et précédé d'un espace, c'est le cas dans les descriptions des cours. Ce qui fait qu'il est possible qu'un caractere de retour a la ligne `\n` soit coupé en deux par un retour a la ligne, comme c'est le cas ci dessus...

### "Parser" les fichier .ics

Le "Parsing" des fichier ICalendar est plutôt simple, lire le fichier sous la forme d'une liste de lignes :

```py
# On lit tout le fichiers ICS
with open(filename, "r", encoding="utf-8") as f:
    lines = f.readlines()
```

On parcours ensuite le fichier en initialisant un nouveau dictionnaire pour stocker les champs qui nous sont utile pour chaque event
```py
# Balise Ouvrante, crée un nouveau dictionnaire vide.
for line in lines :
    if line.startswith("BEGIN:VEVENT"):
        event = {"DTSTART":"", "DTEND":"", "SUMMARY":"", "LOCATION":"", "DESCRIPTION":"", "UID":""}
```

si la ligne commence par un espace on l'ajoute au champs description, et si il commence par un des mots clé on l'ajoute au dictionnaire 

```py
elif line.startswith(" "):
    event["DESCRIPTION"] += line.removeprefix(" ").removesuffix("\n")
else:
    for prefix in ("DTSTART:", "DTEND:", "SUMMARY:", "LOCATION:", "DESCRIPTION:", "UID:") :
        if line.startswith(prefix):
            event[prefix.removesuffix(":")] = line.removeprefix(prefix).removesuffix("\n")
            break
```

> Remarque : si une autre ligne excede les 75 octets, les octets restant pourrait être incorrectement placé dans description, cela n'est jamais arrivé dans aucun des fichiers, mais c'est une erreur possible, qui est sur la liste des choses qui seront a changer lors de la v2 de ce projet, (cf section "Améliorations")

enfin si on arrive a la fin d'un event, on appelle la fonction qui extrait le les info importantes (le groupe de TD/TP, la matière, l'enseignant·e·s) des champs bruts avant de stocker l'evenement dans un dictionnaire qui les associe a leur UID.

```py
elif line.startswith("END:VEVENT"):
    e = get_event_from_data(
        self.convert_timestamp(event["DTSTART"]),
        self.convert_timestamp(event["DTEND"]),
        event["SUMMARY"],
        event["LOCATION"],
        event["DESCRIPTION"],
        event["UID"],
        flag_exam
    )
    events[e.uid] = e
```

### `get_event_from_data()`:

la méthode `get_event_from_data` qui extrait les info n'est pas trés agréable a lire, mais essentiellement elle split la description par chaque `\n`, puis comme tout les evenements de la fac suivent un semblant de cohérence, essaie d'extraire les informations en cherchant certains caracteres dans un champs précis etc etc.

Cette maniere de proceder est trés peu optimale car si la fac decide de changer de format ou de rajouter des champs, toute cette partie arrete de fonctionner, et il faudrait recoder cette fonction... (*Garbage in, Garbage out*)

Mais il est difficile de trouver des alternatives, une solution alternative basée sur les regex sera probablement utilisée pour la v2, mais pour l'instant, le code bien que dépendant du bon vouloir de la fac de rester cohérent, semble suffisament robuste pour avoir résisté a deux mises a jours, (avec de trés mineures modifications de notre côté) et parser des emplois du temps de filieres autres que celles d'informatiques.

Cette fonction a cependant été concue pour ne pas s'arreter en cas d'echec de parsing, mais plutot de produire un résultat aberrant, qui est facilement remarquable et qui nous permet rapidement de trouver la cause et de regler le probleme, plutot que de masquer des evenement imparsables.

### Filtrage :

J'ai donc implémenté une classe 'abstraite' `Filter` comprenant une seule méthode `filter`

```py
class Filter:
    """Classe de Base pour les filtres."""
    def filter(self, e:Event) -> bool:
        """Prend en argument un événement `e` et retourne `True` si `e` passe le filtre défini par la sous classe."""
        return True
```

De laquelle des implémentations pour les différent type de filtres ont été dérivés.

ex : Filtre par timestamp

```py
class TimeFilter(Filter):
    """Classe pour les filtres temporelle."""
    def __init__(self, date:date, timing:Timing) -> None:
        self.date   = date
        self.timing = timing
        # Timing est un enum contenant les 

    def filter(self, e: Event) -> bool:
        """Permet de savoir si l'Event passe le filtre."""
        match self.timing:
            case Timing.BEFORE:
                return e.end_timestamp.date() <= self.date
            case Timing.AFTER:
                return e.start_timestamp.date() >= self.date
            case Timing.DURING:
                return e.start_timestamp.date() == self.date
```

le filtrage des evenement consiste donc a parcourir la liste des evenement et de leur appliquer chaqu'un des filtres d'une liste. 

```py
def filter_events(events:list[Event], filters:list[Filter]) -> list[Event]:
    """Applique une liste de filtres à la liste d'événements passée en paramètre et retourne une nouvelle liste filtrée."""
    output = events.copy()
    for e in events:
        for f in filters:
            if not f.filter(e):
                output.remove(e)
                break
    return output
```

Cette approche peut paraitre sous optimale, (O(n*m)) mais lors de nos stress test, nous avons découvert qu'au pire des cas, avec un emploi du temps de plusieurs centaines d'evenements (ce qui est plus haut que le maximum théorique pour une année) et une dizaine de filtres, le filtrage complet prend toujours moins d'une demi seconde.

### Récupération des ICalendar

Heureusement pour ce projet, le lien d'export reste stable pour une période donnée, par exemple si on souhaite exporter les evenemements du 1er Septembre au 31 aout, le lien reste le même, meme si le contenu du fichier change.

Il est donc possible de retelecharger le fichier a intervalles régulier grace à une librairie comme urllib ou requests.

### Pourquoi faire simple quand on peut faire compliqué ?

Pendant le développement de ce script, je discutais du processus avec mes amis de la fac, et Dany m'a proposé de convertir le script simple en un bot Discord pour que les autres étudiants puisse profiter aussi de ma solution pour ce problemes.

Nous avons décidé d'utiliser la librairie interactions.py qui est un wrapper de l'API de Discord prenant en charge de maniere relativement simple les <a href="https://discord.com/developers/docs/interactions/overview#commands">nouvelles slash commandes</a>, de l'API discord, contrairement a d'autres librairies comme Discord.py que j'avais pu utiliser par le passé.

Cette partie du developpement a été majoritairement dirigée par Dany, mais j'ai également ajouté ma pierre à l'édifice.

Nous avons décidé de nous concentrer dans un premier temps uniquement sur l'année de L3, etant donné que nous avions la charge du serveur Discord de promo.

Des rôles correspondant aux différents groupes et filieres on été créé sur le serveur, et le bot peut ainsi lire les roles de chacun. Il garde egalement en mémoire une sauvegarde au format pickle des différents groupe de chaque utilisateurs afin de permettre l'utilisation en messages privé (les messages qu'envoient le bot sont de toute maniere configuré pour n'être visible que par la personne concernée.)

Au fur et a mesure du développement, de plus en plus de features on été ajoutées, comme par exemple un systeme d'abonnement pour recevoir l'emploi du temps et un fichier ics propre en début de semaine, et ce de maniere automatique, mais aussi la détection de changements dans la semaine a venir, ou encore un moyen de rentrer manuellement les dates d'examen pour les mettre en valeur sur l'emploi du temps

Pour plus de details sur cette section, vous pouvez consulter la <a href="https://danydudiot.github.io/fr/projet/edtbot.html"> page du projet sur le site de Dany</a>.

### Qu'est ce qui a marché ?

Le Bot discord est un franc succés, nous n'avons encore eu aucune panne ou bug catastrophique depuis la sortie publique du bot. La très grande majorité de la promo l'utilise, et les retours sont très positifs. Pour le cas d'utilisation qui est le notre, le bot fonctionne parfaitement, et sans accroc.

**MAIS**

### Qu'est-ce qui ne va pas ?

Ce projet etant de base un simple prototype en python a explosé en taille et en complexité sous un scope-creep de plus en plus important, pour arriver au plus gros projets sur lequel Dany et moi avions travaillé, et notre manque d'experience quand a la préparation de la croissance d'un tel projet nous a fait cruellement défaut.

De nombreuses parties interne du projets sont brouillonnes et parfois fragile, le principe ouvert-fermé que nous avions essayé d'implementer a tout simplement disparu, avec le script original séparé en petit morceau disséminés et intrinséquement mélangé avec la partie Bot du projets.

l'Ajout de nouvelle feature est un processus long et difficile car la complexité et les ramifications du code sont importantes.

La maintenance, bien que relativement faible en volume peut impliquer de réecrire completement certaines parties du code...

Nous avons essayé d'etendre le principe a la promo des L2, mais de nombreux mécanisme etait trop dépendant de la configuration actuelle du serveur des L3, et adapter le bot pour les L2 reviendrait a réecrire une large partie du bot, et de faire tourner un deuxieme instance en parrallele...


### Futur 

Comme mentionné si dessus une version 2.0 d'EDT bot est prévue, idéalement développée l'été prochain, avec un ou deux collaborateurs supplémentaires, pour réecrire complemetement le bot de 0, mais en visant le produit fini actuel.

Nous mettrons alors l'accent sur l'application de design pattern, et d'un vrai principe ouvert fermé, pour pouvoir garder le noyeau de traitement des emploi du temps séparé de la sortie sur discord, mais aussi pouvoir permettre de facilement supporter les autres filieres, avec idéalement seulement la méthode "get_event_from_data" a réecrire pour chaque nouvelle filiere.

Nous aimerions aussi inclure une vraie base de donnée, plutot qu'un objet python et une sauvegarde pickle de cet objet, et eventuellement une interface web.

