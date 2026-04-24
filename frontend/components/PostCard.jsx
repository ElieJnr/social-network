"use client";

import { useToast } from "@/components/ui/use-toast";
import { fetchLike } from '@/app/actions/post';
import CommentCard from './CommentCard';
import useSWR, { mutate } from 'swr';
import Image from "next/image";
const {HeartIcon, MessageCircleIcon} = require("lucide-react");
const { AvatarFallback, AvatarImage, Avatar } = require("./ui/avatar");
const { CardContent, Card, CardFooter } = require("./ui/card");
const { Button } = require("./ui/button");
const { useState } = require('react');

const fetcher = (url) => fetch(url, { credentials: 'include' }).then((res) => res.json());

export default function Posts() {

  const { data: posts, mutate, isValidating } = useSWR('http://localhost:8080/posts', fetcher);
  
  return (
    <div className="space-y-4">
      {isValidating && !posts ? (
        <>
          <SkeletonPostCard />
          <SkeletonPostCard />
          <SkeletonPostCard />
        </>
      ) : (
        posts && posts.length > 0 ? (
          posts.map(post =>
            post.Can_see ? <PostCard key={post.PostID} post={post} /> : null
          )
        ) : null
      )}
    </div>
  );
}

export function PostCard({ post }) {
  const [showComments, setShowComments] = useState(false);
  const { toast } = useToast();

  const handleToggleComments = () => {
    setShowComments(!showComments);
  };

  const handleLikeClick = async () => {
    const formData = new FormData();
    formData.append("like-postId", post.PostID);

    try {
      await fetchLike(formData);
      mutate('http://localhost:8080/posts');
    } catch (error) {
      toast({
        title: "Error liking post",
        description: error,
      });
    }
  };


  return (
    <Card >
      <CardContent className="space-y-4 ">
        <div className="flex items-center gap-4 mt-2">
          <Avatar className="w-10 h-10">
            <AvatarImage src={post?.Author.Avatar === "" ? "/placeholder-user.jpg": `/uploads/${post?.Author.Avatar}`} alt={post.Author.Username} />
            <AvatarFallback>{post.Author.Username ? post.Author.Username[0].toUpperCase() : 'U'}</AvatarFallback>
          </Avatar>

          <div className="flex items-center gap-2">
            <span className="text-lg font-semibold">{post.Author.Firstname + " " + post.Author.Lastname || "Anonymous"}</span>
            <span className="text-muted-foreground text-lg">•</span>
            <span className="text-muted-foreground">{post.Formated_date || "just now"}</span>
          </div>

          {/* {post.IsFollower ? '' : <Button variant="outline" size="sm" className="ml-auto">
            Follow
          </Button>} */}
        </div>

        <div className="text-lg grid gap-2 p-4">
          {post.Content || "No content available."}
        </div>
        {post.HasImage && (
          <Image
            src={`/uploads/${post.Image_url}`}
            width={800}
            height={450}
            alt="Project preview"
            className="rounded-lg object-cover"
            style={{ aspectRatio: "800/450", objectFit: "cover" }}
          />
        )}
        <CardFooter className="grid gap-2 p-4">
          <div className="flex items-center gap-2">
            <div className="flex items-center gap-1">
              <Button variant="ghost" size="icon" onClick={handleLikeClick}>
                <HeartIcon
                  className={`h-5 w-5 ${post.Like_status ? 'text-red-500' : ''}`}
                />
              </Button>
              <span>{post.Like_nbr || 0}</span>
            </div>
            <div className="flex items-center gap-1">
              <Button variant="ghost" size="icon" onClick={handleToggleComments}>
                <MessageCircleIcon className="h-5 w-5" />
              </Button>
              <span>{post.Comments_nbr || 0}</span>
            </div>
          </div>
        </CardFooter>
        {showComments && <CommentCard postId={post.PostID} />}
      </CardContent>
    </Card>

  );
}

export function SkeletonPostCard() {
  return (
    <Card className="animate-pulse">
      <CardContent className="space-y-4">
        <div className="flex items-center gap-4 mt-2">
          <div className="w-10 h-10 bg-gray-300 rounded-full"></div>
          <div className="flex-1 space-y-2">
            <div className="h-4 bg-gray-300 rounded w-3/4"></div>
            <div className="h-4 bg-gray-300 rounded w-1/4"></div>
          </div>
        </div>

        <div className="space-y-2">
          <div className="h-4 bg-gray-300 rounded w-full"></div>
          <div className="h-4 bg-gray-300 rounded w-5/6"></div>
        </div>

        <div className="h-48 bg-gray-300 rounded"></div>

        <CardFooter className="grid gap-2 p-4">
          <div className="flex items-center gap-2">
            <div className="flex items-center gap-1">
              <div className="w-5 h-5 bg-gray-300 rounded-full"></div>
              <div className="h-4 bg-gray-300 rounded w-6"></div>
            </div>
            <div className="flex items-center gap-1">
              <div className="w-5 h-5 bg-gray-300 rounded-full"></div>
              <div className="h-4 bg-gray-300 rounded w-6"></div>
            </div>
          </div>
        </CardFooter>
      </CardContent>
    </Card>
  );
}
