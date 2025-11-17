# Aufgaben

Die folgende Aufgaben kannst du direkt in die angegebene Datei lösen. Beachte hierzu die `// TODO: `-Kommentare im Code und die folgenden Instruktionen:

## Steckbrief II

`go-2-ex-1/main.go`: Im letzten Block hast du einen Steckbrief mithilfe loser Variablen definiert (`firstName`, `lastName`, `dayOfBirth` usw.). In dieser Übung sollst du diese Informationen zu `struct`s gruppieren. Erstelle die notwendigen Datenstrukturen und ergänze die Variable `me` mit deinen persönlichen Steckbriefinformationen.

Passen anschliessend die Anzahl deiner Geschwister an, indem du diese um eins erhöhst.

## Fibonacci-Zahlen

`go-2-ex-2/main.go`: Die Fibonacci-Reihe beginnt mit der Folge `1, 1`. Jede weitere Zahl der Reihe ist definiert durch die Summe des Vorgängers und des Vor-Vorgängers, also `1, 1, 2, 3`, weil `1+1=2` und `1+2=3` ist, usw. Geht man von einem 0-basierenden Index aus, kann ein Slice, welches die Fibonacci-Folge enthält, folgendermassen definiert werden:

- Index `0`: Wert `1`
- Index `1`: Wert `1`
- Index `n`: Wert an Index `n-1` plus Wert an Index `n-2`

In der gegebenen Datei wird ein Slice von fünf Elementen definiert. Die ersten beiden Werte lauten `1`. Erweitere das Fibonacci-Slice `fibs` folgendermassen:

1. Indem du die verbleibenden drei Werte berechnest und direkt ins bestehende Slice schreibst.
2. Indem du noch drei weitere Werte berechnest und diese dem Slice anhängst.

## Modulbezeichnungen

`go-2-ex-3/main.go`: Die Module haben jeweils eine Nummer (z.B. `346`) und eine Bezeichnung (z.B. "Cloud-Lösungen konzipieren und realisieren"). Erstelle eine `map` namens `modules`, welche alle nötigen Angaben enthält, um die drei `fmt.Println()`-Aufrufe unten am Programm die richtigen Informationen ausgeben zu lassen. Welche Datentypen musst du hierzu verwenden?

Anschliessend nimmst du folgende Manipulationen an der `map` vor:

1. Ein Element anhand des Schlüssels entfernen.
2. Ein Element hinzufügen.
3. Ein Element durch ein anderes ersetzen.

Du kannst dir Module ausdenken oder im [Modulbaukasten](https://www.modulbaukasten.ch/) nachschauen, wenn dir die Modulinformationen nicht geläufig sind.

## Klassenverwaltung

`go-2-ex-4/main.go`: Erstelle eine kleine Klassenverwaltung, indem du mithilfe von Strukturen, Maps und Slices folgende Sachverhalte modellierst:

- Ein Schüler (`Student`) hat einen Vor- und Nachnamen.
- Eine Klasse (`Class`) besteht aus einer Reihe von Schülern.
- Ein Modul hat eine eindeutige Nummer (z.B. `346`) und wird von einer Reihe von Klassen besucht.

Erstelle die notwendigen Datenstrukturen mit entsprechenden Beispieldaten (d.h. mindestens zwei Klassen mit je drei Schülern und mindestens drei Module, die von einer oder mehreren Klassen besucht werden). Gib die Daten anschliessend per `fmt.Println()` auf die Konsole aus.
