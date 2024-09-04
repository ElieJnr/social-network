"use client";


import { CreateGroupCard } from "@/components/create-group-card";
import AllGroupsComponent from "@/components/groupe";
import { useWebSocket } from "../actions/message";
import ProfileCard from "@/components/ProfileCard";
import NavBar from "@/components/Nav";
import { GroupHomePage } from "@/components/group-home-page";
export default function AllGroup() {
   const socket = useWebSocket('ws://localhost:8080/ws');
    
    return (
        <div className="flex flex-col h-screen">
            <NavBar socket={socket} />
            <div className="flex-1 grid grid-cols-[0.5fr_2fr] gap-6 p-6">
                <div className="space-y-6">
                    <ProfileCard />
                    <CreateGroupCard />
                </div>
                <div className="space-y-6">
                    <AllGroupsComponent />
                </div>
            </div>
        </div>
    );
}

export function SeeGroup() {
    return (
        <div className="flex flex-col h-screen">
            <NavBar socket={socket} />
            <div className="flex-1 ">
                <GroupHomePage />
            </div>
        </div>
    );
}