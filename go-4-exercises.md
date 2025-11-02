# Aufgaben

Die folgende Aufgaben kannst du direkt in die angegebene Datei lösen. Beachte hierzu die `// TODO: `-Kommentare im Code und die folgenden Instruktionen:

## Note berechnen

`go-4-ex-1/main.go`: Eine Note `N` von 1 (schlechteste) bis 6 (beste) kann linear anhand der erreichten Punktzahl `E` und maximalen Punktezahl `M` berechnet werden:

![Berechnung der Note](formulas/grade.png)

Schreibe eine Funktion namens `computeGrade`, welche die beiden Parameter `gotPoints` (erreichte Punktzahl) und `maxPoints` (maximale Punktzahl) erwartet und eine Note im Bereich `[1.0;6.0]` zurückgibt.

Rufen die Funktion dreimal mit verschiedenen Werten auf und schreibe das Ergebnis als Kommentar hinter den jeweiligen Funktionsaufruf, z.B.:

    computeGrade(17.5, 28.0) // 4.125

### Zusatzaufgabe

Erweitere die Funktion `computeGrade`, sodass sie auf sinnlose Argumente reagiert (z.B.  negative Punktzahlen, mehr Punkte erreicht als möglich) und einen entsprechenden `error` zurückgibt. Passe deine Funktionsaufrufe entsprechend an, um auf Fehler reagieren zu können.

## Hypotenuse berechnen

`go-4-ex-2/main.go`: Der Satz vom Pythagoras ist definiert als:

![Satz vom Pythagoras](formulas/pythagoras.png)

Stellt man die Formel auf `c` um, kann man mit der Formel die Länge der Hypotenuse bei gegebenen Kathetenlängen berechnen:

![Berechnung der Hypotenuse](formulas/hypotenuse.png)

Schreibe eine Funktion namens `computeHypotenuse`, welche die beiden Kathetenlängen `a` und `b` als Parameter erwartet und die Länge der Hypotenuse zurückgibt.

Tipp: Verwende die Funktion [`math.Pow`](https://pkg.go.dev/math#Pow) und [`math.Sqrt`](https://pkg.go.dev/math#Sqrt) für die Berechnung.

Teste die Funktion mit verschiedenen Kathetenlängen und dokumentiere die Aufrufe wie bei Aufgabe 1 mit Kommentaren.

### Zusatzaufgabe

Die beiden Kathetenlängen `a` und `b` können zu einem neuen Typ namens `ShortSides` gruppiert werden:

```go
type ShortSides struct {
    a float64
    b float64
}
```

Schreiben eine _Methode_ `Hypotenuse`, welche auf `ShortSides` operiert und keine weiteren Parameter erwartet.

Testen die Methode mit den gleichen Kathetenpaaren wie `computeHypotenuse` und notiere dir die Ergebnisse als Kommentare hinter den Methodenaufrufen.

## Quadratische Gleichungen lösen

`go-4-ex-3/main.go`: Quadratische Gleichungen der Form `ax² + bx + c = 0` lassen sich mithilfe der sogenannten _Mitternachtsformel_ lösen:

![Lösung(en) von quadratischen Gleichungen berechnen](formulas/quadratic-formula.png)

Den Teil unter der Wurzel bezeichnet man als _Diskriminante_:

![Berechnung der Diskriminanten](formulas/discriminant.png)

Je nach Wert der Diskriminante `D` gibt es eine unterschiedliche Anzahl an Lösungen:

- `D > 0`: zwei Lösungen
- `D = 0`: eine Lösung
- `D < 0`: keine Lösung

Schreibe eine Funktion namens `computeQuadraticFormula`, welche die Parameter `a`, `b` und `c` erwartet, und ein Slice mit den berechneten Lösungen zurückgibt.

Tipp: Verwende die Funktion [`math.Pow`](https://pkg.go.dev/math#Pow) und [`math.Sqrt`](https://pkg.go.dev/math#Sqrt) für die Berechnung.

Rufe die Funktion dreimal auf, sodass je ein Ergebnis mit zwei Lösungen, mit einer Lösung und ohne Lösung entsteht. Dokumentiere die Ergebnisse wiederum als Kommentare. Du kannst hierzu die folgenden Testdaten verwenden:

| D | a | b | c |
|--:|--:|--:|--:|
| + | 3 | 4 | 1 |
| 0 | 2 | 4 | 2 |
| - | 3 | 4 | 2 |

### Zusatzaufgabe

Lagere die Berechnung der Diskriminanten in eine eigene Funktion namens `computeDiscriminant` aus. Teste die Funktion separat mit den Testdaten von vorher und dokumentiere die Ergebnisse als Kommentar.

Schreibe anschliessend die Funktion `computeQuadraticFormula` so um, dass sie zur Berechnung die Funktion `computeDiscriminant` verwendet:

![Lösung(en) von quadratischen Gleichungen mithilfe der Diskriminanten berechnen](formulas/quadratic-formula-discriminant.png)

## Temperaturumrechnungen durchführen

`go-4-ex-4/main.go`: Schreiben zwei Funktionen `convertCelsiusToFahrenheit` und `convertFahrenheitToCelsius`, welche Temperaturangaben von Celsius (`C`) nach Fahrenheit (`F`) bzw. von Fahrenheit nach Celsius umrechnen. Verwende hierzu folgende Formeln:

![Fahrenheit in Celsius umrechnen](formulas/fahrenheit-to-celsius.png)

![Celsius in Fahrenheit umrechnen](formulas/celsius-to-fahrenheit.png)

Rufe `convertCelsiusToFahrenheit` mit drei verschiedenen Werten auf und dokumentiere die Ergebnisse als Kommentar.

Rufe anschliessend `convertFahrenheitToCelsius` mit den Ergebnissen von vorher auf, sodass du wieder die ursprüngliche Angabe in Celsius erhalten solltest.

### Zusatzaufgabe

Gegeben sind folgende Typen:

```go
type Celsius float64
type Fahrenheit float64
```

Schreibe zwei Methoden `ConvertToFahrenheit` und `ConvertToCelsius` für `Celsius` bzw. `Fahrenheit`, sodass folgende Aufrufe möglich werden:

```go
var cozy Celsius = 23.0
cozy.ConvertToFahrenheit()

var cold Fahrenheit = 15.3
cold.ConvertToCelsius()
```

Teste die Methoden mit den Daten von vorher und dokumentiere die Ergebnisse als Kommentare.

Zusatzfrage: Welche Notation findest du übersichtlicher?

```go
fmt.Println(convertFahrenheitToCelsius(convertCelsiusToFahrenheit(23.0)))
```

Oder:

```go
var c Celsius = 23
fmt.Println(c.ConvertToFahrenheit().ConvertToCelsius())
```
