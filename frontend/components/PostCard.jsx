"use client";

const { ThumbsUp, ThumbsDown, MessageCircle, ShareIcon, HeartIcon, MessageCircleIcon, ChevronDown, ChevronUp } = require("lucide-react");
const { AvatarFallback, AvatarImage, Avatar } = require("./ui/avatar");
const { CardContent, Card, CardFooter } = require("./ui/card");
import { fetchAllPosts, fetchLike } from '@/app/actions/post';
const { default: Image } = require("next/image");
const { useEffect, useState } = require('react');
const { Textarea } = require("./ui/textarea");
const { Button } = require("./ui/button");
import Link from "next/link";
import CommentCard from './CommentCard';

export default function Posts() {
  const [posts, setPosts] = useState([]);
  const [isMounted, setIsMounted] = useState(false);

  useEffect(() => {
    if (!isMounted) {
      fetchAllPosts(setPosts);
      // console.log("Component mounted");
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
  const [showComments, setShowComments] = useState(false);

  const handleToggleComments = () => {
    setShowComments(!showComments);
  };


  const handleLikeClick = async () => {
    const formData = new FormData();
    formData.append("like-postId", post.PostID);

    try {
      await fetchLike(formData);
      window.location.href = "/";
    } catch (error) {
      console.error("Error liking post:", error);
    }
  };

  return (
    <Card>
      <CardContent className="space-y-4">
        {/* <br /> */}
        <div className="flex items-center gap-4">
          <Avatar className="w-10 h-10">
            <AvatarImage src={post.Author.Avatar || "/placeholder-user.jpg"} alt={post.Author.Username} />
            <AvatarFallback>{post.Author.Username ? post.Author.Username[0].toUpperCase() : 'U'}</AvatarFallback>
          </Avatar>

          <div className="flex items-center gap-2">
            <span className="font-semibold">{post.Author.Firstname + " " + post.Author.Lastname || "Anonymous"}</span>
            <span className="text-muted-foreground">@{post.Author.Username || ""}</span>
            <span className="text-muted-foreground">•</span>
            <span className="text-muted-foreground">{post.Formated_date || "just now"}</span>
          </div>

          {/* <Button  variant="outline" size="sm" className="ml-auto">
            Follow
          </Button> */}
        </div>
        <div className="text-sm grid gap-2 p-4">
          {post.Content || "No content available."}
        </div>
        {post.HasImage && (
          <img
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
        {showComments && <CommentCard postId={post.PostID} commentData={post.Comments} />}
      </CardContent>
    </Card>
  );
}
