"use client";

const { BellIcon, LogOutIcon, MountainIcon } = require("lucide-react");
const { Button } = require("./ui/button");
const { default: Link } = require("next/link");
const { useRouter } = require("next/navigation");

export default function NavBar() {
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
      <header className="bg-primary text-primary-foreground py-4 px-6">
        <div className="container mx-auto flex items-center justify-between">
          <div className="flex items-center gap-4">
            <Link href="#" prefetch={false}>
              <MountainIcon className="h-6 w-6" />
              <span className="sr-only">Acme Inc</span>
            </Link>
            <nav className="hidden md:flex items-center gap-4">
              <Link href="#" className="hover:underline" prefetch={false}>
                Home
              </Link>
              <Link href="#" className="hover:underline" prefetch={false}>
                Groups
              </Link>
              <Link href="#" className="hover:underline" prefetch={false}>
                Chat
              </Link>
            </nav>
          </div>
          <div className="flex items-center">
            <Button variant="ghost" >
              <BellIcon className="h-5 w-5" />
            </Button>
            <Button variant="ghost" onClick={handleLogout}>
              <LogOutIcon className="h-5 w-5" />
            </Button>
          </div>
        </div>
      </header>
    );
  }
  