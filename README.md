# Web Crawler

In diesem Repo wird ein einfacher Web-Crawler entwickelt.

Es werden Datenstrukturen und Funktionen definiert, um HTML-Seiten zu repräsentieren
und auszuwerten, welche Links es darin gibt.
Die so definierten Objekte bilden einen Graphen, den man daraufhin untersuchen kann,
welche Seiten von einer Startseite aus erreichbar sind, wie viele Links auf eine Seite
zeigen etc., um dann auf dieser Grundlage eine einfache Suchmaschine zu bauen.

## Nächste Schritte

Der HTML-Parser in diesem Package ist noch in Entwicklung.
Hier sind die nächsten TODOs:

* Testdaten konsolidieren
  * Konstanten aus Quellcode entfernen, dafür HTML-Dateien einbetten.
  * Bisherige Beispiele vereinfachen und passender benennen.
  * Weitere Beispiele mit Text im Body machen.
* Baum-Traversierung einbauen, um den HTML-Baum sichtbar zu machen.
* Texte im HTML-Body erfassbar machen.
* URLs aus dem echten Netz öffnen.
