"use client";

import CreatePostCard from "@/components/CreatePostCard";
import NavBar from "@/components/NavBar";
import PostCard from "@/components/PostCard";
import ProfileCard from "@/components/ProfileCard";
import SuggestionsCard from "@/components/SuggestionsCard";
import TrendingCard  from "@/components/TrendingCard";
import { userConnect } from "./actions";

export default function HomePage() {
 console.log( userConnect());
 
  return (
    <div className="flex flex-col h-screen">
      <NavBar />
      <div className="flex-1 grid grid-cols-[350px_1fr_400px] gap-6 p-6">
        <div className="space-y-6">
          <ProfileCard />
        </div>
        <div className="space-y-6">
          <CreatePostCard />
          <PostCard />
          <PostCard />
        </div>
        <div className="space-y-6">
          <TrendingCard />
          <SuggestionsCard />
        </div>
      </div>
    </div>
  );
}
