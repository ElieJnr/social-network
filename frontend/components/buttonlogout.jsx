import { useRouter } from "next/navigation";

import { LogOutIcon } from "lucide-react";

export default function LogoutButton() {
  const router = useRouter();
  const handleLogout = async () => {
    try {
      const response = await fetch("http://localhost:8080/logout", {
        method: "POST",
        credentials: "include",
      });

      if (response.ok) {
        router.push("/auth/login");
      } else {
        console.error("Failed to log out");
      }
    } catch (error) {
      console.error("An error occurred during logout:", error);
    }
  };

  return (
    <button className="flex items-center gap-2" onClick={handleLogout}>
      <LogOutIcon className="h-5 w-5" />
    Logout
    </button>
  );
}
