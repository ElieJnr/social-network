"use client";

import { follow } from "@/app/actions/follow";
import { allUsers, userConnect } from "@/app/actions/users";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
const { Avatar, AvatarImage, AvatarFallback } = require("./ui/avatar");
const { Button } = require("./ui/button");
const { Card, CardHeader, CardTitle, CardContent } = require("./ui/card");

export default function SuggestionsCard() {
  const [users, setUsers] = useState(null)
  const [userOnLine, setUserOnLine] = useState(null)
  const [hiddenUsers, setHiddenUsers] = useState([]);
  const router = useRouter()

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await userConnect();
        setUserOnLine(response.user);
      } catch (error) {
        console.error("Error fetching user:", error);
      }
    };
    fetchUser();
  }, []);
  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response = await allUsers();
        console.log(response);
        setUsers(response.users);
        setUsers(users.filter(user => 
          user?.follows?.some(follow => follow.followedUser === userConnect.id)
        ));
      } catch (error) {
        console.error("Error fetching user:", error);
      }
    };
    fetchUsers();
  }, []);
  console.log(users);
 
  
  const handleClick = (user, statut) => {
    console.log(user.id);
    setHiddenUsers(prev => [...prev, user]);
    follow(userOnLine.id, user.id, statut )
  }
  if (!users || !userOnLine) {
    return <div>Loading...</div>;
  }
  return (
    <Card>
      <CardHeader>
        <CardTitle>Suggestions</CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
      {users?.filter(user => user.id !== userOnLine.id && !hiddenUsers.includes(user)).map(user => (
        <div key={user.id} className="flex items-center gap-4">
          <Avatar className="w-10 h-10">
            <AvatarImage src="/placeholder-user.jpg" alt="@shadcn" />
            <AvatarFallback>CN</AvatarFallback>
          </Avatar>
          <div className="space-y-1">
            <div className="font-semibold">{user ? user.firstname : "loading"} {user ? user.lastname : ""}</div>
            <div className="text-muted-foreground">{ ""}</div>
          </div>
          <Button onClick={() => handleClick(user, !user.isPrivate)} variant="outline" size="sm" className="ml-auto">
            Follow
          </Button>
        </div>
      ))}
      </CardContent>
    </Card>
  );
}