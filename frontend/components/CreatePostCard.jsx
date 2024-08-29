"use client";

import { useState } from "react";

const { ImageIcon } = require("lucide-react");
const { Button } = require("./ui/button");
const { Card, CardHeader, CardTitle, CardContent } = require("./ui/card");
const { Textarea } = require("./ui/textarea");
const { Input } = require("./ui/input");
const { Label } = require("./ui/label");
import { fetchCreatePost } from '@/app/actions/post';
import { mutate } from "swr";

export default function CreatePostCard() {
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
        const data = await fetchCreatePost(formData);
        console.log('Post created:', data);
        

        // mutate('http://localhost:8080/posts');
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
  