import { Button } from "@/components/ui/button";
import { LockIcon } from "lucide-react";
import { EyeIcon } from "lucide-react";
import { editProfil } from "../actions/users";

export default function Edit({ onClose, user }) {
    
     const handlePrivate = (isPrivate) => {
         editProfil(user.id, isPrivate)
         onClose()
     };

    return (

        <div className="flex items-center gap-2">
            <Button onClick={() => handlePrivate(false)} variant="outline" size="sm">
                <EyeIcon className="w-4 h-4 mr-2" />
                Public
            </Button>
            <Button onClick={() => handlePrivate(true)} variant="outline" size="sm">
                <LockIcon className="w-4 h-4 mr-2" />
                Private
            </Button>
        </div>

    );
}

