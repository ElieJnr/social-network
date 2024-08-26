import { NextResponse } from 'next/server';

export async function POST(req) {
  try {
    const response = await fetch("http://localhost:8080/post", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include", // Inclut les cookies automatiquement
    });

    console.log("Response from Go server:", response.status);
    const data = await response.json();
    console.log("Data from Go server:", data);

    return NextResponse.json(data, { status: response.status });
  } catch (error) {
    console.error("Error in /api/post:", error);
    return NextResponse.json(
      { error: "Une erreur est survenue" },
      { status: 500 }
    );
  }
}

export async function fetchPosts() {
  try {
    const response = await fetch("http://localhost:8080/posts", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include", // Inclut les cookies automatiquement
    });

    console.log("Response from Go server:", response.status);
    const data = await response.json();
    console.log("Data from Go server:", data);

    return data;
  } catch (error) {
    console.error("Error in /api/posts:", error);
    return { error: "Une erreur est survenue" };
  }
}

export async function fetchCreatePost() {
  try {
    const response = await fetch("http://localhost:8080/post", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include", // Inclut les cookies automatiquement
    });

    console.log("Response from Go server:", response.status);
    const data = await response.json();
    console.log("Data from Go server:", data);

    return data;
  } catch (error) {
    console.error("Error in /api/post:", error);
    return { error: "Une erreur est survenue" };
  }
}