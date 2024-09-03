
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "@/components/ui/card"
import { Button } from "@/components/ui/button"


export function GroupCards({ nom, description, estMembre, onEntrer, onRejoindre }) {
  return (
    <Card className="flex flex-col">
      <CardHeader>
        <div className="flex justify-between items-center">
          <CardTitle className="text-xl">{nom}</CardTitle>
          <Button className=" text-white" onClick={estMembre ? onEntrer : onRejoindre}>
            {estMembre ? 'Voir' : 'Rejoindre'}
          </Button>
        </div>
      </CardHeader>
      <CardContent className="mt-2">
        <p className="text-sm text-muted-foreground">{description}
        </p>
        <div className="flex justify-between mt-4">
          <div className="text-center">
            <p className="text-lg font-bold">989 k</p>
            <p className="text-sm text-muted-foreground">Membres</p>
          </div>
          <div className="text-center">
            <p className="text-lg font-bold">40</p>
            <div className="flex items-center justify-center">
              <span className="h-2 w-2 bg-green-500 rounded-full mr-1" />
              <p className="text-sm text-muted-foreground">En ligne</p>
            </div>
          </div>
          <div className="text-center">
            <p className="text-lg font-bold">Premiers 1 %</p>
            <p className="text-sm text-muted-foreground">Classer par popularité</p>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}


