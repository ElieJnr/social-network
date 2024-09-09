import { useState } from "react";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import { ImageIcon } from "lucide-react"; // Import corrected
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { useToast } from "@/components/ui/use-toast"; // For toasts
import { ValidateInput } from "@/app/actions/input"; // Assuming ValidateInput exists

export function CreateGroupCard() {
  const { toast } = useToast();
  const [groupName, setGroupName] = useState("");
  const [description, setDescription] = useState("");

  // Function to validate individual fields
  const validateField = (field, value) => {
    const { isValid, message } = ValidateInput(value);
    if (!isValid) {
      toast({
        title: "Validation Error",
        description: message,
      });
      return false;
    }
    return true;
  };

  const handleCreateGroup = async () => {
    // Validate group name and description
    if (!validateField("groupName", groupName)) return;
    if (!validateField("description", description)) return;

    const newGroup = {
      title: groupName,
      description: description,
    };

    // Appel à l'API pour créer le groupe
    try {
      const response = await fetch("http://localhost:8080/group/createGroup", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newGroup),
      });

      if (response.ok) {
        toast({
          title: "Group Created",
          description: "The group has been created successfully.",
        });
        window.location.reload();
      } else {
        toast({
          title: "Failed to Create Group",
          description: "An error occurred while creating the group.",
        });
      }
    } catch (error) {
      toast({
        title: "Error",
        description:
          "An error occurred while creating the group. Please try again.",
      });
      console.error("An error occurred while creating the group:", error);
    }
  };

  return (
    <Card className="w-full max-w-md">
      <CardHeader>
        <CardTitle>Create a New Group</CardTitle>
        <CardDescription>
          Fill out the form to create a new group.
        </CardDescription>
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
