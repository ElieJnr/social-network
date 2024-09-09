"use client";

import { Button } from "@/components/ui/button";
import { useState, useEffect } from "react";
import {
    GetAllInfoForUserById,
    search,
    userConnect,
    searchUsers,
} from "../actions/users";
import NavBar from "@/components/Nav";
import UserListModal from "./popup";
import Image from "next/image";
import { follow } from "../actions/follow";
import Edit from "./edit";
import { MailIcon } from "lucide-react";
import { CalendarIcon } from "lucide-react";
import { UserIcon } from "lucide-react";
import { InfoIcon } from "lucide-react";
import { Card, CardContent } from "@/components/ui/card";
import { UsersIcon } from "lucide-react";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import { socketSend, useWebSocket } from "../actions/message";
import { PostGroupCard } from "@/components/group-home-page";
import { PostCard, SkeletonPostCard } from "@/components/PostCard";
import { filterPost } from "../actions/post";
import { useToast } from "@/components/ui/use-toast";

export default function Profil() {
    const socket = useWebSocket("ws://localhost:8080/ws");
    const [userId, setUserId] = useState(null);
    const [user, setUser] = useState(null);
    const [userOnLine, setUserOnLine] = useState(null);
    const [refreshTrigger, setRefreshTrigger] = useState(false);
    const [posts, setPosts] = useState([]);
    const [isFollowersOpen, setFollowersOpen] = useState(false);
    const [isFollowingOpen, setFollowingOpen] = useState(false);
    const [edit, setEdit] = useState(false);
    const [tabRequest, setTabRequest] = useState([]);
    const { toast } = useToast();
    useEffect(() => {
        const urlParams = new URLSearchParams(window.location.search);
        const userIdFromQuery = urlParams.get("userId");
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
        if (userOnLine && userOnLine.requestF) {
            const fetchRequests = async () => {
                try {
                    const response = await searchUsers(
                        userOnLine.requestF,
                        "followedUser"
                    );
                    setTabRequest(response.map((item) => item.user));
                } catch (error) {
                    console.error("Error fetching requests:", error);
                }
            };
            fetchRequests();
        }
    }, [userOnLine, refreshTrigger]);

    useEffect(() => {
        const fetchUserData = async () => {
            if (userId) {
                try {
                    const userData = await GetAllInfoForUserById(userId);
                    setUser(userData.user);
                } catch (error) {
                    console.error(
                        "Erreur lors de la récupération de l'utilisateur:",
                        error
                    );
                }
            }
        };

        fetchUserData();
    }, [userId, refreshTrigger]);

    if (!user || !userOnLine) {
        return <div>Loading...</div>;
    } else {
        console.log(userOnLine.posts == null ? "0" : userOnLine.posts.length);
        console.log(userOnLine);
        console.log(tabRequest);
    }
    const handleClick = async (user, statut, value) => {
        // console.log(statut);
        let ok = statut;
        try {
            const reponse = GetAllInfoForUserById(user.id);
            ok = await reponse.user.isPrivate;
            console.log(ok);
        } catch (error) { }
        console.log(ok);
        follow(userOnLine.id, user.id, ok, false);
        setRefreshTrigger((prev) => !prev);
        ok
            ? toast({
                title: "Follow Successful",
                description: `You follow now ${user.firstname} .`,
            })
            : toast({
                title: "Request",
                description: "Your request has been send.",
            });
        if (value == "Follow") {
            const Message = {
                Type: "notifications",
                ReceiverId: user.id,
                SubType: "sendFollow",
                Content: "wants to follow you",
                IsPrivate: user.IsPrivate,
            };
            setTimeout(() => {
                socketSend(socket, Message);
            }, 350);
        }
    };
    const CanSee = (user) => {
        return search([user], userOnLine?.follows || []).length === 0;
    };

    return (
        <>
            <NavBar socket={socket} />
            <div className="flex justify-between">
                <div className="bg-background text-foreground w-[100%] min-h-screen flex flex-col overflow-x-auto">
                    <div className="flex-1 grid gap-6 p-6">
                        <div className="grid gap-4">
                            <div className="flex items-center justify-between">
                                <div className="flex flex-col">
                                    <div className="flex ">
                                        <Avatar className="w-[8rem] h-[8rem]">
                                            <AvatarImage
                                                src={
                                                    user?.avatar === ""
                                                        ? "/placeholder-user.jpg"
                                                        : `/uploads/${user?.avatar}`
                                                }
                                                alt="@shadcn"
                                            />
                                            <AvatarFallback>CN</AvatarFallback>
                                        </Avatar>
                                        <div className="flex flex-col">
                                            <div className="font-medium text-2xl">
                                                {user ? user.firstname : "loading"}{" "}
                                                {user ? user.lastname : ""}
                                            </div>
                                            {CanSee(user) ||
                                                !user.isPrivate ||
                                                userOnLine.id == user.id ? (
                                                <>
                                                    <div className="flex items-center gap-2">
                                                        <MailIcon className="h-4 w-4 text-muted-foreground" />
                                                        <span className="text-muted-foreground">
                                                            {user ? user.email : ""}
                                                        </span>
                                                    </div>
                                                    <div className="flex items-center gap-2">
                                                        <CalendarIcon className="h-4 w-4 text-muted-foreground" />
                                                        <span className="text-muted-foreground">
                                                            {user ? user.dateOfBirth : ""}
                                                        </span>
                                                    </div>
                                                    <div className="flex items-center gap-2">
                                                        {user && user.username ? (
                                                            <UserIcon className="h-4 w-4 text-muted-foreground" />
                                                        ) : (
                                                            ""
                                                        )}
                                                        <span className="text-muted-foreground">
                                                            {user ? user.username : ""}
                                                        </span>
                                                    </div>
                                                    <div className="flex items-start gap-2">
                                                        {user && user.bio ? (
                                                            <InfoIcon className="h-4 w-4 text-muted-foreground" />
                                                        ) : (
                                                            ""
                                                        )}
                                                        <span className="text-muted-foreground">
                                                            {user ? user.bio : ""}
                                                        </span>
                                                    </div>
                                                </>
                                            ) : (
                                                ""
                                            )}
                                        </div>
                                    </div>
                                </div>
                                <div className="flex items-center gap-4">
                                    {!tabRequest.some((request) => request.id === user.id) ? (
                                        userOnLine.id !== user.id && (
                                            <Button
                                                onClick={(e) => handleClick(user, !user.isPrivate, e.target.textContent)}
                                                variant="outline"
                                                size="sm"
                                            >
                                                {CanSee(user) ? "UnFollow" : "Follow"}
                                            </Button>
                                        )
                                    ) : (
                                        <Button variant="outline" size="sm">
                                            Demande
                                        </Button>
                                    )}

                                    {userOnLine.id === user.id && (
                                        <>
                                            <Button
                                                onClick={() => setEdit(!edit)}
                                                variant="outline"
                                                size="sm"
                                            >
                                                Edit Profil
                                            </Button>
                                            {edit && (
                                                <Edit
                                                    isOpen={edit}
                                                    onClose={() => setEdit(false)}
                                                    user={userOnLine}
                                                />
                                            )}
                                        </>
                                    )}

                                    <Button variant="ghost" size="icon" className="rounded-full">
                                        <MoveHorizontalIcon className="w-5 h-5" />
                                    </Button>
                                </div>
                            </div>
                        </div>
                        <div className="grid sm:grid-cols-3 gap-4 text-center">
                            <div className="bg-[#e2e1e1] rounded-lg p-4">
                                <div className="font-medium">
                                    {user.posts == null ? "0" : filterPost(user.posts).length}
                                </div>
                                <div className=" font-medium ">Posts</div>
                            </div>
                            <div
                                onClick={() => setFollowersOpen(true)}
                                className="bg-[#e2e1e1] rounded-lg p-4 cursor-pointer"
                            >
                                <div className="font-medium">
                                    {user?.followers ? user.followers.length : "0"}
                                </div>
                                <UserListModal
                                    isOpen={isFollowersOpen}
                                    onClose={() => setFollowersOpen(false)}
                                    title="Followers"
                                    users={user?.followers || []}
                                    statut="userId"
                                />
                            </div>
                            <div
                                className="bg-[#e2e1e1] rounded-lg p-4 cursor-pointer"
                                onClick={() => setFollowingOpen(true)}
                            >
                                <div className="font-medium">
                                    {user && user.follows ? user.follows.length : "0"}
                                </div>
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
                                    <Button variant={user.isPrivate ? "outline" : ""} size="sm">
                                        <EyeIcon className="w-4 h-4 mr-2" />
                                        Public
                                    </Button>
                                    <Button variant={!user.isPrivate ? "outline" : ""} size="sm">
                                        <LockIcon className="w-4 h-4 mr-2" />
                                        Private
                                    </Button>
                                </div>
                            </div>
                            <div className="flex justify-center  min-h-screen">
                                <div className="w-[60%] flex flex-col  ">
                                    {CanSee(user) ||
                                        !user.isPrivate ||
                                        userOnLine.id == user.id ? (
                                        <PostsProfil posts={filterPost(user.posts)} />
                                    ) : (
                                        ""
                                    )}
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
            strokeLinejoin="round"
        >
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
            strokeLinejoin="round"
        >
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
            strokeLinejoin="round"
        >
            <polyline points="18 8 22 12 18 16" />
            <polyline points="6 8 2 12 6 16" />
            <line x1="2" x2="22" y1="12" y2="12" />
        </svg>
    );
}

function PostsProfil({ posts }) {
    return (
        <div className="space-y-4">
            {posts && posts != null ? (
                posts.map((post) =>
                    post.Can_see ? <PostCard key={post.PostID} post={post} /> : null
                )
            ) : (
                <div className="flex justify-center">No Posts</div>
            )}
        </div>
    );
}
