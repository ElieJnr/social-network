"use client";
import React, { useState } from 'react';

export default function CreatePostForm() {
  const [thread, setThread] = useState('');
  const [privacy, setPrivacy] = useState('public');
  const [file, setFile] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData();
    formData.append('thread', thread);
    formData.append('privacy', privacy);
    if (file) {
      formData.append('file', file);
    }

    try {
      const response = await fetch('http://localhost:8080/post/create', {
        method: 'POST',
        body: formData,
        credentials: 'include',
      });

      const data = await response.json();
      console.log('Post created:', data);
    } catch (error) {
      console.error('Error creating post:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label htmlFor="thread">Contenu:</label>
        <textarea
          id="thread"
          name="thread"
          value={thread}
          onChange={(e) => setThread(e.target.value)}
          required
        />
      </div>
      <div>
        <label>Confidentialité:</label>
        <div>
          <input
            type="radio"
            id="public"
            name="privacy"
            value="public"
            checked={privacy === 'public'}
            onChange={(e) => setPrivacy(e.target.value)}
          />
          <label htmlFor="public">Public</label>
        </div>
        <div>
          <input
            type="radio"
            id="private"
            name="privacy"
            value="private"
            checked={privacy === 'private'}
            onChange={(e) => setPrivacy(e.target.value)}
          />
          <label htmlFor="private">Privé</label>
        </div>
        <div>
          <input
            type="radio"
            id="almost-private"
            name="privacy"
            value="almost-private"
            checked={privacy === 'almost-private'}
            onChange={(e) => setPrivacy(e.target.value)}
          />
          <label htmlFor="almost-private">Presque privé</label>
        </div>
      </div>
      <div>
        <label htmlFor="file">Image:</label>
        <input
          type="file"
          id="file"
          name="file"
          onChange={(e) => setFile(e.target.files[0])}
        />
      </div>
      <button type="submit">Créer le post</button>
    </form>
  );
}

// app/post/createPostForm.js
// import { useState } from 'react';
// app/post/createPost.js


// export default function CreatePost() {
//   const [title, setTitle] = useState('');
//   const [content, setContent] = useState('');

//   const handleSubmit = async (event) => {
//     event.preventDefault();

//     const postData = {
//       title,
//       content,
//     };

//     try {
//       const response = await fetch('http://localhost:8080/post/create', {
//         method: 'POST',
//         headers: {
//           'Content-Type': 'application/json',
//         },
//         body: JSON.stringify(postData),
//       });

//       if (response.ok) {
//         alert('Post created successfully!');
//         setTitle('');
//         setContent('');
//       } else {
//         alert('Failed to create post.');
//       }
//     } catch (error) {
//       console.error('Error:', error);
//       alert('An error occurred while creating the post.');
//     }
//   };

//   return (
//     <form onSubmit={handleSubmit}>
//       <div>
//         <label>Title:</label>
//         <input
//           type="text"
//           value={title}
//           onChange={(e) => setTitle(e.target.value)}
//           required
//         />
//       </div>
//       <div>
//         <label>Content:</label>
//         <textarea
//           value={content}
//           onChange={(e) => setContent(e.target.value)}
//           required
//         ></textarea>
//       </div>
//       <button type="submit">Create Post</button>
//     </form>
//   );
// }
