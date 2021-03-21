package main

import (
    "encoding/csv"
    "flag"
    "fmt"
    "io"
    "log"
    "os"
    "time"
    //"reflect"
)

func main() {

    filename := flag.String("csv", "problems.csv", "csv file with questions and answers")
    timeLimit := flag.Int("limit", 30, "the time limit for the quiz")
    flag.Parse()

    // Keep track of correct answers
    var total int = 0
    var correct int = 0
    var wrong int = 0

    // structure to hold questions and answers
    type question struct {
        q, a interface{}
    }

    // a slice to hold the questions
    var s []question

    // Open the csv file
    csvfile, err := os.Open(*filename)
	  if err != nil {
		    log.Fatalln("Couldn't open the csv file", err)
	  }

    // Parse the file
	  r := csv.NewReader(csvfile)

    // Iterate through the records and save them as questions in s.
	  for {
		    // Read each record from csv
		    record, err := r.Read()
		    if err == io.EOF {
			      break
		    }
		    if err != nil {
			      log.Fatal(err)
		    }
        s = append(s, question{record[0], record[1]})
        total++
    }

    // Start the quiz!
    fmt.Println("Weclome to my quiz!\n")
    timer := time.NewTimer(time.Duration(*timeLimit)* time.Second)

    for _, question := range s {
      fmt.Printf("Question: %s\n", question.q)
      answerCh := make(chan string)
      go func() {
        var answer = ""
        fmt.Scanf("%s\n", &answer)
        answerCh <- answer
      }()
      select {
      case <-timer.C:
        fmt.Printf("Correct Answers: %d\n", correct)
        fmt.Printf("Incorrect Answers: %d\n", wrong)
        fmt.Printf("Score: %d out of %d = %.2f", correct, total, (float32(correct) / float32(total))*float32(100))
        return
      case answer := <-answerCh:
        if answer == question.a {
            correct++
        } else {
            wrong++
        }
      }
	  }

    fmt.Printf("Correct Answers: %d\n", correct)
    fmt.Printf("Incorrect Answers: %d\n", wrong)
    fmt.Printf("Score: %d out of %d = %.2f", correct, total, (float32(correct) / float32(total))*float32(100))

}
