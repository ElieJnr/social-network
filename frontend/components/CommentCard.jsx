"use client";

import { useState } from 'react';
import useSWR, { mutate } from 'swr';
import { fetchPostCreateComments, fetchAllPostComments } from '@/app/actions/post';
import { AvatarFallback, AvatarImage, Avatar } from "./ui/avatar";
import { useToast } from "@/components/ui/use-toast";
import { Textarea } from "./ui/textarea";
import { Button } from "./ui/button";
import Image from 'next/image';

export default function CommentCard({ postId }) {
    const { data: comments = [], mutate: mutateComments } = useSWR(
        `http://localhost:8080/comments?postId=${postId}`,
        () => fetchAllPostComments(postId)
    );

    return (
        <div className="mx-auto px-4 md:px-6">
            <hr className="flex-grow" />
            <br />
            <CommentList comments={comments} />
            <CommentForm postId={postId} />
        </div>
    );
}

function CommentList({ comments }) {
    return (
        <div className="grid gap-6">
            {comments.map((comment, index) => (
                <Comment
                    key={index}
                    name={`${comment.Author.Firstname} ${comment.Author.Lastname}`}
                    avatar={comment.Author.Avatar}
                    timeAgo={comment.Formated_date}
                    text={comment.Content}
                    hasImage={comment.HasImage}
                    imageUrl={comment.Image_url ? `/uploads/${comment.Image_url}` : null}
                />
            ))}
        </div>
    );
}

function Comment({ name, avatar, timeAgo, text, hasImage, imageUrl }) {
    return (
        <div className="flex items-start gap-4">
            <Avatar className="w-10 h-10 border">
                <AvatarImage src={`/uploads/${avatar}`|| "/placeholder-user.jpg"}alt={name} />
                <AvatarFallback>{name.charAt(0)}</AvatarFallback>
            </Avatar>
            <div className="grid gap-1.5">
                <div className="flex items-center gap-2">
                    <div className="font-semibold">{name}</div>
                    <span className="text-muted-foreground">•</span>
                    <div className="text-xs text-muted-foreground">{timeAgo}</div>
                </div>
                <div className="text-sm leading-loose text-muted-foreground">
                    <p>{text}</p>
                    {hasImage && (
                        <Image
                            src={imageUrl}
                            width={800}
                            height={450}
                            alt="Comment preview"
                            className="rounded-lg object-cover"
                            style={{ aspectRatio: "800/450", objectFit: "cover" }}
                        />
                    )}
                </div>
            </div>
        </div>
    );
}

function CommentForm({ postId }) {
    const [commentContent, setCommentContent] = useState('');
    const [file, setFile] = useState(null);
    const { toast } = useToast();

    const handleSubmit = async (e) => {
        e.preventDefault();
        const formData = new FormData();
        formData.append('postId', postId);
        formData.append('commentContent', commentContent);
        if (file) {
            formData.append('file', file);
        }

        if (!checkComment(commentContent, file, toast)) {
            return;
        }

        try {
            await fetchPostCreateComments(formData);

            setCommentContent('');
            setFile(null);

            mutate(`http://localhost:8080/comments?postId=${postId}`);
        } catch (error) {
            toast({
                title: "Error creating comment",
                description: error,
            });
        }
    }

    const handleFileChange = (e) => {
        setFile(e.target.files[0]);
    }

    return (
        <div className="grid gap-4">
            <h2 className="text-xl font-semibold">Add a comment</h2>
            <div className="w-full">
                <form onSubmit={handleSubmit}>
                    <div className="flex items-start gap-4">
                        <div className="flex-1">
                            <Textarea
                                placeholder="Write a new comment..."
                                value={commentContent}
                                onChange={(e) => setCommentContent(e.target.value)}
                                required
                                className="resize-none p-4 rounded-lg border border-neutral-300 focus:border-primary focus:ring-primary"
                            />
                            <div className="flex justify-between mt-2">
                                <div className="flex items-center">
                                    <label className="bg-primary text-primary-foreground rounded-lg px-4 py-2 hover:bg-primary/90 transition-colors cursor-pointer">
                                        <ImageIcon className="w-5 h-5" />
                                        {file && <span className="text-sm">{file.name}</span>}
                                        <input
                                            type="file"
                                            accept="image/*"
                                            onChange={handleFileChange}
                                            className="hidden"
                                        />
                                    </label>
                                    <Button type="submit" className="ml-2">Comment</Button>
                                </div>
                            </div>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    );
}

function ImageIcon(props) {
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
            <rect width="18" height="18" x="3" y="3" rx="2" ry="2" />
            <circle cx="9" cy="9" r="2" />
            <path d="m21 15-3.086-3.086a2 2 0 0 0-2.828 0L6 21" />
        </svg>
    );
}

function checkComment(commentContent, file, toast) {
    if (!commentContent) {
        toast({
            title: "Error creating Comment",
            description: "Thread must not be empty and must not exceed 500 characters.",
        });
        return false;
    }
    if (file) {
        if (!file.type.includes('image')) {
            toast({
                title: "Error creating Comment",
                description: "File must be an image.",
            });
            return false;
        }

        const validExtensions = ["jpg", "jpeg", "png", "gif"];
        const fileExtension = file.name.split('.').pop().toLowerCase();
        if (!validExtensions.includes(fileExtension)) {
            toast({
                title: "Error creating Comment",
                description: "File must be in jpg, png, or gif format.",
            });
            return false;
        }


        const maxSizeMB = 5;
        const maxSizeBytes = maxSizeMB * 1024 * 1024;
        if (file.size > maxSizeBytes) {
            toast({
                title: "Error creating Comment",
                description: `File size must not exceed ${maxSizeMB} MB.`,
            });
            return false;
        }
    }
    return true;
}