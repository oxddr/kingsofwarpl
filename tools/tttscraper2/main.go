package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	databasePath = flag.String("database_path", "kingsofwar-pl.sqlite3", "Path to Sqlite database")
	eventID      = flag.Int64("event_id", 0, "Event ID")
	eventURL     = flag.String("event_url", "", "URL of the event on tabletop.to")
	dryRun       = flag.Bool("dry_run", true, "If true, no changes to database are made")
)

func save(results *TournamentResults) error {
	db, err := sql.Open("sqlite3", *databasePath)
	if err != nil {
		return fmt.Errorf("unable to open database from %q: %v", *databasePath, err)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("unable to open transaction: %v", err)
	}

	err = saveTx(tx, results)
	if err != nil {
		rErr := tx.Rollback()
		if rErr != nil {
			return fmt.Errorf("unable to rollback transaction: %v (rollback-causing error: %v)", rErr, err)
		}
		return fmt.Errorf("unable to commit transaction: %v", err)
	}
	return nil
}

func saveTx(tx *sql.Tx, results *TournamentResults) error {
	eid := *eventID
	if eid == 0 {
		eventStmt, err := tx.Prepare("INSERT INTO Events(name, date, location, tabletop_url) values(?,?,?,?)")
		if err != nil {
			return fmt.Errorf("unable to prepare statement: %v", err)
		}

		event := results.Tournament
		result, err := eventStmt.Exec(event.Name, event.Date.Format("2006-01-02"), event.Location, event.URL)
		if err != nil {
			return fmt.Errorf("unable to save Event %q: %v", event.Name, err)
		}
		eid, err = result.LastInsertId()
		if err != nil {
			return fmt.Errorf("unable to get LastInsertId: %v", err)
		}
	}

	resultStmt, err := tx.Prepare("INSERT INTO Results(player, event, faction, tp, bonus_tp, attrition_points) values(?,?,?,?,?,?)")
	if err != nil {
		return fmt.Errorf("unable to prepare statement: %v", err)
	}

	playerStmt, err := tx.Prepare("INSERT INTO Players(tabletop_id, name) values(?, ?)")
	if err != nil {
		return fmt.Errorf("unable to prepate statement: %v", err)
	}

	for _, p := range results.Players {
		_, err := playerStmt.Exec(p.ID, p.Name)
		if err != nil {
			return fmt.Errorf("unable to save Player %q: %v", p.ID, err)
		}

		_, err = resultStmt.Exec(p.ID, eid, p.Faction, p.TP, p.BonusTP, p.AttritionPoints)
		if err != nil {
			return fmt.Errorf("unable to save Result %v: %v", p, err)
		}
	}
	return tx.Commit()
}

func descResults(results *TournamentResults) string {
	res := fmt.Sprintf("Tournament: %s\n", results.Tournament.Name)
	res += "id faction tp bonus_tp attr\n"
	for _, p := range results.Players {
		res += fmt.Sprintf("%s %s %d %d %d\n", p.ID, p.Faction, p.TP, p.BonusTP, p.AttritionPoints)
	}
	return res
}

func main() {
	flag.Parse()

	if *eventURL == "" {
		log.Fatal("--event_url must be provided")
	}
	results, err := Scrape(*eventURL)
	if err != nil {
		log.Fatalf("Unable to scraper %q: %v", *eventURL, err)
	}

	log.Print(descResults(results))
	if !*dryRun {
		log.Println("Saving to database...")
		err := save(results)
		if err != nil {
			log.Fatalf("Unable to save in database: %v", err)
		}
	}
}
