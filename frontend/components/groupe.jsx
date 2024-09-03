'use client'

import { useState } from 'react'
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"

export default function AllGroupsComponent() {
  // État pour stocker les groupes (normalement, cela viendrait d'une API)
  const [groupes, setGroupes] = useState([
    { id: '1', nom: 'Groupe A', description: 'Description du groupe A', estMembre: true },
    { id: '2', nom: 'Groupe B', description: 'Description du groupe B', estMembre: false },
    { id: '3', nom: 'Groupe C', description: 'Description du groupe C', estMembre: true },
    { id: '4', nom: 'Groupe D', description: 'Description du groupe D', estMembre: false },
  ])

  // Fonction pour rejoindre un groupe
  const rejoindreGroupe = (id) => {
    setGroupes(groupes.map(groupe => 
      groupe.id === id ? { ...groupe, estMembre: true } : groupe
    ))
  }

  // Fonction pour entrer dans un groupe (à implémenter selon vos besoins)
  const entrerGroupe = (id) => {
    console.log(`Entrer dans le groupe ${id}`)
    // Implémentez ici la logique pour entrer dans le groupe
  }

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-6">Groupes disponibles</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {groupes.map((groupe) => (
          <Card key={groupe.id} className="flex flex-col">
            <CardHeader>
              <CardTitle>{groupe.nom}</CardTitle>
            </CardHeader>
            <CardContent className="flex-grow">
              <p>{groupe.description}</p>
            </CardContent>
            <CardFooter>
              {groupe.estMembre ? (
                <Button className="w-full" onClick={() => entrerGroupe(groupe.id)}>
                  Entrer
                </Button>
              ) : (
                <Button className="w-full" onClick={() => rejoindreGroupe(groupe.id)}>
                  Rejoindre
                </Button>
              )}
            </CardFooter>
          </Card>
        ))}
      </div>
    </div>
  )
}