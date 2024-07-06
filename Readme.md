## Aufgabe 1 

Wir haben beide Funktionen in einer "for-loop" aufgerufen und diese 1 Miliarde mal durchlaufen. Da kamen wir zu dem Entschluss aufgrund der Ausgabe das die "Run-time method" signifikant schneller ist. Nach Analyse des Codes kamen wir zu der Vermutung das die unterschielichen Ausfuehrungszeiten darauf zurückzuführen sind das in der "sumArea_Dict"-Funktion eine Struktur(shape_Value) verwendet wird die einen Funktionszeiger für die Flächenberechnung enthält. Da der Funktionszeiger, der in der Struktur gespeichert ist, aufgerufen werden muss, was im Vergleich zu einem direkten Funktionsaufruf zusätzlichen Overhead verursacht.


Ausgabe:

Dauer der Funktion Lookup:  249.599ms

Dauer der Funktion Dict:  4.8561749s









## Aufgabe 2
Bei der zweiten Aufgabe sollten wir uns ein eigenes Beispiel überlegen, auf das wir den RT und DT Ansatz anwenden. Hierfür haben wir uns für ein Szenario entschieden, bei dem der Lohn von verschiedenen Positionen innerhalb eines Unternehmens berechnet wird. 

Jede Struktur (CEO, Developer, Janitor) hat eine eigene Implementierung der 'calculatePay'-Methode, die den Stundenlohn auf Basis der gearbeiteten Stunden und des spezifischen Stundensatzes berechnet. 

Am Ende wird zufällig ein Mitarbeiter ausgewählt, dessen Gehalt berechnet wird. 




## Aufgabe 3
Der die Möglichkeit der TypeAssertions wird eine Möglichkeit gegeben den konkreten Typ eines Interface-Wertes zur Laufzeit zu überprüfen und zu verwenden.  Der gegebene Code wurde so angepasst, das nun in den Funktionen sumArea_Dict und sumAreaVariant Typbestimmungen druchgeführt werden, um die Flächen von den verschiedenen geometrischen Formeln zu berechnen. Damit kann überprüft werden, ob es sich um einen bekannten Typen handelt, der verarbeitet werden kann. Sollte der Typ nicht bekannt sein tritt ein Fehler auf.


## Aufgabe 4
In dieser Aufgabe werden TypeBounds implementiert, um generische Datentypen mit bestimmten Methoden zu definieren. Dies ermöglicht es FUnktionen und Datentypen flexibler und gleichzeitig sicherer zu gestalten. 
Die generische Struktur 'node' verwendet einen Typparameter 'T', der durch 'any' repräsentiert wird. Dies bedeuetet, dass 'node' jede Art von Wert speichern kann. Die Funktion 'showNode' verwendet einen Typparameter 'T', der durch die Schnittstelle 'Show' beschränkt ist. Dies bedeutet, dass 'showNode' nur mit Typen funktioniert, die die 'Show' Schnittstelle implementieren. Es wurden noch einige Testfälle implementiert, die den Einsatz von TypeBounds zeigen sollen. 