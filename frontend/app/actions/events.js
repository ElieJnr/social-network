"use client";
import { domain } from "..";

const { fetchPost } = require(".");

export const CreateEvent = async (formData) => {
  const data = fetchPost(domain + "/group/createEvent", formData);
  return data;
};

export const fetchEvents = async (id) => {
  try {
    const response = await fetch(domain + `/group/getEvents?groupId=${id}`, {
      method: "GET",
      credentials: "include",
      cache: "no-store",
    });

    const data = response.json();
    return data;
  } catch (error) {
    console.error("Error fetching events:", error);
    throw error;
  }
};
