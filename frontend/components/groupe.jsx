'use client'

import { useState, useEffect } from 'react'
import { GroupCards } from './group-cards'

export default function AllGroupsComponent() {
  const [groupes, setGroupes] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    const fetchGroupes = async () => {
      try {
        const response = await fetch('http://localhost:8080/group/getGroups')
        if (!response.ok) {
          throw new Error('Erreur lors de la récupération des groupes')
        }
        const data = await response.json()

        console.log('Données récupérées:', data) // Log the fetched data

        setGroupes(data)
      } catch (error) {
        setError(error.message)
      } finally {
        setLoading(false)
      }
    }

    fetchGroupes()
  }, [])

  const rejoindreGroupe = (id) => {
    setGroupes(groupes.map(groupe => 
      groupe.id === id ? { ...groupe, estMembre: true } : groupe
    ))
  }

  const entrerGroupe = (id) => {
    console.log(`Entrer dans le groupe ${id}`)
  }

  if (loading) return <p>Chargement...</p>
  if (error) return <p>Erreur: {error}</p>

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-6">Groupes disponibles</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {groupes.map((groupe) => (
          <GroupCards
            key={groupe.id} // Ensure groupe.id is unique and defined
            nom={groupe.title}
            description={groupe.description}
            estMembre={groupe.isMember}
            onEntrer={() => entrerGroupe(groupe.id)}
            onRejoindre={() => rejoindreGroupe(groupe.id)}
          />
        ))}
      </div>
    </div>
  )
}