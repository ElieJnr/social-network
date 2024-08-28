
export async function fetchPost(formData) {
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