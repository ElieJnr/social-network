"use client";

import { useState } from "react";

const { Card, CardHeader, CardTitle, CardContent } = require("./ui/card");
import { fetchCreatePost } from '@/app/actions/post';
import { useToast } from "@/components/ui/use-toast";
const { ImageIcon } = require("lucide-react");
const { Textarea } = require("./ui/textarea");
const { Button } = require("./ui/button");
const { Input } = require("./ui/input");
const { Label } = require("./ui/label");
import { mutate } from "swr";
import useStore from "@/app/store/useStore";


export default function CreatePostCard() {
  const { toast } = useToast();
  const [thread, setThread] = useState('');
  const [privacy, setPrivacy] = useState('public');
  const [file, setFile] = useState(null);
  const { user, setUser } = useStore();

  // console.log("user user: ", user);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData();
    formData.append('thread', thread);
    formData.append('privacy', privacy);
    if (file) {
      formData.append('file', file);
    }

    if (!checkPost(thread, privacy, file, toast)) {
      return;
    }

    try {
      await fetchCreatePost(formData);

      mutate('http://localhost:8080/posts');

      setThread('');
      setPrivacy('public');
      setFile(null);
    } catch (error) {
      toast({
        title: "Error creating post",
        description: error,
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
          <div className="flex items-center gap-4 mb-4">
            {['public', 'private', 'almost-private'].map((option) => (
              <div key={option} className="flex items-center gap-2">
                <Input
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

function checkPost(thread, privacy, file, toast) {
  if (thread.length === 0 || thread.length > 500) {
    toast({
      title: "Error creating post",
      description: "Thread must not be empty and must not exceed 500 characters.",
    });
    return false; 
  }

  
  const validPrivacyOptions = ["public", "private", "almost-private"];
  if (!validPrivacyOptions.includes(privacy)) {
    toast({
      title: "Error creating post",
      description: "Privacy must be one of the following: public, private, or almost-private.",
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
