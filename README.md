# Web Crawler

In diesem Repo wird ein einfacher Web-Crawler entwickelt.

## Vorgegebener Code

Es sind Datenstrukturen und Funktionen definiert, um HTML-Seiten zu repräsentieren
auszuwerten, welche Links es darin gibt, den Seitentext zu bestimmen und eine Schlagwortsuche
in diesem Text durchzuführen.

Die Benutzung dieser Infrastruktur wird in den Tests in `htmlparser` sowie in der Datei `main.go`
demonstiert.

## Aufgaben

Entwickeln Sie auf Basis dieser Datenstrukturen einen einfachen Webcrawler.
Dieser soll folgende Aufgaben erledigen:

* Einen Link entgegennehmen, die Seite parsen und Links auf dieser Seite sammeln.
* Jeden dieser Links verfolgen und erneut die Seite parsen und Links sammeln.
* Dies soll immer weiter fortgesetzt werden, so dass eine Liste aller Links entsteht,
  die von der ursprünglichen Startseite aus erreichbar sind.

**Hinweise:**

* Sie sollten ein Abbruchkriterium für die Suche haben, damit es nicht endlos weitergeht.
  Das Web ist sehr groß und enthält auch Kreise in den Links. Diese müssen Sie vermeiden oder aufbrechen.
* Sie müssen nicht unbedingt alle Links auf der jeweiligen Seite beachten,
  es können auch Einschränkungen definiert werden.
  Vgl. auch das Beispiel aus der `main.go`: Dort werden die gesammelten Links auf die Domain
  der Start-Url eingeschränkt. Eine Suche mit diesen Filtern würde die Start-Domain also
  nicht verlassen.
* Sie können z.B. auch eine Maximalzahl an Links festlegen, die verfolgt werden, bevor das Programm abbricht.
* Alternativ kann auch eine maximale Suchtiefe festgelegt werden.
