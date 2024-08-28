"use client";

const { Avatar, AvatarImage, AvatarFallback } = require("./ui/avatar");
const { Button } = require("./ui/button");
const { Card, CardHeader, CardTitle, CardContent } = require("./ui/card");

export default function SuggestionsCard() {
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
              Follow
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
              Follow
            </Button>
          </div>
        </CardContent>
      </Card>
    );
  }