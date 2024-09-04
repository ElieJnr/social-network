"use client";

import { Button } from "@/components/ui/button";
import { useState, useEffect } from "react";
import { GetAllInfoForUserById, search, userConnect } from "../actions/users";
import NavBar from "@/components/Nav";
import UserListModal from "./popup";
import Image from "next/image";
import { follow } from "../actions/follow";
import { useWebSocket } from "../actions/message";
export default function Profil() {
    const socket = useWebSocket('ws://localhost:8080/ws');
    const [userId, setUserId] = useState(null);
    const [user, setUser] = useState(null);
    const [userOnLine, setUserOnLine] = useState(null)
    const [refreshTrigger, setRefreshTrigger] = useState(false);
    const [posts, setPosts] = useState([])
    const [isFollowersOpen, setFollowersOpen] = useState(false);
    const [isFollowingOpen, setFollowingOpen] = useState(false);

    useEffect(() => {
        const urlParams = new URLSearchParams(window.location.search);
        const userIdFromQuery = urlParams.get('userId');
        setUserId(userIdFromQuery);
    }, []);


    useEffect(() => {
        const fetchUser = async () => {
            try {
                const response = await userConnect();
                setUserOnLine(response.user);
            } catch (error) {
                console.error("Error fetching user:", error);
            }
        };
        fetchUser();
    }, [refreshTrigger]);

    useEffect(() => {
        const fetchUserData = async () => {
            if (userId) {
                try {
                    const userData = await GetAllInfoForUserById(userId);
                    setUser(userData.user);
                } catch (error) {
                    console.error("Erreur lors de la récupération de l'utilisateur:", error);
                }
            }
        };

        fetchUserData();
    }, [userId, refreshTrigger]);

    if (!user || !userOnLine) {
        return <div>Loading...</div>;
    }
    const handleClick = (user, statut) => {
        console.log(user.id);
        follow(userOnLine.id, user.id, statut, false);
        setRefreshTrigger(prev => !prev);

    };


    return (
        <>
            <NavBar socket={socket} />
            <div className="bg-background text-foreground min-h-screen flex flex-col">
                <div className="flex-1 grid gap-6 p-6">
                    <div className="grid gap-4">
                        <div className="flex items-center justify-between">
                            <div className="grid gap-1">
                                <div className="font-medium text-2xl">
                                    {user ? user.firstname : "loading"} {user ? user.lastname : ""}
                                </div>
                                <div className="text-muted-foreground">Software Engineer</div>
                            </div>
                            <div className="flex items-center gap-4">
                                {userOnLine.id !== user.id && (
                                    <Button onClick={() => handleClick(user, !user.isPrivate)} variant="outline" size="sm">
                                        {(search([user], userOnLine?.follows || [])).length === 0 ? "UnFollow" : "Follow"}
                                    </Button>
                                )}
                                <Button variant="ghost" size="icon" className="rounded-full">
                                    <MoveHorizontalIcon className="w-5 h-5" />
                                </Button>
                            </div>

                        </div>
                        <div className="text-sm leading-loose text-muted-foreground">
                            {user ? user.bio : ""}
                        </div>
                    </div>
                    <div className="grid sm:grid-cols-3 gap-4 text-center">
                        <div className="bg-[#e2e1e1] rounded-lg p-4">
                            <div className="font-medium">100</div>
                            <div className="text-xs text-muted-foreground">Posts</div>
                        </div>
                        <div onClick={() => setFollowersOpen(true)} className="bg-[#e2e1e1] rounded-lg p-4 cursor-pointer">
                            <div className="font-medium">{user?.followers ? user.followers.length : "0"}</div>
                            <UserListModal
                                isOpen={isFollowersOpen}
                                onClose={() => setFollowersOpen(false)}
                                title="Followers"
                                users={user?.followers || []}
                                statut="userId"
                            />
                        </div>
                        <div className="bg-[#e2e1e1] rounded-lg p-4 cursor-pointer" onClick={() => setFollowingOpen(true)}>
                            <div className="font-medium">{user && user.follows ? user.follows.length : "0"}</div>
                            <UserListModal
                                isOpen={isFollowingOpen}
                                onClose={() => setFollowingOpen(false)}
                                title="Following"
                                users={user?.follows || []}
                                statut="followedUser"
                            />
                        </div>
                    </div>
                    <div className="grid gap-6">
                        <div className="flex items-center justify-between">
                            <div className="font-medium">Posts</div>
                            <div className="flex items-center gap-2">
                                <Button variant="outline" size="sm">
                                    <EyeIcon className="w-4 h-4 mr-2" />
                                    Public
                                </Button>
                                <Button variant="outline" size="sm">
                                    <LockIcon className="w-4 h-4 mr-2" />
                                    Private
                                </Button>
                            </div>
                        </div>
                        <div className="flex justify-center items-center min-h-screen">
                            <div className="w-[50%] flex flex-col justify-center items-center">
                                <div className="bg-card rounded-lg overflow-hidden">
                                    <Image
                                        src="/placeholder.svg"
                                        alt="Post Image"
                                        width={600}
                                        height={400}
                                        className="w-full aspect-[3/2] object-cover"
                                    />
                                    <div className="p-4">
                                        <div className="font-medium line-clamp-2">Introducing the new Vercel platform</div>
                                        <div className="text-xs text-muted-foreground">2 days ago</div>
                                    </div>
                                </div>
                                <div className="bg-card rounded-lg overflow-hidden">
                                    <Image
                                        src="/placeholder.svg"
                                        alt="Post Image"
                                        width={600}
                                        height={400}
                                        className="w-full aspect-[3/2] object-cover"
                                    />
                                    <div className="p-4">
                                        <div className="font-medium line-clamp-2">Introducing the new Vercel platform</div>
                                        <div className="text-xs text-muted-foreground">2 days ago</div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
}

function EyeIcon(props) {
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
            <path d="M2 12s3-7 10-7 10 7 10 7-3 7-10 7-10-7-10-7Z" />
            <circle cx="12" cy="12" r="3" />
        </svg>
    );
}

function LockIcon(props) {
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
            <rect width="18" height="11" x="3" y="11" rx="2" ry="2" />
            <path d="M7 11V7a5 5 0 0 1 10 0v4" />
        </svg>
    );
}

function MoveHorizontalIcon(props) {
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
            <polyline points="18 8 22 12 18 16" />
            <polyline points="6 8 2 12 6 16" />
            <line x1="2" x2="22" y1="12" y2="12" />
        </svg>
    );
}
