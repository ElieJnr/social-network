'use client';

export default function GlobalError({ error, reset }) {
  return (
    <html className="h-full">
      <body className="flex h-full items-center justify-center bg-background text-foreground">
        <div className="max-w-md p-8 bg-card text-card-foreground rounded-lg shadow-lg">
          <h2 className="text-3xl font-bold mb-4">Oops! Something went wrong</h2>
          <p className="text-lg text-muted-foreground mb-6">{error.message}</p>
          <button
            onClick={() => reset()}
            className="px-4 py-2 bg-primary text-primary-foreground rounded hover:bg-primary/80"
          >
            Try Again
          </button>
        </div>
      </body>
    </html>
  );
}
