
import * as z from "zod";
import { makePostFetch } from ".";
import { domain } from "..";

const registerSchema = z.object({
  email: z.string().email({ message: "Invalid email address" }),
  password: z
    .string()
    .min(8, { message: "Password must be at least 8 characters long" }),
  firstname: z.string().min(1, { message: "First name is required" }),
  lastname: z.string().min(1, { message: "Last name is required" }),
  dateOfBirth: z
    .string()
    .refine((date) => !isNaN(Date.parse(date)), { message: "Invalid date" }),
  avatar: z.any().optional(),
  username: z.string().optional(),
  bio: z.string().optional(),
});

const loginSchema = z.object({
  emailOrUsername: z
    .string()
    .min(1, { message: "Email or username is required" }),
  password: z
    .string()
    .min(8, { message: "Password must be at least 8 characters long" }),
});

export async function authentification(formData) {
  const parsedBody = registerSchema.parse({
    email: formData.get("email"),
    password: formData.get("password"),
    firstname: formData.get("firstname"),
    lastname: formData.get("lastname"),
    dateOfBirth: formData.get("dateOfBirth"),
    avatar: formData.get("avatar"),
    username: formData.get("username") || "",
    bio: formData.get("bio") || "",
  });

  const data = await makePostFetch(domain + "/signin", formData);
  return data;
}

export async function authentificationLogin(formData) {
  const parsedBody = loginSchema.parse( {
    "emailOrUsername" : formData.get("emailOrUsername"),
    "password" : formData.get("password")
});
  const data = await makePostFetch(domain + "/login", formData);
  return data;
}
