package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"socialNetwork/pkg/models"
	"socialNetwork/pkg/services"
	"socialNetwork/utils"
)

func CreateEventHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event models.Event
		eventService := services.NewEventService()
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		user, err := utils.CurrentUser(w, r)
		if err != nil {
			http.Error(w, "Could not create event", http.StatusInternalServerError)
			return
		}
		event.MemberID = user.UserId
		// Create the event
		err = eventService.CreateEvent(event)
		fmt.Println("err", err)
		if err != nil {
			http.Error(w, "Could not create event", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func RespondToEventHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventService := services.NewEventService()
		var response models.EventResponse
		err := json.NewDecoder(r.Body).Decode(&response)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		user, err := utils.CurrentUser(w, r)
		if err != nil {
			http.Error(w, "Could not identify user", http.StatusInternalServerError)
			return
		}
		response.MemberID = user.UserId

		// Check if the user has already responded to the event
		existingResponse, err := eventService.GetResponseByEventAndMember(response.EventID, user.UserId)
		if err == nil && existingResponse != nil {
			// If the user has responded, update the response
			err = eventService.UpdateResponse(response)
		} else {
			// Create a new response
			err = eventService.RespondToEvent(response)
		}
		if err != nil {
			http.Error(w, "Could not respond to event", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetEventsByGroupHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventService := services.NewEventService()

		groupId := r.URL.Query().Get("groupId")
		if groupId == "" {
			http.Error(w, "Group ID is required", http.StatusBadRequest)
			return
		}

		events, err := eventService.GetEventsByGroupId(groupId)
		if err != nil {
			http.Error(w, "Could not fetch events", http.StatusInternalServerError)
			return
		}
		fmt.Println("events",events)
		w.Header().Set("Content-Type", "application/json")
		iss := len(events) > 0
		json.NewEncoder(w).Encode(struct {
			Event   []models.Event `json:"Event"`
			Isevent bool           `json:"Isevent"`
		}{
			Event:   events,
			Isevent: iss,
		})
	}
}

func GetResponsesByGroupAndMemberHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventService := services.NewEventService()
		groupId := r.URL.Query().Get("groupId")
		if groupId == "" {
			http.Error(w, "Group ID is required", http.StatusBadRequest)
			return
		}

		user, err := utils.CurrentUser(w, r)
		if err != nil {
			http.Error(w, "Could not identify user", http.StatusInternalServerError)
			return
		}
		responses, err := eventService.GetResponseByEventAndMember(groupId, user.UserId)
		if err != nil {
			http.Error(w, "Could not fetch responses", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responses)
	}
}
