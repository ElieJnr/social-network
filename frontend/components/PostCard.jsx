"use client";

const { ThumbsUp, ThumbsDown, MessageCircle, ShareIcon, HeartIcon, MessageCircleIcon, ChevronDown, ChevronUp } = require("lucide-react");
const { AvatarFallback, AvatarImage, Avatar } = require("./ui/avatar");
const { CardContent, Card, CardFooter } = require("./ui/card");
import { fetchAllPosts } from '@/app/actions/post';
const { default: Image } = require("next/image");
const { useEffect, useState } = require('react');
const { Textarea } = require("./ui/textarea");
const { Button } = require("./ui/button");
import Link from "next/link";

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

export function CommentCard() {
  return (
    <div className="mx-auto px-4 md:px-6">
      <div className="gap-4">
        <hr className="flex-grow" />
        <br />
      </div>

      <div className="grid gap-6">
        <div className="flex items-start gap-4">
          <Avatar className="w-10 h-10 border">
            <AvatarImage src="/placeholder-user.jpg" alt="@shadcn" />
            <AvatarFallback>AC</AvatarFallback>
          </Avatar>
          <div className="grid gap-1.5">
            <div className="flex items-center gap-2">
              <div className="font-semibold">Sarah Johnson</div>
              <div className="text-xs text-muted-foreground">2 days ago</div>
            </div>
            <div className="text-sm leading-loose text-muted-foreground">
              <p>I love the new design updates! The layout is so clean and modern. Great job on the UI.</p>
            </div>
          </div>
        </div>
        <div className="flex items-start gap-4">
          <Avatar className="w-10 h-10 border">
            <AvatarImage src="/placeholder-user.jpg" alt="@shadcn" />
            <AvatarFallback>AC</AvatarFallback>
          </Avatar>
          <div className="grid gap-1.5">
            <div className="flex items-center gap-2">
              <div className="font-semibold">Alex Smith</div>
              <div className="text-xs text-muted-foreground">3 weeks ago</div>
            </div>
            <div className="text-sm leading-loose text-muted-foreground">
              <p>The new features are really impressive. I can't wait to try them out. Keep up the great work!</p>
            </div>
          </div>
        </div>
        <div className="flex items-start gap-4">
          <Avatar className="w-10 h-10 border">
            <AvatarImage src="/placeholder-user.jpg" alt="@shadcn" />
            <AvatarFallback>AC</AvatarFallback>
          </Avatar>
          <div className="grid gap-1.5">
            <div className="flex items-center gap-2">
              <div className="font-semibold">Emily Parker</div>
              <div className="text-xs text-muted-foreground">2 days ago</div>
            </div>
            <div className="text-sm leading-loose text-muted-foreground">
              <p>The new design is really sleek and modern. I love the attention to detail. Great job!</p>
            </div>
          </div>
        </div>
      </div>
      <div className="grid gap-4">
        <h2 className="text-xl font-semibold">Add a comment</h2>
        <div className="w-full">
          <div className="flex items-start gap-4">
            <div className="flex-1">
              <Textarea
                placeholder="Write a new comment..."
                className="resize-none p-4 rounded-lg border border-neutral-300 focus:border-primary focus:ring-primary"
              />
              <div className="flex justify-between mt-2">
                <div className="flex items-center">
                  <button className="bg-primary text-primary-foreground rounded-lg px-4 py-2 hover:bg-primary/90 transition-colors">
                    <ImageIcon className="w-5 h-5" />
                    <span className="sr-only">Upload image</span>
                  </button>
                  <Button className="ml-2">Comment</Button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

function ImageIcon(props) {
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
      strokeLinejoin="round"
    >
      <rect width="18" height="18" x="3" y="3" rx="2" ry="2" />
      <circle cx="9" cy="9" r="2" />
      <path d="m21 15-3.086-3.086a2 2 0 0 0-2.828 0L6 21" />
    </svg>
  )
}

function PostCard({ post }) {
  const [showComments, setShowComments] = useState(false);

  const handleToggleComments = () => {
    setShowComments(!showComments);
  };

  return (
    <Card>
      <CardContent className="space-y-4">
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
              <Button variant="ghost" size="icon">
                <HeartIcon className="h-5 w-5" />
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
        {showComments && <CommentCard />}
      </CardContent>
    </Card>
  );
}
