"use client";

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";
import { CalendarDaysIcon } from "lucide-react";

export default function TrendingCard() {
    return (
      <Card>
        <CardHeader>
          <CardTitle>Trending</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="flex items-center gap-4">
            <Avatar className="w-10 h-10">
              <AvatarImage src="/placeholder-user.jpg" alt="@shadcn" />
              <AvatarFallback>CN</AvatarFallback>
            </Avatar>
            <div className="space-y-1">
              <div className="font-semibold">Rachel Green</div>
              <div className="text-muted-foreground">@rachelg</div>
            </div>
            <div className="ml-auto text-muted-foreground">
              <CalendarDaysIcon className="h-4 w-4" />
              <span>3d</span>
            </div>
          </div>
          <div className="prose prose-sm">
            <p>
              Just landed my dream job at Ralph Lauren! So excited for this new chapter. Thanks to everyone for the
              support.
            </p>
          </div>
        </CardContent>
      </Card>
    );
}
  