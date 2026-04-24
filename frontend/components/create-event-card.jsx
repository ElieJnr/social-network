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
import { ValidateInput } from "@/app/actions/input"; // Assume ValidateInput exists

export function CreateEventCard({ id, socket }) {
  const { toast } = useToast();
  const [eventName, setEventName] = useState("");
  const [eventDescription, setEventDescription] = useState("");
  const [eventDate, setEventDate] = useState("");

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

  const handleCreateEvent = async () => {
    // Validate each field before proceeding
    if (!validateField("eventName", eventName)) return;
    if (!validateField("eventDescription", eventDescription)) return;
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

    const message = {
      Type: "notifications",
      SubType: "event",
      GroupeId: id,
      Content: eventData.title,
    };
    socketSend(socket, message);
  };

  return (
    <Card className="w-full">
      <CardHeader>
        <CardTitle>Créer un événement</CardTitle>
        <CardDescription>
          Remplissez les informations pour créer un nouvel événement.
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="space-y-2">
          <Label htmlFor="event-name">Nom de l&apos;événement</Label>
          <Input
            id="event-name"
            placeholder="Entrez le nom de l'événement"
            value={eventName}
            onChange={(e) => setEventName(e.target.value)}
          />
        </div>
        <div className="space-y-2">
          <Label htmlFor="event-description">Description</Label>
          <Textarea
            id="event-description"
            placeholder="Décrivez l'événement"
            value={eventDescription}
            onChange={(e) => setEventDescription(e.target.value)}
          />
        </div>
        <div className="space-y-2">
          <Label htmlFor="event-date">Date</Label>
          <Input
            type="date"
            id="event-date"
            value={eventDate}
            onChange={(e) => setEventDate(e.target.value)}
            min={new Date().toISOString().split("T")[0]} // Disable past dates
          />
        </div>
      </CardContent>
      <CardFooter className="flex justify-end">
        <Button onClick={handleCreateEvent}>Créer l&apos;événement</Button>
      </CardFooter>
    </Card>
  );
}
