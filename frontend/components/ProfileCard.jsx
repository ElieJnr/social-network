"use client";

import { useState, useEffect } from "react";
//import { useRouter } from "next/navigation";
import { userConnect } from "@/app/actions/users";

const { CalendarDaysIcon, UsersIcon } = require("lucide-react");
const { Card, CardContent } = require("./ui/card");
const { Avatar, AvatarImage, AvatarFallback } = require("./ui/avatar");

export default function ProfileCard() {
  // const router = useRouter()
  const [user, setUser] = useState(null)
  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await userConnect();
        setUser(response.user);
      } catch (error) {
        console.error("Error fetching user:", error);
      }
    };
    fetchUser();
  }, []);

  return (
    <Card>
      <CardContent className="flex flex-col items-center gap-4 p-6">
        <Avatar className="w-20 h-20">
          <AvatarImage src={user?.avatar === "" ? "/placeholder-user.jpg": `/uploads/${user?.avatar}`} alt="@shadcn" />
          <AvatarFallback>CN</AvatarFallback>
        </Avatar>
        <div className="text-center space-y-1">
          <div className="font-semibold">{user ? user.firstname : "loading"} {user ? user.lastname : ""}</div>
          <div className="text-muted-foreground">{user ? user.email : ""}</div>
        </div>
        <div className="flex items-center gap-4">
          <div className="flex items-center gap-1 text-muted-foreground">
            <UsersIcon className="h-4 w-4" />
            <span>{user && user.follows ? user.follows.length : "0"} following</span>
          </div>
          <div className="flex items-center gap-1 text-muted-foreground">
            <UsersIcon className="h-4 w-4" />
            <span>{user?.followers ? user.followers.length : "0"} followers</span>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}