package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	newEvents, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": fmt.Sprintf("%v", http.StatusInternalServerError), "message": "Internal Server Error"})
		return
	}

	context.JSON(http.StatusOK, newEvents)
}

func getEvent(context *gin.Context) {
	rawId := context.Param("id")
	eventId, err := strconv.ParseInt(rawId, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": fmt.Sprintf("%v", http.StatusInternalServerError), "message": "Data of Integer type required"})
		return
	}

	newEvent, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": fmt.Sprintf("%v", http.StatusInternalServerError), "message": "Internal Server Error"})
		return
	}

	context.JSON(http.StatusOK, newEvent)

}

func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": fmt.Sprintf("%v", http.StatusBadRequest), "message": "Incorrect data passed in POST request"})
		return
	}

	id := context.GetInt64("id")

	event.OrganizerId = id
	event.CreatedAt = time.Now()
	err = event.SaveEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": fmt.Sprintf("%v", http.StatusInternalServerError), "message": "Internal Server Error, Could not create new event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"code": fmt.Sprintf("%v", http.StatusCreated), "message": "Successfully created event", "event": event})
}

func updateEvent(context *gin.Context) {
	/* First step in an update function is to get the event that is to be updated, like I did here i first extracted the id coming through the api call url through context.Param and accessing the name/identifier of the particular value. In our situation it is an integer we need but context.Param always returns a string. so we use strconv.ParseInt to convert our string to an integer.
	 */
	rawId := context.Param("id")
	eventId, err := strconv.ParseInt(rawId, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": fmt.Sprintf("%v", http.StatusInternalServerError), "message": "Data of Integer type required"})
		return
	}

	/* After retrieving the id passed, what we need to do next is get the event at that id. Remember that we had defined a getEventById function so we would use that to retrieve our function */
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": fmt.Sprintf("%v", http.StatusInternalServerError), "message": "Could not fetch record"})
		return
	}

	// Here we extract the id coming from the request(Remember we added the id and email to the jwt as a claim). Then we check it against the id from the event we got from GetEventById
	setId := context.GetInt64("id")

	if event.Id != setId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update/edit this event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": fmt.Sprintf("%v", http.StatusBadRequest), "message": "Incorrect data passed in POST request"})
		return
	}

	updatedEvent.Id = eventId

	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": fmt.Sprintf("%v", http.StatusInternalServerError), "message": "Could not update record"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully updayed record"})
}

func deleteEvent(context *gin.Context) {
	rawId := context.Param("id")
	eventId, err := strconv.ParseInt(rawId, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Data of Integer type required"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch record"})
		return
	}

	setId := context.GetInt64("id")

	if event.Id != setId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update/edit this event"})
		return
	}

	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully Deleted Event"})
}

func emptyEvents(context *gin.Context) {
	err := models.EmptyEventTable()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An error occured trying to delete all records."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully Deleted All Records"})
}

func registerForEvent(context *gin.Context) {
	rawId := context.Param("id")
	eventId, err := strconv.ParseInt(rawId, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Data of Integer type required"})
		return
	}

	userId := context.GetInt64("id")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed at retrieving event", "error": err})
		return
	}

	err = event.RegisterForEvent(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register for event", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered Successfully!"})
}
