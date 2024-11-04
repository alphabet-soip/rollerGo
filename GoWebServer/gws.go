package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func helloWorldHTML(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type to HTML
	w.Header().Set("Content-Type", "")
	// Write the HTML response
	fmt.Fprintf(w, "<h1>Hello World! - GWS</h1>")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	// Text response
	fmt.Fprintf(w, "Hello World! - GWS")
}

func helloWorldJSON(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello World! - GWS"}
	json.NewEncoder(w).Encode(response)
}

// Template for the HTML page
var tmpl = template.Must(template.New("dice").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dice Roller</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            text-align: center;
        }
        .dice-container {
            margin: 20px;
        }
        .dice-value {
            font-size: 2em;
            text-align: right;
        }
        .dice-label {
            text-align: right;
            display: inline-block;
            width: 150px;
        }
        .readonly-field {
            text-align: right;
            background-color: lightgray;
        }
    </style>
</head>
<body onload="rollDice()">
    <h1>Dice Roller</h1>
	<h2>Production</h2>
    <div class="dice-container">
        <label for="dice1" class="dice-label">Dice:</label>
        <input id="dice1" type="text" class="readonly-field" readonly><br><br>

        <button id="rollButton" autofocus onclick="rollDice()">Roll Dice</button>
    </div>

    <footer>
		<p>Press enter to roll dice.</p>
		<p>If lost focus, click 'Roll Dice'.</p>
        <p>Environment: Roller Production</p>
        <p>Version: 1.01</p>
    </footer>

    <script>
        function rollDice() {
            fetch('/roll-dice')
                .then(response => response.text())
                .then(dice1 => {
                    document.getElementById('dice1').value = dice1;
                    document.getElementById('rollButton').focus();
                })
                .catch(error => console.error('Error:', error));
        }
    </script>
</body>
</html>
`))

// Handler for serving the HTML page
func serveTemplate(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Handler for rolling the dice
func rollDiceHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	diceRoll := rand.Intn(6) + 1
	w.Write([]byte(strconv.Itoa(diceRoll)))
}

func main() {
	// Route for the root path
	http.HandleFunc("/hello-world-html", helloWorldHTML)
	http.HandleFunc("/hello-world", helloWorld)
	http.HandleFunc("/hello-world-json", helloWorldJSON)
	http.HandleFunc("/", serveTemplate)
	http.HandleFunc("/roll-dice", rollDiceHandler)

	// Start the server on port 8080
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
