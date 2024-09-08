
"use client"

import { useEffect, useState } from "react"
import { Popover, PopoverTrigger, PopoverContent } from "@/components/ui/popover";
import { Button } from "@/components/ui/button"
import { fetchNotifs } from "@/app/actions/notifications";
import { Check } from 'lucide-react'
import { cn } from "@/lib/utils";
import { useRouter } from "next/navigation";
import { BellIcon, GetIcon } from "./IconNotif";
import { socketSend } from "@/app/actions/message";
import { fetchForAddingInAGroupe, fetchForNotAddingInAGroupe } from "@/app/actions/groupe";
import { follow } from "@/app/actions/follow";
export let setNotifs
export function Notifications({ socket }) {
    const ok = true
    const router = useRouter()
    const [isOpen, setIsOpen] = useState(false)
    const [isMounted, setIsMounted] = useState(false);
    const [notifications, setNotifications] = useState([]);
    const unreadCount = notifications.filter(n => !n.IsRead && n.Type != "msg").length
    setNotifs = setNotifications
    const TitleNotif = {
        event: "Upcoming event",
        follow: "New following request",
        invitation: "New group invitation",
        addGroupe: "Request to join a group"
    }
    useEffect(() => {
        setIsMounted(true);
        fetchNotifs(setNotifications);
    }, []);

    const toggleDropdown = () => {
        setIsOpen(true)
    }

    const markAsRead = (id, userId, Type, GroupId, SenderID, ok) => {
        const Message = {
            Type: "notifications",
            ReceiverId: id,
            SubType: "read"
        }
        console.log(id);

        socketSend(socket, Message)
        fetchNotifs(setNotifications);
        if (Type == "follow") {
            console.log(ok, SenderID, userId, true, true);
            ok ? follow(SenderID, userId, true, ok) : follow(SenderID, userId, false, ok)
        }
        if (Type == "addGroupe") {
            if (ok) {
                fetchForAddingInAGroupe("member", GroupId, SenderID)
            } else {
                fetchForNotAddingInAGroupe(GroupId, SenderID)
            }
        }
    }


    return (
        (<Popover>
            <PopoverTrigger asChild>
                <Button variant="ghost" size="icon" className="relative" onClick={toggleDropdown}>
                    <BellIcon className={cn("h-6 w-6 text-muted-foreground", { 'animate-bounce': unreadCount > 0 })} />
                    {unreadCount > 0 && <span
                        className="absolute -top-1 -right-1 bg-red-500 text-white rounded-full px-1.5 py-0.5 text-xs font-medium">
                        {unreadCount}
                    </span>}
                </Button>
            </PopoverTrigger>
            {isOpen && (
                <PopoverContent align="end" className="w-80 p-4 rounded-lg shadow-lg" >
                    <div className="text-lg font-semibold">Notifications</div>
                    <div />
                    {notifications.map((n) => ((!n.IsRead && n.Type != "msg") && (
                        <div className="space-y-4 mb-2" key={n.Id}>
                            <div className="flex items-start gap-3">
                                <GetIcon type={n.Type} className="h-5 w-5" />
                                <div className="text-sm">
                                    <p className="font-medium">{TitleNotif[n.Type]}</p>
                                    <p className="text-sm text-muted-foreground">{n.SenderInfo.Email + " " + n.Message}.</p>
                                    <p className="text-xs text-muted-foreground">{n.Formated_date}</p>
                                    {((n.Type == "follow" && (n.ReceiverInfo.IsPrivate)) || n.Type == "invitation" || n.Type == "addGroupe") && (<div className="flex gap-2 mt-2">
                                        <Button variant="outline" size="sm" onClick={() => markAsRead(n.Id, n.ReceiverID, n.Type, n.GroupId, n.SenderID, ok)}>
                                            Accept
                                        </Button>
                                        <Button variant="outline" size="sm" onClick={() => markAsRead(n.Id, n.ReceiverID, n.Type, n.GroupId, n.SenderID, !ok)}>
                                            Decline
                                        </Button>
                                    </div>)}
                                </div>
                                <div className="flex items-center space-x-2">
                                    {((n.Type == "event") || ((n.Type == "follow") && (!n.ReceiverInfo.IsPrivate))) && (
                                        <Button size="icon" variant="ghost" onClick={() => markAsRead(n.Id)}>
                                            <Check className="h-4 w-4" />
                                            <span className="sr-only">Marquer comme lu</span>
                                        </Button>
                                    )}
                                </div>
                            </div>
                        </div>
                    )))}
                </PopoverContent>
            )}
        </Popover>)
    );
}


