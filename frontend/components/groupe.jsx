'use client'
import { useRouter } from 'next/navigation'
import { useState, useEffect } from 'react'
import { GroupCards } from './group-cards'
import { Content } from 'next/font/google'
import { fetchGroupes } from '@/app/actions/groupe'
import { socketSend } from '@/app/actions/message'
export default function AllGroupsComponent({ socket }) {
  const router = useRouter()
  const [groupes, setGroupes] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    fetchGroupes(setGroupes,setLoading,setError)
  }, [])

  const rejoindreGroupe = (id, UserId, title) => {
    const Message = {
      Type: "notifications",
      ReceiverId: UserId,
      GroupeId: id,
      SubType: "addGroupe",
      Content: title,
    }
    socketSend(socket,Message)
    window.location.reload()
  }

  const entrerGroupe = (id) => {
    router.push(`/groups/${id}`)

  }

  if (loading) return <p>Chargement...</p>
  if (error) return <p>Erreur: {error}</p>

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-6">Groupes disponibles</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {groupes?.map((groupe) => (
          <GroupCards
            key={groupe.id}
            nom={groupe.title}
            description={groupe.description}
            estMembre={groupe.isMember}
            onEntrer={() => entrerGroupe(groupe.id)}
            onRejoindre={() => rejoindreGroupe(groupe.id, groupe.userId, groupe.title)}
          />
        ))}
      </div>
    </div>
  )
}