import { useState } from "react";
import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from "@/components/ui/card";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
const { ImageIcon } = require("lucide-react");
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";

export function CreateGroupCard() {
  const [groupName, setGroupName] = useState("");
  const [description, setDescription] = useState("");
  const handleCreateGroup = async () => {
    // Exemple de logique pour créer le groupe
    const newGroup = {
      title: groupName,
      description: description,
    };

    // Appel à l'API pour créer le groupe
    try {
      const response = await fetch("http://localhost:8080/group/createGroup", {
        method: "POST",
        credentials:"include",
        
        body: JSON.stringify(newGroup),
      });

      if (response.ok) {
        console.log("Group created successfully");
        window.location.reload()
      } else {
        console.error("Failed to create group");
      }
    } catch (error) {
      console.error("An error occurred while creating the group:", error);
    }
  };

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
            <Input 
              id="name" 
              placeholder="Enter group name" 
              value={groupName} 
              onChange={(e) => setGroupName(e.target.value)} 
            />
            <ImageIcon className="h-5 w-5 ml-2" />
          </div>
        </div>

        <div className="grid gap-2">
          <Label htmlFor="description">Description</Label>
          <Textarea 
            id="description" 
            placeholder="Enter group description" 
            value={description} 
            onChange={(e) => setDescription(e.target.value)} 
          />
        </div>
      </CardContent>
      <CardFooter className="flex justify-end">
        <Button className="w-full" type="button" onClick={handleCreateGroup}>
          Create Group
        </Button>
      </CardFooter>
    </Card>
  );
}
