

export const fetchGroupes = async (setGroupes,setLoading,setError) => {
    try {
      const response = await fetch('http://localhost:8080/group/getGroups', {
        method: "GET",
        credentials: "include",
      });
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