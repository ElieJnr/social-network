"use client";
import React, { useState } from 'react';
import { useRouter } from 'next/navigation';

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import Link from "next/link";
import { Textarea } from "@/components/ui/textarea";
import { Label } from "@/components/ui/label";
import { MountainIcon, BellIcon, LogOutIcon, UsersIcon, CalendarDaysIcon, ImageIcon } from "lucide-react";
import { fetchPost } from '@/app/postActions';
import { mutate } from "swr";
import Image from 'next/image';
// import LogoutButton from './butonlgout';
import dynamic from 'next/dynamic';
import LogoutButton from './buttonlogout';
import { ClickMessageApp } from './ui/message';

function NavBar() {
  const router = useRouter();
  const handleLogout = async () => {
    try {
      const response = await fetch('http://localhost:8080/logout', {
        method: 'POST',
        credentials: 'include',
      });
      if (response.ok) {
        router.push('/auth');
      } else {
        console.error('Failed to log out');
      }
    } catch (error) {
      console.error('An error occurred during logout:', error);
    }
  };

  return (
    <header className="bg-primary text-primary-foreground py-4 px-6">
      <div className="container mx-auto flex items-center justify-between">
        <div className="flex items-center gap-4">
          <Link href="#" prefetch={false}>
            <MountainIcon className="h-6 w-6" />
            <span className="sr-only">Acme Inc</span>
          </Link>
          <nav className="hidden md:flex items-center gap-4">
            <Link href="#" className="hover:underline" prefetch={false}>
              Home
            </Link>
            <Link href="#" className="hover:underline" prefetch={false}>
              Groups
            </Link>
            <Link href="#" className="hover:underline" prefetch={false}>
              Chat
            </Link>
          </nav>
        </div>
        <div className="flex items-center">
          <Button variant="ghost" >
            <BellIcon className="h-5 w-5" />
          </Button>
          <Button variant="ghost" onClick={handleLogout}>
            <LogOutIcon className="h-5 w-5" />
          </Button>
          {/* <LogoutButton variant="ghost" /> */}
        </div>
      </div>
    </header>
  );
}

function ProfileCard() {
  //const [user, SetUser] = useState(null)
  console.log(userConnect());
  console.log("ok");
  return (
    <Card>
      <CardContent className="flex flex-col items-center gap-4 p-6">
        <Avatar className="w-20 h-20">
          <AvatarImage src="/placeholder-user.jpg" alt="@shadcn" />
          <AvatarFallback>CN</AvatarFallback>
        </Avatar>
        <div className="text-center space-y-1">
          <div className="font-semibold">Chandler Bing</div>
          <div className="text-muted-foreground">@chandlerbing</div>
        </div>
        <div className="flex items-center gap-4">
          <div className="flex items-center gap-1 text-muted-foreground">
            <CalendarDaysIcon className="h-4 w-4" />
            <span>Joined June 2023</span>
          </div>
          <div className="flex items-center gap-1 text-muted-foreground">
            <UsersIcon className="h-4 w-4" />
            <span>100 followers</span>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}

function CreatePostCard() {
  const [thread, setThread] = useState('');
  const [privacy, setPrivacy] = useState('public');
  const [file, setFile] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData();
    formData.append('thread', thread);
    formData.append('privacy', privacy);
    if (file) {
      formData.append('file', file);
    }

    try {
      const data = await fetchPost(formData);
      console.log('Post created:', data);
      
      mutate('http://localhost:8080/posts');
      window.location.href = "/";

      setThread('');
      setPrivacy('public');
      setFile(null);
    } catch (error) {
      console.error('Error creating post:', error);
    }
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle>Create a new post</CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        <form onSubmit={handleSubmit}>
          <Textarea
            placeholder="What's on your mind?"
            className="w-full resize-none mb-4"
            rows={3}
            value={thread}
            onChange={(e) => setThread(e.target.value)}
            required
          />
          <div className="flex items-center gap-4 mb-4">
            <Button
              type="button"
              variant="ghost"
              size="icon"
              onClick={() => document.getElementById('file-input').click()}
            >
              <ImageIcon className="h-5 w-5" />
              <span className="sr-only">Add image</span>
            </Button>
            <input
              id="file-input"
              type="file"
              onChange={(e) => setFile(e.target.files[0])}
              className="hidden"
            />
            {file && <span className="text-sm">{file.name}</span>}
          </div>
          <div className="flex items-center gap-4 mb-4">
            {['public', 'private', 'almost-private'].map((option) => (
              <div key={option} className="flex items-center gap-2">
                <input
                  type="radio"
                  id={option}
                  name="privacy"
                  value={option}
                  checked={privacy === option}
                  onChange={(e) => setPrivacy(e.target.value)}
                  className="hidden"
                />
                <Label
                  htmlFor={option}
                  className={`cursor-pointer px-3 py-1 rounded ${privacy === option
                    ? 'bg-primary text-primary-foreground'
                    : 'bg-secondary'
                    }`}
                >
                  {option.charAt(0).toUpperCase() + option.slice(1)}
                </Label>
              </div>
            ))}
          </div>
          <Button type="submit" className="w-full">Post</Button>
        </form>
      </CardContent>
    </Card>
  );
}

function PostCard() {
  return (
    <Card>
      <CardContent className="space-y-4">
        <div className="flex items-center gap-4">
          <Avatar className="w-10 h-10">
            <AvatarImage src="/placeholder-user.jpg" alt="@shadcn" />
            <AvatarFallback>CN</AvatarFallback>
          </Avatar>
          <div className="space-y-1">
            <div className="font-semibold">Chandler Bing</div>
            <div className="text-muted-foreground">@chandlerbing</div>
          </div>
          <div className="ml-auto text-muted-foreground">
            <CalendarDaysIcon className="h-4 w-4" />
            <span>2h</span>
          </div>
        </div>
        <div className="prose prose-sm">
          <p>
            Hey everyone! Just wanted to share a quick update. I&apos;m working on a new project that I&apos;m really excited
            about. It&apos;s going to be a game-changer in the industry. Stay tuned for more details!
          </p>
        </div>
        <Image
          src="/placeholder.svg"
          width={800}
          height={450}
          alt="Project preview"
          className="rounded-lg object-cover"
          style={{ aspectRatio: "800/450", objectFit: "cover" }}
        />
      </CardContent>
    </Card>
  );
}

function TrendingCard() {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Trending</CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="flex items-center gap-4">
          <Avatar className="w-10 h-10">
            <AvatarImage src="/placeholder-user.jpg" alt="@shadcn" />
            <AvatarFallback>CN</AvatarFallback>
          </Avatar>
          <div className="space-y-1">
            <div className="font-semibold">Rachel Green</div>
            <div className="text-muted-foreground">@rachelg</div>
          </div>
          <div className="ml-auto text-muted-foreground">
            <CalendarDaysIcon className="h-4 w-4" />
            <span>3d</span>
          </div>
        </div>
        <div className="prose prose-sm">
          <p>
            Just landed my dream job at Ralph Lauren! So excited for this new chapter. Thanks to everyone for the
            support.
          </p>
        </div>
      </CardContent>
    </Card>
  );
}

function SuggestionsCard() {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Suggestions</CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="flex items-center gap-4">
          <Avatar className="w-10 h-10">
            <AvatarImage src="/placeholder-user.jpg" alt="@shadcn" />
            <AvatarFallback>CN</AvatarFallback>
          </Avatar>
          <div className="space-y-1">
            <div className="font-semibold">Joey Tribbiani</div>
            <div className="text-muted-foreground">@joeyt</div>
          </div>
          <Button variant="outline" size="sm" className="ml-auto">
            Follow
          </Button>
        </div>
        <div className="flex items-center gap-4">
          <Avatar className="w-10 h-10">
            <AvatarImage src="/placeholder-user.jpg" alt="@shadcn" />
            <AvatarFallback>CN</AvatarFallback>
          </Avatar>
          <div className="space-y-1">
            <div className="font-semibold">Phoebe Buffay</div>
            <div className="text-muted-foreground">@phoebeb</div>
          </div>
          <Button variant="outline" size="sm" className="ml-auto">
            Follow
          </Button>
        </div>
      </CardContent>
    </Card>
  );
}

export function HomePage() {
  return (
    <div className="flex flex-col h-screen">
      <Header />

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
          {/* <TrendingCard /> */}
          <SuggestionsCard />
        </div>
      </div>
    </div>
  );
}