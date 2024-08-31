package handlers

import (
	"fmt"
	"net/http"
)

func LikeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Like(w, r)
	}
}

func Like(w http.ResponseWriter, r *http.Request) {
	postId := r.FormValue("like-postId")
	if postId == "" {
		fmt.Println("error getting postId:")
		return
	}

	err := LikeService.LikeTreatment(postId, w, r)
	if err != nil {
		fmt.Println("Error LikeTreatment: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
