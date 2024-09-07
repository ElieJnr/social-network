import React, { useState } from "react";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import { useToast } from "./ui/use-toast";
import { CreateEvent } from "@/app/actions/events";
import { mutate } from "swr";
import { domain } from "@/app";
import { socketSend } from "@/app/actions/message";
import { Popover, PopoverContent, PopoverTrigger } from "radix-ui";
import { Calendar } from "lucide-react";

export function CreateEventCard({ id, socket }) {
  const { toast } = useToast();
  const [eventName, setEventName] = useState("");
  const [eventDescription, setEventDescription] = useState("");
  const [eventDate, setEventDate] = useState("");

  const handleCreateEvent = async () => {
    if (!eventName.trim()) {
      toast({
        title: "Champ manquant",
        description: "Veuillez entrer un nom pour l'événement.",
      });
      return;
    }
    if (!eventDescription.trim()) {
      toast({
        title: "Champ manquant",
        description: "Veuillez entrer une description pour l'événement.",
      });
      return;
    }
    if (!eventDate) {
      toast({
        title: "Champ manquant",
        description: "Veuillez sélectionner une date pour l'événement.",
      });
      return;
    }

    const eventData = {
      groupid: id,
      title: eventName,
      description: eventDescription,
      date: new Date(eventDate).toISOString(),
    };
    try {
      const response = await CreateEvent(eventData);
      if (!response.ok) {
        toast({
          title: "Erreur",
          description: "Erreur lors de la création de l'événement.",
        });
        return;
      }

      // Mutate the correct SWR key after the event is created
      mutate(`${domain}/group/getEvents?groupId=${id}`);

      // Clear the form after successful creation
      setEventName("");
      setEventDescription("");
      setEventDate("");

      toast({
        title: "Événement créé",
        description: "Votre événement a été créé avec succès.",
      });
    } catch (error) {
      toast({
        title: "Erreur",
        description: "Une erreur s'est produite. Veuillez réessayer.",
      });
    }

    const Message = {
      Type: "notifications",
      SubType: "event",
      GroupeId: id,
      Content:eventData.title,
    }
    socketSend(socket, Message)
  };

  return (
    (<Card className="w-full">
      <CardHeader>
        <CardTitle>Créer un événement</CardTitle>
        <CardDescription>Remplissez les informations pour créer un nouvel événement.</CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="space-y-2">
          <Label htmlFor="event-name">Nom de l&apos;événement</Label>
          <Input id="event-name" placeholder="Entrez le nom de l&apos;événement" />
        </div>
        <div className="space-y-2">
          <Label htmlFor="event-description">Description</Label>
          <Textarea id="event-description" placeholder="Décrivez l&apos;événement" />
        </div>
        <div className="flex items-center justify-between">
          <div className="space-y-2">
            <Label htmlFor="event-date-time">Date et heure</Label>
            <Popover>
              <PopoverTrigger asChild>
                <Button variant="outline" className="flex items-center justify-between w-full">
                  <div className="space-y-1">
                    <p className="text-sm font-medium text-muted-foreground">Vendredi 15 septembre, 19h00</p>
                  </div>
                  <CalendarIcon className="h-6 w-6 text-muted-foreground" />
                </Button>
              </PopoverTrigger>
              <PopoverContent className="p-0">
                <Calendar />
              </PopoverContent>
            </Popover>
          </div>
        </div>
      </CardContent>
      <CardFooter className="flex justify-end">
        <Button>Créer l&apos;événement</Button>
      </CardFooter>
    </Card>)
  );
}

function CalendarIcon(props) {
  return (
    (<svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round">
      <path d="M8 2v4" />
      <path d="M16 2v4" />
      <rect width="18" height="18" x="3" y="4" rx="2" />
      <path d="M3 10h18" />
    </svg>)
  );
}