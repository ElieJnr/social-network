
import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from "@/components/ui/card"
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
const { ImageIcon } = require("lucide-react");
import { Label } from "@/components/ui/label"
import { Input } from "@/components/ui/input"

export function CreateGroupCard() {
  return (
    <Card className="w-full max-w-md">
      <CardHeader>
      <CardTitle>Create a New Group</CardTitle>
      <CardDescription>Fill out the form to create a new group.</CardDescription>
      </CardHeader>
      <CardContent className="grid gap-4">
      <div className="grid gap-2">
        <Label htmlFor="name">Group Name</Label>
        <div className="flex items-center">
        <Input id="name" placeholder="Enter group name" />
        <ImageIcon className="h-5 w-5 ml-2" />
        </div>
      </div>
      
      <div className="grid gap-2">
        <Label htmlFor="description">Description</Label>
        <Textarea id="description" placeholder="Enter group description" />
      </div>
      </CardContent>
      <CardFooter className="flex justify-end">
      <Button className="w-full" type="submit">Create Group</Button>
      </CardFooter>
    </Card>
    );
}
