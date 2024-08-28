"use client";

const { CalendarDaysIcon, UsersIcon } = require("lucide-react");
const { Card, CardContent } = require("./ui/card");
const { Avatar, AvatarImage, AvatarFallback } = require("./ui/avatar");

export default function ProfileCard() {
  console.log( userConnect());
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