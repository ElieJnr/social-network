"use client";

const { ThumbsUp, ThumbsDown, MessageCircle } = require("lucide-react");
const { AvatarFallback, AvatarImage, Avatar } = require("./ui/avatar");
const { CardContent, Card } = require("./ui/card");
import { fetchAllPosts } from '@/app/actions/post';
const { default: Image } = require("next/image");
const { useEffect, useState } = require('react');

export default function Posts() {
  const [posts, setPosts] = useState([]);
  const [isMounted, setIsMounted] = useState(false);

  useEffect(() => {
    if (!isMounted) {
      fetchAllPosts(setPosts);
      console.log("Component mounted");
      setIsMounted(true);
    }
  }, [isMounted]);

  return (
    <div className="space-y-4">
      {posts && posts.length > 0 ? (
        posts.map(post =>
          post.Can_see ? <PostCard key={post.PostID} post={post} /> : null
        )
      ) : null}
    </div>
  );
}

function PostCard({ post }) {
  return (
    <Card>
      <CardContent className="space-y-4">
        <div className="flex items-center gap-4">
          <Avatar className="w-10 h-10">
            <AvatarImage src={post.Author.Avatar || "/placeholder-user.jpg"} alt={post.Author.Username} />
            <AvatarFallback>{post.Author.Username ? post.Author.Username[0].toUpperCase() : 'U'}</AvatarFallback>
          </Avatar>

          <div className="flex items-center gap-2">
            <span className="font-semibold">{post.Author.Firstname+" "+post.Author.Lastname || "Anonymous"}</span>
            <span className="text-muted-foreground">@{post.Author.Username || ""}</span>
            <span className="text-muted-foreground">•</span>
            <span className="text-muted-foreground">{post.Formated_date || "just now"}</span>
          </div>
        </div>
        <div className="prose prose-sm">
          <p>{post.Content || "No content available."}</p>
        </div>
        {post.image && (
          <img
            src={post.Image_url }
            width={800}
            height={450}
            alt="Project preview"
            className="rounded-lg object-cover"
            style={{ aspectRatio: "800/450", objectFit: "cover" }}
          />
        )}
        <div className="flex items-center justify-around mt-4">
          <div className="flex items-center space-x-2">
            <ThumbsUp className="h-5 w-5" />
            <span>{post.Like_nbr  || 0}</span>
          </div>
          <div className="flex items-center space-x-2">
            <ThumbsDown className="h-5 w-5" />
            <span>{post.Dislike_nbr || 0}</span>
          </div>
          <div className="flex items-center space-x-2">
            <MessageCircle className="h-5 w-5" />
            <span>{post.comments || 0}</span>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}


