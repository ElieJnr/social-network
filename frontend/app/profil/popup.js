import { Popover, PopoverTrigger, PopoverContent } from "@/components/ui/popover";
import { Button } from "@/components/ui/button";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import { searchUsers } from "../actions/users";
import { useEffect, useState } from "react";

export default function UserListModal({ isOpen, onClose, title, users, statut }) {
    const [tabUsers, setTabUsers] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        async function fetchUsers(users, statut) {
            if (users) {
                setLoading(true);
                try {
                    const usersList = await searchUsers(users, statut);
                    setTabUsers(usersList.map(item => item.user));
                } catch (error) {
                    console.error("Erreur lors du chargement des utilisateurs:", error);
                } finally {
                    setLoading(false);
                }
            }
        }
        if (isOpen) {
            fetchUsers(users, statut);
        }
    }, [users, statut, isOpen]);

    const handleUserClick = (userId) => {
        window.location.href = `/profil?userId=${userId}`;
    };
   
    return (
        <>
            {isOpen && <div className="fixed inset-0 bg-gray-800 opacity-50 z-40"></div>}
            <Popover open={isOpen} onOpenChange={onClose}>
                <PopoverTrigger asChild>
                    <Button variant="outline" className="w-full bg-[#e2e1e1] border-none">
                        {title}
                    </Button>
                </PopoverTrigger>
                <PopoverContent className="w-[400px] p-6 z-50 relative">
                    <div className="flex items-center justify-between mb-4">
                        <h3 className="text-lg font-medium">{title}</h3>
                    </div>
                    <div className="space-y-4">
                        {loading ? (
                            <div>Loading...</div>
                        ) : (
                            tabUsers.length === 0 ? (
                                <div>No users found.</div>
                            ) : (
                                tabUsers.map((user) => (
                                    <div key={user.id} className="flex items-center gap-4">
                                        <Avatar className="cursor-pointer" onClick={() => handleUserClick(user.id)}>
                                            <AvatarImage src={"/placeholder-user.jpg"} alt={user.firstname} />
                                            <AvatarFallback>{user.firstname.charAt(0)}</AvatarFallback>
                                        </Avatar>
                                        <div className="flex-1">
                                            <div className="font-medium">{user.firstname}</div>
                                        </div>
                                    </div>
                                ))
                            )
                        )}
                    </div>
                </PopoverContent>
            </Popover>
        </>
    );
}

function XIcon(props) {
    return (
        <svg
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
            <path d="M18 6 6 18" />
            <path d="m6 6 12 12" />
        </svg>
    );
}


