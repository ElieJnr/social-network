
export async function makeGetFetch(url) {
  const response = await fetch(url, {
    credentials: "include",
  });
  const data = await response.json();
  return data;
}

export async function makePostFetch(url, formData) {
  const response = await fetch(url, {
    method: "POST",
    credentials: "include",
    // cache: "no-store",
    body: formData,
  });

  const data = await response.json();
  return data;
}


export async function userConnect() {
  const response = await fetch("http://localhost:8080/userConnect", {
    method: "GET",
    credentials: "include",
    // cache: "no-store",
  });

  const data = await response.json();
  console.log(data);
  return data;
}


// export async function fetchPost(formData) {
//   try {
//     const response = await fetch("http://localhost:8080/post/create", {
//       method: "POST",
//       body: formData,
//       credentials: "include",
//     });

//     if (!response.ok) {
//       throw new Error("Échec de la création du post");
//     }

//     return;
//   } catch (error) {
//     console.error("Erreur lors de la création du post :", error);
//     throw error;
//   }
// }
