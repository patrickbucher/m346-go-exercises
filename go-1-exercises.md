# Aufgaben

Die folgende Aufgaben kannst du in die angegebene Datei lösen. Beachte hierzu die `// TODO: `-Kommentare im Code und die folgenden Instruktionen:

## Steckbrief

`go-1-ex-1/main.go`: Es werden verschiedene Variablen ausgegeben, die jedoch noch nicht deklariert und initialisiert worden sind. Deklariere die Variablen mit einem passenden Datentyp und verwende deine persönlichen Angaben um die Variablen zu initialisieren.

Tipp: Das Sternzeichen kannst du mit der Unicode-Schreibweise `'\uXXXX'` eingeben, wobei du `XXXX` durch die vier letzten Zeichen des Werts in der [Spalte "Unicode"](https://en.wikipedia.org/wiki/Astrological_symbols#Signs_of_the_zodiac) ersetzen kannst.

## Umrechnungen

`go-1-ex-2/main.go`: Es werden verschiedene imperiale Einheiten (Meilen, Grad Fahrenheit) in die entsprechenden metrischen Einheiten umgerechnet. Du musst im unteren Teil des Programms zwei Berechnungen in die andere Richtung ausführen, d.h. vom metrischen ins imperiale System.

Falls du die Aufgabe nicht lösen kannst, notierst du dir die Umrechnungen auf ein Blatt Papier und stellst die Formeln entsprechend um.

## Standardausgabe, Standardfehler

**Hinweis**: Dieses Programm muss mit der Git Bash oder einer vergleichbaren Unix-Shell ausgeführt werden (`csh`, `ksh`, `zsh`), nicht mit der PowerShell!

`go-1-ex-3/main.go`: Das Programm würfelt eine zufällige Zahl (zwischen 1 und 6) und zeigt diese an. Weiter zeigt das Programm an, wann die Zahl gewürfelt worden ist. Anstelle von `fmt.Println` verwendest du die Funktion `fmt.Fprintln` um die beiden Informationen auszugeben.

Als erstes Argument musst du hierzu `os.Stdout` oder `os.Stderr` verwenden.  Für welche Art der Information (erwürfelte Augen oder Zeitpunkt des Würfelns) ist welche Ausgabe (`os.Stdout`/`os.Stderr`) besser geeignet?

Rufe anschliessend das Programm so auf, dass die Augenzahl in die Datei `eyes.txt` geschrieben wird, und der Zeitpunkt in die Datei `dice.log`. Ergänze den Kommentar unten am Programm um diesen Aufruf entsprechend zu dokumentieren.

**Wichtig**: Implementiere hierzu keinen Dateizugriff! Die Dateinamen dürfen im Programm nur als Kommentare auftauchen. Der Aufrufer entscheidet, wo welche Ausgabe landet!

## Unterlagen

Durchstöbere die folgenden Webseiten:

- [go.dev](https://go.dev/): Die offizielle Webseite von Go
- [go.dev/doc](https://go.dev/doc/): Die offizielle Dokumentation zu Go
- [pkg.go.dev](https://pkg.go.dev/): Suche nach Packages
- [The Go Playground](https://go.dev/play/): Interaktive Online-Entwicklungsumgebung für kleine Beispiele
- [Go by Example](https://gobyexample.com/): Beispiele zu verschiedenen Bereichen von Go

Gib in der Kommandozeile die folgenden Befehle ein:

- `go doc`
- `go doc fmt`
- `go doc fmt.Println`

Mache dich mit dieser Form der Dokumentation vertraut und betrachte ein paar weitere Funktionen aus dem `fmt`-Package.

Du kannst mit den Pfeiltasten durch die Ausgabe navigieren. Mit `q` kannst du die Anzeige verlassen.
