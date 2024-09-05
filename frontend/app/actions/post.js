
export async function fetchCreatePost(formData) {
  try {
    const response = await fetch("http://localhost:8080/post/create", {
      method: "POST",
      body: formData,
      credentials: "include",
      cache: "no-store",
    });

    if (!response.ok) {
      throw new Error("Échec de la création du post");
    }
    return;
  } catch (error) {
    console.error("Erreur lors de la création du post :", error);
    throw error;
  }
}

export async function fetchAllPosts() {
  try {
    const response = await fetch('http://localhost:8080/posts', {
      method: 'GET',
      credentials: 'include',
      cache: "no-store",
    });
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    const data = await response.json();
    // setPosts(data);
    // setPosts(data);
    // console.log('Posts fetched:', data);
  } catch (error) {
    console.error('Error fetching posts:', error);
  }
}

export async function fetchPostCreateComments(formData) {
  try {
    const response = await fetch("http://localhost:8080/comment/create", {
      method: "POST",
      body: formData,
      credentials: "include",
      cache: "no-store",
    });

    console.log("REPONSE", response);
    

    console.log("REPONSE", response);
    
    if (!response.ok) {
      throw new Error("Échec de la création du comment");
      // throw new Error("Échec de la création du comment");
    }

    return;
  } catch (error) {
    console.error("Erreur lors de la création du comment :", error);
    throw error;
  }
}

export async function fetchAllPostComments(postId) {
  try {
      const response = await fetch(`http://localhost:8080/comments?postId=${postId}`, {
          method: 'GET',
          credentials: 'include',
          cache: "no-store",
      });

      if (!response.ok) {
          throw new Error('Network response was not ok');
      }

      const data = await response.json();
      return data; // Retourner les commentaires
  } catch (error) {
      console.error('Error fetching comments:', error);
      throw error;
  }
}

export async function fetchLike(formData) {
  try {
    const response = await fetch("http://localhost:8080/like", {
      method: "POST",
      body: formData,
      credentials: "include",
      cache: "no-store",
    });

    if (!response.ok) {
      throw new Error("Échec de la création du post");
    }

    return;
  } catch (error) {
    console.error("Erreur lors de la création du post :", error);
    throw error;
  }
}