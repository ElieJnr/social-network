"use client";

import CreatePostCard from "@/components/CreatePostCard";
import PostCard from "@/components/PostCard";
import ProfileCard from "@/components/ProfileCard";
import SuggestionsCard from "@/components/SuggestionsCard";
import { ClickMessageApp } from "@/components/ui/message";
import { useWebSocket } from "./actions/message";
import NavBar from "@/components/Nav";
export let socketGlobal
export default function HomePage() {
  const socket = useWebSocket('ws://localhost:8080/ws');
  return (
    <div className="flex flex-col h-screen">
      <NavBar socket={socket} />
      <div className="flex-1 grid grid-cols-[350px_1fr_400px] gap-6 p-6">
        <div className="space-y-6">
          <ProfileCard />
          <SuggestionsCard socket={socket}/>
        </div>
        <div className="space-y-6">
          <div className="w-full p-2">
            <CreatePostCard />
          </div>
          <PostCard />
        </div>
        <div className="space-y-6">
          <ClickMessageApp socket={socket} />
        </div>
      </div>
    </div>
  );
}

