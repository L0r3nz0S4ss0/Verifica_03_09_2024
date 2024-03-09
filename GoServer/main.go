package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Connessione al database MySQL
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/esercizio")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inizializzazione del router gin
	router := gin.Default()

	// Endpoint per ottenere i commenti in base alla data di aggiunta
	router.GET("/comments/api/get/date/:date", getCommentsByDate)

	// Endpoint per ottenere i commenti in base all'ora di aggiunta
	router.GET("/comments/api/get/time/:time", getCommentsByTime)

	// Endpoint per aggiornare un commento
	router.PATCH("/comments/api/update/:id", updateComment)

	// Endpoint per eliminare un commento
	router.DELETE("/comments/api/delete/:id", deleteComment)

	// Endpoint per creare un nuovo commento
	router.POST("/comments/api/create", createComment)

	// Avvio del server
	router.Run(":8080")
}

func getCommentsByDate(c *gin.Context) {
	date := c.Param("date")
	rows, err := db.Query("SELECT * FROM comments WHERE DATE(data_ora) =?", date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	comments := []map[string]interface{}{}
	for rows.Next() {
		var postId, id, name, email, body string
		var dataOraStr string
		if err := rows.Scan(&postId, &id, &name, &email, &body, &dataOraStr); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Converti la stringa dataOraStr in time.Time
		dataOra, err := time.Parse("2006-01-02 15:04:05", dataOraStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		comment := map[string]interface{}{
			"postId":   postId,
			"id":       id,
			"name":     name,
			"email":    email,
			"body":     body,
			"data_ora": dataOra,
		}
		comments = append(comments, comment)
	}

	c.JSON(http.StatusOK, comments)
}

// ...

func getCommentsByTime(c *gin.Context) {
	timeStr := c.Param("time")
	timeInt, err := strconv.Atoi(timeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format"})
		return
	}

	hour := timeInt / 100
	minute := timeInt % 100

	timeObj := time.Date(0, 1, 1, hour, minute, 0, 0, time.UTC)

	rows, err := db.Query("SELECT * FROM comments WHERE TIME(data_ora) = ?", timeObj.Format("15:04:05"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	comments := []map[string]interface{}{}
	for rows.Next() {
		var postId, id, name, email, body string
		var dataOraStr string
		if err := rows.Scan(&postId, &id, &name, &email, &body, &dataOraStr); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Converti la stringa dataOraStr in time.Time
		dataOra, err := time.Parse("2006-01-02 15:04:05", dataOraStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		comment := map[string]interface{}{
			"postId":   postId,
			"id":       id,
			"name":     name,
			"email":    email,
			"body":     body,
			"data_ora": dataOra,
		}
		comments = append(comments, comment)
	}

	c.JSON(http.StatusOK, comments)
}

func updateComment(c *gin.Context) {
	id := c.Param("id")

	// Parsing dei dati JSON dal corpo della richiesta
	var requestBody map[string]interface{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Aggiornamento del commento nel database
	_, err := db.Exec("UPDATE comments SET name = ?, email = ?, body = ? WHERE id = ?", requestBody["name"], requestBody["email"], requestBody["body"], id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Commento con ID %s aggiornato con successo", id)})
}

func deleteComment(c *gin.Context) {
	id := c.Param("id")

	// Eliminazione del commento nel database
	_, err := db.Exec("DELETE FROM comments WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Commento con ID %s eliminato con successo", id)})
}

func createComment(c *gin.Context) {
	// Parsing dei dati JSON dal corpo della richiesta
	var requestBody map[string]interface{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Inserimento del nuovo commento nel database
	_, err := db.Exec("INSERT INTO comments (postId, id, name, email, body, data_ora) VALUES (?, ?, ?, ?, ?, ?)", requestBody["postId"], requestBody["id"], requestBody["name"], requestBody["email"], requestBody["body"], time.Now())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Commento creato con successo"})
}
