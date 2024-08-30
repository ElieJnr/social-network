"use client"
/**
 * v0 by Vercel.
 * @see https://v0.dev/t/Dk4wHU0kBRP
 * Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
 */
import { Button } from "@/components/ui/button"
import { fetchNotifs } from "../actions/notifications"
import { useEffect, useState } from "react"
import { XIcon, IconNotif } from "@/components/IconNotif"
import { useRouter } from "next/navigation"

export default function AllNotifs() {
    const router = useRouter()
    const [Len, setLen] = useState(0)
    const [isMounted, setIsMounted] = useState(false);
    const [notifications, setNotifications] = useState([]);

    useEffect(() => {
        setIsMounted(true)
        fetchNotifs(setLen, setNotifications);
    }, []);

    const Home = () => {
        if (isMounted) {
            router.push("/")
        }
    }
    return (
        <div className="bg-white shadow-sm">
            <div className="container mx-auto px-4 py-6">
                <div className="flex items-center justify-between mb-6">
                    <h1 className="text-2xl font-bold">Notifications</h1>
                    <Button variant="ghost" size="icon" className="relative" onClick={Home}>
                        <XIcon className="h-6 w-6 text-muted-foreground" />
                    </Button>
                </div>
                <div className="space-y-4">
                    {notifications.map((n) => (
                        <div className="flex items-start gap-3" key={n.Id}>
                            <IconNotif type={n.Type} className="h-5 w-5" />
                            <div>
                                <p className="font-medium">{n.Message}</p>
                                <p className="text-sm text-muted-foreground">Hey, I just wanted to say hi and see how you're doing.</p>
                                <p className="text-xs text-muted-foreground">{n.CreateAt}</p>
                                {(n.Type == "follow" || n.Type == "invitation") && (<div className="flex gap-2 mt-2">
                                    <Button variant="outline" size="sm" id={n.Id}>
                                        Accept
                                    </Button>
                                    <Button variant="outline" size="sm" id={n.Id}>
                                        Decline
                                    </Button>
                                </div>)}
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    )
}




