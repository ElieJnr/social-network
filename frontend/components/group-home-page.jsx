"use client"
import { CardContent, Card, CardFooter , CardHeader, CardTitle } from "@/components/ui/card"
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import { fetchGroupCreatePost } from "@/app/actions/post";
import { CreateEventCard } from './create-event-card';
import { useWebSocket } from '@/app/actions/message';
import { useToast } from "@/components/ui/use-toast";
import { Textarea } from "@/components/ui/textarea"
import { Button } from "@/components/ui/button";
const { ImageIcon } = require("lucide-react");
import { Input } from "@/components/ui/input";
import { useParams } from "next/navigation";
import { fetchLike } from '@/app/actions/post';
import CommentCard from './CommentCard';
import useSWR, { mutate } from 'swr';
import GroupChat from "./MessageComponent/MessageApp";
const {HeartIcon, MessageCircleIcon} = require("lucide-react");
const { useState } = require('react');
// import PostCard from "@/components/PostCard";
import EventList from './EventList';

const fetcher = (url) => fetch(url, { credentials: 'include' }).then((res) => res.json());

function Sidebar({ setActiveComponent, socket }) {
  return (
    <aside className="w-full space-y-4">
      <Card className="flex flex-col w-full">
        <CardHeader>
          <div className="flex justify-between items-center">
            <p className="text-lg text-muted-foreground">Ceci est une description de groupe.</p>
          </div>
        </CardHeader>
        <CardContent className="mt-2">
          <div className="flex justify-between mt-4">
            <div className="text-center">
              <p className="text-lg font-bold">989 k</p>
              <p className="text-sm text-muted-foreground">Membres</p>
            </div>
            <div className="text-center">
              <p className="text-lg font-bold">40</p>
              <div className="flex items-center justify-center">
                <span className="h-2 w-2 bg-green-500 rounded-full mr-1" />
                <p className="text-sm text-muted-foreground">En ligne</p>
              </div>
            </div>
            <div className="text-center">
              <p className="text-lg font-bold">Premiers 1 %</p>
              <p className="text-sm text-muted-foreground">Classer par popularité</p>
            </div>
          </div>
        </CardContent>
      </Card>

      <Card className="w-full max-w">
        <div className="w-full flex space-x-3">
          <Button className="flex-1" onClick={() => setActiveComponent('event')}>Create Event</Button>
          <Button className="flex-1" onClick={() => setActiveComponent('post')}>Create Post</Button>
        </div>
      </Card>

      <div className="space-y-6">
        <GroupChat socket={socket} />
      </div>
    </aside>
  );
}

export function GroupHomePage({id}) {
  const [activeComponent, setActiveComponent] = useState('post');
  const socket = useWebSocket('ws://localhost:8080/ws');

  return (
    <div className="flex flex-col h-screen">
      <div className="flex-1 grid grid-cols-[1fr_2fr_1.5fr] gap-6 p-6">
        <div className="space-y-6">
          {/* <EventsCards /> */}
          <SuggestionsGroupCard />
          <EventList id={id}/>
        </div>
        <div className="space-y-6">
          {activeComponent === 'post' ? <CreatePostGroupCard /> : <CreateEventCard id={id} />}
          <PostGroupCard />
          {/* <PostCard /> */}
        </div>
        <div className="space-y-6">
          <Sidebar setActiveComponent={setActiveComponent} />
        </div>
      </div>
    </div>
  );
}

export function SuggestionsGroupCard() {
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
            Invite
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
            Invite
          </Button>
        </div>
      </CardContent>
    </Card>
  );
}

export function CreatePostGroupCard() {
  const [thread, setThread] = useState('');
  const [file, setFile] = useState(null);
  const { toast } = useToast();
  const { id } = useParams()

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData();
    formData.append('thread', thread);
    if (file) {
      formData.append('file', file);
    }

    formData.append('groupId', id);

    if (!checkPosts(thread, file, toast)) {
      return;
    }

    try {
      await fetchGroupCreatePost(formData);
      mutate(`http://localhost:8080/group/posts?groupId=${id}`);
      setThread('');
      setFile(null);
      toast({
        title: "Post created successfully",
        description: "Your post has been created and shared.",
      });
    }
    catch (error) {
      toast({
        title: "Error creating post",
        description: error.message || "An unexpected error occurred",
      });
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
            <Input
              id="file-input"
              type="file"
              onChange={(e) => setFile(e.target.files[0])}
              className="hidden"
            />
            {file && <span className="text-sm">{file.name}</span>}
          </div>
          <Button type="submit" className="w-full">Post</Button>
        </form>
      </CardContent>
    </Card>
  );
}

export function PostGroupCard() {
  const { id } = useParams();
  const { data: posts, mutate, isValidating } = useSWR(`http://localhost:8080/group/posts?groupId=${id}`, fetcher);

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

function PostCard({ post }) {
  const [showComments, setShowComments] = useState(false);
  const { toast } = useToast();
  const { id } = useParams()

  const handleToggleComments = () => {
    setShowComments(!showComments);
  };

  const handleLikeClick = async () => {
    const formData = new FormData();
    formData.append("like-postId", post.PostID);

    try {
      await fetchLike(formData);
      mutate(`http://localhost:8080/group/posts?groupId=${id}`);
    } catch (error) {
      toast({
        title: "Error liking post",
        description: error,
      });
    }
  };

  return (
    <Card>
      <CardContent className="space-y-4">
        <div className="flex items-center gap-4 mt-2">
          <Avatar className="w-10 h-10">
            <AvatarImage src={post?.Author.Avatar || "/placeholder-user.jpg"} alt={post?.Author.Username} />
            <AvatarFallback>{post?.Author.Username ? post?.Author.Username[0].toUpperCase() : 'U'}</AvatarFallback>
          </Avatar>

          <div className="flex items-center gap-2">
            <span className="text-lg font-semibold">{post?.Author.Firstname + " " + post?.Author.Lastname || "Anonymous"}</span>
            <span className="text-muted-foreground">@{post?.Author.Username || ""}</span>
            <span className="text-muted-foreground text-lg">•</span>
            <span className="text-muted-foreground">{post?.Formated_date || "just now"}</span>
          </div>

          {post?.IsFollower ? '' : <Button variant="outline" size="sm" className="ml-auto">
            Follow
          </Button>}
        </div>

        <div className="text-lg grid gap-2 p-4">
          {post?.Content || "No content available."}
        </div>
        {post?.HasImage && (
          <img
            src={`/uploads/${post?.Image_url}`}
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
                  className={`h-5 w-5 ${post?.Like_status ? 'text-red-500' : ''}`}
                />
              </Button>
              <span>{post?.Like_nbr || 0}</span>
            </div>
            <div className="flex items-center gap-1">
              <Button variant="ghost" size="icon" onClick={handleToggleComments}>
                <MessageCircleIcon className="h-5 w-5" />
              </Button>
              <span>{post?.Comments_nbr || 0}</span>
            </div>
          </div>
        </CardFooter>
        {showComments && <CommentCard postId={post?.PostID} />}
      </CardContent>
    </Card>
  );
}

function SkeletonPostCard() {
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

function checkPosts(thread, file, toast) {
  if (thread.length === 0 || thread.length > 500) {
    toast({
      title: "Error creating post",
      description: "Thread must not be empty and must not exceed 500 characters.",
    });
    return false;
  }


  if (file) {
    if (!file.type.includes('image')) {
      toast({
        title: "Error creating post",
        description: "File must be an image.",
      });
      return false;
    }

    const validExtensions = ["jpg", "jpeg", "png", "gif"];
    const fileExtension = file.name.split('.').pop().toLowerCase();
    if (!validExtensions.includes(fileExtension)) {
      toast({
        title: "Error creating post",
        description: "File must be in jpg, png, or gif format.",
      });
      return false;
    }

    const maxSizeMB = 5;
    const maxSizeBytes = maxSizeMB * 1024 * 1024;
    if (file.size > maxSizeBytes) {
      toast({
        title: "Error creating post",
        description: `File size must not exceed ${maxSizeMB} MB.`,
      });
      return false;
    }
  }

  return true;
}