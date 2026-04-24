"use client";
import { useForm } from "react-hook-form";
import React, { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useToast } from "@/components/ui/use-toast";
import { useRouter } from "next/navigation";
import { authentificationLogin } from "@/app/actions/authentification";
import Link from "next/link";
import { LockIcon, AtSignIcon } from "lucide-react";
import { ValidateInput } from "@/app/actions/input"; // Assume ValidateInput exists

const LoginForm = () => {
  const router = useRouter();
  const { toast } = useToast();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();
  const [loading, setLoading] = useState(false);

  // Function to validate individual fields
  const validateField = (field, value) => {
    const { isValid, message } = ValidateInput(value);
    if (!isValid) {
      toast({
        title: "Validation Error",
        description: message,
      });
      return false;
    }
    return true;
  };

  const onSubmit = async (formdata) => {
    // Validate individual fields
    if (!validateField("emailOrUsername", formdata.emailOrUsername)) return;
    if (!validateField("password", formdata.password)) return;

    // If fields are valid, proceed with login
    const data = new FormData();
    data.append("emailOrUsername", formdata.emailOrUsername);
    data.append("password", formdata.password);
    setLoading(true);
    try {
      const response = await fetch("http://localhost:8080/login", {
        method: "POST",
        credentials: "include",
        body: data,
      });
      console.log(response);
      if (response.status !== 200) {
        toast({
          title: "Login Failed",
          description:
            response.message || "Invalid credentials. Please try again.",
          status: "error",
        });
        return;
      }

      toast({
        title: "Login Successful",
        description: "You have been logged in successfully."
      });
      router.push("/");
    } catch (error) {
      toast({
        title: "Login Failed",
        description: "An error occurred during login. Please try again.",
      });
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="w-full flex justify-center items-center flex-col min-h-screen">
      <form className="w-full max-w-md" onSubmit={handleSubmit(onSubmit)}>
        <div className="space-y-2">
          <Label htmlFor="emailOrUsername">Email or Username</Label>
          <div className="relative">
            <AtSignIcon
              className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"
              size={18}
            />
            <Input
              id="emailOrUsername"
              placeholder="Enter your email or username"
              className="pl-10"
              {...register("emailOrUsername", {
                required: "Email or username is required",
              })}
            />
          </div>
          {errors.emailOrUsername && (
            <p className="text-sm text-red-500">
              {errors.emailOrUsername.message}
            </p>
          )}
        </div>

        {/* Password Field */}
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

        {/* Submit Button */}
        <Button className="my-4 w-full" type="submit" disabled={loading}>
          {loading ? "Logging in..." : "Login"}
        </Button>
      </form>
      <div className="text-center">
        Need an account ?
        <Link className="text-primary font-bold" href="/auth">
          {" "}
          Register
        </Link>
        .
      </div>
    </div>
  );
};

export default LoginForm;
