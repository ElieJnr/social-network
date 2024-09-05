"use client"

import { useState, useEffect } from "react";
import { Checkbox } from "@/components/ui/checkbox";
import { Label } from "@/components/ui/label";
import * as Dialog from "@radix-ui/react-dialog";
import { Button } from "@/components/ui/button";

export function AlmostPrivateUserCardModal({ isOpen, onClose, users, onSelectUsers }) {
  const [selectedUserIds, setSelectedUserIds] = useState([]);

  const handleUserSelect = (userId) => {
    if (selectedUserIds.includes(userId)) {
      setSelectedUserIds(selectedUserIds.filter((id) => id !== userId));
    } else {
      setSelectedUserIds([...selectedUserIds, userId]);
    }
  };

  // Met à jour la sélection d'utilisateurs dans le parent
  useEffect(() => {
    onSelectUsers(selectedUserIds);
  }, [selectedUserIds]);

  return (
    <Dialog.Root open={isOpen} onOpenChange={onClose}>
      <Dialog.Portal>
        <Dialog.Overlay className="fixed inset-0 bg-black bg-opacity-50 z-50" />
        <Dialog.Content className="fixed top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 bg-white p-6 rounded-lg z-50 w-full max-w-xl">
          <Dialog.Title className="text-xl font-bold mb-3">Select Users</Dialog.Title>
          <Dialog.Description className="mb-4">Select users who can see the post.</Dialog.Description>
          <div className="bg-card p-4 rounded-lg shadow-lg mb-4">
            <div className="flex flex-wrap gap-2">
              {selectedUserIds.map((userId) => {
                const user = users.find(user => user.Id === userId);
                return (
                  user && (
                    <div
                      key={userId}
                      className="bg-primary text-primary-foreground px-3 py-1 rounded-full flex items-center"
                    >
                      {user.Firstname} {user.Lastname}
                      <button
                        type="button"
                        className="ml-2 text-primary-foreground hover:text-red-500 transition-colors"
                        onClick={() => handleUserSelect(userId)}
                      >
                        <XIcon className="w-3 h-3" />
                      </button>
                    </div>
                  )
                );
              })}
            </div>
          </div>
          <div className="bg-card p-4 rounded-lg shadow-lg">
            <h2 className="text-xl font-bold mb-4">Followers</h2>
            <div className="grid gap-3">
              {users.map((user) => (
                <div key={user.Id} className="flex items-center space-x-2">
                  <Checkbox
                    id={`user-${user.Id}`}
                    checked={selectedUserIds.includes(user.Id)}
                    onCheckedChange={() => handleUserSelect(user.Id)}
                  />
                  <Label htmlFor={`user-${user.Id}`} className="cursor-pointer">
                    {user.Firstname} {user.Lastname}
                  </Label>
                </div>
              ))}
            </div>
          </div>
          <Dialog.Close asChild>
            <Button variant="outline" className="mt-4 w-full">Close</Button>
          </Dialog.Close>
        </Dialog.Content>
      </Dialog.Portal>
    </Dialog.Root>
  );
}

function XIcon(props) {
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
      <path d="M18 6 6 18" />
      <path d="m6 6 12 12" />
    </svg>
  );
}