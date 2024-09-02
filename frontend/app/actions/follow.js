import { fetchPost } from ".";

export async function follow(userId, followedId, statut, followU) {
	let formData = {
	  userId: userId,
	  followedId: followedId,
	  statut: statut,
	  followU : followU
	};
	
	let url = "http://localhost:8080/follow";
	try {
	   await fetchPost(url, formData);
	} catch (error) {
	  console.error("Error following user:", error);
	}
}

