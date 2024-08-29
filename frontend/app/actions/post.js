
export async function fetchCreatePost(formData) {
  try {
    const response = await fetch("http://localhost:8080/post/create", {
      method: "POST",
      body: formData,
      credentials: "include",
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

export async function fetchAllPosts(setPosts) {
  try {
    const response = await fetch('http://localhost:8080/posts', {
      method: 'GET',
      credentials: 'include',
    });
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    const data = await response.json();
    setPosts(data);
    console.log('Posts fetched:', data);
  } catch (error) {
    console.error('Error fetching posts:', error);
  }
}