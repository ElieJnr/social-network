import { useState } from 'react';
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "@/components/ui/card"
import { Button } from "@/components/ui/button";
import { CreateEventCard } from './create-event-card';
import { useWebSocket } from '@/app/actions/message';
import { Label } from "@/components/ui/label"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
const { ImageIcon } = require("lucide-react");
import { Popover, PopoverTrigger, PopoverContent } from "@/components/ui/popover"
import { EventsCards } from './events-cards';
import PostCard from "@/components/PostCard";
import { ClickMessageApp } from './ui/message';


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
        <ClickMessageApp socket={socket} />
      </div>
    </aside>
  );
}

export function GroupHomePage() {
  const [activeComponent, setActiveComponent] = useState('post');
  const socket = useWebSocket('ws://localhost:8080/ws');

  return (
    <div className="flex flex-col h-screen">
      <div className="flex-1 grid grid-cols-[1fr_2fr_1.5fr] gap-6 p-6">
        <div className="space-y-6">
          <EventsCards />
          <SuggestionsGroupCard />
        </div>
        <div className="space-y-6">
          {activeComponent === 'post' ? <CreatePostGroupCard /> : <CreateEventCard />}
          <PostCard />
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

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData();
    formData.append('thread', thread);
    if (file) {
      formData.append('file', file);
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
