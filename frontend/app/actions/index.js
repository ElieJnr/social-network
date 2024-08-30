
export async function makeGetFetch(url) {
  const response = await fetch(url, {
    credentials: "include",
  });
  const data = await response.json();
  return data;
}

export async function makeGetFetchUsers(url) {
  const response = await fetch(url, {
    next: { revalidate: 2 },
    credentials: "include",
  });
  const data = await response.json();
  return data;
}

export async function makePostFetch(url, formData) {
  const response = await fetch(url, {
    method: "POST",
    credentials: "include",
    cache: "no-store",
    body: formData,
  });

  const data = await response.json();
  return data;
}




export async function fetchPost(url, formData) {
  const response = await fetch(url, {
    method: "POST",
    credentials: "include",
    cache: "no-store",
    body: JSON.stringify(formData),
  });
  return response
}
