package lobby

import (
	"Waffle/bancho/osu/base_packet_structures"
)

// GetSlotFromUserId is a utility function to get a slot from a players ID
func (multiLobby *MultiplayerLobby) GetSlotFromUserId(userId int32) int {
	for i := 0; i != 8; i++ {
		if multiLobby.MatchInformation.SlotUserId[i] == userId {
			return i
		}
	}

	return -1
}

// GetSlotFromUsername is a utility function to get a slot from a players username
func (multiLobby *MultiplayerLobby) GetSlotFromUsername(username string) int {
	for i := 0; i != 8; i++ {
		if multiLobby.MultiClients[i].GetUsername() == username {
			return i
		}
	}

	return -1
}

// GetOpenSlotCount is a utility function which returns the amount of slots that players can occupy
func (multiLobby *MultiplayerLobby) GetOpenSlotCount() int {
	count := 0

	for i := 0; i != 8; i++ {
		if multiLobby.MatchInformation.SlotStatus[i] != base_packet_structures.MultiplayerMatchSlotStatusLocked {
			count++
		}
	}

	return count
}

// HaveAllPlayersSkipped is a utility function which checks if everyone skipped
func (multiLobby *MultiplayerLobby) HaveAllPlayersSkipped() bool {
	for i := 0; i != 8; i++ {
		if multiLobby.MatchInformation.SlotStatus[i] == base_packet_structures.MultiplayerMatchSlotStatusPlaying && !multiLobby.PlayerSkipRequested[i] {
			return false
		}
	}

	return true
}

// HaveAllPlayersCompleted is a utility function which checks if everyone completed the map
func (multiLobby *MultiplayerLobby) HaveAllPlayersCompleted() bool {
	count := 0

	for i := 0; i != 8; i++ {
		if multiLobby.PlayerCompleted[i] {
			count++
		}
	}

	return count == multiLobby.GetUsedUpSlots()
}

// HaveAllPlayersLoaded is a utility function which checks if everyone loaded in fine
func (multiLobby *MultiplayerLobby) HaveAllPlayersLoaded() bool {
	count := 0

	for i := 0; i != 8; i++ {
		if multiLobby.PlayersLoaded[i] {
			count++
		}
	}

	return count == multiLobby.GetUsedUpSlots()
}

// GetUsedUpSlots is a utility function which returns the slots that are occupied by players
func (multiLobby *MultiplayerLobby) GetUsedUpSlots() int {
	count := 0

	for i := 0; i != 8; i++ {
		if (multiLobby.MatchInformation.SlotStatus[i] & base_packet_structures.MultiplayerMatchSlotStatusHasPlayer) > 0 {
			count++
		}
	}

	return count
}
