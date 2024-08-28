"use client";

const { CalendarDaysIcon } = require("lucide-react");
const { default: Image } = require("next/image");
const { AvatarFallback, AvatarImage, Avatar } = require("./ui/avatar");
const { CardContent, Card } = require("./ui/card");

export default function PostCard() {
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
  