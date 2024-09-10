import { useState } from "react";
import useSWR from "swr";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { domain } from "@/app";
import Skeleton from "react-loading-skeleton";
import "react-loading-skeleton/dist/skeleton.css";
import { fetchEvents } from "@/app/actions/events";

const fetcher = (url) =>
  fetch(url, { credentials: "include", cache: "no-store" }).then((res) =>
    res.json()
  );

export default function EventList({ id }) {
  const [showGoingEvents, setShowGoingEvents] = useState(false);

  const {
    data: events,
    error,
    mutate,
  } = useSWR(`${domain}/group/getEvents?groupId=${id}`, () => fetchEvents(id));

  const { data: responses, mutate: mutateResponses } = useSWR(
    `${domain}/group/GetrespondEvent?groupId=${id}`,
    fetcher
  );

  const respondToEvent = async (eventid, response) => {
    try {
      // Optimistically update the UI
      mutateResponses(
        (currentResponses) =>
          currentResponses?.map((res) =>
            res.eventid === eventid ? { ...res, response } : res
          ),
        false
      );

      const res = await fetch(`${domain}/group/respondEvent`, {
        method: "POST",
        credentials: "include",
        cache: "no-store",
        body: JSON.stringify({ eventid, response }),
      });

      if (!res.ok) throw new Error("Error responding to event");

      // Revalidate the response list to ensure data consistency
      mutateResponses();
      // Optionally revalidate the events list as well
      mutate();
    } catch (error) {
      console.error("Error responding to event:", error);
      // Revert optimistic update in case of an error
      mutateResponses();
    }
  };

  if (error) {remainingEvents
    return <p>Error loading events. Please try again.</p>;
  }

  if (!events) {
    return (
      <div>
        {Array(3)
          .fill()
          .map((_, index) => (
            <Card key={index} className="w-full max-w-md mx-auto mb-4">
              <CardHeader>
                <CardTitle>
                  <Skeleton width={200} />
                </CardTitle>
                <CardDescription>
                  <Skeleton count={2} />
                </CardDescription>
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="flex items-center justify-between">
                  <Skeleton width={150} />
                </div>
              </CardContent>
              <CardFooter className="flex justify-end">
                <Skeleton width={100} height={36} className="mr-2" />
                <Skeleton width={100} height={36} />
              </CardFooter>
            </Card>
          ))}
      </div>
    );
  }

  const eventsArray = Array.isArray(events.Event)
    ? events.Event
    : [events.Event];

  if (eventsArray.length === 0 || !events.Isevent) {
    return <p className="text-center">No events available.</p>;
  }

  const filteredEvents = eventsArray.filter((event) => {
    const response = responses?.find(
      (res) => String(res.eventid) === String(event.id)
    );
    return !response || response.response !== "not_going";
  });

  const goingEvents = filteredEvents.filter((event) => {
    const response = responses?.find(
      (res) => String(res.eventid) === String(event.id)
    );
    return response && response.response === "going";
  });

  const remainingEvents = filteredEvents.filter((event) => {
    const response = responses?.find(
      (res) => String(res.eventid) === String(event.id)
    );
    return !response || response.response !== "going";
  });

  return (
    <div>
      <Button
        variant="outline"
        onClick={() => setShowGoingEvents(!showGoingEvents)}
        className="mt-8 mx-auto block"
      >
        {showGoingEvents ? "Masquer les Événements" : "Voir les Événements"}
      </Button>

      {showGoingEvents && goingEvents.length > 0 && (
        <div className="mt-4">
          <ul className="space-y-4">
            {goingEvents?.map((event) => (
              <li key={event?.id} className="flex flex-col">
                <span className="font-bold">{event?.title}
                </span>
                <span>
                  {new Date(event?.date).toLocaleString("fr-FR", {
                    weekday: "long",
                    year: "numeric",
                    month: "long",
                    day: "numeric",
                    hour: "numeric",
                    minute: "numeric",
                  })}
                </span>
              </li>
            ))}
          </ul>
        </div>
      )}

      <h2 className="font-semibold text-center mb-4">Tous les Événements</h2>
      {remainingEvents?.map((event) => (
        <Card key={event.id} className="w-full max-w-md mx-auto mb-4">
          <CardHeader>
            <CardTitle>{event.title}</CardTitle>
            <CardDescription>{event.description}</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <p>
              {new Date(event?.date).toLocaleString("fr-FR", {
                weekday: "long",
                year: "numeric",
                month: "long",
                day: "numeric",
                hour: "numeric",
                minute: "numeric",
              })}
            </p>
          </CardContent>
          <CardFooter className="flex justify-end ml-1">
            <Button
            //   variant="outline"
              onClick={() => respondToEvent(event.id, "going")}
            >
              Participer
            </Button>
            <Button
              variant="outline"
              onClick={() => respondToEvent(event.id, "not_going")}
            >
              Ne pas participer
            </Button>
          </CardFooter>
        </Card>
      ))}
    </div>
  );
}
