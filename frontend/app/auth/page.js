"use client";
import { useState } from "react";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { useToast } from "@/components/ui/use-toast";
import { useRouter } from "next/navigation";
import {
  LockIcon,
  MailIcon,
  UserIcon,
  CalendarIcon,
  ImageIcon,
  AtSignIcon,
  FileTextIcon,
} from "lucide-react";
import { useForm } from "react-hook-form";
import { authentification } from "../actions/authentification";

export default function RegisterForm() {
  const { toast } = useToast();
  const router = useRouter();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();
  const [loading, setLoading] = useState(false);

  const onSubmit = async (formData) => {
    setLoading(true);
    const data = new FormData();

    // Ajout des champs texte
    data.append("firstname", formData.firstname);
    data.append("lastname", formData.lastname);
    data.append("email", formData.email);
    data.append("username", formData.username || ""); // Optionnel
    data.append("bio", formData.bio || ""); // Optionnel
    data.append("dateOfBirth", formData.dateOfBirth);
    data.append("password", formData.password);

    // Ajout du fichier d'avatar s'il existe
    if (formData.avatar && formData.avatar.length > 0) {
      data.append("avatar", formData.avatar[0]);
    }

    try {
      const response = await authentification(data);
      
      if (response.status != 201) {
        toast({
          title: "Registration Failed",
          description: response.message,
        });
        return;
      }
      toast({
        title: "Registration Successful",
        description: "Your account has been created.",
      });
      router.push("/auth/login");
    } catch (error) {
      toast({
        title: "Registration Failed",
        description: "An error occurred during registration. Please try again.",
      });
    } finally {
      setLoading(false);
    }
  };
  return (
    <div className="w-full flex justify-center items-center min-h-screen">
    <div className="w-full max-w-md">
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="grid grid-cols-2 gap-4">
          {/* Prénom */}
          <div className="space-y-2">
            <Label htmlFor="firstname">First Name</Label>
            <div className="relative">
              <UserIcon
                className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
                size={18}
              />
              <Input
                id="firstname"
                placeholder="Enter your first name"
                className="pl-10"
                {...register("firstname", {
                  required: "First name is required",
                })}
              />
            </div>
            {errors.firstname && (
              <p className="text-sm text-red-500">{errors.firstname.message}</p>
            )}
          </div>

          {/* Nom */}
          <div className="space-y-2">
            <Label htmlFor="lastname">Last Name</Label>
            <div className="relative">
              <UserIcon
                className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
                size={18}
              />
              <Input
                id="lastname"
                placeholder="Enter your last name"
                className="pl-10"
                {...register("lastname", { required: "Last name is required" })}
              />
            </div>
            {errors.lastname && (
              <p className="text-sm text-red-500">{errors.lastname.message}</p>
            )}
          </div>

          {/* Email */}
          <div className="space-y-2">
            <Label htmlFor="email">Email</Label>
            <div className="relative">
              <MailIcon
                className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
                size={18}
              />
              <Input
                id="email"
                type="email"
                placeholder="Enter your email"
                className="pl-10"
                {...register("email", { required: "Email is required" })}
              />
            </div>
            {errors.email && (
              <p className="text-sm text-red-500">{errors.email.message}</p>
            )}
          </div>

          {/* Username (optionnel) */}
          <div className="space-y-2">
            <Label htmlFor="username">Username (Optional)</Label>
            <div className="relative">
              <AtSignIcon
                className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
                size={18}
              />
              <Input
                id="username"
                placeholder="Enter your username"
                className="pl-10"
                {...register("username")}
              />
            </div>
          </div>

          {/* Avatar (optionnel) */}
          <div className="space-y-2">
            <Label htmlFor="avatar">Avatar/Image (Optional)</Label>
            <div className="relative">
              <ImageIcon
                className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
                size={18}
              />
              <Input
                id="avatar"
                type="file"
                className="pl-10"
                {...register("avatar")}
              />
            </div>
          </div>

          {/* Date de naissance */}
          <div className="space-y-2">
            <Label htmlFor="dateOfBirth">Date of Birth</Label>
            <div className="relative">
              <CalendarIcon
                className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
                size={18}
              />
              <Input
                id="dateOfBirth"
                type="date"
                className="pl-10"
                {...register("dateOfBirth", {
                  required: "Date of birth is required",
                })}
              />
            </div>
            {errors.dateOfBirth && (
              <p className="text-sm text-red-500">
                {errors.dateOfBirth.message}
              </p>
            )}
          </div>
        </div>

        <div className="space-y-2">
          <Label htmlFor="password">Password</Label>
          <div className="relative">
            <LockIcon
              className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
              size={18}
            />
            <Input
              id="password"
              type="password"
              placeholder="Enter your password"
              className="pl-10"
              {...register("password", { required: "Password is required" })}
            />
          </div>
          {errors.password && (
            <p className="text-sm text-red-500">{errors.password.message}</p>
          )}
        </div>
        {/* Bio (optionnel) */}
        <div className="space-y-2">
          <Label htmlFor="bio">About Me (Optional)</Label>
          <div className="relative">
            <FileTextIcon
              className="absolute left-3 top-3 text-gray-400"
              size={18}
            />
            <Textarea
              id="bio"
              placeholder="Tell us about yourself"
              className="pl-10 min-h-[100px]"
              {...register("bio")}
            />
          </div>
        </div>
        

        {/* Bouton d'inscription */}
        <Button className="my-4 w-full" type="submit" disabled={loading}>
          {loading ? "Registering..." : "Register"}
        </Button>
      </form>
      <div className="text-center">
        Already have an account ?
        <Link className="text-primary font-bold" href="/auth/login"> log in </Link>
        .
      </div>
    </div>
    </div>
  );
}
// import React from 'react'
// import AuthForm from './auth-form'

// const page = () => {
//   return (
//     <div><AuthForm/></div>
//   )
// }

// export default page