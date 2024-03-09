# Java JSONPlaceholder Comments Fetcher and Database Inserter

Questo progetto Java si propone di recuperare i commenti dall'API JSONPlaceholder (https://jsonplaceholder.typicode.com/comments), iterando attraverso i postId e inserendo i commenti recuperati in un database MySQL.

## Requisiti

- Java 8 o versioni successive
- Una connessione internet per accedere all'API JSONPlaceholder
- Un database MySQL in cui memorizzare i dati

## Utilizzo

1. Assicurati di avere un database MySQL configurato correttamente.
2. Modifica le informazioni di connessione al database nel file `Main.java` nel metodo `main`.
3. Esegui il file `Main.java`.
4. Il programma recupererà i commenti dall'API JSONPlaceholder, li inserirà nel database e stamperà un messaggio di conferma.

## Struttura del Progetto

- `Main.java`: il file principale contenente il codice Java per recuperare i commenti e inserirli nel database.
- `README.md`: questo file.

## Dipendenze

- Java Standard Library
- Libreria `org.json` per il parsing dei dati JSON
- Connessione Internet per accedere all'API JSONPlaceholder

## Contribuisci

Senti libero di migliorare questo progetto aprendo issue e pull requests!
