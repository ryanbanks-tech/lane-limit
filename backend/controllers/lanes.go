package controllers

import (
	"net/http"
	"strconv"

	"lane-limit/models"

	"github.com/gin-gonic/gin"
)

// For demo: store lanes in a global map
var lanes = make(map[int]models.Lane)

// To generate incremental lane IDs
var currentLaneID = 0

func generateLaneID() int {
	currentLaneID++
	return currentLaneID
}

// UpsertLane creates or updates a lane's info
func UpsertLane(c *gin.Context) {
	var lane models.Lane
	if err := c.ShouldBindJSON(&lane); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// If ID isn't provided, generate a new one
	if lane.ID == 0 {
		lane.ID = generateLaneID()
	}

	// If you receive StartTime/EndTime as strings, parse them here if needed.
	// Example (if your JSON has them as strings):
	//
	//   parsedStart, err := time.Parse(time.RFC3339, someStartTimeString)
	//   if err == nil {
	//       lane.StartTime = parsedStart
	//   }

	// Upsert into map
	lanes[lane.ID] = lane

	c.JSON(http.StatusOK, lane)
}

// GetAllLanes returns all lane info
func GetAllLanes(c *gin.Context) {
	laneList := make([]models.Lane, 0, len(lanes))
	for _, v := range lanes {
		laneList = append(laneList, v)
	}
	c.JSON(http.StatusOK, laneList)
}

// ClearLaneInfo clears a single lane
func ClearLaneInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lane ID"})
		return
	}

	delete(lanes, id)
	c.JSON(http.StatusOK, gin.H{"message": "Lane cleared", "laneID": id})
}

// ClearAllLanes clears all lanes
func ClearAllLanes(c *gin.Context) {
	lanes = make(map[int]models.Lane)
	currentLaneID = 0
	c.JSON(http.StatusOK, gin.H{"message": "All lanes cleared"})
}
