package database

import (
	"Waffle/helpers"

	_ "github.com/go-sql-driver/mysql"
)

// UserStatsFromDatabase retrieves a users stats given their user id and the mode that it should be retrieved for
func UserStatsFromDatabase(id uint64, mode int8) (int8, UserStats) {
	returnStats := UserStats{}

	queryResult, queryErr := Database.Query("SELECT * FROM (SELECT user_id, mode, ROW_NUMBER() OVER (ORDER BY ranked_score DESC) AS 'rank', ranked_score, total_score, user_level, accuracy, playcount, count_ssh, count_ss, count_sh, count_s, count_a, count_b, count_c, count_d, hit300, hit100, hit50, hitMiss, hitGeki, hitKatu, replays_watched, playtime FROM waffle.stats WHERE mode = ? AND user_id != 1) t WHERE user_id = ?", mode, id)

	if queryErr != nil {
		helpers.Logger.Printf("[Database] Failed to Fetch User Stats from Database, MySQL query failed.\n")

		if queryResult != nil {
			queryResult.Close()
		}

		return -2, returnStats
	}

	if queryResult.Next() {
		scanErr := queryResult.Scan(&returnStats.UserID, &returnStats.Mode, &returnStats.Rank, &returnStats.RankedScore, &returnStats.TotalScore, &returnStats.Level, &returnStats.Accuracy, &returnStats.Playcount, &returnStats.CountSSH, &returnStats.CountSS, &returnStats.CountSH, &returnStats.CountS, &returnStats.CountA, &returnStats.CountB, &returnStats.CountC, &returnStats.CountD, &returnStats.Hit300, &returnStats.Hit100, &returnStats.Hit50, &returnStats.HitMiss, &returnStats.HitGeki, &returnStats.HitKatu, &returnStats.ReplaysWatched, &returnStats.Playtime)

		queryResult.Close()

		if scanErr != nil {
			helpers.Logger.Printf("[Database] Failed to Scan Database results onto UserStats object.\n")

			return -2, returnStats
		}

		return 0, returnStats
	}

	return -1, returnStats
}

// UserStatsGetWaffleBot retrieves WaffleBot's stats
func UserStatsGetWaffleBot(mode int8) (int8, UserStats) {
	returnStats := UserStats{}

	queryResult, queryErr := Database.Query("SELECT * FROM (SELECT user_id, mode, ROW_NUMBER() OVER (ORDER BY ranked_score DESC) AS 'rank', ranked_score, total_score, user_level, accuracy, playcount, count_ssh, count_ss, count_sh, count_s, count_a, count_b, count_c, count_d, hit300, hit100, hit50, hitMiss, hitGeki, hitKatu, replays_watched FROM waffle.stats WHERE mode = ?) t WHERE user_id = 1", mode)

	if queryErr != nil {
		helpers.Logger.Printf("[Database] Failed to Fetch User Stats from Database, MySQL query failed.\n")

		if queryResult != nil {
			queryResult.Close()
		}

		return -2, returnStats
	}

	if queryResult.Next() {
		scanErr := queryResult.Scan(&returnStats.UserID, &returnStats.Mode, &returnStats.Rank, &returnStats.RankedScore, &returnStats.TotalScore, &returnStats.Level, &returnStats.Accuracy, &returnStats.Playcount, &returnStats.CountSSH, &returnStats.CountSS, &returnStats.CountSH, &returnStats.CountS, &returnStats.CountA, &returnStats.CountB, &returnStats.CountC, &returnStats.CountD, &returnStats.Hit300, &returnStats.Hit100, &returnStats.Hit50, &returnStats.HitMiss, &returnStats.HitGeki, &returnStats.HitKatu, &returnStats.ReplaysWatched)

		queryResult.Close()

		if scanErr != nil {
			helpers.Logger.Printf("[Database] Failed to Scan Database results onto UserStats object.\n")

			return -2, returnStats
		}

		return 0, returnStats
	}

	return -1, returnStats
}
