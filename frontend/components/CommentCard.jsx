"use client";

const { AvatarFallback, AvatarImage, Avatar } = require("./ui/avatar");
import { fetchPostComments } from '@/app/actions/post';
const { useState } = require('react');
const { Textarea } = require("./ui/textarea");
const { Button } = require("./ui/button");
import { mutate } from 'swr';

export default function CommentCard({ postId, commentData }) {
    console.log('Received commentData:', commentData);

    const comments = commentData ? commentData.map(comment => ({
        name: comment.Author.Firstname + ' ' + comment.Author.Lastname,
        timeAgo: comment.Formated_date,
        text: comment.Content,
        hasImage: !!comment.Image_url,
        imageUrl: comment.Image_url ? `/uploads/${comment.Image_url}` : null,
    })) : [];

    console.log('Formatted comments:', comments);
    
    return (
        <div className="mx-auto px-4 md:px-6">
            <hr className="flex-grow" />
            <br />
            <CommentList comments={comments} />
            <CommentForm postId={postId} />
        </div>
    );
}

export function CommentList({ comments }) {
    console.log('Comments in CommentList:', comments);
    
    return (
        <div className="grid gap-6">
            {comments.map((comment, index) => (
                <Comment
                    key={index}
                    name={comment.name}
                    timeAgo={comment.timeAgo}
                    text={comment.text}
                    hasImage={comment.hasImage}
                    imageUrl={comment.imageUrl}
                />
            ))}
        </div>
    );
}

export function Comment({ name, timeAgo, text, hasImage, imageUrl }) {
    console.log('Rendering Comment:', { name, timeAgo, text, hasImage, imageUrl });

    return (
        <div className="flex items-start gap-4">
            <Avatar className="w-10 h-10 border">
                <AvatarImage src="/placeholder-user.jpg" alt={name} />
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
                        <img
                            src={imageUrl}
                            width={800}
                            height={450}
                            alt="Project preview"
                            className="rounded-lg object-cover"
                            style={{ aspectRatio: "800/450", objectFit: "cover" }}
                        />
                    )}
                </div>
            </div>
        </div>
    );
}

export function CommentForm({ postId }) {
    const [commentContent, setCommentContent] = useState('');
    const [file, setFile] = useState(null);

    const handleSubmit = async (e) => {
        e.preventDefault();
        const formData = new FormData();
        formData.append('postId', postId);
        formData.append('commentContent', commentContent);
        if (file) {
            formData.append('file', file);
        }

        try {
            const data = await fetchPostComments(formData);
            // console.log('Comment created:', data);
            mutate('http://localhost:8080/posts');

            setCommentContent('');
            setFile(null);
        } catch (error) {
            console.error('Error creating comment:', error);
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
    )
}