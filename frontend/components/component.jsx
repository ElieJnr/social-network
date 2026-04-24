
import { Popover, PopoverTrigger, PopoverContent } from "@/components/ui/popover";
import { Button } from "@/components/ui/button"
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar"

export function Component() {
  return (
    (<Popover>
      <PopoverTrigger asChild>
        <Button variant="outline" className="w-full">
          View Users
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[400px] p-6">
        <div className="flex items-center justify-between mb-4">
          <h3 className="text-lg font-medium">Users</h3>
          <div>
            <Button variant="ghost" size="icon" className="rounded-full">
              <XIcon className="w-4 h-4" />
            </Button>
          </div>
        </div>
        <div className="space-y-4">
          <div className="flex items-center gap-4">
            <Avatar>
              <AvatarImage src="/placeholder-user.jpg" alt="@shadcn" />
              <AvatarFallback>JD</AvatarFallback>
            </Avatar>
            <div className="flex-1">
              <div className="font-medium">John Doe</div>
              <div className="text-sm text-muted-foreground">Admin</div>
            </div>
          </div>
          <div className="flex items-center gap-4">
            <Avatar>
              <AvatarImage src="/placeholder-user.jpg" alt="@shadcn" />
              <AvatarFallback>JA</AvatarFallback>
            </Avatar>
            <div className="flex-1">
              <div className="font-medium">Jane Appleseed</div>
              <div className="text-sm text-muted-foreground">Editor</div>
            </div>
          </div>
          <div className="flex items-center gap-4">
            <Avatar>
              <AvatarImage src="/placeholder-user.jpg" alt="@shadcn" />
              <AvatarFallback>LS</AvatarFallback>
            </Avatar>
            <div className="flex-1">
              <div className="font-medium">Liam Smith</div>
              <div className="text-sm text-muted-foreground">Viewer</div>
            </div>
          </div>
        </div>
      </PopoverContent>
    </Popover>)
  );
}

function XIcon(props) {
  return (
    (<svg
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
      <path d="M18 6 6 18" />
      <path d="m6 6 12 12" />
    </svg>)
  );
}
