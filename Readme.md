Aufgabe 1 

Wir haben beide Funktionen in einer "for-loop" aufgerufen und diese 1 Miliarde mal durchlaufen. Da kamen wir zu dem Entschluss aufgrund der Ausgabe das die "Run-time method" signifikant schneller ist. Nach Analyse des Codes kamen wir zu der Vermutung das die unterschielichen Ausfuehrungszeiten darauf zurückzuführen sind das in der "sumArea_Dict"-Funktion eine Struktur(shape_Value) verwendet wird die einen Funktionszeiger für die Flächenberechnung enthält. Da der Funktionszeiger, der in der Struktur gespeichert ist, aufgerufen werden muss, was im Vergleich zu einem direkten Funktionsaufruf zusätzlichen Overhead verursacht.

Ausgabe:

Dauer der Funktion Lookup:  518.9µs
Dauer der Funktion Dict:  4.9904ms