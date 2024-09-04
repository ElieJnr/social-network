import { GroupHomePage } from "@/components/group-home-page";
import NavBar from "./Nav";
import { useWebSocket } from "@/app/actions/message";

export default function SeeGroup() {
    const socket = useWebSocket('ws://localhost:8080/ws');
    return (
        <div className="flex flex-col h-screen">
            <NavBar socket={socket} />
            <div className="flex-1 ">
                <GroupHomePage />
            </div>
        </div>
    );
}