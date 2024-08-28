import React from "react";
import { Button } from "./ui/button";
import { LogOutIcon } from "lucide-react";
import { useRouter } from "next/navigation";

const Logout = () => {
    const router = useRouter();
    const handleLogout = async () => {
      try {
        const response = await fetch('http://localhost:8080/logout', {
          method: 'POST',
          credentials: 'include',
        });
        if (response.ok) {
          router.push('/auth');
        } else {
          console.error('Failed to log out');
        }
      } catch (error) {
        console.error('An error occurred during logout:', error);
      }
    };
  return (
    <Button variant="ghost" onClick={handleLogout}>
      <LogOutIcon className="h-5 w-5" />
    </Button>
  );
};

export default Logout;
