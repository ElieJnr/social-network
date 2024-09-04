
import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { Input } from "@/components/ui/input"
import { Textarea } from "@/components/ui/textarea"
import { Popover, PopoverTrigger, PopoverContent } from "@/components/ui/popover"
import { Button } from "@/components/ui/button"
import { Calendar } from "@/components/ui/calendar"

export function CreateEventCard() {
  return (
    (<Card className="w-full">
      <CardHeader>
        <CardTitle>Créer un événement</CardTitle>
        <CardDescription>Remplissez les informations pour créer un nouvel événement.</CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="space-y-2">
          <Label htmlFor="event-name">Nom de l'événement</Label>
          <Input id="event-name" placeholder="Entrez le nom de l'événement" />
        </div>
        <div className="space-y-2">
          <Label htmlFor="event-description">Description</Label>
          <Textarea id="event-description" placeholder="Décrivez l'événement" />
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
        <Button>Créer l'événement</Button>
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
