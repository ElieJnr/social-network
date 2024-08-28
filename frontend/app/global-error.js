'use client' 

export default function GlobalError({ error, reset }) {
  console.log(error)
  return (
    <html>
      <body>
        <h2>Something went wrong!</h2>
        <p>{error}</p>
        <button className="bg-red-200" onClick={() => reset()}>Try again</button>
      </body>
    </html>
  )
}