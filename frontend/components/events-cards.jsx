import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from "@/components/ui/card"
import { Button } from "@/components/ui/button"

export function EventsCards() {
  return (
    (<Card className="w-full max-w-md mx-auto">
      <CardHeader>
        <CardTitle>Soirée de lancement</CardTitle>
        <CardDescription>Venez célébrer le lancement de notre nouvelle application avec nous !</CardDescription>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="flex items-center justify-between">
          <div className="space-y-1">
            <p className="text-sm font-medium text-muted-foreground">Date et heure</p>
            <p>Vendredi 15 septembre, 19h00</p>
          </div>
          <CalendarIcon className="h-6 w-6 text-muted-foreground" />
        </div>
      </CardContent>
      <CardFooter className="flex justify-end">
        <Button variant="outline" className="mr-2">
          Se désinscrire
        </Button>
        <Button>S&apos;inscrire</Button>
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

function MapPinIcon(props) {
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
      <path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z" />
      <circle cx="12" cy="10" r="3" />
    </svg>)
  );
}