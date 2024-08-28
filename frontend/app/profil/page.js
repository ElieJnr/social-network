"use client";

import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { useState, useEffect } from "react";
import { GetAllInfoForUserById } from "../actions/users";


export default function Profil() {
  const [userId, setUserId] = useState(null);
  const [user, setUser] = useState(null)

  useEffect(() => {
    const urlParams = new URLSearchParams(window.location.search);
    const userIdFromQuery = urlParams.get('userId');
    setUserId(userIdFromQuery);
  }, []); 

  useEffect(() => {
    const fetchUserData = async () => {
      if (userId) {
        try {
          const userData = await GetAllInfoForUserById(userId);
          setUser(userData);
        } catch (error) {
          console.error("Erreur lors de la récupération de l'utilisateur:", error);
        }
      }
    };

    fetchUserData();
  }, [userId]); 


  if (!user) {
    return <div>Loading...</div>; 
  }

  

  return (
    <div className="bg-background text-foreground min-h-screen flex flex-col">
      <header className="bg-card py-4 px-6 flex items-center justify-between shadow-sm">
        <div className="flex items-center gap-4">
          <Avatar className="h-10 w-10">
            <AvatarImage src="/placeholder-user.jpg" alt="@shadcn" />
            <AvatarFallback>JD</AvatarFallback>
          </Avatar>
          <div className="grid gap-1">
            <div className="font-medium">John Doe</div>
            <div className="text-xs text-muted-foreground">@johndoe</div>
          </div>
        </div>
        <Button variant="outline" size="sm">
          Edit Profile
        </Button>
      </header>
      <div className="flex-1 grid gap-6 p-6">
        <div className="grid gap-4">
          <div className="flex items-center justify-between">
            <div className="grid gap-1">
              <div className="font-medium text-2xl">John Doe</div>
              <div className="text-muted-foreground">Software Engineer</div>
            </div>
            <div className="flex items-center gap-4">
              <Button variant="outline" size="sm">
                Follow
              </Button>
              <Button variant="ghost" size="icon" className="rounded-full">
                <MoveHorizontalIcon className="w-5 h-5" />
              </Button>
            </div>
          </div>
          <div className="text-sm leading-loose text-muted-foreground">
            I'm a software engineer with a passion for building innovative products. In my free time, I enjoy exploring
            new technologies and reading about the latest industry trends.
          </div>
        </div>
        <div className="grid sm:grid-cols-3 gap-4 text-center">
          <div className="bg-card rounded-lg p-4">
            <div className="font-medium">100</div>
            <div className="text-xs text-muted-foreground">Posts</div>
          </div>
          <div className="bg-card rounded-lg p-4">
            <div className="font-medium">1.2K</div>
            <div className="text-xs text-muted-foreground">Followers</div>
          </div>
          <div className="bg-card rounded-lg p-4">
            <div className="font-medium">500</div>
            <div className="text-xs text-muted-foreground">Following</div>
          </div>
        </div>
        <div className="grid gap-6">
          <div className="flex items-center justify-between">
            <div className="font-medium">Posts</div>
            <div className="flex items-center gap-2">
              <Button variant="outline" size="sm">
                <EyeIcon className="w-4 h-4 mr-2" />
                Public
              </Button>
              <Button variant="outline" size="sm">
                <LockIcon className="w-4 h-4 mr-2" />
                Private
              </Button>
            </div>
          </div>
          <div className="grid sm:grid-cols-2 md:grid-cols-3 gap-4">
            <div className="bg-card rounded-lg overflow-hidden">
              <img
                src="/placeholder.svg"
                alt="Post Image"
                width={600}
                height={400}
                className="w-full aspect-[3/2] object-cover" />
              <div className="p-4">
                <div className="font-medium line-clamp-2">Introducing the new Vercel platform</div>
                <div className="text-xs text-muted-foreground">2 days ago</div>
              </div>
            </div>
            <div className="bg-card rounded-lg overflow-hidden">
              <img
                src="/placeholder.svg"
                alt="Post Image"
                width={600}
                height={400}
                className="w-full aspect-[3/2] object-cover" />
              <div className="p-4">
                <div className="font-medium line-clamp-2">Building a Serverless API with Next.js</div>
                <div className="text-xs text-muted-foreground">1 week ago</div>
              </div>
            </div>
            <div className="bg-card rounded-lg overflow-hidden">
              <img
                src="/placeholder.svg"
                alt="Post Image"
                width={600}
                height={400}
                className="w-full aspect-[3/2] object-cover" />
              <div className="p-4">
                <div className="font-medium line-clamp-2">Mastering Tailwind CSS: A Comprehensive Guide</div>
                <div className="text-xs text-muted-foreground">2 weeks ago</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

function EyeIcon(props) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round">
      <path d="M2 12s3-7 10-7 10 7 10 7-3 7-10 7-10-7-10-7Z" />
      <circle cx="12" cy="12" r="3" />
    </svg>
  );
}

function LockIcon(props) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round">
      <rect width="18" height="11" x="3" y="11" rx="2" ry="2" />
      <path d="M7 11V7a5 5 0 0 1 10 0v4" />
    </svg>
  );
}

function MoveHorizontalIcon(props) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round">
      <polyline points="18 8 22 12 18 16" />
      <polyline points="6 8 2 12 6 16" />
      <line x1="2" x2="22" y1="12" y2="12" />
    </svg>
  );
}
