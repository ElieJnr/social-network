"use client";

import { follow } from "@/app/actions/follow";
import { allUsers, search, searchUsers, userConnect } from "@/app/actions/users";
import { useState, useEffect } from "react";
import Link from "next/link"
import { socketSend } from "@/app/actions/message";

const { Card, CardHeader, CardTitle, CardContent } = require("./ui/card");
const { Avatar, AvatarImage, AvatarFallback } = require("./ui/avatar");
const { Button } = require("./ui/button");

export default function SuggestionsCard({ socket }) {
  const [users, setUsers] = useState(null)
  const [userOnLine, setUserOnLine] = useState(null)
  const [hiddenUsers, setHiddenUsers] = useState([]);
  const [tabFilter, setTabFilter] = useState([])
  const [tabRequest, setTabRequest] = useState([])
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
    if (userOnLine && userOnLine.requestF) {
      const fetchRequests = async () => {
        try {
          const response = await searchUsers(userOnLine.requestF, "followedUser");
          setTabRequest(response.map(item => item.user));

        } catch (error) {
          console.error("Error fetching requests:", error);
        }
      };
      fetchRequests();
    }
  }, [userOnLine]);


  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response = await allUsers();
        setUsers(response.users);
      } catch (error) {
        console.error("Error fetching user:", error);
      }
    };
    fetchUsers();
  }, []);
  useEffect(() => {
    if (users) {
      setTabFilter(search(users, userOnLine?.follows || []));
    } else {
      console.log("Something went wrong! users is null");
    }
  }, [users, userOnLine]);
  const handleClick = (user, statut) => {
    console.log(user.id);
    setHiddenUsers(prev => [...prev, user]);
    follow(userOnLine.id, user.id, statut, true)
    const Message = {
      Type: "notifications",
      ReceiverId: user.id,
      SubType: "sendFollow",
      Content: "wants to follow you",
      IsPrivate: user.IsPrivate
    }

    socketSend(socket, Message)
  }
  if (!users || !userOnLine) {
    return (
      <Card>
        <CardHeader>
          <CardTitle>Suggestions</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          {/* Skeleton for loading suggestions */}
          <div className="space-y-4 animate-pulse">
            {/* Skeleton user suggestion 1 */}
            <div className="flex items-center gap-4">
              <div className="w-10 h-10 bg-gray-300 rounded-full"></div>
              <div className="space-y-1 flex-1">
                <div className="h-4 bg-gray-300 rounded w-3/4"></div>
                <div className="h-3 bg-gray-300 rounded w-1/2"></div>
              </div>
              <div className="w-20 h-8 bg-gray-300 rounded ml-auto"></div>
            </div>

            {/* Skeleton user suggestion 2 */}
            <div className="flex items-center gap-4">
              <div className="w-10 h-10 bg-gray-300 rounded-full"></div>
              <div className="space-y-1 flex-1">
                <div className="h-4 bg-gray-300 rounded w-3/4"></div>
                <div className="h-3 bg-gray-300 rounded w-1/2"></div>
              </div>
              <div className="w-20 h-8 bg-gray-300 rounded ml-auto"></div>
            </div>

            {/* Skeleton user suggestion 3 */}
            <div className="flex items-center gap-4">
              <div className="w-10 h-10 bg-gray-300 rounded-full"></div>
              <div className="space-y-1 flex-1">
                <div className="h-4 bg-gray-300 rounded w-3/4"></div>
                <div className="h-3 bg-gray-300 rounded w-1/2"></div>
              </div>
              <div className="w-20 h-8 bg-gray-300 rounded ml-auto"></div>
            </div>
          </div>
        </CardContent>
      </Card>

    )
  } 

  return (
    <Card>
      <CardHeader>
        <CardTitle>Suggestions</CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        {tabFilter?.filter(user =>
          (user.id !== userOnLine.id && !hiddenUsers.includes(user)) &&
          !tabRequest.some(request => request.id === user.id)
        ).map(user => (
          <div key={user.id} className="flex items-center gap-4">
            <Link href={user ? `/profil?userId=${user.id}` : "#"}
              prefetch={false} className="flex items-center gap-2">
              <Avatar className="w-10 cursor-pointer h-10">
                <AvatarImage src={`/uploads/${user?.avatar}`|| "/placeholder-user.jpg"} alt="@shadcn" />
                <AvatarFallback>CN</AvatarFallback>
              </Avatar>
            </Link>
            <div className="space-y-1">
              <div className="font-semibold">{user ? user.firstname : "loading"} {user ? user.lastname : ""}</div>
              <div className="text-muted-foreground">{""}</div>
            </div>
            <Button onClick={() => handleClick(user, !user.isPrivate)} variant="outline" size="sm" className="ml-auto">
              Follow
            </Button>
          </div>
        ))}
        {tabRequest.map((user) =>
          <div key={user.id} className="flex items-center gap-4">
            <Link href={user ? `/profil?userId=${user.id}` : "#"}
              prefetch={false} className="flex items-center gap-2">
              <Avatar className="w-10 cursor-pointer h-10">
                <AvatarImage src={`/uploads/${user?.avatar}`|| "/placeholder-user.jpg"} alt="@shadcn" />
                <AvatarFallback>CN</AvatarFallback>
              </Avatar>
            </Link>
            <div className="space-y-1">
              <div className="font-semibold">{user ? user.firstname : "loading"} {user ? user.lastname : ""}</div>
              <div className="text-muted-foreground">{""}</div>
            </div>
            <Button variant="outline" size="sm" className="ml-auto">
              Demande
            </Button>
          </div>

        )}

      </CardContent>
    </Card>
  );
}